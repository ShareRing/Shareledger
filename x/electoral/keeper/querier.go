package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	abci "github.com/tendermint/tendermint/abci/types"
)

const QueryVoter = "voter"

func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err error) {
		switch path[0] {
		case QueryVoter:
			return queryVoter(ctx, path[1:], req, keeper)
		default:
			return nil, sdkerrors.ErrInvalidRequest
		}
	}
}

func queryVoter(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, error) {
	Booking := keeper.GetVoter(ctx, path[0])
	res, err := codec.MarshalJSONIndent(keeper.cdc, Booking)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return res, nil
}
