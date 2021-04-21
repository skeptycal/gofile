package main

import (
	"errors"
	"fmt"

	"github.com/skeptycal/ansi"
	"github.com/skeptycal/cli"
	"github.com/skeptycal/gofile/errorlogger"
)

var (
	EL  errorlogger.ErrorLogger = errorlogger.EL
	Err func(err error) error   = EL.Err
	// sb strings.Builder         = strings.Builder{}
)

func main() {

	cli := cli.New()

	color := ansi.NewColor(2, 0, 1)

	const message = `
%sExample for package errorlogger:
The first line is an error message displayed by the logrus logger package.
The second line is a fmt.Println acknowledgement.
The third line is a log.info message from the logger.

`
	cli.SetColor(color)
	cli.Printf(message, color)

	err := errors.New("fake error")

	fmt.Printf("the fake error occurred: %s\n", Err(err))

}
