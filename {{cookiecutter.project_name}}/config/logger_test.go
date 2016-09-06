package config

import (
	"reflect"
	"testing"

	"{{ cookiecutter.project_name|lower }}/logger"

	"github.com/spf13/viper"
)

func TestNewLoggerConfig(t *testing.T) {
	tt := []struct {
		level         string
		logstash      bool
		logstash_type string
		logfile       string
		disableStdOut bool
		version       string
		expected      *logger.Config
	}{
		{
			"error",
			false,
			"",
			"",
			false,
			"",
			&logger.Config{
				Level:         "error",
				Logstash:      false,
				LogstashType:  "",
				LogFile:       "",
				DisableStdOut: false,
				Version:       "unknown",
			},
		},
		{
			"",
			true,
			"",
			"",
			false,
			"",
			&logger.Config{
				Level:         "info",
				Logstash:      true,
				LogstashType:  "",
				LogFile:       "",
				DisableStdOut: false,
				Version:       "unknown",
			},
		},
		{
			"",
			false,
			"foo",
			"",
			false,
			"",
			&logger.Config{
				Level:         "info",
				Logstash:      false,
				LogstashType:  "foo",
				LogFile:       "",
				DisableStdOut: false,
				Version:       "unknown",
			},
		},
		{
			"",
			false,
			"",
			"/path/to/log.file",
			false,
			"",
			&logger.Config{
				Level:         "info",
				Logstash:      false,
				LogstashType:  "",
				LogFile:       "/path/to/log.file",
				DisableStdOut: false,
				Version:       "unknown",
			},
		},
		{
			"",
			false,
			"",
			"",
			true,
			"",
			&logger.Config{
				Level:         "info",
				Logstash:      false,
				LogstashType:  "",
				LogFile:       "",
				DisableStdOut: true,
				Version:       "unknown",
			},
		},
		{
			"",
			false,
			"",
			"",
			false,
			"1.2.3",
			&logger.Config{
				Level:         "info",
				Logstash:      false,
				LogstashType:  "",
				LogFile:       "",
				DisableStdOut: false,
				Version:       "1.2.3",
			},
		},
	}
	for i, tc := range tt {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			defer func() {
				version = ""
				viper.Reset()
			}()
			if tc.level != "" {
				viper.Set(VLOGGER_LEVEL, tc.level)
			}
			if tc.logstash {
				viper.Set(VLOGGER_LOGSTASH, true)
			}
			if tc.logstash_type != "" {
				viper.Set(VLOGGER_LOGSTASH_TYPE, tc.logstash_type)
			}
			if tc.logfile != "" {
				viper.Set(VLOGGER_LOGFILE, tc.logfile)
			}
			if tc.disableStdOut {
				viper.Set(VLOGGER_DISABLE_STDOUT, tc.disableStdOut)
			}
			if tc.version != "" {
				version = tc.version
			}
			c := NewLoggerConfig()
			if !reflect.DeepEqual(tc.expected, c) {
				t.Errorf("Unexpected logger config: want %#v, got %#v", tc.expected, c)
			}
		})
	}
}
