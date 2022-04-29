package types

import "encoding/binary"

var _ binary.ByteOrder

const (
    // FormatKeyPrefix is the prefix to retrieve all Format
	FormatKeyPrefix = "Format/value/"
)

// FormatKey returns the store key to retrieve a Format from the index fields
func FormatKey(
network string,
) []byte {
	var key []byte
    
    networkBytes := []byte(network)
    key = append(key, networkBytes...)
    key = append(key, []byte("/")...)
    
	return key
}