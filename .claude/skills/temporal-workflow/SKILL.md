---
name: temporal-workflow
description: Create, manage, and monitor Temporal workflows with AI agent orchestration. Use when developing workflow definitions, monitoring execution, or troubleshooting workflow issues.
argument-hint: "[action] [workflowName] [parameters]"
disable-model-invocation: false
user-invocable: true
allowed-tools: 
  - Bash
  - Read
  - Write
  - Grep
---

# Temporal Workflow Skill

Comprehensive Temporal workflow management for the AI Agents sandbox. This skill provides end-to-end workflow lifecycle management from creation to monitoring and debugging.

## Usage
```bash
/temporal-workflow create my-workflow "Go workflow for data processing"
/temporal-workflow status my-workflow
/temporal-workflow monitor my-workflow --live
/temporal-workflow debug my-workflow --history=50
/temporal-workflow test my-workflow --unit
```

## Core Capabilities

### 1. Workflow Creation & Development
```bash
# Create new workflow with template
/temporal-workflow create order-processing "Handles e-commerce order lifecycle"

# Generate workflow skeleton
/temporal-workflow scaffold payment-workflow --activities=validate,process,notify

# Add activity to existing workflow
/temporal-workflow add-activity my-workflow data-validation
```

### 2. Workflow Monitoring
```bash
# Real-time monitoring
/temporal-workflow monitor my-workflow --live

# Historical analysis
/temporal-workflow history my-workflow --days=7

# Performance metrics
/temporal-workflow metrics my-workflow --detailed
```

### 3. Debugging & Troubleshooting
```bash
# Debug failed workflow
/temporal-workflow debug my-workflow --error=timeout

# Query workflow state
/temporal-workflow query my-workflow currentState

# Replay workflow execution
/temporal-workflow replay my-workflow --run-id=12345
```

## Workflow Templates

### Standard Workflow Pattern
```go
func StandardWorkflow(ctx workflow.Context, input WorkflowInput) error {
    // Setup
    logger := workflow.GetLogger(ctx)
    logger.Info("Starting workflow", "input", input)
    
    // Activity execution with retries
    retryOptions := workflow.RetryOptions{
        InitialInterval: time.Second * 1,
        BackoffCoefficient: 2.0,
        MaximumInterval: time.Second * 30,
        MaximumAttempts: 5,
    }
    
    result, err := workflow.ExecuteActivity(ctx, retryOptions, ProcessActivity, input.Data)
    if err != nil {
        return err
    }
    
    // Completion
    logger.Info("Workflow completed", "result", result)
    return nil
}
```

### Workflow with Compensation
```go
func CompensationWorkflow(ctx workflow.Context, input WorkflowInput) error {
    // Execute activities with compensation
    ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
        StartToCloseTimeout: time.Minute * 10,
    })
    
    var compensation workflow.Compensation
    workflow.SetCompensationHook(ctx, func(ctx workflow.Context) error {
        return compensation(ctx)
    })
    
    // Main workflow logic
    err := workflow.ExecuteActivity(ctx, CreateResourceActivity, input.Resource)
    if err != nil {
        return err
    }
    
    compensation = func(ctx workflow.Context) error {
        return workflow.ExecuteActivity(ctx, DeleteResourceActivity, input.Resource)
    }
    
    // Continue with workflow...
    return nil
}
```

## Integration Points

### Backend Integration
- **Workflows**: `backend/workflows/`
- **Activities**: `backend/activities/`
- **Main Entry**: `backend/main.go`
- **Tests**: `backend/workflows/*_test.go`

### API Endpoints
```bash
# Start workflow
curl -X POST http://localhost:8081/api/v1/workflows/start \
  -H "Content-Type: application/json" \
  -d '{"workflow": "my-workflow", "input": {...}}'

# Get workflow status
curl http://localhost:8081/api/v1/workflows/{workflow-id}/status

# Cancel workflow
curl -X DELETE http://localhost:8081/api/v1/workflows/{workflow-id}
```

## Development Workflow

### 1. Create Workflow Definition
```bash
/temporal-workflow create my-new-workflow "Description of workflow purpose"
```

### 2. Define Activities
```go
// In backend/activities/my_activities.go
func MyActivity(ctx context.Context, input ActivityInput) (ActivityOutput, error) {
    // Activity implementation
    return result, nil
}
```

### 3. Register Workflow
```go
// In backend/main.go
worker.RegisterWorkflow(MyNewWorkflow)
worker.RegisterActivity(MyActivity)
```

### 4. Test Implementation
```bash
/temporal-workflow test my-new-workflow --unit
/temporal-workflow test my-new-workflow --integration
```

## Monitoring Dashboard

### Key Metrics
- **Workflow Success Rate**: Percentage of successful executions
- **Average Execution Time**: Performance tracking
- **Activity Latency**: Per-activity performance metrics
- **Error Rates**: Failure analysis by category

### Real-time Status
```bash
# Live monitoring dashboard
/temporal-workflow dashboard --refresh=5s

# Workflow topology view
/temporal-workflow topology --workflow=my-workflow
```

## Error Handling Patterns

### Retry Strategies
```go
// Exponential backoff
retryOptions := workflow.RetryOptions{
    InitialInterval:    time.Second,
    BackoffCoefficient: 2.0,
    MaximumInterval:    time.Minute,
    MaximumAttempts:    10,
}

// Linear backoff
retryOptions := workflow.RetryOptions{
    InitialInterval:    time.Second * 5,
    BackoffCoefficient: 1.0,
    MaximumAttempts:    5,
}
```

### Circuit Breaker Pattern
```go
func CircuitBreakerWorkflow(ctx workflow.Context, input WorkflowInput) error {
    // Track failure count
    failureCount := 0
    maxFailures := 3
    
    for attempt := 0; attempt < 5; attempt++ {
        err := workflow.ExecuteActivity(ctx, RiskyActivity, input)
        if err == nil {
            failureCount = 0 // Reset on success
            return nil
        }
        
        failureCount++
        if failureCount >= maxFailures {
            return workflow.NewContinueAsNewError(ctx, FallbackWorkflow, input)
        }
        
        // Wait before retry
        workflow.Sleep(ctx, time.Duration(attempt)*time.Second)
    }
    
    return errors.New("max attempts exceeded")
}
```

## Best Practices

1. **Idempotent Activities**: Design activities to be safely retryable
2. **Proper Logging**: Use structured logging with workflow context
3. **Timeout Management**: Set appropriate timeouts for workflows and activities
4. **Error Handling**: Implement comprehensive error handling and retries
5. **Testing**: Write unit tests for workflows and integration tests for activities
6. **Monitoring**: Set up alerts for workflow failures and performance issues

## Troubleshooting Guide

### Common Issues
- **Workflow Not Starting**: Check worker registration and activity availability
- **Activity Timeouts**: Review timeout configurations and activity performance
- **Memory Leaks**: Monitor workflow state and ensure proper cleanup
- **Deadlocks**: Check for circular dependencies in activity calls

### Debug Commands
```bash
# Check worker status
/temporal-workflow worker-status

# Validate workflow definition
/temporal-workflow validate my-workflow

# Check activity registration
/temporal-workflow activities --list
```

## Related Skills

- `/workflow-management`: Orchestrate multiple workflows
- `/ai-agent-orchestration`: Coordinate AI agent interactions
- `/compliance-check`: Validate workflow compliance
- `/cost-optimization`: Optimize workflow resource usage
