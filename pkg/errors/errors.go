package errors

//
// This is a copy of majordomusio/commons/pkg/error, with improvements.
// Will be eventually backported.
//

import (
	ee "errors"
	"fmt"
	"net/http"
	"runtime"
	"strings"
)

type (
	errorWrapper struct {
		err error
		msg string
		// should match semantics of HTTP status codes. See https://en.wikipedia.org/wiki/List_of_HTTP_status_codes
		status int
		// location of the error
		pkg string
		fn  string
	}

	// ErrorObject is used to report errors in an API request
	ErrorObject struct {
		Status  int    `json:"status" binding:"required"`
		Message string `json:"message" binding:"required"`
	}
)

func (e *errorWrapper) Error() string {
	return e.msg
}

func (e *errorWrapper) Unwrap() error {
	return e.err
}

func (e *errorWrapper) FullError() string {
	return fmt.Sprintf("%s. Status=%d, package='%s',f='%s'", e.msg, e.status, e.pkg, e.fn)
}

// ToErrorObject returns a struct ready to use in a API response
func (e *errorWrapper) ToErrorObject() *ErrorObject {
	return &ErrorObject{
		Status:  e.status,
		Message: e.msg,
	}
}

// New returns an error that formats as the given text
func New(text string) error {
	p, f := packageAndFunc()
	return &errorWrapper{err: ee.New(text), msg: text, status: http.StatusInternalServerError, pkg: p, fn: f}
}

// Wrap adds some context to an error
func Wrap(e error) error {
	p, f := packageAndFunc()
	return &errorWrapper{err: e, msg: e.Error(), status: http.StatusInternalServerError, pkg: p, fn: f}
}

// NewOperationError wraps an error with additional metadata
func NewOperationError(operation string, s int, e error) error {
	p, f := packageAndFunc()
	return &errorWrapper{err: e, msg: fmt.Sprintf("Error in operation '%s", operation), status: s, pkg: p, fn: f}
}

// see https://stackoverflow.com/questions/25262754/how-to-get-name-of-current-package-in-go
func packageAndFunc() (string, string) {
	pc, _, _, _ := runtime.Caller(2)
	parts := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	pl := len(parts)
	pkg := ""
	fn := parts[pl-1]
	if parts[pl-2][0] == '(' {
		fn = parts[pl-2] + "." + fn
		pkg = strings.Join(parts[0:pl-2], ".")
	} else {
		pkg = strings.Join(parts[0:pl-1], ".")
	}
	return pkg, fn
}
