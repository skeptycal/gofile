package gofile

import (
	"bytes"
	"fmt"
	"os"
)

const (
	NormalMode        os.FileMode = 0644
	DirMode           os.FileMode = 0755
	MinBufferSize                 = 16
	SmallBufferSize               = 64
	Chunk                         = 512
	DefaultBufferSize             = 1024
	DefaultBufSize                = 4096
	MaxInt                        = int(^uint(0) >> 1)
	MinRead                       = bytes.MinRead
)

const (
	PathSep = os.PathSeparator
	ListSep = os.PathListSeparator
	NewLine = '\n'
)

type TimeZone int

// American Time Zones
const (
	EST TimeZone = -(5 + iota)
	CST
	MST
	PST
)

func (tz TimeZone) String() string {
	return fmt.Sprintf("GMT%+dh", tz)
}

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
