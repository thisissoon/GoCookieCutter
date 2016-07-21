package logger

import (
	"bytes"
	"io"
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/Sirupsen/logrus"
	"github.com/Sirupsen/logrus/formatters/logstash"
)

var logbk *Logger

func setup() {
	logbk = log
}

func teardown() {
	log = logbk
}

type fmtr struct {
	msg  string
	data logrus.Fields
}

func (f *fmtr) Format(e *logrus.Entry) ([]byte, error) {
	f.msg = e.Message
	f.data = e.Data
	return []byte(e.Message), nil
}

func TestFields(t *testing.T) {
	setup()

	tt := []struct {
		logger       *Logger
		msg          string
		fields       F
		expectedMsg  string
		expectedData logrus.Fields
	}{
		{
			New(), "foo", F{"foo": "bar"}, "foo", logrus.Fields{"foo": "bar"},
		},
	}

	for _, tc := range tt {
		b := &bytes.Buffer{}
		log = tc.logger
		log.entry.Logger.Out = b
		f := &fmtr{}
		log.entry.Logger.Formatter = f

		Fields(tc.fields).Error(tc.msg)

		if !reflect.DeepEqual(tc.expectedMsg, f.msg) {
			t.Errorf("Unexpected msg, want %s, got %s", tc.expectedMsg, f.msg)
		}

		if !reflect.DeepEqual(tc.expectedData, f.data) {
			t.Errorf("Unexpected data, want %#v, got %#v", tc.expectedData, f.data)
		}
	}

	teardown()
}

func TestLogMessage(t *testing.T) {
	setup()

	tt := []struct {
		logger   *Logger
		lvl      string
		msg      string
		args     []interface{}
		expected []byte
	}{
		{New(), "debug", "foo", []interface{}{}, []byte("foo")},
		{New(), "debug", "foo %s", []interface{}{"bar"}, []byte("foo bar")},
		{New(), "info", "foo", []interface{}{}, []byte("foo")},
		{New(), "info", "foo %s", []interface{}{"bar"}, []byte("foo bar")},
		{New(), "warn", "foo", []interface{}{}, []byte("foo")},
		{New(), "warn", "foo %s", []interface{}{"bar"}, []byte("foo bar")},
		{New(), "error", "foo", []interface{}{}, []byte("foo")},
		{New(), "error", "foo %s", []interface{}{"bar"}, []byte("foo bar")},
	}

	for _, tc := range tt {
		b := &bytes.Buffer{}
		log = tc.logger
		log.entry.Logger.Out = b
		log.entry.Logger.Formatter = &fmtr{}

		SetLevel(tc.lvl)
		switch tc.lvl {
		case "debug":
			Debug(tc.msg, tc.args...)
		case "info":
			Info(tc.msg, tc.args...)
		case "warn":
			Warn(tc.msg, tc.args...)
		case "error":
			Error(tc.msg, tc.args...)
		}

		if !bytes.Equal(tc.expected, b.Bytes()) {
			t.Errorf("Unexpected log message, want %s, got %s", tc.expected, b.Bytes())
		}
	}

	teardown()
}

func TestSetLevel(t *testing.T) {
	setup()

	tt := []struct {
		logger   *Logger
		level    string
		expected logrus.Level
		err      error
	}{
		{New(), "debug", logrus.DebugLevel, nil},
		{New(), "info", logrus.InfoLevel, nil},
		{New(), "warn", logrus.WarnLevel, nil},
		{New(), "error", logrus.ErrorLevel, nil},
		{New(), "DEBUG", logrus.DebugLevel, nil},
		{New(), "INFO", logrus.InfoLevel, nil},
		{New(), "WARN", logrus.WarnLevel, nil},
		{New(), "ERROR", logrus.ErrorLevel, nil},
		{New(), "foo", logrus.InfoLevel, EUNSUPPORTEDLEVEL},
	}

	for _, tc := range tt {
		log = tc.logger
		err := SetLevel(tc.level)
		lvl := log.entry.Logger.Level

		if !reflect.DeepEqual(tc.expected, lvl) {
			t.Errorf("Unexpected level, want %s, got %s", tc.expected, lvl)
		}

		if !reflect.DeepEqual(tc.err, err) {
			t.Errorf("Unexpected err, want %#v, got %#v", tc.err, err)
		}
	}

	teardown()
}

func TestSetLogstashFormat(t *testing.T) {
	setup()

	tt := []struct {
		logger   *Logger
		t        string
		expected *logstash.LogstashFormatter
	}{
		{New(), "", &logstash.LogstashFormatter{Type: DEFAULT_LOGSTASH_TYPE}},
		{New(), "foo", &logstash.LogstashFormatter{Type: "foo"}},
	}

	for _, tc := range tt {
		log = tc.logger
		SetLogstashFormat(tc.t)
		fmtr := log.entry.Logger.Formatter

		if !reflect.DeepEqual(tc.expected, fmtr) {
			t.Errorf("Unexpected formatter, want %#v, got %#v", tc.expected, fmtr)
		}
	}

	teardown()
}

