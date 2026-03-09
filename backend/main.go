package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
	
	// Import our custom packages
	"github.com/lloydchang/backstage-temporal/backend/activities"
	"github.com/lloydchang/backstage-temporal/backend/emulators"
	"github.com/lloydchang/backstage-temporal/backend/workflows"
)

// CORS middleware
func corsMiddleware(next http.Handler) http.Handler {
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
}

func HelloBackstageWorkflow(ctx workflow.Context, name string) (string, error) {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    time.Second,
			BackoffCoefficient: 2.0,
			MaximumInterval:    time.Minute,
			MaximumAttempts:    5,
		},
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	var result string
	err := workflow.ExecuteActivity(ctx, FetchDataActivity, name).Get(ctx, &result)
	if err != nil {
		return "", err
	}

	err = workflow.ExecuteActivity(ctx, ProcessDataActivity, result).Get(ctx, &result)
	if err != nil {
		return "", err
	}

	return result, nil
}

func ComplianceCheckWorkflow(ctx workflow.Context, data string) (string, error) {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    time.Second,
			BackoffCoefficient: 2.0,
			MaximumInterval:    time.Minute,
			MaximumAttempts:    5,
		},
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	var fetchedData string
	err := workflow.ExecuteActivity(ctx, FetchDataActivity, data).Get(ctx, &fetchedData)
	if err != nil {
		return "", err
	}

	var checkResult string
	err = workflow.ExecuteActivity(ctx, AgentCheckActivity, fetchedData).Get(ctx, &checkResult)
	if err != nil {
		return "", err
	}

	var aggregatedResult string
	err = workflow.ExecuteActivity(ctx, AggregateResultsActivity, []string{checkResult}).Get(ctx, &aggregatedResult)
	if err != nil {
		return "", err
	}

	// If issues found, wait for human review
	if aggregatedResult != "All Compliant" {
		// Set up signal channel for human approval
		approvalCh := workflow.GetSignalChannel(ctx, "human-approval")

		// Set up timer for 24 hours
		timerCtx, cancelTimer := workflow.WithCancel(ctx)
		timer := workflow.NewTimer(timerCtx, 24*time.Hour)

		selector := workflow.NewSelector(ctx)
		var approval string
		selector.AddReceive(approvalCh, func(c workflow.Channel, more bool) {
			c.Receive(ctx, &approval)
			cancelTimer() // Cancel timer if approval received
		})
		selector.AddFuture(timer, func(f workflow.Future) {
			// Timer expired, default to rejected
			approval = "Rejected"
		})

		// Wait for either signal or timer
		selector.Select(ctx)

		return approval, nil
	}

	return "Approved", nil
}

func FetchDataActivity(ctx context.Context, name string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("FetchDataActivity", "name", name)
	return "Fetched data for " + name, nil
}

func ProcessDataActivity(ctx context.Context, data string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("ProcessDataActivity", "data", data)
	return "Processed: " + data, nil
}

func AgentCheckActivity(ctx context.Context, data string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("AgentCheckActivity", "data", data)
	// Simulate AI agent compliance check
	// In real implementation, call Azure Foundry or local AI API
	return "Compliant", nil
}

func AggregateResultsActivity(ctx context.Context, results []string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("AggregateResultsActivity", "results", results)
	// Aggregate compliance results
	compliantCount := 0
	for _, result := range results {
		if result == "Compliant" {
			compliantCount++
		}
	}
	if compliantCount == len(results) {
		return "All Compliant", nil
	}
	return "Issues Found", nil
}

func HumanReviewActivity(ctx context.Context, aggregatedResult string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("HumanReviewActivity", "aggregatedResult", aggregatedResult)
	// Mock human review - in real implementation, pause workflow and wait for signal
	if aggregatedResult == "All Compliant" {
		return "Approved", nil
	}
	return "Rejected", nil
}

