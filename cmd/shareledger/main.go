package main

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sharering/shareledger/cmd/shareledger/subcommands"
	"github.com/sharering/shareledger/constants"
)

func main() {

	rootCmd := subcommands.RootCmd

	// configure SDK Cosmos
	configureCosmosSDK()
	

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
		subcommands.SendCoinCmd,
	)

	rootCmd.Execute()
}

func configureCosmosSDK() {
	//set bench32 prefix
	sdkConfig := sdk.GetConfig()
	sdkConfig.SetBech32PrefixForAccount(constants.Bech32PrefixAccAddr, constants.Bech32PrefixAccPub)
	sdkConfig.SetBech32PrefixForValidator(constants.Bech32PrefixValAddr, constants.Bech32PrefixValPub)
	sdkConfig.SetBech32PrefixForConsensusNode(constants.Bech32PrefixConsAddr, constants.Bech32PrefixConsPub)
	sdkConfig.Seal()
}
