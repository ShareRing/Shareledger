package pos

import (
	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
	// "github.com/tendermint/tendermint/crypto"
	// "github.com/tendermint/tendermint/crypto/ed25519"

	"github.com/sharering/shareledger/types"
)

var (
	// pk1   = ed25519.GenPrivKey().PubKey()
	// pk2   = ed25519.GenPrivKey().PubKey()
	// pk3   = ed25519.GenPrivKey().PubKey()
	pk1, _ = types.GenerateKeyPair()
	pk2, _ = types.GenerateKeyPair()
	pk3, _ = types.GenerateKeyPair()
	addr1  = pk1.Address()
	addr2  = pk2.Address()
	addr3  = pk3.Address()

	emptyAddr   sdk.Address
	emptyPubkey types.PubKey
)
