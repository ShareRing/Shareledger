package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/sharering/shareledger/testutil/sample"
	denom "github.com/sharering/shareledger/x/utils/denom"
)

func TestMsgUpdateSwapFee_ValidateBasic(t *testing.T) {
	amount := sdk.NewDecCoin(denom.Base, sdk.NewInt(100000))
	tests := []struct {
		name string
		msg  MsgUpdateSwapFee
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateSwapFee{
				Creator: "invalid_address",
				In:      &amount,
				Out:     &amount,
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateSwapFee{
				Creator: sample.AccAddress(),
				In:      &amount,
				Out:     &amount,
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
