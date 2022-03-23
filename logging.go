package gofile

import "github.com/skeptycal/goutil/errorlogger"

// errorlogger implements error logging to a logrus log
// by default. It is completely compatible with the standard library
// 'log' logger API. It provides an efficient and convenient way to
// log errors with minimal overhead and to temporarily disable or
// enable logging.
//
// A global GFlog and Err with default behaviors are supplied that
// may be aliased if you wish:
//
//  GFlog = errorlogger.New() // implements the ErrorLogger interface
//  Err = errorlogger.Err // trigger function
//
// ErrorLogger is:
//  type ErrorLogger interface {
// 	 	Disable()
// 	 	Enable()
// 	 	Err(err error) error
// 	 	SetLoggerFunc(fn func(args ...interface{}))
// 	 	SetErrorType(errType error)
//  }

// Log is the default global ErrorLogger. It
// implements the ErrorLogger interface as well
// as the logrus.Logger interface, which is
// compatible with the standard library "log" package.
//
// Err is the default errorlogger.ErrorLogger for
// the gofile package.
var GFlog = errorlogger.New()

// Err is the logging function for the
// global ErrorLogger.
//
// Err logs an error to the provided logger,
// if it is enabled, and returns the error
// unchanged.
var Err = GFlog.Err
