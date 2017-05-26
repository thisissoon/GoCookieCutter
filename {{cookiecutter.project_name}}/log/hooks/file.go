package hooks

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

type FileHook struct {
	wcFn func() (io.WriteCloser, error)
}

func (h *FileHook) Fire(entry *logrus.Entry) error {
	wc, err := h.wcFn()
	if err != nil {
		return err
	}
	defer wc.Close()
	serialized, err := entry.Logger.Formatter.Format(entry)
	if err != nil {
		return err
	}
	if _, err := wc.Write(serialized); err != nil {
		return err
	}
	return nil
}

func (hook *FileHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func NewFileHook(path string) *FileHook {
	fn := func() (io.WriteCloser, error) {
		return os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	}
	return &FileHook{fn}
}
