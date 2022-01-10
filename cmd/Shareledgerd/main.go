package main

import (
	"os"

	"github.com/sharering/shareledger/cmd/Shareledgerd/sub"
	"github.com/sharering/shareledger/cmd/Shareledgerd/tools"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/sharering/shareledger/app"
	"github.com/spf13/cobra"
	"github.com/tendermint/spm/cosmoscmd"
)

func main() {
	rootCmd, _ := cosmoscmd.NewRootCmd(
		app.Name,
		app.AccountAddressPrefix,
		app.DefaultNodeHome,
		app.Name,
		app.ModuleBasics,
		app.New,
		cosmoscmd.AddSubCmd(
			sub.NewImportKeyCmd(),
			getGenesisCmd(app.DefaultNodeHome),
		),
	)

	//fmt.Println("PreRun E", rootCmd.PreRunE)

	//for _, c := range rootCmd.Commands() {
	//	if c.Name() == "query" {
	//		for _, moduleCommands := range c.Commands() {
	//			subTxCommands := moduleCommands.Commands()
	//			for i, _ := range subTxCommands {
	//				fmt.Println("subTxCommands", subTxCommands[i].Name())
	//				subTxCommands[i].PreRunE = func(cmd *cobra.Command, args []string) error {
	//					fmt.Println("--- IN PreRunE ----")
	//					return nil
	//				}
	//			}
	//			//fmt.Println("subTxCommands", subTxCommands[i].Name(), subTxCommands[i].PreRunE)
	//		}
	//	}
	//}

	//rootCmd.PreRunE = func(cmd *cobra.Command, args []string) error {
	//
	//	fmt.Println("IN PreRunE", cmd.UsageString())
	//	return nil
	//}

	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}

func getGenesisCmd(defaultNodeHome string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "genesis",
		Short: "genesis file subcommands",
	}
	cmd.AddCommand(
		tools.NewGenesisAddAuthorityAccountCmd(defaultNodeHome),
		tools.NewGenesisAddTreasureAccountCmd(defaultNodeHome),
		tools.NewGenesisAddValidatorAccountCmd(defaultNodeHome),
		tools.NewGenesisAddAccountOperatorCmd(defaultNodeHome),
	)
	return cmd
}
