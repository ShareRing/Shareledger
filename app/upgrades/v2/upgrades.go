package v2

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	controllerkeeper "github.com/cosmos/ibc-go/v5/modules/apps/27-interchain-accounts/controller/keeper"
	icacontrollertypes "github.com/cosmos/ibc-go/v5/modules/apps/27-interchain-accounts/controller/types"
	icatypes "github.com/cosmos/ibc-go/v5/modules/apps/27-interchain-accounts/types"
	"github.com/sharering/shareledger/app/keepers"
)

// Please put devpoolaccount in plan.info
func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	keepers *keepers.AppKeepers,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		ctx.Logger().Info("start to run module migrations version:", UpgradeName, " ...")

		// this will fix export genesis fail when we upgrade to v2
		controllerGenesisState := icatypes.ControllerGenesisState{
			ActiveChannels:     []icatypes.ActiveChannel{},
			InterchainAccounts: []icatypes.RegisteredInterchainAccount{},
			Ports:              []string{},
			Params: icacontrollertypes.Params{
				ControllerEnabled: true,
			},
		}

		controllerkeeper.InitGenesis(ctx, keepers.ICAControllerKeeper, controllerGenesisState)
		vm, err := mm.RunMigrations(ctx, configurator, vm)

		p := keepers.DistributionxKeeper.GetParams(ctx)
		p.DevPoolAccount = plan.Info
		keepers.DistributionxKeeper.SetParams(ctx, p)
		return vm, err
	}
}
