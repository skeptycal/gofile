package gofile

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// A BasicFile provides access to a single file.
// The BasicFile interface is the minimum implementation required of the file.
// A file may implement additional interfaces, such as
// ReadDirFile, ReaderAt, or Seeker, to provide additional or optimized functionality.
//
// Reference: standard library fs.go
type BasicFile interface {
	io.ReadWriteCloser
	FileInfo
	Stat() (FileInfo, error)
	Read([]byte) (int, error)
	Close() error
}

// type File interface {
// FileInfo // use cached FileInfo from Stat() instead
// Abs() string // this is useful to have ...
// String() string // this is useful to have
// }

type basicfile struct {
	providedName string        // file name as provided by various constructor methods
	name         string        // Base name of the file
	abs          string        // Full absolute path of the file
	bak          string        // Full absolute path of the file with an additional '.bak' suffix
	tmp          string        // Full absolute path of the file with an additional '~' suffix
	size         int64         // Size of the file
	FileInfo                   // os.FileInfo interface
	data         *bytes.Buffer // buffered io.ReadWriter
	f            *os.File
	t            *os.File
}

// Close resets the in memory buffer that is used to cache the file contents.
// The actual file was closed after reading the data into memory.
//
// Reset resets the buffer to be empty, but it retains the underlying
// storage for use by future writes. This will cause problems for long
// running programs that access many files. Use Purge() to release the buffer.
func (d *basicfile) Close() error {
	// TODO - this will cause problems for long running programs ...
	d.data.Reset()
	if d.f != nil {
		d.f.Close()
	}
	if d.t != nil {
		d.t.Close()
		err := os.Remove(d.t.Name())
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *basicfile) Purge() error {
	d.data.Reset()
	d.data = bytes.NewBuffer(make([]byte, 0, 0))
	if d.data.Cap() != 0 {
		return bytes.ErrTooLarge
	}
	return nil
}

func (d *basicfile) Stat() (FileInfo, error) {
	if d.FileInfo == nil {
		fi, err := os.Stat(d.Abs())
		if err != nil {
			log.Error(err)
			return nil, Err(err)
		}
		d.FileInfo = fi
	}
	return d.FileInfo, nil
}

func (d *basicfile) String() string {
	return fmt.Sprintf("datafile: %s", d.Name())
}

func (d *basicfile) Data() ([]byte, error) {
	if d.data.Len() == 0 {
		_, err := d.load()
		if err != nil {
			return nil, err
		}
	}
	buf := make([]byte, 0, d.buffersize())
	buf = append(buf, d.data.Bytes()...)

	return buf, nil
}

// IsRegular is a convenie
func (d *basicfile) IsRegular() bool {
	// return d.FileInfo.Mode().IsRegular()
	return d.Mode().IsRegular()
}

func (d *basicfile) Abs() string {
	if d.abs == "" {
		chk, err := filepath.Abs(d.providedName)
		if err != nil {
			log.Errorf("provided filename '%s' not found: %v", d.providedName, err)
			return ""
		}
		d.abs = chk
	}
	return d.abs
}

func (d *basicfile) SetData(p []byte) (n int, err error) {
	d.data.Reset()
	return d.data.Write(p)
}

func (d *basicfile) File() (*os.File, error) {
	if d.f != nil {
		_, err := d.load()
		if err != nil {
			return nil, err
		}
	}
	return d.f, nil
}

// load reads the file from the disk and returns the number
// of bytes read and any error encountered.
func (d *basicfile) load() (n int64, err error) {
	f, err := os.Open(d.Name())
	if err != nil {
		return 0, err
	}
	defer f.Close()

	n, err = d.data.ReadFrom(f)
	if err != nil {
		return n, err
	}

	if n != d.Size() {
		return n, fmt.Errorf("bad read count: (%d != %d)", d.Size(), n)
	}

	return n, nil

}

// buffersize calculates the buffer size for the file.
func (d *basicfile) buffersize() int64 {
	// TODO - should analyze different buffersize values
	return d.Size() + minBufferSize
}

// tmpName creates a temporary file on disk with a trailing ~ suffix
// added to the name. The file is removed and the name is returned.
func (d *basicfile) tmpName() string {
	if d.tmp == "" {
		d.tmp = filepath.Join(d.Name(), "~")

	}
	return d.tmp
}

func (d *basicfile) tmpFile() (*os.File, error) {
	return os.CreateTemp("", d.Name())
}

func (d *basicfile) bakName() string {
	if d.bak == "" {
		d.bak = fmt.Sprintf("%s.bak", d.Name())
		f, err := os.Create(d.bak)
		if err != nil {
			log.Fatalf("provided filename '%s' could not be created: %v", d.bak, err)
		}
		f.Close()
	}
	return d.bak
}

func (d *basicfile) replace(old, new string) error {
	p := bytes.ReplaceAll(d.data.Bytes(), []byte(old), []byte(new))

	n, err := d.SetData(p)
	if err != nil {
		return err
	}

	if n != len(p) {
		return bufio.ErrBadReadCount
	}
	return nil
}

func (d *basicfile) writeBak() error {

	_, err := Copy(d.Name(), d.bakName())

	if err != nil {
		return err
	}
	return nil
}

//******************* Reference: standard library fs.go

// A ReadDirFile is a directory file whose entries can be read with the ReadDir method.
// Every directory file should implement this interface.
// (It is permissible for any file to implement this interface,
// but if so ReadDir should return an error for non-directories.)
type ReadDirFile interface {
	BasicFile

	// ReadDir reads the contents of the directory and returns
	// a slice of up to n DirEntry values in directory order.
	// Subsequent calls on the same file will yield further DirEntry values.
	//
	// If n > 0, ReadDir returns at most n DirEntry structures.
	// In this case, if ReadDir returns an empty slice, it will return
	// a non-nil error explaining why.
	// At the end of a directory, the error is io.EOF.
	//
	// If n <= 0, ReadDir returns all the DirEntry values from the directory
	// in a single slice. In this case, if ReadDir succeeds (reads all the way
	// to the end of the directory), it returns the slice and a nil error.
	// If it encounters an error before the end of the directory,
	// ReadDir returns the DirEntry list read until that point and a non-nil error.
	ReadDir(n int) ([]DirEntry, error)
}

// A DirEntry is an entry read from a directory
// (using the ReadDir function or a ReadDirFile's ReadDir method).
type DirEntry interface {
	// Name returns the name of the file (or subdirectory) described by the entry.
	// This name is only the final element of the path (the base name), not the entire path.
	// For example, Name would return "hello.go" not "/home/gopher/hello.go".
	Name() string

	// IsDir reports whether the entry describes a directory.
	IsDir() bool

	// Type returns the type bits for the entry.
	// The type bits are a subset of the usual FileMode bits, those returned by the FileMode.Type method.
	Type() FileMode

	// Info returns the FileInfo for the file or subdirectory described by the entry.
	// The returned FileInfo may be from the time of the original directory read
	// or from the time of the call to Info. If the file has been removed or renamed
	// since the directory read, Info may return an error satisfying errors.Is(err, ErrNotExist).
	// If the entry denotes a symbolic link, Info reports the information about the link itself,
	// not the link's target.
	Info() (FileInfo, error)
}
