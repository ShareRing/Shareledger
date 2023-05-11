package types

import (
	"strings"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgEnrollApprover_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgEnrollApprovers
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgEnrollApprovers{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgEnrollApprovers{
				Creator:   sample.AccAddress(),
				Addresses: []string{sample.AccAddress()},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				isContain := strings.Contains(err.Error(), tt.err.Error())
				require.True(t, isContain)
				return
			}
			require.NoError(t, err)

		},
		)
	}
}
