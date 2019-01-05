package main

import (
	"github.com/sharering/shareledger/cmd/shareledger/subcommands"
)

func main() {
	rootCmd := subcommands.RootCmd

	rootCmd.AddCommand(
		subcommands.InitFilesCmd,
		subcommands.ShowNodeIDCmd,
		subcommands.ResetAllCmd,
		subcommands.NodeCmd,
		subcommands.VersionCmd,
		subcommands.TestnetFilesCmd,
		subcommands.ShowPrivKeyCmd,
		subcommands.RegisterValidatorCmd,
		subcommands.ShowAddressCmd,
		subcommands.ShowBalanceCmd,
		subcommands.ShowVdiCmd,
		subcommands.WithdrawBlockRewardCmd,
		subcommands.BeginUnbondingCmd,
		subcommands.CompleteUnbondingCmd,
	)

	rootCmd.Execute()
}
