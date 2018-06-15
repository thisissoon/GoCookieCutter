package config

import (
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// application name
const APP_NAME = "{{cookiecutter.name}}"

// configuration path viper lookup keys
const (
	CONFIG_PATH_KEY = "config"
	ENVIORNMENT_KEY = "environment"
)

// init sets default configuration file settings such as path look
// up values and environment variable binding
func init() {
	// Config file lookup locations
	viper.SetConfigType("toml")
	viper.SetConfigName("{{cookiecutter.name}}")
	viper.AddConfigPath("$HOME/.config/")
	{% if cookiecutter.project is not none -%}
	viper.AddConfigPath("/etc/{{cookiecutter.project}}/{{cookiecutter.name}}")
	viper.AddConfigPath("$HOME/.config/{{cookiecutter.project}}/{{cookiecutter.name}}")
	{% else -%}
	viper.AddConfigPath("/etc/{{cookiecutter.name}}")
	viper.AddConfigPath("$HOME/.config/{{cookiecutter.name}}")
	{% endif -%}
	// Environment variables
	viper.AutomaticEnv()
	viper.SetEnvPrefix("{{cookiecutter.name}}")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	// Environment variable binding
	viper.SetDefault(ENVIORNMENT_KEY, "prod") // Default env - always assume prod
	viper.BindEnv(CONFIG_PATH_KEY, ENVIORNMENT_KEY)
}

// FromFile reads configuration from a file, bind a CLI flag to
func FromFile() error {
	path := viper.GetString(CONFIG_PATH_KEY)
	if path != "" {
		viper.SetConfigFile(path)
	}
	return viper.ReadInConfig()
}

// BindPFlag binds a config key to a CLI pflag
func BindPFlags(flags map[string]*pflag.Flag) {
	for k, f := range flags {
		viper.BindPFlag(k, f)
	}
}

// Environment returns the current configured environment
func Environment() string {
	return viper.GetString(ENVIORNMENT_KEY)
}
