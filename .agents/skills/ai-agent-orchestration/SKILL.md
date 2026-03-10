---
name: ai-agent-orchestration
description: Orchestrate AI agents and manage agent workflows. Use when designing agent interactions, managing agent execution, coordinating multiple agents, or implementing agent communication patterns.
---

# AI Agent Orchestration

When orchestrating AI agents, follow these principles:

## 1. Agent Design
- Define clear agent responsibilities and boundaries
- Design agent interfaces and communication protocols
- Implement proper agent state management
- Handle agent lifecycle events

## 2. Workflow Coordination
- Create agent workflow definitions
- Define agent execution sequences
- Handle agent dependencies and prerequisites
- Implement agent failover and recovery

## 3. Communication Patterns
- Design agent-to-agent messaging
- Implement event-driven communication
- Handle asynchronous agent interactions
- Manage agent conversation context

## 4. Resource Management
- Allocate agent resources appropriately
- Monitor agent performance and health
- Handle agent scaling and load balancing
- Manage agent configuration and secrets

## Agent Types in This Sandbox
- **Compliance Agents**: Validate policies and regulations
- **Security Agents**: Perform security analysis and monitoring
- **Cost Agents**: Optimize resource usage and spending
- **Infrastructure Agents**: Manage and discover infrastructure
- **Workflow Agents**: Coordinate and execute workflows

## Implementation Patterns
```go
// Agent workflow example
func AgentWorkflow(ctx workflow.Context, input AgentInput) error {
    // Initialize agent
    agent := NewAgent(input.Type)
    
    // Execute agent logic
    result, err := workflow.ExecuteActivity(ctx, agent.Run, input.Params)
    if err != nil {
        return err
    }
    
    // Handle agent output
    return workflow.ExecuteActivity(ctx, ProcessAgentResult, result.Value)
}
```

## Configuration Files
- Agent definitions: `backend/agents/`
- Workflow configurations: `backend/workflows/`
- Skill definitions: `.agents/skills/` and `.claude/skills/`

## Monitoring and Debugging
- Track agent execution metrics
- Monitor agent communication patterns
- Log agent decision points
- Debug agent failures with detailed logs

## Best Practices
- Keep agent responsibilities focused and single-purpose
- Implement proper error handling and retries
- Design for agent isolation and fault tolerance
- Use structured logging for agent activities
- Test agent interactions thoroughly
