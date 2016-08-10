package server

import (
	"testing"

	"{{ cookiecutter.project_name|lower }}/testutil"

	"github.com/labstack/echo"
)

func TestHelathcheckHandler(t *testing.T) {
	tt := []struct {
		expected string
	}{
		{`{"_links":{"self":{"href":"/__healthcheck__"}},"version":"unknown"}`},
	}

	for _, tc := range tt {
		_, rec, ctx := testutil.EchoRecorder(echo.GET, "/__healthcheck__", nil, nil)
		healthcheckHandler(ctx)

		if tc.expected != rec.Body.String() {
			t.Errorf("wanted %s got %s", tc.expected, rec.Body.String())
		}
	}
}
