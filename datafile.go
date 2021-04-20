package gofile

import (
	"os"
	"path/filepath"
)

type DataFile interface {
	File
	Data() ([]byte, error)
}

// datafile is a file type that is specialized for binary data
type datafile struct {
	basicfile
}

func NewDataFile(filename string) (DataFile, error) {
	src, err := os.Stat(filename)
	if err != nil {
		return nil, err
	}

	if !src.Mode().IsRegular() {
		return nil, ErrNotRegular
	}

	name, err := filepath.Abs(src.Name())
	if err != nil {
		return nil, &PathError{Op: "abs", Path: src.Name(), Err: err}
	}

	df := &datafile{}

	df.providedName = filename
	df.name = name
	df.size = src.Size()
	// df.info = src
	df.FileInfo = src

	return df, nil
}
