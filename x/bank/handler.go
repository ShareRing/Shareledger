package bank

import (
	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/bank/handlers"
	"github.com/sharering/shareledger/x/bank/messages"
	"reflect"
)

func NewHandler(key *sdk.KVStoreKey) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case messages.MsgCheck:
			return handlers.HandleMsgCheck(key)(ctx, msg)
		case messages.MsgLoad:
			return handlers.HandleMsgLoad(key)(ctx, msg)
		case messages.MsgSend:
			return handlers.HandleMsgSend(key)(ctx, msg)
		default:
			errMsg := "Unrecognized bank Msg type" + reflect.TypeOf(msg).Name()
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}
