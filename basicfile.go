package gofile

import (
	"os"
)

// A BasicFile provides access to a single file as an in
// memory buffer.
//
// The BasicFile interface is the minimum implementation
// required of the file and may be extended to specific file
// types. (e.g. CSV, JSON, Esri Shapefile, config files, etc.)
//
// It may also be implemented as an abstract "file" interface
// that provides access to a single file that is too large to
// fit in memory at once.
//
// An implementation for large files should include a way to
// cache one section at a time, perhaps using a maxAlloc
// value or a mutex of file sections.
//
// Caching write requests will likely be the bottleneck and
// collecting multiple write requests and then writing the results
// of the most recent or most active areas of the file may be
// effective. However, performance profiling and some research
// into whether a database is more efficient is warranted.
//
// It could also be implemented as a way to access a database,
// API, buffer, or other storage.
//
// A file may implement additional interfaces, such as
// ReadDirFile, ReaderAt, or Seeker, to provide additional
// or optimized functionality.
//
// Reference: standard library fs.go
type BasicFile interface {
	File
	FileInfo
}

type basicFile struct {
	os.File
	FileMode
}

// // Name - returns the base name of the file
// func (f *basicFile) Name() string {
// 	return f.Name()
// }

// // Size - returns the length in bytes for regular files; system-dependent for others
// func (f *basicFile) Size() int64 {
// 	return f.Size()
// }

// // Mode - returns the file mode bits
// func (f *basicFile) Mode() FileMode {
// 	return f.Mode()
// }

// // ModTime - returns the modification time
// func (f *basicFile) ModTime() time.Time {
// 	return f.ModTime()
// }

// // IsDir - returns true if the file is a directory
// func (f *basicFile) IsDir() bool {
// 	return f.IsDir()
// }

// // Sys - returns the underlying data source (can return nil)
// func (f *basicFile) Sys() interface{} {
// 	return f.Sys()
// }
