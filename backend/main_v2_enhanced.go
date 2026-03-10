package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"

	"github.com/gorilla/mux"
	"github.com/lloydchang/backstage-temporal/backend/activities"
	"github.com/lloydchang/backstage-temporal/backend/config"
	"github.com/lloydchang/backstage-temporal/backend/emulators"
	"github.com/lloydchang/backstage-temporal/backend/humanloop"
	"github.com/lloydchang/backstage-temporal/backend/mcp"
	"github.com/lloydchang/backstage-temporal/backend/monitoring"
	"github.com/lloydchang/backstage-temporal/backend/performance"
	"github.com/lloydchang/backstage-temporal/backend/security"
	"github.com/lloydchang/backstage-temporal/backend/types"
	"github.com/lloydchang/backstage-temporal/backend/workflows"
)

// Enhanced configuration with environment variables
type AppConfig struct {
	TemporalHost      string        `json:"temporalHost"`
	ServerPort        string        `json:"serverPort"`
	TaskQueue         string        `json:"taskQueue"`
	EnableMetrics     bool          `json:"enableMetrics"`
	EnableSecurity    bool          `json:"enableSecurity"`
	MaxConcurrent     int           `json:"maxConcurrent"`
	ShutdownTimeout   time.Duration `json:"shutdownTimeout"`
	LogLevel          string        `json:"logLevel"`
	
	// MCP Configuration
	EnableMCP         bool          `json:"enableMCP"`
	MCPServerName     string        `json:"mcpServerName"`
	MCPServerVersion  string        `json:"mcpServerVersion"`
	MCPTransport      string        `json:"mcpTransport"`
	MCPPort           string        `json:"mcpPort"`
	MCPEnableAuth     bool          `json:"mcpEnableAuth"`
	MCPAPIKey         string        `json:"mcpApiKey"`
	MCPAllowedTools   []string      `json:"mcpAllowedTools"`
	MCPAllowedResources []string    `json:"mcpAllowedResources"`
}

// Global application state
type AppState struct {
	config           *AppConfig
	temporalClient   client.Client
	worker           worker.Worker
	metricsCollector monitoring.MetricsCollector
	securityManager  security.SecurityManager
	mcpServer        *mcp.MCPServer
	shutdownOnce     sync.Once
}

var appState *AppState

// Enhanced CORS middleware with security headers
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
		w.Header().Set("Access-Control-Max-Age", "86400")
		
		// Security headers
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		
		next.ServeHTTP(w, r)
	})
}

// Enhanced logging middleware
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		logged := &responseLogger{ResponseWriter: w, statusCode: http.StatusOK}
		
		next.ServeHTTP(logged, r)
		
		duration := time.Since(start)
		log.Printf("Request: %s %s - Status: %d - Duration: %v", 
			r.Method, r.URL.Path, logged.statusCode, duration)
	})
}

type responseLogger struct {
	http.ResponseWriter
	statusCode int
}

func (rl *responseLogger) WriteHeader(code int) {
	rl.statusCode = code
	rl.ResponseWriter.WriteHeader(code)
}

