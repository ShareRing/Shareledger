package cli

import (
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/sharering/shareledger/x/document/types"
	myutils "github.com/sharering/shareledger/x/utils"
)

var _ = strconv.Itoa(0)

func CmdUpdateDocument() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update [holder id] [proof] [extra data]",
		Short: "Update a document",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argHolder := args[0]
			argProof := args[1]
			argData := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// seed implementation
			keySeed := viper.GetString(myutils.FlagKeySeed)
			if keySeed != "" {
				clientCtx, err = myutils.CreateContextFromSeed(keySeed, clientCtx)
				if err != nil {
					return err
				}
			}

			msg := types.NewMsgUpdateDocument(
				argData,
				argHolder,
				clientCtx.GetFromAddress().String(),
				argProof,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().String(myutils.FlagKeySeed, "", myutils.KeySeedUsage)

	return cmd
}
