package bank

import (
	"fmt"

	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
	"github.com/sharering/shareledger/types"
	"github.com/sharering/shareledger/x/auth"
)

type Keeper struct {
	am auth.AccountMapper
}

func NewKeeper(_am auth.AccountMapper) Keeper {
	return Keeper{am: _am}
}

func (k Keeper) SubtractCoins(
	ctx sdk.Context,
	addr sdk.Address,
	amt types.Coins,
) (types.Coins, sdk.Error) {

	return substractCoins(ctx, k.am, addr, amt)
}

func (k Keeper) AddCoins(
	ctx sdk.Context,
	addr sdk.Address,
	amt types.Coins,
) (types.Coins, sdk.Error) {

	return addCoins(ctx, k.am, addr, amt)
}

func (k Keeper) SubtractCoin(
	ctx sdk.Context,
	addr sdk.Address,
	amt types.Coin,
) (types.Coins, sdk.Error) {

	return subtractCoin(ctx, k.am, addr, amt)
}

func (k Keeper) AddCoin(
	ctx sdk.Context,
	addr sdk.Address,
	amt types.Coin,
) (types.Coins, sdk.Error) {

	return addCoin(ctx, k.am, addr, amt)
}

func subtractCoin(
	ctx sdk.Context,
	am auth.AccountMapper,
	addr sdk.Address,
	amt types.Coin,
) (types.Coins, sdk.Error) {

	oldCoins := getCoins(ctx, am, addr)

	newCoins := oldCoins.Minus(amt)

	if !newCoins.IsNotNegative() {
		return oldCoins, sdk.ErrInsufficientCoins(fmt.Sprintf("%s < %s", oldCoins, amt))
	}

	err := setCoins(ctx, am, addr, newCoins)

	return newCoins, err

}

func substractCoins(
	ctx sdk.Context,
	am auth.AccountMapper,
	addr sdk.Address,
	amt types.Coins,
) (types.Coins, sdk.Error) {

	oldCoins := getCoins(ctx, am, addr)

	newCoins := oldCoins.MinusMany(amt)

	if !newCoins.IsNotNegative() {
		return oldCoins, sdk.ErrInsufficientCoins(fmt.Sprintf("%s < %s", oldCoins, amt))
	}

	err := setCoins(ctx, am, addr, newCoins)

	return newCoins, err
}

func addCoin(
	ctx sdk.Context,
	am auth.AccountMapper,
	addr sdk.Address,
	amt types.Coin,
) (types.Coins, sdk.Error) {

	oldCoins := getCoins(ctx, am, addr)

	newCoins := oldCoins.Plus(amt)

	if !newCoins.IsNotNegative() {
		return oldCoins, sdk.ErrInsufficientCoins(fmt.Sprintf("Error during coins addition: %s + %s", oldCoins, amt))
	}

	err := setCoins(ctx, am, addr, newCoins)
	return newCoins, err
}

func addCoins(
	ctx sdk.Context,
	am auth.AccountMapper,
	addr sdk.Address,
	amt types.Coins,
) (types.Coins, sdk.Error) {

	oldCoins := getCoins(ctx, am, addr)

	newCoins := oldCoins.PlusMany(amt)

	if !newCoins.IsNotNegative() {
		return oldCoins, sdk.ErrInsufficientCoins(fmt.Sprintf("Error during coins addition: %s + %s", oldCoins, amt))
	}

	err := setCoins(ctx, am, addr, newCoins)
	return newCoins, err
}

//______________________________________________________________________________________________

func getCoins(ctx sdk.Context, am auth.AccountMapper, addr sdk.Address) types.Coins {

	acc := am.GetAccount(ctx, addr)

	if acc == nil {
		return types.Coins{}
	}

	return acc.GetCoins()
}

func setCoins(ctx sdk.Context, am auth.AccountMapper, addr sdk.Address, amt types.Coins) sdk.Error {

	acc := am.GetAccount(ctx, addr)

	if acc == nil {
		acc = am.NewAccountWithAddress(ctx, addr)
	}

	// UPdate new coins to  accounts
	acc.SetCoins(amt)

	// Update new account to AccountMapper
	am.SetAccount(ctx, acc)

	return nil
}
