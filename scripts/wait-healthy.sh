#!/usr/bin/env bash

set -exuo pipefail

while ! docker exec -i sandbox_postgres bash -c 'psql -U lutheran -d lutherize -l &>/dev/null'
do
  echo "waiting for postgres ..."
  sleep 1
done

docker exec -i sandbox_postgres bash -c 'psql -U postgres -d lutherize -c "GRANT CREATE ON SCHEMA public TO lutheran;"'
