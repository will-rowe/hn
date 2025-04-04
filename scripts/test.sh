#!/usr/bin/env bash
set -e

echo "linting"
go vet ./...
buf lint

echo "running tests"
go test ./... -v

echo "Done!"