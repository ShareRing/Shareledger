package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	idtypes "github.com/sharering/shareledger/x/id/types"
)

type GentlemintKeeper interface {
	GetExchangeRateF(ctx sdk.Context) sdk.Dec
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
	IsAccountOperator(ctx sdk.Context, address sdk.AccAddress) bool
	IsVoter(ctx sdk.Context, address sdk.AccAddress) bool
}
type IDKeeper interface {
	GetFullIDByIDString(ctx sdk.Context, id string) (*idtypes.Id, bool)
}
