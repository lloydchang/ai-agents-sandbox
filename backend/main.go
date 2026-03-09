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

func main() {
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatal("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, "backstage-task-queue", worker.Options{})
	w.RegisterWorkflow(HelloBackstageWorkflow)
	w.RegisterActivity(FetchDataActivity)
	w.RegisterActivity(ProcessDataActivity)

	err = w.Start()
	if err != nil {
		log.Fatal("Unable to start worker", err)
	}

	// HTTP server for endpoints
	r := mux.NewRouter()
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
