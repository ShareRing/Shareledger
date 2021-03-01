package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/x/gentlemint/types"

	abci "github.com/tendermint/tendermint/abci/types"
)

const (
	QueryLoader            = "loader"
	QueryExchange          = "exchange"
	QuerySigner            = "id-signer"
	QuerySigners           = "id-signers"
	QueryAccountOprator    = "account-operator"
	QueryAllAccountOprator = "account-operators"
	QueryDocumentIssuer    = "document-issuer"
	QueryAllDocumentIssuer = "document-issuers"
)

func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err error) {
		switch path[0] {
		case QueryLoader:
			return queryLoader(ctx, path[1:], req, keeper)
		case QueryExchange:
			return queryExchange(ctx, path[1:], req, keeper)
		case QuerySigner:
			return querySigner(ctx, path[1:], req, keeper)
		case QuerySigners:
			return queryAllSigner(ctx, path[1:], req, keeper)
		case QueryAccountOprator:
			return queryAccountOperator(ctx, path[1:], req, keeper)
		case QueryAllAccountOprator:
			return queryAllAccountOperator(ctx, req, keeper)
		case QueryDocumentIssuer:
			return queryDocumentIssuer(ctx, path[1:], req, keeper)
		case QueryAllDocumentIssuer:
			return queryAllDocumentIssuer(ctx, req, keeper)
		default:
			return nil, sdkerrors.ErrInvalidRequest
		}
	}
}

func queryLoader(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, error) {
	loader := keeper.GetSHRPLoader(ctx, path[0])
	res, err := codec.MarshalJSONIndent(keeper.cdc, loader)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return res, nil
}

func queryExchange(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, error) {
	rate := keeper.GetExchangeRate(ctx)
	res, err := codec.MarshalJSONIndent(keeper.cdc, rate)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return res, nil
}

func querySigner(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, error) {
	addr, err := sdk.AccAddressFromBech32(path[0])

	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	idSigner := keeper.GetIdSigner(ctx, addr)
	res, err := codec.MarshalJSONIndent(keeper.cdc, idSigner)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return res, nil
}

func queryAllSigner(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, error) {
	accs := []types.AccState{}
	cb := func(acc types.AccState) (stop bool) {
		accs = append(accs, acc)
		return false
	}

	keeper.IterateIdSigners(ctx, cb)

	res, err := codec.MarshalJSONIndent(keeper.cdc, accs)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return res, nil
}

func queryAccountOperator(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, error) {
	addr, err := sdk.AccAddressFromBech32(path[0])

	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	acc := keeper.GetAccOp(ctx, addr)
	res, err := codec.MarshalJSONIndent(keeper.cdc, acc)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return res, nil
}

func queryAllAccountOperator(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, error) {
	accs := []types.AccState{}
	cb := func(acc types.AccState) (stop bool) {
		accs = append(accs, acc)
		return false
	}

	keeper.IterateAccOps(ctx, cb)

	res, err := codec.MarshalJSONIndent(keeper.cdc, accs)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return res, nil
}

func queryDocumentIssuer(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, error) {
	addr, err := sdk.AccAddressFromBech32(path[0])

	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	acc := keeper.GetDocIssuer(ctx, addr)
	res, err := codec.MarshalJSONIndent(keeper.cdc, acc)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return res, nil
}

func queryAllDocumentIssuer(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, error) {
	accs := []types.AccState{}
	cb := func(acc types.AccState) (stop bool) {
		accs = append(accs, acc)
		return false
	}

	keeper.IterateDocIssuers(ctx, cb)

	res, err := codec.MarshalJSONIndent(keeper.cdc, accs)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return res, nil
}
