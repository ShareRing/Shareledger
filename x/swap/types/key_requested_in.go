package types

import (
	"encoding/binary"
	"fmt"
)

var _ binary.ByteOrder

const (
	// RequestedInKeyPrefix is the prefix to retrieve all RequestedIn
	RequestedInKeyPrefix = "RequestedIn/value/"
)

// RequestedInKey returns the store key to retrieve a RequestedIn from the index fields
func RequestedInKey(
	txHashEventInx string,
	logIndex uint64,
) []byte {
	var key []byte

	txHashBytes := []byte(txHashEventInx)
	key = append(key, txHashBytes...)
	key = append(key, []byte("/")...)
	key = append(key, []byte(fmt.Sprintf("%d", logIndex))...)
	key = append(key, []byte("/")...)
	return key
}
