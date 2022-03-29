package gofile

import (
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
	ErrNoAlloc          = basicfile.ErrNoAlloc          // NewGoFileError("memory allocation failure", "", ErrInvalid)
	ErrNotImplemented   = basicfile.ErrNotImplemented   // NewGoFileError("feature not implemented", "", ErrInvalid)
	ErrFileLocked       = basicfile.ErrFileLocked       // NewGoFileError("file locked", "", ErrClosed)
	ErrExist            = basicfile.ErrExist            // NewGoFileError("", "", fs.ErrExist)
	ErrNotExist         = basicfile.ErrNotExist         // NewGoFileError("", "", fs.ErrNotExist)
	ErrPermission       = basicfile.ErrPermission       // NewGoFileError("", "", fs.ErrPermission)
	ErrClosed           = basicfile.ErrClosed           // NewGoFileError("", "", fs.ErrClosed)
	ErrInvalid          = basicfile.ErrInvalid          // NewGoFileError("", "", fs.ErrInvalid)
	ErrNoDeadline       = basicfile.ErrNoDeadline       // NewGoFileError("", "", os.ErrNoDeadline)
	ErrDeadlineExceeded = basicfile.ErrDeadlineExceeded // NewGoFileError("", "", os.ErrDeadlineExceeded)
	ErrProcessDone      = basicfile.ErrProcessDone      // NewGoFileError("", "", os.ErrProcessDone)
	ErrClosedPipe       = basicfile.ErrClosedPipe       // NewGoFileError("", "", io.ErrClosedPipe)
	ErrNoProgress       = basicfile.ErrNoProgress       // NewGoFileError("", "", io.ErrNoProgress)
	ErrShortBuffer      = basicfile.ErrShortBuffer      // NewGoFileError("", "", io.ErrShortBuffer)
	ErrShortWrite       = basicfile.ErrShortWrite       // NewGoFileError("", "", io.ErrShortWrite)
	ErrUnexpectedEOF    = basicfile.ErrUnexpectedEOF    // NewGoFileError("", "", io.ErrUnexpectedEOF)
	ErrBadPattern       = basicfile.ErrBadPattern       // NewGoFileError("", "", filepath.ErrBadPattern)
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

// // ErrExist            = fs.ErrExist
// // ErrNotExist         = fs.ErrNotExist
// // ErrPermission       = fs.ErrPermission
// // ErrClosed           = fs.ErrClosed
// // ErrInvalid          = fs.ErrInvalid
// // ErrNoDeadline       = os.ErrNoDeadline
// // ErrDeadlineExceeded = os.ErrDeadlineExceeded
// // ErrProcessDone      = os.ErrProcessDone
// // ErrClosedPipe       = io.ErrClosedPipe
// // ErrNoProgress       = io.ErrNoProgress
// // ErrShortBuffer      = io.ErrShortBuffer
// // ErrShortWrite       = io.ErrShortWrite
// // ErrUnexpectedEOF    = io.ErrUnexpectedEOF
// // ErrBadPattern       = filepath.ErrBadPattern
