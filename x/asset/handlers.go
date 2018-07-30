package asset



import (
	"fmt"
	"reflect"

	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/asset/messages"

)


func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {

		case messages.MsgCreate:
			return handleAssetCreation(ctx, k, msg)
		case messages.MsgRetrieve:
			return handleAssetRetrieval(ctx, k, msg)
		case messages.MsgUpdate:
			return handleAssetUpdate(ctx, k, msg)
		case messages.MsgDelete:
			return handleAssetDelete(ctx, k, msg)

		default:
			errMsg := fmt.Sprintf("Unrecognized trace Msg type: %v", reflect.TypeOf(msg).Name())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}


func handleAssetCreation(ctx sdk.Context, k Keeper, msg messages.MsgCreate) sdk.Result {

	asset, err := k.CreateAsset(ctx, msg)
	if err != nil {
		return sdk.ErrInternal(err.Error()).Result()
	}

	return sdk.Result{
		Log: fmt.Sprintf("%s", asset),
		Tags: msg.Tags(),
	}
}


func handleAssetRetrieval(ctx sdk.Context, k Keeper, msg messages.MsgRetrieve) sdk.Result {

	asset, err := k.RetrieveAsset(ctx, msg)
	if err != nil {
		return sdk.ErrInternal(err.Error()).Result()
	}

	return sdk.Result{
		Log: fmt.Sprintf("%s", asset),
		Tags: msg.Tags(),
	}
}

func handleAssetUpdate(ctx sdk.Context, k Keeper, msg messages.MsgUpdate) sdk.Result {

	asset, err := k.UpdateAsset(ctx, msg)
	if err != nil {
		return sdk.ErrInternal(err.Error()).Result()
	}

	return sdk.Result{
		Log: fmt.Sprintf("%s", asset),
		Tags: msg.Tags(),
	}
}

func handleAssetDelete(ctx sdk.Context, k Keeper, msg messages.MsgDelete) sdk.Result {

	asset, err := k.DeleteAsset(ctx, msg)
	if err != nil {
		return sdk.ErrInternal(err.Error()).Result()
	}

	return sdk.Result{
		Log: fmt.Sprintf("%s", asset),
		Tags: msg.Tags(),
	}
}
