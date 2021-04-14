// Package gofile provides support for file operations.
package gofile

import (
	"os"

	log "github.com/sirupsen/logrus"
)

// var Err func(e error) error = zsh.Err

// PWD returns a rooted path name corresponding to the
// current directory. If the current directory can be
// reached via multiple paths (due to symbolic links),
// Getwd may return any one of them.
func PWD() string {
	// func Getwd() (dir string, err error) {
	dir, err := os.Getwd()
	if err != nil {
		log.Error("PWD could not be determined, using default.")
		return ""
	}
	return dir
}
