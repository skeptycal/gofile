// Package redlogger implements logging to an io.Writer (default stderr)
// with text wrapped in ANSI escape codes
package redlogger

import (
	"bufio"
	"io"
	"os"

	"github.com/sirupsen/logrus"
	ansi "github.com/skeptycal/ansi"
)

var (
	defaultRedLogColor  ansi.Ansi      = ansi.NewColor(ansi.Red, ansi.BlackBackground, ansi.Bold)
	defaultLogLevel     logrus.Level   = logrus.InfoLevel
	defaultRedLogOutput ansi.CLI       = ansi.NewStderr(os.Stderr)
	log                 *logrus.Logger = &logrus.Logger{
		Out: New(defaultRedLogOutput, defaultRedLogColor),
	}
)

func init() {

	l := logrus.New()
	var log = &logrus.Logger{
		Out:       New(nil, defaultRedLogColor),
		Formatter: new(logrus.TextFormatter),
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.InfoLevel,
	}
	// log.SetFormatter(new(logrus.TextFormatter))
	log.SetLevel(defaultLogLevel)

	log.Info("RedLogger enabled...")
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
	nn, err := l.Writer.WriteString(l.color)                 // test
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
