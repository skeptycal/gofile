package gofile

import (
	"errors"
	"os"
	"testing"
)

var (
	fakeError         = errors.New("fake error")
	fakePathError     = &os.PathError{"fake path error", "os.fs", fakeError}
	fakeGoFileError   = NewGoFileError("gofile error test", "gofile.fs", fakePathError)
	targetPathError   = &GoFileError{"target PathError", "os.fs", fakeError}
	targetGoFileError = &GoFileError{"target GoFileError", "gofile.fs", fakeError}
)

func Test_GoFileError(t *testing.T) {
	got := fakeGoFileError
	if errors.Is(got, targetGoFileError) {
		t.Errorf("error should be wrapped as a GofileError: %w", got)
	}
}
