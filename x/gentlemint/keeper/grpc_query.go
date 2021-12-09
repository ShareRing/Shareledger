package keeper

import (
	"github.com/ShareRing/Shareledger/x/gentlemint/types"
)

var _ types.QueryServer = Keeper{}
