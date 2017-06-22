package cli

import (
	"{{cookiecutter.project_name}}/app"
	"{{cookiecutter.project_name}}/build"
	"{{cookiecutter.project_name}}/config"
	"{{cookiecutter.project_name}}/log"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// Custom configuration file path
var configPath string

// Root CLI Command
var rootCmd = &cobra.Command{
	Use:   "{{cookiecutter.project_name}}",
	Short: "{{cookiecutter.description}}",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Bind log-level flag to config to allow log level to be set from the CLI Flag
		config.LogLevelFlag(cmd.Flags().Lookup("log-level"))
		// Load custom configuration file if path given
		if configPath != "" {
			if err := config.Read(configPath); err != nil {
				log.WithError(err).Error("error reading configuration")
			}
		}
		// Create a new logger
		l := log.New(log.Config{
			Level:        config.LogLevel(),
			Format:       config.LogFormat(),
			File:         config.LogFile(),
			Console:      config.LogConsole(),
			LogstashType: config.LogstashType(),
		})
		// Always log these fields
		l.PersistentFields(logrus.Fields{
			"version":   build.Version(),
			"buildTime": build.BuildTime(),
		})
		// Set global logger
		log.SetGlobalLogger(l)
	},
	Run: func(*cobra.Command, []string) {
		if err := app.Run(); err != nil {
			log.WithError(err).Error("application run error")
		}
	},
}

// Initialiser
func init() {
	// Add CLI Flags to Root Command
	rootCmd.PersistentFlags().StringVarP(
		&configPath,
		"config",
		"c",
		"",
		"Absolute path to configuration file")
	rootCmd.PersistentFlags().StringP(
		"log-level",
		"l",
		"",
		"Log Level (debug,info,warn,error)")
	// Add Sub Commands
	rootCmd.AddCommand(versionCmd)
}

// Execute CLI
func Execute() error {
	return rootCmd.Execute()
}
