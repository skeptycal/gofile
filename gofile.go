// Package gofile provides support for file operations.
package gofile

import (
	"errors"
	"os"

	"github.com/skeptycal/gofile/errorlogger"
)

var (
	ErrBadCount   = errors.New("datafile: bad read count")
	ErrNotRegular = errors.New("datafile: not regular file")
	log           = errorlogger.Log
)

// PWD returns a rooted path name corresponding to the
// current directory. If the current directory can be
// reached via multiple paths (due to symbolic links),
// PWD may return any one of them.
//
// If an error occurs, the empty string is returned and
// the error is logged.
func PWD() string {
	// func Getwd() (dir string, err error) {
	dir, err := os.Getwd()
	if err != nil {
		log.Error(NewPathError("current directory could not be determined", "os.Getwd()", err))
		return ""
	}
	return dir
}

func IsDir(name string) bool {
	fi, err := os.Stat(name)
	if err != nil {
		return false
	}
	return fi.IsDir()
}

func IsRegular(name string) bool {
	fi, err := os.Stat(name)
	if err != nil {
		return false
	}
	return fi.Mode().IsRegular()
}
