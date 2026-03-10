# Temporal AI Agents - Multi-Interface Platform

A comprehensive AI-native engineering platform that combines Temporal durable workflows with AI agent orchestration, providing multiple interfaces for seamless integration: **REST APIs**, **MCP Server**, **WebMCP Client**, **CLI Tool**, and **Enhanced GUI**.

## 🏗️ Architecture Overview

```
AI Assistant (Claude/Codex)
    |
    v
Coordination Layer
├── AGENTS.md (Agent behavior rules)
├── SKILL.md (Specialized capabilities)
└── tools/ (Safe execution boundaries)
    |
    v
8 Interface Types
├── REST APIs (Programmatic access)
├── MCP Server (Standardized protocol)
├── WebMCP Client (Browser interface)
├── CLI Tool (Command-line skills)
├── Enhanced GUI (Visual management)
├── SKILL.md Integration (AI assistant)
├── AGENTS.md Rules (Behavior constraints)
└── Tool Configuration (Safety boundaries)
    |
    v
Temporal Engine (Go Backend)
├── AI Agent Orchestration
├── Multi-Agent Collaboration
├── Human-in-the-Loop Workflows
├── Compliance Automation
└── Infrastructure Emulation
    |
    v
PostgreSQL + Temporal Server
```

## 🎯 Key Features

### 🤖 AI Agent Orchestration
- **Multi-Agent Collaboration**: Coordinate specialized AI agents for complex tasks
- **Human-in-the-Loop**: Seamless integration of human decision points in workflows
- **Compliance Automation**: Built-in support for SOC2, GDPR, HIPAA standards
- **Infrastructure Emulation**: Safe simulation of AWS, Azure, GCP resources

### 🔧 Multiple Interfaces
- **REST APIs**: Full programmatic access to workflows and agent capabilities
- **MCP Server**: Standardized Model Context Protocol for AI tool interoperability
- **WebMCP Client**: Browser-based interface for interactive agent management
- **CLI Tool**: Command-line interface with skill invocation (`/skill-name` syntax)
- **Enhanced GUI**: Visual workflow builder and real-time monitoring dashboard

### 📋 Coordination Layer
- **AGENTS.md**: Global agent behavior rules and repository policies
- **SKILL.md**: Specialized workflow capabilities with structured I/O
- **Tool Configurations**: Safe execution boundaries for bash, git, kubectl, terraform, docker

## 🚀 Quick Start

### Prerequisites
- Go 1.21+
- Node.js 18+
- Docker & Docker Compose

### 1. Start Infrastructure
```bash
cd backend && docker-compose up -d
```
Starts PostgreSQL, Temporal Server, and Temporal UI.

### 2. Start Backend (All Interfaces)
```bash
cd backend && go run main.go
```
Launches:
- REST API server on `:8081`
- MCP server for AI tool integration
- WebMCP client interface
- CLI-ready endpoints

### 3. Start Frontend (GUI)
```bash
cd frontend && yarn start
```
Backstage app with AI agent dashboard on `http://localhost:3000`

### 4. Test AI Agent Workflows
**Via GUI:**
1. Navigate to `http://localhost:3000/temporal`
2. Select workflow type (AI Orchestration, Multi-Agent, Compliance)
3. Click "Start [Workflow Type]" and monitor execution

**Via CLI:**
```bash
cd cli && go run main.go skill invoke compliance-check vm-web-server-001 SOC2
```

**Via REST API:**
```bash
curl -X POST http://localhost:8081/workflow/start-ai-orchestration \
  -H "Content-Type: application/json"
```

## 🛠️ CLI Usage

The CLI provides comprehensive skill management and workflow control:

### Skill Management
```bash
# List all available skills
./cli skill list

# Get detailed skill information
./cli skill info compliance-check

# Invoke a skill with arguments
./cli skill invoke compliance-check vm-web-server-001 SOC2 high

# Use slash syntax for skills
./cli /compliance-check vm-web-server-001
```

### Interactive Mode
```bash
./cli interactive
```
Available commands:
- `skill <name> [args]` - Invoke skills
- `skills list` - List available skills
- `skills info <name>` - Get skill details
- `start <type>` - Start workflows
- `status <id>` - Check workflow status
- `signal <id> <name> <value>` - Send signals
- `health` - Check server health

### Workflow Management
```bash
# Start different workflow types
./cli workflow start ai-orchestration
./cli workflow start multi-agent
./cli workflow start compliance

# Monitor workflow status
./cli workflow status <workflow-id>

# Send signals to running workflows
./cli workflow signal <workflow-id> approval true
```

## 🔌 API Endpoints

### REST APIs
```bash
# Start workflows
POST /workflow/start-ai-orchestration
POST /workflow/start-multi-agent
POST /workflow/start-compliance

# Workflow management
GET /workflow/status?id=<workflow_id>
POST /workflow/signal/<workflow_id>

# Skill execution
POST /api/skills/{skill_name}/execute
GET /api/skills
GET /api/skills/{skill_name}

# MCP integration
POST /mcp
GET /mcp/resources
GET /mcp/tools

# Health and monitoring
GET /health
GET /metrics
```

### MCP Server
The MCP server provides standardized AI tool integration:
- **Tools**: start_compliance_workflow, get_workflow_status, signal_workflow, get_infrastructure_info
- **Resources**: workflow_results, agent_capabilities, compliance_reports
- **Protocol**: JSON-RPC 2.0 over WebSocket/HTTP

## 📁 Repository Structure

