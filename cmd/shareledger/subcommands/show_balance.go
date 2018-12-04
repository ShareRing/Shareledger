package subcommands

import (
	"fmt"

	"github.com/sharering/shareledger/client"
	"github.com/spf13/cobra"
)

var (
	nodeAddress string
)

var ShowBalanceCmd = &cobra.Command{
	Use:   "show_balance",
	Short: "show current balance of this masternode",
	RunE:  showBalance,
}

func init() {
	ShowBalanceCmd.Flags().StringVar(&nodeAddress, "client", "", "Node address to query info. Example: tcp://123.123.123.123:46657")
}

func showBalance(cmd *cobra.Command, args []string) (err error) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			err = fmt.Errorf("Error in balance retrieval")
		}
	}()

	var context client.CoreContext
	if nodeAddress == "" {

		context = client.NewCoreContextFromConfig(config)

	} else {

		context = client.NewCoreContextFromConfigWithClient(config, nodeAddress)
	}

	err = context.CheckBalance()
	if err != nil {
		return err
	}

	return nil

}
