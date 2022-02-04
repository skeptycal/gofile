package gofile

import (
	"errors"
	"testing"
)

var (
	errFakeError = errors.New("fake error")

	errFakePathError = &PathError{
		Op:   "fake path error",
		Path: "os.fs",
		Err:  errFakeError,
	}

	fakeGoFileError = NewGoFileError("gofile error test", "gofile.fs", errFakePathError)

	targetPathError = &GoFileError{"target PathError", "os.fs", errFakeError}

	targetGoFileError = &GoFileError{"target GoFileError", "gofile.fs", errFakeError}
)

func Test_GoFileError(t *testing.T) {
	got := fakeGoFileError
	if errors.Is(got, targetGoFileError) {
		t.Errorf("error should be wrapped as a GofileError: %v", got)
	}
}
