package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	abci "github.com/tendermint/tendermint/abci/types"
)

const (
	QueryID     = "id"
	QuerySigner = "signer"
)

func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err error) {
		switch path[0] {
		case QueryID:
			return queryId(ctx, path[1:], req, keeper)
		case QuerySigner:
			return querySigner(ctx, path[1:], req, keeper)
		default:
			return nil, sdkerrors.ErrInvalidRequest
		}
	}
}

func queryId(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, error) {
	Id := keeper.GetId(ctx, path[0])
	res, err := codec.MarshalJSONIndent(keeper.cdc, Id)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return res, nil
}

func querySigner(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, error) {
	IdSigner := keeper.GetIdSigner(ctx, path[0])
	res, err := codec.MarshalJSONIndent(keeper.cdc, IdSigner)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return res, nil
}
