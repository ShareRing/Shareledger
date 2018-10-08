package keeper

// import (
// 	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
// 	types "github.com/sharering/shareledger/types"
// )

// // Perform a delegation, set/update everything necessary within the store.
// func (k Keeper) Delegate(ctx sdk.Context, delAddr sdk.Address, bondAmt types.Coin,
// 	validator types.Validator, subtractAccount bool) (newShares sdk.Dec, err sdk.Error) {

// 	// Get or create the delegator delegation
// 	//Work-Around, refactor later
// 	delegation, found := nil, false //k.GetDelegation(ctx, delAddr, validator.OperatorAddr)
// 	if !found {
// 		delegation = types.Delegation{
// 			DelegatorAddr: delAddr,
// 			ValidatorAddr: validator.OperatorAddr,
// 			Shares:        sdk.ZeroDec(),
// 		}
// 	}

// 	if subtractAccount {
// 		// Account new shares, save
// 		_, _, err = k.bankKeeper.SubtractCoins(ctx, delegation.DelegatorAddr, types.Coins{bondAmt})
// 		if err != nil {
// 			return
// 		}
// 	}

// 	validator, newShares = k.AddValidatorTokensAndShares(ctx, validator, bondAmt.Amount)

// 	// Update delegation
// 	delegation.Shares = delegation.Shares.Add(newShares)
// 	delegation.Height = ctx.BlockHeight()
// 	//k.SetDelegation(ctx, delegation)
// 	return newShares, nil
// }
