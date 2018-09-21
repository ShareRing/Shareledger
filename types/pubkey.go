package types

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"

	secp256k1 "github.com/btcsuite/btcd/btcec"
	sha3 "github.com/ethereum/go-ethereum/crypto/sha3"

	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
	crypto "github.com/tendermint/go-crypto"

	// Library for test
	"github.com/landonia/crypto/bip39"
)

var (
	BITS = 256
)

//----------------------------------------

type PubKey interface {
	Address() sdk.Address
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

//------------------------------------------------------------

func GetTestPubKey() PubKeySecp256k1 {
	pkBytes, err := hex.DecodeString("ab83994cf95abe45b9d8610524b3f8f8fd023d69f79449011cb5320d2ca180c5")

	if err != nil {
		fmt.Println("Error in DecodeString: ", err)
		panic("Error in DecodeString")
	}

	_, pubKey_ := secp256k1.PrivKeyFromBytes(secp256k1.S256(), pkBytes)

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

	_, pubK := secp256k1.PrivKeyFromBytes(secp256k1.S256(), privKeyBytes)

	serPubKey := pubK.SerializeUncompressed()

	var pubKey PubKeySecp256k1
	copy(pubKey[:], serPubKey[:65])

	var privKey PrivKeySecp256k1
	copy(privKey[:], privKeyBytes[:32])

	return pubKey, privKey

}
