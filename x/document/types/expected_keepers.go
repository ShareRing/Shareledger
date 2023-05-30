package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/id/types"
)

type IDKeeper interface {
	GetFullIDByIDString(ctx sdk.Context, id string) (*types.Id, bool)
}
