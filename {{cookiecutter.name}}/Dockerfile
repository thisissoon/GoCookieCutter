# Stage 1 - Binary Build
# BUILD_X args should be passed at build time as docker build args
FROM golang:{{cookiecutter.go}}-alpine AS builder
ARG BUILD_TIME
ARG BUILD_VERSION
ARG BUILD_COMMIT
RUN apk update && apk add build-base libressl-dev
WORKDIR /go/src/{{cookiecutter.name}}
COPY ./ /go/src/{{cookiecutter.name}}
RUN BIN_NAME=bin make build

# Stage 2 - Final Image
# The application should be statically linked
FROM alpine:3.6
RUN apk update && apk add --no-cache ca-certificates && rm -rf /var/cache/apk/*
COPY --from=builder /go/src/{{cookiecutter.name}}/bin /usr/bin/{{cookiecutter.name}}
VOLUME ["/etc/{{cookiecutter.name}}"]
ENTRYPOINT ["{{cookiecutter.name}}"]
