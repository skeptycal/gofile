package gofile

import "github.com/skeptycal/gofile/errorlogger"

// Package errorlogger implements error logging to a logrus log
// by default. It is completely compatible with the standard library
// 'log' logger API. It provides an efficient and convenient way to
// log errors with minimal overhead and to temporarily disable or
// enable logging.
//
// A global EL and Err with default behaviors are supplied that
// may be aliased if you wish:
//
//  EL = errorlogger.EL // implements the ErrorLogger interface
//  Err = errorlogger.Err // trigger function
//
// ErrorLogger is:
//  type ErrorLogger interface {
// 	 Disable()
// 	 Enable()
// 	 Err(err error) error
// 	 SetLoggerFunc(fn func(args ...interface{}))
//
// 	// SetErrorType allows ErrorLogger to wrap errors in a
// 	// specified custom type. For example, if you want all errors
//  // returned to be of type *os.PathError
// 	 SetErrorType(errType error)
// }
var (
	EL  = errorlogger.EL
	Err = errorlogger.Err
)

type ErrorLogger interface {

	// Disable disables logging and sets a no-op function for
	// Err() to prevent slowdowns while logging is disabled.
	Disable()

	// Enable enables logging and restores the Err() logging functionality.
	Enable()

	// Err logs an error to the provided logger, if it is enabled,
	// and returns the error unchanged.
	Err(err error) error

	// SetLoggerFunc allows setting of the logger function.
	// The default is log.Error(), which is compatible with
	// the standard library log package and logrus.
	SetLoggerFunc(fn func(args ...interface{}))

	Wrap(err error, message string) error

	// TODO - not implemented
	// SetErrorType allows ErrorLogger to wrap errors in a
	// specified custom type.
	// SetErrorType(errType error)
}
