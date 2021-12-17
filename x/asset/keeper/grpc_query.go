package keeper

import (
	"github.com/ShareRing/Shareledger/x/asset/types"
)

var _ types.QueryServer = Keeper{}
