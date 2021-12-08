package cli

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/ShareRing/Shareledger/x/gentlemint/types"
	utils "github.com/ShareRing/Shareledger/x/utils"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
)

var _ = strconv.Itoa(0)

func CmdEnrollLoaders() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "enroll-loaders [address1] [address2] [address3]",
		Short: "Broadcast message enroll-loaders",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAddresses := args[:]
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgEnrollLoaders(
				clientCtx.GetFromAddress().String(),
				argAddresses,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdEnrollLoadersFromFile() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "enroll-loaders-from-file [filepath]",
		Short: "Broadcast message enroll-loaders",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			addrList, err := utils.GetAddressFromFile(args[0])
			if err != nil {
				return err
			}
			lenAddr := len(addrList)
			reqAddr := make([]string, 0, 5)
			for i := 0; i < lenAddr; i++ {
				reqAddr = append(reqAddr, addrList[i])
				// Send 5 addresses per time. Following the old logic of cli
				if (i+1)%5 == 0 || i == lenAddr-1 {
					msg := types.NewMsgEnrollLoaders(
						clientCtx.GetFromAddress().String(),
						reqAddr[:],
					)
					if err := msg.ValidateBasic(); err != nil {
						return err
					}
					if err := tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg); err != nil {
						return err
					}
					reqAddr = make([]string, 0, 5)
				}
			}
			return nil
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
