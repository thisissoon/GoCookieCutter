// Configuration handling for the logger
//
// Usage:
//
// import (
//     "{{ cookiecutter.project_name|lower }}/config"
//     "{{ cookiecutter.project_name|lower }}/logger"
// )
//
// func main() {
//     c := config.NewLoggerConfig()
//     if err := logger.Configure(c); err != nil {
//         fmt.Println(err)
//         return
//     }
//     logger.Debug("Some Logger Message")
// }

package config

import (
	"{{ cookiecutter.project_name|lower }}/logger"

	"github.com/spf13/viper"
)

// Logger viper path lookup, example:
// These are exported to CLI flags can be found to these paths
//
// [logger]
// level = "debug"
// logstash = true
// logstash_type = "foo"
// logfile = "/path/to/file.log"
// disable_stdout = true
const (
	VLOGGER_LEVEL          = "logger.level"
	VLOGGER_LOGSTASH       = "logger.logstash"
	VLOGGER_LOGSTASH_TYPE  = "logger.logstash_type"
	VLOGGER_LOGFILE        = "logger.logfile"
	VLOGGER_DISABLE_STDOUT = "logger.disable_stdout"
)

// Returns a logger config instance configured according to
// values in viper, sets sensible defaults.
func NewLoggerConfig() *logger.Config {
	// Logger Verbosity
	if !viper.IsSet(VLOGGER_LEVEL) {
		viper.Set(VLOGGER_LEVEL, "info")
	}

	// Logstash
	if viper.Get(VLOGGER_LOGSTASH) == nil {
		viper.Set(VLOGGER_LOGSTASH, false)
	}

	// Logstash Type
	if !viper.IsSet(VLOGGER_LOGSTASH_TYPE) {
		viper.Set(VLOGGER_LOGSTASH_TYPE, "")
	}

	// File Logging
	if !viper.IsSet(VLOGGER_LOGFILE) {
		viper.Set(VLOGGER_LOGFILE, "")
	}

	// Disable Std Out
	if viper.Get(VLOGGER_DISABLE_STDOUT) == nil {
		viper.Set(VLOGGER_DISABLE_STDOUT, false)
	}

	// Always log with versions
	return &logger.Config{
		Level:         viper.GetString(VLOGGER_LEVEL),
		Logstash:      viper.GetBool(VLOGGER_LOGSTASH),
		LogstashType:  viper.GetString(VLOGGER_LOGSTASH_TYPE),
		LogFile:       viper.GetString(VLOGGER_LOGFILE),
		DisableStdOut: viper.GetBool(VLOGGER_DISABLE_STDOUT),
		Version:       Version(),
	}
}
