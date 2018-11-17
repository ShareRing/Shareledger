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
	)

	rootCmd.Execute()
}
