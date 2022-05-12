package keeper

import (
	"context"
	denom "github.com/sharering/shareledger/x/utils/demo"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/x/swap/types"
)

func (k msgServer) CreateSchema(goCtx context.Context, msg *types.MsgCreateSchema) (*types.MsgCreateSchemaResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetSchema(
		ctx,
		msg.Network,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	inF, err := denom.NormalizeToBaseCoin(denom.Base, sdk.NewDecCoins(*msg.In), sdk.NewDec(0), false)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid fee")
	}
	outF, err := denom.NormalizeToBaseCoin(denom.Base, sdk.NewDecCoins(*msg.Out), sdk.NewDec(0), false)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid fee")
	}

	var signSchema = types.Schema{
		Creator: msg.Creator,
		Network: msg.Network,
		Schema:  msg.Schema,
		Fee: &types.Fee{
			In:  &inF,
			Out: &outF,
		},
		ContractExponent: msg.GetContractExponent(),
	}

	k.SetSchema(
		ctx,
		signSchema,
	)
	return &types.MsgCreateSchemaResponse{}, nil
}

func (k msgServer) UpdateSchema(goCtx context.Context, msg *types.MsgUpdateSchema) (*types.MsgUpdateSchemaResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetSchema(
		ctx,
		msg.Network,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var schema = types.Schema{
		Creator: msg.Creator,
		Network: msg.Network,
		Schema:  valFound.Schema,
		Fee:     valFound.Fee,
	}
	if msg.Schema != "" {
		schema.Schema = msg.Schema
	}

	if !msg.Out.IsZero() {
		outF, err := denom.NormalizeToBaseCoin(denom.Base, sdk.NewDecCoins(*msg.Out), sdk.NewDec(0), false)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid fee")
		}
		schema.Fee.Out = &outF

	}
	if !msg.In.IsZero() {
		inF, err := denom.NormalizeToBaseCoin(denom.Base, sdk.NewDecCoins(*msg.In), sdk.NewDec(0), false)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid fee")
		}
		schema.Fee.In = &inF
	}

	if msg.GetContractExponent() != 0 {
		schema.ContractExponent = msg.GetContractExponent()
	}

	k.SetSchema(ctx, schema)

	return &types.MsgUpdateSchemaResponse{}, nil
}

func (k msgServer) DeleteSchema(goCtx context.Context, msg *types.MsgDeleteSchema) (*types.MsgDeleteSchemaResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetSchema(
		ctx,
		msg.Network,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveSchema(
		ctx,
		msg.Network,
	)

	return &types.MsgDeleteSchemaResponse{}, nil
}
