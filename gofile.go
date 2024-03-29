// Package gofile provides support for file operations.
package gofile

import (
	"os"

	"github.com/skeptycal/basicfile"
)

var NewFileWithErr = basicfile.NewFileWithErr

// PWD returns a rooted path name corresponding to the
// current directory. If the current directory can be
// reached via multiple paths (due to symbolic links),
// PWD may return any one of them.
//
// If an error occurs, the empty string is returned and
// the error is logged.
func PWD() string {
	dir, err := os.Getwd()
	if err != nil {
		Err(NewGoFileError("current directory could not be determined", "os.Getwd()", err))
		return ""
	}
	return dir
}

// IsDir reports whether a file is a directory.
// That is, it tests for the ModeDir bit being
// set in m.
func IsDir(name string) bool {

	// TODO: use basicfile
	// return NewFileWithErr(name).IsDir()
	fi, err := os.Stat(name)
	if err != nil {
		return false
	}
	return fi.Mode().IsDir()
}

// IsRegular reports whether a file describes
// a regular file. That is, it tests that no
// mode type bits are set.
func IsRegular(name string) bool {

	// TODO: use basicfile
	// return NewFileWithErr(name).IsRegular()

	fi, err := os.Stat(name)
	if err != nil {
		return false
	}
	return fi.Mode().IsRegular()
}
