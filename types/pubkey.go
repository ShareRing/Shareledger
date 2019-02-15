package types

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/btcsuite/btcd/btcec"
	// sha3 "github.com/ethereum/go-ethereum/crypto/sha3"
	sha3 "golang.org/x/crypto/sha3"

	sdk "github.com/cosmos/cosmos-sdk/types"
	crypto "github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/secp256k1"

	// Library for test
	"github.com/landonia/crypto/bip39"
)

var (
	BITS = 256
	//ADDRESSLENGTH = 20 //byte length
)

const (
	ADDRESSLENGTH = 20 //byte length
)

//----------------------------------------

type PubKey interface {
	Address() sdk.AccAddress
	Bytes() []byte
	VerifyBytes(msg []byte, sig Signature) bool
	Equals(PubKey) bool
	String() string
}

//----------------------------------------

var _ PubKey = PubKeySecp256k1{}

// Implements PubKey.
// Compressed pubkey (just the x-cord),
// prefixed with 0x02 or 0x03, depending on the y-cord.

// Normal key 65 byte
type PubKeySecp256k1 [65]byte

func NewPubKeySecp256k1(b []byte) PubKeySecp256k1 {
	if len(b) != 65 {
		panic("Length of input to create PubKeySecp256k1 should be 65")
	}
	var pubK PubKeySecp256k1

	copy(pubK[:], b[:65])

	return pubK
}

func NilPubKeySecp256k1() PubKeySecp256k1 {
	return NewPubKeySecp256k1(make([]byte, 65))
}

// Implements Bitcoin style addresses: RIPEMD160(SHA256(pubkey))
func (pubKey PubKeySecp256k1) Address() sdk.AccAddress {
	hasherSHA256 := sha3.NewLegacyKeccak256()
	//hasherSHA256.Write([]byte("0x"))
	hasherSHA256.Write(pubKey[:]) // does not error
	var sha []byte
	sha = hasherSHA256.Sum(sha)
	return sdk.AccAddress(sha[12:])
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
		panic("Signature is not of tupe SignatureSecp256k1")
		return false
	}

	pub__, err := btcec.ParsePubKey(pubKey[:], btcec.S256())
	if err != nil {
		return false
	}

	sig__, err := btcec.ParseDERSignature(sig[:], btcec.S256())
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

func (pubKey PubKeySecp256k1) ToABCIPubKey() secp256k1.PubKeySecp256k1 {
	var pk secp256k1.PubKeySecp256k1

	pub__, err := btcec.ParsePubKey(pubKey[:], btcec.S256())

	if err != nil {
		panic(err)
	}

	copy(pk[:], pub__.SerializeCompressed())
	return pk
}

//------------------------------------------------------------

func GetTestPubKey() PubKeySecp256k1 {
	pkBytes, err := hex.DecodeString("ab83994cf95abe45b9d8610524b3f8f8fd023d69f79449011cb5320d2ca180c5")

	if err != nil {
		panic("Error in DecodeString")
	}

	_, pubKey_ := btcec.PrivKeyFromBytes(btcec.S256(), pkBytes)

	serPubKey := pubKey_.SerializeUncompressed()
	var pubKey PubKeySecp256k1
	copy(pubKey[:], serPubKey[:65])

	//address := pubKey.Address()
	return pubKey
}

func GenerateKeyPair() (PubKeySecp256k1, PrivKeySecp256k1) {

	entropy, err := bip39.GenerateRandomEntropy(BITS)

	if err != nil {
		panic(err)
	}

	mnemonics, err1 := entropy.GenerateMnemonics(bip39.English)

	if err1 != nil {
		panic(err1)
	}

	privKeyBytes := crypto.Sha256([]byte(mnemonics.String())) // 32 bytes

	_, pubK := btcec.PrivKeyFromBytes(btcec.S256(), privKeyBytes)

	serPubKey := pubK.SerializeUncompressed()

	var pubKey PubKeySecp256k1
	copy(pubKey[:], serPubKey[:65])

	var privKey PrivKeySecp256k1
	copy(privKey[:], privKeyBytes[:32])

	return pubKey, privKey

}

// convert from Tendermint PubKey to ShareLedger PubKey
func ConvertToPubKey(pubKey []byte) PubKeySecp256k1 {

	// Convert to Tendermint PubKeySecp256k1
	// Tendermint use compressed version
	// We need to convert to non-compressed version
	// tmPubKey, ok := pubKey.(crypto.PubKeySecp256k1)
	// if !ok {
	// panic("Key is not of type crypto.PubKeySecp156k1")
	// }

	// Convert to PubKey in btcec
	// btPubKey, err := btcec.ParsePubKey(tmPubKey[:], btcec.S256())
	btPubKey, err := btcec.ParsePubKey(pubKey[:], btcec.S256())

	if err != nil {
		panic("Cannot parse PubKey in Secp256k1")
	}

	return NewPubKeySecp256k1(btPubKey.SerializeUncompressed())
}
