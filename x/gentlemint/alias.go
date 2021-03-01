package gentlemint

import (
	"github.com/sharering/shareledger/x/gentlemint/keeper"
	"github.com/sharering/shareledger/x/gentlemint/types"
)

const (
	ModuleName              = types.ModuleName
	RouterKey               = types.RouterKey
	StoreKey                = types.StoreKey
	QuerierRoute            = types.QuerierRoute
	TypeLoadSHRMsg          = types.TypeLoadSHRMsg
	TypeLoadSHRPMsg         = types.TypeLoadSHRPMsg
	TypeBuyCentMsg          = types.TypeBuyCent
	TypeBurnSHRPMsg         = types.TypeBurnSHRPMsg
	TypeBurnSHRMsg          = types.TypeBurnSHRMsg
	TypeEnrollSHRPLoaderMsg = types.TypeEnrollSHRPLoaderMsg
	TypeRevokeSHRPLoaderMsg = types.TypeRevokeSHRPLoaderMsg
	TypeBuySHRMsg           = types.TypeBuySHRMsg
	TypeSetExchangeMsg      = types.TypeMsgSetExchange
	TypesSendSHRP           = types.TypeMsgSendSHRP
	TypeSendSHR             = types.TypeMsgSendSHR

	IdSignerPrefix = types.IdSignerPrefix
	ActiveStatus   = types.Active
	InactiveStatus = types.Inactive

	TypeEnrollIDSignersMsg = types.TypeEnrollIDSignersMsg
	TypeRevokeIDSignersMsg = types.TypeRevokeIDSignersMsg
)

var (
	NewKeeper         = keeper.NewKeeper
	NewQuerier        = keeper.NewQuerier
	NewMsgLoadSHR     = types.NewMsgLoadSHR
	NewMsgBurnSHRP    = types.NewMsgBurnSHRP
	NewMsgBurnSHR     = types.NewMsgBurnSHR
	NewMsgBuySHR      = types.NewMsgBuySHR
	NewMsgSetExchange = types.NewMsgSetExchange
	NewMsgSendSHRP    = types.NewMsgSendSHRP
	NewMsgSendSHR     = types.NewMsgSendSHR
	ModuleCdc         = types.ModuleCdc
	RegisterCodec     = types.RegisterCodec
	ParseCoinStr      = types.ParseCoinStr
)

type (
	Keeper              = keeper.Keeper
	MsgLoadSHR          = types.MsgLoadSHR
	MsgBurnSHRP         = types.MsgBurnSHRP
	MsgBurnSHR          = types.MsgBurnSHR
	MsgLoadSHRP         = types.MsgLoadSHRP
	MsgBuySHR           = types.MsgBuySHR
	MsgSetExchange      = types.MsgSetExchange
	MsgBuyCent          = types.MsgBuyCent
	MsgEnrollSHRPLoader = types.MsgEnrollSHRPLoaders
	MsgRevokeSHRPLoader = types.MsgRevokeSHRPLoaders
	MsgSendSHRP         = types.MsgSendSHRP
	MsgSendSHR          = types.MsgSendSHR

	MsgRevokeIDSigners = types.MsgRevokeIDSigners
	MsgEnrollIDSigners = types.MsgEnrollIDSigners

	MsgEnrollDocIssuers = types.MsgEnrollDocIssuers
	MsgRevokeDocIssuers = types.MsgRevokeDocIssuers

	MsgEnrollAccOperators = types.MsgEnrollAccOperators
	MsgRevokeAccOperators = types.MsgRevokeAccOperators

	AccState = types.AccState
)
