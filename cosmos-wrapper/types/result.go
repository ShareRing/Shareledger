package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// cosmos-sdk v0.27.0 remove FeeDenom, FeeAmount
// This wraps around their old implementation
// The new Result is embedded inside Shareledger's Result
type Result struct {
	sdk.Result
	FeeDenom string
	FeeAmount int64
}

func NewResult(r sdk.Result) Result {
	return Result{
		Result: r,
	}
}

func (r Result) CosmosResult() sdk.Result {
	return r.Result
}