```
repo/
├── AGENTS.md              # Agent behavior rules
├── SKILL.md               # AI skill definitions
├── backend/               # Go Temporal workflows
│   ├── activities/        # AI agent activities
│   ├── mcp/              # MCP server implementation
│   └── main.go           # Multi-interface server
├── frontend/              # React/TypeScript Backstage
│   ├── src/components/   # WebMCP client, AgentManagement
│   └── plugins/          # Temporal integration plugin
├── cli/                  # Command-line interface
├── tools/                # Tool safety configurations
│   ├── bash.yaml         # Bash command restrictions
│   ├── git.yaml          # Git operation limits
│   ├── kubectl.yaml      # Kubernetes inspection
│   ├── terraform.yaml    # Infrastructure planning
│   └── docker.yaml       # Container management
├── docs/                 # Documentation and guides
└── scripts/              # Development automation
```

## 🏃‍♂️ Development Scripts

### Full Environment
```bash
./scripts/dev.sh    # Start infrastructure + frontend
```

### Build Everything
```bash
./scripts/build.sh  # Docker images for all components
```

### Validation
```bash
./scripts/validate.sh  # Comprehensive environment check
```

## 🔒 Coordination Layer

### AGENTS.md
Global agent operating manual defining:
- **Repository Access Rules**: Allowed/forbidden directories and operations
- **Workflow Execution Rules**: Safe execution patterns and error handling
- **Code Generation Standards**: TypeScript/JavaScript best practices
- **Security Constraints**: Command restrictions and data protection
- **Interface-Specific Rules**: REST API, MCP, CLI, GUI guidelines

### SKILL.md
Structured skill definitions with:
- **Input/Output Schemas**: Standardized parameters and returns
- **Tool Requirements**: Required permissions and resources
- **Usage Examples**: Concrete invocation patterns
- **Error Handling**: Failure modes and recovery

### Tool Configurations
Safe execution boundaries:
- **bash.yaml**: npm, git, docker, terraform commands with restrictions
- **git.yaml**: Read-only operations (status, diff, log, blame)
- **kubectl.yaml**: Cluster inspection without destructive operations
- **terraform.yaml**: Planning/safety operations only
- **docker.yaml**: Container inspection with resource limits

## 🧪 Testing

### Backend Tests
```bash
cd backend && go test ./...
```

### Frontend Tests
```bash
cd frontend && yarn test
```

### CLI Tests
```bash
cd cli && go test ./...
```

### Integration Tests
```bash
./scripts/validate.sh
```

## 🔍 Monitoring & Observability

### Temporal UI
- **Workflow Execution**: Real-time workflow visualization at `http://localhost:8080`
- **Activity Logs**: Detailed execution history and performance metrics
- **Worker Status**: Health and capacity monitoring

### Agent Dashboard
- **Workflow Status Table**: Real-time execution monitoring
- **Performance Metrics**: Agent execution times and success rates
- **Resource Health**: Infrastructure utilization and compliance scores

### Audit Trails
- **Complete Workflow History**: All executions with full context
- **Agent Decisions**: Reasoning and decision logs
- **Compliance Reports**: Regulatory compliance validation results

## 🐛 Troubleshooting

### Docker Issues
```bash
# Check container status
docker-compose ps

# View logs
docker-compose logs temporal
docker-compose logs postgres

# Restart services
docker-compose restart
```

### Backend Issues
```bash
# Check Temporal connectivity
curl http://localhost:7233/health

# Verify worker registration
docker-compose logs temporal | grep "Worker registered"
```

### CLI Issues
```bash
# Test basic connectivity
./cli health

# Check skill availability
./cli skills list
```

### Frontend Issues
```bash
# Clear cache and reinstall
cd frontend && rm -rf node_modules && yarn install

# Check configuration
cat frontend/app-config.yaml
```

## 🚀 Extensions & Integration

### AI Assistant Integration
- **Claude Skills**: Direct skill invocation via `/skill-name` syntax
- **Codex Skills**: GitHub Copilot integration with workflow orchestration
- **Custom GPTs**: ChatGPT plugin for workflow management

### Enterprise Integration
- **GitHub Actions**: CI/CD pipelines with workflow triggers
- **Slack/Discord Bots**: Real-time workflow notifications
- **Jira Integration**: Ticket-based workflow initiation
- **ServiceNow**: IT service management integration

### Advanced Workflows
- **Multi-Cloud Orchestration**: Cross-provider infrastructure management
- **Compliance Automation**: Automated regulatory compliance checking
- **Security Scanning**: Continuous security posture analysis
- **Cost Optimization**: Automated resource optimization workflows

---

## 📚 Documentation

- **[AGENTS.md](AGENTS.md)**: Agent operating manual and behavior rules
- **[SKILL.md](SKILL.md)**: AI skill definitions and usage patterns
- **[IMPLEMENTATION_SUMMARY.md](IMPLEMENTATION_SUMMARY.md)**: Technical implementation details
- **[docs/](docs/)**: Additional guides and specifications

## 🤝 Contributing

This is an AI-native engineering platform designed for:
- **Safe experimentation** with AI agents and workflows
- **Multi-interface development** supporting various integration patterns
- **Enterprise-grade orchestration** with durable execution guarantees
- **Regulatory compliance** built into the core architecture

## 📄 License

[GNU Affero General Public License v3.0 or later](LICENSE)

---

**Open Source**: https://github.com/lloydchang/sandbox-backstage-temporal
