# OpenAI Codex Skills

This directory contains OpenAI Codex-style skills for the AI Agents Sandbox. These skills follow the OpenAI Codex skills specification and can be invoked using the `$` prefix or through the `/skills` command.

## Available Skills

### 🔄 temporal-workflow
Manage and monitor Temporal workflows with AI agent orchestration.

**Usage:**
```bash
$temporal-workflow
```

**Capabilities:**
- Workflow creation and development
- Activity integration
- Testing and validation
- Monitoring and debugging

**Files:**
- `SKILL.md` - Main skill instructions
- `agents/openai.yaml` - UI configuration and dependencies
- `examples/basic-workflow.go` - Example workflow implementation
- `scripts/validate-workflow.sh` - Workflow validation script

### 📋 backstage-catalog
Manage Backstage software catalog and components.

**Usage:**
```bash
$backstage-catalog
```

**Capabilities:**
- Entity management and registration
- Catalog organization
- API documentation
- Component relationships

**Files:**
- `SKILL.md` - Main skill instructions
- `agents/openai.yaml` - UI configuration and dependencies
- `examples/payment-service.yaml` - Example component definition
- `templates/component-template.yaml` - Component creation template

### 🤖 ai-agent-orchestration
Orchestrate AI agents and manage agent workflows.

**Usage:**
```bash
$ai-agent-orchestration
```

**Capabilities:**
- Agent design and coordination
- Workflow orchestration
- Communication patterns
- Resource management

**Files:**
- `SKILL.md` - Main skill instructions
- `agents/openai.yaml` - UI configuration and dependencies

### 🔍 compliance-check
Perform automated compliance checks and policy validation.

**Usage:**
```bash
$compliance-check
```

**Capabilities:**
- Policy definition and validation
- Configuration analysis
- Automated checks
- Remediation guidance

**Files:**
- `SKILL.md` - Main skill instructions
- `agents/openai.yaml` - UI configuration and dependencies

### 💰 cost-optimization
Optimize resource costs and spending across infrastructure.

**Usage:**
```bash
$cost-optimization
```

**Capabilities:**
- Cost analysis and trends
- Resource assessment
- Optimization opportunities
- Implementation tracking

**Files:**
- `SKILL.md` - Main skill instructions
- `agents/openai.yaml` - UI configuration and dependencies

## Configuration

### agents/openai.yaml Structure
Each skill includes an `agents/openai.yaml` configuration file that defines:

```yaml
interface:
  display_name: "Display Name"
  short_description: "Brief description"
  icon_small: "./assets/icon.svg"
  icon_large: "./assets/logo.png"
  brand_color: "#COLOR"
  default_prompt: "Default prompt text"

policy:
  allow_implicit_invocation: true

dependencies:
  tools:
    - type: "mcp"
      value: "tool-name"
      description: "Tool description"
      transport: "streamable_http"
      url: "http://localhost:PORT"
```

## Skill Development

### Creating New Skills

1. Create skill directory:
```bash
mkdir -p .agents/skills/new-skill
```

2. Create SKILL.md:
```markdown
---
name: new-skill
description: What this skill does and when to use it
---

# Skill instructions here...
```

3. Create agents/openai.yaml for UI configuration
4. Add examples and supporting files as needed

### Skill Structure
```
skill-name/
├── SKILL.md                    # Main skill instructions (required)
├── agents/
│   └── openai.yaml            # UI configuration (optional)
├── examples/                   # Usage examples (optional)
├── scripts/                    # Helper scripts (optional)
├── templates/                  # Template files (optional)
└── assets/                     # Icons and images (optional)
```

## Integration

### Temporal Integration
Skills integrate with Temporal workflows through:
- Workflow definitions in `backend/workflows/`
- Activity implementations in `backend/activities/`
- API endpoints at `http://localhost:8081`

### Backstage Integration
Skills connect to Backstage through:
- Catalog entities in `catalog-info.yaml` files
- App configuration in `frontend/app-config.yaml`
- Plugin integrations in `frontend/plugins/`

## Best Practices

1. **Clear Descriptions**: Write detailed skill descriptions for implicit invocation
2. **Structured Content**: Use clear sections and code examples
3. **Error Handling**: Include comprehensive error handling guidance
4. **Integration Points**: Clearly define integration with other systems
5. **Examples**: Provide practical usage examples
6. **Testing**: Include validation and testing procedures

## Troubleshooting

### Skill Not Detected
- Restart Codex after adding new skills
- Check SKILL.md syntax
- Verify directory structure

### UI Configuration Issues
- Validate openai.yaml syntax
- Check icon file paths
- Verify URL endpoints

### Integration Problems
- Check backend service status
- Verify API endpoints
- Review network connectivity
