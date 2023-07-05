package app_test

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime/debug"
	"strings"
	"testing"

	"github.com/CosmWasm/wasmd/x/wasm"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/simapp"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/kv"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/sharering/shareledger/app"
	"github.com/sharering/shareledger/app/params"
	distributionxType "github.com/sharering/shareledger/x/distributionx/types"
	elecTypes "github.com/sharering/shareledger/x/electoral/types"
	gentlemintmoduletypes "github.com/sharering/shareledger/x/gentlemint/types"
	swapmoduletypes "github.com/sharering/shareledger/x/swap/types"
)

// Get flags every time the simulator is run
func init() {
	simapp.GetSimulatorFlags()

}

type StoreKeysPrefixes struct {
	A        storetypes.StoreKey
	B        storetypes.StoreKey
	Prefixes [][]byte
}

var emptyWasmOpts []wasm.Option = nil

func fauxMerkleModeOpt(bapp *baseapp.BaseApp) {
	bapp.SetFauxMerkleMode()
}

type (
	SkipCheckVoterOpts struct {
	}
)

func (opt SkipCheckVoterOpts) Get(o string) interface{} {
	if o == app.FlagAppOptionSkipCheckVoter {
		return true
	}
	return false
}

func TestAppImportExport(t *testing.T) {
	config, db, dir, logger, skip, err := simapp.SetupSimulation("leveldb-app-sim", "Simulation")
	if skip {
		t.Skip("skipping application import/export simulation")
	}
	require.NoError(t, err, "simulation setup failed")
	params.SetAddressPrefixes()
	defer func() {
		_ = db.Close()
		require.NoError(t, os.RemoveAll(dir))
	}()

	simApp := app.New(logger,
		db,
		nil,
		true,
		map[int64]bool{},
		app.DefaultNodeHome,
		simapp.FlagPeriodValue,
		app.MakeTestEncodingConfig(),
		SkipCheckVoterOpts{},
		fauxMerkleModeOpt,
	)

	_, simParams, simErr := simulation.SimulateFromSeed(
		t,
		os.Stdout,
		simApp.BaseApp,
		simapp.AppStateFn(simApp.AppCodec(), simApp.SimulationManager()),
		simtypes.RandomAccounts, // Replace with own random account function if using keys other than secp256k1
		simapp.SimulationOperations(simApp, simApp.AppCodec(), config),
		simApp.ModuleAccountAddrs(),
		config,
		simApp.AppCodec(),
	)

	require.NoError(t, simErr)

	err = simapp.CheckExportSimulation(simApp, config, simParams)
	require.NoError(t, err)
	if config.Commit {
		simapp.PrintStats(db)
	}
	fmt.Printf("exporting genesis...\n")

	exported, err := simApp.ExportAppStateAndValidators(false, []string{})
	require.NoError(t, err)

	fmt.Printf("importing genesis...\n")

	_, newDb, newDir, _, _, err := simapp.SetupSimulation("leveldb-app-sim-2", "Simulation-2")
	require.NoError(t, err, "simulation setup failed")

	defer func() {
		newDb.Close()
		require.NoError(t, os.RemoveAll(newDir))
	}()
	newSimApp := app.New(
		log.NewNopLogger(),
		newDb,
		nil,
		true,
		map[int64]bool{},
		app.DefaultNodeHome,
		simapp.FlagPeriodValue,
		app.MakeTestEncodingConfig(),
		simapp.EmptyAppOptions{},
		fauxMerkleModeOpt,
	)

	var genesisState app.GenesisState
	err = json.Unmarshal(exported.AppState, &genesisState)
	require.NoError(t, err)

	defer func() {
		if r := recover(); r != nil {
			err := fmt.Sprintf("%v", r)
			if !strings.Contains(err, "validator set is empty after InitGenesis") {
				panic(r)
			}
			logger.Info("Skipping simulation as all validators have been unbonded")
			logger.Info("err", err, "stacktrace", string(debug.Stack()))
		}
	}()

	ctxA := simApp.NewContext(true, tmproto.Header{Height: simApp.LastBlockHeight()})
	ctxB := newSimApp.NewContext(true, tmproto.Header{Height: simApp.LastBlockHeight()})
	newSimApp.GetMM().InitGenesis(ctxB, simApp.AppCodec(), genesisState)
	newSimApp.StoreConsensusParams(ctxB, exported.ConsensusParams)
	fmt.Printf("comparing stores...\n")

	storeKeysPrefixes := []StoreKeysPrefixes{
		{simApp.AppKeepers.GetKey(authtypes.StoreKey), newSimApp.AppKeepers.GetKey(authtypes.StoreKey), [][]byte{}},
		{simApp.AppKeepers.GetKey(stakingtypes.StoreKey), newSimApp.AppKeepers.GetKey(stakingtypes.StoreKey),
			[][]byte{
				stakingtypes.UnbondingQueueKey, stakingtypes.RedelegationQueueKey, stakingtypes.ValidatorQueueKey,
				stakingtypes.HistoricalInfoKey,
			}}, // ordering may change but it doesn't matter
		{simApp.AppKeepers.GetKey(slashingtypes.StoreKey), newSimApp.AppKeepers.GetKey(slashingtypes.StoreKey), [][]byte{}},
		{simApp.AppKeepers.GetKey(minttypes.StoreKey), newSimApp.AppKeepers.GetKey(minttypes.StoreKey), [][]byte{}},
		{simApp.AppKeepers.GetKey(distrtypes.StoreKey), newSimApp.AppKeepers.GetKey(distrtypes.StoreKey), [][]byte{}},
		{simApp.AppKeepers.GetKey(banktypes.StoreKey), newSimApp.AppKeepers.GetKey(banktypes.StoreKey), [][]byte{banktypes.BalancesPrefix}},
		{simApp.AppKeepers.GetKey(paramtypes.StoreKey), newSimApp.AppKeepers.GetKey(paramtypes.StoreKey), [][]byte{}},
		{simApp.AppKeepers.GetKey(govtypes.StoreKey), newSimApp.AppKeepers.GetKey(govtypes.StoreKey), [][]byte{}},
		{simApp.AppKeepers.GetKey(evidencetypes.StoreKey), newSimApp.AppKeepers.GetKey(evidencetypes.StoreKey), [][]byte{}},
		{simApp.AppKeepers.GetKey(capabilitytypes.StoreKey), newSimApp.AppKeepers.GetKey(capabilitytypes.StoreKey), [][]byte{}},
		{simApp.AppKeepers.GetKey(authzkeeper.StoreKey), newSimApp.AppKeepers.GetKey(authzkeeper.StoreKey), [][]byte{authzkeeper.GrantKey, authzkeeper.GrantQueuePrefix}},
		{simApp.AppKeepers.GetKey(distributionxType.StoreKey), newSimApp.AppKeepers.GetKey(distributionxType.StoreKey), [][]byte{}},
		{simApp.AppKeepers.GetKey(swapmoduletypes.StoreKey), newSimApp.AppKeepers.GetKey(swapmoduletypes.StoreKey), [][]byte{}},
		{simApp.AppKeepers.GetKey(elecTypes.StoreKey), newSimApp.AppKeepers.GetKey(elecTypes.StoreKey), [][]byte{}},
		{simApp.AppKeepers.GetKey(gentlemintmoduletypes.StoreKey), newSimApp.AppKeepers.GetKey(gentlemintmoduletypes.StoreKey), [][]byte{}},
	}

	for _, skp := range storeKeysPrefixes {
		storeA := ctxA.KVStore(skp.A)
		storeB := ctxB.KVStore(skp.B)

		failedKVAs, failedKVBs := sdk.DiffKVStores(storeA, storeB, skp.Prefixes)
		require.Equal(t, len(failedKVAs), len(failedKVBs), "unequal sets of key-values to compare")

		fmt.Printf("compared %d different key/value pairs between %s and %s\n", len(failedKVAs), skp.A, skp.B)
		require.Equal(t, 0, len(failedKVAs), GetSimulationLog(skp.A.Name(), simApp.SimulationManager().StoreDecoders, failedKVAs, failedKVBs))
	}
}

