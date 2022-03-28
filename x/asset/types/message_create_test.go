package types

import (
	"github.com/google/uuid"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgCreate_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateAsset
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateAsset{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateAsset{
				Creator: sample.AccAddress(),
				UUID:    uuid.New().String(),
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
