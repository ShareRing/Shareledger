package simulation

import (
	"math/rand"
	"sync"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/sharering/shareledger/testutil"
	"github.com/sharering/shareledger/x/electoral/types"
)

var (
	electoralAccount = map[string][]sdk.AccAddress{}
	rwMux            = sync.RWMutex{}
)

func GetElectoralAddress(r *rand.Rand, addressRole string) sdk.Address {
	rwMux.RLock()
	defer rwMux.RUnlock()
	accounts := electoralAccount[addressRole]
	if len(accounts) == 0 {
		return nil
	}
	if len(accounts) == 1 {
		return accounts[0]
	}

	return testutil.RandPick(r, accounts)
}

func MustGenRandGenesis(simState *module.SimulationState) {

	authority := testutil.RandPick(simState.Rand, simState.Accounts)
	treasure := testutil.RandPick(simState.Rand, simState.Accounts)

	electoralAccount["authority"] = []sdk.AccAddress{authority.Address}
	electoralAccount["treasure"] = []sdk.AccAddress{treasure.Address}

	accState := make([]types.AccState, 0, 24)

	for i := 0; i < 3; i++ {
		acc := testutil.RandPick(simState.Rand, simState.Accounts)
		accState = append(accState, types.AccState{
			Key:     string(types.GenAccStateIndexKey(acc.Address, types.AccStateKeyAccOp)),
			Address: acc.Address.String(),
			Status:  string(types.StatusActive),
		})
		electoralAccount["operator"] = append(electoralAccount["operator"], acc.Address)

		acc = testutil.RandPick(simState.Rand, simState.Accounts)
		accState = append(accState, types.AccState{
			Key:     string(types.GenAccStateIndexKey(acc.Address, types.AccStateKeyApprover)),
			Address: acc.Address.String(),
			Status:  string(types.StatusActive),
		})
		electoralAccount["approver"] = append(electoralAccount["approver"], acc.Address)

		acc = testutil.RandPick(simState.Rand, simState.Accounts)
		accState = append(accState, types.AccState{
			Key:     string(types.GenAccStateIndexKey(acc.Address, types.AccStateKeyDocIssuer)),
			Address: acc.Address.String(),
			Status:  string(types.StatusActive),
		})
		electoralAccount["docIssuer"] = append(electoralAccount["docIssuer"], acc.Address)

		acc = testutil.RandPick(simState.Rand, simState.Accounts)
		accState = append(accState, types.AccState{
			Key:     string(types.GenAccStateIndexKey(acc.Address, types.AccStateKeyIdSigner)),
			Address: acc.Address.String(),
			Status:  string(types.StatusActive),
		})
		electoralAccount["idSigner"] = append(electoralAccount["idSigner"], acc.Address)

		acc = testutil.RandPick(simState.Rand, simState.Accounts)
		accState = append(accState, types.AccState{
			Key:     string(types.GenAccStateIndexKey(acc.Address, types.AccStateKeyShrpLoaders)),
			Address: acc.Address.String(),
			Status:  string(types.StatusActive),
		})
		electoralAccount["loader"] = append(electoralAccount["loader"], acc.Address)

		acc = testutil.RandPick(simState.Rand, simState.Accounts)
		accState = append(accState, types.AccState{
			Key:     string(types.GenAccStateIndexKey(acc.Address, types.AccStateKeyRelayer)),
			Address: acc.Address.String(),
			Status:  string(types.StatusActive),
		})
		electoralAccount["relayer"] = append(electoralAccount["relayer"], acc.Address)

		acc = testutil.RandPick(simState.Rand, simState.Accounts)
		accState = append(accState, types.AccState{
			Key:     string(types.GenAccStateIndexKey(acc.Address, types.AccStateKeySwapManager)),
			Address: acc.Address.String(),
			Status:  string(types.StatusActive),
		})
		electoralAccount["swapManager"] = append(electoralAccount["swapManager"], acc.Address)

		acc = testutil.RandPick(simState.Rand, simState.Accounts)
		accState = append(accState, types.AccState{
			Key:     string(types.GenAccStateIndexKey(acc.Address, types.AccStateKeyVoter)),
			Address: acc.Address.String(),
			Status:  string(types.StatusActive),
		})
		electoralAccount["voter"] = append(electoralAccount["voter"], acc.Address)
	}

	genState := types.GenesisState{
		Authority: &types.Authority{
			Address: authority.Address.String(),
		},
		Treasurer:    &types.Treasurer{Address: treasure.Address.String()},
		AccStateList: accState,
	}

	electGenBz := simState.Cdc.MustMarshalJSON(&genState)

	simState.GenState[types.ModuleName] = electGenBz
	return

}
