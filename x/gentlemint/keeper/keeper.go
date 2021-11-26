package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/ShareRing/Shareledger/x/gentlemint/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type (
	Keeper struct {
		cdc      codec.BinaryCodec
		storeKey sdk.StoreKey
		memKey   sdk.StoreKey

		bankKeeper    types.BankKeeper
		accountKeeper types.AccountKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,

	bankKeeper types.BankKeeper, accountKeeper types.AccountKeeper,
) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,

		bankKeeper: bankKeeper, accountKeeper: accountKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) ShrMintPossible(ctx sdk.Context, amt sdk.Int) bool {
	total := k.bankKeeper.Get
}

func (k Keeper) LoadCoins(ctx sdk.Context, toAddr sdk.AccAddress, amt sdk.Coins) error {
	if !amt.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, amt.String())
	}
	// moduleAcct := sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, toAddr, amt); err != nil {
		return sdkerrors.Wrap(err, fmt.Sprintf("send coins to account %s", toAddr.String()))
	}
	return k.bankKeeper.MintCoins(ctx, types.ModuleName, amt)
}

func (k Keeper) BurnCoins(ctx sdk.Context, addr sdk.AccAddress, amt sdk.Coins) error {
	if !amt.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, amt.String())
	}

	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, addr, types.ModuleName, amt); err != nil {
		return sdkerrors.Wrapf(err, "send coins to module, amt %s", amt.String())
	}

	return k.bankKeeper.BurnCoins(ctx, types.ModuleName, amt)
}
