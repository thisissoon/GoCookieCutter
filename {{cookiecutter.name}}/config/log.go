package config

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

// viper lookup keys
const LOG_FORMAT_KEY = "log.format"

// intit sets up default values and binds environment variables
func init() {
	// Set logger level field to severity for stack driver support
	zerolog.LevelFieldName = "severity"
	// Defaults
	viper.SetDefault(LOG_FORMAT_KEY, "json")
	// Environment variable binding
	viper.BindEnv(LOG_FORMAT_KEY)
}

// LogFormat returns the configured log format:w
func LogFormat() string {
	return viper.GetString(LOG_FORMAT_KEY)
}

// LogWriter returns an appropriate io.Writer for the
// configured log format
func LogWriter() io.Writer {
	var w io.Writer = os.Stdout
	switch LogFormat() {
	case "console", "terminal":
		w = zerolog.ConsoleWriter{
			Out: os.Stdout,
		}
	case "discard":
		w = ioutil.Discard
	}
	return w
}
