package keeper

import (
	"github.com/ShareRing/Shareledger/x/gentlemint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) IsAuthority(ctx sdk.Context, address sdk.AccAddress) bool {
	value, found := k.GetAuthority(ctx)
	return found && value.Address == address.String()
}

func (k Keeper) IsTreasurer(ctx sdk.Context, address sdk.AccAddress) bool {
	value, found := k.GetTreasurer(ctx)
	return found && value.Address == address.String()
}

func (k Keeper) IsAccountOperator(ctx sdk.Context, address sdk.AccAddress) bool {
	return k.isActive(ctx, address, types.AccStateKeyAccOp)
}

func (k Keeper) IsSHRPLoader(ctx sdk.Context, address sdk.AccAddress) bool {
	return k.isActive(ctx, address, types.AccStateKeyShrpLoaders)
}

func (k Keeper) IsIDSigner(ctx sdk.Context, address sdk.AccAddress) bool {
	return k.isActive(ctx, address, types.AccStateKeyIdSigner)
}

func (k Keeper) IsDocIssuer(ctx sdk.Context, address sdk.AccAddress) bool {
	return k.isActive(ctx, address, types.AccStateKeyDocIssuer)
}
