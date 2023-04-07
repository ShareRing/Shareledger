package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

type (
	BankKeeper interface {
		SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	}
	AccountKeeper interface {
		GetAccount(sdk.Context, sdk.AccAddress) authtypes.AccountI
	}
)
