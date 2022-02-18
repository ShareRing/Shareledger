package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgRevokeAccountOperator_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgRevokeAccountOperators
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgRevokeAccountOperators{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgRevokeAccountOperators{
				Creator:   sample.AccAddress(),
				Addresses: []string{sample.AccAddress()},
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
