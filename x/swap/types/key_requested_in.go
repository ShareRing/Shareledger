package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// RequestedInKeyPrefix is the prefix to retrieve all RequestedIn
	RequestedInKeyPrefix = "RequestedIn/value/"
)

// RequestedInKey returns the store key to retrieve a RequestedIn from the index fields
func RequestedInKey(
	txHash string,
) []byte {
	var key []byte

	txHashBytes := []byte(txHash)
	key = append(key, txHashBytes...)
	key = append(key, []byte("/")...)
	return key
}
