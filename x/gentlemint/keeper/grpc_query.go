package keeper

import (
	"github.com/sharering/shareledger/x/gentlemint/types"
)

var _ types.QueryServer = Keeper{}
