package utils

import (
	"bufio"
	"fmt"
	"os"

	"github.com/sharering/shareledger/types"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func WriteFile(path string, pubKeys []types.PubKeySecp256k1, privKeys []types.PrivKeySecp256k1) {
	f, err := os.Create(path)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	for i := 0; i < len(pubKeys); i++ {
		// fmt.Printf("KeyPair: %v %v\n", pubKeys[i], privKeys[i])

		_, e := f.Write(pubKeys[i][:])
		Check(e)
		// _, e = f.Write([]byte(","))
		// Check(e)
		_, e = f.Write(privKeys[i][:])
		Check(e)
		// _, e = f.Write([]byte(";"))
		// Check(e)

	}
}

func ReadFile(path string) ([]types.PubKeySecp256k1, []types.PrivKeySecp256k1) {
	f, err := os.Open(path)

	defer f.Close()

	Check(err)

	reader := bufio.NewReader(f)

	pubKeys := make([]types.PubKeySecp256k1, 1)
	privKeys := make([]types.PrivKeySecp256k1, 1)

	for _, err1 := reader.Peek(1); err1 == nil; _, err1 = reader.Peek(1) {

		// pubK, err := reader.ReadBytes(byte(','))
		pubK, err := reader.Peek(65)
		Check(err)

		reader.Discard(65)

		fmt.Printf("%v\n", pubK)

		pubKeys = append(pubKeys, types.NewPubKeySecp256k1(pubK))

		// privK, err := reader.ReadBytes(byte(';'))
		privK, err := reader.Peek(32)
		Check(err)

		reader.Discard(32)

		privKeys = append(privKeys, types.NewPrivKeySecp256k1(privK))

	}

	return pubKeys, privKeys
}
