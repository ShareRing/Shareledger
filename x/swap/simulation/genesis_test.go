package simulation

import (
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/sharering/shareledger/x/swap/types"
	"math/rand"
	"testing"
)

func TestMustGenRandGenesis(t *testing.T) {
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(interfaceRegistry)
	s := rand.NewSource(1)
	r := rand.New(s)

	simState := &module.SimulationState{
		AppParams: make(simtypes.AppParams),
		Cdc:       cdc,
		Rand:      r,
		NumBonded: 3,
		Accounts:  simtypes.RandomAccounts(r, 3),
		GenState:  make(map[string]json.RawMessage),
	}
	MustGenRandGenesis(simState)
	genStr, err := simState.GenState[types.ModuleName].MarshalJSON()
	if err != nil {
		t.Failed()
	}

	fmt.Println("genesis swap", string(genStr))
}
