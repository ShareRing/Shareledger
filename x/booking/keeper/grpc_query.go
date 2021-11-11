package keeper

import (
	"github.com/ShareRing/Shareledger/x/booking/types"
)

var _ types.QueryServer = Keeper{}
