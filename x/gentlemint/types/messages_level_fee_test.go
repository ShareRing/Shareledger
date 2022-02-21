package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/constant"
	denom "github.com/sharering/shareledger/x/utils/demo"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateLevelFee_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgSetLevelFee
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgSetLevelFee{
				Level:   string(constant.HighFee),
				Fee:     sdk.NewDecCoinFromDec(denom.ShrP, sdk.MustNewDecFromStr("10.9")),
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgSetLevelFee{
				Level:   string(constant.HighFee),
				Fee:     sdk.NewDecCoinFromDec(denom.ShrP, sdk.MustNewDecFromStr("10.9")),
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

func TestMsgSetLevelFee_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgSetLevelFee
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgSetLevelFee{
				Level:   string(constant.HighFee),
				Fee:     sdk.NewDecCoinFromDec(denom.ShrP, sdk.MustNewDecFromStr("10.9")),
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgSetLevelFee{
				Level:   string(constant.HighFee),
				Fee:     sdk.NewDecCoinFromDec(denom.ShrP, sdk.MustNewDecFromStr("10.9")),
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
				Level:   string(constant.HighFee),
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteLevelFee{
				Level:   string(constant.HighFee),
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
