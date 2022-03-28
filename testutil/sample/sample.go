package sample

import (
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
)

// AccAddress returns a sample account address
func AccAddress() string {
	_, _, addr := testdata.KeyTestPubAddr()
	return addr.String()
}
