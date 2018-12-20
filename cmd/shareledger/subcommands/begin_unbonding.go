package subcommands

import (
	"fmt"

	"github.com/sharering/shareledger/client"
	"github.com/spf13/cobra"
)

var BeginUnbondingCmd = &cobra.Command{
	Use:   "begin_unbonding",
	Short: "begin unbonding",
	RunE:  beginUnbonding,
}

var (
	unbondedTokens int64
)

func init() {
	BeginUnbondingCmd.Flags().StringVar(&nodeAddress, "client", "", "Node address to query info. Example: tcp://127.0.0.1:46657")
	BeginUnbondingCmd.Flags().Int64Var(&unbondedTokens, "tokens", 200000, "Amount of unbonded tokens.")
}

func beginUnbonding(cmd *cobra.Command, args []string) (err error) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			err = fmt.Errorf("Error in showing this masternode earning")
		}
	}()

	var context client.CoreContext
	if nodeAddress == "" {

		context = client.NewCoreContextFromConfig(config)

	} else {

		context = client.NewCoreContextFromConfigWithClient(config, nodeAddress)
	}

	err = context.BeginUnbonding(unbondedTokens)
	if err != nil {
		return err
	}

	return nil

}
