package auth

import (
	"testing"

	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
	"bitbucket.org/shareringvn/cosmos-sdk/wire"

	"github.com/sharering/shareledger/types"
)

func TestQueryTransaction(t *testing.T) {
	pubKey := types.GetTestPubKey()

	// Create a Nonce Msg
	nonceMsg := MsgNonce{
		Address: pubKey.Address(),
	}

	queryTx := types.NewQueryTx(nonceMsg)

	// register codec
	cdc := wire.NewCodec()
	cdc.RegisterInterface((*types.SHRTx)(nil), nil)
	cdc.RegisterConcrete(types.BasicTx{}, "shareledger/BasicTx", nil)
	cdc.RegisterConcrete(types.QueryTx{}, "shareledger/QueryTx", nil)

	cdc.RegisterInterface((*sdk.Msg)(nil), nil)
	cdc.RegisterConcrete(MsgNonce{}, "shareledger/auth/MsgNonce", nil)

	val, err := cdc.MarshalJSON(queryTx)
	if err != nil {
		t.Errorf("Marshal QueryTx failing. %s", err)
	}
	t.Logf("QueryTransaction: %s\n", val)

	//return nil,
	//sdk.ErrInternal(fmt.Sprintf("%s doesnt exist", addr.String())).Result()
	shrA := NewSHRAccountWithAddress(pubKey.Address())
	t.Logf("Nonce: %d\n", shrA.GetNonce())
}
