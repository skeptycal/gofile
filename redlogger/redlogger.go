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
)

type Logger = logrus.Logger

var LogLevel Level = defaultLogLevel

var (
	defaultRedLogColor ansi.Ansi = ansi.NewColor(ansi.Red, ansi.Black, ansi.Bold)
	defaultLogLevel    Level     = logrus.DebugLevel
	// defaultDevLogLevel  Level     = logrus.InfoLevel
	// defaultProdLogLevel Level     = logrus.DebugLevel
	AllLevels []Level = logrus.AllLevels
)

var Log = &logrus.Logger{
	Out:       New(nil, defaultRedLogColor),
	Formatter: new(logrus.TextFormatter),
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

	Log.SetLevel(LogLevel)

	Log.Info("Logger enabled...")
}

// New returns a new instance of RedLogger. The zero value contains
// color as the default wrap color and DevMode as false.
func New(w io.Writer, color ansi.Ansi) *RedLogger {

	if color == nil {
		color = defaultRedLogColor
	}

	r := cli.NewStderr(w)
	r.SetColor(color)
	r.DevMode(false)

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

// Write wraps p with Ansi color codes and writes the result to the buffer.
func (l *RedLogger) Write(p []byte) (int, error) {
	// nn, err := l.Writer.WriteString("--> redlogger Write()") // test
	_, wraperr := l.Writer.WriteString(l.color.String()) // test
	if wraperr != nil {
		return 0, wraperr
	}

	n, err := l.Writer.Write(p)
	if err != nil {
		return n, err
	}

	_, wraperr = l.Writer.WriteString(ansi.Reset)
	if wraperr != nil {
		return n, wraperr
	}

	return n, nil
}

// WriteString wraps p with Ansi color codes and writes the result to the buffer.
func (l *RedLogger) WriteString(s string) (n int, err error) {
	return l.Write([]byte(s))
}
