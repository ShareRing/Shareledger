package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/tendermint/tendermint/libs/log"

	assetkeeper "github.com/sharering/shareledger/x/asset/keeper"
	"github.com/sharering/shareledger/x/booking/types"
)

type (
	Keeper struct {
		cdc         codec.BinaryCodec
		storeKey    storetypes.StoreKey
		memKey      storetypes.StoreKey
		assetKeeper assetkeeper.Keeper
		bankKeeper  bankkeeper.Keeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ask assetkeeper.Keeper,
	bankKeeper bankkeeper.Keeper,
) *Keeper {
	return &Keeper{
		cdc:         cdc,
		storeKey:    storeKey,
		memKey:      memKey,
		assetKeeper: ask,
		bankKeeper:  bankKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
