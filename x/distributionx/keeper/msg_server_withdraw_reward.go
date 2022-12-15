package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/distributionx/types"
)

func (k msgServer) WithdrawReward(goCtx context.Context, msg *types.MsgWithdrawReward) (*types.MsgWithdrawRewardResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// send reward to the first signer than reset amount
	signer := msg.GetSigners()[0]
	reward, _ := k.GetReward(ctx, signer.String())
	k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.FeeWasmName, signer, reward.Amount)
	k.SetReward(ctx, types.Reward{
		Index:  signer.String(),
		Amount: []sdk.Coin{},
	})

	return &types.MsgWithdrawRewardResponse{
		Amount: reward.Amount,
	}, nil
}
