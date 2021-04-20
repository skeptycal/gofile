package gofile

import "os"

type PathError = os.PathError

// Reference: standard library fs.go

// PathError records an error and the operation and file path that caused it.
// type PathError struct {
// 	Op   string
// 	Path string
// 	Err  error
// }

// func (e *PathError) Error() string { return e.Op + " " + e.Path + ": " + e.Err.Error() }

// func (e *PathError) Unwrap() error { return e.Err }

// Timeout reports whether this error represents a timeout.
// func (e *PathError) Timeout() bool {
// 	t, ok := e.Err.(interface{ Timeout() bool })
// 	return ok && t.Timeout()
// }
