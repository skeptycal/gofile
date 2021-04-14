// Package json implements json serialization and deserialization.
package json

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/skeptycal/gofile"
)

type jsonMap map[string]interface{}

// Load loads and returns a JSON structure representing the json file.
func Load(filename string) (JSON, error) {
	fi := gofile.Stat(filename)
	if fi == nil {
		return nil, nil
	}

	j := &jsonStruct{fi, &jsonMap{}}
	err := j.ReadFile()
	if err != nil {
		return nil, err
	}
	return j, err
}

func New(filename string) (JSON, error) {
	fi := gofile.Stat(filename)
	if fi != nil {
		return nil, os.ErrExist
	}

	j := &jsonStruct{fi, &jsonMap{}}
	err := j.ReadFile()
	if err != nil {
		return nil, err
	}
	return j, err
}

// JSON describes a JSON file and data structure object.
type JSON interface {
	ReadFile() error
	Name() string
	Save() error
	Size() int64
	json.Marshaler
	json.Unmarshaler
}

// jsonStruct implements a JSON mapping with os.FileInfo included.
type jsonStruct struct {
	os.FileInfo
	v *jsonMap
}

// Load loads JSON data from the underlying file
//
// note: variable/field names should begin with an
// uppercase letter or they will not load correctly
func (j *jsonStruct) ReadFile() error {
	data, err := ioutil.ReadFile(j.Name())
	if err != nil {
		return err
	}
	return j.UnmarshalJSON(data)
}

// Save saves JSON data to the underlying file
//
// note: variable/field names should begin with an
// uppercase letter or they will not load correctly
func (j *jsonStruct) Save() error {
	data, err := j.MarshalJSON()
	if err != nil {
		return err
	}
	return ioutil.WriteFile(j.Name(), data, 0644)
}

// Unmarshaler is the interface implemented by types
// that can unmarshal a JSON description of themselves.
// The input can be assumed to be a valid encoding of
// a JSON value. UnmarshalJSON must copy the JSON data
// if it wishes to retain the data after returning.
//
// By convention, to approximate the behavior of Unmarshal,
// Unmarshalers implement UnmarshalJSON([]byte("null")) as
// a no-op.
func (j *jsonStruct) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, j.v)
}

// MarshalJSON implements the json.Marshaler interface
// and returns the JSON encoding of itself.
//
func (j *jsonStruct) MarshalJSON() ([]byte, error) {
	return json.Marshal(j.v)
}
