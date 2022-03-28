package keeper

import (
	"fmt"
	denom "github.com/sharering/shareledger/x/utils/demo"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/x/gentlemint/types"
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

func (k Keeper) BaseMintPossible(ctx sdk.Context, amt sdk.Int) bool {
	total := k.bankKeeper.GetSupply(ctx, denom.Base)
	newAmt := total.Amount.Add(amt)
	return newAmt.LT(types.MaxBaseSupply)
}

// loadCoins mint amt coins to module address and then send coins to account toAddr
func (k Keeper) loadCoins(ctx sdk.Context, toAddr sdk.AccAddress, amt sdk.Coins) error {
	if !amt.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, amt.String())
	}
	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, amt); err != nil {
		return sdkerrors.Wrapf(err, "mint %v coins to module %v", amt, types.ModuleName)
	}
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, toAddr, amt); err != nil {
		return sdkerrors.Wrapf(err, "send coins to account %s", toAddr.String())
	}
	return nil
}

// burnCoins send amt from address to module address then burning
func (k Keeper) burnCoins(ctx sdk.Context, addr sdk.AccAddress, amt sdk.Coins) error {
	if !amt.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, amt.String())
	}

	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, addr, types.ModuleName, amt); err != nil {
		return sdkerrors.Wrapf(err, "send coins to module, amt %s", amt.String())
	}

	return k.bankKeeper.BurnCoins(ctx, types.ModuleName, amt)
}

// LoadAllowanceLoader loads allowance coins to loader
func (k Keeper) LoadAllowanceLoader(ctx sdk.Context, addr sdk.AccAddress) error {
	return k.loadCoins(ctx, addr, types.AllowanceLoader)
}

func (k Keeper) buyBaseDenom(ctx sdk.Context, base sdk.Coin, buyer sdk.AccAddress) error {
	if base.Denom != denom.Base || !base.IsValid() {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "%v", base)
	}
	if !k.BaseMintPossible(ctx, base.Amount) {
		return sdkerrors.Wrap(types.ErrBaseSupplyExceeded, base.String())
	}

	rate := k.GetExchangeRateD(ctx)

	currentBalance := k.bankKeeper.GetAllBalances(ctx, buyer)

	cost, err := denom.NormalizeToBaseCoin(denom.BaseUSD, sdk.NewDecCoinsFromCoins(base), rate, true)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrLogic, "+%v", err)
	}
	if !currentBalance.IsAllGTE(sdk.NewCoins(cost)) {
		return sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds, "current balances %v, Cost: %v", currentBalance, cost)
	}
	if err := k.burnCoins(ctx, buyer, sdk.NewCoins(cost)); err != nil {
		return sdkerrors.Wrapf(err, "charge %v coins", cost)
	}

	if err := k.loadCoins(ctx, buyer, sdk.NewCoins(base)); err != nil {
		return sdkerrors.Wrapf(err, "load %v coins", base)
	}
	return nil
}
