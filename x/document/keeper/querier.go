package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/ShareRing/Shareledger/x/document/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func NewQuerier(k Keeper, legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case types.QueryByProof:
			return queryDocByProof(ctx, req, k, legacyQuerierCdc)
		case types.QueryByHolder:
			return queryAllDocsOfAHolder(ctx, req, k, legacyQuerierCdc)
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unknown query path: %s", path[0])
		}
	}
}

// Get Doc by proof
func queryDocByProof(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryDocByProofParams

	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	// Return empty doc if the doc does not exist
	queryDoc := types.Document{Proof: params.Proof}
	doc := k.GetDocByProof(ctx, queryDoc)

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, doc)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

func queryAllDocsOfAHolder(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryDocByHolderParams

	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	// Return empty doc if the doc does not exist
	docs := make([]types.Document, 0)

	cb := func(doc types.Document) bool {
		docs = append(docs, doc)
		return false
	}

	k.IterateAllDocsOfAHolder(ctx, params.Holder, cb)

	// docsType := types.Docs(docs)

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, docs)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil

}
