package fs

import (
	"errors"
	"io/fs"
	"os"

	"github.com/skeptycal/errorlogger"
)

var (
	Err           = errorlogger.Err
	ErrNoAlloc    = errors.New("failed to allocate memory for file")
	ErrInvalid    = fs.ErrInvalid
	ErrPermission = fs.ErrPermission
	ErrExist      = fs.ErrExist
	ErrNotExist   = fs.ErrNotExist
	ErrClosed     = fs.ErrClosed
)

// PathErrorWrapper impliments an interface to fs.PathError
// which records an error and the operation and file path that caused it.
//
// Any error passed in is wrapped with *fs.PathError as well
// as GoFileError and may be tested with errors.Is().
//
// Use NewPathError to create a new PathErrorWrapper.
type PathErrorWrapper interface {
	Error() string
	Unwrap() error
	Timeout() bool
}

func NewPathError(op, path string, err error) PathErrorWrapper {
	op = "(gofile.fs) " + op
	err =
	return &PathError{op, path, err}
}

// PathError records an error and the operation and file path that caused it.

type PathError os.PathError

// type PathError struct {
// 	Op   string
// 	Path string
// 	Err  error
// }

func (e *PathError) Error() string { return e.Op + " " + e.Path + ": " + e.Err.Error() }

func (e *PathError) Unwrap() error { return e.Err }

// Timeout reports whether this error represents a timeout.
func (e *PathError) Timeout() bool {
	t, ok := e.Err.(interface{ Timeout() bool })
	return ok && t.Timeout()
}
