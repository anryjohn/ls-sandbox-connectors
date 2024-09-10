# Copyright © 2021 Luther Systems, Ltd. All right reserved.

# Makefile
#
# The primary project makefile that should be run from the root directory and is
# able to build and run the entire application.

PROJECT_REL_DIR=.
include ${PROJECT_REL_DIR}/common.mk
BUILD_IMAGE_PROJECT_DIR=/go/src/${PROJECT_PATH}

GO_SERVICE_PACKAGES=./portal/... ./phylum/...
GO_API_PACKAGES=./api/...
GO_PACKAGES=${GO_SERVICE_PACKAGES} ${GO_API_PACKAGES}

SUBSTRATEHCP_FILE ?= ${PWD}/${SUBSTRATE_PLUGIN_PLATFORM_TARGETED}

export SUBSTRATEHCP_FILE

.DEFAULT_GOAL := default
.PHONY: default
default: all

.PHONY: all push clean test

clean:
	rm -rf build

all: plugin
.PHONY: plugin plugin-linux plugin-darwin
plugin: ${SUBSTRATE_PLUGIN}

plugin-linux: ${SUBSTRATE_PLUGIN_LINUX}

plugin-darwin: ${SUBSTRATE_PLUGIN_DARWIN}

all: tests-api
.PHONY: tests-api
tests-api:
	cd tests && $(MAKE)
.PHONY: tests-api-clean
tests-api-clean:
	cd tests && $(MAKE) clean

all: api
.PHONY: api
api:
	cd api && $(MAKE)

.PHONY: apiclean
apiclean:
	cd api && $(MAKE) clean

all: phylum
.PHONY: phylum
phylum:
	cd phylum && $(MAKE)
test: phylumtest
.PHONY: phylumtest
phylumtest:
	cd phylum && $(MAKE) test
clean: phylumclean
.PHONY: phylumclean
phylumclean:
	cd phylum && $(MAKE) clean

all: portal
.PHONY: portal
portal: plugin
	cd ${SERVICE_DIR} && $(MAKE)
clean: portalclean
.PHONY: portalclean
portalclean:
	cd ${SERVICE_DIR} && $(MAKE) clean

.PHONY: fabric
all: fabric
fabric:
	cd fabric && $(MAKE)
.PHONY: fabricclean
clean: fabricclean
fabricclean:
	cd fabric && $(MAKE) clean

.PHONY: storage-up
storage-up:
	cd fabric && $(MAKE) up install init

.PHONY: storage-down
storage-down:
	-cd fabric && $(MAKE) down
	-${DOCKER} kill sandbox_ch
	-bash ./scripts/stop-postfix.sh

.PHONY: service-up
service-up: api portal
	./blockchain_compose.py local up -d
	bash ./scripts/start-postfix.sh
	bash ./scripts/start-camunda.sh
	${DOCKER} build -t luthersystems/mem-up-ready -f Dockerfile.mem-up-ready .
	${DOCKER_RUN} --rm -v "/var/run/docker.sock:/var/run/docker.sock" --network byfn luthersystems/mem-up-ready bash /opt/wait-healthy.sh
	${DOCKER_RUN} --rm --name sandbox_ch --detach --network byfn -v ${PWD}":"${PWD} -w ${PWD} -v /tmp/.luther/_go-ch:/root/go ${BUILD_IMAGE_GO_ALPINE}:${BUILDENV_TAG} bash ./scripts/start-ch.sh
	bash -c 'while ! docker logs sandbox_ch 2>&1 | grep -E -q "listen to fabric events"; do sleep 3; done'
	sleep 60
	make integration
	sleep 10
	-./psql.sh -c 'select * from summary;'
#docker logs fnb-sandbox-peer0-1

