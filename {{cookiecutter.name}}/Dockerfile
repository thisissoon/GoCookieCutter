# Stage 1 - Binary Build
# BUILD_X args should be passed at build time as docker build args
FROM golang:{{cookiecutter.go}}-alpine AS builder
ARG BIN_VERSION
ARG GIT_COMMIT
ARG GIT_SHA
ARG GIT_TAG
ARG GIT_DIRTY
ENV BIN_OUTDIR=
ENV BIN_NAME={{cookiecutter.name}}
RUN apk update && apk add build-base git libressl-dev
WORKDIR /usr/src/{{cookiecutter.name}}
COPY ./ /usr/src/{{cookiecutter.name}}
RUN make static

# Stage 2 - Final Image
# The application should be statically linked
FROM alpine:latest
RUN apk update \
	&& apk add --no-cache ca-certificates \
	&& rm -rf /var/cache/apk/* \
    && addgroup {{cookiecutter.name}} \
	&& adduser -D -H -G {{cookiecutter.name}} {{cookiecutter.name}}
COPY --from=builder /usr/src/{{cookiecutter.name}} /usr/bin/{{cookiecutter.name}}
VOLUME ["/etc/{{cookiecutter.name}}"]
ENTRYPOINT ["{{cookiecutter.name}}"]
EXPOSE 5000
USER {{cookiecutter.name}}
