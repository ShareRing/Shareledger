package utils

import (
	"reflect"
	"strings"

	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"

	"github.com/sharering/shareledger/constants"
)

// GetMsgType return type of message in string
// to be used in fee calculation
func GetMsgType(msg sdk.Msg) string {
	msgTypeFull := reflect.TypeOf(msg).String()

	parts := strings.Split(msgTypeFull, ".")
	msgType := parts[len(parts)-1]

	return msgType
}

func GetMsgFee(msg sdk.Msg) (int64, string) {

	msgType := GetMsgType(msg)

	msgLevel := constants.LEVELS[msgType]

	return int64(constants.FEE_LEVELS[msgLevel]), constants.FEE_DENOM
}
