package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ActionLevelFeeKeyPrefix is the prefix to retrieve all ActionLevelFee
	ActionLevelFeeKeyPrefix = "ActionLevelFee/value/"
)

// ActionLevelFeeKey returns the store key to retrieve a ActionLevelFee from the index fields
func ActionLevelFeeKey(
	action string,
) []byte {
	var key []byte

	actionBytes := []byte(action)
	key = append(key, actionBytes...)
	key = append(key, []byte("/")...)

	return key
}
