# Temporal AI Agents - Comprehensive Multi-Interface Platform

A production-ready AI-native engineering platform that combines **Temporal durable workflows** with **AI agent orchestration**, providing comprehensive interfaces for seamless integration: **REST APIs**, **MCP Server**, **WebMCP Client**, **CLI Tools**, and **Enhanced GUI**.

## 🏗️ Architecture Overview

```
AI Assistant (Claude/Codex/CustomGPT)
    |
    v
Coordination Layer
├── AGENTS.md (Agent behavior rules)
├── SKILL.md (Specialized capabilities)
├── .agents/skills/ (Skill definitions & scripts)
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
├── Skills System (Claude/Codex compatible)
├── MCP Integration (Model Context Protocol)
├── Infrastructure Emulation
└── Multi-Interface Support
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
- **MCP Integration**: Standardized Model Context Protocol for tool interoperability

### 🔧 Skills System (Claude/Codex Compatible)
- **SKILL.md Format**: YAML frontmatter + markdown content based on Claude Code specifications
- **Automatic Discovery**: Scans `.agents/skills/` directories with priority-based conflict resolution
- **Skill Execution**: Argument substitution, forked contexts, and tool restrictions
- **Core Skills**: compliance-check, security-scan, cost-analysis, workflow-management
- **Multi-Interface Support**: CLI (`/skill-name`), GUI, API, and AI assistant integration

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
- **Interface Standards**: Consistent patterns across all client types

## 🚀 Quick Start

### Prerequisites
- **Go 1.25+** (for backend services and Temporal workflows)
- **Node.js 16+** (for frontend React/TypeScript application)  
- **Docker & Docker Compose** (for infrastructure services)
- **Make** (optional, for build automation)

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
**Via Web UI (Skills):**
1. Navigate to `http://localhost:3000/skills`
2. Browse available skills (compliance-check, security-scan, cost-analysis, workflow-management)
3. Click "Execute" on any skill with arguments
4. Monitor execution results and status

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
temporal-agents skill list

# Get detailed skill information
temporal-agents skill info compliance-check

# Invoke a skill with arguments
temporal-agents skill invoke compliance-check vm-web-server-001 SOC2 high

# Use slash syntax for skills (Claude/Codex compatible)
temporal-agents skill invoke /compliance-check vm-web-server-001
```

### Interactive Mode
```bash
temporal-agents --interactive
```
Available commands:
- `skill <name> [args]` - Invoke skills with `/skill-name` syntax
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

# Skills management
GET /api/skills                           # List all skills
GET /api/skills/invocable                 # List user-invocable skills
GET /api/skills/{skill_name}              # Get skill details
POST /api/skills/{skill_name}/execute     # Execute skill with arguments
POST /api/skills/discover                 # Trigger skill discovery

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

## � Workflow Templates

The `examples/workflow-templates.yaml` file contains pre-defined templates for:

- **Security Compliance Scan**: Comprehensive security and compliance validation
- **Multi-Cloud Compliance**: Cross-cloud compliance validation
- **Continuous Compliance Monitoring**: Automated ongoing compliance checks
- **Incident Response Compliance**: Compliance-focused incident response
- **Vendor Risk Assessment**: Automated vendor risk evaluation

## 🤖 Agent Types

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

## ⚙️ Configuration

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

```
repo/
├── AGENTS.md              # Agent behavior rules
├── SKILL.md               # AI skill definitions
├── .agents/               # Skills directory structure
│   └── skills/            # Skill definitions (Claude/Codex compatible)
│       ├── compliance-check/
│       │   ├── SKILL.md          # Compliance checking skill
│       │   ├── scripts/          # Supporting scripts
│       │   └── templates/        # Output templates
│       ├── security-scan/        # Security analysis skill
│       ├── cost-analysis/        # Cost optimization skill
│       └── workflow-management/  # Workflow control skill
├── backend/               # Go Temporal workflows
│   ├── activities/        # AI agent activities
│   ├── skills/            # Skills system implementation
│   │   ├── skill.go       # Skill parsing and management
│   │   └── service.go     # HTTP API for skills
│   ├── mcp/              # MCP server implementation
│   └── main.go           # Multi-interface server
├── frontend/              # React/TypeScript Backstage
│   ├── src/components/
│   │   ├── SkillsManagement.tsx  # Skills UI component
│   │   └── AgentManagement.tsx   # Workflow management
│   └── plugins/          # Temporal integration plugin
├── cli/                  # Command-line interface with skills
├── tools/                # Tool safety configurations
│   ├── bash.yaml         # Bash command restrictions
│   ├── git.yaml          # Git operation limits
│   ├── kubectl.yaml      # Kubernetes inspection
│   ├── terraform.yaml    # Infrastructure planning
│   └── docker.yaml       # Container management
├── docs/
│   ├── skills-system-guide.md    # Complete skills documentation
│   ├── comprehensive-interfaces-guide.md
│   └── claude-and-codex-skills.txt
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
- **[docs/skills-system-guide.md](docs/skills-system-guide.md)**: Complete skills system documentation (Claude/Codex compatible)
- **[IMPLEMENTATION_SUMMARY.md](IMPLEMENTATION_SUMMARY.md)**: Technical implementation details
- **[docs/comprehensive-interfaces-guide.md](docs/comprehensive-interfaces-guide.md)**: Complete interfaces documentation (SKILL.md, AGENTS.md, MCP, APIs, CLIs, GUIs)
- **[docs/claude-and-codex-skills.txt](docs/claude-and-codex-skills.txt)**: AI assistant skills and workflows reference
- **[docs/](docs/)**: Additional guides and specifications

