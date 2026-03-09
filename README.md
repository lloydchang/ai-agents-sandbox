# Backstage + Temporal Sandbox

This repository provides a sandbox environment for experimenting with Backstage (TypeScript frontend) and Temporal (Go backend) integration, including Dockerized environments, sample workflows, Backstage plugin, and local development scripts.

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

## Setup Instructions

1. Clone the repository.

2. Ensure Go 1.21+, Node.js, Docker, Docker Compose are installed.

3. Create the Backstage app: `cd frontend && npx @backstage/create-app backstage-temporal`

4. Install dependencies: `cd frontend && yarn install`

5. Add the temporal-integration plugin: `npx @backstage/cli create-plugin --name temporal-integration`

6. Implement the plugin to trigger workflows via backend endpoints.

## Local Development

1. Run `./scripts/dev.sh` to start Temporal server and Backstage dev server.

2. Access Backstage at http://localhost:3000

3. Access Temporal UI at http://localhost:8080

4. Temporal worker runs on http://localhost:8081

## Example Workflow

Use the temporal-integration plugin in Backstage to trigger and monitor workflows.

## Repository Structure

- `/backend`: Temporal worker code (Go)

- `/frontend`: Backstage app + plugins (TypeScript)

- `/scripts`: Dev and build automation

- `/docs`: Documentation and notes

## Build

Run `./scripts/build.sh` to build Docker images.
