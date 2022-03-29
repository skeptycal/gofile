package gofile

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/skeptycal/basicfile"
)

const (
	smallBufferSize = 64
	maxInt          = int(^uint(0) >> 1)
	minRead         = bytes.MinRead
)

func Exists(filename string) bool {
	_, err := os.Stat(filename)
	return errors.Is(err, os.ErrNotExist)
}

func NotExists(filename string) bool {
	_, err := os.Stat(filename)
	return errors.Is(err, os.ErrNotExist)
}

// Stat returns the os.FileInfo for file if it exists.
//
// It is a convenience wrapper for os.Stat that traps
// and processes errors that may occur using the
// the ErrorLogger package.
//
// If the file does not exist, nil is returned.
// Errors are logged if Err is active.
func Stat(filename string) os.FileInfo {
	fi, err := os.Stat(filename)
	if err != nil {
		Err(basicfile.NewGoFileError("gofile.Stat()", filename, err))
		return nil
	}
	return fi
}

// Mode returns the filemode of file.
func Mode(filename string) os.FileMode {
	fi, err := os.Stat(filename)
	if err != nil {
		Err(NewGoFileError("gofile.Mode()", filename, err))
		return 0
	}
	return fi.Mode()
}

// StatCheck returns file information (after symlink evaluation)
// using os.Stat(). If the file does not exist, is not a regular file,
// or if the user lacks adequate permissions, an error is returned.
// StatCheck returns file information (after symlink evaluation
// and path cleaning) using os.Stat().
//
// If the file does not exist, is not a regular file,
// or if the user lacks adequate permissions, an error is
// returned.
//
// It is a convenience wrapper for os.Stat that traps
// and processes errors that may occur using the
// the ErrorLogger package.
//
// If the file does not exist, nil is returned.
// Errors are logged if Err is active.
func StatCheck(filename string) (os.FileInfo, error) {

	// EvalSymlinks also calls Abs and Clean as well as
	// checking for existance of the file.
	filename, err := filepath.EvalSymlinks(filename)
	if err != nil {
		return nil, Err(NewGoFileError("gofile.StatCheck()#EvalSymlinks", filename, err))

	}

	fi, err := os.Stat(filename)
	if err != nil {
		return nil, Err(NewGoFileError("gofile.StatCheck()#os.Stat", filename, err))
	}

	//Check 'others' permission
	m := fi.Mode()
	if m&(1<<2) == 0 {
		// err = Wrap(err, "")
		return nil, Err(NewGoFileError("gofile.StatCheck()#insufficient_permissions", filename, ErrPermission))
		fmt.Errorf("insufficient permissions: %v", filename)
	}

	if fi.IsDir() {
		return nil, fmt.Errorf("the filename %s refers to a directory", filename)
	}

	if !fi.Mode().IsRegular() {
		return nil, fmt.Errorf("the filename %s is not a regular file", filename)
	}

	return fi, err
}
