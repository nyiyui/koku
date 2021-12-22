package api

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	// ErrMultipleLateStarts says that there are multiple late start events found.
	ErrMultipleLateStarts = errors.New("multiple late starts found")

	// ErrSchoolNotOpen says that school isn't open.
	ErrSchoolNotOpen = errors.New("school not open")
)

// StatusError is an error from an HTTP status code.
type StatusError struct {
	code int
}

var _ error = StatusError{}

func (err StatusError) Error() string {
	return fmt.Sprintf("status code %d %s", err.code, http.StatusText(err.code))
}
