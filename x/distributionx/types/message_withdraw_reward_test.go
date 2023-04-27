package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgWithdrawReward_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgWithdrawReward
		err  error
	}{
		{
			name: "valid address",
			msg: MsgWithdrawReward{
				Creator: "shareledger18g8x9censnr3k2y7x6vwntlhvz254ym4qflcak",
			},
		},
		{
			name: "invalid address",
			msg: MsgWithdrawReward{
				Creator: "",
			},
			err: sdkerrors.ErrInvalidAddress,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
