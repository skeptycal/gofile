package main

import (
	"errors"
	"fmt"

	"github.com/skeptycal/gofile/errorlogger"
)

var (
	EL  errorlogger.ErrorLogger = errorlogger.EL
	Err func(err error) error   = errorlogger.Err
	// sb strings.Builder         = strings.Builder{}
)

func main() {

	const message = `
Example for package errorlogger:
The first line is an error message displayed by the logrus logger package.
The second line is a fmt.Println acknowledgement.
The third line is a log.info message from the logger.

`

	fmt.Println("")

	err := errors.New("fake error")

	fmt.Printf("the fake error occurred: %s\n", Err(err))

}
