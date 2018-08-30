package auth

import (
	"testing"

	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
	"bitbucket.org/shareringvn/cosmos-sdk/wire"
)

func TestBaseAccount(t *testing.T) {

	cdc := wire.NewCodec()
	cdc.RegisterInterface((*BaseAccount)(nil), nil)
	cdc.RegisterConcrete(SHRAccount{}, "shareledger/SHRAccount", nil)

	addr := sdk.Address([]byte("405C725BC461DCA455B8AA84769E8ACE6B3763F4"))
	acc := NewSHRAccountWithAddress(addr)
	t.Logf("Account: %s", acc)
	acc.SetNonce(2)

}
