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

func CmdCreateIds() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-ids [id] [backup-address] [owner-address] [extra-data]",
		Short: "create ids",
		Long: strings.TrimSpace(fmt.Sprintf(`
CreateAsset new batch of IDs by given information
Example:
$ %s tx %s create uid-159654,uid-159655 shareledger1s432..,shareledgerzv95wpluxhf.. shareledger1s432,shareledgerzv95wpluxhf.. http://sharering.network/id1,http://sharering.network/id2`, version.Name, types.ModuleName)),
		Args: cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			sep := ","
			argId := strings.Split(args[0], sep)
			argBackupAddress := strings.Split(args[1], sep)
			argOwnerAddress := strings.Split(args[2], sep)
			argExtraData := strings.Split(args[3], sep)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateIds(
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
