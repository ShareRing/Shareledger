package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

type BankKeeper interface {
	// Methods imported from bank should be defined here
	MintCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	LockedCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
	BurnCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
}

type AccountKeeper interface {
	// Methods imported from account should be defined here
	SetAccount(ctx sdk.Context, acc types.AccountI)
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
}
