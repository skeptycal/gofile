package gofile

import (
	"math"
)

const (
	defaultBufSize         = 4096
	chunk          float64 = 512.0
)

// chunkMultiple returns a multiple of chunk size closest to but greater than size.
func chunkMultiple(size int) int {
	return int(math.Ceil(float64(size)/chunk) * chunk)
}

// InitialCapacity returns the multiple of 'chunk' one more than needed to
// accomodate the given capacity.
func InitialCapacity(capacity int) int {
	if capacity <= defaultBufSize {
		return defaultBufSize
	}
	return chunkMultiple(capacity)
}
