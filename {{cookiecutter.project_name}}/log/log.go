package log

import (
	"io/ioutil"
	"os"
	"strings"

	"{{cookiecutter.project_name}}/log/formatters"
	"{{cookiecutter.project_name}}/log/hooks"

	"github.com/sirupsen/logrus"
)

// Global logger
var log Logger = New(Config{
	Level:   "info",
	Format:  "text",
	Console: true,
})

// Set the global logger to a new log instances
func SetGlobalLogger(l *Log) {
	log = l
}

type Logger interface {
	Debugf(string, ...interface{})
	Infof(string, ...interface{})
	Warnf(string, ...interface{})
	Errorf(string, ...interface{})
	WithError(error) *logrus.Entry
	WithField(string, interface{}) *logrus.Entry
	WithFields(logrus.Fields) *logrus.Entry
}

// Log configuration
type Config struct {
	Level        string
	Format       string
	File         string
	Console      bool
	LogstashType string
}

// Extends logrus.Entry
type Log struct {
	*logrus.Entry
}

// Disable console output
func (log *Log) DisableConsoleOutput() {
	log.Logger.Out = ioutil.Discard
}

// Enable console outout
func (log *Log) EnableConsoleOutput() {
	log.Logger.Out = os.Stdout
}

// Set the Log Level
func (log *Log) SetLevel(lvl string) {
	lvl = strings.ToLower(lvl)
	switch lvl {
	case "debug":
		log.Logger.Level = logrus.DebugLevel
	case "info":
		log.Logger.Level = logrus.InfoLevel
	case "warn":
		log.Logger.Level = logrus.WarnLevel
	case "error":
		log.Logger.Level = logrus.ErrorLevel
	default:
		log.Logger.Level = logrus.InfoLevel
	}
}

// Set the format of the log
func (log *Log) SetFormat(format string, args map[string]interface{}) {
	format = strings.ToLower(format)
	switch format {
	case "json":
		log.Logger.Formatter = &logrus.JSONFormatter{}
	case "logstash":
		typ, ok := args["logstash.type"].(string)
		if !ok {
			typ = "{{cookiecutter.project_name}}"
		}
		log.Logger.Formatter = &formatters.LogstashFormatter{
			Type: typ,
		}
	default:
		log.Logger.Formatter = &logrus.TextFormatter{
			FullTimestamp: true,
		}
	}
}

// Log to a file
func (log *Log) WriteToFile(path string) {
	log.Logger.Hooks.Add(hooks.NewFileHook(path))
}

// Log a persitent field with all log messages
func (log *Log) PersistentField(k string, v interface{}) {
	log.Entry = log.WithField(k, v)
}

// Log persitent fields with all log messages
func (log *Log) PersistentFields(f logrus.Fields) {
	log.Entry = log.WithFields(f)
}

// Add Error
func WithError(err error) *logrus.Entry {
	return log.WithError(err)
}

// Add one field to the log context
func WithField(k string, v interface{}) *logrus.Entry {
	return log.WithField(k, v)
}

// Log multiple fields
func WithFields(f logrus.Fields) *logrus.Entry {
	return log.WithFields(f)
}

// Debug logging
func Debug(fmt string, args ...interface{}) {
	log.Debugf(fmt, args...)
}

// Info Logging
func Info(fmt string, args ...interface{}) {
	log.Infof(fmt, args...)
}

// Warn logging
func Warn(fmt string, args ...interface{}) {
	log.Warnf(fmt, args...)
}

// Error Logging
func Error(fmt string, args ...interface{}) {
	log.Errorf(fmt, args...)
}

// Constructs a new logger
func New(config Config) *Log {
	l := &Log{
		Entry: logrus.NewEntry(logrus.New()),
	}
	l.SetLevel(config.Level)
	l.SetFormat(config.Format, map[string]interface{}{
		"logstash.type": config.LogstashType,
	})
	if config.Console {
		l.EnableConsoleOutput()
	} else {
		l.DisableConsoleOutput()
	}
	return l
}
