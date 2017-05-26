package config

import (
	"os"
	"strings"

	"{{cookiecutter.project_name}}/log"

	"github.com/spf13/viper"
)

// Configuration defaults
func init() {
	viper.SetTypeByDefaultValue(true)
	viper.SetConfigType("toml")
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/{{cookiecutter.project_name}}")
	{% if cookiecutter.local_config_path != "" -%}
	viper.AddConfigPath("{{cookiecutter.local_config_path}}")
	{%- endif %}
	viper.SetEnvPrefix("{{cookiecutter.project_name}}")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := viper.ReadInConfig(); err != nil {
		log.WithError(err).Warn("error reading configuration, using default values")
	}
}

// Load custom configuration file
func Read(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	return viper.ReadConfig(f)
}
