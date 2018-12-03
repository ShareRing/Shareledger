package pos

import (
	"fmt"
	"github.com/pkg/errors"

	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
	abci "github.com/tendermint/abci/types"

	"github.com/sharering/shareledger/constants"
	"github.com/sharering/shareledger/types"
	"github.com/sharering/shareledger/x/pos/keeper"
	posTypes "github.com/sharering/shareledger/x/pos/type"
)

// GenesisState - all staking state that must be provided at genesis
type GenesisState struct {
	Pool       posTypes.Pool         `json:"pool"`
	Params     posTypes.Params       `json:"params"`
	Validators []posTypes.Validator  `json:"validators"`
	Bonds      []posTypes.Delegation `json:"bonds"`
}

func InitGenesis(ctx sdk.Context, keeper keeper.Keeper, data GenesisState) ([]abci.Validator, error) {

	var abciVals []abci.Validator
	keeper.SetPool(ctx, data.Pool)
	keeper.SetParams(ctx, data.Params)
	keeper.InitIntraTxCounter(ctx)

	for _, validator := range data.Validators {

		constants.LOGGER.Info("Validator",
			"ShareledgerAddress", fmt.Sprintf("%X", validator.Owner),
			"TendermintAddress", fmt.Sprintf("%X", validator.ABCIValidator().Address),
		)

		if validator.DelegatorShares.IsZero() {
			return abciVals, errors.Errorf("genesis validator cannot have zero delegator shares, validator: %v", validator)
		}
		abciVals = append(abciVals, validator.ABCIValidator())

		keeper.SetValidator(ctx, validator)
		keeper.SetValidatorByPowerIndex(ctx, validator, data.Pool)
		// Manually set indices for the first time
		//keeper.SetValidatorByConsAddr(ctx, validator)
		//keeper.OnValidatorCreated(ctx, validator.OperatorAddr)

		vdi := posTypes.NewValidatorDistInfo(validator.Owner, int64(0))

		// Store ValidatorDistInfo
		keeper.SetValidatorDistInfo(ctx, vdi)
	}

	for _, delegation := range data.Bonds {
		keeper.SetDelegation(ctx, delegation)
		//keeper.OnDelegationCreated(ctx, delegation.DelegatorAddr, delegation.ValidatorAddr)
	}

	return abciVals, nil

}

func NewGenesisState(pool posTypes.Pool, params posTypes.Params, validators []posTypes.Validator, bonds []posTypes.Delegation) GenesisState {
	return GenesisState{
		Pool:       pool,
		Params:     params,
		Validators: validators,
		Bonds:      bonds,
	}
}

func GenerateGenesis(pubKey types.PubKeySecp256k1) GenesisState {
	validator := posTypes.NewValidator(
		pubKey.Address(),
		pubKey,
		posTypes.NewDescription("sharering", "", "sharering.network", ""))


	validator.Tokens = types.OneDec() // avoid zero tokens

	pool := posTypes.InitialPool()
	pool.LooseTokens = types.NewDec(3000000000) //hard-code with 3 billion loose-token
	pool.BondedTokens = types.ZeroDec()

	gs := GenesisState{
		Pool:       pool, //posTypes.InitialPool(),
		Params:     posTypes.DefaultParams(),
		Validators: []posTypes.Validator{validator},
	}
	return gs
}
