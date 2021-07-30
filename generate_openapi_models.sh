#!/bin/bash
# Bash strict mode: http://redsymbol.net/articles/unofficial-bash-strict-mode/
set -euo pipefail

# Run linter
docker run --user ${UID} --rm -v "$(pwd)/docs:/data" jamescooke/openapi-validator -ev openapi.yml

# Generate models
docker run --user ${UID} --rm -v "$(pwd):/local" \
openapitools/openapi-generator-cli generate \
-i /local/docs/openapi.yml \
-g go \
-o /local/api/openapi

# Remove unused files
find api/openapi -mindepth 1 -not -iname "model_*.go" -not -name utils.go -delete

# Format generated files
gofmt -s -w ./api/openapi