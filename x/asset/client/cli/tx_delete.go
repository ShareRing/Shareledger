package cli

import (
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/sharering/shareledger/x/asset/types"
	myutils "github.com/sharering/shareledger/x/utils"
)

var _ = strconv.Itoa(0)

func CmdDelete() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete [uuid]",
		Short: "Broadcast message DeleteAsset",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argUUID := args[0]

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

			msg := types.NewMsgDelete(
				clientCtx.GetFromAddress().String(),
				argUUID,
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
