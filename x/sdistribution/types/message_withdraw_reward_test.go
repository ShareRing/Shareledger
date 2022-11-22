package types

import (
	"testing"

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
			msg:  MsgWithdrawReward{},
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
