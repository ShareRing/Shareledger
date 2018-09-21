package accounts

import (


	"github.com/sharering/shareledger/types"
)


// GetKeyPair - return KeyPair
func GetKeyPair() (types.PubKeySecp256k1, types.PrivKeySecp256k1) {
	return types.GenerateKeyPair()
}
