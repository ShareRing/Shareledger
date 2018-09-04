package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"

	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
	"github.com/btcsuite/btcd/btcec"
	crypto "github.com/tendermint/go-crypto"

	"github.com/sharering/shareledger/app"
	"github.com/sharering/shareledger/types"
	"github.com/sharering/shareledger/x/bank"
	"github.com/sharering/shareledger/x/bank/messages"

	"github.com/sharering/shareledger/x/asset"
	amsg "github.com/sharering/shareledger/x/asset/messages"

	//ahdl "github.com/sharering/shareledger/x/asset/handlers"

	"github.com/sharering/shareledger/x/booking"
	bmsg "github.com/sharering/shareledger/x/booking/messages"

	sha3 "github.com/ethereum/go-ethereum/crypto/sha3"
	"github.com/sharering/shareledger/x/auth"
)

func usingJson() {
	msg := messages.MsgSend{
		To: sdk.Address([]byte("234")),
		Amount: types.Coin{
			Denom:  "SHR",
			Amount: 3,
		},
	}
	fmt.Println("Message:", msg)
	b, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Marshalled:", b)
	fmt.Printf("String format: %s\n", b)
	fmt.Printf("ToString: %s\n", hex.EncodeToString(b))

	var dmsg messages.MsgSend
	nerr := json.Unmarshal(b, &dmsg)
	if nerr != nil {
		fmt.Println("Unmarsjal error:", err)
	}
	fmt.Println("Unmarshalled:", dmsg)
}

func usingCodec() {

	cdc := app.MakeCodec()
	cdc = bank.RegisterCodec(cdc)

	msg := messages.MsgSend{
		To: sdk.Address([]byte("234")),
		Amount: types.Coin{
			Denom:  "SHR",
			Amount: 3,
		},
	}
	fmt.Println("Message:", msg)

	res, err := cdc.MarshalJSON(msg)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println("Marshalled:", res)
	fmt.Printf("String format: %s\n", res)
	fmt.Printf("ToString: %s\n", hex.EncodeToString(res))

	fmt.Println("*****")
	msg1 := messages.MsgCheck{
		Account: sdk.Address([]byte("123")),
	}

	fmt.Println("Check Message:", msg1)
	res1, err1 := cdc.MarshalJSON(msg1)
	if err1 != nil {
		fmt.Println("Error", err1)
		return
	}

	fmt.Println("Marshalled:", res1)
	fmt.Printf("String format: %s\n", res1)
	fmt.Printf("ToString: %s\n", hex.EncodeToString(res1))

	fmt.Println("********")
	var a messages.MsgCheck
	err = cdc.UnmarshalJSON(res1, &a)
	fmt.Println("Type:", a.Type())
	fmt.Println("Unmarshalled:", a)

	fmt.Println("********")
	var b sdk.Msg
	err = cdc.UnmarshalJSON(res1, &b)
	fmt.Println("Type:", b.Type())
	fmt.Println("Unmarshalled:", b)

	printMsgLoad()
}

func printMsgLoad() {
	cdc := app.MakeCodec()
	cdc = bank.RegisterCodec(cdc)

	fmt.Println("*****MsgLoad")
	msg1 := messages.MsgLoad{
		Account: sdk.Address([]byte("123")),
		Amount:  types.Coin{"SHR", 100},
	}

	fmt.Println("Load Message:", msg1)
	res1, err1 := cdc.MarshalJSON(msg1)
	if err1 != nil {
		fmt.Println("Error", err1)
		return
	}

	fmt.Println("Marshalled:", res1)
	fmt.Printf("String format: %s\n", res1)
	fmt.Printf("ToString: %s\n", hex.EncodeToString(res1))

}

