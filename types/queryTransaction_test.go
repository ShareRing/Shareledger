package types

import (
	"testing"

	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
	"bitbucket.org/shareringvn/cosmos-sdk/wire"

	"github.com/sharering/shareledger/x/auth"
)

func TestQueryTransaction(t *testing.T) {
	pubKey := GetTestPubKey()

	// Create a Nonce Msg
	nonceMsg := auth.NewMsgNonce(pubKey.Address())

	queryTx := NewQueryTx(nonceMsg)

	// register codec
	cdc := wire.NewCodec()
	cdc.RegisterInterface((*types.SHRTx)(nil), nil)
	cdc.RegisterConcrete(types.BasicTx{}, "shareledger/BasicTx", nil)
	cdc.RegisterConcrete(types.QueryTx{}, "shareledger/QueryTx", nil)

	val, err := cdc.MarshalJSON(queryTx)
	if err != nil {
		t.Errorf("Marshal QueryTx failing. %s", err)
	}
	fmt.Printf("QueryTransaction: %s\n", val)

}
