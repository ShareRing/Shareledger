package sample

import (
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto/secp256k1"
)

// AccAddress returns a sample account address
func AccAddress() string {
	pk := ed25519.GenPrivKey().PubKey()
	addr := pk.Address()
	return sdk.AccAddress(addr).String()
}
func RandomAddr(amount int) (prvs []secp256k1.PrivKey, addrs []sdk.AccAddress, addStr []string) {
	for i := 0; i < amount; i++ {
		prv := secp256k1.GenPrivKey()
		addr := sdk.AccAddress(prv.PubKey().Address())

		addrs = append(addrs, addr)
		prvs = append(prvs, prv)
		addStr = append(addStr, addr.String())
	}
	return
}
