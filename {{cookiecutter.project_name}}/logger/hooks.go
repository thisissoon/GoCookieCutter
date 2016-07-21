// Logrus Hooks, for example a file system logger

package logger

import (
	"os"

	"github.com/Sirupsen/logrus"
)

//
// Local File System Log Hook
// This hook will log entries to a log file on the local file system
//

type fshook struct {
	path string
}

// Log entries to a file
func (h *fshook) Fire(entry *logrus.Entry) error {
	fd, err := os.OpenFile(h.path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return err
	}

	defer fd.Close()

	msg, err := entry.String()
	if err != nil {
		return err
	}
	fd.WriteString(msg)

	return nil
}

// Returns the log levels support by this hook
func (hook *fshook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.DebugLevel,
		logrus.InfoLevel,
		logrus.WarnLevel,
		logrus.ErrorLevel,
	}
}

func NewFSHook(p string) *fshook {
	return &fshook{path: p}
}
