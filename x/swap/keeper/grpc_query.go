package keeper

import (
	"github.com/sharering/shareledger/x/swap/types"
)

var _ types.QueryServer = Keeper{}
