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

func JSONFlag()string{
	return fmt.Sprintf("--%s=%s", cli.OutputFlag, "json")
}


func SkipConfirmation()string{
	return fmt.Sprintf("--%s", flags.FlagSkipConfirmation)
}

func SHRFee1()string{
	return fmt.Sprintf("--%s=%s", flags.FlagFees, "1shr")
}

func SHRFee2()string{
	return fmt.Sprintf("--%s=%s", flags.FlagFees, "2shr")
}
func SHRFee3()string{
	return fmt.Sprintf("--%s=%s", flags.FlagFees, "3shr")
}

func BlockBroadcast()string{
	return fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock)
}


func MakeByAccount(k string)string{
	return fmt.Sprintf("--%s=%s", flags.FlagFrom, k)
}



func GetDefaultFlags(txCreator string) []string {
	return []string{
		SkipConfirmation(),
		MakeByAccount(txCreator),
		JSONFlag(),
		SHRFee1(),
		BlockBroadcast(),
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
