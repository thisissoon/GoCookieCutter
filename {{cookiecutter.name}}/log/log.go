package log

import (
	"{{cookiecutter.name}}/config"

	"github.com/rs/zerolog"
)

// Default returns a new default logger setup from configuration
func Default() zerolog.Logger {
	f := map[string]interface{}{
		"app":     config.APP_NAME,
		"env":     config.Environment(),
		"version": config.Version(),
		"commit":  config.Commit(),
	}
	return zerolog.New(config.LogWriter()).
		With().
		Timestamp().
		Fields(f).
		Logger()
}
