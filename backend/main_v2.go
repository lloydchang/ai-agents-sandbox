package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"

	"github.com/gorilla/mux"
	"github.com/lloydchang/backstage-temporal/backend/activities"
	"github.com/lloydchang/backstage-temporal/backend/humanloop"
	"github.com/lloydchang/backstage-temporal/backend/performance"
	"github.com/lloydchang/backstage-temporal/backend/types"
	"github.com/lloydchang/backstage-temporal/backend/workflows"
)

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
	w := worker.New(c, "ai-agent-task-queue", worker.Options{})

	// Register workflows
	w.RegisterWorkflow(ComplianceCheckWorkflow)
	w.RegisterWorkflow(workflows.AIOrchestrationWorkflow)
	w.RegisterWorkflow(workflows.HumanInTheLoopWorkflow)
	w.RegisterWorkflow(workflows.MultiAgentCollaborationWorkflow)

	// Register enhanced workflows
	w.RegisterWorkflow(humanloop.EnhancedHumanInTheLoopWorkflow)
	w.RegisterWorkflow(performance.OptimizedWorkflow)
	w.RegisterWorkflow(performance.PerformanceMonitoringWorkflow)

	// Register activities
	w.RegisterActivity(activities.DiscoverInfrastructureActivity)
	w.RegisterActivity(activities.SecurityAgentActivity)
	w.RegisterActivity(activities.ComplianceAgentActivity)
	w.RegisterActivity(activities.CostOptimizationAgentActivity)
	w.RegisterActivity(activities.AggregateAgentResultsActivity)
	w.RegisterActivity(activities.GenerateComplianceReportActivity)
	w.RegisterActivity(activities.HumanReviewActivity)

	// Register human loop activities
	w.RegisterActivity(humanloop.RouteTaskActivity)
	w.RegisterActivity(humanloop.SendNotificationActivity)

	// Start worker
	go func() {
		err := w.Run(worker.InterruptCh())
		if err != nil {
			log.Fatalf("Failed to start worker: %v", err)
		}
	}()

	// Create HTTP server with enhanced endpoints
	r := mux.NewRouter()

	// CORS middleware
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		})
	})

	// Enhanced workflow endpoints
	r.HandleFunc("/workflow/start", func(w http.ResponseWriter, r *http.Request) {
		var data struct {
			Data string `json:"data"`
		}
		json.NewDecoder(r.Body).Decode(&data)

		workflowOptions := client.StartWorkflowOptions{
			ID:        "compliance-check-" + time.Now().Format("20060102150405"),
			TaskQueue: "ai-agent-task-queue",
		}

		we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, ComplianceCheckWorkflow, data.Data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(map[string]string{
			"workflowId": we.GetID(),
			"runId":      we.GetRunID(),
		})
	}).Methods("POST")

	// Enhanced AI orchestration endpoint
	r.HandleFunc("/workflow/ai-orchestration/start", func(w http.ResponseWriter, r *http.Request) {
		var request types.ComplianceRequest
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		workflowOptions := client.StartWorkflowOptions{
			ID:        "ai-orchestration-" + time.Now().Format("20060102150405"),
			TaskQueue: "ai-agent-task-queue",
		}

		we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, workflows.AIOrchestrationWorkflow, request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(map[string]string{
			"workflowId": we.GetID(),
			"runId":      we.GetRunID(),
		})
	}).Methods("POST")

	// Enhanced signal endpoint
	r.HandleFunc("/workflow/signal/{workflowId}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		workflowId := vars["workflowId"]

		var signalData struct {
			Signal string `json:"signal"`
			Data   string `json:"data"`
		}
		json.NewDecoder(r.Body).Decode(&signalData)

		err := c.SignalWorkflow(context.Background(), workflowId, "", signalData.Signal, signalData.Data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(map[string]string{"status": "signal sent"})
	}).Methods("POST")

	// Enhanced status endpoint
	r.HandleFunc("/workflow/status/{workflowId}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		workflowId := vars["workflowId"]

		resp, err := c.DescribeWorkflowExecution(context.Background(), workflowId, "")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(resp)
	}).Methods("GET")

	// Enhanced human task endpoints
	r.HandleFunc("/human-tasks/pending", func(w http.ResponseWriter, r *http.Request) {
		// Mock data for pending human tasks
		tasks := []types.HumanTask{
			{
				ID:          "task-1",
				Title:       "Security Policy Violation Review",
				Description: "AI agent detected potential security policy violations in cloud infrastructure",
				Priority:    "high",
				AssignedTo:  "security-team@company.com",
				DueAt:       time.Now().Add(24 * time.Hour),
				Status: types.HumanTaskStatus{
					State:     "pending",
					UpdatedAt: time.Now(),
					UpdatedBy: "system",
					Notes:     "Awaiting human review",
				},
				Data: map[string]interface{}{
					"workflowId": "workflow-1",
					"riskLevel":  "medium",
					"findings":   []string{"Unencrypted S3 bucket", "Open security group"},
				},
			},
		}

		json.NewEncoder(w).Encode(map[string]interface{}{
			"tasks": tasks,
			"count": len(tasks),
		})
	}).Methods("GET")

	r.HandleFunc("/human-tasks/{taskId}/complete", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		taskId := vars["taskId"]

		var result types.HumanTaskResult
		if err := json.NewDecoder(r.Body).Decode(&result); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// In a real implementation, this would update the task status in a database
		result.TaskID = taskId
		result.CompletedAt = time.Now()

		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "completed",
			"result": result,
		})
	}).Methods("POST")

	// Enhanced approvals endpoint
	r.HandleFunc("/approvals/pending", func(w http.ResponseWriter, r *http.Request) {
		// Mock data for pending approvals
		approvals := []types.AggregatedResult{
			{
				OverallScore:       92.5,
				RequiresHumanReview: true,
				RiskLevel:          "Medium",
				Summary:            "AI agents detected compliance issues requiring human review",
				AgentResults: []types.AgentResult{
					{
						AgentID:        "security-agent-1",
						AgentType:      "Security",
						Status:         "completed",
						Score:          88.0,
						Findings:       []string{"Unencrypted data storage", "Weak password policy"},
						Recommendations: []string{"Enable encryption", "Strengthen password requirements"},
						ExecutedAt:     time.Now().Add(-1 * time.Hour),
					},
					{
						AgentID:        "compliance-agent-1",
						AgentType:      "Compliance",
						Status:         "completed",
						Score:          95.0,
						Findings:       []string{"Missing audit logs"},
						Recommendations: []string{"Enable comprehensive logging"},
						ExecutedAt:     time.Now().Add(-1 * time.Hour),
					},
				},
			},
		}

		json.NewEncoder(w).Encode(approvals)
	}).Methods("GET")

	r.HandleFunc("/approvals/{approvalId}/decide", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		approvalId := vars["approvalId"]

		var decision types.HumanReviewResult
		if err := json.NewDecoder(r.Body).Decode(&decision); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// In a real implementation, this would update the approval status
		decision.ReviewedAt = time.Now()

		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "decision recorded",
			"approvalId": approvalId,
			"decision": decision,
		})
	}).Methods("POST")

	// Enhanced monitoring endpoints
	r.HandleFunc("/monitoring/metrics", func(w http.ResponseWriter, r *http.Request) {
		metrics := map[string]interface{}{
			"activeWorkflows": 5,
			"completedWorkflows": 150,
			"failedWorkflows": 2,
			"averageExecutionTime": "45s",
			"successRate": 0.987,
			"timestamp": time.Now(),
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(metrics)
	}).Methods("GET")

	// Health check endpoint
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		health := map[string]interface{}{
			"status": "healthy",
			"timestamp": time.Now(),
			"version": "2.0",
			"temporalConnection": "connected",
			"features": []string{
				"AI Agent Orchestration",
				"Human-in-the-Loop",
				"Compliance Checking",
				"Real-time Monitoring",
				"Auto-approval",
			},
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(health)
	}).Methods("GET")

	log.Printf("Starting enhanced HTTP server on :8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}

// ComplianceCheckWorkflow with enhanced human-in-the-loop
func ComplianceCheckWorkflow(ctx workflow.Context, data string) (string, error) {
	logger := workflow.GetLogger(ctx)
	logger.Info("Starting enhanced ComplianceCheckWorkflow", "data", data)

	// Set activity options with retry policy
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute * 10,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    time.Second,
			BackoffCoefficient: 2.0,
			MaximumInterval:    time.Minute,
			MaximumAttempts:    5,
		},
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	// Step 1: Infrastructure discovery
	var infraResult types.InfrastructureResult
	err := workflow.ExecuteActivity(ctx, activities.DiscoverInfrastructureActivity, data).Get(ctx, &infraResult)
	if err != nil {
		return "", err
	}

	// Step 2: AI agent check
	var checkResult types.AgentResult
	err = workflow.ExecuteActivity(ctx, activities.SecurityAgentActivity, infraResult).Get(ctx, &checkResult)
	if err != nil {
		return "", err
	}

	// Step 3: Aggregate results
	var aggregatedResult types.AggregatedResult
	err = workflow.ExecuteActivity(ctx, activities.AggregateAgentResultsActivity, []types.AgentResult{checkResult}).Get(ctx, &aggregatedResult)
	if err != nil {
		return "", err
	}

	// Step 4: Enhanced human review with timeout and signals
	if aggregatedResult.RequiresHumanReview {
		logger.Info("Issues detected, requiring human review")

		// Set up signal channels for different decision types
		approvalCh := workflow.GetSignalChannel(ctx, "human-approval")
		rejectionCh := workflow.GetSignalChannel(ctx, "human-rejection")
		escalationCh := workflow.GetSignalChannel(ctx, "human-escalation")

		// Set up timer with configurable timeout
		timerCtx, cancelTimer := workflow.WithCancel(ctx)
		timeout := time.Hour * 24 // Default 24 hours
		timer := workflow.NewTimer(timerCtx, timeout)

		selector := workflow.NewSelector(ctx)
		var decision string
		var decisionReceived bool

		// Handle approval signal
		selector.AddReceive(approvalCh, func(c workflow.ReceiveChannel, more bool) {
			c.Receive(ctx, &decision)
			decision = "Approved"
			decisionReceived = true
			cancelTimer()
		})

		// Handle rejection signal
		selector.AddReceive(rejectionCh, func(c workflow.ReceiveChannel, more bool) {
			c.Receive(ctx, &decision)
			decision = "Rejected"
			decisionReceived = true
			cancelTimer()
		})

		// Handle escalation signal
		selector.AddReceive(escalationCh, func(c workflow.ReceiveChannel, more bool) {
			c.Receive(ctx, &decision)
			decision = "Escalated"
			decisionReceived = true
			cancelTimer()
		})

		// Handle timer expiration
		selector.AddFuture(timer, func(f workflow.Future) {
			if !decisionReceived {
				decision = "Timeout - Auto-rejected"
			}
		})

		// Wait for decision
		selector.Select(ctx)

		logger.Info("Human review completed", "decision", decision)
		return decision, nil
	}

	return "Auto-approved - No issues found", nil
}
