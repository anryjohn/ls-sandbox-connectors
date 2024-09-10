#!/usr/bin/env bash

set -o xtrace
set -o errexit
set -o nounset
set -o pipefail

cd ./build/_postfix/docker-postfix/demo
make down
