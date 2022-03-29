package gofile

import (
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

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
	providedName string
	pathname     string
	count        int
	opts         dirOptions
	list         []BasicFile // []DataFile // fs.FileInfo
}

func (l *dirList) Len() int {
	if l.count == 0 {
		l.count = len(l.list)
	}
	return l.count
}

func (l *dirList) Dirs() (list []fs.FileInfo) {
	for _, f := range l.list {
		if f.IsDir() {
			list = append(list, f)
		}
	}
	return
}

func (l *dirList) Path() string {
	if l.pathname == "" {
		if !IsDir(l.providedName) {
			log.Errorf("directory %s not found", l.providedName)
			return ""
		}
		dir, err := filepath.Abs(l.providedName)
		if err != nil {
			log.Error(err)
			return ""
		}
		l.pathname = dir
	}
	return l.pathname
}

func fi2bf(fi fs.FileInfo) (BasicFile, error) {
	f, err := os.Open(fi.Name())

	if err != nil {
		return nil, err
	}

	return &basicFile{
		File:    f,
		fi:      fi,
		modTime: time.Now(),
		lock:    false,
	}, nil
}

// Returns the list of files in the directory.
//
// If an error is encountered, that file will be
// skipped and processing will continue.
func (l *dirList) List() (fi []BasicFile, err error) {
	if l.Len() == 0 {
		list, err := ioutil.ReadDir(l.Path())
		if err != nil {
			return nil, err
		}
		for _, file := range list {
			bf, err := fi2bf(file)
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
