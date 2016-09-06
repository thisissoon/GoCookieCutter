package logger

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
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
	tt := []struct {
		msg          string
		fields       F
		expectedMsg  string
		expectedData logrus.Fields
	}{
		{
			"foo",
			F{"foo": "bar"},
			"foo",
			logrus.Fields{"foo": "bar"},
		},
	}
	for i, tc := range tt {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			setup()
			defer teardown()
			b := &bytes.Buffer{}
			log = New()
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
		})
	}
}

func TestLogMessage(t *testing.T) {
	tt := []struct {
		lvl      string
		msg      interface{}
		args     []interface{}
		expected []byte
	}{
		{"debug", "foo", nil, []byte("foo")},
		{"debug", "foo %s", []interface{}{"bar"}, []byte("foo bar")},
		{"debug", errors.New("foo"), nil, []byte("foo")},
		{"info", "foo", nil, []byte("foo")},
		{"info", "foo %s", []interface{}{"bar"}, []byte("foo bar")},
		{"info", errors.New("foo"), nil, []byte("foo")},
		{"warn", "foo", nil, []byte("foo")},
		{"warn", "foo %s", []interface{}{"bar"}, []byte("foo bar")},
		{"warn", errors.New("foo"), nil, []byte("foo")},
		{"error", "foo", nil, []byte("foo")},
		{"error", "foo %s", []interface{}{"bar"}, []byte("foo bar")},
		{"error", errors.New("foo"), nil, []byte("foo")},
		{"print", "", []interface{}{"foo"}, []byte("foo")},
		{"println", "", []interface{}{"foo"}, []byte("foo")},
		{"printf", "foo %s", []interface{}{"bar"}, []byte("foo bar")},
		{"fatalf", "foo %s", []interface{}{"bar"}, []byte("foo bar")},
	}
	for i, tc := range tt {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			setup()
			defer teardown()
			b := &bytes.Buffer{}
			log = New()
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
			case "print":
				Print(tc.args...)
			case "println":
				Println(tc.args...)
			case "printf":
				Printf(tc.msg.(string), tc.args...)
			case "fatalf":
				Fatalf(tc.msg.(string), tc.args...)
			}
			if !bytes.Equal(tc.expected, b.Bytes()) {
				t.Errorf("Unexpected log message, want %s, got %s", tc.expected, b.Bytes())
			}
		})
	}
}

func TestSetLevel(t *testing.T) {
	tt := []struct {
		level    string
		expected logrus.Level
		err      error
	}{
		{"debug", logrus.DebugLevel, nil},
		{"info", logrus.InfoLevel, nil},
		{"warn", logrus.WarnLevel, nil},
		{"error", logrus.ErrorLevel, nil},
		{"DEBUG", logrus.DebugLevel, nil},
		{"INFO", logrus.InfoLevel, nil},
		{"WARN", logrus.WarnLevel, nil},
		{"ERROR", logrus.ErrorLevel, nil},
		{"foo", logrus.InfoLevel, EUNSUPPORTEDLEVEL},
	}
	for i, tc := range tt {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			setup()
			defer teardown()
			log = New()
			err := SetLevel(tc.level)
			lvl := log.entry.Logger.Level
			if !reflect.DeepEqual(tc.expected, lvl) {
				t.Errorf("Unexpected level, want %s, got %s", tc.expected, lvl)
			}
			if !reflect.DeepEqual(tc.err, err) {
				t.Errorf("Unexpected err, want %#v, got %#v", tc.err, err)
			}
		})
	}
}

func TestSetLogstashFormat(t *testing.T) {
	tt := []struct {
		t        string
		expected *logstash.LogstashFormatter
	}{
		{"", &logstash.LogstashFormatter{Type: DEFAULT_LOGSTASH_TYPE}},
		{"foo", &logstash.LogstashFormatter{Type: "foo"}},
	}
	for i, tc := range tt {
		setup()
		defer teardown()
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			log = New()
			SetLogstashFormat(tc.t)
			fmtr := log.entry.Logger.Formatter
			if !reflect.DeepEqual(tc.expected, fmtr) {
				t.Errorf("Unexpected formatter, want %#v, got %#v", tc.expected, fmtr)
			}
		})
	}
}

