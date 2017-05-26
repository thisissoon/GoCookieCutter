# Lightweight Hardended Linux Distro
FROM alpine:{{cookiecutter.alpine_version|default('3.5')}}

# Update and Install OS level packages
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

# Default build arguments
ARG BINLOC=./{{cookiecutter.project_name}}.linux-amd64
ARG BINDEST=/usr/local/bin/{{cookiecutter.project_name}}

# Copy Binary
COPY ${BINLOC} ${BINDEST}

# Volumes
VOLUME ["/etc/{{cookiecutter.project_name}}", "/var/log/{{cookiecutter.project_name}}"]

# Set our Application Entrypoint
ENTRYPOINT ["{{cookiecutter.project_name}}"]
