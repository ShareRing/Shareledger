package types

import (
    "encoding/hex"
    "testing"
    //"fmt"

    crypto "github.com/tendermint/go-crypto"
	"github.com/btcsuite/btcd/btcec"

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
    var pubKey PubKeySecp256k1
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

    signBytes := msgCreate.GetSignBytes()
    t.Logf("SignBytes: %s\n", signBytes)
    messageHash := crypto.Sha256(signBytes)
    t.Logf("Hash: %x\n", messageHash)

    signature, err := privKey_.Sign(messageHash)
    if err != nil {
        t.Log("Error in Sign:", err)
        return
    }
    t.Logf("Raw Sig: %x %x\n", signature.R, signature.S)
    serSig := signature.Serialize()
    t.Logf("Signature: %x\n", serSig)

    var ecSig SignatureSecp256k1
    ecSig = append(ecSig, serSig...)

    shrSig := NewSHRSignature(pubKey, ecSig)
    t.Log("SHRSig:", shrSig)
    t.Log("Verify:", shrSig.Verify(signBytes))


    tx := NewSHRTx(msgCreate, shrSig)
    t.Log("Verify Signature:", tx.VerifySignature())

    t.Logf("Ox: %x\n", []byte("0x"))
    }
