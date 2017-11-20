package config

import (
	"time"

	"github.com/spf13/viper"
)

// viper lookup keys
const (
	VERSION_KEY   = "version"
	COMMIT_KEY    = "commit"
	TIMESTAMP_KEY = "timestamp"
)

// Build time variables with default values
// These are set at compile time
var (
	version   string
	timestamp string
	commit    string
)

// init sets defaults and binds enviornment variables
func init() {
	// set defaults
	viper.SetDefault(VERSION_KEY, version)
	viper.SetDefault(COMMIT_KEY, commit)
	viper.SetDefault(TIMESTAMP_KEY, timestamp)
	// bind environment variables
	viper.BindEnv(
		VERSION_KEY,
		COMMIT_KEY,
		TIMESTAMP_KEY,
	)
}

// Version returns the build version
func Version() string {
	return viper.GetString(VERSION_KEY)
}

// Commit returns the build commit hash
func Commit() string {
	return viper.GetString(COMMIT_KEY)
}

// BuildTimestamp returns the build timestamp
func BuildTimestamp() time.Time {
	return time.Unix(viper.GetInt64(TIMESTAMP_KEY), 0).UTC()
}