.PHONY: service-up-ch
service-up-ch: api
	-${DOCKER} kill sandbox_ch
	${DOCKER_RUN} --rm --name sandbox_ch --detach --network byfn -v ${PWD}":"${PWD} -w ${PWD} -v /tmp/.luther/_go-ch:/root/go ${BUILD_IMAGE_GO_ALPINE}:${BUILDENV_TAG} bash ./scripts/start-ch.sh
	docker logs -f sandbox_ch

.PHONY: tidy-ch
tidy-ch:
	touch /tmp/.luther/_go-ch/clean

.PHONY: service-down
service-down:
	-${DOCKER} kill sandbox_ch
	-./blockchain_compose.py local down

.PHONY: up
up: all service-down storage-down storage-up service-up
	@

.PHONY: down
down: explorer-down service-down storage-down
	@

.PHONY: init
init:
	-cd fabric && $(MAKE) init

.PHONY: upgrade
upgrade: all service-down init service-up
	@

.PHONY: mem-up
mem-up: all mem-down
	./blockchain_compose.py mem up -d
	${DOCKER} build -t luthersystems/mem-up-ready -f Dockerfile.mem-up-ready .
	${DOCKER_RUN} --rm -v "/var/run/docker.sock:/var/run/docker.sock" --network byfn luthersystems/mem-up-ready bash /opt/wait-healthy.sh
	-${DOCKER} kill sandbox_ch
	${DOCKER_RUN} --rm --name sandbox_ch --detach --network byfn -v ${PWD}":"${PWD} -w ${PWD} ${BUILD_IMAGE_GO_ALPINE}:${BUILDENV_TAG} bash ./scripts/start-ch.sh

.PHONY: mem-down
mem-down: explorer-down
	-./blockchain_compose.py mem down
	${DOCKER} kill sandbox_ch

# citest runs all tests within containers, as in CI.
.PHONY: citest
citest: plugin unit integrationcitest
	@

.PHONY: unit-portal
unit-portal:
	go test -v ./...

.PHONY: unit
unit: unit-portal unit-other
	@echo "all tests passed"

.PHONY: unit-other
unit-other: phylumtest
	@echo "phylum tests passed"


# NOTE:  The `citest` target manages creating/destroying a compose network.  To
# run tests repeatedly execute the `integration` target directly.
.PHONY: integrationcitest
# The `down` wouldn't execute without this syntax
integrationcitest:
	$(MAKE) up
	$(MAKE) integration
	$(MAKE) down

.PHONY: integration
integration:
	cd tests && $(MAKE) test-docker

.PHONY: repl
repl:
	cd phylum && $(MAKE) repl

# this target is called by git-hooks/pre-push. It's separated into its own target
# to allow us to update the git-hooks without having to reinstall the hook
# It generates postman artifacts and protobuf artifacts.
.PHONY: pre-push
pre-push:
	$(MAKE) tests-api
	cd api && $(MAKE)

.PHONY:
download: ${SUBSTRATE_PLUGIN}

.PHONY: print-export-path
print-export-path:
	@echo "export SUBSTRATEHCP_FILE=${SUBSTRATEHCP_FILE}"

${STATIC_PLUGINS_DUMMY}:
	${MKDIR_P} $(dir $@)
	./scripts/obtain-plugin.sh
	touch $@

${SUBSTRATE_PLUGIN}: ${STATIC_PLUGINS_DUMMY}
	@

.PHONY: explorer
explorer: explorer-up-clean

.PHONY: explorer-up
explorer-up:
	cd ${PROJECT_REL_DIR}/explorer && make up

.PHONY: explorer-up-clean
explorer-up-clean:
	cd ${PROJECT_REL_DIR}/explorer && make up-clean

.PHONY: explorer-down
explorer-down:
	cd ${PROJECT_REL_DIR}/explorer && make down

.PHONY: explorer-clean
explorer-clean:
	cd ${PROJECT_REL_DIR}/explorer && make down-clean

.PHONY: explorer-watch
explorer-watch:
	cd ${PROJECT_REL_DIR}/explorer && make watch
