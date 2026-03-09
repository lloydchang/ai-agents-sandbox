# Backstage-Temporal Repository Scaffold Plan

This plan will fully scaffold a GitHub repository integrating Backstage (TypeScript) and Temporal (Go) with Dockerized environments, sample workflows, plugin integration, and local development scripts.

## Phase 1: Repository Setup & Structure
- Initialize GitHub repository with proper README, LICENSE, and .gitignore
- Create directory structure: backend/, frontend/, scripts/, docs/
- Set up initial git commit and push

## Phase 2: Temporal Backend (Go)
- Initialize Go module with Temporal SDK
- Implement HelloBackstageWorkflow with FetchDataActivity and ProcessDataActivity
- Add logging and retry policies
- Create Dockerfile for the worker
- Set up docker-compose.yml with Temporal server, UI, and PostgreSQL
- Add HTTP endpoints: POST /workflow/start, GET /workflow/status

## Phase 3: Backstage Frontend (TypeScript)
- Create Backstage app using @backstage/create-app
- Develop temporal-integration plugin with:
  - Workflow trigger button
  - Workflow status table
  - Optional logs view
- Configure connection to Temporal backend (REST/gRPC)

## Phase 4: Integration & Testing
- Connect frontend plugin to backend endpoints
- Implement end-to-end workflow validation
- Test complete workflow execution flow

## Phase 5: Development Tooling
- Create /scripts/dev.sh for starting Temporal server and Backstage dev server
- Create /scripts/build.sh for building Docker images
- Add comprehensive documentation in README.md and /docs/

## Phase 6: Documentation & Extensions
- Update README with architecture diagram and setup instructions
- Add experimental notes and changelog to /docs/
- Prepare for optional extensions (AI agent integration, multi-cloud hooks, GitHub Actions CI)

## Technical Considerations
- Use Go 1.21-alpine for backend Docker image
- Implement proper error handling and logging
- Ensure Docker and Docker Compose compatibility
- Keep repository private initially per security best practices

## Notes
- Ensure Go, Node.js, Docker, Docker Compose are installed.
- Repo starts private; make public later if stable.
- Validate end-to-end workflow functionality.
