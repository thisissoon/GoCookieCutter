#
# Application Makefile
#

VERSION			?= $(shell git rev-parse --short HEAD)
BUILD_DIR		?= ./.build
BUILD_NAME		?= {{ cookiecutter.project_name }}
BUILD_FLAGS		?= -ldflags "-X {{ cookiecutter.project_name }}/config.version=${VERSION}"
DOCKER_IMAGE	?= {{ cookiecutter.docker_image_name }}
DOCKER_TAG		?= latest
TEST_VERBOSE	?= 0
TEST_COVERAGE	?= 0

#
# Testing
# Redis is required to run the test suite against
#

test: deps
ifeq ($(TEST_VERBOSE),1)
ifeq ($(TEST_COVERAGE),1)
	./test.sh -v --html
else
	./test.sh -v
endif
else
ifeq ($(TEST_COVERAGE),1)
	./test.sh --html
else
	./test.sh
endif
endif

#
# Build all binaries
#

all: deps test build_linux build_osx

deps:
	govendor sync

#
# Linux Build
#

build_linux: build_linux64

build_linux64:
	mkdir -p ${BUILD_DIR}/linux_amd64
	GOOS=linux GOARCH=amd64 go build -o ${BUILD_DIR}/linux_amd64/$(BUILD_NAME) ${BUILD_FLAGS}

#
# OSX Build
#

build_osx: build_osx64

build_osx64:
	mkdir -p ${BUILD_DIR}/darwin_amd64
	GOOS=darwin GOARCH=amd64 go build -o ${BUILD_DIR}/darwin_amd64/$(BUILD_NAME) ${BUILD_FLAGS}

#
# Docker Images
#

image:
	docker build --build-arg BUILD_DIR=$(BUILD_DIR)/linux_amd64 --build-arg BUILD_NAME=$(BUILD_NAME) --force-rm -t $(DOCKER_IMAGE):$(DOCKER_TAG) .
