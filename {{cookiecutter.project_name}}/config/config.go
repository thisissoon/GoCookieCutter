// Configuration loading methods, allows config to be loaded from pre defined
// paths or from a specific file path.
// CLI Flags can also be bound to speicifc viper lookup paths

package config

import (
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Allows CLI Flags to be bound to configuration
// For example allowing a --level/-l for overriding logger verbosity
// This bining should be performing in cli package init methods
func BindFlag(s string, f *pflag.Flag) {
	viper.BindPFlag(s, f)
}

// Read configuration into viper from file
func Read() error {
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/{{ cookiecutter.client_name|lower }}/{{ cookiecutter.project_name|lower }}")
	viper.AddConfigPath("$HOME/.config/{{ cookiecutter.client_name|lower }}/{{ cookiecutter.project_name|lower }}")
	return viper.ReadInConfig()
}

// Read config from a specifc file
func ReadFromPath(p string) error {
	viper.SetConfigFile(p)
	return viper.ReadInConfig()
}

// Package initialiser
func init() {
	viper.SetConfigType("toml") // We only support toml, cos it's awesome!
	viper.SetEnvPrefix("{{ cookiecutter.client_name|lower }}_{{ cookiecutter.project_name|lower }}")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}
