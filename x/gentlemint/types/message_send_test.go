package types

import (
	"github.com/cosmos/cosmos-sdk/types"
	denom "github.com/sharering/shareledger/x/utils/demo"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgSend_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgSend
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgSend{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgSend{
				Creator: sample.AccAddress(),
				Address: sample.AccAddress(),
				Coins:   types.NewDecCoins(types.NewDecCoinFromDec(denom.Shr, types.MustNewDecFromStr("1.1"))),
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
