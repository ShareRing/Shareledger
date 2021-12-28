package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/sharering/shareledger/x/document/types"
	myutils "github.com/sharering/shareledger/x/utils"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/version"
)

var _ = strconv.Itoa(0)

func CmdCreateDocumentInBatch() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-batch [holder id] [proof] [extra data]",
		Short: "Create new document in batch",
		Long: strings.TrimSpace(fmt.Sprintf(`
Example:
$ %s tx %s create-batch uuid-5122,uuid-0218 c89efdaa54c0f20c7adf6,c89efdaa54c0f20c7adf6 https://sharering.network/id/463,https://sharering.network/id/463`, version.Name, types.ModuleName)),
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			sep := ","
			argHolder := strings.Split(args[0], sep)
			argProof := strings.Split(args[1], sep)
			argData := strings.Split(args[2], sep)

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

			msg := types.NewMsgCreateDocumentInBatch(
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
