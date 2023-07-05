package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/sharering/shareledger/testutil/sample"
)

func TestMsgReplaceIdOwner_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgReplaceIdOwner
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgReplaceIdOwner{
				BackupAddress: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgReplaceIdOwner{
				BackupAddress: sample.AccAddress(),
				Id:            "id",
				OwnerAddress:  sample.AccAddress(),
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
