package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/sharering/shareledger/testutil/sample"
	"github.com/sharering/shareledger/x/utils/denom"
)

func TestMsgDeposit_ValidateBasic(t *testing.T) {
	amount := sdk.NewDecCoin(denom.Base, sdk.NewInt(100000))
	tests := []struct {
		name string
		msg  MsgDeposit
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeposit{
				Creator: "invalid_address",
				Amount:  &amount,
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeposit{
				Creator: sample.AccAddress(),
				Amount:  &amount,
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
