package network

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/tendermint/tendermint/libs/cli"
)

var (
	CommonFeeFlags          = "--fee 1shr"
	CommonResponseTypeFlags = "--output=json"
)

func GetDefaultFlags(txCreator string) []string {
	return []string{
		fmt.Sprintf("--%s", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagFrom, txCreator),
		fmt.Sprintf("--%s=%s", cli.OutputFlag, "json"),
		fmt.Sprintf("--%s=%s", flags.FlagFees, "1shr"),
	}
}

func GetFlagsQuery() []string {
	return []string{
		fmt.Sprintf("--%s=%s", cli.OutputFlag, "json"),
	}
}

func GetDefaultFlags2SHR(txCreator string) []string {
	return []string{
		fmt.Sprintf("--%s", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagFrom, txCreator),
		fmt.Sprintf("--%s=%s", cli.OutputFlag, "json"),
		fmt.Sprintf("--%s=%s", flags.FlagFees, "2shr"),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
	}
}
