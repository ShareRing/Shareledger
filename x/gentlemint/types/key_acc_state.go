package types

import (
	"encoding/binary"
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ binary.ByteOrder

type AccStateKeyType string

const (
	AccStateKeyIdSigner    AccStateKeyType = "IDS"
	AccStateKeyDocIssuer   AccStateKeyType = "DOCIS"
	AccStateKeyAccOp       AccStateKeyType = "ACCOP"
	AccStateKeyShrpLoaders AccStateKeyType = "shrploader"
)

const (
	// AccStateKeyPrefix is the prefix to retrieve all AccState
	AccStateKeyPrefix = "AccState/value/"
)

// AccStateKey returns the store key to retrieve a AccState from the index fields
func AccStateKey(
	key string,
) []byte {
	var keyb []byte

	keyBytes := []byte(string(key))
	keyb = append(keyb, keyBytes...)
	keyb = append(keyb, []byte("/")...)

	return keyb
}

func GenAccStateIndexKey(addr sdk.AccAddress, key AccStateKeyType) string {
	return fmt.Sprintf("%s%s", key, addr)
}
