package main

import (
	"{{cookiecutter.name}}/config"
	"{{cookiecutter.name}}/log"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/rs/zerolog"
)

// Application entry point
func main() {
	{{cookiecutter.name}}Cmd().Execute()
}

// New constructs a new CLI interface for execution
func {{cookiecutter.name}}Cmd() *cobra.Command {
	logger := log.Defaults(zerolog.New(log.Writer()))
	cmd := &cobra.Command{
		Use:   "{{cookiecutter.name}}",
		Short: "Run the service",
		PersistentPreRun: func(*cobra.Command, []string) {
			if err := config.FromFile(); err != nil {
				logger.Error().Err(err).Msg("failed to read configuration file")
			}
		},
		Run:   {{cookiecutter.name}}Run,
	}
	// Global flags
	pflags := cmd.PersistentFlags()
	pflags.StringP("config-file", "c", "", "path to configuration file")
	pflags.String("log-format", "", "log format [console|json]")
	// Local Flags
	flags := cmd.Flags()
	flags.StringP("listen", "l", "", "server listen address")
	// Bind flags to config options
	config.BindPFlags(map[string]*pflag.Flag{
		config.CONFIG_PATH_KEY: pflags.Lookup("config-file"),
		config.LOG_FORMAT_KEY:  pflags.Lookup("log-format"),
	})
	// Add sub commands
	cmd.AddCommand(versionCmd())
	return cmd
}

// {{cookiecutter.name}}Run is executed when the CLI executes
// the {{cookiecutter.name}} command
func {{cookiecutter.name}}Run(cmd *cobra.Command, _ []string) {
	cmd.Help()
}
