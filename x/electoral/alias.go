package electoral

import (
	"github.com/sharering/shareledger/x/electoral/keeper"
	"github.com/sharering/shareledger/x/electoral/types"
)

const (
	ModuleName          = types.ModuleName
	RouterKey           = types.RouterKey
	StoreKey            = types.StoreKey
	QuerierRoute        = types.QuerierRoute
	TypeEnrollVoter     = types.TypeEnrollVoterMsg
	TypeRevokeVoter     = types.TypeRevokeVoterMsg
	StatusVoterEnrolled = types.StatusVoterEnrolled
)

var (
	NewKeeper         = keeper.NewKeeper
	NewQuerier        = keeper.NewQuerier
	NewMsgEnrollVoter = types.NewMsgEnrollVoter
	NewMsgRevokeVoter = types.NewMsgRevokeVoter
	ModuleCdc         = types.ModuleCdc
	RegisterCodec     = types.RegisterCodec
)

type (
	Keeper         = keeper.Keeper
	MsgEnrollVoter = types.MsgEnrollVoter
	MsgRevokeVoter = types.MsgRevokeVoter
	Voter          = types.Voter
)
