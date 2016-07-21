// Version Sub Command

package cli

import (
	"fmt"
	"{{ cookiecutter.project_name|lower }}/config"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display the build version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Build Version:", config.Version())
	},
}
