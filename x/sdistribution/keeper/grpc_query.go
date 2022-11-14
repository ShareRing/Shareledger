package keeper

import (
	"github.com/sharering/shareledger/x/sdistribution/types"
)

var _ types.QueryServer = Keeper{}
