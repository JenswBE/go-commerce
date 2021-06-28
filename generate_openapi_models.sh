#!/bin/bash
# Bash strict mode: http://redsymbol.net/articles/unofficial-bash-strict-mode/
set -euo pipefail

# Expand YAML
docker run --rm -v "$(pwd):/local" python bash -c "pip install ruamel.yaml.cmd && yaml merge-expand /local/docs/openapi.yml /local/docs/openapi.yml.tmp"

# Generate models
docker run --user ${UID} --rm -v "$(pwd):/local" openapitools/openapi-generator-cli generate -i /local/docs/openapi.yml.tmp -g go -o /local/api/openapi

# Remove unused files
rm -f docs/openapi.yml.tmp
find api/openapi -mindepth 1 -not -iname "model_*.go" -not -name utils.go -delete

# Format generated files
gofmt -s -w ./api/openapi