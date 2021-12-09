package main

import (
	"os"

	"github.com/ShareRing/Shareledger/app"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
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
			NewImportKeyCmd(),
			getStakingCmd(),
			getSlashingCmd(),
			getDistributionCmd(),
			getFeegrantCmd(),
		),
	)
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
		NewCreateValidatorCmd(),
		NewEditValidatorCmd(),
		NewDelegateCmd(),
		NewRedelegateCmd(),
		NewUnbondCmd(),
	)

	return cmd
}

func getSlashingCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "slashing",
		Short: "Slashing transaction subcommands",
	}

	cmd.AddCommand(
		NewUnjailTxCmd(),
	)

	return cmd
}

func getDistributionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "distribution",
		Short: "Distribution transaction subcommands",
	}

	cmd.AddCommand(
		NewWithdrawRewardsCmd(),
		NewWithdrawAllRewardsCmd(),
	)

	return cmd
}

func getFeegrantCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "feegrant",
		Short: "Feegrant transaction subcommands",
	}

	cmd.AddCommand(
		NewCmdFeeGrant(),
		NewCmdRevokeFeegrant(),
	)

	return cmd
}
