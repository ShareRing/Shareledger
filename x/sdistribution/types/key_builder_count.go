package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// BuilderCountKeyPrefix is the prefix to retrieve all BuilderCount
	BuilderCountKeyPrefix = "BuilderCount/value/"
)

// BuilderCountKey returns the store key to retrieve a BuilderCount from the index fields
func BuilderCountKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
