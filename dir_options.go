package gofile

// SortType is a list of constants representing sort
// methods for directory listings.
type SortType int

const (
	Alpha SortType = iota + 1
	Size
	Version
	Extension
	Atime
	Ctime
)

var sortNames = map[SortType]string{
	1: "Alpha",
	2: "Size",
	3: "Version",
	4: "Extension",
	5: "Atime",
	6: "Ctime",
}

func (s SortType) String() string {
	return sortNames[s]
}

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
