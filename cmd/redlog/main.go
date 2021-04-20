package main

import (
	"os"

	"github.com/skeptycal/gofile/redlogger"
)

var r = redlogger.New(os.Stderr, nil)

func main() {
	defer r.Flush()
	_, _ = r.WriteString("Hello World!")
}
