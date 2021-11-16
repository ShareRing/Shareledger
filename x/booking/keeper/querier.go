package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	abci "github.com/tendermint/tendermint/abci/types"
)

const QueryBooking = "booking"

func NewQuerier(keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err error) {
		switch path[0] {
		case QueryBooking:
			return queryBooking(ctx, path[1:], req, keeper, legacyQuerierCdc)
		default:
			return nil, sdkerrors.ErrInvalidRequest
		}
	}
}

func queryBooking(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	Booking := keeper.GetBooking(ctx, path[0])
	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, Booking)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return res, nil
}
