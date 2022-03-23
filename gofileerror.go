package gofile

import (
	"fmt"
	"os"
	"reflect"

	"github.com/pkg/errors"
)

var errorType = reflect.TypeOf((*error)(nil)).Elem()

type (

	// GoFileError records an error from a specific
	// GoFile function call. If there is a relevant
	// path, the path will be set and the error will
	// be of type *os.PathError
	GoFileError struct {
		Op   string
		Path string
		Err  error
	}
)

// NewGoFileError returns, as an error, a new
// GoFileError which implements the gofile.Errer
// interface. If there is a relevant path,
// the path will be set and the error will
// wrap *os.PathError.
// As a convenience, if err is nil, NewGoFileError
// returns nil.
func NewGoFileError(op, path string, err error) *GoFileError {
	if err == nil {
		return nil
	}

	if path != "" {
		err = &PathError{Op: op, Path: path, Err: err}

	} else {
		var wderr error
		path, wderr = os.Getwd()
		if wderr != nil {
			path = "unknown"
		}
	}

	return &GoFileError{
		Op:   "gofile: " + op,
		Path: path,
		Err:  &PathError{Op: op, Path: path, Err: err},
	}
}

const goFileErrorPrefix = "gofile: "

func prependGoFilePrefix(msg string) string {
	return goFileErrorPrefix + msg
}

func checkGoFilePath(path string) string {
	if path == "" {
		return PWD()
	}
	if Exists(path) {
		return "invalid path"
	}
	return path
}

func gferr(path, op string, err error) error {
	var gfe = new(GoFileError)

	v, ok := err.(*os.PathError)
	if ok {
		op = v.Op
		path = v.Path
		err = v
	}

	gfe = &GoFileError{Op: prependGoFilePrefix(op), Path: path, Err: err}
	gfe.Wrap("gofile stack trace")

	return gfe
}

func (e *GoFileError) Error() string {
	return "gofile error: - " + e.Op + " (path: " + e.Path + "): " + e.Err.Error()
}

// Wrap replaces the underlying error (err) with a
// wrapper annotating it with a stack trace at the
// point Wrap is called, and the supplied message.
//
// A pointer to the new, wrapped error is returned.
// If err is nil, Wrap returns nil and performs no
// other operations.
func (e *GoFileError) Wrap(message string) *GoFileError {
	if e.Err == nil {
		return nil
	}

	e.Err = errors.Wrap(e.Err, goFileErrorPrefix+message)

	return e
}

func (e *GoFileError) Wrapf(format string, args ...interface{}) *GoFileError {
	return e.Wrap(fmt.Sprintf(format, args...))
}

func (e *GoFileError) WithMessage(msg string) *GoFileError {
	e.Err = errors.WithMessage(e.Err, msg)
	return e
}

// Unwrap returns the result of calling the Unwrap
// method on err, if err's type contains an Unwrap
// method.
// Otherwise, Unwrap returns nil.
func (e *GoFileError) Unwrap() error {
	eerr := errors.Unwrap(e.Err)
	return eerr
}

// As finds the first error in err's chain that matches target, and if one is found, sets
// target to that error value and returns true. Otherwise, it returns false.
//
// The chain consists of err itself followed by the sequence of errors obtained by
// repeatedly calling Unwrap.
//
// An error matches target if the error's concrete value is assignable to the value
// pointed to by target, or if the error has a method As(interface{}) bool such that
// As(target) returns true. In the latter case, the As method is responsible for
// setting target.
//
// An error type might provide an As method so it can be treated as if it were a
// different error type.
//
// As panics if target is not a non-nil pointer to either a type that implements
// error, or to any interface type.
func (e *GoFileError) As(any) {
	return
}

// Timeout reports whether this error represents a timeout.
func (e *GoFileError) Timeout() bool {
	t, ok := e.Err.(timeout)
	return ok && t.Timeout()
}
