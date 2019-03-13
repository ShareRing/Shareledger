package asset

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	amino "github.com/tendermint/go-amino"
	abci "github.com/tendermint/tendermint/abci/types"
)

const (
	QueryAsset = "asset"
)

func NewQuerier(k Keeper, cdc *amino.Codec) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryAsset:
			return queryAsset(ctx, cdc, req, k)
		default:
			return nil, sdk.ErrUnknownRequest("Unknown asset query endpoint")
		}
	}
}

type QueryAssetParams struct {
	UUID string
}

func queryAsset(
	ctx sdk.Context, cdc *amino.Codec, req abci.RequestQuery, k Keeper,
) (res []byte, err sdk.Error) {
	var params QueryAssetParams

	errRes := cdc.UnmarshalBinaryLengthPrefixed(req.Data, &params)
	if errRes != nil {
		return []byte{}, sdk.ErrTxDecode(fmt.Sprintf("Unknown Asset query: %s", errRes.Error()))
	}

	asset, errR := k.RetrieveAsset(ctx, params)
	if errR != nil {
		return []byte{}, sdk.ErrInternal(err.Error())
	}

	res = []byte(asset.String())

	return res, nil
}
