package bank

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	amino "github.com/tendermint/go-amino"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/sharering/shareledger/x/auth"
)

// query endpoints supported by auth querier
const (
	QueryBalance = "balance"
)

func NewQuerier(am auth.AccountMapper, cdc *amino.Codec) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryBalance:
			return queryBalance(ctx, cdc, req, am)
		default:
			return nil, sdk.ErrUnknownRequest("unknown auth query endpoint")
		}
	}
}

type QueryBalanceParams struct {
	Address sdk.AccAddress
}

func queryBalance(
	ctx sdk.Context, cdc *amino.Codec, req abci.RequestQuery, am auth.AccountMapper,
) (res []byte, err sdk.Error) {
	var params QueryBalanceParams

	errRes := cdc.UnmarshalBinaryLengthPrefixed(req.Data, &params)
	if errRes != nil {
		return []byte{}, sdk.ErrUnknownAddress(fmt.Sprintf("Malform address: %s", errRes.Error()))
	}

	account := am.GetAccount(ctx, params.Address)
	if account == nil {
		account = auth.NewSHRAccountWithAddress(params.Address)
	}

	res = []byte(account.GetCoins().String())
	// if err1 != nil {
	// 	return []byte{}, sdk.ErrInternal(fmt.Sprintf("couldnot marshal result to JSON: %s", err1.Error()))
	// }

	return res, nil
}
