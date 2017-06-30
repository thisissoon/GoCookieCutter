# Stage 1 - Binary Build
FROM golang:{{cookiecutter.go_version|default('1.8')}}-alpine AS builder
RUN apk update && apk add make git {{cookiecutter.onbuild_pkgs}}
WORKDIR /go/src/{{cookiecutter.project_name}}
COPY ./ /go/src/{{cookiecutter.project_name}}
ENV GOOS=linux GOARCH=amd64
RUN make build

# Stage 2 - Final Image
FROM alpine:{{cookiecutter.alpine_version|default('3.5')}}
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
COPY --from=builder /go/src/{{cookiecutter.project_name}}/{{cookiecutter.project_name}}.linux-amd64 /usr/bin/{{cookiecutter.project_name}}
VOLUME ["/etc/{{cookiecutter.project_name}}", "/var/log/{{cookiecutter.project_name}}"]
ENTRYPOINT ["{{cookiecutter.project_name}}"]
