package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	posTypes "github.com/sharering/shareledger/x/pos/type"
)

// GetValidatorByTDMAddress - get validator by Tendermint Address
func (k Keeper) GetValidatorByTDMAddress(
	ctx sdk.Context, tdmAddress []byte,
) (
	validator posTypes.Validator, found bool,
) {
	store := ctx.KVStore(k.storeKey)
	value := store.Get(GetTdmAddressKey(tdmAddress))
	if value == nil {
		return validator, false
	}

	return k.GetValidator(ctx, sdk.AccAddress(value))
}

// SetAddressByTDMAddress - set ShareledgerAddress by TDMAddress to later retrieve during POS processing
func (k Keeper) SetAddressByTDMAddress(
	ctx sdk.Context, tdmAddress []byte, address sdk.AccAddress,
) {
	store := ctx.KVStore(k.storeKey)
	store.Set(GetTdmAddressKey(tdmAddress), address[:])
}
