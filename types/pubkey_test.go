package types

import (
	"encoding/hex"
	"github.com/btcsuite/btcd/btcec"
)

func GetTestPubKey() PubKeySecp256k1 {
	pkBytes, err := hex.DecodeString("ab83994cf95abe45b9d8610524b3f8f8fd023d69f79449011cb5320d2ca180c5")

	if err != nil {
		fmt.Println("Error in DecodeString: ", err)
		return
	}

	privKey_, pubKey_ := btcec.PrivKeyFromBytes(btcec.S256(), pkBytes)

	serPubKey := pubKey_.SerializeUncompressed()
	var pubKey PubKeySecp256k1
	copy(pubKey[:], serPubKey[:65])

	//address := pubKey.Address()
	return pubKey
}
