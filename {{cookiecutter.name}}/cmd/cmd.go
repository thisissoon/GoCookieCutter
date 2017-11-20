package cmd

import (
	"fmt"
	"time"

	"{{cookiecutter.name}}/config"
	"{{cookiecutter.name}}/log"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// New constructs a new CLI interface for execution
func New() *cobra.Command {
	cmd := &cobra.Command{
		PersistentPreRun: func(*cobra.Command, []string) {
			if err := config.FromFile(); err != nil {
				log.Default().
					Error().
					Err(err).
					Msg("failed to read configuration file")
			}
		},
		Run: func(*cobra.Command, []string) {
			fmt.Println("{{cookiecutter.name}}")
		},
	}
	// Global flags
	pflags := cmd.PersistentFlags()
	pflags.StringP("config-file", "c", "", "path to configuration file")
	pflags.String("log-format", "", "log format [console|json]")
	// Bind flags to config options
	config.BindPFlags(map[string]*pflag.Flag{
		config.CONFIG_PATH_KEY: pflags.Lookup("config-file"),
		config.LOG_FORMAT_KEY:  pflags.Lookup("log-format"),
	})
	// Add sub commands
	cmd.AddCommand(version())
	return cmd
}

// version returns a command for printing the current version
func version() *cobra.Command {
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
