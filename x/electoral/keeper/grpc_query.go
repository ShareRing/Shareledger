package keeper

import (
	"github.com/ShareRing/Shareledger/x/electoral/types"
)

var _ types.QueryServer = Keeper{}
