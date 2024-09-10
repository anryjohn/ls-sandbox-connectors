#!/usr/bin/env bash

set -o xtrace
set -o errexit
set -o nounset
set -o pipefail

FILE_INNER_SH="$(readlink -f ./scripts/start-camunda-inner.sh)"

mkdir -p ./build/_camunda
cd ./build/_camunda
if false
then
[ -f ./camunda.zip ] || ( wget -O./camunda.zip.tmp https://github.com/camunda/camunda/archive/refs/heads/main.zip && mv ./camunda.zip{.tmp,} )
if [ ! -d ./camunda-main ]
then
  unzip camunda.zip
  (
    cd ./camunda-main
    patch -p1 <./../../../scripts/camunda2.patch
    docker build -t camunda-mine/operate -f operate.Dockerfile .
  )
fi
fi
[ -d ./camunda-platform ] || git clone https://github.com/camunda/camunda-platform.git
cd ./camunda-platform
git reset --hard HEAD
patch -p1 <./../../../scripts/camunda.patch
docker kill start-camunda-inner || true
docker compose -f docker-compose-core.yaml down -v || true
for i in zeebe elastic
do
  docker volume rm camunda-platform_"$i" || true
done
docker compose -f docker-compose-core.yaml up -d

while [ "$(docker inspect zeebe --format '{{ .State.Health.Status }}')" != "healthy" ]
do
  echo "... waiting for zeebe.healthy ..."
  sleep 1
done

#while [ "$(docker inspect camunda-platform-rest-api --format '{{ .State.Health.Status }}')" != "healthy" ]
#do
#  echo "... waiting for camunda-platform-rest-api.healthy ..."
#  sleep 1
#done

docker build -t start_camunda - <<'EOF'
FROM alpine:edge

RUN apk update && echo start_camunda.001
RUN apk add bash npm
RUN mkdir /root/project
WORKDIR /root/project
RUN npm init -y
RUN npm install -D typescript
RUN npm install -D ts-node
RUN npm install @camunda8/sdk
EOF

docker run --name start-camunda-inner --rm --detach --network byfn -v "$FILE_INNER_SH":/root/project/start-camunda-inner.sh start_camunda bash ./start-camunda-inner.sh
