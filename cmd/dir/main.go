package dirtest

import (
	"fmt"
	"os"
	"time"

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

	// green := ansi.NewColor(2, 0, 1)
	// blue := ansi.NewColor(33, 0, 1)

	// var color = ansi.NewColor(ansi.White, ansi.Black, ansi.Bold)

	log.Info("log started...")

	var testpath string

	if len(os.Args) > 2 {
		testpath = os.Args[1]
	} else {
		testpath = gofile.PWD()
	}

	// d, err := gofile.NewDIR(testpath)
	d, err := gofile.NewFileWithErr(testpath)

	if err != nil {
		log.Fatal(err)
	}

	list, err := d.List()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(color)

	fmt.Printf("directory of %s\n", d.Path())
	for _, f := range list {
		fi, err := f.Stat()
		if err != nil {
			log.Fatal(err)
		}
		if fi.IsDir() {
			color = blue
		} else {
			color = green
		}
		// fmt.Printf("%s%s\n", color, fi.Name())
		// fmt.Printf("%s\n", f.Base())

		fmt.Printf("%s%v %7d %v %s\n", color, fi.Mode(), fi.Size(), fi.ModTime().Format(time.Stamp), fi.Name())
		// fmt.Printf("%s\n", f.Base())

	}

	fmt.Println("")

	// file := list[0].FileInfo()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("first file: %v\n", file.Mode())
}
