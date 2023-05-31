package network

import (
	"fmt"

	"github.com/cometbft/cometbft/libs/cli"
	"github.com/cosmos/cosmos-sdk/client/flags"
)

var (
	JSONFlag         = fmt.Sprintf("--%s=%s", cli.OutputFlag, "json")
	SkipConfirmation = fmt.Sprintf("--%s", flags.FlagSkipConfirmation)
	SyncBroadcast    = fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync)
	GasAdjustment    = fmt.Sprintf("--%s=%s", flags.FlagGasAdjustment, "1.7")
	SHRFee2          = fmt.Sprintf("--%s=%s", flags.FlagFees, "2shr")
	SHRFee1          = fmt.Sprintf("--%s=%s", flags.FlagFees, "1shr")
	SHRFee10         = fmt.Sprintf("--%s=%s", flags.FlagFees, "10000000000nshr")
)

func SHRFee(number int) string {
	return fmt.Sprintf("--%s=%dshr", flags.FlagFees, number)
}

func AccountSeq(i int) string {
	return fmt.Sprintf("--%s=%d", flags.FlagSequence, i)
}

func MakeByAccount(k string) string {
	return fmt.Sprintf("--%s=%s", flags.FlagFrom, k)
}

func GetFlagsQuery() []string {
	return []string{
		fmt.Sprintf("--%s=%s", cli.OutputFlag, "json"),
	}
}
