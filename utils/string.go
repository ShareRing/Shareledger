package utils

import (
	"encoding/hex"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func ByteToString(inp []byte) string {
	return  fmt.Sprintf("%x", inp)
}

func StringToAddress(input string) sdk.AccAddress {
	decoded, err := hex.DecodeString(input)
	if err != nil {
		panic(err)
	}
	var address sdk.AccAddress
	copy(address[:], decoded)
	return address
}