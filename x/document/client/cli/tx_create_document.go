package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/sharering/shareledger/x/document/types"
)

var _ = strconv.Itoa(0)

func CmdCreateDocument() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [holder id] [proof] [extra data]",
		Short: "CreateAsset a new document",
		Long: strings.TrimSpace(fmt.Sprintf(`
Example:
$ %s tx %s create uuid-5132 c89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc6 https://sharering.network/id/463`, version.Name, types.ModuleName)),
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argHolder := args[0]
			argProof := args[1]
			argData := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateDocument(
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

	return cmd
}
