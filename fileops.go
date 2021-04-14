package gofile

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

const (
	defaultBufSize  = 4096
	smallBufferSize = 64
	chunk           = 512
	maxInt          = int(^uint(0) >> 1)
	minRead         = bytes.MinRead
)

// Err logs errors and passes them through unchanged.
func Err(err error) error {
	if err != nil {
		log.Error(err)
	}
	return err
}

// Stat returns the os.FileInfo for file if it exists.
// If the file does not exist, nil is returned.
// Errors are logged if Err is active.
func Stat(file string) os.FileInfo {
	fi, err := os.Stat(file)
	if err != nil {
		Err(err)
		return nil
	}
	return fi
}

// StatCheck returns file information (after symlink evaluation)
// using os.Stat(). If the file does not exist, is not a regular file,
// or if the user lacks adequate permissions, an error is returned.
func StatCheck(filename string) (os.FileInfo, error) {

	// EvalSymlinks also calls Abs and Clean as well as
	// checking for existance of the file.
	filename, err := filepath.EvalSymlinks(filename)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(filename)
	if err != nil {
		return nil, err
	}

	//Check 'others' permission
	m := fi.Mode()
	if m&(1<<2) == 0 {
		return nil, fmt.Errorf("insufficient permissions: %v", filename)
	}

	if fi.IsDir() {
		return nil, fmt.Errorf("the filename %s refers to a directory", filename)
	}

	if !fi.Mode().IsRegular() {
		return nil, fmt.Errorf("the filename %s is not a regular file", filename)
	}

	return fi, err
}

// chunkMultiple returns a multiple of chunk size closest to but greater than size.
func chunkMultiple(size int64) int64 {
	return (size/chunk + 1) * chunk
}

// InitialCapacity returns the multiple of 'chunk' one more than needed to
// accomodate the given capacity.
func InitialCapacity(capacity int64) int {
	if capacity < defaultBufSize {
		return defaultBufSize
	}
	return int(chunkMultiple(capacity))
}

// Mode returns the filemode of file.
func Mode(file string) os.FileMode { return Stat(file).Mode() }

// Create creates or truncates the named file and returns an opened file as io.ReadCloser.
//
// If the file already exists, it is truncated. If the file
// does not exist, it is created with mode 0666 (before umask).
// If successful, methods on the returned File can be used
// for I/O; the associated file descriptor has mode O_RDWR. If
// there is an error, it will be of type *PathError.
//
// If the file cannot be created, an error of type *PathError
// is returned.
//
// Errors are logged if gofile.Err is active.
func Create(filename string) io.ReadWriteCloser {

	// OpenFile is the generalized open call; most users will use Open or Create instead. It opens the named file with specified flag (O_RDONLY etc.). If the file does not exist, and the O_CREATE flag is passed, it is created with mode perm (before umask). If successful, methods on the returned File can be used for I/O. If there is an error, it will be of type *PathError.
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		Err(err)
		return nil
	}

	return f
}

// CreateSafe creates the named file and returns an opened file as io.ReadCloser.
//
// If successful, methods on the returned File can be used
// for I/O; the associated file descriptor has mode O_RDWR.
//
// If the file already exists, nil is returned.
// Errors are logged if Err is active.
//
// If the file already exists, of an error occurs, it returns
// nil and an error is sent to Err. If there is an error, it
// will be of type *PathError.
//
func CreateSafe(filename string) io.ReadWriteCloser {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		Err(fmt.Errorf("file already exists (%s): %v", filename, err))
		return nil
	}
	return f
}
