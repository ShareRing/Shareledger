package keeper

import (
	"github.com/sharering/shareledger/x/electoral/types"
)

var _ types.QueryServer = Keeper{}
