package keeper

import (
	"github.com/sharering/shareledger/x/booking/types"
)

var _ types.QueryServer = Keeper{}
