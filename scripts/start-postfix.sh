#!/usr/bin/env bash

set -o xtrace
set -o errexit
set -o nounset
set -o pipefail

mkdir -p ./build/_postfix
cd ./build/_postfix
[ -d ./docker-postfix ] || git clone https://github.com/mlan/docker-postfix.git
cd ./docker-postfix
git reset --hard HEAD
patch -p1 <./../../../scripts/postfix.patch
cd ./demo
make down
for i in app-atch app-conf app-spam auto db flt mta
do
  docker volume rm demo_"$i" || true
done
make up
