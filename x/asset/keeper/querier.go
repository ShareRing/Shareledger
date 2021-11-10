package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	abci "github.com/tendermint/tendermint/abci/types"
)

const (
	QueryAsset = "asset"
)

func NewQuerier(keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err error) {
		switch path[0] {
		case QueryAsset:
			return queryAsset(ctx, path[1:], req, keeper, legacyQuerierCdc)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Unknown query to asset module")
		}
	}
}

func queryAsset(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	asset := keeper.GetAsset(ctx, path[0])
	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, asset)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return res, nil
}
