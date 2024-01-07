#!/bin/bash
# Bash strict mode: http://redsymbol.net/articles/unofficial-bash-strict-mode/
set -euo pipefail

# Get real user ID
if [ -z "${SUDO_UID:-}" ]; then
    echo "Must be run with sudo ..."
    exit 1
fi

# Run linter
podman run --pull always --user ${SUDO_UID:?} --rm -v "$(pwd)/docs:/data:z" \
-e "NO_UPDATE_NOTIFIER=true" \
docker.io/ibmdevxsdk/openapi-validator:latest \
--errors-only \
openapi.yml

# Clean directory if exists
rm api/openapi/* || true

# Generate models
podman run --pull always --user ${SUDO_UID:?} --rm -v "$(pwd):/local:z" \
-e "GO_POST_PROCESS_FILE=gofmt -s -w" \
docker.io/openapitools/openapi-generator-cli generate \
--input-spec /local/docs/openapi.yml \
--generator-name go \
--output /local/api/openapi \
--additional-properties enumClassPrefix=true

# Remove unused files
find ./api/openapi ! -name '*.go' -delete || true
rm -rf ./api/openapi/test
go mod tidy
