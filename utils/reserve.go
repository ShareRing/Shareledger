package utils

import (
	"bytes"
	"encoding/hex"

	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"

	"github.com/sharering/shareledger/constants"
)

// IsValidReserve - check whether an address is a valid reserve
func IsValidReserve(address sdk.Address) bool {
	for _, resStr := range constants.RESERVE_ACCOUNTS {
		decoded, err := hex.DecodeString(resStr)
		if err != nil {
			panic(err)
		}
		if bytes.Equal(address[:], decoded) {
			return true
		}
	}
	return false
}
