package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/sharering/shareledger/testutil/sample"
)

func TestMsgCreateActionLevelFee_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateActionLevelFee
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateActionLevelFee{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateActionLevelFee{
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

func TestMsgUpdateActionLevelFee_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateActionLevelFee
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateActionLevelFee{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateActionLevelFee{
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

func TestMsgDeleteActionLevelFee_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteActionLevelFee
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteActionLevelFee{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteActionLevelFee{
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
