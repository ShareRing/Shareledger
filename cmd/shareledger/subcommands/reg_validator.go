package subcommands

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/sharering/shareledger/client"
)

const (
	MINIMUN_STAKED int64 = int64(2000)
)

var (
	moniker  string
	identity string
	website  string
	details  string
	amount   int64
)

// ShowNodeIDCmd dumps node's ID to the standard output.
var RegisterValidatorCmd = &cobra.Command{
	Use:   "register_masternode",
	Short: "register this node as a masternode",
	RunE:  registerValidator,
}

func init() {
	RegisterValidatorCmd.Flags().StringVar(&moniker, "moniker", "", "Node name")
	// RegisterValidatorCmd.Flags().StringVar(&identity, "identity", "", "Identity Signature (ex: uPort or Keybase)")
	RegisterValidatorCmd.Flags().StringVar(&website, "website", "sharering.network", "Website link")
	RegisterValidatorCmd.Flags().StringVar(&details, "details", "ShareLedger Masternode", "Details of your MasterNode")
	RegisterValidatorCmd.Flags().Int64Var(&amount, "tokens", 0, "Amount of tokens to be staked.")
	RegisterValidatorCmd.MarkFlagRequired("tokens")
	RegisterValidatorCmd.Flags().StringVar(&nodeAddress, "client", "", "Node address to query info. Example: tcp://127.0.0.1:46657")
}

func registerValidator(cmd *cobra.Command, args []string) error {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	var context client.CoreContext

	if nodeAddress == "" {
		context = client.NewCoreContextFromConfig(config)
	} else {
		context = client.NewCoreContextFromConfigWithClient(config, nodeAddress)
	}

	if moniker == "" {
		moniker = config.BaseConfig.Moniker
	}

	fmt.Printf("Amount=%d Moniker=%s Website=%s Details=%s\n", amount, moniker, website, details)

	res, err := context.RegisterValidator(amount, moniker, "", website, details)
	if err != nil {
		return err
	}

	fmt.Printf("%v\n", res)

	return nil
}
