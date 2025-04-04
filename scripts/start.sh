#!/usr/bin/env bash
set -euo pipefail

# Build and start
docker-compose up --build

# Run in background
docker-compose up -d

# Tear down
docker-compose down

# Check logs
docker-compose logs -f