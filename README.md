# AI Agents' Sandbox - Enterprise Edition

[![License: AGPL v3](https://img.shields.io/badge/License-AGPL%20v3-blue.svg)](LICENSE)
[![GitHub](https://img.shields.io/badge/GitHub-View_Repository-black?logo=github)](https://github.com/lloydchang/ai-agents-sandbox)
[![Production Ready](https://img.shields.io/badge/Production-Ready-brightgreen.svg)](https://github.com/lloydchang/ai-agents-sandbox)
[![7/7 Integrations](https://img.shields.io/badge/Integrations-7%2F7-brightgreen.svg)](https://github.com/lloydchang/ai-agents-sandbox)

🚀 **Enterprise-Grade AI Agent Platform with Complete Repository Integration** 🚀

A bleeding-edge laboratory for governed AI agent orchestration, enhanced with **7 production-ready integrations** from leading AI repositories.

## What is AI Agents Sandbox?

`ai-agents-sandbox` is a comprehensive, enterprise-ready platform that serves as both a **playground for multiple AI agents** and a **secure walled garden** for experimentation. It provides:

- **30 Specialized Skills** - Infrastructure provisioning, compliance scanning, cost optimization, and more
- **Multi-Agent Orchestration** - Coordinate complex workflows
- **Safe Execution Environment** - Isolated tool boundaries prevent unintended consequences
- **Human-in-the-Loop Controls** - Critical decisions require human approval
- **Multiple Interfaces** - REST APIs, CLI, WebMCP, and direct AI assistant integration

## 🎯 Repository Integration

Successfully integrated **7 high-value repositories** to create a complete AI agent ecosystem:

### Phase 1 (High Priority) - 3/3 Complete ✅
1. **MCP Tool Support** (temporal-ai-agent) - Goal-based agents, multi-agent workflows
2. **RAG AI Plugin Suite** (roadie-backstage-plugins) - Interactive chat, source attribution
3. **ReAct Patterns** (durable-react-agent-gemini) - Structured reasoning, thought-action-observation

### Phase 2 (Medium Priority) - 4/4 Complete ✅
4. **Research Workflows** (ai-iceberg-demo) - Multi-agent analysis, knowledge graphs
5. **AWS Bedrock Integration** (aws-samples) - Claude, Titan, Jurassic models
6. **WebSocket Real-Time Updates** (gorilla/websocket) - Live monitoring, event streaming
7. **Multi-Model AI** (spring-projects/spring-ai) - Intelligent selection, ensemble methods

**Results:** 100% integration success rate, 35+ activities, 25+ API endpoints, production-ready architecture.

## 🚀 Quick Start

### Prerequisites
- Go 1.25+
- Node.js 16+
- Docker & Docker Compose

### Local Development
```bash
# Clone and setup
git clone https://github.com/lloydchang/ai-agents-sandbox.git
cd ai-agents-sandbox

# Start everything
./scripts/dev.sh

# Access interfaces
# Frontend: http://localhost:3000
# Backend API: http://localhost:8081
# Temporal UI: http://localhost:8080
```

### Development Setup & Validation

For comprehensive development and testing, run these validation steps:

```bash
# 1. Validate environment and skill suite integrity
./bootstrap.sh

# 2. Run skill evaluations and tests
python3 eval/run_evals.py

# 3. Start development environment
./scripts/dev.sh
```

**What these do:**
- **`bootstrap.sh`** - Validates prerequisites, checks all 28 skills exist, ensures environment readiness
- **`run_evals.py`** - Tests skill definitions, validates SKILL.md formats, checks trigger phrases
- **`./scripts/dev.sh`** - Starts all services (Temporal, PostgreSQL, backend, frontend)

### Try It Out
```bash
# CLI example
./cli skill invoke /compliance-check vm-web-server-001 SOC2

# REST API example
curl -X POST http://localhost:8081/api/skills/compliance-check/execute \
  -H "Authorization: Bearer $AI_AGENTS_API_KEY" \
  -d '{"targetResource": "vm-web-server-001", "complianceType": "SOC2"}'
```

## 📚 Documentation

Comprehensive documentation is organized in the [`docs/`](./docs/) directory:

### 🏃‍♂️ [User Guide](./docs/user-guide/)
- **[Getting Started](./docs/user-guide/getting-started.md)** - Setup and basic usage
- **[Skills Reference](./docs/user-guide/skills-reference.md)** - All 28 skills and workflows
- **[Troubleshooting](./docs/user-guide/troubleshooting.md)** - Common issues and solutions

### 🛠️ [Developer Guide](./docs/developer-guide/)
- **[Agent Behavior](./docs/developer-guide/agent-behavior.md)** - Governance and rules
- **[Operational Procedures](./docs/developer-guide/operational-procedures.md)** - Workflow execution
- **[Skills API](./docs/developer-guide/skills-api.md)** - Technical skill integration
- **[Implementation](./docs/developer-guide/implementation.md)** - Architecture and deployment
- **[Extending](./docs/developer-guide/extending.md)** - Adding new skills and integrations

### 📖 [Reference](./docs/reference/)
- **[CLI Commands](./docs/reference/cli-commands.md)** - Complete command reference
- **[API Reference](./docs/reference/api-reference.md)** - REST API documentation
- **[Configuration](./docs/reference/configuration.md)** - All configuration options

## 🔑 Key Features

### Safe Execution Environment
- **Infrastructure Emulation** - Test against simulated AWS, Azure, GCP resources
- **Strict Tool Boundaries** - Configurable blast radius for every agent action
- **Human-in-the-Loop** - Built-in workflow pauses for critical decisions

### AI Agent Orchestration
- **Durable Execution** - Temporal ensures workflows survive crashes and timeouts
- **Skill System** - Auto-discovery of tools from `.agents/skills/` directory
- **Multi-Interface Support** - REST APIs, MCP server, CLI, WebMCP client, Backstage GUI

### Enterprise Ready
- **Comprehensive Skills** - 28 skills covering infrastructure, operations, security, compliance
- **Multi-Agent Coordination** - Parallel and sequential agent execution
- **Audit Trails** - Complete logging and compliance reporting
- **Production Architecture** - Dockerized deployment and monitoring

## 🏗️ Architecture Overview

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   AI Assistants │    │   Coordination  │    │   Execution     │
│   (Claude/GPT)  │───▶│   Layer         │───▶│   Layer         │
│                 │    │   (AGENTS.md)   │    │   (Temporal)    │
└─────────────────┘    └─────────────────┘    └─────────────────┘
         │                       │                       │
         ▼                       ▼                       ▼
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   Interfaces    │    │   Skills        │    │   Data Layer    │
│   (REST/MCP/CLI)│    │   (28 skills)   │    │   (PostgreSQL)  │
└─────────────────┘    └─────────────────┘    └─────────────────┘
```

## 🤝 Contributing

This sandbox is built for safe, bleeding-edge experimentation. Contributions welcome:

- **New AI Skills** - Add to `.agents/skills/` with SKILL.md specifications
- **Interface Development** - Extend WebMCP, add new integrations
- **Advanced Orchestration** - New multi-agent collaboration patterns
- **Compliance Frameworks** - Additional regulatory standards

See **[Extending](./docs/developer-guide/extending.md)** for detailed contribution guidelines.

## 📋 Example Use Cases

- **Tenant Onboarding**: "Onboard enterprise customer in Azure East US"
- **Compliance Automation**: Weekly SOC2 scans across all infrastructure
- **Incident Response**: "P0/P1 — service returning 503s"
- **Cost Optimization**: Monthly cloud spend analysis and recommendations
- **Security Monitoring**: Real-time vulnerability scanning and alerting

## 🛡️ Governance & Safety

All agent operations follow strict governance rules defined in **[Agent Behavior](./docs/developer-guide/agent-behavior.md)**:

- **Safety First**: Never execute destructive operations without approval
- **Audit Trail**: All actions logged and traceable
- **Human Oversight**: Critical decisions require human review
- **Idempotency**: Operations safe to retry

## 📈 Monitoring & Observability

- **Temporal UI** (`http://localhost:8080`) - Deep workflow execution visibility
- **Agent Dashboard** (`http://localhost:3000/temporal`) - Visual workflow builder
- **Audit Trails** - Complete history of all agent actions
- **Metrics & Logging** - Comprehensive monitoring and alerting

## 🎯 Why This Name?

The name "AI Agents Sandbox" deliberately carries multiple meanings:

- **Sandbox**: Safe, isolated environment for experimentation
- **Agents**: Multiple specialized AI agents working together
- **Playground**: High-velocity testbed for bleeding-edge patterns

Both "AI Agents Sandbox" (a sandbox for agents) and "AI Agent's Sandbox" (the agent's own environment) are accurate and intentional readings.

## 📄 License

[GNU Affero General Public License v3.0 or later](LICENSE)

---

**Ready to explore?** Start with **[Getting Started](./docs/user-guide/getting-started.md)** or dive into **[Skills Reference](./docs/user-guide/skills-reference.md)** to see what the agents can do.

---

## Architecture & The Sandbox Boundary

```
Backstage Frontend (TypeScript)
    │
    ▼
[ AI Assistants ] (Claude / Codex / GPT)
    │
    ▼
[ Coordination Layer ] (AGENTS.md & SKILL.md rules)
    │
    ▼
[ Unified Interface Fabric ]
├── REST API  |  MCP Server  |  WebMCP Client  |  CLI  |  Backstage GUI
    │
    ▼
==================== SANDBOX BOUNDARY ====================
[ Temporal Engine ] (Go Backend)
├── AI Agent Orchestration & Multi-Agent Collaboration
├── Human-in-the-Loop Workflow Pauses
├── MCP Integration (Model Context Protocol)
├── Safe Tool Execution Limits (bash, git, docker, kubectl)
└── Infrastructure Emulation Layer (AWS / Azure / GCP)
==========================================================
    │
    ▼
Temporal Server <──> PostgreSQL (Durable State & Workflows)
```

---

## The Laboratory Concept

This is not a production-ready framework — it is a bleeding-edge testbed. Use this repository to experiment with:

- **The Coordination Layer** — Define agent behaviors via `AGENTS.md` and structured `SKILL.md` schemas
- **Interface Interoperability** — Test how the same Skill behaves across CLI, REST API, WebMCP, and Backstage GUI
- **Tool Boundaries** — Push the limits of what agents can safely do using isolated execution configs that define the blast radius for every agent action
- **Multi-Agent Orchestration** — Run Security, Compliance, and Cost Optimization agents in parallel, with consensus-building and human-in-the-loop escalation

---

## Key Features

**The Sandbox (Safe Execution)**
- *Infrastructure Emulation* — Run compliance and security scans against simulated AWS, Azure, and GCP resources without touching your actual cloud accounts
- *Strict Tool Boundaries* — Granular config files (`tools/bash.yaml`, `tools/kubectl.yaml`, etc.) restrict agents to read-only, planning, or inspection operations
- *Human-in-the-Loop* — Built-in workflow pauses that require human authorization before agents execute critical decisions

**AI Agent Orchestration**
- *Durable Execution* — Powered by Temporal, ensuring long-running multi-agent workflows survive server crashes and timeouts
- *Skill System (Claude/Codex Compatible)* — Priority-based auto-discovery of tools defined via `SKILL.md` (YAML frontmatter + markdown) in `.agents/skills/`

**Multi-Interface Platform**
- Interact with agents however you want: REST APIs, MCP Server, WebMCP Client, CLI, Enhanced Backstage GUI, and direct AI assistant integration

---

## Quick Start

### Prerequisites
- Go 1.25+
- Node.js 16+
- Docker & Docker Compose
- Make *(optional)*

### 1. Start Infrastructure
```bash
cd backend && docker-compose up -d
```
Starts:
- PostgreSQL on port 5432
- Temporal Server on port 7233
- Temporal UI on port 8080

### 2. Start Backend
```bash
cd backend && go mod tidy && go run main.go
```
Launches the REST API on `:8081`, MCP server, WebMCP client interface, infrastructure emulator, and CLI-ready endpoints.

### 3. Start Frontend
```bash
cd frontend && yarn install && yarn start
```
Backstage with AI agent dashboard at `http://localhost:3000`.

Alternatively, start everything at once:
```bash
./scripts/dev.sh
```

### 4. Try It

**Web UI:**
1. Navigate to `http://localhost:3000/temporal`
2. Click "Start HelloBackstage Workflow"
3. Monitor workflow status in the table
4. View detailed execution in Temporal UI at `http://localhost:8080`

**CLI:**
```bash
./cli skill invoke /compliance-check vm-web-server-001 SOC2
```

**REST API:**
```bash
curl -X POST http://localhost:8081/api/skills/compliance-check/execute \
  -d '{"target": "vm-web-server-001", "standard": "SOC2"}'
```

---

## Example Workflow

**HelloBackstageWorkflow** is the simplest entry point, demonstrating the core pattern:

1. **FetchDataActivity** — retrieve data for a given name
2. **ProcessDataActivity** — process the fetched data
3. **Retry Policy** — automatic retries with exponential backoff
4. **Logging** — detailed activity logging throughout

From there, the sandbox scales up to multi-agent orchestration, compliance automation, and human-in-the-loop patterns. See `examples/workflow-templates.yaml` for pre-configured templates including Continuous Compliance Monitoring, Multi-Cloud Compliance, and Vendor Risk Assessment.

---

## CLI Usage

```bash
# Skill management
./cli skill list
./cli skill info compliance-check
./cli skill invoke /compliance-check vm-web-server-001 SOC2 high

# Workflow management
./cli workflow start ai-orchestration
./cli workflow start multi-agent
./cli workflow start compliance
./cli workflow status <workflow-id>
./cli workflow signal <workflow-id> approval true

# Interactive mode
./cli --interactive
```

Interactive commands: `skill`, `skills list`, `skills info`, `start`, `status`, `signal`, `health`

---

## API Reference

### Workflows
```
POST /workflow/start
POST /workflow/start-ai-orchestration
POST /workflow/start-multi-agent
POST /workflow/start-compliance
POST /workflow/start-human-in-loop
GET  /workflow/status?id=<workflow_id>
POST /workflow/signal/<workflow_id>
```

### Skills
```
GET  /api/skills
GET  /api/skills/invocable
GET  /api/skills/{skill_name}
POST /api/skills/{skill_name}/execute
POST /api/skills/discover
```

### Infrastructure Emulator
```
GET  /emulator/resources
GET  /emulator/resources/{id}
GET  /emulator/resources/{id}/security
GET  /emulator/resources/{id}/compliance
```

### MCP & Health
```
POST /mcp
GET  /mcp/resources
GET  /mcp/tools
GET  /health
GET  /metrics
```

The MCP server exposes standardized tools so local AI assistants (e.g., Claude Desktop) can safely interact with the sandbox. Tools: `start_compliance_workflow`, `get_workflow_status`, `signal_workflow`, `get_infrastructure_info`. Resources: `workflow_results`, `agent_capabilities`, `compliance_reports`. Protocol: JSON-RPC 2.0.

---

## Agents

**Security Agent** — Vulnerability scanning, threat detection, security posture analysis, access control validation.

**Compliance Agent** — SOC2/GDPR/HIPAA validation, control assessment, gap analysis, audit trail generation.

**Cost Optimization Agent** — Resource utilization analysis, right-sizing suggestions, reserved instance analysis.

---

## Core Governance

`AGENTS.md` and `SKILL.md` are the most important assets in the repository — the rules of the road for all agents in the lab.

**AGENTS.md** — Global behavioral constraints and repository-wide safety policies: allowed operations, forbidden directories, code generation standards, and interface-specific rules.

**SKILL.md** — Standardized input/output schemas for agentic tools, with usage examples, tool requirements, and error handling patterns.

**`/tools`** — Sandboxed execution configs that define the blast radius for agent actions:
- `bash.yaml` — permitted npm, git, docker, terraform commands
- `git.yaml` — read-only operations only (status, diff, log, blame)
- `kubectl.yaml` — cluster inspection, no destructive operations
- `terraform.yaml` — plan and safety operations only
- `docker.yaml` — container inspection with resource limits

---

## Repository Structure

```
ai-agents-sandbox/
├── AGENTS.md                   # Agent behavior rules
├── SKILL.md                    # AI skill definitions
├── .agents/skills/             # Skill definitions (Claude/Codex compatible)
│   ├── compliance-check/
│   ├── security-scan/
│   ├── cost-analysis/
│   └── workflow-management/
├── backend/                    # Go Temporal workflows
│   ├── activities/
│   ├── skills/                 # skill.go, service.go
│   ├── mcp/
│   ├── docker-compose.yml
│   ├── Dockerfile
│   └── main.go
├── frontend/                   # React/TypeScript Backstage
│   ├── src/
│   │   ├── components/
│   │   │   ├── SkillsManagement.tsx
│   │   │   └── AgentManagement.tsx
│   │   └── plugins/temporal-integration/
│   └── app-config.yaml
├── cli/
├── tools/                      # bash.yaml, git.yaml, kubectl.yaml,
│                               # terraform.yaml, docker.yaml
├── examples/workflow-templates.yaml
├── docs/
│   ├── skills-system-guide.md
│   ├── comprehensive-interfaces-guide.md
│   └── claude-and-codex-skills.txt
└── scripts/                    # dev.sh, build.sh, validate.sh
```

---

## Configuration

**Backend:**
- Temporal Server: `localhost:7233`
- Task Queue: `backstage-task-queue`
- HTTP Server: `localhost:8081`

**Frontend:**
- Backstage: `http://localhost:3000`
- Temporal UI: `http://localhost:8080`
- Backend API: `http://localhost:8081`

**Environment variables:**
- `TEMPORAL_HOST` — Temporal server address (default: `localhost:7233`)
- `DB_HOST` — PostgreSQL host (default: `localhost:5432`)
- `LOG_LEVEL` — Logging verbosity (default: `info`)

---

## Development Scripts

```bash
./scripts/dev.sh       # Start infrastructure + frontend
./scripts/build.sh     # Build Docker images for all components
./scripts/validate.sh  # Comprehensive environment validation
```

---

## Testing

```bash
cd backend  && go test ./...
cd backend  && go test -tags=integration ./...
cd cli      && go test ./...
cd frontend && yarn test
./scripts/validate.sh   # Full-stack integration tests
```

---

## Monitoring & Observability

**Temporal UI** (`http://localhost:8080`) — Deep dive into activity logs, AI decision reasoning, retry policies, and worker health.

**Agent Dashboard** (`http://localhost:3000/temporal`) — Visual workflow builder, agent execution times, compliance scores, and resource health metrics.

**Audit Trails** — Complete workflow history, agent decision logs, and compliance reports are stored for every execution.

---

## Troubleshooting

**Docker:**
```bash
docker-compose ps
docker-compose logs temporal
docker-compose logs postgres
docker-compose restart
```
Check for port conflicts: 5432, 7233, 8080, 8081, 3000.

**Backend:**
```bash
curl http://localhost:7233/health
docker-compose logs temporal | grep "Worker registered"
```

**Agent decision tracing:**
```bash
LOG_LEVEL=debug go run main.go
```

**Frontend:**
```bash
cd frontend && rm -rf node_modules && yarn install
cat frontend/app-config.yaml
```

**CLI:**
```bash
./cli health
./cli skill list
```

---

## Contributing & Extensions

This sandbox is built for safe, bleeding-edge experimentation. Contributions are welcome in:

**New AI Skills** — Create a directory under `.agents/skills/your-skill-name/`, add a `SKILL.md` with YAML frontmatter defining inputs, outputs, and tool requirements, and optionally add supporting scripts and templates. Skills are auto-discovered on next backend startup.

**Interface Development** — Expand WebMCP, add Jira/Slack/Discord hooks, or integrate GitHub Actions CI/CD.

**Advanced Orchestration** — New multi-agent collaboration patterns, additional cloud providers in the infrastructure emulator, or multi-cloud hooks (AWS Proton, Azure Foundry).

**Compliance Frameworks** — Additional regulatory standards beyond SOC2/GDPR/HIPAA and new validation rules.

---

## Documentation

| File | Description |
|------|-------------|
| [AGENTS.md](AGENTS.md) | Agent operating manual and behavior rules |
| [SKILL.md](SKILL.md) | Skill definitions and usage patterns |
| [docs/skills-system-guide.md](docs/skills-system-guide.md) | Full skills system reference |
| [docs/comprehensive-interfaces-guide.md](docs/comprehensive-interfaces-guide.md) | All interface types documented |
| [IMPLEMENTATION_SUMMARY.md](IMPLEMENTATION_SUMMARY.md) | Technical implementation details |

---

## Why This Name?

**Why `ai-agents-sandbox`?**

The name was chosen deliberately, and every word carries meaning.

**`agents` (plural)** — The repo contains multiple specialized AI agents: Security, Compliance, and Cost Optimization. They collaborate, build consensus, and orchestrate work together. Singular (`ai-agent-sandbox`) would imply a contained environment for testing a single chatbot. Plural signals variety, scope, and a multi-actor system — which is what this actually is.

**`sandbox`** — "Sandbox" does double duty here. In software, a sandbox is a safe, isolated environment for experimentation without production consequences. It also signals that this is a high-velocity testbed: patterns here are designed for rapid iteration, the specifications are evolving, and breaking changes are expected. If this were called a "platform" or "orchestrator," users would expect stability guarantees it doesn't claim to provide. "Sandbox" sets honest expectations while still being attractive to developers who want to work on the bleeding edge in a contained, safe way.

**The bonus ambiguity** — Some readers will parse it as *AI Agents Sandbox* (a sandbox for multiple AI agents), others as *AI Agent's Sandbox* (the agent's own playground). Both readings are accurate and intentional. The first describes what you do with it; the second describes what it is from the agent's perspective — a home environment governed by `AGENTS.md`, `SKILL.md`, and defined tool boundaries.

**Why not something else?**
- `temporal-ai-agents` — Temporal already uses "AI agents" heavily in their own marketing; the name would look like an official or affiliated project
- `agentic-sandbox` — An adjective without a noun; less searchable; dates faster
- `ai-agent-platform` — Oversells stability and completeness
- `lab` / `edge` / `experimental` — Overused, too vague, or implies instability rather than intentional exploration

`ai-agents-sandbox` is clean, searchable, accurate, and carries no branding conflicts.

---

## License

[GNU Affero General Public License v3.0 or later](LICENSE)
