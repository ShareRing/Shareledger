package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/sharering/shareledger/app/params"
	"github.com/sharering/shareledger/testutil/sample"
	denom "github.com/sharering/shareledger/x/utils/denom"
)

func init() {
	params.SetAddressPrefixes()
}
func TestMsgWithdraw_ValidateBasic(t *testing.T) {
	amount := sdk.NewDecCoin(denom.Base, sdk.NewInt(100000))
	tests := []struct {
		name string
		msg  MsgWithdraw
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgWithdraw{
				Creator:  "invalid_address",
				Receiver: "shareledger1uf4vhge3qte80k0v74vepxz9rssfl09l7clknh",
				Amount:   amount,
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgWithdraw{
				Creator:  sample.AccAddress(),
				Amount:   amount,
				Receiver: "shareledger1uf4vhge3qte80k0v74vepxz9rssfl09l7clknh",
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
