package types

import (
	"fmt"

	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
	"github.com/btcsuite/btcd/btcec"
	crypto "github.com/tendermint/go-crypto"
)

type PrivKey interface {
	Sign(sdk.Msg) Signature
	SignWithNonce(sdk.Msg, int64) Signature
	String() string
}

//--------------------------
// Implement interface

var _ PrivKey = PrivKeySecp256k1{}

type PrivKeySecp256k1 [32]byte

func NewPrivKeySecp256k1(b []byte) PrivKeySecp256k1 {
	if len(b) != 32 {
		panic("Length of input to create PrivKeySecp256k1 should be 32")
	}
	var privK PrivKeySecp256k1

	copy(privK[:], b[:32])

	return privK
}

func (privKey PrivKeySecp256k1) Sign(msg sdk.Msg) Signature {

	signBytes := msg.GetSignBytes()

	msgHash := crypto.Sha256(signBytes)

	btcecPrivKey, _ := btcec.PrivKeyFromBytes(btcec.S256(), privKey[:])

	signature, err := btcecPrivKey.Sign(msgHash)

	if err != nil {
		panic(err)
	}

	serSig := signature.Serialize()

	var ecSig SignatureSecp256k1

	ecSig = append(ecSig, serSig...)

	return ecSig
}

func (privKey PrivKeySecp256k1) SignWithNonce(msg sdk.Msg, nonce int64) Signature {
	signBytes := msg.GetSignBytes()

	signBytesWithNonce := append([]byte(fmt.Sprintf("%d", nonce)), signBytes...)

	msgHash := crypto.Sha256(signBytesWithNonce)

	btcecPrivKey, _ := btcec.PrivKeyFromBytes(btcec.S256(), privKey[:])

	signature, err := btcecPrivKey.Sign(msgHash)

	if err != nil {
		panic(err)
	}

	serSig := signature.Serialize()

	var ecSig SignatureSecp256k1

	ecSig = append(ecSig, serSig...)

	return ecSig

}

func (privKey PrivKeySecp256k1) String() string {
	return fmt.Sprintf("PrivKeySecp256k1{%X}", privKey[:])
}

func (privKey PrivKeySecp256k1) PubKey() PubKeySecp256k1 {
	_, pubK := btcec.PrivKeyFromBytes(btcec.S256(), privKey[:])

	serPubKey := pubK.SerializeUncompressed()

	var pubKey PubKeySecp256k1
	copy(pubKey[:], serPubKey[:65])

	return pubKey
}

func ConvertToPrivKey(privKey crypto.PrivKey) PrivKeySecp256k1 {
	privK, ok := privKey.(crypto.PrivKeySecp256k1)

	if !ok {
		panic("PrivateKey is not of type Secp256k1")
	}

	privateKey := NewPrivKeySecp256k1(privK[:])

	return privateKey
}
