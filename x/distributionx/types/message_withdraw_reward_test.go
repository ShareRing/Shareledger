package types

import (
	"strings"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/app/params"
	"github.com/stretchr/testify/require"
)

func init() {
	params.SetAddressPrefixes()
}
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
				isContain := strings.Contains(err.Error(), tt.err.Error())
				require.True(t, isContain)
			} else {
				require.NoError(t, err)
			}

		})
	}
}
