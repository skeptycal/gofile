package gofile

import (
	"io"
	"io/ioutil"
	"os"
)

func Copy(src, dest string) (int64, error) {
	return copy(src, dest)
}

func copy(src, dst string) (written int64, err error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, Err(err)
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, NewGoFileError("source file not a regular file", src, err)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, NewGoFileError("unable to open source file", src, err)
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, NewGoFileError("unable to create destination file", dst, err)
	}
	defer destination.Close()

	nBytes, err := io.Copy(destination, source)
	return nBytes, NewGoFileError("invalid gofile copy result", src+" to "+dst, err)
}

func CopyUtil(src, dst string) (written int64, err error) {
	buf, err := ioutil.ReadFile(src)
	if err != nil {
		return 0, NewGoFileError("unable to read source file into buffer", src, err)
	}

	n := len(buf)

	err = ioutil.WriteFile(dst, buf, NormalMode)
	if err != nil {
		return 0, NewGoFileError("unable to write destination file from buffer", dst, err)
	}
	return int64(n), nil
}

func CopyBuffer(src, dst string, buffersize int) (written int64, err error) {

	fi, err := os.Stat(src)
	if err != nil {
		return 0, NewGoFileError("unable to read source file", src, err)
	}

	if !fi.Mode().IsRegular() {
		return 0, NewGoFileError("source file not a regular file", src, err)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, NewGoFileError("unable to open source file", src, err)
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, NewGoFileError("unable to create destination file", dst, err)
	}
	defer destination.Close()

	// TODO - test buffersize fi.Size / 10 ... fi.Size / 100, etc. with minimum
	if buffersize == 0 {
		buffersize = DefaultBufferSize
	}

	var nn int64 = 0
	buf := make([]byte, buffersize)

	for {
		n, err := source.Read(buf)
		if err != nil && err != io.EOF {
			return 0, NewGoFileError("error reading source file", src, err)
		}
		if n == 0 {
			break
		}
		nn += int64(n)
		if _, err := destination.Write(buf[:n]); err != nil {
			return nn, NewGoFileError("error writing destination file", dst, err)
		}
	}
	return nn, nil
}
