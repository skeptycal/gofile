package copybenchmarks

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

const (
	normalMode        os.FileMode = 0644
	dirMode           os.FileMode = 0755
	defaultBufferSize int         = 1024
	minBufferSize     int64       = 16
)

func Copy(src, dest string) (int64, error) {
	return copy(src, dest)
}

// NewPathError records an error and the operation and file path that caused it.
//  type PathError struct {
//  	Op   string
//  	Path string
//  	Err  error
//  }
func NewPathError(op, path string, err error) *PathError {
	return &PathError{
		Op:   op,
		Path: path,
		Err:  err,
	}
}

func copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()

	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

func copyutil(src, dst string) error {
	buf, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(dst, buf, normalMode)
}

func copybuffer(src, dst string, buffersize int) error {

	fi, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !fi.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()

	// TODO - test buffersize fi.Size / 10 ... fi.Size / 100, etc. with minimum
	if buffersize == 0 {
		buffersize = defaultBufferSize
	}

	buf := make([]byte, buffersize)

	for {
		n, err := source.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}

		if _, err := destination.Write(buf[:n]); err != nil {
			return err
		}
	}
	return nil
}

func copyBenchmark() {
	// https://opensource.com/article/18/6/copying-files-go
}
