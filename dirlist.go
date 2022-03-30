package gofile

import (
	"io/fs"
	"os"
	"path/filepath"

	"github.com/skeptycal/basicfile"
)

type (
	BasicFile = basicfile.BasicFile

	DIR interface {
		Len() int
		Path() string
		List() ([]BasicFile, error)
		SetOpts(opts dirOptions)

		// Chdir changes the current working directory to the file, which must be a directory. If there is an error, it will be of type *PathError.
		Chdir() error
	}
)

// type datafile =

type dirList struct {
	providedName string      // original name provided
	name         string      // JIT absolute file name
	count        int         // JIT cached file count
	opts         dirOptions  // options for directory listing
	list         []BasicFile // or []DataFile // fs.FileInfo
}

func (l *dirList) Len() int {
	if l.count == 0 {
		l.count = len(l.list)
	}
	return l.count
}

// Dirs returns a list of all objects
// in the dirList that are directories.
func (l *dirList) Dirs() []fs.FileInfo {
	list := make([]fs.FileInfo, 0, l.Len())

	for _, f := range l.list {
		if f.IsDir() {
			list = append(list, f)
		}
	}
	return list
}

func (l *dirList) Path() string {
	if l.name == "" {
		if !IsDir(l.providedName) {
			log.Errorf("%s is not a directory", l.providedName)
			return ""
		}
		var err error
		l.name, err = filepath.Abs(l.providedName)
		if err != nil {
			log.Error(err)
			return ""
		}
	}
	return l.name
}

func (l *dirList) Abs() string {
	if l.name == "" {
		l.name, _ = filepath.Abs(l.providedName)
	}
	return l.name
}

func (l *dirList) Dir() string  { return filepath.Dir(l.Abs()) }
func (l *dirList) Base() string { return filepath.Base(l.Abs()) }

// Returns the list of files in the directory.
//
// If an error is encountered, that file will be
// skipped and processing will continue.
func (l *dirList) List() (fi []BasicFile, err error) {
	if l.Len() == 0 {

		// path := l.Dir()

		list, err := os.ReadDir(l.Path())
		// list, err := ioutil.ReadDir(l.Path())
		if err != nil {
			return nil, err
		}
		for _, dir := range list {
			dir.Name()
			bf, err := basicfile.NewBasicFile(dir.Name())
			if err != nil {
				continue
			}
			list = append(list, bf)
		}
	}
	return l.list, nil
}

func (l *dirList) SetOpts(opts dirOptions) {
	l.opts = opts
}

//  DIR interface {
// 		Len() int
// 		Path() string
// 		List() ([]DataFile, error)
// 		SetOpts(opts dirOpts)
// 	}

// func NewDIR(name string) (DIR, error) {
// 	name, err := filepath.Abs(name)
// 	if err != nil {
// 		return nil, err
// 	}

// 	fi, err := os.Stat(name)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if !fi.IsDir() {
// 		return nil, fmt.Errorf("%s is not a directory", name)
// 	}

// 	return &dirList{
// 		providedName: name,
// 	}, nil
// }
