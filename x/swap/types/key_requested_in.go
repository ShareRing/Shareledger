package types

import (
	"encoding/binary"
	"fmt"
)

var _ binary.ByteOrder

const (
	// PastTxEventsKeyPrefix is the prefix to retrieve all PastTxEvent
	PastTxEventsKeyPrefix = "PastTxEvents/"
)

// PastTxEventKey returns the store key to retrieve a PastTxEvent from the index fields
// txhash/logindex/ -> value
func PastTxEventKey(
	txHash string,
	logIndex uint64,
) []byte {
	key := []byte{}

	key = append(key, []byte(Seperator)...)
	key = append(key, []byte(fmt.Sprintf("%d", logIndex))...)
	key = append([]byte(txHash), key...)

	return key
}

func PastTxEventByTxHashKey(txHash string) []byte {
	key := []byte{}

	key = append(key, []byte(Seperator)...)
	key = append([]byte(txHash), key...)

	return key
}
