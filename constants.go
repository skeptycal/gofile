package gofile

import "fmt"

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
