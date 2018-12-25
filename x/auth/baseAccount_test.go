package auth

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/go-amino"
)

func TestBaseAccount(t *testing.T) {

	cdc := amino.NewCodec()
	cdc.RegisterInterface((*BaseAccount)(nil), nil)
	cdc.RegisterConcrete(SHRAccount{}, "shareledger/SHRAccount", nil)

	addr := sdk.AccAddress([]byte("405C725BC461DCA455B8AA84769E8ACE6B3763F4"))
	acc := NewSHRAccountWithAddress(addr)
	t.Logf("Account: %s", acc)
	acc.SetNonce(2)

}
