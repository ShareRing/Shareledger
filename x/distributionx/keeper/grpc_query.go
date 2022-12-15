package keeper

import (
	"github.com/sharering/shareledger/x/distributionx/types"
)

var _ types.QueryServer = Keeper{}
