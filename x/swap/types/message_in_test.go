package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/testutil/sample"

	denom "github.com/sharering/shareledger/x/utils/denom"
	"github.com/stretchr/testify/require"
)

func TestMsgIn_ValidateBasic(t *testing.T) {
	amount := sdk.NewDecCoin(denom.Base, sdk.NewInt(100000))
	tests := []struct {
		name string
		msg  MsgRequestIn
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgRequestIn{
				Creator: "invalid_address",
				Amount:  &amount,
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgRequestIn{
				Creator:     sample.AccAddress(),
				DestAddress: sample.AccAddress(),
				Amount:      &amount,
				TxEvents: []*TxEvent{
					{
						TxHash:   "hash1",
						Sender:   "sender1",
						LogIndex: 0,
					},
				},
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
