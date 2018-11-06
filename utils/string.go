package utils

import (
	"encoding/hex"
	"fmt"

	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
)

func ByteToString(inp []byte) string {
	return  fmt.Sprintf("%x", inp)
}

func StringToAddress(input string) sdk.Address {
	decoded, err := hex.DecodeString(input)
	if err != nil {
		panic(err)
	}
	var address sdk.Address
	copy(address[:], decoded)
	return address
}