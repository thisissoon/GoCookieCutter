// Default HTTP Handlers serving common routes

package server

import (
	"net/http"

	"{{ cookiecutter.project_name|lower }}/config"
	"{{ cookiecutter.project_name|lower }}/hal"

	"github.com/labstack/echo"
)

// GET /__healthcheck__
// The Healthcheck endpoint returns the current version with a 200 response
func healthcheckHandler(ctx echo.Context) error {
	r := hal.Resource(ctx, struct {
		Version string `json:"version"`
	}{
		Version: config.Version(),
	})

	return hal.Response(http.StatusOK, ctx, r)
}
