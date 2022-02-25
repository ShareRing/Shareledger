package sample

import (
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	"github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// AccAddress returns a sample account address
func AccAddress() string {
	pk := ed25519.GenPrivKey().PubKey()
	addr := pk.Address()
	return sdk.AccAddress(addr).String()
}
func RandomAddr(amount int) (prvs []types.PrivKey, addrs []sdk.AccAddress, addStr []string) {
	for i := 0; i < amount; i++ {
		prv, _, addr := testdata.KeyTestPubAddr()
		addrs = append(addrs, addr)
		prvs = append(prvs, prv)
		addStr = append(addStr, addr.String())
	}

	return
}
