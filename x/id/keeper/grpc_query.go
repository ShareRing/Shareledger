package keeper

import (
	"github.com/ShareRing/Shareledger/x/id/types"
)

var _ types.QueryServer = Keeper{}
