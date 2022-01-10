package main

import (
	"os"

	"github.com/sharering/shareledger/cmd/Shareledgerd/sub"
	"github.com/sharering/shareledger/cmd/Shareledgerd/tools"

	"github.com/cosmos/cosmos-sdk/client/flags"
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
			getStakingCmd(),
			getSlashingCmd(),
			getDistributionCmd(),
			getFeegrantCmd(),
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

func getStakingCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "staking",
		Short: "Staking transaction subcommands",
	}

	cmd.AddCommand(
		sub.NewCreateValidatorCmd(),
		sub.NewEditValidatorCmd(),
		sub.NewDelegateCmd(),
		sub.NewRedelegateCmd(),
		sub.NewUnbondCmd(),
	)
	cmd.PersistentFlags().String(flags.FlagChainID, "", "The network chain ID")

	return cmd
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

func getSlashingCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "slashing",
		Short: "Slashing transaction subcommands",
	}

	cmd.AddCommand(
		sub.NewUnjailTxCmd(),
	)
	cmd.PersistentFlags().String(flags.FlagChainID, "", "The network chain ID")

	return cmd
}

func getDistributionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "distribution",
		Short: "Distribution transaction subcommands",
	}

	cmd.AddCommand(
		sub.NewWithdrawRewardsCmd(),
		sub.NewWithdrawAllRewardsCmd(),
	)
	cmd.PersistentFlags().String(flags.FlagChainID, "", "The network chain ID")

	return cmd
}

func getFeegrantCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "feegrant",
		Short: "Feegrant transaction subcommands",
	}

	cmd.AddCommand(
		sub.NewCmdFeeGrant(),
		sub.NewCmdRevokeFeegrant(),
	)
	cmd.PersistentFlags().String(flags.FlagChainID, "", "The network chain ID")

	return cmd
}
