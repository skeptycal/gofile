package gofile

import (
	el "github.com/skeptycal/errorlogger"
)

// errorlogger implements error logging to a logrus log
// by default. It is completely compatible with the standard library
// 'log' logger API. It provides an efficient and convenient way to
// log errors with minimal overhead and to temporarily disable or
// enable logging.
//
// A global Log and Err with default behaviors are supplied that
// may be aliased if you wish:
//
//  Log = errorlogger.Log // implements the ErrorLogger interface
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
	// Log is the default global ErrorLogger. It implements the ErrorLogger interface as well as the logrus.Logger interface, which is compatible with the standard library "log" package gofile
	Log = el.Log

	// Err is the logging function for the global ErrorLogger.
	Err = el.Err
)
