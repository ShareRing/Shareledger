package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/testutil/sample"
	"github.com/sharering/shareledger/x/utils/denom"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateFormat_ValidateBasic(t *testing.T) {
	dec := sdk.NewDecCoin(denom.Base, sdk.NewInt(10000))
	tests := []struct {
		name string
		msg  MsgCreateSchema
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateSchema{
				Creator: "invalid_address",
				In:      dec,
				Out:     dec,
				Schema:  "schema",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateSchema{
				Creator: sample.AccAddress(),
				In:      dec,
				Out:     dec,
				Schema:  "schema",
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

func TestMsgUpdateFormat_ValidateBasic(t *testing.T) {
	dec := sdk.NewDecCoin(denom.Base, sdk.NewInt(10000))
	tests := []struct {
		name string
		msg  MsgUpdateSchema
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateSchema{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateSchema{
				Creator: sample.AccAddress(),
				Network: "eth",
				In:      &dec,
				Schema:  "schema",
				Out:     &dec,
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

func TestMsgDeleteFormat_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteSchema
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteSchema{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteSchema{
				Network: "eth",
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
