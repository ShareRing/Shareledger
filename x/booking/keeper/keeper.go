package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	assetkeeper "github.com/ShareRing/Shareledger/x/asset/keeper"
	"github.com/ShareRing/Shareledger/x/booking/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
)

type (
	Keeper struct {
		cdc         codec.BinaryCodec
		storeKey    sdk.StoreKey
		memKey      sdk.StoreKey
		assetKeeper assetkeeper.Keeper
		bankKeeper  bankkeeper.Keeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
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
