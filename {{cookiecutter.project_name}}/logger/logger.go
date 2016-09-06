// Common logger interface wrapper around Logrus
// This package uses the singleton pattern to allow a single instance
// wrapper around the logrus logging library
//
// Usage:
//
// import "{{ cookiecutter.project_name|lower }}/logger"
//
// func main() {
//     logger.Debug("Debug Message")
// }
//
// To log with the version number
//
// import "{{ cookiecutter.project_name|lower }}/logger"
//
// func main() {
//     logger.LogWithVersion("1.2.3")
//     logger.Debug("Debug Message")
// }
//
// To log with extra fields
//
// import "{{ cookiecutter.project_name|lower }}/logger"
//
// func main() {
//     logger.LogWithVersion("1.2.3")
//     logger.Fields(logger.F{"foo": "bar"}).Debug("Debug Message")
// }

package logger

import (
	"errors"
	"io"
	"io/ioutil"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/Sirupsen/logrus/formatters/logstash"
)

// Default logstash type
const DEFAULT_LOGSTASH_TYPE = "{{ cookiecutter.project_name|lower }}"

// Default Formatter
var defaultFormatter = &logrus.TextFormatter{FullTimestamp: true}

// Errors
var EUNSUPPORTEDLEVEL = errors.New("logger: unsupported level")

// Logger instance
var log *Logger

// Logger configuration type
type Config struct {
	Level         string
	Logstash      bool
	LogstashType  string
	LogFile       string
	DisableStdOut bool
	Version       string
}

// Logger type which implements logrus functionality
type Logger struct {
	// Configuration
	Config *Config

	// Raw logrus logger
	entry *logrus.Entry
}

// A custom type for logrus.Fields short hand
type F logrus.Fields

// logger.Fields(logger.f{"foo": "bar"}).Debug("foo")
func Fields(f F) *Logger { return log.Fields(f) }
func (l *Logger) Fields(f F) *Logger {
	return &Logger{
		entry: l.entry.WithFields(logrus.Fields(f)),
	}
}

// Log a error message
// logger.Err("Foo") or logger.Err("%s", "bar")
func Fatal(msg interface{}, v ...interface{}) { log.Fatal(msg, v...) }
func (l *Logger) Fatal(msg interface{}, args ...interface{}) {
	str, ok := msg.(string)
	if !ok {
		args = append([]interface{}{msg}, args...)
		l.entry.Fatal(args...)
	} else {
		l.entry.Fatalf(str, args...)
	}
}

// Log a error message
// logger.Err("Foo") or logger.Err("%s", "bar")
func Error(msg interface{}, v ...interface{}) { log.Error(msg, v...) }
func (l *Logger) Error(msg interface{}, args ...interface{}) {
	str, ok := msg.(string)
	if !ok {
		args = append([]interface{}{msg}, args...)
		l.entry.Error(args...)
	} else {
		l.entry.Errorf(str, args...)
	}
}

// Log a warn message
// logger.Warn("Foo") or logger.Warn("foo: %s", "bar")
func Warn(msg interface{}, v ...interface{}) { log.Warn(msg, v...) }
func (l *Logger) Warn(msg interface{}, args ...interface{}) {
	str, ok := msg.(string)
	if !ok {
		args = append([]interface{}{msg}, args...)
		l.entry.Warn(args...)
	} else {
		l.entry.Warnf(str, args...)
	}
}

// Log an info level message
// logger.Info("Foo") or logger.Info("foo: %s", "bar")
func Info(msg interface{}, v ...interface{}) { log.Info(msg, v...) }
func (l *Logger) Info(msg interface{}, args ...interface{}) {
	str, ok := msg.(string)
	if !ok {
		args = append([]interface{}{msg}, args...)
		l.entry.Info(args...)
	} else {
		l.entry.Infof(str, args...)
	}
}

