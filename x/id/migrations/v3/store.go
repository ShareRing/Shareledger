package v1

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/id/types"
)

// MigrateStore performs in-place store migrations from v2 to v3. The
// migration includes:
//
// - Remove test data (ID have field extraData: "https://sharering.network/")
func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
	store := ctx.KVStore(storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.IDKeyPrefix))

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		bId := types.MustUnmarshalBaseID(cdc, iterator.Value())
		idKey := iterator.Key()[len(types.IDKeyPrefix):]
		id := types.NewIDFromBaseID(string(idKey), &bId)

		if id.Data.ExtraData == "https://sharering.network" {
			store.Delete(iterator.Key())
		}
	}
	return nil
}
