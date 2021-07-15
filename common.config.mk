# Copyright © 2021 Luther Systems, Ltd. All right reserved.

# config.mk
#
# General project configuration that configures make targets and tracks
# dependency versions.

# PROEJECT and VERSION are attached to docker images and phylum deployment
# artifacts created during the build process.
PROJECT=sandbox
VERSION=0.1.0-SNAPSHOT

# The makefiles use docker images to build artifacts in this project.  These
# variables configure the images used for builds.
BUILDENV_TAG=0.0.41

# These variables control the version numbers for parts of the LEIA platform
# and should be kept up-to-date to leverage the latest platform features.
SUBSTRATE_VERSION=2.160.0-fabric2-SNAPSHOT-4bdb08c
SHIROCLIENT_VERSION=2.159.0-fabric2-SNAPSHOT
SHIROTESTER_VERSION=2.159.0-fabric2-SNAPSHOT
CHAINCODE_VERSION=${SUBSTRATE_VERSION}
NETWORK_BUILDER_VERSION=2.159.0-fabric2-SNAPSHOT
MARTIN_VERSION=0.1.0-SNAPSHOT

# A golang module proxy server can greatly help speed up docker builds but the
# official proxy at https://proxy.golang.org only works for public modules.
# When your application needs private go module dependencies consider running a
# local athens-proxy server with an ssh/http configuration which can access
# private source repositories, otherwise set GOPRIVATE (or GONOPROXY and
# GONOSUMDB) if private modules are needed.  Though be aware that GOPRIVATE
# requires credentials (e.g. for github ssh) be available during builds which
# complicates things considerably.
# 		https://docs.gomods.io/
# 		https://golang.org/ref/mod#private-modules
GOPROXY ?= https://proxy.golang.org
GOPRIVATE ?=
GONOPROXY ?= ${GOPRIVATE}
GONOSUMDB ?= ${GOPRIVATE}
