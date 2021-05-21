package gofile

import (
	"io/fs"
	"os"
)

// SameFile reports whether fi1 and fi2 describe the same file.
// For example, on Unix this means that the device and inode fields
// of the two underlying structures are identical; on other systems
// the decision may be based on the path names.
// SameFile only applies to results returned by this package's Stat.
// It returns false in other cases.
var SameFile = os.SameFile

// A FileMode represents a file's mode and permission bits.
// The bits have the same definition on all systems, so that
// information about files can be moved from one system
// to another portably. Not all bits apply to all systems.
// The only required bit is ModeDir for directories.
type FileMode = os.FileMode

// A FileInfo describes a file and is returned by Stat and Lstat.
type FileInfo = fs.FileInfo

// A FileInfo describes a file and is returned by Stat and Lstat.
//
// type FileInfo interface {
// 	Name() string       // base name of the file
// 	Size() int64        // length in bytes for regular files; system-dependent for others
// 	Mode() FileMode     // file mode bits
// 	ModTime() time.Time // modification time
// 	IsDir() bool        // abbreviation for Mode().IsDir()
// 	Sys() interface{}   // underlying data source (can return nil)
// }

// fi returns the FileInfo associated with the file and serves
// as a lazy cache of os.FileInfo
func (d *basicfile) fi() FileInfo {
	if d.FileInfo == nil {
		fi, err := os.Stat(d.Abs())
		if err != nil {
			log.Error(err)
			return nil
		}
		d.FileInfo = fi
	}
	return d.FileInfo
}

// Name is the base name of the file
// func (d *basicfile) Name() string {

// 	// if d.FileInfo != nil {
// 	// 	return d.FileInfo.Name()
// 	// }
// 	return filepath.Base(d.Abs())
// }

// Size returns the length in bytes for regular files; system-dependent for others
// func (d *basicfile) Size() int64 {
// 	if d.size == 0 {
// 		d.size = d.FileInfo().Size()
// 	}
// 	return d.size
// }

// Mode returns the file mode bits
// func (d *basicfile) Mode() FileMode {
// 	return d.FileInfo().Mode()
// }

// ModTime returns the modification time
// func (d *basicfile) ModTime() time.Time {
// 	return d.FileInfo().ModTime()
// }

// IsDir is an abbreviation for Mode().IsDir()
// func (d *basicfile) IsDir() bool {
// 	return d.info.Mode().IsDir()
// }

// Sys returns the underlying data source (can return nil)
// func (d *basicfile) Sys() interface{} {
// 	return d.FileInfo().Sys()
// }
