#!/bin/bash
# Bash strict mode: http://redsymbol.net/articles/unofficial-bash-strict-mode/
set -euo pipefail

# Get real user ID
USER_ID=${UID}
if [ -n ${SUDO_UID+x} ]; then
USER_ID=${SUDO_UID}
fi

# Run linter
docker run --pull always --user ${USER_ID:?} --rm -v "$(pwd)/docs:/data" \
-e "NO_UPDATE_NOTIFIER=true" \
jamescooke/openapi-validator:latest \
--errors_only \
--verbose \
openapi.yml

# Clean directory if exists
rm api/openapi/* || true

# Generate models
docker run --pull always --user ${USER_ID:?} --rm -v "$(pwd):/local" \
-e "GO_POST_PROCESS_FILE=gofmt -s -w" \
openapitools/openapi-generator-cli generate \
--input-spec /local/docs/openapi.yml \
--generator-name go \
--output /local/api/openapi \
--additional-properties enumClassPrefix=true

# Remove unused files
rm api/openapi/go.mod api/openapi/go.sum 
go mod tidy