## 🤝 Contributing

This is an AI-native engineering platform designed for:
- **Safe experimentation** with AI agents and workflows
- **Multi-interface development** supporting various integration patterns
- **Enterprise-grade orchestration** with durable execution guarantees
- **Regulatory compliance** built into the core architecture
- **AI assistant integration** supporting Claude Skills, Codex Skills, and Custom GPTs

### Interface Development
- **SKILL.md**: Create reusable workflow capabilities for AI assistants
- **AGENTS.md**: Define agent behavior rules and safety constraints
- **MCP Integration**: Implement standardized Model Context Protocol interfaces
- **Multi-Client Support**: Build consistent experiences across CLI, GUI, API, and web interfaces

### Contribution Areas
- **AI Agent Workflows**: New orchestration patterns and agent collaborations
- **Interface Implementations**: Additional client types and protocol support
- **Compliance Frameworks**: New regulatory standards and validation rules
- **Infrastructure Emulation**: Additional cloud providers and services
- **Documentation**: Guides and examples for all interface types

## 📄 License

[GNU Affero General Public License v3.0 or later](LICENSE)

---

**Open Source**: https://github.com/lloydchang/sandbox-backstage-temporal


# Temporal AI Agents - Comprehensive Multi-Interface Platform

A production-ready AI-native engineering platform that combines **Temporal durable workflows** with **AI agent orchestration**, providing comprehensive interfaces for seamless integration: **REST APIs**, **MCP Server**, **WebMCP Client**, **CLI Tools**, and **Enhanced GUI**.

## 🏗️ Architecture Overview

```
AI Assistant (Claude/Codex/CustomGPT)
    |
    v
Coordination Layer
├── AGENTS.md (Agent behavior rules)
├── SKILL.md (Specialized capabilities)
├── .agents/skills/ (Skill definitions & scripts)
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
├── Skills System (Claude/Codex compatible)
├── MCP Integration (Model Context Protocol)
├── Infrastructure Emulation
└── Multi-Interface Support
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
- **MCP Integration**: Standardized Model Context Protocol for tool interoperability

### 🔧 Skills System (Claude/Codex Compatible)
- **SKILL.md Format**: YAML frontmatter + markdown content based on Claude Code specifications
- **Automatic Discovery**: Scans `.agents/skills/` directories with priority-based conflict resolution
- **Skill Execution**: Argument substitution, forked contexts, and tool restrictions
- **Core Skills**: compliance-check, security-scan, cost-analysis, workflow-management
- **Multi-Interface Support**: CLI (`/skill-name`), GUI, API, and AI assistant integration

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
- **Interface Standards**: Consistent patterns across all client types

## 🚀 Quick Start

### Prerequisites
- **Go 1.25+** (for backend services and Temporal workflows)
- **Node.js 16+** (for frontend React/TypeScript application)  
- **Docker & Docker Compose** (for infrastructure services)
- **Make** (optional, for build automation)

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
**Via Web UI (Skills):**
1. Navigate to `http://localhost:3000/skills`
2. Browse available skills (compliance-check, security-scan, cost-analysis, workflow-management)
3. Click "Execute" on any skill with arguments
4. Monitor execution results and status

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
temporal-agents skill list

# Get detailed skill information
temporal-agents skill info compliance-check

# Invoke a skill with arguments
temporal-agents skill invoke compliance-check vm-web-server-001 SOC2 high

