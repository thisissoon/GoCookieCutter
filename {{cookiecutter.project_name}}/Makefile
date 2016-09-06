#
# Application Makefile
#

VERSION			?= $(shell git rev-parse HEAD)
BUILD_DIR		?= ./.build
BUILD_NAME		?= {{ cookiecutter.project_name }}
BUILD_FLAGS		?= -ldflags "-X {{ cookiecutter.project_name }}/config.version=${VERSION}"
DOCKER_IMAGE	?= registry.soon.build/crest/studentprocessorsvc
DOCKER_TAG		?= latest
TEST_ENVS 		?=
TEST_FLAGS 		?=

ifdef VERBOSE
ifeq ($(VERBOSE),1)
TEST_FLAGS += -v
endif
endif

#
# Testing
#

test:
	$(TEST_ENVS) ./test.sh$(TEST_FLAGS)

#
# Build all binaries
#

all: test build_linux build_osx

#
# Linux Build
#

build_linux: build_linux64

build_linux64:
	mkdir -p ${BUILD_DIR}/linux_amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BUILD_DIR}/linux_amd64/$(BUILD_NAME) ${BUILD_FLAGS}

#
# OSX Build
#

build_osx: build_osx64

build_osx64:
	mkdir -p ${BUILD_DIR}/darwin_amd64
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ${BUILD_DIR}/darwin_amd64/$(BUILD_NAME) ${BUILD_FLAGS}

#
# Docker Images
#

image:
	docker build --build-arg BUILD_DIR=$(BUILD_DIR)/linux_amd64 --build-arg BUILD_NAME=$(BUILD_NAME) --force-rm -t $(DOCKER_IMAGE):$(DOCKER_TAG) .
