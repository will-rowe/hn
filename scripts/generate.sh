#!/usr/bin/env bash
set -e

echo "generating protobuf bindings"
buf generate --clean --include-imports

echo "generating Go mocks"
go install github.com/vektra/mockery/v2@v2.51.1
mockery --config mockery.yaml

echo "running go mod tidy"
go mod tidy

echo "Done!"