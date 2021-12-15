package types

import (
	"encoding/binary"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ binary.ByteOrder

const (
	// AccStateKeyPrefix is the prefix to retrieve all AccState
	AccStateKeyPrefix = "AccState/value/"
)

type AccStateKeyType string

const (
	AccStateKeyIdSigner    AccStateKeyType = "idsigner"
	AccStateKeyDocIssuer   AccStateKeyType = "docIssuer"
	AccStateKeyAccOp       AccStateKeyType = "accop"
	AccStateKeyShrpLoaders AccStateKeyType = "shrploader"
	AccStateKeyVoter       AccStateKeyType = "voter"
)

type IndexKeyAccState string

// AccStateKey returns the store key to retrieve a AccState from the index fields
func AccStateKey(
	k IndexKeyAccState,
) []byte {
	var key []byte

	keyBytes := []byte(string(k))
	key = append(key, keyBytes...)
	key = append(key, []byte("/")...)

	return key
}

func GenAccStateIndexKey(addr sdk.AccAddress, key AccStateKeyType) IndexKeyAccState {
	return IndexKeyAccState(fmt.Sprintf("%s%s", key, addr))
}
