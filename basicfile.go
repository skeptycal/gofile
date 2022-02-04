package gofile

import (
	"io/fs"
	"os"
	"time"
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
//  type FileModer interface {
//  	String() string
//  	IsDir() bool
//  	IsRegular() bool
//  	Perm() FileMode
//  	Type() FileMode
//  }
//
// A FileInfo describes a file and is returned by Stat.
//
//  type FileInfo interface {
//      Name() string       // base name of the file
//      Size() int64        // length in bytes for regular files; system-dependent for others
//      Mode() FileMode     // file mode bits
//      ModTime() time.Time // modification time
//      IsDir() bool        // abbreviation for Mode().IsDir()
//      Sys() interface{}   // underlying data source (can return nil)
//  }
//
// Reference: standard library fs.go
type BasicFile interface {
	File
	FileModer
	FileInfo
}

type blankFile struct {
	rw   *os.File
	mode FileMode
	fi   FileInfo
}

type basicFile struct {
	rw   *os.File
	mode FileMode
	fi   FileInfo

	// A File provides access to a single file.
	// The File interface is the minimum implementation
	// required of the file.
	// A file may implement additional interfaces, such
	// as ReadDirFile, ReaderAt, or Seeker, to provide
	// additional or optimized functionality.
	//
	//  type File interface {
	//  	Stat() (fs.FileInfo, error)
	//  	Read([]byte) (int, error)
	//  	Close() error
	//  }
	//
	// Reference: standard library fs.go

	// f *os.File
	// mode FileMode
	// fi   FileInfo
}

func (f *basicFile) Read(p []byte) (n int, err error) {
	return f.rw.Read(p)
}

func (f *basicFile) Write(p []byte) (n int, err error) {
	return f.rw.Write(p)
}

// Close closes the File, rendering it unusable for I/O. On files that support SetDeadline, any pending I/O operations will be canceled and return immediately with an error. Close will return an error if it has already been called.
func (f *basicFile) Close() error {
	return f.rw.Close()
}

func (f *basicFile) Stat() (fs.FileInfo, error) {
	if f == nil {
		fi, err := Stat(f.Name())
		if err != nil {
			return nil, err
		}
		f.fi = fi
	}
	return f.fi, nil
}

// Size - returns the length in bytes for regular files; system-dependent for others
func (f *basicFile) Size() int64 {
	return f.fi.Size()
}

// Mode - returns the file mode bits
func (f *basicFile) Mode() FileMode {
	return f.mode
}

// ModTime - returns the modification time
func (f *basicFile) ModTime() time.Time {
	return f.fi.ModTime()
}

// Sys - returns the underlying data source (can return nil)
func (f *basicFile) Sys() interface{} {
	return f.fi.Sys()
}

// Name - returns the base name of the file
func (f *basicFile) Name() string {
	return f.fi.Name()
}

///********************************** FileMode

func (f *basicFile) String() string {
	return f.mode.String()
}

// IsDir reports whether m describes a directory. That is, it tests for the ModeDir bit being set in m.
func (f *basicFile) IsDir() bool {
	return f.mode.IsDir()
}

// IsRegular reports whether m describes a regular file. That is, it tests that no mode type bits are set.
func (f *basicFile) IsRegular() bool {
	return f.mode.IsRegular()
}

// Perm returns the Unix permission bits in m (m & ModePerm).
func (f *basicFile) Perm() FileMode {
	return f.mode.Perm()
}

// Type returns type bits in m (m & ModeType).
func (f *basicFile) Type() FileMode {
	return f.mode.Type()
}
