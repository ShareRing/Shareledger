package pos

import (
	"bytes"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/pos/keeper"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/sharering/shareledger/constants"
	"github.com/sharering/shareledger/types"
	posTypes "github.com/sharering/shareledger/x/pos/type"
)

func EndBlocker(ctx sdk.Context, k keeper.Keeper, proposerAddress []byte) []abci.ValidatorUpdate {

	// Proposer exists
	if !bytes.Equal(proposerAddress, []byte{}) {

		validator, found := k.GetValidatorByTDMAddress(ctx, proposerAddress)

		if !found {
			panic(posTypes.ErrNoValidatorFound(posTypes.DefaultCodespace).Error())
		}

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

	var valUpdates []abci.ValidatorUpdate
	if ValidatorChanged {
		valUpdates = k.GetValidatorSetUpdates(ctx) //work-around to get all ABCIValidators -> need to update
		/*for _, val := range valUpdates {
			fmt.Printf("Validator Update/abci Address=%X Power=%d\n", val.Address, val.Power)
		}*/
	} else {
		valUpdates = []abci.ValidatorUpdate{}
	}
	//TODO: return updated Validators list
	return valUpdates
}
