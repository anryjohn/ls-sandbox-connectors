#!/usr/bin/env bash

set -euxo pipefail

mkdir -p /opt/gomodcache
export GOMODCACHE=/opt/gomodcache

go mod tidy

mkdir /opt/gomodcache/tidied
cp go.{mod,sum} /opt/gomodcache/tidied

tar -cf /mnt/project_rel_dir/build/gomodcache.tar -C /opt/gomodcache -c .