func TestAppFullSimulation(t *testing.T) {
	params.SetAddressPrefixes()
	config, db, dir, logger, skip, err := simapp.SetupSimulation("leveldb-app-sim", "Simulation")
	if skip {
		t.Skip("skipping application import/export simulation")
	}

	require.NoError(t, err, "simulation setup failed")

	defer func() {
		require.NoError(t, db.Close())
		require.NoError(t, os.RemoveAll(dir))
	}()

	simApp := app.New(logger,
		db,
		nil,
		true,
		map[int64]bool{},
		app.DefaultNodeHome,
		simapp.FlagPeriodValue,
		app.MakeTestEncodingConfig(),
		SkipCheckVoterOpts{},
		fauxMerkleModeOpt,
	)

	_, simParams, simErr := simulation.SimulateFromSeed(
		t,
		os.Stdout,
		simApp.BaseApp,
		simapp.AppStateFn(simApp.AppCodec(), simApp.SimulationManager()),
		simtypes.RandomAccounts, // Replace with own random account function if using keys other than secp256k1
		simapp.SimulationOperations(simApp, simApp.AppCodec(), config),
		simApp.ModuleAccountAddrs(),
		config,
		simApp.AppCodec(),
	)

	require.NoError(t, simErr)

	err = simapp.CheckExportSimulation(simApp, config, simParams)
	require.NoError(t, err)
	if config.Commit {
		simapp.PrintStats(db)
	}

}

// GetSimulationLog unmarshals the KVPair's Value to the corresponding type based on the
// each's module store key and the prefix bytes of the KVPair's key.
func GetSimulationLog(storeName string, sdr sdk.StoreDecoderRegistry, kvAs, kvBs []kv.Pair) (log string) {
	for i := 0; i < len(kvAs); i++ {
		if len(kvAs[i].Value) == 0 && len(kvBs[i].Value) == 0 {
			// skip if the value doesn't have any bytes
			continue
		}

		decoder, ok := sdr[storeName]
		if ok {
			log += decoder(kvAs[i], kvBs[i])
		} else {
			log += fmt.Sprintf("store A %s => %s\nstore B %s => %s\n", string(kvAs[i].Key), string(kvAs[i].Value), string(kvBs[i].Key), string(kvBs[i].Value))
		}
	}

	return log
}
