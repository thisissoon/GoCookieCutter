// Common HAL Formatted Error Handlers

package hal

import (
	"errors"
	"reflect"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/nvellon/hal"
)

// Common 500 error response
var EDefaultError = &Error{
	Message: "An unexpected error has occured",
}

// Common 404 error response
var ENotFound = &Error{
	Message: "This is not the resource you are looking for",
}

// This type allows us to produce a map of specifc errors, this
// is useful for validation error messages for example where we
// need to list errors field by field
type ErrorMap map[string]string

func (m ErrorMap) Add(field string, err error) {
	m[field] = err.Error()
}

func (m ErrorMap) Delete(field string) {
	delete(m, field)
}

// A HAL error type that satisfies the nvellon/hal mapper interface
// and the standard go error interface, this allows us to produce
// hal formatted error responses
type Error struct {
	Message  string   `json:"message"`
	ErrorMap ErrorMap `json:"errors,omitempty"`
}

func (e *Error) Error() string {
	return e.Message
}

func (e *Error) GetMap() hal.Entry {
	return hal.Entry{"_errors": e}
}

// Constructs an ErrorMap from govalidator validation errors, if json tags also exist
// the error map will be populated with the json field name instead of the struct field name
func ValidationErrorMap(i interface{}, err error) ErrorMap {
	// Ensure we have a govalidator error
	err, ok := err.(govalidator.Errors)
	if !ok {
		return nil
	}

	// Use refect to get the raw struct element
	typ := reflect.TypeOf(i).Elem()
	if typ.Kind() != reflect.Struct {
		return nil
	}

	// Errors found by the validator
	errsByField := govalidator.ErrorsByField(err.(govalidator.Errors))

	// Make an error map
	m := make(ErrorMap, len(errsByField))

	// Loop over our struct fields
	for i := 0; i < typ.NumField(); i++ {
		// Get the field
		f := typ.Field(i)
		// Do we have an error for the field
		e, ok := errsByField[f.Name]
		if ok {
			// Try and get the `json` struct tag
			name := strings.Split(f.Tag.Get("json"), ",")[0]
			// If the name is - we should ignore the field
			if name == "-" {
				continue
			}
			// If the name is not blank we add it our error map
			if name != "" {
				m.Add(name, errors.New(e))
				continue
			}
			// Finall if all else has failed just add the raw field name to the
			// error map
			m.Add(f.Name, errors.New(e))
		}
	}

	return m
}
