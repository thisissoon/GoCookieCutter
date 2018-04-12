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
LDFLAGS ?= -d
LDFLAGS += -X {{cookiecutter.name}}/config.timestamp=$(BUILD_TIME)
LDFLAGS += -X {{cookiecutter.name}}/config.version=$(BUILD_VERSION)
LDFLAGS += -X {{cookiecutter.name}}/config.commit=$(BUILD_COMMIT)
# Go Build Flags
GOBUILD_FLAGS ?= -tags netgo -installsuffix netgo
GOBUILD_FLAGS += -installsuffix netgo
{% if cookiecutter.image != "n" %}
# Docker Configuration
DOCKER_IMAGE    ?= {{cookiecutter.image}}
DOCKER_TAG      ?= latest
{% endif %}
# Binary Name
BIN_NAME        ?= {{cookiecutter.name}}.$(BUILD_VERSION).$(GOOS)-$(GOARCH)
# Compress Binry
COMPRESS_BINARY ?= 0
# Verbose build output
GOBUILD_VERBOSE ?= 0

.PHONY: build testdata

# Run the go application
run:
	{{cookiecutter.name|upper}}_VERSION=$(BUILD_VERSION) \
	{{cookiecutter.name|upper}}_COMMIT=$(BUILD_COMMIT) \
	{{cookiecutter.name|upper}}_TIMESTAMP=$(BUILD_TIME) \
	go run ./main.go $(FLAGS)

# Sets an extra build variables before compilation
build-env:
ifeq ($(COMPRESS_BINARY),1)
	$(eval LDFLAGS += -a -w -s)
endif

# Go build fags
build-flags:
ifeq ($(GOBUILD_VERBOSE),1)
	$(eval GOBUILD_FLAGS += -v)
endif

# Build a binary
# Note: This builds a static binary and thus OS level dependant packages
# are required for build to work
build: build-env build-flags |
	CGO_ENABLED=0 \
	GOOS=$(GOOS) \
	GOARCH=$(GOARCH) \
	go build $(GOBUILD_FLAGS) \
		-ldflags "$(LDFLAGS)" \
		-o "$(BIN_NAME)" \
		./cmd/{{cookiecutter.name}}

{% if cookiecutter.image != "n" %}
# Build docker image
image:
	docker build \
		--force-rm \
		--build-arg BUILD_TIME=$(BUILD_TIME) \
		--build-arg BUILD_VERSION=$(BUILD_VERSION) \
		--build-arg BUILD_COMMIT=$(BUILD_COMMIT) \
		-t $(DOCKER_IMAGE):$(DOCKER_TAG) .

{% endif %}
# Run test suite
test:
ifeq ("$(wildcard $(shell which gocov))","")
	go get github.com/axw/gocov/gocov
endif
	{{cookiecutter.name|upper}}_LOG_FORMAT=discard \
		gocov test -v ./... | gocov report
