package fs

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"./copybenchmarks"
)

const (
	NormalMode        os.FileMode = 0644
	DirMode           os.FileMode = 0755
	MinBufferSize                 = 16
	SmallBufferSize               = 64
	Chunk                         = 512
	DefaultBufferSize             = 1024
	DefaultBufSize                = 4096
	MaxInt                        = int(^uint(0) >> 1)
	MinRead                       = bytes.MinRead
)

var Copy = copybenchmarks.Copy

// Basicfile provides implements BasicFile by providing access
// to information and data from a single local file. It is
// designed to cache file information and contents in memory.
type Basicfile struct {
	ProvidedName  string // file name as provided by various constructor methods
	name          string // Base name of the file
	abs           string // Full absolute path of the file
	bak           string // Full absolute path of the file with an additional '.bak' suffix
	tmp           string // Full absolute path of the file with an additional '~' suffix
	size          int64  // Size of the file
	writecache    bool   // Enable write caching (FSync() must be called explicitly.)
	filetype      string // Type of the file (CSV, JSON, etc.)
	os.FileInfo          // os.FileInfo interface
	*bytes.Buffer        // buffered io.ReadWriter
	f             *os.File
	t             *os.File
}

const defaultWriteCache = true

// NewBasicFile creates a new basic file with the provided name.
//
// The underlying file will be created if it does not exist.
// Create creates or truncates the named file. If the file
// already exists, it is truncated. If the file does not exist,
// it is created with mode 0666 (before umask). If successful,
// methods on the returned File can be used for I/O; the associated
// file descriptor has mode O_RDWR. If there is an error, it will
// be of type *PathError.
func New(filename string) (*Basicfile, error) {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, NormalMode)
	if err != nil {
		return nil, err
	}
	b := Basicfile{
		ProvidedName: filename,
		name:         f.Name(),
		f:            f,
		writecache:   defaultWriteCache,
		filetype:     "basic",
	}
	return &b, nil
}

// Close resets the in memory buffer that is used to cache the file contents.
// The actual file was closed after reading the data into memory.
//
// Reset resets the buffer to be empty, but it retains the underlying
// storage for use by future writes. This will cause problems for long
// running programs that access many files. Use Purge() to release the buffer.
func (d *Basicfile) Close() (err error) {
	// TODO - this will cause problems for long running programs ...
	defer d.Reset()
	if d.f != nil {
		err = FSErr(d.f.Close())
	}
	if d.t != nil {
		cerr := Err(d.t.Close())
		rerr := Err(os.Remove(d.t.Name()))

		if rerr != nil {
			return rerr
		}
		if cerr != nil {
			return cerr
		}
	}
	return err
}

// clear will clear all struct fields. This has the effect of
// removing the cached values and forcing any individual
// function calls to retrieve the fresh values.
//
// Used when the underlying data structure has been changed
// in some material way. (e.g. renaming a file)
func (d *Basicfile) clear() error {
	d.name = ""
	d.abs = ""
	d.bak = ""
	d.tmp = ""
	d.size = 0
	d.FileInfo = nil
	d.f = nil
	d.t = nil
	return d.Purge()
}

// newdata returns a new, empty data structure of the appropriate format.
func (d *Basicfile) newdata() interface{} {
	return bytes.NewBuffer(make([]byte, 0, 0))
}

// Purge clears the in memory buffer and resets the capacity to zero.
func (d *Basicfile) Purge() error {
	d.Reset()
	d.Buffer = d.newdata().(*bytes.Buffer)
	if d.Len() != 0 || d.Cap() != 0 { // TODO: this is a little sketchy...
		return ErrNoAlloc
	}
	return nil
}

// Stat returns the cached file information for the underlying
// BasicFile. It is intended to be used to cache permanent
// information, such as file name, Mode(), IsDir(), etc.
//
// While this works fine with permanent values such as the
// file name, it will not give useful results for data that
// are often changed. (e.g. Last Access, Size) In those cases,
// use os.Stat() directly instead.
func (d *Basicfile) Stat() (FileInfo, error) {
	if d.FileInfo == nil {
		fi, err := os.Stat(d.Abs())
		if err != nil {
			return nil, Err(err)
		}
		d.FileInfo = fi
	}
	return d.FileInfo, nil
}

func (d *Basicfile) String() string {
	return fmt.Sprintf("%s: %s", d.filetype, d.Name())
}

// Data returns the contents of the data source buffer
// in the format expected in common use cases.
//
// For typical files, this is []byte, but it may be
// implemented as a CSV string, Shapefile object,
// serialized JSON, or other format.
func (d *Basicfile) Data() []byte {
	if d.Len() == 0 {
		_, err := d.load()
		if err != nil {
			return nil
		}
	}
	buf := make([]byte, 0, d.buffersize())
	buf = append(buf, d.Bytes()...)

	return buf
}

// IsRegular reports whether the file is a regular file.
// That is, it tests that no mode type bits are set.
func (d *Basicfile) IsRegular() bool {
	return d.Mode().IsRegular()
}

