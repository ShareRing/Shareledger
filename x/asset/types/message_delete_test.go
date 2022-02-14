package types

import (
	"github.com/google/uuid"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgDelete_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteAsset
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteAsset{
				Owner: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteAsset{
				Owner: sample.AccAddress(),
				UUID:  uuid.New().String(),
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
