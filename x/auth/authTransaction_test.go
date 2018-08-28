package auth

import (
	"encoding/hex"
	"strconv"
	"testing"
	//"fmt"

	"github.com/btcsuite/btcd/btcec"
	crypto "github.com/tendermint/go-crypto"

	"github.com/sharering/shareledger/types"
	"github.com/sharering/shareledger/x/asset/messages"
)

func TestTransaction(t *testing.T) {
	pkBytes, err := hex.DecodeString("ab83994cf95abe45b9d8610524b3f8f8fd023d69f79449011cb5320d2ca180c5")

	if err != nil {
		t.Log("Error in DecodeString: ", err)
		return
	}

	privKey_, pubKey_ := btcec.PrivKeyFromBytes(btcec.S256(), pkBytes)

	serPubKey := pubKey_.SerializeUncompressed()
	var pubKey types.PubKeySecp256k1
	t.Logf("PubKey: %x\n", serPubKey)
	copy(pubKey[:], serPubKey[:65])

	address := pubKey.Address()
	t.Logf("Address: %x\n", address)
	uuid := "112233"
	status := true
	fee := int64(1)
	hash := []byte("111111")

	msgCreate := messages.NewMsgCreate(
		address,
		hash,
		uuid,
		status,
		fee,
	)

	nonce := 1

	// Signing
	signBytes := msgCreate.GetSignBytes()
	t.Logf("SignBytes: %s\n", signBytes)

	signBytesWithNonce := append([]byte(strconv.Itoa(nonce)), signBytes...)
	t.Logf("SignBytes with Nonce: %s\n", signBytesWithNonce)

	messageHash := crypto.Sha256(signBytesWithNonce)
	t.Logf("Hash: %x\n", messageHash)

	signature, err := privKey_.Sign(messageHash)
	if err != nil {
		t.Log("Error in Sign:", err)
		return
	}
	serSig := signature.Serialize()
	t.Logf("Signature: %x\n", serSig)

	var ecSig types.SignatureSecp256k1
	ecSig = append(ecSig, serSig...)

	shrSig := NewAuthSig(pubKey, ecSig, int64(nonce))
	t.Log("AuthSig:", shrSig)
	t.Log("Verify:", shrSig.Verify(signBytesWithNonce))

	if !shrSig.Verify(signBytes) {
		t.Error("Signature verification failed.")
	}

	tx := NewAuthTx(msgCreate, shrSig)
	t.Logf("Nonce Signed Tx: %s\n", tx)
	t.Log("Verify Signature:", tx.VerifySignature())

	if !tx.VerifySignature() {
		t.Error("Signature verification failed.")

	}

}
