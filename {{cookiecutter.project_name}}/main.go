package main

import (
	"{{cookiecutter.project_name}}/cli"
	"{{cookiecutter.project_name}}/log"
)

func main() {
	if err := cli.Execute(); err != nil {
		log.WithError(err).Error("application execution error")
	}
}
