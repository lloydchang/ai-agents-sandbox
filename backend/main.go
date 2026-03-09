package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
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

	var reviewResult string
	err = workflow.ExecuteActivity(ctx, HumanReviewActivity, aggregatedResult).Get(ctx, &reviewResult)
	if err != nil {
		return "", err
	}

	return reviewResult, nil
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
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatal("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, "backstage-task-queue", worker.Options{})
	w.RegisterWorkflow(HelloBackstageWorkflow)
	w.RegisterWorkflow(ComplianceCheckWorkflow)
	w.RegisterActivity(FetchDataActivity)
	w.RegisterActivity(ProcessDataActivity)
	w.RegisterActivity(AgentCheckActivity)
	w.RegisterActivity(AggregateResultsActivity)
	w.RegisterActivity(HumanReviewActivity)

	err = w.Start()
	if err != nil {
		log.Fatal("Unable to start worker", err)
	}

	// HTTP server for endpoints
	r := mux.NewRouter()
	
	// Apply CORS middleware
	r.Use(corsMiddleware)
	
	// Add explicit OPTIONS handler for CORS preflight
	r.HandleFunc("/workflow/start", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}).Methods("OPTIONS")
	
	r.HandleFunc("/workflow/status", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}).Methods("OPTIONS")
	
	r.HandleFunc("/workflow/start", func(w http.ResponseWriter, r *http.Request) {
		we, err := c.ExecuteWorkflow(context.Background(), client.StartWorkflowOptions{
			ID:        "backstage-workflow-" + time.Now().Format("20060102150405"),
			TaskQueue: "backstage-task-queue",
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
			TaskQueue: "backstage-task-queue",
		}, ComplianceCheckWorkflow, "Compliance Data")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte(we.GetID()))
	}).Methods("POST")

	r.HandleFunc("/workflow/status", func(w http.ResponseWriter, r *http.Request) {
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
		w.Write([]byte(resp.WorkflowExecutionInfo.Status.String()))
	}).Methods("GET")

	log.Fatal(http.ListenAndServe(":8081", r))
}
