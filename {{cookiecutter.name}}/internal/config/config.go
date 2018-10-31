package config

import (
	"reflect"
	"strings"

	"github.com/rs/zerolog"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// application name
const APP_NAME = "{{cookiecutter.name}}"

// Config stores configuration options set by configuration file or env vars
type Config struct {
	Log Log `mapstructure:"log"`
}

// Log contains logging configuration
type Log struct {
	Console bool   `mapstructure:"console"`
	Verbose bool   `mapstructure:"verbose"`
	Level   string `mapstructure:"level"`
}

// An Option function can provide extra viper configuration
type Option func(v *viper.Viper) error

// ConfigFile will override implict configuration file lookups and specify an
// absolute path to a config file to load
func ConfigFile(p string) Option {
	return func(v *viper.Viper) error {
		if p != "" {
			v.SetConfigFile(p)
		}
		return nil
	}
}

// BindFlag returns an Option function allowing the binding of CLI flags to
// confugration values
func BindFlag(key string, flag *pflag.Flag) Option {
	return func(v *viper.Viper) error {
		if flag == nil {
			return nil
		}
		return v.BindPFlag(key, flag)
	}
}

// BindEnvs takes an interface and an optional slice of strings. It recurses
// over the interface which should be a struct and extracts the mapstructure
// tags, if the field is a struct the name of the field is appended to the
// parts slice and BindEnvs is called again with the nested struct and the parts
// slice, if it is not a struct the tag name is joined with the parts slice with
// a . and viper.BindEnv is called with that name
// This allows us to use environment variable bindings with viper.Unmarshal
// which cannot use viper.AutomaticEnv
func BindEnvs(v *viper.Viper, iface interface{}, parts ...string) error {
	ifv := reflect.ValueOf(iface)
	ift := reflect.TypeOf(iface)
	for i := 0; i < ift.NumField(); i++ {
		val := ifv.Field(i)
		tv, ok := ift.Field(i).Tag.Lookup("mapstructure")
		if !ok {
			continue
		}
		switch val.Kind() {
		case reflect.Struct:
			err := BindEnvs(v, val.Interface(), append(parts, tv)...)
			if err != nil {
				return err
			}
		default:
			err := v.BindEnv(strings.Join(append(parts, tv), "."))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Default returns a default configuration setup with sane defaults
func Default() Config {
	return Config{
		Log{
			Level: zerolog.InfoLevel.String(),
		},
	}
}

// New constructs a new Config instance
func New(opts ...Option) (Config, error) {
	c := Default()
	v := viper.New()
	v.SetConfigType("toml")
	v.SetConfigName("{{cookiecutter.name}}")
	// Set default config paths
	{% if cookiecutter.project is not none -%}
	viper.AddConfigPath("/etc/{{cookiecutter.project}}")
	viper.AddConfigPath("$HOME/.config/{{cookiecutter.project}}")
	{% else -%}
	viper.AddConfigPath("/etc/{{cookiecutter.name}}")
	viper.AddConfigPath("$HOME/.config")
	{% endif -%}
	v.SetEnvPrefix("{{cookiecutter.name}}")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	for _, opt := range opts {
		err := opt(v)
		if err != nil {
			return c, err
		}
	}
	err := BindEnvs(v, c)
	if err != nil {
		return c, err
	}
	switch err := v.ReadInConfig(); err.(type) {
	case nil, viper.ConfigFileNotFoundError:
		break
	default:
		return c, err
	}
	if err := v.Unmarshal(&c); err != nil {
		return c, err
	}
	return c, nil
}
