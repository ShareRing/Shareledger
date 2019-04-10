package identity

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	amino "github.com/tendermint/go-amino"
	abci "github.com/tendermint/tendermint/abci/types"
)

const (
	QueryIdentity = "identity"
)

func NewQuerier(k Keeper, cdc *amino.Codec) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryIdentity:
			return queryIdentity(ctx, cdc, req, k)
		default:
			return nil, sdk.ErrUnknownRequest("Unknown identity query endpoint")
		}
	}
}

type QueryIdentityParams struct {
	Address sdk.AccAddress
}

func queryIdentity(
	ctx sdk.Context, cdc *amino.Codec, req abci.RequestQuery, k Keeper,
) (res []byte, err sdk.Error) {
	var params QueryIdentityParams

	errRes := cdc.UnmarshalBinaryLengthPrefixed(req.Data, &params)

	if errRes != nil {
		return []byte{}, sdk.ErrUnknownAddress(fmt.Sprintf("Malform address: %s", errRes.Error()))
	}

	hash, ok := k.Get(ctx, params.Address)
	if !ok {
		return []byte{}, sdk.ErrInternal(fmt.Sprintf("Identity for address doesn't exist"))
	}

	res = []byte(hash)
	return res, nil
}
