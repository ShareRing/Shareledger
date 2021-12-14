package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// AccStateKeyPrefix is the prefix to retrieve all AccState
	AccStateKeyPrefix = "AccState/value/"
)

// AccStateKey returns the store key to retrieve a AccState from the index fields
func AccStateKey(
	key string,
) []byte {
	var key []byte

	keyBytes := []byte(key)
	key = append(key, keyBytes...)
	key = append(key, []byte("/")...)

	return key
}
