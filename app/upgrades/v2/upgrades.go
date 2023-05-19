package v2

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	controllerkeeper "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/controller/keeper"
	genesistypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/genesis/types"

	"github.com/sharering/shareledger/app/keepers"
)

// Please put devpoolaccount in plan.info
func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	keepers *keepers.AppKeepers,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _ upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		ctx.Logger().Info("start to run module migrations version:", UpgradeName, " ...")

		controllerGenesisState := genesistypes.DefaultControllerGenesis()

		controllerkeeper.InitGenesis(ctx, keepers.ICAControllerKeeper, controllerGenesisState)
		vm, err := mm.RunMigrations(ctx, configurator, vm)

		p := keepers.DistributionxKeeper.GetParams(ctx)
		// TODO: update this
		p.DevPoolAccount = "shareledger18pf3zdwqjntd9wkvfcjvmdc7hua6c0q2eck5h5"
		keepers.DistributionxKeeper.SetParams(ctx, p)
		return vm, err
	}
}