func main() {
	// Initialize the infrastructure emulator
	emulator := emulators.GetGlobalEmulator()
	log.Printf("Infrastructure emulator initialized")

	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatal("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, "ai-agent-task-queue", worker.Options{})
	
	// Register workflows
	w.RegisterWorkflow(workflows.AIOrchestrationWorkflow)
	w.RegisterWorkflow(workflows.HumanInTheLoopWorkflow)
	w.RegisterWorkflow(workflows.MultiAgentCollaborationWorkflow)
	w.RegisterWorkflow(HelloBackstageWorkflow)
	w.RegisterWorkflow(ComplianceCheckWorkflow)

	// Register activities
	w.RegisterActivity(activities.DiscoverInfrastructureActivity)
	w.RegisterActivity(activities.SecurityAgentActivity)
	w.RegisterActivity(activities.ComplianceAgentActivity)
	w.RegisterActivity(activities.CostOptimizationAgentActivity)
	w.RegisterActivity(activities.AggregateAgentResultsActivity)
	w.RegisterActivity(activities.GenerateComplianceReportActivity)
	w.RegisterActivity(activities.PrimaryAgentActivity)
	w.RegisterActivity(activities.ValidationAgentActivity)
	w.RegisterActivity(activities.BuildConsensusActivity)
	w.RegisterActivity(activities.GenerateFinalRecommendationActivity)
	w.RegisterActivity(activities.HumanReviewActivity)
	
	// Register legacy activities
	w.RegisterActivity(FetchDataActivity)
	w.RegisterActivity(ProcessDataActivity)
	w.RegisterActivity(AgentCheckActivity)
	w.RegisterActivity(AggregateResultsActivity)

	err = w.Start()
	if err != nil {
		log.Fatal("Unable to start worker", err)
	}

	// HTTP server for endpoints
	r := mux.NewRouter()
	
	// Apply CORS middleware
	r.Use(corsMiddleware)
	
	// Add explicit OPTIONS handlers for CORS preflight
	r.HandleFunc("/workflow/start", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}).Methods("OPTIONS")
	
	r.HandleFunc("/workflow/signal/{workflowId}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}).Methods("OPTIONS")
	
	r.HandleFunc("/workflow/start-ai-orchestration", func(w http.ResponseWriter, r *http.Request) {
		request := workflows.ComplianceRequest{
			TargetResource: "vm-web-server-001",
			ComplianceType: "full-scan",
			Parameters:     make(map[string]string),
			RequesterID:    "backstage-user",
			Priority:       "normal",
		}
		
		we, err := c.ExecuteWorkflow(context.Background(), client.StartWorkflowOptions{
			ID:        "ai-orchestration-" + time.Now().Format("20060102150405"),
			TaskQueue: "ai-agent-task-queue",
		}, workflows.AIOrchestrationWorkflow, request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte(we.GetID()))
	}).Methods("POST")

	r.HandleFunc("/workflow/start-human-in-loop", func(w http.ResponseWriter, r *http.Request) {
		task := workflows.HumanTask{
			ID:          "task-" + time.Now().Format("20060102150405"),
			Title:       "Security Review Required",
			Description: "Please review the security compliance findings",
			Priority:    "high",
			AssignedTo:  "security-team",
			DueAt:       time.Now().Add(24 * time.Hour),
			Status:      workflows.HumanTaskStatus{State: "pending", UpdatedAt: time.Now()},
			Data:        make(map[string]interface{}),
		}
		
		we, err := c.ExecuteWorkflow(context.Background(), client.StartWorkflowOptions{
			ID:        "human-in-loop-" + time.Now().Format("20060102150405"),
			TaskQueue: "ai-agent-task-queue",
		}, workflows.HumanInTheLoopWorkflow, task)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte(we.GetID()))
	}).Methods("POST")

	r.HandleFunc("/workflow/start-multi-agent", func(w http.ResponseWriter, r *http.Request) {
		request := workflows.CollaborationRequest{
			TaskID:           "collab-" + time.Now().Format("20060102150405"),
			PrimaryAgent:     "security",
			ValidationAgents: []string{"compliance", "cost-optimization"},
			Data:             make(map[string]interface{}),
			ConsensusType:    "majority",
		}
		
		we, err := c.ExecuteWorkflow(context.Background(), client.StartWorkflowOptions{
			ID:        "multi-agent-" + time.Now().Format("20060102150405"),
			TaskQueue: "ai-agent-task-queue",
		}, workflows.MultiAgentCollaborationWorkflow, request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte(we.GetID()))
	}).Methods("POST")
	
	r.HandleFunc("/workflow/start", func(w http.ResponseWriter, r *http.Request) {
		we, err := c.ExecuteWorkflow(context.Background(), client.StartWorkflowOptions{
			ID:        "backstage-workflow-" + time.Now().Format("20060102150405"),
			TaskQueue: "ai-agent-task-queue",
		}, HelloBackstageWorkflow, "Backstage")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte(we.GetID()))
	}).Methods("POST")

	r.HandleFunc("/workflow/start-compliance", func(w http.ResponseWriter, r *http.Request) {
		we, err := c.ExecuteWorkflow(context.Background(), client.StartWorkflowOptions{
			ID:        "compliance-workflow-" + time.Now().Format("20060102150405"),
			TaskQueue: "ai-agent-task-queue",
		}, ComplianceCheckWorkflow, "Compliance Data")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte(we.GetID()))
	}).Methods("POST")

	r.HandleFunc("/workflow/signal/{workflowId}", func(w http.ResponseWriter, r *http.Request) {
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
		w.Write([]byte("Signal sent"))
	}).Methods("POST")

	// Infrastructure emulator endpoints
	r.HandleFunc("/emulator/resources", func(w http.ResponseWriter, r *http.Request) {
		resourceType := r.URL.Query().Get("type")
		resources, err := emulator.ListResources(context.Background(), resourceType)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
		w.Header().Set("Content-Type", "application/json")
		// Simple JSON encoding
		json.NewEncoder(w).Encode(resources)
	}).Methods("GET")

	r.HandleFunc("/emulator/resources/{id}", func(w http.ResponseWriter, r *http.Request) {
		resourceID := mux.Vars(r)["id"]
		resource, err := emulator.GetResource(context.Background(), resourceID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resource)
	}).Methods("GET")

	r.HandleFunc("/emulator/resources/{id}/security", func(w http.ResponseWriter, r *http.Request) {
		resourceID := mux.Vars(r)["id"]
		posture, err := emulator.GetSecurityPosture(context.Background(), resourceID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(posture)
	}).Methods("GET")

	r.HandleFunc("/emulator/resources/{id}/compliance", func(w http.ResponseWriter, r *http.Request) {
		resourceID := mux.Vars(r)["id"]
		standards := []string{"SOC2", "GDPR", "HIPAA"} // Default standards
		status, err := emulator.GetComplianceStatus(context.Background(), resourceID, standards)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(status)
	}).Methods("GET")

	log.Printf("Starting HTTP server on :8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}
