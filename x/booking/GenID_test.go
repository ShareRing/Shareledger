package booking

import (
	"log"
	"testing"

	"github.com/ShareRing/Shareledger/x/booking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestSum(t *testing.T) {
	booker, err := sdk.AccAddressFromHex("23AE01BEE3B45C85AD85BF88AEADFC04B7C633CD")
	if err != nil {
		t.Fatal(err)
	}

	log.Println(booker.String())
	duration := int64(10)
	uuid := "1111"
	// msg := types.NewMsgBook(booker, uuid, duration)
	msg := types.MsgBook{
		Booker:   string(booker),
		UUID:     uuid,
		Duration: duration,
	}

	// msg.Duration = duration
	// msg.Booker = booker
	// msg.UUID = uuid

	bookId, err := GenBookID(msg)

	if err != nil {
		t.Fatal(err)
	}
	t.Log(bookId)
	t.Fail()
	// fmt.Println(bookId)
}
