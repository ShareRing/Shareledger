package main

import (
	"fmt"
	"encoding/json"
	"encoding/hex"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sharering/shareledger/x/bank/messages"
	"github.com/sharering/shareledger/x/bank"
	"github.com/sharering/shareledger/types"
	"github.com/cosmos/cosmos-sdk/wire"
)

func usingJson(){
	msg := messages.MsgSend{
		From: sdk.Address([]byte("123")),
		To: sdk.Address([]byte("234")),
		Amount: types.Coin{
			Denom: "SHR",
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

func usingCodec(){

	cdc := bank.MakeCodec()

	msg := messages.MsgSend{
		From: sdk.Address([]byte("123")),
		To: sdk.Address([]byte("234")),
		Amount: types.Coin{
			Denom: "SHR",
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
		Denom: "SHR",
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



	printMsgLoad(cdc)
}


func printMsgLoad(cdc *wire.Codec){
	fmt.Println("*****MsgLoad")
	msg1 := messages.MsgLoad{
		Nonce: 1,
		Account: sdk.Address([]byte("123")),
		Amount: types.Coin{"SHR", 100},
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


func main(){
	usingCodec()
	fmt.Println("******")
	usingJson()
}
