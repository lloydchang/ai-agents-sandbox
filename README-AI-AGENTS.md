# Temporal + AI Agents Implementation

This implementation provides a comprehensive local development environment for AI agent experimentation using Temporal and Backstage, focusing on compliance workflows and agent orchestration.

## Architecture Overview

### Core Components

1. **Backend Services (Go)**
   - Temporal workflow orchestration
   - AI agent workers with structured I/O
   - Infrastructure emulator for safe cloud simulation
   - REST API endpoints for workflow management

2. **Frontend Plugin (React/TypeScript)**
   - AI Agent Workflow Builder
   - Real-time Monitoring Dashboard
   - Human-in-the-loop interfaces
   - Compliance visualization

3. **Infrastructure**
   - Local Temporal cluster with PostgreSQL
   - Docker Compose for service orchestration
   - Safe sandboxed environment

## Features

### 🤖 AI Agent Workflows

- **Multi-Agent Orchestration**: Coordinate multiple specialized AI agents
- **Collaboration Patterns**: Agent-to-agent communication and consensus building
- **Human-in-the-Loop**: Seamless integration of human decision points
- **Compliance Focus**: Built-in support for SOC2, GDPR, HIPAA standards

### 🏗️ Infrastructure Emulation

- **Safe Simulation**: Emulate AWS, Azure, GCP resources without production impact
- **Real-time Metrics**: CPU, memory, disk utilization monitoring
- **Security Posture Analysis**: Automated security scanning and compliance checking
- **Multi-cloud Support**: Unified view across cloud providers

### 📊 Monitoring & Visualization

- **Real-time Dashboard**: Live workflow execution status
- **Agent Performance**: Individual agent metrics and scoring
- **Resource Health**: Infrastructure monitoring with health scores
- **Audit Trails**: Complete workflow history for compliance

## Quick Start

### Prerequisites

- Docker and Docker Compose
- Go 1.25+
- Node.js 16+
- Make (optional)

### 1. Start Infrastructure Services

```bash
cd backend
docker-compose up -d
```

This starts:
- PostgreSQL (port 5432)
- Temporal Server (port 7233)
- Temporal UI (port 8080)

### 2. Start Backend Services

```bash
cd backend
go mod tidy
go run main.go
```

The backend will start on port 8081 with:
- Workflow execution endpoints
- Infrastructure emulator API
- Agent management interfaces

### 3. Start Frontend

```bash
cd frontend
yarn install
yarn start
```

The frontend will be available at http://localhost:3000

### 4. Access Temporal UI

Visit http://localhost:8080 to view:
- Running workflows
- Execution history
- Worker status
- Workflow metrics

## Usage Examples

### 1. AI Agent Orchestration

```bash
curl -X POST http://localhost:8081/workflow/start-ai-orchestration
```

This triggers a workflow that:
1. Discovers infrastructure resources
2. Runs Security, Compliance, and Cost Optimization agents in parallel
3. Aggregates results
4. Routes to human review if needed

### 2. Multi-Agent Collaboration

```bash
curl -X POST http://localhost:8081/workflow/start-multi-agent
```

This demonstrates:
1. Primary agent analysis
2. Validation by multiple secondary agents
3. Consensus building
4. Final recommendation generation

### 3. Human-in-the-Loop Workflow

```bash
curl -X POST http://localhost:8081/workflow/start-human-in-loop
```

This creates a workflow that:
1. Performs automated compliance checks
2. Pauses for human decision
3. Resumes based on human input
4. Records audit trail

## API Endpoints

### Workflow Management

- `POST /workflow/start-ai-orchestration` - Start AI orchestration workflow
- `POST /workflow/start-multi-agent` - Start multi-agent collaboration
- `POST /workflow/start-human-in-loop` - Start human-in-the-loop workflow
- `GET /workflow/status?id={workflowId}` - Check workflow status

### Infrastructure Emulator

- `GET /emulator/resources` - List all emulated resources
- `GET /emulator/resources/{id}` - Get specific resource details
- `GET /emulator/resources/{id}/security` - Get security posture
- `GET /emulator/resources/{id}/compliance` - Get compliance status

## Workflow Templates

The `examples/workflow-templates.yaml` file contains pre-defined templates for:

- **Security Compliance Scan**: Comprehensive security and compliance validation
- **Multi-Cloud Compliance**: Cross-cloud compliance validation
- **Continuous Compliance Monitoring**: Automated ongoing compliance checks
- **Incident Response Compliance**: Compliance-focused incident response
- **Vendor Risk Assessment**: Automated vendor risk evaluation

## Agent Types

### Security Agent
- Vulnerability scanning
- Security posture analysis
- Threat detection
- Access control validation

### Compliance Agent
- Regulatory standard validation (SOC2, GDPR, HIPAA)
- Control assessment
- Gap analysis
- Audit trail generation

### Cost Optimization Agent
- Resource utilization analysis
- Cost optimization recommendations
- Right-sizing suggestions
- Reserved instance analysis

## Configuration

### Backend Configuration

Environment variables:
- `TEMPORAL_HOST`: Temporal server address (default: localhost:7233)
- `DB_HOST`: PostgreSQL host (default: localhost:5432)
- `LOG_LEVEL`: Logging level (default: info)

### Agent Configuration

Each agent can be configured via:
- JSON configuration files
- Environment variables
- Runtime parameters

## Development

### Adding New Agents

1. Create agent activity in `activities/ai_agent_activities.go`
2. Register activity in `main.go`
3. Add agent type to frontend components
4. Update workflow templates

### Extending Workflows

1. Define new workflow in `workflows/ai_agent_workflows.go`
2. Add corresponding activities
3. Create API endpoint in `main.go`
4. Update frontend workflow builder

### Adding Compliance Standards

1. Update `emulators/infrastructure_emulator.go`
2. Add standard to `examples/workflow-templates.yaml`
3. Update frontend compliance visualizations

## Testing

### Unit Tests

```bash
cd backend
go test ./...
```

### Integration Tests

```bash
cd backend
go test -tags=integration ./...
```

### Frontend Tests

```bash
cd frontend
yarn test
```

## Monitoring

### Metrics Collection

The system collects:
- Workflow execution times
- Agent performance scores
- Resource utilization metrics
- Error rates and retry counts

### Logging

Structured logging includes:
- Workflow execution context
- Agent decision reasoning
- Human interaction events
- System health status

### Alerting

Configurable alerts for:
- Failed workflow executions
- Low agent scores
- High resource utilization
- Compliance violations

## Security Considerations

### Data Protection

- All data remains local during development
- No external API calls by default
- Encrypted communication between components
- Audit logging for all actions

### Access Control

- Role-based access controls in workflows
- Human approval gates for critical actions
- Immutable audit trails
- Secure credential handling

## Troubleshooting

### Common Issues

1. **Temporal Connection Failed**
   - Ensure Docker services are running
   - Check network connectivity
   - Verify port accessibility

2. **Agent Timeouts**
   - Increase activity timeouts
   - Check resource availability
   - Review agent configuration

3. **Frontend Build Errors**
   - Install dependencies with `yarn install`
   - Check TypeScript configuration
   - Verify environment variables

### Debug Mode

Enable debug logging:
```bash
LOG_LEVEL=debug go run main.go
```

## Contributing

1. Fork the repository
2. Create feature branch
3. Implement changes with tests
4. Update documentation
5. Submit pull request

## License

This project is licensed under the Apache License 2.0 - see the LICENSE file for details.

## Support

For questions and support:
- Check the documentation
- Review the examples
- Open an issue on GitHub
- Join the community discussions
