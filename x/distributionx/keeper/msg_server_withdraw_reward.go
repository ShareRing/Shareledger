package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sharering/shareledger/x/distributionx/types"
)

func (k msgServer) WithdrawReward(goCtx context.Context, msg *types.MsgWithdrawReward) (*types.MsgWithdrawRewardResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	signer := msg.GetSigners()[0]
	reward, _ := k.GetReward(ctx, signer.String())
	err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, signer, reward.Amount)
	if err != nil {
		return nil, err
	}
	// reset reward
	k.SetReward(ctx, types.Reward{
		Index:  signer.String(),
		Amount: []sdk.Coin{},
	})

	return &types.MsgWithdrawRewardResponse{
		Amount: reward.Amount,
	}, nil
}
