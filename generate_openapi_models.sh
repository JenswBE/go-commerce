#!/bin/bash
# Bash strict mode: http://redsymbol.net/articles/unofficial-bash-strict-mode/
set -euo pipefail

# Run linter
docker run --user ${UID} --rm -v "$(pwd)/docs:/data" \
-e "NO_UPDATE_NOTIFIER=true" \
jamescooke/openapi-validator \
--errors_only \
--verbose \
openapi.yml

# Generate models
docker run --user ${UID} --rm -v "$(pwd):/local" \
-e "GO_POST_PROCESS_FILE=gofmt -s -w" \
openapitools/openapi-generator-cli generate \
--input-spec /local/docs/openapi.yml \
--generator-name go \
--output /local/api/openapi

# Remove unused files
find api/openapi -mindepth 1 -not -iname "model_*.go" -not -name utils.go -delete