package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// RequestedInKeyPrefix is the prefix to retrieve all RequestedIn
	RequestedInKeyPrefix = "RequestedIn/value/"
)

// RequestedInKey returns the store key to retrieve a RequestedIn from the index fields
func RequestedInKey(
	address string,
) []byte {
	var key []byte

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	return key
}
