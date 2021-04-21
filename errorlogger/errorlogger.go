// Package errorlogger implements error logging to a logrus log
// (or a standard library log) by providing a convenient way to
// log errors and to temporarily disable/enable logging.
//
// A global EL and Err with default behaviors are supplied that
// may be aliased if you wish:
//  EL = errorlogger.EL
//  Err = errorlogger.Err
//
// If you do not intend to use any options or disable the logger,
// it may be more convenient to use a function alias to call the
// most common method, Err(), like this:
//  var Err = errorlogger.New().Err
// then, just call the function:
//  err := someProcess(stuff)
//  if err != nil {
//   return Err(err)
//  }
//
// Either way, the default ErrorLogger is enabled and ready to go:
//  EL := errorlogger.New() // enabled by default
//  Err := EL.Err
//
// If a private ErrorLogger is desired, or if name collisions with
// Err cause conflicts, you may implement your own.
//  myErr := errorlogger.New()
//  err := myErr.Err
//
// Example:
//  f, err := os.Open("somefile.txt")
//  if err != nil {
// 	 return nil, e.Err(err) // avoids additional logging steps
//  }
//  e.Disable() // can be disabled and enabled as desired
package errorlogger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Level type values: Panic, Fatal, Error, Warn, Info, Debug, Trace
type Level = logrus.Level

const (
	defaultLogLevel    logrus.Level = logrus.InfoLevel
	defaultEnabled     bool         = true
	defaultWrapMessage string       = ""
)

// Log implements the default logrus error logger
//
// Reference: https://github.com/sirupsen/logrus/
var Log *logrus.Logger = defaultLogger

// Defaults for ErrorLogger
var (
	defaultLogFunc       loggerFunc       = defaultLogger
	defaultTextFormatter logrus.Formatter = new(logrus.TextFormatter)
	defaultJSONFormatter logrus.Formatter = new(logrus.JSONFormatter)
	defaultFormatter     logrus.Formatter = defaultTextFormatter
	defaultLogger                         = &logrus.Logger{

		Out: os.Stderr,

		Formatter: defaultFormatter,

		Hooks: make(logrus.LevelHooks),

		Level: defaultLogLevel,
	}
)

// New returns a new ErrorLogger with logging enabled.
func New() ErrorLogger {
	return NewWithOptions(defaultEnabled, defaultLogFunc, defaultWrapMessage)
}

func NewWithOptions(enabled bool, fn loggerFunc, wrap string) ErrorLogger {
	e := errorLogger{}
	if enabled {
		e.Enable()
	} else {
		e.Disable()
	}

	e.SetLoggerFunc(fn)
	e.SetErrorType(wrap)

	return &e
}

// ErrorLogger implements error logging to a logrus log
// (or a standard library log) by providing a convenient way to
// log errors and to temporarily disable/enable logging.
type ErrorLogger interface {

	// Disable disables logging and sets a no-op function for
	// Err() to prevent slowdowns while logging is disabled.
	Disable()

	// Enable enables logging and restores the Err() logging functionality.
	Enable()

	// EnableText enables text formatting of log errors (default)
	EnableText()

	// EnableJSON enables JSON formatting of log errors
	EnableJSON()

	// LogLevel sets the logging level from a string value.
	// Allowed values: Panic, Fatal, Error, Warn, Info, Debug, Trace
	SetLogLevel(lvl string)

	// Err logs an error to the provided logger, if it is enabled,
	// and returns the error unchanged.
	Err(err error) error

	// SetLoggerFunc allows setting of the logger function.
	// The default is log.Error(), which is compatible with
	// the standard library log package and logrus.
	SetLoggerFunc(fn loggerFunc)

	// // SetErrorType allows ErrorLogger to wrap errors in a
	// // specified custom type. For example, if you want all errors
	// // returned to be of type *os.PathError
	// SetErrorType(errType error)
}

// errorLogger implements ErrorLogger with logrus or the
// standard library log package.
type errorLogger struct {
	enabled bool                  // `default:"true"`
	wrap    string                // `default:""` // "" = disabled
	errFunc func(err error) error // `default:"()yesErr"`
	logFunc loggerFunc            // `default:"logrus.New()"`
	logger  *logrus.Logger
}

// SetErrorType allows ErrorLogger to wrap errors in a specified custom message.
// Setting wrap == "" will disable wrapping of errors.
func (e *errorLogger) SetErrorType(message string) {
	// TODO - not completely implemented in Err()
	e.wrap = message
}

func (e *errorLogger) EnableText() {
	e.logger.SetFormatter(defaultTextFormatter)
}

func (e *errorLogger) EnableJSON() {
	e.logger.SetFormatter(defaultJSONFormatter)
}

func (e *errorLogger) SetLogLevel(lvl string) {
	level, err := logrus.ParseLevel(lvl)
	if err != nil {
		Err(err)
		return
	}
	e.logger.SetLevel(level)
}