func printMsgCreate() {
	cdc := app.MakeCodec()
	cdc = asset.RegisterCodec(cdc)

	msg := amsg.MsgCreate{
		Creator: sdk.Address([]byte("333333")),
		Hash:    []byte("333333"),
		UUID:    "333333",
		Fee:     1,
		Status:  true,
	}
	//msg := amsg.MsgDelete{
	//	UUID: "333333",
	//}
	//msg := amsg.MsgUpdate{
	//	Creator: sdk.Address([]byte("333333")),
	//	Hash: []byte("333333"),
	//	UUID: "333333",
	//}
	res, err := cdc.MarshalJSON(msg)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println("Marshalled:", res)
	fmt.Printf("String format: %s\n", res)
	fmt.Printf("ToString: %s\n", hex.EncodeToString(res))

}

func printMsgBook() {
	cdc := app.MakeCodec()
	cdc = booking.RegisterCodec(cdc)

	msg := bmsg.MsgBook{
		UUID:     "112233",
		Duration: 12,
	}

	res, err := cdc.MarshalJSON(msg)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println("Marshalled:", res)
	fmt.Printf("String format: %s\n", res)
	fmt.Printf("ToString: %s\n", hex.EncodeToString(res))
}

func printMsgComplete() {
	cdc := app.MakeCodec()
	cdc = booking.RegisterCodec(cdc)

	msg := bmsg.MsgComplete{
		BookingID: "121212",
	}

	res, err := cdc.MarshalJSON(msg)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println("Marshalled:", res)
	fmt.Printf("String format: %s\n", res)
	fmt.Printf("ToString: %s\n", hex.EncodeToString(res))
}

func testMsgCreateSignedTx() {
	pkBytes, err := hex.DecodeString("ab83994cf95abe45b9d8610524b3f8f8fd023d69f79449011cb5320d2ca180c5")

	if err != nil {
		fmt.Println("Error in DecodeString: ", err)
		return
	}

	privKey_, pubKey_ := btcec.PrivKeyFromBytes(btcec.S256(), pkBytes)

	serPubKey := pubKey_.SerializeUncompressed()
	var pubKey types.PubKeySecp256k1
	copy(pubKey[:], serPubKey[:65])
	fmt.Printf("SerPubKey: %x\n", serPubKey)
	fmt.Printf("PubKey: %x\n", serPubKey[:65])

	address := pubKey.Address()
	fmt.Printf("Address: %x\n", address)
	uuid := "112233"
	status := true
	fee := int64(1)
	hash := []byte("111111")

	msgCreate := amsg.NewMsgCreate(
		address,
		hash,
		uuid,
		status,
		fee,
	)

	signBytes := msgCreate.GetSignBytes()
	fmt.Printf("SignBytes: %s\n", signBytes)
	messageHash := crypto.Sha256(signBytes)
	fmt.Printf("Hash: %x\n", messageHash)

	signature, err := privKey_.Sign(messageHash)
	if err != nil {
		fmt.Println("Error in Sign:", err)
		return
	}
	fmt.Printf("Raw Sig: %x %x\n", signature.R, signature.S)
	serSig := signature.Serialize()
	fmt.Printf("Signature: %x\n", serSig)

	var ecSig types.SignatureSecp256k1
	ecSig = append(ecSig, serSig...)

	shrSig := types.NewBasicSig(pubKey, ecSig)
	fmt.Println("Verify:", shrSig.Verify(signBytes))

	tx := types.NewBasicTx(msgCreate, shrSig)
	fmt.Println("Verify Signature:", tx.VerifySignature())

	cdc := app.MakeCodec()
	cdc = asset.RegisterCodec(cdc)
	b, e := cdc.MarshalJSON(tx)
	if e != nil {
		fmt.Println("Error in encoding tx", e)
	}
	fmt.Printf("Tx encoded: %s\n", b)

	bA, eA := cdc.MarshalJSON(serSig)
	if eA != nil {
		fmt.Println("Error:", eA)
	}
	fmt.Printf("SerializedSig: %s\n", bA)

}

