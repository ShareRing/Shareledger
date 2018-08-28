package auth

import (
	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
)

type contextKey int // local to auth module

const (
	contextKeySigner contextKey = iota
)

// WithSigners add the signer to the context
func WithSigners(ctx sdk.Context, account BaseAccount) sdk.Context {
	return ctx.WithValue(contextKeySigner, account)
}

// Get the signers from the context
func GetSigner(ctx sdk.Context) BaseAccount {
	v := ctx.Value(contextKeySigner)
	if v == nil {
		var b BaseAccount
		return b
	}
	return v.(BaseAccount)
}
