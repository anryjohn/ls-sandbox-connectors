#!/usr/bin/env bash

set -exuo pipefail

docker exec -i sandbox_postgres psql -U lutheran -d lutherize "$@"
