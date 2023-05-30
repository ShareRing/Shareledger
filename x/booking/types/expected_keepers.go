package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	assettypes "github.com/sharering/shareledger/x/asset/types"
)

type AssetKeeper interface {
	SetAssetStatus(ctx sdk.Context, uuid string, status bool)
	GetAsset(ctx sdk.Context, uuid string) (assettypes.Asset, bool)
}

type BankKeeper interface {
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
}
