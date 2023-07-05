package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/sharering/shareledger/testutil/sample"
)

func TestMsgRevokeDocIssuer_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgRevokeDocIssuers
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgRevokeDocIssuers{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgRevokeDocIssuers{
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
