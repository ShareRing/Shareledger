package pos

import (
	"fmt"

	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/pos/keeper"
	abci "github.com/tendermint/abci/types"

	"github.com/sharering/shareledger/constants"
	"github.com/sharering/shareledger/types"
	posTypes "github.com/sharering/shareledger/x/pos/type"
)

func EndBlocker(ctx sdk.Context, k keeper.Keeper, proposer types.PubKeySecp256k1) []abci.Validator {

	// Proposer exists
	if !proposer.Equals(types.NilPubKeySecp256k1()) {

		address := proposer.Address()

		validator, found := k.GetValidator(ctx, address)

		if !found {
			panic(posTypes.ErrNoValidatorFound(posTypes.DefaultCodespace).Error())
		}

		// txt, _ := validator.HumanReadableString()
		// fmt.Println("UpdateBlockReward", txt)

		vdi, err := k.UpdateBlockReward(
			ctx,
			validator.Owner,
			validator.CommissionRate,
			types.NewPOSCoin(constants.POS_BLOCK_REWARD),
		)
		if err != nil {
			panic(err.Error())
		}

		constants.LOGGER.Info(fmt.Sprintf("Proposer %X", vdi.ValidatorAddr),
			"RewardAccum", vdi.RewardAccum.String(),
			"Commission", vdi.Commission.String(),
			"WithdrawHeight", vdi.WithdrawalHeight,
			"ValidatorReward", vdi.ValidatorReward.String(),
		)
		// fmt.Printf("ValidatorDistInfo: %v\n", vdi.HumanReadableString())
	}

	var valUpdates []abci.Validator
	/*if ValidatorChanged {
		valUpdates = k.GetValidatorSetUpdates(ctx) //work-around to get all ABCIValidators -> need to update
		for _, val := range valUpdates {
			fmt.Printf("Validator Update/abci Address=%X Power=%d\n", val.Address, val.Power)
		}
	} else {
		valUpdates = []abci.Validator{}
	}*/
	valUpdates = k.GetValidatorSetUpdates(ctx)
	return valUpdates
}
