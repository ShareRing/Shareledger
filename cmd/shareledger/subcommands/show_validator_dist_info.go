package subcommands

import (
	"fmt"

	"github.com/sharering/shareledger/client"
	"github.com/spf13/cobra"
)

var ShowVdiCmd = &cobra.Command{
	Use:   "show_val_dist_info",
	Short: "show current earning of this masternode",
	RunE:  showVdiCmd,
}

func init() {
	ShowVdiCmd.Flags().StringVar(&nodeAddress, "client", "", "Node address to query info. Example: tcp://127.0.0.1:46657")
}

func showVdiCmd(cmd *cobra.Command, args []string) (err error) {

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

	err = context.CheckValidatorDistInfo()
	if err != nil {
		return err
	}

	return nil

}
