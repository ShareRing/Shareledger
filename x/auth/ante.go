package auth

import (
	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
)

func NewAnteHandler(am AccountMapper) sdk.AnteHandler {
	return func(
		ctx sdk.Context, tx types.SHRTx,
	) (_ sdk.Context, _ sdk.Result, abort bool) {

		authTx, ok := tx.(AuthTx)
		if !ok {
			return ctx, sdk.ErrInternal("tx must be AuthTx").Result(), true
		}

		sig := autTx.GetSignature()
		if sig == nil {
			return ctx,
				sdk.ErrInternal("Null signature").Result(),
				true
		}

		authSig, ok := sig.(AuthSig)
		if ok != nil {
			return ctx,
				sdk.ErrInternal("Sig must be AuthSig").Result(),
				true
		}

		// verify Nonce and Signature
		_, res := verifySignature(ctx, am, authSig, tx.GetSignBytes())
		if res != nil {
			return ctx, res, true
		}

		return ctx, sdk.Result{}, false // abort = false

	}
}

func verifySignature(ctx sdk.Context,
					am AccountMapper,
					sig AuthSig,
					signBytes []byte) (acc BaseAccount,
									   res sdk.Result) {

	addr = sig.GetPubKey().Address()
	acc := am.GetAccount(ctx, addr)
	
	// acc exists
	if acc == nil {
		return nil,
			sdk.ErrUnknownAddress(addr.String()).Result()
	}
	
	// verify nonce
	currentNonce := acc.GetNonce()
	if sig.GetNonce() < currentNonce {
		return nil,
			sdk.ErrInternal(fmt.Sprintf("Invalid nonce. Current nonce is: %d", currentNonce)).Result(),
	}


	// verify signature
	if !sig.Verify(signBytes) {
		return nil,
		sdk.ErrUnauthorized("Signature Verification failed.").Result()
	}

	// Update nonce
	acc.SetAccount(sig.GetNonce() + 1)

	return acc, nil
}
