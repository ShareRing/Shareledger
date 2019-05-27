package utils

import (
	"encoding/hex"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func ByteToString(inp []byte) string {
	return fmt.Sprintf("%x", inp)
}

func StringToAddress(input string) sdk.AccAddress {
	decoded, err := hex.DecodeString(input)
	if err != nil {
		panic(err)
	}
	return sdk.AccAddress(decoded)
}

func CleanupTDMLog(input string) string {
	return strings.TrimLeft(input, "Msg 0 failed: ")
}
