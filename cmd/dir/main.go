package main

import (
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
	log "github.com/sirupsen/logrus"
	"github.com/skeptycal/basicfile"
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

	green := color.New(color.FgHiGreen)
	blue := color.New(color.FgHiBlue)
	color := color.New(color.Bold, color.FgHiWhite, color.BgBlack)

	log.Info("log started...")

	var testpath string

	if len(os.Args) > 2 {
		testpath = os.Args[1]
	} else {
		var err error
		testpath, err = os.Getwd()
		if err != nil {
			testpath = "."
		}
	}

	log.Info("testpath: ", testpath)

	// d, err := gofile.NewDIR(testpath)
	d, err := basicfile.NewBasicFile(testpath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(color)

	root := d.FileOps().Abs()

	log.Info("d.Name(): ", root)

	rdList, err := os.ReadDir(root)
	if err != nil {
		log.Fatalf("os.ReadDir error: %v", err)
	}

	fmt.Printf("directory of %s\n", d.Name())
	for _, f := range rdList {
		fi, err := f.Info()
		if err != nil {
			log.Error(err)
			continue
		}
		if fi.IsDir() {
			color = blue
		} else {
			color = green
		}
		// fmt.Printf("%s%s\n", color, fi.Name())
		// fmt.Printf("%s\n", f.Base())

		color.Printf("%v %7d %v %s\n", fi.Mode(), fi.Size(), fi.ModTime().Format(time.Stamp), fi.Name())
		// fmt.Printf("%s\n", f.Base())

	}

	fmt.Println("")

	// file := list[0].FileInfo()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("first file: %v\n", file.Mode())
}
