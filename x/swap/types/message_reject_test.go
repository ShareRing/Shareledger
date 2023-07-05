package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/sharering/shareledger/testutil/sample"
)

func TestMsgReject_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgReject
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgReject{
				Creator: "invalid_address",
				Ids:     []uint64{23},
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgReject{
				Creator: sample.AccAddress(),
				Ids:     []uint64{23},
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
