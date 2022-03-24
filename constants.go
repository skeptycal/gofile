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
	PathSep      = os.PathSeparator
	ListSep      = os.PathListSeparator
	NL      byte = '\n'
	TAB     byte = '\t'
	NUL     byte = 0
)

type TimeZone int

// American Time Zones
const (
	EST TimeZone = -(5 + iota)
	CST
	MST
	PST
)

var tzNames = map[TimeZone]string{
	1: "EST",
	2: "CST",
	3: "MST",
	4: "PST",
}

func (tz TimeZone) String() string {
	if s, ok := tzNames[tz]; ok {
		return s
	}
	return fmt.Sprintf("GMT%+dh", tz)
}
