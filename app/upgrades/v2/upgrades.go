package v2

import (
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	tmtypes "github.com/cometbft/cometbft/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	govv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	controllerkeeper "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/controller/keeper"
	genesistypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/genesis/types"

	"github.com/sharering/shareledger/app/keepers"
	idtypes "github.com/sharering/shareledger/x/id/types"
)

// Please put devpoolaccount in plan.info
func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	keepers *keepers.AppKeepers,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _ upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		ctx.Logger().Info("start to run module migrations version:", UpgradeName, " ...")

		// update genesis for ICAController
		controllerGenesisState := genesistypes.DefaultControllerGenesis()
		controllerkeeper.InitGenesis(ctx, keepers.ICAControllerKeeper, controllerGenesisState)
		// ================================================

		// modify DistributionxParams
		p := keepers.DistributionxKeeper.GetParams(ctx)
		p.DevPoolAccount = "shareledger1qda5j5mgvmesceg8h68przllcxhz5lzvwz38v5"
		keepers.DistributionxKeeper.SetParams(ctx, p)
		// ================================================

		// move from legacy params module to new in-module store
		for _, subspace := range keepers.ParamsKeeper.GetSubspaces() {
			subspace := subspace

			var keyTable paramstypes.KeyTable
			switch subspace.Name() {
			case authtypes.ModuleName:
				keyTable = authtypes.ParamKeyTable() //nolint:staticcheck
			case banktypes.ModuleName:
				keyTable = banktypes.ParamKeyTable() //nolint:staticcheck
			case stakingtypes.ModuleName:
				keyTable = stakingtypes.ParamKeyTable() //nolint:staticcheck
			case minttypes.ModuleName:
				keyTable = minttypes.ParamKeyTable() //nolint:staticcheck
			case distrtypes.ModuleName:
				keyTable = distrtypes.ParamKeyTable() //nolint:staticcheck
			case slashingtypes.ModuleName:
				keyTable = slashingtypes.ParamKeyTable() //nolint:staticcheck
			case govtypes.ModuleName:
				keyTable = govv1.ParamKeyTable() //nolint:staticcheck
			case crisistypes.ModuleName:
				keyTable = crisistypes.ParamKeyTable() //nolint:staticcheck
			case idtypes.ModuleName:
				keyTable = crisistypes.ParamKeyTable() //nolint:staticcheck
			case wasmtypes.ModuleName:
				keyTable = wasmtypes.ParamKeyTable() //nolint:staticcheck
			default:
				continue
			}
			if !subspace.HasKeyTable() {
				subspace.WithKeyTable(keyTable)
			}
		}
		baseAppLegacySS := keepers.ParamsKeeper.Subspace(baseapp.Paramspace).WithKeyTable(paramstypes.ConsensusParamsKeyTable())
		baseapp.MigrateParams(ctx, baseAppLegacySS, &keepers.ConsensusParamsKeeper)
		consensusParams := baseapp.GetConsensusParams(ctx, baseAppLegacySS)
		// make sure the consensus params are set
		if consensusParams.Block == nil || consensusParams.Evidence == nil || consensusParams.Validator == nil {
			defaultParams := tmtypes.DefaultConsensusParams().ToProto()
			keepers.ConsensusParamsKeeper.Set(ctx, &defaultParams)
		}
		// ======================================================

		return mm.RunMigrations(ctx, configurator, vm)
	}
}
