package gofile

import (
	"errors"
	"io/fs"
	"os"

	"github.com/skeptycal/errorlogger"
)

var (
	FSErr         = errorlogger.Err
	ErrNoAlloc    = errors.New("failed to allocate memory for file")
	ErrClosed     = fs.ErrClosed
	ErrInvalid    = fs.ErrInvalid
	ErrPermission = fs.ErrPermission
	ErrExist      = NewGoFileError("gofile", "", fs.ErrExist)
	ErrNotExist   = NewGoFileError("gofile", "", fs.ErrNotExist)
)

// NewPathError records an error and the operation and file path that caused it.
//  type PathError struct {
//  	Op   string
//  	Path string
//  	Err  error
//  }
func NewPathError(op, path string, err error) *PathError {
	return &PathError{
		Op:   op,
		Path: path,
		Err:  err,
	}
}

func gferr(path, op string, eerr error) error {
	if eerr == nil {
		return nil
	}

	path = "(gofile error) " + path

	err, ok := eerr.(*os.PathError)
	if !ok {
		return NewGoFileError(path, op, eerr)
	}

	if op == "" {
		op = err.Op
	}

	if path == "" {
		path = err.Path
	}

	pe := &os.PathError{path, op, eerr}

	return &GoFileError{path, op, pe}
}

func opErr(op string, err error) error {
	if err == nil {
		return nil
	}

	return gferr("", op, err)
}

type GoFileError struct {
	Op   string
	Path string
	Err  error
}

func (e *GoFileError) Error() string {
	return "gofile error:" + e.Op + " " + e.Path + ": " + e.Err.Error()
}

func (e *GoFileError) Unwrap() error { return e.Err }

// Timeout reports whether this error represents a timeout.
func (e *GoFileError) Timeout() bool {
	t, ok := e.Err.(interface{ Timeout() bool })
	return ok && t.Timeout()
}

// NewGoFileError returns a new GoFileError which is also
// an os.PathError
func NewGoFileError(op, path string, err error) *GoFileError {
	fse := &PathError{op, path, err}
	op = "(gofile error) " + op
	return &GoFileError{op, path, fse}
}

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

// PathError records an error and the operation and file path that caused it.
//  type PathError struct {
//  	Op   string
//  	Path string
//  	Err  error
//  }
//
//  interface {
// 		Error() string
// 		Unwrap() error
// 		Timeout() bool
//	}
type PathError = os.PathError