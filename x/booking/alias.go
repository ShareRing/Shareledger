package booking

import (
	"github.com/sharering/shareledger/x/booking/keeper"
	"github.com/sharering/shareledger/x/booking/types"
)

const (
	ModuleName          = types.ModuleName
	RouterKey           = types.RouterKey
	StoreKey            = types.StoreKey
	QuerierRoute        = types.QuerierRoute
	TypeBookMsg         = types.TypeBookMsg
	TypeBookCompleteMsg = types.TypeBookCompleteMsg
)

var (
	NewKeeper          = keeper.NewKeeper
	NewQuerier         = keeper.NewQuerier
	NewBookMsg         = types.NewMsgBook
	NewBookCompleteMsg = types.NewMsgComplete
	NewBooking         = types.NewBooking
	ModuleCdc          = types.ModuleCdc
	RegisterCodec      = types.RegisterCodec
)

type (
	Keeper          = keeper.Keeper
	MsgBook         = types.MsgBook
	MsgBookComplete = types.MsgComplete
	QueryResResolve = types.QueryResResolve
	QueryResNames   = types.QueryResNames
	Booking         = types.Booking
)
