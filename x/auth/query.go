package auth

import (
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	amino "github.com/tendermint/go-amino"
	abci "github.com/tendermint/tendermint/abci/types"
)

// query endpoints supported by auth querier
const (
	QueryNonce = "nonce"
)

func NewQuerier(am AccountMapper, cdc *amino.Codec) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryNonce:
			return queryNonce(ctx, cdc, req, am)
		default:
			return nil, sdk.ErrUnknownRequest("unknown auth query endpoint")
		}
	}
}

type QueryNonceParams struct {
	Address sdk.AccAddress
}

func queryNonce(
	ctx sdk.Context, cdc *amino.Codec, req abci.RequestQuery, am AccountMapper,
) (res []byte, err sdk.Error) {
	var params QueryNonceParams

	errRes := cdc.UnmarshalBinaryLengthPrefixed(req.Data, &params)
	if errRes != nil {
		return []byte{}, sdk.ErrUnknownAddress(fmt.Sprintf("Malform address: %s", errRes.Error()))
	}

	nonce, err := am.GetNonce(ctx, params.Address)
	if err != nil {
		return []byte{}, sdk.ErrInternal(fmt.Sprintf("Unable to retrieve nonce: %s", err.Error()))
	}

	res, err1 := json.Marshal(nonce)
	if err1 != nil {
		return []byte{}, sdk.ErrInternal(fmt.Sprintf("couldnot marshal result to JSON: %s", err1.Error()))
	}

	return res, nil
}
