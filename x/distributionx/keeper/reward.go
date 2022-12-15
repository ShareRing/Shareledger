package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/distributionx/types"
)

// SetReward set a specific reward in the store from its index
func (k Keeper) SetReward(ctx sdk.Context, reward types.Reward) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RewardKeyPrefix))
	b := k.cdc.MustMarshal(&reward)
	store.Set(types.RewardKey(reward.Index), b)
}

// GetReward returns a reward from its index
func (k Keeper) GetReward(ctx sdk.Context, index string) (val types.Reward, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RewardKeyPrefix))

	b := store.Get(types.RewardKey(index))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveReward removes a reward from the store
func (k Keeper) RemoveReward(ctx sdk.Context, index string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RewardKeyPrefix))
	store.Delete(types.RewardKey(index))
}

// GetAllReward returns all reward
func (k Keeper) GetAllReward(ctx sdk.Context) (list []types.Reward) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RewardKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Reward
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) IncReward(ctx sdk.Context, address string, coins sdk.Coins) {
	val, found := k.GetReward(ctx, address)
	if !found {
		val = types.Reward{
			Index:  address,
			Amount: nil,
		}
	}
	val.Amount = val.Amount.Add(coins...)
	k.SetReward(ctx, val)
}
