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
	reserveAddress string
)

// LoadCoinCmd - load coin to an account
var LoadCoinCmd = &cobra.Command{
	Use:   "load_coin",
	Short: "send coin to other account",
	RunE:  loadCoin,
}

func init() {
	LoadCoinCmd.Flags().StringVar(&address, "address", "", "Receiving account")
	LoadCoinCmd.Flags().StringVar(&coinAmount, "amount", "", "Amount. Decimal is possible.")
	LoadCoinCmd.Flags().StringVar(&denom, "denom", "", "Denomination. Available denoms: %s")
	LoadCoinCmd.Flags().StringVar(&nodeAddress, "client", "", "Node address to query info. Example: tcp://123.123.123.123:46657")
	LoadCoinCmd.Flags().StringVar(&reserveAddress, "reserve", "", "Address of the reserve/faucet.")

	LoadCoinCmd.MarkFlagRequired("address")
	LoadCoinCmd.MarkFlagRequired("amount")
	LoadCoinCmd.MarkFlagRequired("denom")
	LoadCoinCmd.MarkFlagRequired("reserve")
	LoadCoinCmd.MarkFlagRequired("nodeAddress")
}

func loadCoin(cmd *cobra.Command, args []string) error {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	var context client.CoreContext
	context = client.NewCoreContextWithClient(reserveAddress, nodeAddress)

	addressBytes, err := hex.DecodeString(address)
	if err != nil {
		return err
	}

	addr := sdk.AccAddress(addressBytes)

	dec, err := types.NewDecFromStr(coinAmount)

	if err != nil {
		return err
	}

	amount := types.NewCoinFromDec(denom, dec)

	fmt.Printf("Address=%X Amount=%s\n", addr, dec)

	res, err := context.LoadBalance(addr, amount)
	if err != nil {
		return err
	}

	fmt.Printf("%v\n", res)

	return nil
}
