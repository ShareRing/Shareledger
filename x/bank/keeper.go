package bank

import (
	"fmt"

	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
	"bitbucket.org/shareringvn/cosmos-sdk/wire"
	"github.com/sharering/shareledger/types"
	"github.com/sharering/shareledger/x/auth"
)

type Keeper struct {
	am auth.AccountMapper
}

func NewKeeper(bankKey sdk.StoreKey, _am auth.AccountMapper, cdc *wire.Codec) Keeper {
	return Keeper{am: _am}
}

func (k Keeper) SubtractCoins(ctx sdk.Context, addr sdk.Address, amt types.Coins) (types.Coins, sdk.Tags, sdk.Error) {

	return substractCoins(ctx, k.am, addr, amt)
}

func substractCoins(ctx sdk.Context, am auth.AccountMapper, addr sdk.Address, amt types.Coins) (types.Coins, sdk.Tags, sdk.Error) {
	oldCoins := getCoins(ctx, am, addr)
	newCoins := oldCoins.MinusMany(amt)
	if !newCoins.IsNotNegative() {
		return amt, nil, sdk.ErrInsufficientCoins(fmt.Sprintf("%s < %s", oldCoins, amt))
	}
	err := setCoins(ctx, am, addr, newCoins)
	tags := sdk.NewTags("sender", []byte(addr.String()))
	return newCoins, tags, err
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

	acc.SetCoins(amt)

	am.SetAccount(ctx, acc)
	return nil
}
