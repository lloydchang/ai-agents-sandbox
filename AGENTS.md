# Temporal AI Agents - Agent Operating Manual

## Project Overview
This repository implements the Temporal AI Agents system, providing a comprehensive orchestration platform for AI agent workflows in enterprise environments. The system supports multiple interfaces including REST APIs, MCP servers, CLI tools, WebMCP clients, enhanced GUIs, and AI assistant integrations.

## System Architecture
- **Backend**: Go-based Temporal workflows with enhanced activities
- **Frontend**: React/Material-UI dashboard for agent management
- **APIs**: REST endpoints for programmatic access
- **MCP**: Model Context Protocol server for AI tool interoperability
- **CLI**: Command-line interface for workflow operations
- **WebMCP**: Browser-based MCP client interface
- **Skills**: AI assistant integration via SKILL.md specifications

## Repository Structure
```
repo/
├── backend/           # Go Temporal workflows and activities
├── frontend/          # React dashboard and WebMCP client
├── cli/              # Command-line interface
├── docs/             # Documentation and interface specs
├── SKILL.md          # AI assistant skill definitions
├── AGENTS.md         # This file - agent operating rules
└── tools/            # Tool permissions and configurations
```

## Agent Behavior Rules

### Core Principles
- **Safety First**: Never execute destructive operations without explicit approval
- **Audit Trail**: All agent actions must be logged and traceable
- **Human Oversight**: Critical decisions require human review
- **Idempotency**: Operations should be safe to retry

### Repository Access Rules
**ALLOWED Directories:**
- Read/write: `backend/`, `frontend/`, `cli/`, `docs/`
- Read-only: Root configuration files
- Forbidden: System directories, generated files

**FORBIDDEN Modifications:**
- Never modify `dist/`, `build/`, or generated files
- Never edit migration files directly (use skills instead)
- Never change infrastructure configurations without approval
- Never commit to protected branches

### Workflow Execution Rules
**Trigger Conditions:**
- Use SKILL.md files for specialized workflows
- Require explicit user approval for destructive operations
- Validate inputs before executing workflows
- Maintain state consistency across retries

**Error Handling:**
- Log all failures with context
- Retry transient failures automatically
- Escalate critical failures to human operators
- Provide meaningful error messages

### Code Generation Rules
**Standards Compliance:**
- Follow TypeScript/JavaScript best practices
- Use established patterns from existing codebase
- Include proper error handling and logging
- Add tests for new functionality

**File Organization:**
- Place new components in appropriate directories
- Follow naming conventions from existing code
- Update imports and dependencies correctly
- Maintain clean git history

### Security Constraints
**Command Restrictions:**
- No direct shell execution without tool approval
- No network requests to unapproved endpoints
- No file system operations outside allowed paths
- No database modifications without migration workflows

**Data Protection:**
- Never expose sensitive configuration
- Sanitize all user inputs
- Use secure communication channels
- Respect data retention policies

## Testing Requirements
**Before Committing:**
- Run all existing tests: `npm test`
- Lint code: `npm run lint`
- Build successfully: `npm run build`
- Update documentation for API changes

**Workflow Validation:**
- Test workflows with mock data first
- Verify error handling paths
- Check resource cleanup on failures
- Validate state persistence

## Deployment Rules
**Staging Environment:**
- Deploy to staging before production
- Run integration tests in staging
- Verify monitoring and logging
- Perform security scans

**Production Deployment:**
- Require code review approval
- Use blue-green deployment strategy
- Monitor key metrics post-deployment
- Have rollback plan ready

## Communication Standards
**User Interaction:**
- Provide clear, actionable feedback
- Explain complex operations in simple terms
- Offer progress updates for long-running tasks
- Suggest next steps when appropriate

**Error Reporting:**
- Include error codes and descriptions
- Provide troubleshooting guidance
- Suggest contact information for issues
- Log errors with full context

## Interface-Specific Rules

### REST API Usage
- Use proper HTTP methods and status codes
- Validate all inputs and outputs
- Rate limit requests appropriately
- Document API endpoints clearly

### MCP Server Interaction
- Follow MCP protocol specifications
- Handle tool registration correctly
- Manage resource subscriptions
- Provide comprehensive tool metadata

### CLI Operations
- Use consistent command structure
- Provide help text and examples
- Support both interactive and scripted modes
- Handle signals gracefully

### GUI/Dashboard Access
- Ensure responsive design
- Provide accessibility features
- Include loading states and error handling
- Support keyboard navigation

### AI Assistant Integration
- Follow SKILL.md specifications
- Provide clear instructions and examples
- Handle edge cases gracefully
- Maintain conversation context

## Monitoring and Observability
**Required Metrics:**
- Workflow execution times and success rates
- Error rates by component
- Resource utilization
- User interaction patterns

**Logging Standards:**
- Use structured logging with consistent fields
- Include correlation IDs for request tracing
- Log security events appropriately
- Maintain audit trails for compliance

## Scaling Considerations
**Performance Optimization:**
- Optimize workflow execution paths
- Cache frequently accessed data
- Use efficient algorithms and data structures
- Monitor resource consumption

**Concurrency Management:**
- Handle concurrent workflow executions
- Prevent resource conflicts
- Implement proper locking mechanisms
- Support horizontal scaling

## Emergency Procedures
**System Outages:**
- Activate backup workflows if available
- Notify stakeholders of issues
- Provide status updates regularly
- Restore services with minimal downtime

**Security Incidents:**
- Isolate affected systems
- Preserve evidence for investigation
- Communicate transparently with users
- Implement remediation measures

## Development Workflow
**Feature Development:**
1. Create issue/ticket for work
2. Implement changes following rules above
3. Write/update tests
4. Update documentation
5. Submit pull request for review

**Code Review Process:**
- Review code for security issues
- Verify compliance with agent rules
- Test functionality thoroughly
- Ensure documentation is updated
- Approve only after all checks pass

## Contact Information
**For Issues:**
- Create GitHub issue with detailed description
- Include error logs and reproduction steps
- Tag appropriately for routing

**For Security Concerns:**
- Use dedicated security reporting channel
- Provide minimal information initially
- Allow time for investigation

This AGENTS.md file serves as the operating manual for AI agents working in this repository. All agents must follow these rules to ensure safe, reliable, and compliant operation of the Temporal AI Agents system.
