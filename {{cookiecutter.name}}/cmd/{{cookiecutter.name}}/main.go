package main

import (
	"io"
	"os"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	configkit "go.soon.build/kit/config"

	"{{cookiecutter.module}}/internal/config"
	"{{cookiecutter.module}}/internal/version"
)

// Default logger
var log zerolog.Logger

// Global app configuration
var cfg config.Config

// Default log config
var defaultLog = config.Log{
	Console: false,
	Verbose: false,
	Level:   zerolog.DebugLevel.String(),
}

// Application entry point
func main() {
	cmd := {{cookiecutter.name|replace('-', '')|replace('.', '')}}Cmd()
	if err := cmd.Execute(); err != nil {
		log.Error().Err(err).Msg("exiting from fatal error")
		os.Exit(1)
	}
}

// New constructs a new CLI interface for execution
func {{cookiecutter.name|replace('-', '')|replace('.', '')}}Cmd() *cobra.Command {
	var configPath string
	cmd := &cobra.Command{
		Use:           "{{cookiecutter.name}}",
		Short:         "Run the service",
		SilenceErrors: true,
		SilenceUsage:  true,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// setup logger to capture config errors
			log = initLogger(defaultLog)
			// Load config
			var err error
			cfg, err = config.New(
				configkit.WithFile(configPath),
				configkit.BindFlag("log.console", cmd.Flag("console")),
				configkit.BindFlag("log.verbose", cmd.Flag("verbose")),
			)
			if err != nil {
				return err
			}
			// Override logger with user config
			log = initLogger(cfg.Log)
			return nil
		},
		RunE: {{cookiecutter.name|replace('-', '')|replace('.', '')}}Run,
	}
	// Global flags
	pflags := cmd.PersistentFlags()
	pflags.StringVarP(&configPath, "config", "c", "", "path to configuration file (default is $HOME/.config/{{cookiecutter.name}}.toml)")
	pflags.Bool("console", false, "use console log writer")
	pflags.BoolP("verbose", "v", false, "verbose logging")
	// Add sub commands
	cmd.AddCommand(versionCmd())
	return cmd
}

// {{cookiecutter.name}}Run is executed when the CLI executes
// the {{cookiecutter.name}} command
func {{cookiecutter.name|replace('-', '')|replace('.', '')}}Run(cmd *cobra.Command, _ []string) error {
	return cmd.Help()
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
		"app":     config.AppName,
	}).Timestamp().Logger()
}
