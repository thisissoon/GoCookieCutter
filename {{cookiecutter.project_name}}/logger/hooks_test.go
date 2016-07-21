package logger

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

// Tempfile maker
var mkTempFile = func(t *testing.T) (*os.File, func()) {
	f, err := ioutil.TempFile(os.TempDir(), "test")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("Created Temp File @", f.Name())

	return f, func() {
		f.Close()
		if err := os.Remove(f.Name()); err != nil {
			t.Fatal(err)
		}
		fmt.Println("Removed Temp File @", f.Name())
	}
}

func TestFSHookFire(t *testing.T) {
	tt := []struct {
		logger *Logger
		msg    []byte
	}{
		{New(), []byte("Log to file test")},
	}

	for _, tc := range tt {
		f, cleanup := mkTempFile(t)
		tc.logger.entry.Logger.Formatter = &fmtr{}
		tc.logger.DisableStdOut()
		tc.logger.LogToFile(f.Name())
		tc.logger.Info(string(tc.msg))

		b, err := ioutil.ReadFile(f.Name())
		if err != nil {
			t.Fatal(err)
		}

		if !bytes.Equal(tc.msg, b) {
			t.Errorf("Unexpected log file content, want %s, got %s", tc.msg, b)
		}

		cleanup()
	}

}
