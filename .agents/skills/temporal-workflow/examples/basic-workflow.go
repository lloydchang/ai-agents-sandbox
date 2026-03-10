package main

import (
    "time"
    "go.temporal.io/sdk/workflow"
    "go.temporal.io/sdk/activity"
)

// BasicWorkflow demonstrates a simple Temporal workflow pattern
func BasicWorkflow(ctx workflow.Context, input WorkflowInput) (WorkflowOutput, error) {
    logger := workflow.GetLogger(ctx)
    logger.Info("Starting basic workflow", "input", input)

    // Configure activity options
    activityOptions := workflow.ActivityOptions{
        StartToCloseTimeout: time.Minute * 10,
        RetryPolicy: &temporal.RetryPolicy{
            InitialInterval:    time.Second,
            BackoffCoefficient: 2.0,
            MaximumInterval:    time.Minute,
            MaximumAttempts:    5,
        },
    }
    ctx = workflow.WithActivityOptions(ctx, activityOptions)

    // Execute activities sequentially
    var result1 string
    err := workflow.ExecuteActivity(ctx, ProcessDataActivity, input.Data).Get(ctx, &result1)
    if err != nil {
        return WorkflowOutput{}, err
    }

    var result2 string
    err = workflow.ExecuteActivity(ctx, ValidateResultActivity, result1).Get(ctx, &result2)
    if err != nil {
        return WorkflowOutput{}, err
    }

    output := WorkflowOutput{
        Result:      result2,
        ProcessedAt: workflow.Now(ctx),
    }

    logger.Info("Workflow completed successfully", "output", output)
    return output, nil
}

// WorkflowInput represents the input structure
type WorkflowInput struct {
    Data string `json:"data"`
}

// WorkflowOutput represents the output structure
type WorkflowOutput struct {
    Result      string    `json:"result"`
    ProcessedAt time.Time `json:"processed_at"`
}
