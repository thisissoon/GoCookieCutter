package cli

import (
	"fmt"

	"{{cookiecutter.project_name}}/app"

	"github.com/spf13/cobra"
)

// Version CLI Command
// Prints the application build version and time
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print application version and build time",
	Run: func(*cobra.Command, []string) {
		bt := app.BuildTime().Format("Monday January 2 2006 at 15:04:05 MST")
		fmt.Println(fmt.Sprintf("Version: %s", app.Version()))
		fmt.Println(fmt.Sprintf("Build Time: %s", bt))
	},
}