// Enhanced workflow with improved error handling and metrics
func EnhancedComplianceCheckWorkflow(ctx workflow.Context, data string) (string, error) {
	logger := workflow.GetLogger(ctx)
	logger.Info("Starting Enhanced Compliance Check Workflow", "data", data)
	
	// Track workflow metrics
	startTime := workflow.Now(ctx)
	
	// Enhanced activity options with circuit breaker pattern
	activityOptions := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute * 10,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    time.Second * 2,
			BackoffCoefficient: 2.0,
			MaximumInterval:    time.Minute * 2,
			MaximumAttempts:    3,
			NonRetryableErrorTypes: []string{"ValidationError", "AuthenticationError"},
		},
		HeartbeatTimeout: time.Minute * 1,
	}
	ctx = workflow.WithActivityOptions(ctx, activityOptions)
	
	// Stage 1: Data discovery with validation
	var discoveryResult types.InfrastructureResult
	err := workflow.ExecuteActivity(ctx, activities.DiscoverInfrastructureActivity, data).Get(ctx, &discoveryResult)
	if err != nil {
		logger.Error("Data discovery failed", "error", err)
		return "", fmt.Errorf("discovery failed: %w", err)
	}
	
	// Stage 2: Parallel agent execution
	futures := make([]workflow.Future, 0, 3)
	
	// Security Agent
	var securityResult types.AgentResult
	futures = append(futures, workflow.ExecuteActivity(ctx, activities.SecurityAgentActivity, discoveryResult.ResourceID))
	
	// Compliance Agent
	var complianceResult types.AgentResult
	futures = append(futures, workflow.ExecuteActivity(ctx, activities.ComplianceAgentActivity, discoveryResult.ResourceID))
	
	// Cost Optimization Agent
	var costResult types.AgentResult
	futures = append(futures, workflow.ExecuteActivity(ctx, activities.CostOptimizationAgentActivity, discoveryResult.ResourceID))
	
	// Wait for all agents to complete
	agentResults := make([]types.AgentResult, 0, 3)
	for i, future := range futures {
		var result types.AgentResult
		if i == 0 {
			err = future.Get(ctx, &securityResult)
			agentResults = append(agentResults, securityResult)
		} else if i == 1 {
			err = future.Get(ctx, &complianceResult)
			agentResults = append(agentResults, complianceResult)
		} else {
			err = future.Get(ctx, &costResult)
			agentResults = append(agentResults, costResult)
		}
		
		if err != nil {
			logger.Warn("Agent execution failed", "agentIndex", i, "error", err)
			// Continue with other agents instead of failing completely
		}
	}
	
	// Stage 3: Aggregate results with confidence scoring
	var aggregatedResult types.AggregatedResult
	err = workflow.ExecuteActivity(ctx, activities.AggregateAgentResultsActivity, agentResults).Get(ctx, &aggregatedResult)
	if err != nil {
		logger.Error("Result aggregation failed", "error", err)
		return "", fmt.Errorf("aggregation failed: %w", err)
	}
	
	// Stage 4: Human review if needed
	if aggregatedResult.RequiresHumanReview {
		logger.Info("Human review required", "confidence", aggregatedResult.ConfidenceScore)
		
		// Create human task
		task := types.HumanTask{
			ID:          fmt.Sprintf("task-%s", workflow.GetInfo(ctx).WorkflowExecution.ID),
			Title:       "Compliance Review Required",
			Description: fmt.Sprintf("Review compliance findings for %s", discoveryResult.ResourceID),
			Priority:    "high",
			AssignedTo:  "compliance-team",
			DueAt:       workflow.Now(ctx).Add(24 * time.Hour),
			Status:      types.HumanTaskStatus{State: "pending", UpdatedAt: workflow.Now(ctx)},
			Data:        map[string]interface{}{"aggregatedResult": aggregatedResult},
		}
		
		var humanResult types.HumanTaskResult
		err = workflow.ExecuteActivity(ctx, activities.HumanReviewActivity, task).Get(ctx, &humanResult)
		if err != nil {
			logger.Error("Human review failed", "error", err)
			return "", fmt.Errorf("human review failed: %w", err)
		}
		
		// Update aggregated result with human feedback
		if humanResult.Approved {
			aggregatedResult.Summary += " [Human Approved]"
		} else {
			return fmt.Sprintf("Rejected: %s", humanResult.Decision), nil
		}
	}
	
	// Stage 5: Generate final report
	var report types.ComplianceReport
	err = workflow.ExecuteActivity(ctx, activities.GenerateComplianceReportActivity, aggregatedResult).Get(ctx, &report)
	if err != nil {
		logger.Error("Report generation failed", "error", err)
		return "", fmt.Errorf("report generation failed: %w", err)
	}
	
	// Log workflow completion
	duration := workflow.Now(ctx).Sub(startTime)
	logger.Info("Enhanced Compliance Check Workflow completed", 
		"duration", duration, 
		"reportID", report.ID,
		"overallScore", report.Score)
	
	return fmt.Sprintf("Compliance check completed. Report ID: %s, Score: %.2f", report.ID, report.Score), nil
}

