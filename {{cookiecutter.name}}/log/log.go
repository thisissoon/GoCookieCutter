package log

import (
	"io"
	"io/ioutil"
	"os"

	"{{cookiecutter.name}}/config"

	"github.com/rs/zerolog"
)

// Defaults add a set of default fields to a logger
func Defaults(l zerolog.Logger) zerolog.Logger {
	return l.With().Fields(map[string]interface{}{
		"app":     config.APP_NAME,
		"version": config.Version(),
		"commit":  config.Commit(),
	}).Timestamp().Logger()
}

// Writer returns an appropriate io.Writer for the
// configured log format
func Writer() io.Writer {
	var w io.Writer = os.Stdout
	switch config.LogFormat() {
	case "console", "terminal":
		w = zerolog.ConsoleWriter{
			Out: os.Stdout,
		}
	case "discard":
		w = ioutil.Discard
	}
	return w
}
