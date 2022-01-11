package fee

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"reflect"
	"testing"
)

func TestTableFee_GetNameMsg(t *testing.T) {
	a := func(msg sdk.Msg) bool {
		fmt.Println(reflect.ValueOf(&distributiontypes.MsgWithdrawDelegatorReward{}).Type().String())
		return reflect.ValueOf(&distributiontypes.MsgWithdrawDelegatorReward{}).Type() == reflect.TypeOf(msg)
	}
	msg := &distributiontypes.MsgWithdrawDelegatorReward{}
	fmt.Println(a(msg))

}