// Graceful shutdown handler
func (app *AppState) shutdown() {
	app.shutdownOnce.Do(func() {
		log.Println("Initiating graceful shutdown...")
		
		// Stop MCP server if running
		if app.mcpServer != nil {
			if err := app.mcpServer.Stop(); err != nil {
				log.Printf("Error stopping MCP server: %v", err)
			} else {
				log.Println("MCP server stopped")
			}
		}
		
		if app.worker != nil {
			app.worker.Stop()
			log.Println("Temporal worker stopped")
		}
		
		if app.temporalClient != nil {
			app.temporalClient.Close()
			log.Println("Temporal client closed")
		}
		
		log.Println("Graceful shutdown completed")
	})
}

// Enhanced health check endpoint
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	health := map[string]interface{}{
		"status":      "healthy",
		"timestamp":   time.Now(),
		"version":     "v2",
		"uptime":      time.Since(time.Now()), // This would be tracked in real implementation
		"goroutines":  runtime.NumGoroutine(),
		"memory":      map[string]interface{}{
			"alloc":      runtime.MemStats{}.Alloc,
			"total_alloc": runtime.MemStats{}.TotalAlloc,
			"sys":        runtime.MemStats{}.Sys,
		},
	}
	
	if appState != nil && appState.metricsCollector != nil {
		health["metrics"] = appState.metricsCollector.GetSummary()
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(health)
}

// Enhanced workflow endpoints with better error handling
func setupWorkflowRoutes(router *mux.Router, client client.Client) {
	// Enhanced compliance workflow
	router.HandleFunc("/workflow/start-enhanced-compliance", func(w http.ResponseWriter, r *http.Request) {
		var request struct {
			TargetResource string            `json:"targetResource"`
			Parameters    map[string]string `json:"parameters"`
			Priority      string            `json:"priority"`
		}
		
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		
		workflowOptions := client.StartWorkflowOptions{
			ID:        fmt.Sprintf("enhanced-compliance-%s-%d", request.TargetResource, time.Now().Unix()),
			TaskQueue: appState.config.TaskQueue,
		}
		
		we, err := client.ExecuteWorkflow(context.Background(), workflowOptions, 
			EnhancedComplianceCheckWorkflow, request.TargetResource)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to start workflow: %v", err), http.StatusInternalServerError)
			return
		}
		
		response := map[string]interface{}{
			"workflowId": we.GetID(),
			"runId":      we.GetRunID(),
			"status":     "started",
			"timestamp":  time.Now(),
		}
		
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}).Methods("POST")
	
	// Batch workflow execution
	router.HandleFunc("/workflow/start-batch", func(w http.ResponseWriter, r *http.Request) {
		var requests []struct {
			TargetResource string            `json:"targetResource"`
			Parameters    map[string]string `json:"parameters"`
		}
		
		if err := json.NewDecoder(r.Body).Decode(&requests); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		
		results := make([]map[string]interface{}, 0, len(requests))
		
		for _, req := range requests {
			workflowOptions := client.StartWorkflowOptions{
				ID:        fmt.Sprintf("batch-compliance-%s-%d", req.TargetResource, time.Now().UnixNano()),
				TaskQueue: appState.config.TaskQueue,
			}
			
			we, err := client.ExecuteWorkflow(context.Background(), workflowOptions, 
				EnhancedComplianceCheckWorkflow, req.TargetResource)
			
			result := map[string]interface{}{
				"targetResource": req.TargetResource,
				"workflowId":     we.GetID(),
				"status":         "started",
			}
			
			if err != nil {
				result["error"] = err.Error()
				result["status"] = "failed"
			}
			
			results = append(results, result)
		}
		
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"batchId":    fmt.Sprintf("batch-%d", time.Now().Unix()),
			"workflows":  results,
			"totalCount": len(requests),
		})
	}).Methods("POST")
}

