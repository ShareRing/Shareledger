package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/sharering/shareledger/testutil/sample"
)

func TestMsgCreateLevelFee_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateLevelFee
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateLevelFee{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateLevelFee{
				Creator: sample.AccAddress(),
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

func TestMsgUpdateLevelFee_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateLevelFee
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateLevelFee{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateLevelFee{
				Creator: sample.AccAddress(),
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

func TestMsgDeleteLevelFee_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteLevelFee
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteLevelFee{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteLevelFee{
				Creator: sample.AccAddress(),
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
