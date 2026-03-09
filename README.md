# Backstage + Temporal Sandbox

This repository provides a complete sandbox environment for experimenting with Backstage (TypeScript frontend) and Temporal (Go backend) integration, including Dockerized environments, sample workflows, Backstage plugin, and local development scripts.

## Architecture

```

Backstage Frontend (TypeScript)
    |
    v
Temporal Backend (Go Worker) <-> Temporal Server <-> PostgreSQL
    |
    v
Workflows & Activities

```

## Quick Start

### Prerequisites
- Go 1.21+
- Node.js 18+
- Docker & Docker Compose

### 1. Start Infrastructure
```bash
cd backend && docker-compose up -d
```
This starts:
- PostgreSQL on port 5432
- Temporal Server on port 7233
- Temporal UI on port 8080

### 2. Start Backend Worker
```bash
cd backend && go run main.go
```
The worker will be available on port 8081 with endpoints:
- POST /workflow/start
- GET /workflow/status?id=<workflow_id>

### 3. Start Frontend
```bash
cd frontend && yarn start
```
Backstage will be available on http://localhost:3000

### 4. Test Integration
1. Navigate to http://localhost:3000/temporal
2. Click "Start HelloBackstage Workflow"
3. Monitor workflow status in the table
4. View detailed workflow execution in Temporal UI at http://localhost:8080

## Development Scripts

### Start Full Development Environment
```bash
./scripts/dev.sh
```
This automatically starts the infrastructure and frontend.

### Build Docker Images
```bash
./scripts/build.sh
```
Builds both backend and frontend Docker images.

## Example Workflow

The **HelloBackstageWorkflow** demonstrates:
1. **FetchDataActivity**: Retrieves data for a given name
2. **ProcessDataActivity**: Processes the fetched data
3. **Retry Policy**: Automatic retries with exponential backoff
4. **Logging**: Comprehensive activity logging

## Repository Structure

- `/backend`: Temporal worker code (Go)
  - `main.go`: Worker with HTTP endpoints
  - `docker-compose.yml`: Infrastructure setup
  - `Dockerfile`: Container build configuration
- `/frontend`: Backstage app + plugins (TypeScript)
  - `src/plugins/temporal-integration/`: Workflow management UI
  - `app-config.yaml`: Backstage configuration
- `/scripts`: Dev and build automation
- `/docs`: Documentation and notes

## API Endpoints

### Start Workflow
```bash
curl -X POST http://localhost:8081/workflow/start
```
Returns: Workflow ID

### Get Workflow Status
```bash
curl "http://localhost:8081/workflow/status?id=<workflow_id>"
```
Returns: Workflow status (RUNNING, COMPLETED, FAILED, etc.)

## Configuration

### Backend Configuration
- Temporal Server: localhost:7233
- Task Queue: backstage-task-queue
- HTTP Server: localhost:8081

### Frontend Configuration
- Backstage: http://localhost:3000
- Temporal UI: http://localhost:8080
- Backend API: http://localhost:8081

## Testing

### Backend Tests
```bash
cd backend && go test
```

### Frontend Tests
```bash
cd frontend && yarn test
```

## Troubleshooting

### Docker Issues
- Ensure Docker daemon is running
- Check port conflicts (5432, 7233, 8080, 8081, 3000)

### Backend Issues
- Verify Temporal server is running: `docker-compose ps`
- Check logs: `docker-compose logs temporal`

### Frontend Issues
- Clear node_modules and reinstall: `rm -rf node_modules && yarn install`
- Check configuration in app-config.yaml

## Extensions

This sandbox is designed for experimentation and can be extended with:

- **AI Agent Integration**: Add intelligent workflow orchestration
- **Multi-cloud Hooks**: AWS Proton, Azure Foundry integrations
- **GitHub Actions CI**: Automated testing and deployment
- **Additional Workflows**: Complex business logic examples
- **Enhanced UI**: Advanced workflow visualization and management

## License

```
repo-root/                      <- AGPLv3 (overall repo)
├── LICENSE                     <- AGPLv3
├── frontend/                   <- Backstage code
│   ├── LICENSE                 <- Apache 2.0
├── backend/                    <- Temporal code
│   ├── LICENSE                 <- MIT
```

