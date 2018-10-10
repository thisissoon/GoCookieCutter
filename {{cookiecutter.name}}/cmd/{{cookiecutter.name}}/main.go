package main

import (
	"io"
	"os"

	"{{cookiecutter.module}}/internal/config"
	"{{cookiecutter.module}}/internal/version"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

// Default logger
var log zerolog.Logger

// Global app configuration
var configMain config.Config

// Application entry point
func main() {
	{{cookiecutter.name}}Cmd().Execute()
}

// New constructs a new CLI interface for execution
func {{cookiecutter.name}}Cmd() *cobra.Command {
	var configPath string
	cmd := &cobra.Command{
		Use:   "{{cookiecutter.name}}",
		Short: "Run the service",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// Load config
			configMain, err := config.New(
				config.ConfigFile(configPath),
				config.BindFlag("log.console", cmd.Flag("console")),
				config.BindFlag("log.verbose", cmd.Flag("verbose")),
			)
			if err != nil {
				return err
			}
			// Setup default logger
			log = initLogger(configMain.Log)
			return nil
		},
		Run:   {{cookiecutter.name}}Run,
	}
	// Global flags
	pflags := cmd.PersistentFlags()
	{% if cookiecutter.project is not none -%}
	pflags.StringVarP(&configPath, "config", "c", "", "path to configuration file (default is $HOME/.config/{{cookiecutter.project}}/{{cookiecutter.name}}.toml)")
	{% else -%}
	pflags.StringVarP(&configPath, "config", "c", "", "path to configuration file (default is $HOME/.config/{{cookiecutter.name}}.toml)")
	{% endif -%}
	pflags.Bool("console", false, "use console log writer")
	pflags.BoolP("verbose", "v", false, "verbose logging")
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
func initLogger(c config.Log) zerolog.Logger {
	// Set logger level field to severity for stack driver support
	zerolog.LevelFieldName = "severity"
	var w io.Writer = os.Stdout
	if c.Console {
		w = zerolog.ConsoleWriter{
			Out: os.Stdout,
		}
	}
	// Parse level from config
	lvl, err := zerolog.ParseLevel(c.Level)
	if err != nil {
		lvl = zerolog.InfoLevel
	}
	// Override level with verbose
	if c.Verbose {
		lvl = zerolog.DebugLevel
	}
	return zerolog.New(w).Level(lvl).With().Fields(map[string]interface{}{
		"version": version.Version,
		"app":     config.APP_NAME,
	}).Timestamp().Logger()
}
