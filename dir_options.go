package gofile

// dirOpts contains the options for directory listings.
type dirOpts struct {
	dirsfirst     bool   `default:"true"`
	all           bool   `default:"true"`
	almostAll     bool   `default:"true"`
	author        bool   `default:"false"`
	escape        bool   `default:"false"`
	blockSize     string `default:"K"`
	ignoreBackups bool   `default:"false"`
	dirOnly       bool   `default:"false"`
	color         bool   `default:"true"`
	one           bool   `default:"false"`
	columns       int    `default:"0"`
	classify      bool   `default:"true"`
	owner         bool   `default:"true"`
	group         bool   `default:"true"`
	sort          int    `default:"Alpha"`
	revSort       bool   `default:"false"`
	size          string `default:"K"`
	human         bool   `default:"true"`
	si            bool   `default:"true"`
	inode         bool   `default:"true"`
	dereference   bool   `default:"true"`
	numeric       bool   `default:"false"`
	slash         byte   `default:"'/'"`
	quote         bool   `default:"false"`
	recursive     bool   `default:"false"`
	timeStyle     string `default:"time.Stamp"`
}

var defaultOptions = dirOpts{}