func main() {
	// Load configuration with environment variable support
	config := &AppConfig{
		TemporalHost:    getEnv("TEMPORAL_HOST", "localhost:7233"),
		ServerPort:      getEnv("SERVER_PORT", "8081"),
		TaskQueue:       getEnv("TASK_QUEUE", "ai-agent-task-queue-v2"),
		EnableMetrics:   getEnv("ENABLE_METRICS", "true") == "true",
		EnableSecurity:  getEnv("ENABLE_SECURITY", "true") == "true",
		MaxConcurrent:   parseInt(getEnv("MAX_CONCURRENT", "10")),
		ShutdownTimeout: parseDuration(getEnv("SHUTDOWN_TIMEOUT", "30s")),
		LogLevel:        getEnv("LOG_LEVEL", "info"),
		
		// MCP Configuration
		EnableMCP:           getEnv("ENABLE_MCP", "true") == "true",
		MCPServerName:       getEnv("MCP_SERVER_NAME", "Temporal AI Agents"),
		MCPServerVersion:    getEnv("MCP_SERVER_VERSION", "1.0.0"),
		MCPTransport:        getEnv("MCP_TRANSPORT", "stdio"),
		MCPPort:             getEnv("MCP_PORT", "8082"),
		MCPEnableAuth:       getEnv("MCP_ENABLE_AUTH", "false") == "true",
		MCPAPIKey:           getEnv("MCP_API_KEY", ""),
		MCPAllowedTools:     getEnvSlice("MCP_ALLOWED_TOOLS", []string{}),
		MCPAllowedResources: getEnvSlice("MCP_ALLOWED_RESOURCES", []string{}),
	}
	
	// Initialize application state
	appState = &AppState{config: config}
	
	log.Printf("Starting AI Agent Temporal Service V2")
	log.Printf("Configuration: %+v", config)
	
	// Initialize Temporal client
	c, err := client.Dial(client.Options{
		HostPort: config.TemporalHost,
	})
	if err != nil {
		log.Fatalf("Failed to create Temporal client: %v", err)
	}
	appState.temporalClient = c
	
	// Initialize worker with enhanced options
	workerOptions := worker.Options{
		MaxConcurrentActivityExecutionSize: config.MaxConcurrent,
	}
	
	w := worker.New(c, config.TaskQueue, workerOptions)
	appState.worker = w
	
	// Register enhanced workflows
	w.RegisterWorkflow(EnhancedComplianceCheckWorkflow)
	w.RegisterWorkflow(workflows.AIOrchestrationWorkflowV2)
	w.RegisterWorkflow(humanloop.EnhancedHumanInTheLoopWorkflow)
	w.RegisterWorkflow(performance.OptimizedWorkflow)
	
	// Register enhanced activities
	w.RegisterActivity(activities.DiscoverInfrastructureActivity)
	w.RegisterActivity(activities.SecurityAgentActivity)
	w.RegisterActivity(activities.ComplianceAgentActivity)
	w.RegisterActivity(activities.CostOptimizationAgentActivity)
	w.RegisterActivity(activities.AggregateAgentResultsActivity)
	w.RegisterActivity(activities.HumanReviewActivity)
	w.RegisterActivity(activities.GenerateComplianceReportActivity)
	
	// Initialize monitoring if enabled
	if config.EnableMetrics {
		appState.metricsCollector = monitoring.NewMetricsCollector()
		go appState.metricsCollector.Start(context.Background())
	}
	
	// Initialize security if enabled
	if config.EnableSecurity {
		appState.securityManager = security.NewSecurityManager()
	}
	
	// Initialize MCP server if enabled
	if config.EnableMCP {
		mcpConfig := &mcp.MCPConfig{
			ServerName:        config.MCPServerName,
			ServerVersion:     config.MCPServerVersion,
			TransportType:     config.MCPTransport,
			Port:              config.MCPPort,
			EnableAuth:        config.MCPEnableAuth,
			APIKey:            config.MCPAPIKey,
			AllowedTools:      config.MCPAllowedTools,
			AllowedResources:  config.MCPAllowedResources,
			LogLevel:          config.LogLevel,
		}
		
		appState.mcpServer = mcp.NewMCPServer(c, mcpConfig)
		
		// Start MCP server in goroutine
		go func() {
			if err := appState.mcpServer.Start(context.Background()); err != nil {
				log.Printf("MCP server failed to start: %v", err)
			}
		}()
		
		log.Printf("MCP server configured with %s transport on port %s", config.MCPTransport, config.MCPPort)
	}
	
	// Start worker
	err = w.Start()
	if err != nil {
		log.Fatalf("Failed to start worker: %v", err)
	}
	defer w.Stop()
	
	log.Printf("Worker started on task queue: %s", config.TaskQueue)
	
	// Setup HTTP server with enhanced middleware
	router := mux.NewRouter()
	router.Use(corsMiddleware)
	router.Use(loggingMiddleware)
	
	// Health check endpoint
	router.HandleFunc("/health", healthCheckHandler).Methods("GET")
	
	// Setup workflow routes
	setupWorkflowRoutes(router, c)
	
	// Legacy endpoints for backward compatibility
	router.HandleFunc("/workflow/start", func(w http.ResponseWriter, r *http.Request) {
		workflowOptions := client.StartWorkflowOptions{
			ID:        fmt.Sprintf("legacy-workflow-%d", time.Now().Unix()),
			TaskQueue: config.TaskQueue,
		}
		
		we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, 
			EnhancedComplianceCheckWorkflow, "legacy-request")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
		w.Write([]byte(we.GetID()))
	}).Methods("POST")
	
	// Workflow status endpoint
	router.HandleFunc("/workflow/status", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "missing id", http.StatusBadRequest)
			return
		}
		
		resp, err := c.DescribeWorkflowExecution(context.Background(), id, "")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
		status := map[string]interface{}{
			"workflowId": resp.WorkflowExecutionInfo.WorkflowExecution.ID,
			"runId":      resp.WorkflowExecutionInfo.WorkflowExecution.RunID,
			"status":     resp.WorkflowExecutionInfo.Status.String(),
			"startTime":  resp.WorkflowExecutionInfo.StartTime,
			"closeTime":  resp.WorkflowExecutionInfo.CloseTime,
		}
		
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(status)
	}).Methods("GET")
	
	// Signal workflow endpoint
	router.HandleFunc("/workflow/signal/{workflowId}", func(w http.ResponseWriter, r *http.Request) {
		workflowId := mux.Vars(r)["workflowId"]
		
		var signalReq struct {
			Signal string `json:"signal"`
			Value  string `json:"value"`
		}
		
		if err := json.NewDecoder(r.Body).Decode(&signalReq); err != nil {
			http.Error(w, "invalid JSON", http.StatusBadRequest)
			return
		}
		
		err := c.SignalWorkflow(context.Background(), workflowId, "", signalReq.Signal, signalReq.Value)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "signal sent",
			"signal":  signalReq.Signal,
			"value":   signalReq.Value,
		})
	}).Methods("POST")
	
	// Metrics endpoint
	if config.EnableMetrics {
		router.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
			if appState.metricsCollector != nil {
				metrics := appState.metricsCollector.GetAll()
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(metrics)
			} else {
				http.Error(w, "Metrics not enabled", http.StatusServiceUnavailable)
			}
		}).Methods("GET")
	}
	
	// Start HTTP server
	server := &http.Server{
		Addr:         ":" + config.ServerPort,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	
	log.Printf("HTTP server starting on port %s", config.ServerPort)
	
	// Graceful shutdown handling
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server failed: %v", err)
		}
	}()
	
	// Wait for interrupt signal
	sigChan := make(chan os.Signal, 1)
	// signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	
	log.Println("Shutdown signal received")
	
	// Graceful shutdown with timeout
	shutdownCtx, cancel := context.WithTimeout(context.Background(), config.ShutdownTimeout)
	defer cancel()
	
	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Printf("Server shutdown error: %v", err)
	}
	
	appState.shutdown()
	log.Println("Service stopped")
}

// Helper functions
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func parseInt(s string) int {
	if i, err := strconv.Atoi(s); err == nil {
		return i
	}
	return 10 // default
}

func parseDuration(s string) time.Duration {
	if d, err := time.ParseDuration(s); err == nil {
		return d
	}
	return 30 * time.Second // default
}

func getEnvSlice(key string, defaultValue []string) []string {
	if value := os.Getenv(key); value != "" {
		return strings.Split(value, ",")
	}
	return defaultValue
}
