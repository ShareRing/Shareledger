package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// LevelFeeKeyPrefix is the prefix to retrieve all LevelFee
	LevelFeeKeyPrefix = "LevelFee/value/"
)

// LevelFeeKey returns the store key to retrieve a LevelFee from the index fields
func LevelFeeKey(
	level string,
) []byte {
	var key []byte

	levelBytes := []byte(level)
	key = append(key, levelBytes...)
	key = append(key, []byte("/")...)

	return key
}
