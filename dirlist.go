package gofile

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

type DIR interface {
	Len() int
	Path() string
	List() ([]fs.FileInfo, error)
}

type dirList struct {
	providedName string
	pathname     string
	count        int
	list         []fs.FileInfo
}

func (l *dirList) Len() int {
	if l.count == 0 {
		l.count = len(l.list)
	}
	return l.count
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

func (l *dirList) List() ([]fs.FileInfo, error) {
	if l.Len() == 0 {
		list, err := ioutil.ReadDir(l.Path())
		if err != nil {
			return nil, err
		}
		l.list = list
	}
	return l.list, nil
}

func NewDIR(name string) (DIR, error) {

	name, err := filepath.Abs(name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(name)
	if err != nil {
		return nil, err
	}

	if !fi.IsDir() {
		return nil, fmt.Errorf("%s is not a directory", name)
	}

	return &dirList{
		providedName: name,
	}, nil
}