func TestLogVersion(t *testing.T) {
	setup()

	tt := []struct {
		logger   *Logger
		version  string
		expected logrus.Fields
	}{
		{New(), "1.2.3", logrus.Fields{"version": "1.2.3"}},
	}

	for _, tc := range tt {
		log = tc.logger
		LogVersion(tc.version)
		data := log.entry.Data

		if !reflect.DeepEqual(tc.expected, data) {
			t.Errorf("Unexpected log version, want %#v, got %#v", tc.expected, data)
		}
	}

	teardown()
}

func TestConfigure(t *testing.T) {
	setup()

	tt := []struct {
		logger *Logger
		config *Config

		expectedLevel     logrus.Level
		expectedFields    logrus.Fields
		expectedFormatter logrus.Formatter
		expectedErr       error
	}{
		{New(), &Config{Level: "error"}, logrus.ErrorLevel, logrus.Fields{"version": ""}, defaultFormatter, nil},
		{New(), &Config{Level: "foo"}, logrus.InfoLevel, logrus.Fields{"version": ""}, defaultFormatter, EUNSUPPORTEDLEVEL},
		{New(), &Config{Logstash: true}, logrus.InfoLevel, logrus.Fields{"version": ""}, &logstash.LogstashFormatter{Type: DEFAULT_LOGSTASH_TYPE}, nil},
		{New(), &Config{Logstash: true, LogstashType: "foo"}, logrus.InfoLevel, logrus.Fields{"version": ""}, &logstash.LogstashFormatter{Type: "foo"}, nil},
		{New(), &Config{Version: "1.2.3"}, logrus.InfoLevel, logrus.Fields{"version": "1.2.3"}, defaultFormatter, nil},
		{New(), &Config{LogFile: "/path/to.file"}, logrus.InfoLevel, logrus.Fields{"version": ""}, defaultFormatter, nil},
		{New(), &Config{DisableStdOut: true}, logrus.InfoLevel, logrus.Fields{"version": ""}, defaultFormatter, nil},
	}

	for _, tc := range tt {
		log = tc.logger

		err := Configure(tc.config)
		level := log.entry.Logger.Level
		data := log.entry.Data
		fmtr := log.entry.Logger.Formatter

		if !reflect.DeepEqual(tc.expectedLevel, level) {
			t.Errorf("Unexpected log level, want %#v, got %#v", tc.expectedLevel, data)
		}

		if !reflect.DeepEqual(tc.expectedFields, data) {
			t.Errorf("Unexpected log data, want %#v, got %#v", tc.expectedFields, data)
		}

		if !reflect.DeepEqual(tc.expectedFormatter, fmtr) {
			t.Errorf("Unexpected log formatter, want %#v, got %#v", tc.expectedFormatter, fmtr)
		}

		if !reflect.DeepEqual(tc.expectedErr, err) {
			t.Errorf("Unexpected log err, want %#v, got %#v", tc.expectedErr, err)
		}
	}

	teardown()
}

func TestLogToFile(t *testing.T) {
	setup()

	hook := NewFSHook("/tmp/foo.log")

	tt := []struct {
		logger *Logger

		expected logrus.LevelHooks
	}{
		{
			New(),
			map[logrus.Level][]logrus.Hook{
				logrus.DebugLevel: []logrus.Hook{hook},
				logrus.InfoLevel:  []logrus.Hook{hook},
				logrus.WarnLevel:  []logrus.Hook{hook},
				logrus.ErrorLevel: []logrus.Hook{hook},
			},
		},
	}

	for _, tc := range tt {
		log = tc.logger

		LogToFile("/tmp/foo.log")
		hooks := log.entry.Logger.Hooks

		if !reflect.DeepEqual(tc.expected, hooks) {
			t.Errorf("Unexpected hooks, want %#v, got %#v", tc.expected, hooks)
		}
	}

	teardown()
}

func TestDisableStdOut(t *testing.T) {
	setup()

	tt := []struct {
		logger *Logger

		expected io.Writer
	}{
		{New(), ioutil.Discard},
	}

	for _, tc := range tt {
		log = tc.logger

		DisableStdOut()
		w := log.entry.Logger.Out

		if !reflect.DeepEqual(tc.expected, w) {
			t.Errorf("Unexpected writer, want %#v, got %#v", tc.expected, w)
		}
	}

	teardown()

}

func TestWriter(t *testing.T) {
	setup()

	logger := New()

	tt := []struct {
		logger *Logger

		expected func() *io.PipeWriter
	}{
		{logger, func() *io.PipeWriter {
			_, w := io.Pipe()
			return w
		}},
	}

	for _, tc := range tt {
		log = tc.logger

		w := Writer()

		if !reflect.DeepEqual(tc.expected(), w) {
			t.Errorf("Unexpected writer, want %#v, got %#v", tc.expected(), w)
		}
	}

	teardown()
}
