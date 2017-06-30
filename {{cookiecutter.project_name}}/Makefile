# Build Variables
BUILD_TIME      ?= $(shell date +%s)
BUILD_VERSION   ?= $(shell git rev-parse HEAD)

# Docker Configuration
DOCKER_IMAGE    ?= {{cookiecutter.docker_image}}
DOCKER_TAG      ?= latest

# Go Compilation Flags
GOOUTDIR        ?= .
GOOS            ?=
GOARCH          ?=
CGO_ENABLED     ?= 0

# Bin Name
BIN_NAME        ?= {{cookiecutter.project_name}}
BIN_SUFFIX      ?=
ifneq ($(GOOS),)
ifneq ($(GOARCH),)
BIN_SUFFIX      = .$(GOOS)-$(GOARCH)
endif
endif

# LDFlags
BUILD_TIME_LDFLAG       ?= -X {{cookiecutter.project_name}}/build.timestamp=$(BUILD_TIME)
BUILD_VERSION_LDFLAG    ?= -X {{cookiecutter.project_name}}/build.version=$(BUILD_VERSION)
LDFLAGS                 ?= "$(BUILD_TIME_LDFLAG) $(BUILD_VERSION_LDFLAG)"

.PHONY: build

build:
	GOOS=$(GOOS) \
	GOARCH=$(GOARCH) \
	CGO_ENABLED=$(CGO_ENABLED) \
	go build \
		-v \
		-ldflags $(LDFLAGS) \
		-o "$(GOOUTDIR)/$(BIN_NAME)$(BIN_SUFFIX)"

run:
	go run main.go

test:
	go test -race -v -cover $(shell go list ./... | grep -v ./vendor/)

coverage:
	./coverage.sh
	go tool cover -html=coverage.txt

image:
	docker build --force-rm -t $(DOCKER_IMAGE):$(DOCKER_TAG) .

apib:
	docker run -v $(shell pwd)/docs:/docs -t christianbladescb/aglio:latest -i /docs/{{cookiecutter.project_name}}.apib -o /docs/{{cookiecutter.project_name}}.html

godoc:
	godoc -http=:9000
