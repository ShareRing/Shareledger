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

func (k Keeper) PShrMintPossible(ctx sdk.Context, amt sdk.Int) bool {
	total := k.bankKeeper.GetSupply(ctx, denom.PShr)
	newAmt := total.Amount.Add(amt)
	return newAmt.LT(types.MaxPSHRSupply)
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

func (k Keeper) buyPShr(ctx sdk.Context, amount sdk.Int, buyer sdk.AccAddress) error {
	if !k.PShrMintPossible(ctx, amount) {
		return sdkerrors.Wrap(types.ErrPSHRSupplyExceeded, amount.String())
	}

	rate := k.GetExchangeRateD(ctx)

	currentBalance := k.bankKeeper.GetAllBalances(ctx, buyer)
	currentShrpBalance := sdk.NewCoins(
		sdk.NewCoin(denom.ShrP, currentBalance.AmountOf(denom.ShrP)),
		sdk.NewCoin(denom.Cent, currentBalance.AmountOf(denom.Cent)),
	)
	cost, err := types.GetCostShrpForPShr(currentShrpBalance, amount, rate)
	if err != nil {
		return sdkerrors.Wrapf(err, "current %v balance", currentShrpBalance)
	}
	if cost.Sub.Empty() {
		return sdkerrors.ErrInsufficientFunds
	}

	if !cost.Add.Empty() {
		if err := k.loadCoins(ctx, buyer, cost.Add); err != nil {
			return sdkerrors.Wrapf(err, "%v coins in return", cost.Add)
		}
	}
	if err := k.burnCoins(ctx, buyer, cost.Sub); err != nil {
		return sdkerrors.Wrapf(err, "charge %v coins", cost.Sub)
	}
	boughtShr := sdk.NewCoins(sdk.NewCoin(denom.PShr, amount))
	if err := k.loadCoins(ctx, buyer, boughtShr); err != nil {
		return sdkerrors.Wrapf(err, "send %v coins", boughtShr)
	}
	return nil
}
