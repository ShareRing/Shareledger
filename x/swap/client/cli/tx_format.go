package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/sharering/shareledger/x/swap/types"
	"github.com/spf13/cobra"
)

func CmdCreateFormat() *cobra.Command { //...
	cmd := &cobra.Command{
		Use:     "format [network] [data]",
		Short:   "Create a new format following eip712 format.\n[network] network name\n[data] json format for eip712 data",
		Example: `create-format erc20 '{"types":{"EIP712Domain":[{"name":"name","type":"string"},{"name":"version","type":"string"},{"name":"chainId","type":"uint256"},{"name":"verifyingContract","type":"address"}],"Swap":[{"name":"ids","type":"uint256[]"},{"name":"tos","type":"address[]"},{"name":"amounts","type":"uint256[]"}]},"primaryType":"Swap","domain":{"name":"ShareRingSwap","version":"2.0","chainId":"0x7a69","verifyingContract":"0x0165878a594ca255338adfa4d48449f69242eb8f","salt":""}}''`,
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexNetwork := args[0]

			// Get value arguments
			dataStr := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateFormat(
				clientCtx.GetFromAddress().String(),
				indexNetwork,
				dataStr,
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

func CmdUpdateFormat() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "update-format [network] [data]",
		Short:   "Update a [network] data format",
		Example: `update-format erc20 '{"types":{"EIP712Domain":[{"name":"name","type":"string"},{"name":"version","type":"string"},{"name":"chainId","type":"uint256"},{"name":"verifyingContract","type":"address"}],"Swap":[{"name":"ids","type":"uint256[]"},{"name":"tos","type":"address[]"},{"name":"amounts","type":"uint256[]"}]},"primaryType":"Swap","domain":{"name":"ShareRingSwap","version":"2.0","chainId":"0x7a69","verifyingContract":"0x0165878a594ca255338adfa4d48449f69242eb8f","salt":""}}''`,
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexNetwork := args[0]

			// Get value arguments
			dataStr := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateFormat(
				clientCtx.GetFromAddress().String(),
				indexNetwork,
				dataStr,
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

func CmdDeleteFormat() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-format [network]",
		Short: "Delete a format",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			indexNetwork := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteFormat(
				clientCtx.GetFromAddress().String(),
				indexNetwork,
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
