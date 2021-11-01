package keeper

import (
	"github.com/ShareRing/Shareledger/x/document/types"
)

var _ types.QueryServer = Keeper{}
