# Compilation Flags
GOOS            ?= $(shell go env GOOS)
GOARCH          ?= $(shell go env GOARCH)
# Flags
FLAGS           ?=
# Build Vars
BUILD_TIME      ?= $(shell date +%s)
BUILD_VERSION   ?= $(shell head -n 1 VERSION | tr -d "\n")
BUILD_COMMIT    ?= $(shell git rev-parse HEAD)
# LDFlags
LDFLAGS ?= -d -a -w
LDFLAGS += -X {{cookiecutter.name}}/config.timestamp=$(BUILD_TIME)
LDFLAGS += -X {{cookiecutter.name}}/config.version=$(BUILD_VERSION)
LDFLAGS += -X {{cookiecutter.name}}/config.commit=$(BUILD_COMMIT)
# Docker Configuration
DOCKER_IMAGE    ?= {{cookiecutter.image}}
DOCKER_TAG      ?= latest
# Binary Name
BIN_NAME        ?= {{cookiecutter.name}}.$(BUILD_VERSION).$(GOOS)-$(GOARCH)

.PHONY: build testdata

# Run the go application
run:
	{{cookiecutter.name|upper}}_VERSION=$(BUILD_VERSION) \
	{{cookiecutter.name|upper}}_COMMIT=$(BUILD_COMMIT) \
	{{cookiecutter.name|upper}}_TIMESTAMP=$(BUILD_TIME) \
	go run ./main.go $(FLAGS)

# Build a binary
# Note: This builds a static binary and thus OS level dependant packages
# are required for build to work
build:
	CGO_ENABLED=0 \
	GOOS=$(GOOS) \
	GOARCH=$(GOARCH) \
	go build \
		-v \
		-tags netgo \
		-installsuffix netgo \
		-ldflags "$(LDFLAGS)" \
		-o "$(BIN_NAME)"

# Build docker image
image:
	docker build \
		--force-rm \
		--build-arg BUILD_TIME=$(BUILD_TIME) \
		--build-arg BUILD_VERSION=$(BUILD_VERSION) \
		--build-arg BUILD_COMMIT=$(BUILD_COMMIT) \
		-t $(DOCKER_IMAGE):$(DOCKER_TAG) .
