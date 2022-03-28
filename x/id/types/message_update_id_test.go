package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgUpdateId_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateId
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateId{
				IssuerAddress: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateId{
				IssuerAddress: sample.AccAddress(),
				Id:            "id",
				ExtraData:     "extra-data",
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
