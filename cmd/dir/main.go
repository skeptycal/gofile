package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/skeptycal/gofile"
)

// func Dir(path string) string {
// 	command := fmt.Sprintf("ls -R %s", path)
// 	result := zsh.Sh(command)
// 	return result
// }

// func DirTime(args string) string {
// 	command := fmt.Sprintf("time ls -R . %s", args)
// 	result := zsh.Sh(command)
// 	return result
// }

func main() {

	log.Info("log started...")

	var testpath string

	if len(os.Args) < 2 {
		testpath = os.Args[1]
	} else {
		testpath = gofile.PWD()
	}

	d, err := gofile.NewDir(testpath)
	if err != nil {
		log.Fatal(err)
	}

	d.List()

	fmt.Println("Directory Listing Benchmarks:\n ")
	fmt.Println("")
	fmt.Println(Dir(testpath))
}
