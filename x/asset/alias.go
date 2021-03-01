package asset

import (
	"bitbucket.org/shareringvietnam/shareledger-fix/x/asset/keeper"
	"bitbucket.org/shareringvietnam/shareledger-fix/x/asset/types"
)

const (
	ModuleName         = types.ModuleName
	RouterKey          = types.RouterKey
	StoreKey           = types.StoreKey
	QuerierRoute       = types.QuerierRoute
	TypeAssetDeleteMsg = types.TypeAssetDeleteMsg
	TypeAssetCreateMsg = types.TypeAssetCreateMsg
	TypeAssetUpdateMsg = types.TypeAssetUpdateMsg
)

var (
	NewKeeper         = keeper.NewKeeper
	NewQuerier        = keeper.NewQuerier
	NewCreateAssetMsg = types.NewMsgCreate
	NewUpdateAssetMsg = types.NewMsgUpdate
	NewDeleteAssetMsg = types.NewMsgDelete
	NewAsset          = types.NewAsset
	ModuleCdc         = types.ModuleCdc
	RegisterCodec     = types.RegisterCodec
)

type (
	Keeper          = keeper.Keeper
	MsgCreate       = types.MsgCreate
	MsgUpdate       = types.MsgUpdate
	MsgDelete       = types.MsgDelete
	QueryResResolve = types.QueryResResolve
	QueryResNames   = types.QueryResNames
	Asset           = types.Asset
)
