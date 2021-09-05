// Package redlogger implements logging to an io.Writer (default stderr)
// with text wrapped in ANSI escape codes
package redlogger

import (
	"bufio"
	"flag"
	"io"

	"github.com/sirupsen/logrus"
	"github.com/skeptycal/ansi"
	"github.com/skeptycal/cli"
	"github.com/skeptycal/errorlogger"
)

type Logger = logrus.Logger

var LogLevel Level = defaultLogLevel

var (
	defaultRedLogColor ansi.Ansi = ansi.NewColor(ansi.Red, ansi.Black, ansi.Bold)
	defaultDevMode     bool      = false
	defaultLogLevel    Level     = logrus.InfoLevel
	def                          = errorlogger.NewTextFormatter()
	// defaultDevLogLevel  Level     = logrus.InfoLevel
	// defaultProdLogLevel Level     = logrus.DebugLevel
	AllLevels []Level = logrus.AllLevels

	llog         = errorlogger.New()
	RedFormatter = new(logrus.TextFormatter)
)

var logrusLogger = &logrus.Logger{
	Out:       New(nil, defaultRedLogColor, defaultDevMode),
	Formatter: RedFormatter,
	Hooks:     make(logrus.LevelHooks),
	Level:     defaultLogLevel,
}

// Level may be one of PanicLevel, FatalLevel, ErrorLevel, WarnLevel, InfoLevel, DebugLevel, TraceLevel
type Level = logrus.Level

func init() {
	logLevelFlag := flag.String("log level", "INFO", "set the log level. (TRACE, DEBUG, INFO, WARN, ERROR, FATAL, PANIC)")
	// log.SetFormatter(new(logrus.TextFormatter))

	LogLevel, err := logrus.ParseLevel(*logLevelFlag)
	if err != nil {
		LogLevel = defaultLogLevel
	}

	RedFormatter.DisableColors = true // disable colors so only the redlogger wrap colors show

	logrusLogger.SetLevel(LogLevel)

	logrusLogger.Info("Logger enabled...")
}

// New returns a new instance of RedLogger. The zero value contains
// color as the default wrap color and DevMode as false.
func New(w io.Writer, color ansi.Ansi, devMode bool) *RedLogger {

	if color == nil {
		color = defaultRedLogColor
	}

	r := cli.NewStderr(w)
	r.SetColor(color)
	r.DevMode(devMode)

	return &RedLogger{color, bufio.NewWriter(r)}
}

// RedLogger implements buffering for an io.Writer object that
// always wraps output in an ANSI color.
//
// If an error occurs writing to a Writer, no more data will be
// accepted and all subsequent writes, and Flush, will return the error.
// After all data has been written, the client should call the
// Flush method to guarantee all data has been forwarded to
// the underlying io.Writer.
type RedLogger struct {
	color ansi.Ansi // Color(83)
	*bufio.Writer
}

// Write wraps p with Ansi color codes and writes the contents
// of p into the buffer. It returns the number of bytes written
// and any error encountered. If n < len(p), it also returns an
// error explaining why the write is short.
//
// The writing of ANSI wrap codes is checked for validity, but
// these wrrites are not counted in the return value for n.
// This allows 'transparent' ANSI wrapping where the caller is
// unaware of the underlying actions and only receives error
// messages based on the original data.
func (l *RedLogger) Write(p []byte) (n int, err error) {
	_, wraperr := l.Writer.WriteString(l.color.String())
	if wraperr != nil {
		return 0, wraperr
	}

	n, err = l.Writer.Write(p)
	if err != nil {
		return n, err
	}

	_, wraperr = l.Writer.WriteString(ansi.Reset)
	if wraperr != nil {
		return n, wraperr
	}

	if n < len(p) {
		return n, io.ErrShortWrite
	}

	return n, nil
}

// WriteString wraps s with Ansi color codes and writes the contents
// of s into the buffer. It returns the number of bytes written
// and any error encountered. If n < len(s), it also returns an
// error explaining why the write is short.
//
// The writing of ANSI wrap codes is checked for validity, but
// these wrrites are not counted in the return value for n.
// This allows 'transparent' ANSI wrapping where the caller is
// unaware of the underlying actions and only receives error
// messages based on the original data.
func (l *RedLogger) WriteString(s string) (n int, err error) {
	// todo check if l.Writer implements io.StringWriter?
	// return l.Write([]byte(s))
	_, wraperr := l.Writer.WriteString(l.color.String())
	if wraperr != nil {
		return 0, wraperr
	}

	n, err = l.Writer.WriteString(s)
	if err != nil {
		return n, err
	}
	_, wraperr = l.Writer.WriteString(ansi.Reset)
	if wraperr != nil {
		return n, wraperr
	}

	if n < len(s) {
		return n, io.ErrShortWrite
	}

	return n, nil
}
