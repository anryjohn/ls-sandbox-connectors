#!/usr/bin/env bash

set -o xtrace
set -o errexit
set -o nounset
set -o pipefail

GOPATH=/root/go

if [ -f "$GOPATH"/clean ]
then
  go clean -cache -modcache
  go mod tidy
  rm -f "$GOPATH"/clean
fi

go mod tidy
cd ./connectorhub
go run . start
