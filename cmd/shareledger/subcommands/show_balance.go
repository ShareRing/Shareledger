package subcommands

import (
	"fmt"

	"github.com/sharering/shareledger/client"
	"github.com/spf13/cobra"
)

var ShowBalanceCmd = &cobra.Command{
	Use:   "show_balance",
	Short: "show current balance of this masternode",
	RunE:  showBalance,
}

func showBalance(cmd *cobra.Command, args []string) (err error) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			err = fmt.Errorf("Error in balance show.")
		}
	}()

	context := client.NewCoreContextFromConfig(config)

	err = context.CheckBalance()
	if err != nil {
		return err
	}

	return nil

}