# Use slash syntax for skills (Claude/Codex compatible)
temporal-agents skill invoke /compliance-check vm-web-server-001
```

### Interactive Mode
```bash
temporal-agents --interactive
```
Available commands:
- `skill <name> [args]` - Invoke skills with `/skill-name` syntax
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

# Skills management
GET /api/skills                           # List all skills
GET /api/skills/invocable                 # List user-invocable skills
GET /api/skills/{skill_name}              # Get skill details
POST /api/skills/{skill_name}/execute     # Execute skill with arguments
POST /api/skills/discover                 # Trigger skill discovery

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

## � Workflow Templates

The `examples/workflow-templates.yaml` file contains pre-defined templates for:

- **Security Compliance Scan**: Comprehensive security and compliance validation
- **Multi-Cloud Compliance**: Cross-cloud compliance validation
- **Continuous Compliance Monitoring**: Automated ongoing compliance checks
- **Incident Response Compliance**: Compliance-focused incident response
- **Vendor Risk Assessment**: Automated vendor risk evaluation

## 🤖 Agent Types

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

## ⚙️ Configuration

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

```
repo/
├── AGENTS.md              # Agent behavior rules
├── SKILL.md               # AI skill definitions
├── .agents/               # Skills directory structure
│   └── skills/            # Skill definitions (Claude/Codex compatible)
│       ├── compliance-check/
│       │   ├── SKILL.md          # Compliance checking skill
│       │   ├── scripts/          # Supporting scripts
│       │   └── templates/        # Output templates
│       ├── security-scan/        # Security analysis skill
│       ├── cost-analysis/        # Cost optimization skill
│       └── workflow-management/  # Workflow control skill
├── backend/               # Go Temporal workflows
│   ├── activities/        # AI agent activities
│   ├── skills/            # Skills system implementation
│   │   ├── skill.go       # Skill parsing and management
│   │   └── service.go     # HTTP API for skills
│   ├── mcp/              # MCP server implementation
│   └── main.go           # Multi-interface server
├── frontend/              # React/TypeScript Backstage
│   ├── src/components/
│   │   ├── SkillsManagement.tsx  # Skills UI component
│   │   └── AgentManagement.tsx   # Workflow management
│   └── plugins/          # Temporal integration plugin
├── cli/                  # Command-line interface with skills
├── tools/                # Tool safety configurations
│   ├── bash.yaml         # Bash command restrictions
│   ├── git.yaml          # Git operation limits
│   ├── kubectl.yaml      # Kubernetes inspection
│   ├── terraform.yaml    # Infrastructure planning
│   └── docker.yaml       # Container management
├── docs/
│   ├── skills-system-guide.md    # Complete skills documentation
│   ├── comprehensive-interfaces-guide.md
│   └── claude-and-codex-skills.txt
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
- **[docs/skills-system-guide.md](docs/skills-system-guide.md)**: Complete skills system documentation (Claude/Codex compatible)
- **[IMPLEMENTATION_SUMMARY.md](IMPLEMENTATION_SUMMARY.md)**: Technical implementation details
- **[docs/comprehensive-interfaces-guide.md](docs/comprehensive-interfaces-guide.md)**: Complete interfaces documentation (SKILL.md, AGENTS.md, MCP, APIs, CLIs, GUIs)
- **[docs/claude-and-codex-skills.txt](docs/claude-and-codex-skills.txt)**: AI assistant skills and workflows reference
- **[docs/](docs/)**: Additional guides and specifications

## 🤝 Contributing

This is an AI-native engineering platform designed for:
- **Safe experimentation** with AI agents and workflows
- **Multi-interface development** supporting various integration patterns
- **Enterprise-grade orchestration** with durable execution guarantees
- **Regulatory compliance** built into the core architecture
- **AI assistant integration** supporting Claude Skills, Codex Skills, and Custom GPTs

### Interface Development
- **SKILL.md**: Create reusable workflow capabilities for AI assistants
- **AGENTS.md**: Define agent behavior rules and safety constraints
- **MCP Integration**: Implement standardized Model Context Protocol interfaces
- **Multi-Client Support**: Build consistent experiences across CLI, GUI, API, and web interfaces

### Contribution Areas
- **AI Agent Workflows**: New orchestration patterns and agent collaborations
- **Interface Implementations**: Additional client types and protocol support
- **Compliance Frameworks**: New regulatory standards and validation rules
- **Infrastructure Emulation**: Additional cloud providers and services
- **Documentation**: Guides and examples for all interface types

## 📄 License

[GNU Affero General Public License v3.0 or later](LICENSE)

---

**Open Source**: https://github.com/lloydchang/sandbox-backstage-temporal

--

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
