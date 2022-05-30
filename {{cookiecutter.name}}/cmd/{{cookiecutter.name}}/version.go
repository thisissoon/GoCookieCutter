package main

import (
	"os"

	"github.com/spf13/cobra"

	"{{cookiecutter.module}}/internal/version"
)

// versionCmd returns a CLI command that when run prints
// the application build version, commit and time
func versionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Prints the build version",
		Run: func(*cobra.Command, []string) {
			version.Write(os.Stdout)
		},
	}
}
