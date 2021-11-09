package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ShareRing/Shareledger/x/id/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		NewIdTxCmd(),
		// NewIdBatchTxCmd(),
		// UpdateIdTxCmd(),
		// UpdateReplaceIdownerTxCmd(),
	)

	return cmd
}

// NewIdTxCmd returns a CLI command handler for creating a MsgCreateId transaction.
func NewIdTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [id] [backup_address] [owner_address] [extra_data]",
		Short: `Create new Id`,
		Long: `Create a new Id by given information.
eg: 
$ create uid-159654 shareledger1s432u6zv95wpluxhf4qru2ewy58kc3w4tkzm3v shareledger1s432u6zv95wpluxhf4qru2ewy58kc3w4tkzm3v shareledger1s432u6zv95wpluxhf4qru2ewy58kc3w4tkzm3v http://sharering.network
		`,
		Args: cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			// inBuf := bufio.NewReader(cmd.InOrStdin())
			// txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder())

			// cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec()
			clientCtx, err := client.GetClientTxContext(cmd)

			id := args[0]

			backupAddr, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			ownerAddr, err := sdk.AccAddressFromBech32(args[2])
			if err != nil {
				return err
			}

			extraData := args[3]

			// build and sign the transaction, then broadcast to Tendermint
			msg := types.NewMsgCreateId(clientCtx.GetFromAddress(), backupAddr, ownerAddr, id, extraData)

			// return utils.GenerateOrBroadcastMsgs(clientCtx, txBldr, []sdk.Msg{msg})
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// func NewIdBatchTxCmd() *cobra.Command {
// 	cmd := &cobra.Command{
// 		Use:   "create-batch [id] [backup_address] [owner_address] [extra_data]",
// 		Short: `Create id by a batch`,
// 		Long: `Create multiple ids by a bath.
// eg:
// $ create-batch id id1,id2 shareledger1yyc7hnyjqwlxqy9rt8scwuttwa3nt4yn2wl0z0,shareledger1m3tfsfc5hkj78c9kk6k500e3p35tngnl3xymg4 shareledger17kejntc3f8njkermmpgtff0zg44ejlw5aw9jka,shareledger1e37j0radceaa5qp0akffe0j5kxwn9jxfy7ve9k shareledger17kejntc3f8njkermmpgtff0zg44ejlw5aw9jka,shareledger1e37j0radceaa5qp0akffe0j5kxwn9jxfy7ve9k manager,staff
// 		`,
// 		Args: cobra.ExactArgs(4),
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			inBuf := bufio.NewReader(cmd.InOrStdin())
// 			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder())

// 			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec()

// 			seperator := ","
// 			ids := strings.Split(args[0], seperator)
// 			backups := strings.Split(args[1], seperator)
// 			owners := strings.Split(args[2], seperator)
// 			extras := strings.Split(args[3], seperator)

// 			backupAddrs := make([]sdk.AccAddress, 0, len(ids))
// 			ownerAddrs := make([]sdk.AccAddress, 0, len(ids))

// 			for i := 0; i < len(ids); i++ {
// 				backupAddr, err := sdk.AccAddressFromBech32(backups[i])
// 				if err != nil {
// 					return err
// 				}

// 				ownerAddr, err := sdk.AccAddressFromBech32(owners[i])
// 				if err != nil {
// 					return err
// 				}

// 				backupAddrs = append(backupAddrs, backupAddr)
// 				ownerAddrs = append(ownerAddrs, ownerAddr)
// 			}

// 			// build and sign the transaction, then broadcast to Tendermint
// 			msg := idtypes.NewMsgCreateIdBatch(cliCtx.GetFromAddress(), backupAddrs, ownerAddrs, ids, extras)

// 			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
// 		},
// 	}

// 	cmd = flags.PostCommands(cmd)[0]

// 	return cmd
// }

// // NewIdTxCmd returns a CLI command handler for creating a MsgUpdateId transaction.
// func UpdateIdTxCmd() *cobra.Command {
// 	cmd := &cobra.Command{
// 		Use:   "update [id] [extra_data]",
// 		Short: `Update Id`,
// 		Args:  cobra.ExactArgs(2),
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			inBuf := bufio.NewReader(cmd.InOrStdin())
// 			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder())

// 			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec()

// 			id := args[0]

// 			extraData := args[1]

// 			// build and sign the transaction, then broadcast to Tendermint
// 			msg := idtypes.NewMsgUpdateId(cliCtx.GetFromAddress(), id, extraData)

// 			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
// 		},
// 	}

// 	cmd = flags.PostCommands(cmd)[0]

// 	return cmd
// }

// // NewIdTxCmd returns a CLI command handler for creating a MsgCreateId transaction.
// func UpdateReplaceIdownerTxCmd() *cobra.Command {
// 	cmd := &cobra.Command{
// 		Use:   "replace [id] [new_owner_address]",
// 		Short: `Replace owner of an Id`,
// 		Args:  cobra.ExactArgs(2),
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			inBuf := bufio.NewReader(cmd.InOrStdin())
// 			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder())

// 			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec()

// 			id := args[0]

// 			newOwner, err := sdk.AccAddressFromBech32(args[1])
// 			if err != nil {
// 				return err
// 			}

// 			// build and sign the transaction, then broadcast to Tendermint
// 			msg := idtypes.NewMsgReplaceIdOwner(id, newOwner, cliCtx.GetFromAddress())

// 			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
// 		},
// 	}

// 	cmd = flags.PostCommands(cmd)[0]

// 	return cmd
// }
