// Log Configuration

package config

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	// General Log Configuration
	log_lvl     = "log.level"
	log_fmt     = "log.format"
	log_file    = "log.file"
	log_console = "log.console"
	// Logstash Specific
	logstash_type = "logstash.type"
)

func init() {
	viper.SetDefault(log_lvl, "info")
	viper.SetDefault(log_fmt, "text")
	viper.SetDefault(log_file, "")
	viper.SetDefault(log_console, true)
	viper.SetDefault(logstash_type, "{{cookiecutter.project_name}}")
}

func LogLevelFlag(flag *pflag.Flag) {
	viper.BindPFlag(log_lvl, flag)
}

func LogLevel() string {
	viper.BindEnv(log_lvl)
	return viper.GetString(log_lvl)
}

func LogFormat() string {
	viper.BindEnv(log_fmt)
	return viper.GetString(log_fmt)
}

func LogFile() string {
	viper.BindEnv(log_file)
	return viper.GetString(log_file)
}

func LogConsole() bool {
	viper.BindEnv(log_console)
	return viper.GetBool(log_console)
}

func LogstashType() string {
	viper.BindEnv(logstash_type)
	return viper.GetString(logstash_type)
}
