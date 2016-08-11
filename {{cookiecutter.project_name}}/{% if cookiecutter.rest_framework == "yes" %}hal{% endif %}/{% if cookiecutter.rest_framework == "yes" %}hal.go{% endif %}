package hal

import (
	"net/http"

	"{{ cookiecutter.project_name|lower }}/logger"

	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo"
	"github.com/nvellon/hal"
)

// Returns a hal formatted response
func Response(status int, ctx echo.Context, r *hal.Resource) error {
	return ctx.JSON(status, r)
}

// Builds a HAL resource doucment from request context
func Resource(ctx echo.Context, i interface{}) *hal.Resource {
	return hal.NewResource(i, ctx.Request().URI())
}

// Common 500 error handler, logging the error
func ServerError(ctx echo.Context, err error) error {
	logger.Error("%#v", err)
	r := Resource(ctx, EDefaultError)
	return Response(http.StatusInternalServerError, ctx, r)
}

// Common 404 handler
func NotFound(ctx echo.Context) error {
	return Response(http.StatusNotFound, ctx, Resource(ctx, ENotFound))
}

// Common Validation Error handler
func ValidationError(ctx echo.Context, i interface{}, err error) error {
	_, ok := err.(govalidator.Errors)
	if !ok {
		return Response(422, ctx, Resource(ctx, &Error{Message: err.Error()}))
	}

	return Response(422, ctx, Resource(ctx, &Error{
		Message:  "Validation Error",
		ErrorMap: ValidationErrorMap(i, err),
	}))
}
