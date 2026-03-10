---
name: backstage-catalog
description: Manage Backstage software catalog and components. Use when working with Backstage entities, catalog configuration, component registration, or API documentation.
---

# Backstage Catalog Management

When working with Backstage catalog, follow these guidelines:

## 1. Entity Management
- Create and update catalog-info.yaml files
- Define proper component, API, and resource entities
- Use correct entity relationships and dependencies
- Follow Backstage entity specification

## 2. Component Registration
- Register new components in the catalog
- Define component ownership and lifecycle
- Add proper metadata and tags
- Set up component documentation

## 3. API Documentation
- Create API entities with OpenAPI specs
- Define API endpoints and schemas
- Document API contracts and usage
- Link APIs to consuming components

## 4. Catalog Organization
- Organize entities by domain and team
- Use proper naming conventions
- Define entity hierarchies
- Set up catalog filters and views

## Key File Locations
- Catalog config: `frontend/app-config.yaml`
- Entity definitions: `catalog-info.yaml` files
- Plugin configurations: `frontend/plugins/`

## Common Tasks
```yaml
# Component entity example
apiVersion: backstage.io/v1alpha1
kind: Component
metadata:
  name: my-component
  description: Component description
spec:
  type: service
  owner: team-name
  lifecycle: production
```

## Best Practices
- Use consistent entity naming
- Maintain proper ownership information
- Keep entity documentation up to date
- Validate entity YAML syntax
- Test catalog registration in development

## Integration Points
- Link components to Temporal workflows
- Connect APIs to documentation
- Associate resources with infrastructure
- Map teams to organizational structure
