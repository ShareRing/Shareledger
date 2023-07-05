package types

import (
	"strings"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/sharering/shareledger/app/params"
	"github.com/sharering/shareledger/testutil/sample"
)

func init() {
	params.SetAddressPrefixes()
}
func TestMsgEnrollSwapManager_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgEnrollSwapManagers
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgEnrollSwapManagers{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgEnrollSwapManagers{
				Creator: sample.AccAddress(),
				Addresses: []string{
					sample.AccAddress(),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				isContain := strings.Contains(err.Error(), tt.err.Error())
				require.True(t, isContain)
				return
			}
			require.NoError(t, err)
		},
		)
	}
}
