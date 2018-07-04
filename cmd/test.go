package main

import (
	"fmt"
	"encoding/json"
	"encoding/hex"

	"github.com/sharering/shareledger/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func usingJson(){
	msg := types.MsgSend{
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

	var dmsg types.MsgSend
	nerr := json.Unmarshal(b, &dmsg)
	if nerr != nil {
		fmt.Println("Unmarsjal error:", err)
	}
	fmt.Println("Unmarshalled:", dmsg)
}

func usingCodec(){

	cdc := types.MakeCodec()

	msg := types.MsgSend{
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
	msg1 := types.MsgCheck{
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
	var a types.MsgCheck
	err = cdc.UnmarshalJSON(res1, &a)
	fmt.Println("Type:", a.Type())
	fmt.Println("Unmarshalled:", a)

	fmt.Println("********")
	var b sdk.Msg
	err = cdc.UnmarshalJSON(res1, &b)
	fmt.Println("Type:", b.Type())
	fmt.Println("Unmarshalled:", b)
}


func main(){
	usingCodec()
	fmt.Println("******")
	usingJson()
}
