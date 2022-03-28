package keeper

import (
	"github.com/sharering/shareledger/x/id/types"
)

var _ types.QueryServer = Keeper{}
