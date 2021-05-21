package gofile

import (
	"bufio"
	"path/filepath"

	"github.com/skeptycal/gofile/basicfile"
)

const (
	defaultLineSep = newline
	defaultListSep = filepath.ListSeparator // Colon

)

type TextFile interface {
	basicfile.BasicFile
	Text() string
	Lines() (retval []string, err error)
	Sep(c byte)
}

// textfile is a file type that is specialized for utf-8 text
type textfile struct {
	basicfile.Basicfile
	linesep   byte
	recordsep byte
}

func (d *textfile) Sep(c byte) {
	d.linesep = c
}

func (d *textfile) RecordSep(c byte) {
	d.recordsep = c
}

func (d *textfile) Text() string {
	return string(d.Data())
}

func (d *textfile) Lines() (retval []string, err error) {

	scanner := bufio.NewScanner(d)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		retval = append(retval, scanner.Text())
	}

	return retval, scanner.Err()
}
