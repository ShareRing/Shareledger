package identity

import (
	"bitbucket.org/shareringvietnam/shareledger-fix/x/identity/keeper"
	"bitbucket.org/shareringvietnam/shareledger-fix/x/identity/types"
)

const (
	ModuleName             = types.ModuleName
	RouterKey              = types.RouterKey
	StoreKey               = types.StoreKey
	QuerierRoute           = types.QuerierRoute
	TypeCreateIdMsg        = types.TypeCreateIdMsg
	TypeUpdateIdMsg        = types.TypeUpdateIdMsg
	TypeDeleteIdMsg        = types.TypeDeleteIdMsg
	TypeEnrollIdSignersMsg = types.TypeEnrollIDSignersMsg
	TypeRevokeIdSignersMsg = types.TypeRevokeIDSignersMsg
	IdSignerPrefix         = keeper.IdSignerPrefix
	IdPrefix               = keeper.IdPrefix
)

var (
	NewKeeper     = keeper.NewKeeper
	NewQuerier    = keeper.NewQuerier
	ModuleCdc     = types.ModuleCdc
	RegisterCodec = types.RegisterCodec
)

type (
	Keeper             = keeper.Keeper
	MsgEnrollIDSigners = types.MsgEnrollIDSigners
	MsgRevokeIDSigners = types.MsgRevokeIDSigners
	MsgCreateId        = types.MsgCreateId
	MsgUpdateId        = types.MsgUpdateId
	MsgDeleteId        = types.MsgDeleteId
	IdSigner           = types.IdSigner
)
