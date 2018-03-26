package config

import (
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

// viper lookup keys
const LOG_FORMAT_KEY = "log.format"

// init sets up default values and binds environment variables
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
