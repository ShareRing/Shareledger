package ante

import (
	idtypes "github.com/ShareRing/Shareledger/x/id/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type GentlemintKeeper interface {
	GetExchangeRateF(ctx sdk.Context) float64
}
type ElectoralKeeper interface {
}
type BankKeeper interface {
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
}

type RoleKeeper interface {
	IsAuthority(ctx sdk.Context, address sdk.AccAddress) bool
	IsSHRPLoader(ctx sdk.Context, address sdk.AccAddress) bool
	IsTreasurer(ctx sdk.Context, address sdk.AccAddress) bool
	IsIDSigner(ctx sdk.Context, address sdk.AccAddress) bool
	IsDocIssuer(ctx sdk.Context, address sdk.AccAddress) bool
}
type IDKeeper interface {
	GetIDByIdString(ctx sdk.Context, id string) *idtypes.ID
}