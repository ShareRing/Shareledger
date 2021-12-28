package types

import (
	"testing"

	"github.com/sharering/shareledger/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgSendShrp_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgSendShrp
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgSendShrp{
				Creator: "invalid_address",
				Address: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgSendShrp{
				Creator: sample.AccAddress(),
				Address: sample.AccAddress(),
				Amount: "101",
			},
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
