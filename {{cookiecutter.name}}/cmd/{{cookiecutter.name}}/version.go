package main

import (
	"fmt"
	"time"

	"{{cookiecutter.name}}/config"

	"github.com/spf13/cobra"
)

// versionCmd returns a CLI command that when run prints
// the application build version, commit and time
func versionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Prints the build version",
		Run: func(*cobra.Command, []string) {
			fmt.Println("Version:", config.Version())
			fmt.Println("Commit:", config.Commit())
			fmt.Println("Built:", config.BuildTimestamp().Format(time.RFC1123))
		},
	}
}
