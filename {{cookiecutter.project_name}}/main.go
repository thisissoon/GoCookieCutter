// Go Application Entrypoint

package main

import (
	"os"

	"{{ cookiecutter.project_name|lower }}/cli"
	"{{ cookiecutter.project_name|lower }}/logger"
)

func main() {
	if err := cli.Execute(); err != nil {
		logger.Error("Failed to start: %s", err)
		os.Exit(1)
	}
}
