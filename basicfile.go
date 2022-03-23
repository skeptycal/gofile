package gofile

import (
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type (
	basicFile struct {
		providedName string      // original user input
		fi           os.FileInfo // cached file information
		modTime      time.Time   // used to validate cache entries
		*os.File                 // underlying file handle
	}
)

type (
	ReadWriteCloser interface {
		Read(p []byte) (n int, err error)
		Write(p []byte) (n int, err error)
		Close() error
	}

	StringWriter interface {
		ReadWriteCloser
		io.StringWriter
	}

	Seeker interface {
		Seek(offset int64, whence int) (int64, error)
	}

	ToFrom interface {
		io.ReaderFrom
		io.WriterTo
	}

	// ReadWriteAt implements io.ReaderAt and io.WriterAt
	// for concurrent, non-overlapping reads and writes.
	ReadWriterAt interface {
		// ReaderAt is the interface that wraps the basic ReadAt method.
		//
		// ReadAt reads len(p) bytes into p starting at offset off in the
		// underlying input source. It returns the number of bytes
		// read (0 <= n <= len(p)) and any error encountered.
		//
		// When ReadAt returns n < len(p), it returns a non-nil error
		// explaining why more bytes were not returned. In this respect,
		// ReadAt is stricter than Read.
		//
		// Even if ReadAt returns n < len(p), it may use all of p as scratch
		// space during the call. If some data is available but not len(p) bytes,
		// ReadAt blocks until either all the data is available or an error occurs.
		// In this respect ReadAt is different from Read.
		//
		// If the n = len(p) bytes returned by ReadAt are at the end of the
		// input source, ReadAt may return either err == EOF or err == nil.
		//
		// If ReadAt is reading from an input source with a seek offset,
		// ReadAt should not affect nor be affected by the underlying
		// seek offset.
		//
		// Clients of ReadAt can execute parallel ReadAt calls on the
		// same input source.
		//
		// Implementations must not retain p.
		ReadAt(p []byte, off int64) (n int, err error)

		// WriterAt is the interface that wraps the basic WriteAt method.
		//
		// WriteAt writes len(p) bytes from p to the underlying data stream
		// at offset off. It returns the number of bytes written from p (0 <= n <= len(p))
		// and any error encountered that caused the write to stop early.
		// WriteAt must return a non-nil error if it returns n < len(p).
		//
		// If WriteAt is writing to a destination with a seek offset,
		// WriteAt should not affect nor be affected by the underlying
		// seek offset.
		//
		// Clients of WriteAt can execute parallel WriteAt calls on the same
		// destination if the ranges do not overlap.
		//
		// Implementations must not retain p.
		WriteAt(p []byte, off int64) (n int, err error)
	}

	FileOps interface {
		Chmod(mode os.FileMode) error
		Move(path string) error
		Abs() (string, error)
		Split(path string) (dir, file string)
		Base(path string) string
		Dir(path string) string
		Ext(path string) string
	}

	GoFile interface {
		ReadWriteCloser
		ToFrom

		Name() string
		Open(name string) (http.File, error)
		Readdir(count int) ([]os.FileInfo, error)
		Stat() (os.FileInfo, error)
	}
)

func NewFile(providedName string) (BasicFile, error) {
	return nil, ErrNotImplemented
}

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
// using File, FileInfo, and FileModer interfaces
//
// Minimum required to implement fs.File interface:
//  type File interface {
//     Stat() (fs.FileInfo, error)
//     Read([]byte) (int, error)
//     Close() error
//  }
//
// Implements fs.FileInfo interface:
// 	// A FileInfo describes a file and is returned by Stat.
//  type FileInfo interface {
//  	Name() string       // base name of the file
// 	Size() int64        // length in bytes for regular files; sy stem-dependent for others
//  	Mode() FileMode     // file mode bits
//  	ModTime() time.Time // modification time
//  	IsDir() bool        // abbreviation for Mode().IsDir()
//  	Sys() interface{}   // underlying data source (can return nil)
//  }
type BasicFile interface {
	Handle() (*os.File, error)

	// Implements fs.File interface:
	Stat() (fs.FileInfo, error)
	io.ReadCloser

	FileModer
	FileInfo

	// Additional basic methods:
	Abs() string // absolute path of the file
	// IsRegular() bool // is a regular file?
	// String() string

}

// type BasicFile interface {
// 	File
// 	FileInfo
// }

func (f *basicFile) Handle() (*os.File, error) {
	if f.File != nil {
		return f.File, nil
	}
	ff, err := os.Open(f.providedName)
	if err != nil {
		return nil, err
	}
	f.File = ff
	return f.File, nil
}

// Name - returns the base name of the file
func (f *basicFile) Name() string {
	return filepath.Base(f.Abs())
}

func (f *basicFile) Abs() string {
	s, err := filepath.Abs(f.providedName)
	if err != nil {
		return ""
	}
	return s
}

func (f *basicFile) FileInfo() FileInfo {
	if f.fi == nil {
		fi, err := os.Stat(f.Name())
		// log error but do not return
		if err != nil {
			_ = Err(err)
			return nil
		}
		f.fi = fi
	}
	return f.fi
}

// Size - returns the length in bytes for regular files; system-dependent for others
func (f *basicFile) Size() int64 {
	return f.FileInfo().Size()
}

// Mode - returns the file mode bits
func (f *basicFile) Mode() FileMode {
	return f.FileInfo().Mode()
}

// ModTime - returns the modification time
func (f *basicFile) ModTime() time.Time {
	return f.FileInfo().ModTime()
}

// IsDir - returns true if the file is a directory
func (f *basicFile) IsDir() bool {
	return f.FileInfo().IsDir()
}

// Sys - returns the underlying data source (can return nil)
func (f *basicFile) Sys() interface{} {
	return f.FileInfo().Sys()
}

func (f *basicFile) IsRegular() bool {
	return f.FileInfo().Mode().IsRegular()
}

func (f *basicFile) Perm() FileMode {
	return f.FileInfo().Mode().Perm()
}

func (f *basicFile) Type() FileMode {
	return f.FileInfo().Mode().Type()
}

func (f *basicFile) String() string {
	return fmt.Sprintf("%8s %15s", f.Mode(), f.Name())
}
