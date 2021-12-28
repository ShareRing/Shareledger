package keeper

import (
	"github.com/sharering/shareledger/x/gentlemint/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)


func (k Keeper) GetExchangeRateF(ctx sdk.Context) float64 {
	v, found := k.GetExchangeRate(ctx)
	if !found {
		return types.DefaultExchangeRate
	}
	return v.Rate
}

// SetExchangeRate set exchangeRate in the store
func (k Keeper) SetExchangeRate(ctx sdk.Context, exchangeRate types.ExchangeRate) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ExchangeRateKey))
	b := k.cdc.MustMarshal(&exchangeRate)
	store.Set([]byte{0}, b)
}

// GetExchangeRate returns exchangeRate
func (k Keeper) GetExchangeRate(ctx sdk.Context) (val types.ExchangeRate, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ExchangeRateKey))
	b := store.Get([]byte{0})

	val = types.ExchangeRate{Rate: types.DefaultExchangeRate}

	if b != nil {
		k.cdc.MustUnmarshal(b, &val)
	}
	return val, true
}

// RemoveExchangeRate removes exchangeRate from the store
func (k Keeper) RemoveExchangeRate(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ExchangeRateKey))
	store.Delete([]byte{0})
}
