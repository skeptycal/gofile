package main

import (
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/skeptycal/ansi"
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

	green := ansi.NewColor(2, 0, 1)
	blue := ansi.NewColor(33, 0, 1)

	log.Info("log started...")

	var testpath string

	if len(os.Args) > 2 {
		testpath = os.Args[1]
	} else {
		testpath = gofile.PWD()
	}

	d, err := gofile.NewDIR(testpath)
	if err != nil {
		log.Fatal(err)
	}

	list, err := d.List()
	if err != nil {
		log.Fatal(err)
	}

	// cli := cli.New()

	fmt.Printf("directory of %s\n", d.Path())
	for _, f := range list {
		if f.IsDir() {
			fmt.Printf("%s%s\n", blue, f.Name())
			// fmt.Printf("%s\n", f.Base())
		}
	}

	for _, f := range list {
		if !f.IsDir() {
			fmt.Printf("%s%v %7d %v %s\n", green, f.Mode(), f.Size(), f.ModTime().Format(time.Stamp), f.Name())
			// fmt.Printf("%s\n", f.Base())
		}
	}

	fmt.Println("")

	// file := list[0].FileInfo()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("first file: %v\n", file.Mode())
}
