package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sharering/shareledger/x/utils/denom"
)

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	keeper Keeper
}

// NewMigrator returns a new Migrator.
func NewMigrator(keeper Keeper) Migrator {
	return Migrator{keeper: keeper}
}

// Migrate1to2 migrates from version 1 to 2.
func (m Migrator) Migrate1to2(ctx sdk.Context) error {
	return nil
}

// this version add paramSpace min_gas_price for `gentlemint`
func (m Migrator) Migrate2to3(ctx sdk.Context) error {
	// default min-gas-prices is 100_000_000 nshr
	m.keeper.SetMinGasPriceParam(ctx, sdk.NewDecCoins(sdk.NewDecCoin(denom.Base, sdk.NewInt(100_000))))
	return nil
}
