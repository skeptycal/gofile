package gofile

import log "github.com/sirupsen/logrus"

var Err = NewErrorLogger()

// NewErrorLogger returns a new ErrorLogger with logging enabled.
func NewErrorLogger() ErrorLogger {
	e := &errorLogger{}
	e.SetLoggerFunc(log.Error)
	e.Enable()
	return e
}

// ErrorLogger implements error logging to a logrus log by providing
// a convenient way to log errors and to temporarily disable logging.
//
// The default ErrorLogger is enabled and ready to go:
//  Err := errorlogger.New() // enabled by default
//
// This step is not required to use the defaults; a global ErrorLogger
// named Err is provided. If a private ErrorLogger is desired, or if name
// collisions with Err cause conflicts, you may implement your own.
//
// Example:
//  f, err := os.Open("somefile.txt")
//  if err != nil {
// 	 return nil, e.Err(err) // avoids additional logging steps
//  }
//  e.Disable() // can be disabled and enabled as desired
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

	// TODO - not implemented
	// SetErrorType allows ErrorLogger to wrap errors in a
	// specified custom type.
	// SetErrorType(errType error)
}

// errorLogger implements ErrorLogger with logrus or the
// standard library log package.
type errorLogger struct {
	enabled bool                      // `default:"true"`
	errType error                     // TODO - not implemented
	errFunc func(err error) error     // `default:"()yesErr"`
	logFunc func(args ...interface{}) // `default:"log.Error"`
}

// Disable disables logging and sets a no-op function for
// Err() to prevent slowdowns while logging is disabled.
func (e *errorLogger) Disable() {
	e.enabled = false
	e.errFunc = e.noErr
}

// Enable enables logging and restores the Err() logging functionality.
func (e *errorLogger) Enable() {
	e.enabled = true
	e.errFunc = e.yesErr
}

// Err logs an error to the provided logger, if it is enabled,
// and returns the error unchanged.
func (e *errorLogger) Err(err error) error {
	return e.errFunc(err)
}

// SetLoggerFunc allows setting of the logger function.
// The default is log.Error(err), which is compatible with
// the standard library log package and logrus.
//
// The function signature must be
//  func(args ...interface{}).
func (e *errorLogger) SetLoggerFunc(fn func(args ...interface{})) {
	e.logFunc = fn
}

// noErr is a no-op placeholder for Err
func (e *errorLogger) noErr(err error) error {
	return err
}

// Err logs errors and passes them through unchanged.
func (e *errorLogger) yesErr(err error) error {
	if err != nil {
		e.logFunc(err)
	}
	return err
}

// SetErrorType allows ErrorLogger to wrap errors in a specified custom type.
// TODO - make public once implementation is complete
func (e *errorLogger) setErrorType(errType error) {
	// TODO - not implemented
	e.errType = errType
}
