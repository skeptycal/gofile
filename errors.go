package gofile

import (
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/skeptycal/basicfile"
)

type GoFileError = basicfile.GoFileError

var NewGoFileError = basicfile.NewGoFileError

// type (
// 	timeout interface {
// 		Timeout() bool
// 	}

// 	// Errer implements the common error methods
// 	// associated with file and system errors.
// 	Errer interface {
// 		Error() string
// 		Unwrap() error
// 		Timeout() bool
// 		// Wrap() error
// 	}

// 	// PathError records an error and the operation
// 	// and file path that caused it.
// 	// PathError = os.PathError

// 	// SyscallError records an error from a specific
// 	// system call.
// 	SyscallError = os.SyscallError
// )

// // NewSyscallError returns, as an error, a new
// // SyscallError with the given system call name
// // and error details.
// //
// // As a convenience, if err is nil, NewSyscallError
// // returns nil.
// var NewSyscallError = os.NewSyscallError

// Portable analogs of some common errors.
//
// Errors returned from this package may be tested against these errors
// with errors.Is.
var (
	// ErrNoAlloc          = basicfile.ErrNoAlloc        // NewGoFileError("memory allocation failure", "", ErrInvalid)
	// ErrFileLocked       = basicfile.ErrFileLocked     // NewGoFileError("file locked", "", ErrClosed)
	ErrNotImplemented   = basicfile.ErrNotImplemented // NewGoFileError("feature not implemented", "", ErrInvalid)
	ErrExist            = fs.ErrExist                 // NewGoFileError("", "", fs.ErrExist)
	ErrNotExist         = fs.ErrNotExist              // NewGoFileError("", "", fs.ErrNotExist)
	ErrPermission       = fs.ErrPermission            // NewGoFileError("", "", fs.ErrPermission)
	ErrClosed           = fs.ErrClosed                // NewGoFileError("", "", fs.ErrClosed)
	ErrInvalid          = fs.ErrInvalid               // NewGoFileError("", "", fs.ErrInvalid)
	ErrNoDeadline       = os.ErrNoDeadline            // NewGoFileError("", "", os.ErrNoDeadline)
	ErrDeadlineExceeded = os.ErrDeadlineExceeded      // NewGoFileError("", "", os.ErrDeadlineExceeded)
	ErrProcessDone      = os.ErrProcessDone           // NewGoFileError("", "", os.ErrProcessDone)
	ErrClosedPipe       = io.ErrClosedPipe            // NewGoFileError("", "", io.ErrClosedPipe)
	ErrNoProgress       = io.ErrNoProgress            // NewGoFileError("", "", io.ErrNoProgress)
	ErrShortBuffer      = io.ErrShortBuffer           // NewGoFileError("", "", io.ErrShortBuffer)
	ErrShortWrite       = io.ErrShortWrite            // NewGoFileError("", "", io.ErrShortWrite)
	ErrUnexpectedEOF    = io.ErrUnexpectedEOF         // NewGoFileError("", "", io.ErrUnexpectedEOF)
	ErrBadPattern       = filepath.ErrBadPattern      // NewGoFileError("", "", filepath.ErrBadPattern)
)

// func SetError(op, path string, err GoFileError) GoFileError {
// 	if op != "" {
// 		err.Op = op
// 	}
// 	if path != "" {
// 		err.Path = path
// 	}
// 	return err
// }

// // NewPathError returns, as an error, a new
// // PathError with the given operation and file
// // path that caused it.
// // The error will be of type *os.PathError
// // As a convenience, if err is nil, NewPathError
// // returns nil.
// func NewPathError(op, path string, err error) Errer {
// 	if err == nil {
// 		return nil
// 	}
// 	return &os.PathError{
// 		Op:   op,
// 		Path: path,
// 		Err:  err,
// 	}
// }

// var (
// 	ErrExist            = fs.ErrExist
// 	ErrNotExist         = fs.ErrNotExist
// 	ErrPermission       = fs.ErrPermission
// 	ErrClosed           = fs.ErrClosed
// 	ErrInvalid          = fs.ErrInvalid
// 	ErrNoDeadline       = os.ErrNoDeadline
// 	ErrDeadlineExceeded = os.ErrDeadlineExceeded
// 	ErrProcessDone      = os.ErrProcessDone
// 	ErrClosedPipe       = io.ErrClosedPipe
// 	ErrNoProgress       = io.ErrNoProgress
// 	ErrShortBuffer      = io.ErrShortBuffer
// 	ErrShortWrite       = io.ErrShortWrite
// 	ErrUnexpectedEOF    = io.ErrUnexpectedEOF
// 	ErrBadPattern       = filepath.ErrBadPattern
// )