// Abs returns an absolute representation of the file's path. If
// the path is not absolute it will be joined with the current
// working directory to turn it into an absolute path. The absolute
// path name for a given file is not guaranteed to be unique. Abs
// calls Clean on the result.
//
// The path is cached in the basicfile object. If the file is moved to
// another location, a new basicfile object should be created to track it.
//
// If an error occurs, it is logged and the empty string is returned.
func (d *Basicfile) Abs() string {
	if d.abs == "" {
		chk, err := filepath.Abs(d.ProvidedName)
		if err != nil {
			Err(fmt.Errorf("provided filename '%s' not found: %v", d.ProvidedName, err))
			return ""
		}
		d.abs = chk
	}
	return d.abs
}

// Move moves the file to another location and updates the basicfile
// object with the result. If an error occurs, it is logged and
// returned.
//
// If the dst string provided does not contain an absolute path,
// the relative path of the source BasicFile is used. This means
// that the file is actually renamed to dst in the same directory.
//
// If the destination file already exists, an error is returned.
// To replace the destination file with the source file, use Rename().
//
func (d *Basicfile) Move(dst string) error {
	dstabs, err := filepath.Abs(dst)
	if err != nil {
		return err
	}
	if _, err := os.Stat(dstabs); errors.Is(err, os.ErrNotExist) {
		return os.Rename(d.abs, dstabs)
	}
	return os.ErrExist
}

// Rename renames (moves) the underlying file to dst. If dst already
// exists and is not a directory, Rename replaces it. OS-specific
// restrictions may apply when they are in different
// directories. If there is an error, it will be of type *LinkError.
func (d *Basicfile) Rename(dst string) error {
	err := os.Rename(d.Abs(), dst)
	if err != nil {
		return err
	}
	d.clear()
	if err != nil {
		return err
	}
	d.name = dst
	return nil
}

// SetData truncates the buffer and reads p into it.
func (d *Basicfile) SetData(p []byte) (n int, err error) {
	d.Reset()
	if d.Len() < len(p) {
		d.Grow(len(p))
	}
	return d.Write(p)
}

// File returns a file pointer to the underlying file.
func (d *Basicfile) File() (*os.File, error) {
	if d.f != nil {
		_, err := d.load()
		if err != nil {
			return nil, err
		}
	}
	return d.f, nil
}

// ReadFrom reads data from the underlying file until EOF and
// appends it to the buffer, growing the buffer as needed.
// The return value n is the number of bytes read. Any error
// except io.EOF encountered during the read is also returned.
// If the buffer becomes too large, ReadFrom will panic with
// ErrTooLarge.
//
// The buffer contents are not guaranteed to be written to the
// underlying file until FSync() is called.
//
// The buffer is not reset prior to calling ReadFrom. If this is
// desired, use ReadFile instead.
func (d *Basicfile) ReadFrom(r io.Reader) (n int64, err error) {
	return d.Buffer.ReadFrom(r)
}

// ReadFile reads the named file using os.ReadFile and writes
// the data into the buffer.
//
// The buffer is *reset* before calling os.ReadFile. All data in
// the buffer is *truncated* at this time. If the desired outcome
// is to APPEND the data to the current buffer, use ReadFrom() instead.
//
// A successful call returns err == nil, not err == EOF.
// Because ReadFile reads the whole file, it does not treat an
// EOF from Read as an error to be reported.
//
// The buffer contents are not guaranteed to be written to the
// underlying file until FSync() is called. If WriteCache is
// false, FSync() is called immediately.
func (d *Basicfile) ReadFile(filename string) (n int64, err error) {

	d.Reset()

	f, err := os.Open(d.Abs())
	if err != nil {
		return 0, err
	}
	defer f.Close()

	return d.ReadFrom(nil)

	buf, err := os.ReadFile(filename)
	if err != nil {
		return 0, err
	}

	d.SetData(buf)

	n, err = d.Buffer.ReadFrom(f)
	if err != nil {
		return n, err
	}

	if n != d.Size() {
		return n, fmt.Errorf("bad read count: (want %d - got %d)", d.Size(), n)
	}

	return n, nil
}

// buffersize calculates the buffer size for the file.
func (d *Basicfile) buffersize() int64 {
	// TODO - should analyze different buffersize values
	return d.Size() + MinBufferSize
}

// tmpName creates a temporary file on disk with a trailing ~ suffix
// added to the name. The file is removed and the name is returned.
func (d *Basicfile) tmpName() string {
	if d.tmp == "" {
		d.tmp = filepath.Join(d.Name(), "~")

	}
	return d.tmp
}

func (d *Basicfile) tmpFile() (*os.File, error) {
	return os.CreateTemp("", d.Name())
}

func (d *Basicfile) bakName() string {
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

func (d *Basicfile) replace(old, new string) error {
	p := bytes.ReplaceAll(d.Bytes(), []byte(old), []byte(new))

	n, err := d.SetData(p)
	if err != nil {
		return err
	}

	if n != len(p) {
		return bufio.ErrBadReadCount
	}
	return nil
}

func (d *Basicfile) writeBak() error {

	_, err := Copy(d.Name(), d.bakName())

	if err != nil {
		return err
	}
	return nil
}
