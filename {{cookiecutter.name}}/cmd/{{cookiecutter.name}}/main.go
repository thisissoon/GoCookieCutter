package main

import (
	"fmt"
	"io"
	"os"

	"{{cookiecutter.pkg}}/internal/config"
	"{{cookiecutter.pkg}}/internal/version"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Default logger
var log zerolog.Logger

// Application entry point
func main() {
	{{cookiecutter.name}}Cmd().Execute()
}

// New constructs a new CLI interface for execution
func {{cookiecutter.name}}Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "{{cookiecutter.name}}",
		Short: "Run the service",
		PersistentPreRun: func(*cobra.Command, []string) {
			// Setup default logger
			log = initLogger()
			// Init config
			if err := config.FromFile(); err != nil {
				log.Error().Err(err).Msg("failed to read configuration file")
			} else {
				log.Debug().Msg(fmt.Sprintf("using config file: %s", viper.ConfigFileUsed()))
			}
			// Reconfigure logger with config
			log = initLogger()
		},
		Run:   {{cookiecutter.name}}Run,
	}
	// Global flags
	pflags := cmd.PersistentFlags()
	{% if cookiecutter.project is not none -%}
	pflags.StringP("config", "c", "", "path to configuration file (default is $HOME/.config/{{cookiecutter.project}}/{{cookiecutter.name}}.toml)")
	{% else -%}
	pflags.StringP("config", "c", "", "path to configuration file (default is $HOME/.config/{{cookiecutter.name}}.toml)")
	{% endif -%}
	pflags.String("log-format", "", "log format [console|json] (default is json)")
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

// initLogger constructs a default logger from config
func initLogger() zerolog.Logger {
	var w io.Writer = os.Stdout
	switch config.LogFormat() {
	case "console", "terminal":
		w = zerolog.ConsoleWriter{
			Out: os.Stdout,
		}
	}
	return zerolog.New(w).With().Fields(map[string]interface{}{
		"version": version.Version,
		"app":     config.APP_NAME,
	}).Timestamp().Logger()
}
