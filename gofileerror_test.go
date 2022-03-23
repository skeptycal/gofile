package gofile

import (
	"errors"
	"os"
	"testing"
)

var (
	errFake           = errors.New("fake error")
	fakePathError     = &os.PathError{Op: "os.fs", Path: "fake path error", Err: errFake}
	fakeGoFileError   = NewGoFileError("gofile error test", "gofile.fs", fakePathError)
	targetPathError   = NewGoFileError("target PathError", "os.fs", fakePathError)
	targetGoFileError = NewGoFileError("target GoFileError", "gofile.fs", fakePathError)
)

func Test_GoFileError(t *testing.T) {
	got := fakeGoFileError
	if errors.Is(got, targetGoFileError) {
		t.Errorf("error should be wrapped as a GofileError: %v", got)
	}
}
