# SOON\_ Go Cookie Cutter

A Go project template that gives us our standard Go project setup.
Powered by [Cookiecutter](https://github.com/audreyr/cookiecutter).

## Features
 - Uses [dep](https://github.com/golang/dep) for dependancy management
 - Uses [viper](https://github.com/spf13/viper) for configuraiton
 - Uses [cobra](https://github.com/spf13/cobra) for CLI commands in `cmd` package
 - Uses [zerlog](https://github.com/rs/zerolog) for structured logging

## Optional Features
 - Configures GOPATH and installs deps
 - Dockerfile for building go binary and dockerfile with final binary
 - Option of GitlabCI

## Usage

1. Get [cookiecutter](https://github.com/audreyr/cookiecutter) via `pip` or `brew`
2. Get `dep`: https://github.com/golang/dep/releases
3. Use `cookiecutter`: `cookiecutter https://github.com/thisissoon/GoCookieCutter.git`
4. Fill in the information `cookiecutter` asks you
5. The project is now setup in `name/src/name` directory, `cd` into it
6. Read the `README.md`
