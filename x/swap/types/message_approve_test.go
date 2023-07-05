package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/sharering/shareledger/testutil/sample"
)

func TestMsgApprove_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgApproveIn
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgApproveIn{
				Creator: "invalid_address",
				Ids:     []uint64{2, 3, 4},
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgApproveIn{
				Creator: sample.AccAddress(),
				Ids:     []uint64{2, 3, 4},
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
