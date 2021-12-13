package ante

import sdk "github.com/cosmos/cosmos-sdk/types"

type GentlemintKeeper interface {
	GetExchangeRateF(ctx sdk.Context) float64
}
type ElectoralKeeper interface {
}
type BankKeeper interface {
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
}
