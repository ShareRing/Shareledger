package pos

import (
	"fmt"

	"github.com/pkg/errors"

	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
	abci "github.com/tendermint/abci/types"
)

// GenesisState - all staking state that must be provided at genesis
type GenesisState struct {
	Pool       Pool        `json:"pool"`
	Params     Params      `json:"params"`
	Validators []Validator `json:"validators"`
	//Bonds      []Delegation `json:"bonds"`
}

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) ([]abci.Validator, error) {

	var abciVals []abci.Validator
	fmt.Println("Genesis Pool", data.Pool)
	fmt.Println("Genesis Param", data.Params)
	for _, validator := range data.Validators {

		if validator.DelegatorShares.IsZero() {
			return abciVals, errors.Errorf("genesis validator cannot have zero delegator shares, validator: %v", validator)
		}

		fmt.Printf("InitGenesis validator: %s\n", validator.Owner)
		abciVals = append(abciVals, validator.ABCIValidator())
	}

	return abciVals, nil

}
func GenerateGenesis() GenesisState {
	gs := GenesisState{
		Pool:       InitialPool(),
		Validators: []Validator{},
	}
	return gs
}
