// Package json implements encoding and decoding of JSON as defined in
// RFC 7159. The mapping between JSON and Go values is described
// in the documentation for the Marshal and Unmarshal functions.
//
// It relies heavily on the standard library "encoding/json" package.
// See "JSON and Go" for an introduction to that package:
//
// https://golang.org/doc/articles/json_and_go.html
package json

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"

	"github.com/skeptycal/gofile"
)

// New returns a JSON structure from the json file. If the
// file does not exist, it will be created.
//
// If no filename is provided, the structure will only
// be manipulated in memory.
//
func New(filename string) (JSON, error) {
	j := jsonStruct{
		name: defaultJSONstructName,
		fi:   nil,
		v:    &jsonMap{},
	}
	if filename == "" {
		return &j, nil
	}

	err := j.AddFile(filename)
	if err != nil {
		return &j, err
	}

	return &j, nil
}

// JSON describes a JSON file and data structure object.
type JSON interface {
	AddFile(filename string) error
	Name() string
	Load() error
	Save() error
	json.Marshaler
	json.Unmarshaler
}

func (j *jsonStruct) Name() string {
	if j.fi != nil {
		return j.fi.Name()
	}
	return "in-memory JSON"
}

// AddFile creates a file to store the underlying JSON
// data if none is already present. It sets the name and
// FileInfo fields and returns any error encountered,
func (j *jsonStruct) AddFile(filename string) error {

	f, err := os.OpenFile(filename, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	absName, err := filepath.Abs(f.Name())
	if err != nil {
		return err
	}

	fi, err := gofile.Stat(absName)
	if err != nil {
		return err
	}

	j.name = absName
	j.fi = fi

	if fi.Size() == 0 {
		return nil
	}

	err = j.Load()
	if err != nil {
		return err
	}
	return nil
}

// Load loads JSON data from the underlying file. If no file
// is associated with the JSON data, an error is returned.
//
// Use method AddFile() to attach a file to a JSON structure
// that did not originally have one.
//
// note: variable/field names should begin with an
// uppercase letter or they will not load correctly
func (j *jsonStruct) Load() error {
	if j.Name() == defaultJSONstructName {
		return os.ErrNotExist
	}

	data, err := os.ReadFile(j.Name())
	if err != nil {
		return err
	}
	return j.UnmarshalJSON(data)
}

// Save saves JSON data to the underlying file. If no file is associated with the JSON data, an error is returned.
//
// Use method AddFile() to attach a file to a JSON structure that did not originally have one.
//
// note: variable/field names should begin with an uppercase letter or they will not load correctly
func (j *jsonStruct) Save() error {
	if j.Name() == defaultJSONstructName {
		return os.ErrNotExist
	}

	data, err := j.MarshalJSON()
	if err != nil {
		return err
	}
	return os.WriteFile(j.Name(), data, 0644)
}

// A FileMode represents a file's mode and permission bits.
// The bits have the same definition on all systems, so that
// information about files can be moved from one system
// to another portably. Not all bits apply to all systems.
// The only required bit is ModeDir for directories.
type FileMode = os.FileMode

// A FileInfo describes a file and is returned by Stat and Lstat.
// type FileInfo = os.FileInfo

// A FileInfo describes a file and is returned by Stat and Lstat.
type FileInfo interface {
	Name() string       // base name of the file
	Size() int64        // length in bytes for regular files; system-dependent for others
	Mode() FileMode     // file mode bits
	ModTime() time.Time // modification time
	IsDir() bool        // abbreviation for Mode().IsDir()
	Sys() interface{}   // underlying data source (can return nil)
}
