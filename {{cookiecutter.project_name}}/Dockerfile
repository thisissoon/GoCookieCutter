# Lightweight Hardended Linux Distro
FROM alpine:3.3

# Update and Install OS level packages
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

# Default build arguments
ARG BUILD_DIR=./_build/linux_amd64
ARG BUILD_NAME={{ cookiecutter.project_name|lower }}
ARG BUILD_DEST=/usr/local/bin/{{ cookiecutter.project_name|lower }}

# Copy Binary
COPY ${BUILD_DIR}/${BUILD_NAME} ${BUILD_DEST}

# Volumes
VOLUME ["/etc/{{ cookiecutter.client_name|lower }}/{{ cookiecutter.project_name|lower }}", "/var/log/{{ cookiecutter.client_name|lower }}/{{ cookiecutter.project_name|lower }}"]

# Set our Application Entrypoint
ENTRYPOINT ["{{ cookiecutter.project_name|lower }}"]
