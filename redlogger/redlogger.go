// Package redlogger implements logging to an io.Writer (default stderr)
// with text wrapped in ANSI escape codes
package redlogger

import (
	"bufio"
	"flag"
	"io"

	"github.com/sirupsen/logrus"
	ansi "github.com/skeptycal/ansi"
)

type Logger = logrus.Logger

var (
	defaultRedLogColor  ansi.Ansi = ansi.NewColor(ansi.Red, ansi.Black, ansi.Bold)
	defaultDevLogLevel  Level     = logrus.InfoLevel
	defaultProdLogLevel Level     = logrus.DebugLevel
	AllLevels                     = logrus.AllLevels
)

var Log = &logrus.Logger{
	Out:       New(nil, defaultRedLogColor),
	Formatter: new(logrus.TextFormatter),
	Hooks:     make(logrus.LevelHooks),
	Level:     logrus.InfoLevel,
}

type (
	// PanicLevel, FatalLevel, ErrorLevel, WarnLevel, InfoLevel, DebugLevel, TraceLevel,
	Level = logrus.Level
)

func init() {

	logLevelFlags := flag.String("log level", "INFO", "set the log level. (INFO, DEBUG, WARN, ERROR, FATAL)")
	// log.SetFormatter(new(logrus.TextFormatter))
	Log.SetLevel(defaultLogLevel)

	Log.Info("RedLogger enabled...")
}

func New(w io.Writer, color ansi.Ansi) *RedLogger {

	if color == nil {
		color = defaultRedLogColor
	}

	r := ansi.NewStderr(w)
	r.SetColor(color)
	r.DevMode(false)

	return &RedLogger{color, bufio.NewWriter(w)}
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
func (l *RedLogger) Write(p []byte) (n int, err error) {
	nn, err := l.Writer.WriteString("--> redlogger Write()") // test
	nn, err = l.Writer.WriteString(l.color.String())         // test
	if err != nil {
		return 0, err
	}

	n += nn

	nn, err = l.Writer.Write(p)
	if err != nil {
		return n, err
	}

	n += nn

	nn, err = l.Writer.WriteString(ansi.Reset)
	if err != nil {
		return n, err
	}

	return n + nn, nil
}

// WriteString wraps p with Ansi color codes and writes the result to the buffer.
func (l *RedLogger) WriteString(s string) (n int, err error) {
	return l.Write([]byte(s))
}
