package gofile

import (
	"io/fs"
	"io/ioutil"
	"path/filepath"
)

type (
	DIR interface {
		Len() int
		Path() string
		List() ([]DataFile, error)
		SetOpts(opts dirOpts)
	}
)

// type datafile =

type dirList struct {
	providedName string
	pathname     string
	count        int
	opts         dirOpts
	list         []fs.FileInfo // []DataFile // fs.FileInfo
}

func (l *dirList) Len() int {
	if l.count == 0 {
		l.count = len(l.list)
	}
	return l.count
}

func (l *dirList) Dirs() (list []FileInfo) {
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

func (l *dirList) List() (fi []FileInfo, err error) {
	if l.Len() == 0 {
		l.list, err = ioutil.ReadDir(l.Path())
		if err != nil {
			return nil, err
		}
		// for _, file := range list {
		// 	fp := &basicFile{
		// 		// ProvidedName: file.Name(),
		// 		// size:         file.Size(),
		// 		FileMode: file.Mode(),
		// 	}
		// 	l.list = append(l.list, fp)
		// }
	}
	return l.list, nil
}

func (l *dirList) SetOpts(opts dirOpts) {
	l.opts = opts
}

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
