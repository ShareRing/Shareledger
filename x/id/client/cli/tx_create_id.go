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
	"github.com/sharering/shareledger/x/id/types"
)

var _ = strconv.Itoa(0)

func CmdCreateId() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [id] [backup-address] [owner-address] [extra-data]",
		Short: "CreateAsset new id",
		Long: strings.TrimSpace(fmt.Sprintf(`
CreateAsset a new Id by given information
Example:
$ %s tx %s create uid-159654 shareledger1s432u6zv95wpluxhf4qru2ewy58kc3w4tkzm3v shareledger1s432u6zv95wpluxhf4qru2ewy58kc3w4tkzm3v http://sharering.network`, version.Name, types.ModuleName)),
		Args: cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argId := args[0]
			argBackupAddress := args[1]
			argOwnerAddress := args[2]
			argExtraData := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateId(
				clientCtx.GetFromAddress().String(),
				argBackupAddress,
				argExtraData,
				argId,
				argOwnerAddress,
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
