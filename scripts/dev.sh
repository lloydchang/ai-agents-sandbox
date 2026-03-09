#!/bin/bash
set -e
# Start Temporal server & Backstage dev server
docker-compose -f backend/docker-compose.yml up -d
cd frontend
yarn start
