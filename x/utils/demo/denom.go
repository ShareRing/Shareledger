package denom

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"math"
)

var ShrExponent = sdk.NewDec(int64(math.Pow(10, 8)))
