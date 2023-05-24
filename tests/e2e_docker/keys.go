package e2e

import (
	"github.com/cosmos/go-bip39"
)

// createMnemonic creates a random string mnemonic
func createMnemonic() string {
	entropySeed, err := bip39.NewEntropy(256)
	if err != nil {
		panic(err)
	}

	mnemonic, err := bip39.NewMnemonic(entropySeed)
	if err != nil {
		panic(err)
	}

	return mnemonic
}
