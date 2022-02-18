package types

import (
	"github.com/google/uuid"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgBook_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateBooking
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateBooking{
				Booker: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateBooking{
				Booker:   sample.AccAddress(),
				UUID:     uuid.New().String(),
				Duration: 1,
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
