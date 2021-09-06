// Package fs defines basic interfaces to a file system.
// A file system can be provided by the host operating system
// but also by other packages.
package gofile

import (
	"io"
	"io/fs"
	"os"
	"syscall"
	"time"
)

//******************* Reference: standard library fs.go

// SameFile reports whether fi1 and fi2 describe the same file.
// For example, on Unix this means that the device and inode fields
// of the two underlying structures are identical; on other systems
// the decision may be based on the path names.
// SameFile only applies to results returned by this package gofile
// It returns false in other cases.
var SameFile = os.SameFile

type (

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
	BasicFile interface {
		File
		FileInfo
	}

	// type File interface {
	// FileInfo // use cached FileInfo from Stat() instead
	// Abs() string // this is useful to have ...
	// String() string // this is useful to have
	// }

	GoFile interface {
		Name() string
		Read(b []byte) (n int, err error)
		ReadAt(b []byte, off int64) (n int, err error)
		ReadFrom(r io.Reader) (n int64, err error)
		Write(b []byte) (n int, err error)
		WriteAt(b []byte, off int64) (n int, err error)
		Seek(offset int64, whence int) (ret int64, err error)
		WriteString(s string) (n int, err error)
		Chmod(mode FileMode) error
		SetDeadline(t time.Time) error
		SetReadDeadline(t time.Time) error
		SetWriteDeadline(t time.Time) error
		SyscallConn() (syscall.RawConn, error)
	}

	// An FS provides access to a hierarchical file system.
	//
	// The FS interface is the minimum implementation required of the file system.
	// A file system may implement additional interfaces,
	// such as ReadFileFS, to provide additional or optimized functionality.
	//
	//  type FS interface {
	//  	// Open opens the named file.
	//  	//
	//  	// When Open returns an error, it should be of type *PathError
	//  	// with the Op field set to "open", the Path field set to name,
	//  	// and the Err field describing the problem.
	//  	//
	//  	// Open should reject attempts to open names that do not satisfy
	//  	// ValidPath(name), returning a *PathError with Err set to
	//  	// ErrInvalid or ErrNotExist.
	//  	Open(name string) (File, error)
	//  }
	//
	// Reference: standard library fs.go
	FS = fs.FS

	// A File provides access to a single file.
	// The File interface is the minimum implementation required
	// of the file.
	// A file may implement additional interfaces, such as
	// ReadDirFile, ReaderAt, or Seeker, to provide additional
	// or optimized functionality.
	//
	//  type File interface {
	//  	Stat() (fs.FileInfo, error)
	//  	Read([]byte) (int, error)
	//  	Close() error
	//  }
	//
	// Reference: standard library fs.go
	File = fs.File

	// A DirEntry is an entry read from a directory
	// (using the ReadDir function or a ReadDirFile's ReadDir method).
	//
	//  type DirEntry interface {
	//  	// Name returns the name of the file (or subdirectory) described by the entry.
	//  	// This name is only the final element of the path (the base name), not the entire path.
	//  	// For example, Name would return "hello.go" not "/home/gopher/hello.go".
	//  	Name() string
	//
	//   	// IsDir reports whether the entry describes a directory.
	//  	IsDir() bool
	//
	//   	// Type returns the type bits for the entry.
	//  	// The type bits are a subset of the usual FileMode bits, those returned by the FileMode.Type method.
	//  	Type() FileMode
	//
	//   	// Info returns the FileInfo for the file or subdirectory described by the entry.
	//  	// The returned FileInfo may be from the time of the original directory read
	//  	// or from the time of the call to Info. If the file has been removed or renamed
	//  	// since the directory read, Info may return an error satisfying errors.Is(err, ErrNotExist).
	//  	// If the entry denotes a symbolic link, Info reports the information about the link itself,
	//  	// not the link's target.
	//  	Info() (FileInfo, error)
	//  }
	//
	// Reference: standard library fs.go
	DirEntry = fs.DirEntry

	// A ReadDirFile is a directory file whose entries can be read with the ReadDir method.
	// Every directory file should implement this interface.
	// (It is permissible for any file to implement this interface,
	// but if so ReadDir should return an error for non-directories.)
	//   type ReadDirFile interface {
	//   	File
	// 		// ReadDir reads the contents of the directory and returns
	//		// a slice of up to n DirEntry values in directory order.
	//		// Subsequent calls on the same file will yield further DirEntry values.
	//		//
	//		// If n > 0, ReadDir returns at most n DirEntry structures.
	//		// In this case, if ReadDir returns an empty slice, it will return
	//		// a non-nil error explaining why.
	//		// At the end of a directory, the error is io.EOF.
	//		//
	//		// If n <= 0, ReadDir returns all the DirEntry values from the directory
	//		// in a single slice. In this case, if ReadDir succeeds (reads all the way
	//		// to the end of the directory), it returns the slice and a nil error.
	//		// If it encounters an error before the end of the directory,
	//		// ReadDir returns the DirEntry list read until that point and a non-nil error.
	//   	ReadDir(n int) ([]DirEntry, error)
	//   }
	//
	// Reference: standard library fs.go
	ReadDirFile = fs.ReadDirFile

	// A FileInfo describes a file and is returned by Stat.
	//
	//	type FileInfo interface {
	//		Name() string       // base name of the file
	//		Size() int64        // length in bytes for regular files; system-dependent for others
	//		Mode() FileMode     // file mode bits
	//		ModTime() time.Time // modification time
	//		IsDir() bool        // abbreviation for Mode().IsDir()
	//		Sys() interface{}   // underlying data source (can return nil)
	//	}
	//
	// Reference: standard library fs.go
	FileInfo = fs.FileInfo

	// A FileMode represents a file's mode and permission bits.
	// The bits have the same definition on all systems, so that
	// information about files can be moved from one system
	// to another portably. Not all bits apply to all systems.
	// The only required bit is ModeDir for directories.
	//
	//  type FileMode uint32
	//
	// Reference: standard library fs.go
	FileMode = fs.FileMode

	// FileModer implements fs.FileMode functionality
	// ... yes, I hate the name, too
	FileModer interface {
		String() string
		IsDir() bool
		IsRegular() bool
		Perm() FileMode
		Type() FileMode
	}
)