func TestSetPlainTextFormat(t *testing.T) {
	tt := []struct {
		expected *logrus.TextFormatter
	}{
		{&logrus.TextFormatter{}},
	}
	for i, tc := range tt {
		setup()
		defer teardown()
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			log = New()
			SetPlainTextFormat()
			fmtr := log.entry.Logger.Formatter
			if !reflect.DeepEqual(tc.expected, fmtr) {
				t.Errorf("Unexpected formatter, want %#v, got %#v", tc.expected, fmtr)
			}
		})
	}
}

func TestLogVersion(t *testing.T) {
	tt := []struct {
		version  string
		expected logrus.Fields
	}{
		{"1.2.3", logrus.Fields{"version": "1.2.3"}},
	}
	for i, tc := range tt {
		setup()
		defer teardown()
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			log = New()
			LogVersion(tc.version)
			data := log.entry.Data

			if !reflect.DeepEqual(tc.expected, data) {
				t.Errorf("Unexpected log version, want %#v, got %#v", tc.expected, data)
			}
		})
	}
}

func TestConfigure(t *testing.T) {
	tt := []struct {
		config            *Config
		expectedLevel     logrus.Level
		expectedFields    logrus.Fields
		expectedFormatter logrus.Formatter
		expectedErr       error
	}{
		{&Config{Level: "error"}, logrus.ErrorLevel, logrus.Fields{"version": ""}, defaultFormatter, nil},
		{&Config{Level: "foo"}, logrus.InfoLevel, logrus.Fields{"version": ""}, defaultFormatter, EUNSUPPORTEDLEVEL},
		{&Config{Logstash: true}, logrus.InfoLevel, logrus.Fields{"version": ""}, &logstash.LogstashFormatter{Type: DEFAULT_LOGSTASH_TYPE}, nil},
		{&Config{Logstash: true, LogstashType: "foo"}, logrus.InfoLevel, logrus.Fields{"version": ""}, &logstash.LogstashFormatter{Type: "foo"}, nil},
		{&Config{Version: "1.2.3"}, logrus.InfoLevel, logrus.Fields{"version": "1.2.3"}, defaultFormatter, nil},
		{&Config{LogFile: "/path/to.file"}, logrus.InfoLevel, logrus.Fields{"version": ""}, defaultFormatter, nil},
		{&Config{DisableStdOut: true}, logrus.InfoLevel, logrus.Fields{"version": ""}, defaultFormatter, nil},
	}
	for i, tc := range tt {
		setup()
		defer teardown()
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			log = New()
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
		})
	}
}

func TestLogToFile(t *testing.T) {
	hook := NewFSHook("/tmp/foo.log")
	tt := []struct {
		expected logrus.LevelHooks
	}{
		{
			map[logrus.Level][]logrus.Hook{
				logrus.DebugLevel: []logrus.Hook{hook},
				logrus.InfoLevel:  []logrus.Hook{hook},
				logrus.WarnLevel:  []logrus.Hook{hook},
				logrus.ErrorLevel: []logrus.Hook{hook},
			},
		},
	}
	for i, tc := range tt {
		setup()
		defer teardown()
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			log = New()
			LogToFile("/tmp/foo.log")
			hooks := log.entry.Logger.Hooks
			if !reflect.DeepEqual(tc.expected, hooks) {
				t.Errorf("Unexpected hooks, want %#v, got %#v", tc.expected, hooks)
			}
		})
	}
}

func TestEnableStdOut(t *testing.T) {
	tt := []struct {
		expected io.Writer
	}{
		{os.Stderr},
	}
	for i, tc := range tt {
		setup()
		defer teardown()
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			log = New()
			DisableStdOut()
			EnableStdOut()
			w := log.entry.Logger.Out
			if !reflect.DeepEqual(tc.expected, w) {
				t.Errorf("Unexpected writer, want %#v, got %#v", tc.expected, w)
			}
		})
	}
}

func TestDisableStdOut(t *testing.T) {
	tt := []struct {
		expected io.Writer
	}{
		{ioutil.Discard},
	}
	for i, tc := range tt {
		setup()
		defer teardown()
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			log = New()
			DisableStdOut()
			w := log.entry.Logger.Out
			if !reflect.DeepEqual(tc.expected, w) {
				t.Errorf("Unexpected writer, want %#v, got %#v", tc.expected, w)
			}
		})
	}
}

func TestLog(t *testing.T) {
	setup()
	defer teardown()
	n := New()
	log = n
	l := Log()
	if !reflect.DeepEqual(l, n) {
		t.Errorf("Unexpected log, want %#v, got %#v", n, l)
	}
}