func testMsgCreateNonceSignedTx() {
	pkBytes, err := hex.DecodeString("ab83994cf95abe45b9d8610524b3f8f8fd023d69f79449011cb5320d2ca180c5")

	if err != nil {
		fmt.Println("Error in DecodeString: ", err)
		return
	}

	privKey_, pubKey_ := btcec.PrivKeyFromBytes(btcec.S256(), pkBytes)

	serPubKey := pubKey_.SerializeUncompressed()
	var pubKey types.PubKeySecp256k1
	copy(pubKey[:], serPubKey[:65])
	fmt.Printf("SerPubKey: %x\n", serPubKey)
	fmt.Printf("PubKey: %x\n", serPubKey[:65])

	address := pubKey.Address()
	fmt.Printf("Address: %x\n", address)
	uuid := "112233"
	status := true
	fee := int64(1)
	hash := []byte("111111")

	msgCreate := amsg.NewMsgCreate(
		address,
		hash,
		uuid,
		status,
		fee,
	)

	nonce := 1

	signBytes := msgCreate.GetSignBytes()
	fmt.Printf("SignBytes: %s\n", signBytes)

	signBytesWithNonce := append([]byte(strconv.Itoa(nonce)), signBytes...)
	fmt.Printf("SignBytesWithNonce: %s", signBytesWithNonce)

	messageHash := crypto.Sha256(signBytesWithNonce)
	fmt.Printf("Hash: %x\n", messageHash)

	signature, err := privKey_.Sign(messageHash)
	if err != nil {
		fmt.Println("Error in Sign:", err)
		return
	}

	fmt.Printf("Raw Sig: %x %x\n", signature.R, signature.S)
	serSig := signature.Serialize()
	fmt.Printf("Signature: %x\n", serSig)

	var ecSig types.SignatureSecp256k1
	ecSig = append(ecSig, serSig...)

	shrSig := auth.NewAuthSig(pubKey, ecSig, int64(nonce))
	fmt.Println("Verify:", shrSig.Verify(signBytes))

	tx := auth.NewAuthTx(msgCreate, shrSig)
	fmt.Println("Verify Signature:", tx.VerifySignature())

	cdc := app.MakeCodec()
	cdc = asset.RegisterCodec(cdc)
	b, e := cdc.MarshalJSON(tx)
	if e != nil {
		fmt.Println("Error in encoding tx", e)
	}
	fmt.Printf("Tx encoded: %s\n", b)

	bA, eA := cdc.MarshalJSON(serSig)
	if eA != nil {
		fmt.Println("Error:", eA)
	}
	fmt.Printf("SerializedSig: %s\n", bA)

}

func testPubKey() {
	keyStr := "2b4d507ba002c4fa3d6394dbd911bb2ae49d2dbc70e0c80e368e967d409e87dbd8d95b0d50c0fe95410687e9679e334f41efb690853e061c14c00704f8901afc"
	keyByte, err := hex.DecodeString(keyStr)
	if err != nil {
		fmt.Println("Error", err)
	}
	var keySec types.PubKeySecp256k1
	copy(keySec[:], keyByte[:])
	fmt.Printf("Address: %x\n", keySec.Address())

}

func testKeccak() {
	// str := []byte("2b4d507ba002c4fa3d6394dbd911bb2ae49d2dbc70e0c80e368e967d409e87dbd8d95b0d50c0fe95410687e9679e334f41efb690853e061c14c00704f8901afc")
	str := "2b4d507ba002c4fa3d6394dbd911bb2ae49d2dbc70e0c80e368e967d409e87dbd8d95b0d50c0fe95410687e9679e334f41efb690853e061c14c00704f8901afc"
	hash := sha3.NewKeccak256()

	var buf []byte
	keyByte := decodeHex(str)
	var keySec types.PubKeySecp256k1
	copy(keySec[:], keyByte[:])
	hash.Write(keySec[:])
	// hash.Write(str)
	buf = hash.Sum(buf)
	res := fmt.Sprintf("%x", buf)
	fmt.Println(res)
	fmt.Println("Length", len(res))
}

func decodeHex(s string) []byte {
	b, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}

	return b
}

func testSignature() {

}

func main() {
	//usingCodec()
	printMsgCreate()
	//printMsgBook()
	//printMsgLoad()
	//printMsgComplete()
	//testPubKey()
	// testKeccak()
	//testMsgCreateSignedTx()
	//testMsgCreateNonceSignedTx()

}
