package pos

import (
	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
	"github.com/pkg/errors"
	abci "github.com/tendermint/abci/types"
)

// GenesisState - all staking state that must be provided at genesis
type GenesisState struct {
	//Pool       Pool         `json:"pool"`
	//Params     Params       `json:"params"`
	Validators []Validator `json:"validators"`
	//Bonds      []Delegation `json:"bonds"`
}

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) ([]abci.Validator, error) {

	var abciVals []abci.Validator

	for _, validator := range data.Validators {

		if validator.DelegatorShares.IsZero() {
			return abciVals, errors.Errorf("genesis validator cannot have zero delegator shares, validator: %v", validator)
		}
		abciVals = append(abciVals, validator.ABCIValidator())
	}

	return abciVals, nil

}
