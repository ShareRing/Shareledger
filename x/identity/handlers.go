package identity

import (
	"encoding/json"
	"fmt"
	"reflect"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/constants"
	sdkTypes "github.com/sharering/shareledger/cosmos-wrapper/types"
	"github.com/sharering/shareledger/utils"
	"github.com/sharering/shareledger/x/auth"
)

// NewHandler ...
func NewHandler(k Keeper, am auth.AccountMapper) sdkTypes.Handler {

	return func(ctx sdk.Context, msg sdk.Msg) sdkTypes.Result {

		// check Signer, only valid reserve accounts are priviledged to perform Identity
		signer := auth.GetSigner(ctx)

		if !utils.IsValidReserve(signer.GetAddress()) {
			return sdkTypes.NewResult(sdk.ErrInternal(fmt.Sprintf(constants.RES_RESERVE_ONLY)).Result())
		}

		var ret sdk.Result
		var address sdk.AccAddress

		switch msg := msg.(type) {
		case MsgIDCreate:
			address = msg.Address
			ret = handleIDCreation(ctx, k, msg)
		case MsgIDUpdate:
			address = msg.Address
			ret = handleIDUpdate(ctx, k, msg)
		case MsgIDDelete:
			address = msg.Address
			ret = handleIDDeletion(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized Identity Msg: %v", reflect.TypeOf(msg).Name())
			return sdkTypes.NewResult(sdk.ErrUnknownRequest(errMsg).Result())
		}

		if !ret.IsOK() {
			return sdkTypes.NewResult(ret)
		}

		// Get fee
		fee, denom := utils.GetMsgFee(msg)

		fmt.Printf("ADDRESS: %s\n", address)

		return sdkTypes.Result{
			Result:    ret,
			FeeDenom:  denom,
			FeeAmount: fee,
			Signer:    address,
		}
	}
}

type idStruct struct {
	Address sdk.AccAddress `json:"address"`
	Hash    string         `json:"string"`
}

func newIdStruct(address sdk.AccAddress, hash string) idStruct {
	return idStruct{
		Address: address,
		Hash:    hash,
	}
}

func (id idStruct) String() string {
	b, err := json.Marshal(id)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func handleIDCreation(ctx sdk.Context, k Keeper, msg MsgIDCreate) sdk.Result {
	_, isExisted := k.Get(ctx, msg.Address)

	if isExisted {
		return sdk.ErrInternal(fmt.Sprintf("Identity already existed")).Result()
	}

	k.Store(ctx, msg.Address, msg.Hash)

	return sdk.Result{
		Log:  newIdStruct(msg.Address, msg.Hash).String(),
		Tags: msg.Tags(),
	}
}

func handleIDUpdate(ctx sdk.Context, k Keeper, msg MsgIDUpdate) sdk.Result {
	_, isExisted := k.Get(ctx, msg.Address)

	if !isExisted {
		return sdk.ErrInternal(fmt.Sprintf("Identity does not existed")).Result()
	}

	k.Store(ctx, msg.Address, msg.Hash)

	return sdk.Result{
		Log:  newIdStruct(msg.Address, msg.Hash).String(),
		Tags: msg.Tags(),
	}
}

func handleIDDeletion(ctx sdk.Context, k Keeper, msg MsgIDDelete) sdk.Result {
	hash, isExisted := k.Get(ctx, msg.Address)

	if !isExisted {
		return sdk.ErrInternal(fmt.Sprintf("Identity does not existed")).Result()
	}

	k.Delete(ctx, msg.Address)

	return sdk.Result{
		Log:  newIdStruct(msg.Address, hash).String(),
		Tags: msg.Tags(),
	}
}
