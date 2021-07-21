#!/bin/bash
# Bash strict mode: http://redsymbol.net/articles/unofficial-bash-strict-mode/
set -euo pipefail

export USE_DC_HELPERS_INSTEAD_OF_DIRECTLY_CALLING_DOCKER_COMPOSE=true
docker-compose down