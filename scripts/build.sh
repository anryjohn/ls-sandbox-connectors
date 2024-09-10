#!/usr/bin/env bash

set -euxo pipefail

mkdir -p /opt/gomodcache
export GOMODCACHE=/opt/gomodcache
tar -C /opt/gomodcache -xf /src/build/gomodcache.tar
cp /opt/gomodcache/tidied/go.{mod,sum} ./

SVC_PKG="$(go list ./${SERVICE_DIR})" \
GO_LD_FLAGS="-X ${SVC_PKG}/version.Version=${VERSION} -extldflags '-static'"

CGO_ENABLED=1 GOOS=linux CGO_LDFLAGS_ALLOW="-Wl,--no-as-needed" \
    go build -a \
    -installsuffix "${GO_BUILD_TAGS}" \
    -tags "${GO_BUILD_TAGS}" \
    -ldflags "${GO_LD_FLAGS}" \
    -o app \
    "./${SERVICE_DIR}"
