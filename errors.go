package gofile

import (
	"errors"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/skeptycal/errorlogger"
)

var (
	log = errorlogger.Log

	// e is an errorlogger.ErrorLogger with default settings
	e = Err

	FSErr = errorlogger.Err
)

// Portable analogs of some common errors.
//
// Errors returned from this package may be tested against these errors
// with errors.Is.
var (
	ErrNoAlloc          = errors.New("failed to allocate memory for file")
	ErrNotImplemented   = errors.New("not implemented")
	ErrFsInvalid        = fs.ErrInvalid
	ErrPermission       = fs.ErrPermission
	ErrExist            = fs.ErrExist
	ErrNotExist         = fs.ErrNotExist
	ErrClosed           = fs.ErrClosed
	ErrInvalid          = os.ErrInvalid
	ErrNoDeadline       = os.ErrNoDeadline
	ErrDeadlineExceeded = os.ErrDeadlineExceeded
	ErrProcessDone      = os.ErrProcessDone
	ErrClosedPipe       = io.ErrClosedPipe
	ErrNoProgress       = io.ErrNoProgress
	ErrShortBuffer      = io.ErrShortBuffer
	ErrShortWrite       = io.ErrShortWrite
	ErrUnexpectedEOF    = io.ErrUnexpectedEOF
)

type (
	timeout interface {
		Timeout() bool
	}

	// Errer implements the common error methods
	// associated with file and system errors.
	Errer interface {
		Error() string
		Unwrap() error
		Timeout() bool
	}

	// PathError records an error and the operation
	// and file path that caused it.
	PathError = os.PathError

	// SyscallError records an error from a specific
	// system call.
	SyscallError = os.SyscallError
)

var aaa = filepath.ErrBadPattern

// NewSyscallError returns, as an error, a new
// SyscallError with the given system call name
// and error details.
// As a convenience, if err is nil, NewSyscallError
// returns nil.
var NewSyscallError = os.NewSyscallError

// NewPathError returns, as an error, a new
// PathError with the given operation and file
// path that caused it.
// The error will be of type *os.PathError
// As a convenience, if err is nil, NewPathError
// returns nil.
func NewPathError(op, path string, err error) Errer {
	if err == nil {
		return nil
	}
	return &PathError{
		Op:   op,
		Path: path,
		Err:  err,
	}
}
