package keeper

import (
	"github.com/sharering/shareledger/x/asset/types"
)

var _ types.QueryServer = Keeper{}
