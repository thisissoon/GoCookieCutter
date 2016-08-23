// Main CLI Application Entrypoint

package cli

import (
	"errors"
	"fmt"
	"os"

	"{{ cookiecutter.project_name|lower }}/config"
	"{{ cookiecutter.project_name|lower }}/logger"

	"github.com/spf13/cobra"
)

//
// Unexported
//

var {{ cookiecutter.project_name|lower }}Cmd = &cobra.Command{
	Use:   "{{ cookiecutter.project_name|lower }}",
	Short: "{{ cookiecutter.description }}",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if err := configure(cmd); err != nil {
			logger.Error("Configuration Error: %s", err)
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) { cmd.Help() },
}

// Intialiser
func init() {
	// Flags
	{{ cookiecutter.project_name|lower }}Cmd.PersistentFlags().StringP("config", "c", "", "Optional Config File Path")
	{{ cookiecutter.project_name|lower }}Cmd.PersistentFlags().StringP("verbosity", "v", "", "Log Level Verbosity (debug, info, warn, error)")
	// Sub Commands (Add more as required)
	{{ cookiecutter.project_name|lower }}Cmd.AddCommand(versionCmd)
	{% if cookiecutter.restapi == "yes" %}{{ cookiecutter.project_name|lower }}Cmd.AddCommand(runrestapiCmd){% endif %}
}

// Configures static parts of the application
func configure(cmd *cobra.Command) error {
	p, err := cmd.Flags().GetString("config")
	if err != nil {
		return errors.New("Failed to read config CLI flag")
	}

	// Read Config
	if p == "" {
		if err := config.Read(); err != nil {
			logger.Warn("No configuration file found")
		}
	} else {
		if err := config.ReadFromPath(p); err != nil {
			return fmt.Errorf("Failed to read config: %s", err)
		}
	}

	// Configure Logger
	if err := configureLogger(cmd); err != nil {
		return err
	}

	return nil
}

// Configure the central logger
func configureLogger(cmd *cobra.Command) error {
	// Bind a log level flag
	config.BindFlag(config.VLOGGER_LEVEL, cmd.Flag("verbosity"))

	// Configures the logger
	if err := logger.Configure(config.NewLoggerConfig()); err != nil {
		return err
	}

	return nil
}

//
// Exported
//

func Execute() error {
	return {{ cookiecutter.project_name|lower }}Cmd.Execute()
}
