#!/usr/bin/env bash

set -o xtrace
set -o errexit
set -o nounset
set -o pipefail

tries=busted

for i in `seq 1 1`
do
  if make up
  then
    docker logs fnb-sandbox-peer0-1 &>/tmp/.luther_log
    if grep -E -q 'level=info msg="Handler found" endpoint=default op=Handle' /tmp/.luther_log
    then
      continue
    else
      tries="$i"
      break
    fi
  fi
done

if [ "$tries" != "busted" ]
then
  docker logs fnb-sandbox-peer0-1
  echo "+OK (tries=""$tries"")"
fi
