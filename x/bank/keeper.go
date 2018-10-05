package bank

import (
	"fmt"

	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
	"bitbucket.org/shareringvn/cosmos-sdk/wire"
	"github.com/sharering/shareledger/types"
	"github.com/sharering/shareledger/x/auth"
)

type Keeper struct {
}

func NewKeeper(bankKey sdk.StoreKey, cdc *wire.Codec) Keeper {
	return Keeper{}
}

func (k Keeper) subtractCoins(ctx sdk.Context, am auth.AccountMapper, addr sdk.Address, amt types.Coins) (types.Coins, sdk.Tags, sdk.Error) {

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

	err := acc.SetCoins(amt)
	if err != nil {
		// Handle w/ #870
		panic(err)
	}
	am.SetAccount(ctx, acc)
	return nil
}
