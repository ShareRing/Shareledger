package pos

import (
	"fmt"

	"github.com/pkg/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

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

func InitGenesis(ctx sdk.Context, keeper keeper.Keeper, data GenesisState) ([]abci.ValidatorUpdate, error) {

	var abciVals []abci.ValidatorUpdate
	keeper.SetPool(ctx, data.Pool)
	keeper.SetParams(ctx, data.Params)
	keeper.InitIntraTxCounter(ctx)

	for _, validator := range data.Validators {

		constants.LOGGER.Info("Validator",
			"ShareledgerAddress", fmt.Sprintf("%X", validator.Owner),
			"TendermintPubKey", fmt.Sprintf("%X", validator.ABCIValidatorUpdate().PubKey),
			"tokens", fmt.Sprintf("%v", validator.Tokens),
			"power", fmt.Sprintf("%d", validator.ABCIValidatorUpdate().Power),
		)

		if validator.DelegatorShares.IsZero() {
			return abciVals, errors.Errorf("genesis validator cannot have zero delegator shares, validator: %v", validator)
		}
		abciVals = append(abciVals, validator.ABCIValidatorUpdate())

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

	validator.Tokens, _ = types.NewDecFromStr("2000000") // avoid zero tokens
	validator.Status = types.Bonded

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
