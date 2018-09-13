package main

import (
	"encoding/hex"
	"fmt"
	"sort"

	"github.com/btcsuite/btcd/btcec"

	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
	"bitbucket.org/shareringvn/cosmos-sdk/wire"

	crypto "github.com/tendermint/go-crypto"

	"github.com/sharering/shareledger/x/auth"

	"github.com/sharering/shareledger/app"
	"github.com/sharering/shareledger/types"

	"github.com/sharering/shareledger/x/asset"
	amsg "github.com/sharering/shareledger/x/asset/messages"

	"github.com/sharering/shareledger/x/bank"
	"github.com/sharering/shareledger/x/bank/messages"

	"github.com/sharering/shareledger/x/booking"
	bmsg "github.com/sharering/shareledger/x/booking/messages"
)

func testMessage(cdc *wire.Codec, message sdk.Msg) (int, int) {
	jsonRes, err1 := cdc.MarshalJSON(message)
	aminoRes, err2 := cdc.MarshalBinary(message)
	if err1 != nil || err2 != nil {
		return -1, -1
	}
	fmt.Printf("\nJSON-ENCODED:\n%s\n", jsonRes)
	fmt.Printf("\nAMINO-ENCODED:\n0x%v\n", hex.EncodeToString(aminoRes))

	return len(jsonRes), len(aminoRes)
}

func testSignedTx(cdc *wire.Codec, signedTx types.SHRTx) (int, int) {
	jsonRes, err1 := cdc.MarshalJSON(signedTx)
	aminoRes, err2 := cdc.MarshalBinary(signedTx)
	if err1 != nil || err2 != nil {
		return -1, -1
	}
	fmt.Printf("\nJSON-ENCODED:\n%s\n", jsonRes)
	fmt.Printf("\nAMINO-ENCODED:\n0x%v\n", hex.EncodeToString(aminoRes))

	return len(jsonRes), len(aminoRes)
}

func getKeyPair() (*btcec.PrivateKey, types.PubKeySecp256k1) {
	pkBytes, _ := hex.DecodeString("ab83994cf95abe45b9d8610524b3f8f8fd023d69f79449011cb5320d2ca180c5")

	privKey_, pubKey_ := btcec.PrivKeyFromBytes(btcec.S256(), pkBytes)

	serPubKey := pubKey_.SerializeUncompressed()
	var pubKey types.PubKeySecp256k1
	copy(pubKey[:], serPubKey[:65])
	return privKey_, pubKey
}

func signMessage(msg sdk.Msg) (tx auth.AuthTx) {

	nonce := 4

	privKey_, pubKey := getKeyPair()

	address := pubKey.Address()
	fmt.Printf("Address: %x\n", address)

	signBytes := msg.GetSignBytes()
	msgHash := crypto.Sha256(signBytes)

	signature, _ := privKey_.Sign(msgHash)
	serSig := signature.Serialize()

	var ecSig types.SignatureSecp256k1
	ecSig = append(ecSig, serSig...)

	shrSig := auth.NewAuthSig(pubKey, ecSig, int64(nonce))

	tx = auth.NewAuthTx(msg, shrSig)
	return tx
}

func main() {

	cdc := app.MakeCodec()
	cdc = asset.RegisterCodec(cdc)
	cdc = bank.RegisterCodec(cdc)
	cdc = booking.RegisterCodec(cdc)

	_, pubKey := getKeyPair()
	address := pubKey.Address()
	//fmt.Println("Address Length:", len(address))

	msgSend := messages.MsgSend{
		To: address,
		Amount: types.Coin{
			Denom:  "SHRP",
			Amount: 1,
		},
	}

	msgLoad := messages.MsgLoad{
		Account: address,
		Amount:  types.Coin{"SHR", 100},
	}

	msgCreate := amsg.MsgCreate{
		Creator: address,
		Hash:    []byte("333333333"),
		UUID:    "33333333",
		Fee:     1,
		Status:  true,
	}

	msgRetrieve := amsg.MsgRetrieve{
		UUID: "33333333",
	}

	msgUpdate := amsg.MsgUpdate{
		Creator: address,
		Hash:    []byte("333333333"),
		UUID:    "33333333",
		Fee:     1,
		Status:  true,
	}

	msgDelete := amsg.MsgDelete{
		UUID: "33333333",
	}

	msgBook := bmsg.MsgBook{
		UUID:     "112233",
		Duration: 12,
	}

	msgComplete := bmsg.MsgComplete{
		BookingID: "112233",
	}

	msgs := map[string]sdk.Msg{
		"MsgSend":     msgSend,
		"MsgLoad":     msgLoad,
		"MsgCreate":   msgCreate,
		"MsgRetrieve": msgRetrieve,
		"MsgUpdate":   msgUpdate,
		"MsgDelete":   msgDelete,
		"MsgBook":     msgBook,
		"MsgComplete": msgComplete,
	}

	// Sort key
	var keys []string
	for k, _ := range msgs {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	// perform Amino-encoded and record result
	result := map[string][2]int{}
	for _, name := range keys {

		msg := msgs[name]

		fmt.Println("MESSAGE: ", name)
		r1, r2 := testMessage(cdc, msg)

		signedTx := signMessage(msg)
		fmt.Println("\nSIGNED MESSAGE:", signedTx)

		r3, r4 := testSignedTx(cdc, signedTx)

		result[name] = [2]int{r1, r2}
		result[fmt.Sprintf("Signed %s", name)] = [2]int{r3, r4}

		fmt.Println("\n")
	}

	fmt.Printf("%20s  %10s  %10s  %10s\n", "Messages", "JSON-encoded", "Amino-encoded", "Percentages")
	total := 0
	var sum float64
	for _, name := range keys {
		r := result[name]
		ratio := float64(r[0]) * 100 / float64(r[1])
		fmt.Printf("%20s  %10d  %10d  %12.2f%%\n", name, r[0], r[1], ratio)
		signedName := fmt.Sprintf("Signed %s", name)
		total += 1
		sum += ratio

		r = result[signedName]
		ratio = float64(r[0]) * 100 / float64(r[1])
		fmt.Printf("%20s  %10d  %10d  %12.2f%%\n", signedName, r[0], r[1], ratio)
		total += 1
		sum += ratio
	}
	fmt.Printf("\nTOTAL: %d AVERAGE: %.2f%%\n", total, sum/float64(total))
}
