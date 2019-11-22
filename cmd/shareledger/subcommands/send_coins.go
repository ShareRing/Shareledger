package subcommands

import (
	"encoding/hex"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"

	"github.com/sharering/shareledger/client"
	"github.com/sharering/shareledger/types"
)

var (
	address    string
	coinAmount string
	denom      string
)

// SendCoinCmd - send coin to other account
var SendCoinCmd = &cobra.Command{
	Use:   "send_coin",
	Short: "send coin to other account",
	RunE:  sendCoin,
}

func init() {
	SendCoinCmd.Flags().StringVar(&address, "address", "", "Receiving account")
	SendCoinCmd.Flags().StringVar(&coinAmount, "amount", "", "Amount. Decimal is possible.")
	SendCoinCmd.Flags().StringVar(&denom, "denom", "", "Denomination. Available denoms: %s")
	SendCoinCmd.Flags().StringVar(&nodeAddress, "client", "", "Node address to query info. Example: tcp://123.123.123.123:46657")
	SendCoinCmd.MarkFlagRequired("address")
	SendCoinCmd.MarkFlagRequired("amount")
	SendCoinCmd.MarkFlagRequired("denom")
}

func sendCoin(cmd *cobra.Command, args []string) error {

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

	addressBytes, err := hex.DecodeString(address)
	if err != nil {
		return err
	}

	addr := sdk.AccAddress(addressBytes)

	dec, err := types.NewDecFromStr(coinAmount)
	fmt.Printf("coin:%s dec:%v\n", coinAmount, dec)

	if err != nil {
		return err
	}

	amount := types.NewCoinFromDec(denom, dec)

	fmt.Printf("Address=%X Amount=%s\n", addr, dec)

	res, err := context.SendCoins(addr, amount)
	if err != nil {
		return err
	}

	fmt.Printf("%v\n", res)

	return nil
}
