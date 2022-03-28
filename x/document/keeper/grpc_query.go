package keeper

import (
	"github.com/sharering/shareledger/x/document/types"
)

var _ types.QueryServer = Keeper{}
