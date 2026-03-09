package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"

	"github.com/gorilla/mux"
	"github.com/lloydchang/backstage-temporal/backend/activities"
	"github.com/lloydchang/backstage-temporal/backend/emulators"
	"github.com/lloydchang/backstage-temporal/backend/types"
	"github.com/lloydchang/backstage-temporal/backend/workflows"
)

// Simplified V2 implementation without conflicts
func main() {
	// Create Temporal client
	c, err := client.Dial(client.Options{
		HostPort: "localhost:7233",
	})
	if err != nil {
		log.Fatalf("Failed to create Temporal client: %v", err)
	}
	defer c.Close()

	// Create worker
	w := worker.New(c, "ai-agent-task-queue-v2", worker.Options{})

	// Register workflows
	w.RegisterWorkflow(EnhancedComplianceCheckWorkflow)
	w.RegisterWorkflow(workflows.AIOrchestrationWorkflow)
	w.RegisterWorkflow(workflows.HumanInTheLoopWorkflow)
	w.RegisterWorkflow(workflows.MultiAgentCollaborationWorkflow)

	// Register activities
	w.RegisterActivity(activities.DiscoverInfrastructureActivity)
	w.RegisterActivity(activities.SecurityAgentActivity)
	w.RegisterActivity(activities.ComplianceAgentActivity)
	w.RegisterActivity(activities.CostOptimizationAgentActivity)
	w.RegisterActivity(activities.AggregateAgentResultsActivity)
	w.RegisterActivity(activities.HumanReviewActivity)
	w.RegisterActivity(activities.GenerateComplianceReportActivity)

	// Start worker
	err = w.Start()
	if err != nil {
		log.Fatalf("Failed to start worker: %v", err)
	}
	defer w.Stop()

	log.Println("Worker started on task queue: ai-agent-task-queue-v2")

	// Setup HTTP server
	router := mux.NewRouter()
	
	// Enhanced CORS middleware
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
			w.Header().Set("Access-Control-Max-Age", "86400")
			
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}
			
			next.ServeHTTP(w, r)
		})
	})

	// Enhanced health check
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		health := map[string]interface{}{
			"status":    "healthy",
			"timestamp": time.Now(),
			"version":   "v2",
			"uptime":    time.Since(time.Now()),
		}
		
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(health)
	}).Methods("GET")

	// Enhanced workflow endpoints
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
			TaskQueue: "ai-agent-task-queue-v2",
		}
		
		we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, 
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
				TaskQueue: "ai-agent-task-queue-v2",
			}
			
			we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, 
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

	// Enhanced workflow status
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
			"workflowId": resp.WorkflowExecutionInfo.Execution.WorkflowId,
			"runId":      resp.WorkflowExecutionInfo.Execution.RunId,
			"status":     resp.WorkflowExecutionInfo.Status.String(),
			"startTime":  resp.WorkflowExecutionInfo.StartTime,
			"closeTime":  resp.WorkflowExecutionInfo.CloseTime,
		}
		
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(status)
	}).Methods("GET")

	// Enhanced signal workflow
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

	// Infrastructure emulator endpoints
	emulator := emulators.GetGlobalEmulator()
	
	router.HandleFunc("/emulator/resources", func(w http.ResponseWriter, r *http.Request) {
		resourceType := r.URL.Query().Get("type")
		resources, err := emulator.ListResources(context.Background(), resourceType)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resources)
	}).Methods("GET")

	router.HandleFunc("/emulator/resources/{id}", func(w http.ResponseWriter, r *http.Request) {
		resourceID := mux.Vars(r)["id"]
		resource, err := emulator.GetResource(context.Background(), resourceID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resource)
	}).Methods("GET")

	router.HandleFunc("/emulator/resources/{id}/security-posture", func(w http.ResponseWriter, r *http.Request) {
		resourceID := mux.Vars(r)["id"]
		posture, err := emulator.GetSecurityPosture(context.Background(), resourceID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(posture)
	}).Methods("GET")

	router.HandleFunc("/emulator/resources/{id}/compliance", func(w http.ResponseWriter, r *http.Request) {
		resourceID := mux.Vars(r)["id"]
		standards := []string{"SOC2", "GDPR", "HIPAA"}
		status, err := emulator.GetComplianceStatus(context.Background(), resourceID, standards)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(status)
	}).Methods("GET")

	// Legacy endpoints for backward compatibility
	router.HandleFunc("/workflow/start", func(w http.ResponseWriter, r *http.Request) {
		workflowOptions := client.StartWorkflowOptions{
			ID:        fmt.Sprintf("legacy-workflow-%d", time.Now().Unix()),
			TaskQueue: "ai-agent-task-queue-v2",
		}
		
		we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, 
			EnhancedComplianceCheckWorkflow, "legacy-request")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
		w.Write([]byte(we.GetID()))
	}).Methods("POST")

	// Start HTTP server
	server := &http.Server{
		Addr:         ":8081",
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Printf("Enhanced HTTP server starting on port 8081")
	log.Fatal(server.ListenAndServe())
}

// Enhanced compliance check workflow with better error handling
func EnhancedComplianceCheckWorkflow(ctx workflow.Context, data string) (string, error) {
	logger := workflow.GetLogger(ctx)
	logger.Info("Starting Enhanced Compliance Check Workflow", "data", data)
	
	// Enhanced activity options
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
	
	// Stage 1: Data discovery
	var discoveryResult struct {
		ResourceID string `json:"resourceId"`
		Status     string `json:"status"`
	}
	
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
		}
	}
	
	// Stage 3: Aggregate results
	var aggregatedResult types.AggregatedResult
	err = workflow.ExecuteActivity(ctx, activities.AggregateAgentResultsActivity, agentResults).Get(ctx, &aggregatedResult)
	if err != nil {
		logger.Error("Result aggregation failed", "error", err)
		return "", fmt.Errorf("aggregation failed: %w", err)
	}
	
	// Stage 4: Human review if needed
	if aggregatedResult.RequiresHumanReview {
		logger.Info("Human review required", "confidence", aggregatedResult.ConfidenceScore)
		
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
	
	logger.Info("Enhanced Compliance Check Workflow completed", "reportID", report.ID, "score", report.Score)
	
	return fmt.Sprintf("Compliance check completed. Report ID: %s, Score: %.2f", report.ID, report.Score), nil
}
