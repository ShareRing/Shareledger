package auth

import (
	"fmt"

	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
	//"bitbucket.org/shareringvn/cosmos-sdk/wire"
	//"github.com/sharering/shareledger/types"
)

func NewAnteHandler(am AccountMapper) sdk.AnteHandler {
	return func(
		ctx sdk.Context, tx sdk.Tx,
	) (_ sdk.Context, _ sdk.Result, abort bool) {

		authTx, ok := tx.(AuthTx)
		if !ok {
			return ctx, sdk.ErrInternal("tx must be AuthTx").Result(), true
		}

		sig := authTx.GetSignature()
		if sig == nil {
			return ctx,
				sdk.ErrInternal("Null signature").Result(),
				true
		}

		authSig, ok := sig.(AuthSig)
		if !ok {
			return ctx,
				sdk.ErrInternal("Sig must be AuthSig").Result(),
				true
		}

		// verify Nonce and Signature
		signingAccount, res := verifySignature(ctx, am, authSig, authTx.GetSignBytes())

		if signingAccount == nil {
			return ctx, res, true
		}

		// Save account to context
		ctx = WithSigners(ctx, signingAccount)

		return ctx, sdk.Result{}, false // abort = false

	}
}

func verifySignature(ctx sdk.Context,
	am AccountMapper,
	sig AuthSig,
	signBytes []byte) (acc BaseAccount,
	res sdk.Result) {

	fmt.Println("PubKey used to sign:", sig.GetPubKey())
	addr := sig.GetPubKey().Address()

	acc = am.GetAccount(ctx, addr)
	fmt.Println("PubKey of this account:", acc)

	// acc doesn't exist
	if acc == nil {
		//return nil,
		//sdk.ErrInternal(fmt.Sprintf("%s doesnt exist", addr.String())).Result()
		shrA := NewSHRAccountWithAddress(addr)
		acc = shrA
	}

	if acc.GetPubKey() == nil {
		acc.SetPubKey(sig.GetPubKey())
	}

	// verify nonce
	currentNonce := acc.GetNonce()
	if sig.GetNonce() != currentNonce+1 {
		return nil,
			sdk.ErrInternal(fmt.Sprintf("Invalid nonce. Current nonce is: %d", currentNonce)).Result()
	}

	// verify signature
	if !sig.Verify(signBytes) {
		return nil,
			sdk.ErrUnauthorized("Signature Verification failed.").Result()
	}

	// Update nonce
	acc.SetNonce(sig.GetNonce()) // currentNonce + 1

	// Save new nonce
	am.SetAccount(ctx, acc)

	return acc, sdk.Result{}
}
