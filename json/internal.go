package json

import (
	"os"

	"github.com/skeptycal/gofile/errorlogger"
)

const (
	defaultJSONstructName = "in-memory JSON"
)

var log = errorlogger.Log
var EL = errorlogger.EL
var Err = errorlogger.Err

// jsonStruct implements a JSON mapping with os.FileInfo included.
type jsonStruct struct {
	name string `default:"defaultJSONstructName"`
	fi   os.FileInfo
	v    *jsonMap
}
