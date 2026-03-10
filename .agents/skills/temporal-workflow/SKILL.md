---
name: temporal-workflow
description: Manage and monitor Temporal workflows with AI agent orchestration. Use when creating, monitoring, or troubleshooting Temporal workflows, handling workflow execution, or managing workflow state.
---

# Temporal Workflow Management

When working with Temporal workflows, always follow these steps:

## 1. Workflow Analysis
- Examine the current workflow definition in `backend/workflows/`
- Check for existing activities in `backend/activities/`
- Review workflow configuration and dependencies

## 2. Workflow Creation
- Create new workflow files following Go Temporal patterns
- Define proper workflow interfaces and input/output structs
- Implement error handling and retry policies
- Add appropriate logging and monitoring

## 3. Activity Integration
- Connect workflows to appropriate activities
- Ensure activity functions are properly registered
- Handle activity timeouts and heartbeats

## 4. Testing and Validation
- Create unit tests for workflows
- Test workflow execution paths
- Validate error handling scenarios

## 5. Monitoring and Debugging
- Use Temporal UI to monitor workflow execution
- Check workflow history for troubleshooting
- Analyze workflow performance metrics

## Key Commands
```bash
# Start Temporal worker
go run backend/main.go

# Check workflow status
# Use Temporal UI at http://localhost:8233

# Run tests
go test ./backend/workflows/...
```

## Common Patterns
- Use workflow.Sleep for retries with exponential backoff
- Implement proper cancellation handling
- Use workflow.ExecuteActivity for activity calls
- Log workflow events with structured logging

## File Locations
- Workflows: `backend/workflows/`
- Activities: `backend/activities/`
- Main entry: `backend/main.go`
- Tests: `backend/workflows/*_test.go`
