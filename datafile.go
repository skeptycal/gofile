package gofile

import (
	"bufio"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

type DataFile interface {
	File
	Data() ([]byte, error)
}

type TextFile interface {
	File
	Text() string
	Lines() (retval []string, err error)
}

// textfile is a file type that is specialized for utf-8 text
type textfile struct {
	basicfile
	linesep   byte
	recordsep byte
}

func (d *textfile) Text() string {
	buf, err := d.Data()
	if err != nil {
		log.Fatal(err)
	}
	return string(buf)
}

func (d *textfile) Lines() (retval []string, err error) {

	file, err := os.Open(d.Name())
	if err != nil {
		log.Errorf("failed to open %s", d.Name())
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		retval = append(retval, scanner.Text())
	}

	return retval, scanner.Err()
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