// Log a debug message
// logger.Debug("Foo") or logger.Debug("foo: %s", "bar")
func Debug(msg interface{}, v ...interface{}) { log.Debug(msg, v...) }
func (l *Logger) Debug(msg interface{}, args ...interface{}) {
	str, ok := msg.(string)
	if !ok {
		args = append([]interface{}{msg}, args...)
		l.entry.Debug(args...)
	} else {
		l.entry.Debugf(str, args...)
	}
}

// Common Logging Interface
func Print(v ...interface{}) { log.Print(v...) }
func (l *Logger) Print(v ...interface{}) {
	l.entry.Info(v...)
}

func Printf(format string, v ...interface{}) { log.Printf(format, v...) }
func (l *Logger) Printf(format string, v ...interface{}) {
	l.entry.Infof(format, v...)
}

func Println(v ...interface{}) { log.Println(v...) }
func (l *Logger) Println(v ...interface{}) {
	l.entry.Infoln(v...)
}

func Fatalf(format string, v ...interface{}) { log.Fatalf(format, v...) }
func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.entry.Errorf(format, v...)
}

// Sets the logging verbosity level
func SetLevel(lvl string) error { return log.SetLevel(lvl) }
func (l *Logger) SetLevel(lvl string) error {
	lvl = strings.ToLower(lvl)
	switch lvl {
	case "debug":
		l.entry.Logger.Level = logrus.DebugLevel
	case "info":
		l.entry.Logger.Level = logrus.InfoLevel
	case "warn":
		l.entry.Logger.Level = logrus.WarnLevel
	case "error":
		l.entry.Logger.Level = logrus.ErrorLevel
	case "":
		l.entry.Logger.Level = logrus.InfoLevel
	default:
		return EUNSUPPORTEDLEVEL
	}

	return nil
}

// Sets the logger format to use logstash formatter
func SetLogstashFormat(t string) { log.SetLogstashFormat(t) }
func (l *Logger) SetLogstashFormat(t string) {
	if t == "" {
		t = DEFAULT_LOGSTASH_TYPE
	}
	l.entry.Logger.Formatter = &logstash.LogstashFormatter{
		Type: t,
	}
}

func SetPlainTextFormat() { log.SetPlainTextFormat() }
func (l *Logger) SetPlainTextFormat() {
	l.entry.Logger.Formatter = &logrus.TextFormatter{}
}

// Sets the logger to always log the version number
func LogVersion(v string) { log.LogVersion(v) }
func (l *Logger) LogVersion(v string) {
	l.entry = l.entry.WithField("version", v)
}

// Central logger configuration handler. Takes a config.LoggerConfig
// instance and configures the logger according to the settings on that instance
func Configure(c *Config) error { return log.Configure(c) }
func (l *Logger) Configure(c *Config) error {
	l.Config = c
	l.LogVersion(c.Version)

	// Set the log level
	if err := l.SetLevel(c.Level); err != nil {
		return err
	}

	// Set Logstash
	if c.Logstash {
		l.SetLogstashFormat(c.LogstashType)
	}

	// Set the logger to log to file
	if c.LogFile != "" {
		l.LogToFile(c.LogFile)
	}

	// Disable Std Out?
	if c.DisableStdOut {
		l.DisableStdOut()
	}

	return nil
}

// Sets the logger outout to a file
func LogToFile(p string) { log.LogToFile(p) }
func (l *Logger) LogToFile(p string) {
	l.entry.Logger.Hooks.Add(NewFSHook(p))
}

// Disables logging to stdout
func DisableStdOut() { log.DisableStdOut() }
func (l *Logger) DisableStdOut() {
	l.entry.Logger.Out = ioutil.Discard
}

// Disables logging to stdout
func EnableStdOut() { log.EnableStdOut() }
func (l *Logger) EnableStdOut() {
	l.entry.Logger.Out = os.Stderr
}

// Returns the current globally scoped logger
func Log() *Logger {
	return log
}

// Constructor
func New() *Logger {
	l := logrus.New()
	l.Formatter = defaultFormatter
	return &Logger{entry: logrus.NewEntry(l)}
}

// Initialiser, ensures we always have a log instance setup
func init() {
	// Set our singleton log instance
	log = New()
}
