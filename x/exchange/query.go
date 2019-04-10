package exchange

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	amino "github.com/tendermint/go-amino"
	abci "github.com/tendermint/tendermint/abci/types"
)

const (
	QueryExchange = "exchange"
)

func NewQuerier(k Keeper, cdc *amino.Codec) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryExchange:
			return queryExchange(ctx, cdc, req, k)
		default:
			return nil, sdk.ErrUnknownRequest("Unkown exchange query endpoint")
		}
	}
}

type QueryExchangeParams struct {
	FromDenom string
	ToDenom   string
}

func queryExchange(
	ctx sdk.Context, cdc *amino.Codec, req abci.RequestQuery, k Keeper,
) (res []byte, err sdk.Error) {
	var params QueryExchangeParams

	errRes := cdc.UnmarshalBinaryLengthPrefixed(req.Data, &params)
	if errRes != nil {
		return []byte{}, sdk.ErrTxDecode(fmt.Sprintf("Unknown Exchange query: %s", errRes.Error()))
	}

	exr, errR := k.RetrieveExchangeRate(ctx, params.FromDenom, params.ToDenom)
	if errR != nil {
		return []byte{}, sdk.ErrInternal(err.Error())
	}

	res = []byte(exr.String())
	return res, nil	
}
