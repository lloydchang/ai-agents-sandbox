#!/bin/bash
# Start Temporal server & Backstage dev server
docker-compose -f ../backend/docker-compose.yml up -d
cd ../frontend
yarn dev
