package types

import (
	"bytes"
	"encoding/json"
	"fmt"

	secp256k1 "github.com/btcsuite/btcd/btcec"
	sha3 "github.com/ethereum/go-ethereum/crypto/sha3"

	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
	crypto "github.com/tendermint/go-crypto"
)

//----------------------------------------

type PubKey interface {
	Address() sdk.Address
	Bytes() []byte
	VerifyBytes(msg []byte, sig Signature) bool
	Equals(PubKey) bool
}

//----------------------------------------

var _ PubKey = PubKeySecp256k1{}

// Implements PubKey.
// Compressed pubkey (just the x-cord),
// prefixed with 0x02 or 0x03, depending on the y-cord.

// Normal key 65 byte
type PubKeySecp256k1 [65]byte

// Implements Bitcoin style addresses: RIPEMD160(SHA256(pubkey))
func (pubKey PubKeySecp256k1) Address() sdk.Address {
	hasherSHA256 := sha3.NewKeccak256()
	//hasherSHA256.Write([]byte("0x"))
	hasherSHA256.Write(pubKey[:]) // does not error
	var sha []byte
	sha = hasherSHA256.Sum(sha)
	return sdk.Address(sha[12:])
}

func (pubKey PubKeySecp256k1) Bytes() []byte {
	//cdc := amino.NewCodec()
	bz, err := json.Marshal(pubKey)
	if err != nil {
		panic(err)
	}
	return bz
}

func (pubKey PubKeySecp256k1) VerifyBytes(msg []byte, sig_ Signature) bool {
	// and assert same algorithm to sign and verify
	sig, ok := sig_.(SignatureSecp256k1)
	if !ok {
		fmt.Println("signature Is not Secp")
		return false
	}

	pub__, err := secp256k1.ParsePubKey(pubKey[:], secp256k1.S256())
	if err != nil {
		return false
	}

	sig__, err := secp256k1.ParseDERSignature(sig[:], secp256k1.S256())
	if err != nil {
		return false
	}
	return sig__.Verify(crypto.Sha256(msg), pub__)
}

func (pubKey PubKeySecp256k1) String() string {
	return fmt.Sprintf("PubKeySecp256k1{%X}", pubKey[:])
}

func (pubKey PubKeySecp256k1) Equals(other PubKey) bool {
	if otherSecp, ok := other.(PubKeySecp256k1); ok {
		return bytes.Equal(pubKey[:], otherSecp[:])
	} else {
		return false
	}
}
