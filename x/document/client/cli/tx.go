package cli

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"

	"github.com/ShareRing/Shareledger/x/document/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
)

func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		NewDocumentTxCmd(),
		NewDocumentInBatchTxCmd(),
		UpdateDocCmd(),
		RevokeDocCmd(),
	)

	return cmd
}

func NewDocumentTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [holder id] [proof] [extra_data]",
		Short: `Create new docoment`,
		Long: `Create a new document by given information.
eg: 
$ create uid-159654 c89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc6 https://sharering.network/id/463"
		`,
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			// inBuf := bufio.NewReader(cmd.InOrStdin())
			// txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			// cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			issuer := clientCtx.GetFromAddress()
			holderId := args[0]
			proof := args[1]
			data := args[2]

			// build and sign the transaction, then broadcast to Tendermint
			msg := types.NewMsgCreateDoc(issuer.String(), holderId, proof, data)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func NewDocumentInBatchTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-batch [holder id] [proof] [extra_data]",
		Short: `Create new docoment`,
		Long: `Create a new document by given information.
eg: 
$ create uid-159654 c89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc6 https://sharering.network/id/463"
		`,
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			issuer := clientCtx.GetFromAddress()

			sep := ","
			holderId := strings.Split(args[0], sep)
			proof := strings.Split(args[1], sep)
			data := strings.Split(args[2], sep)

			// build and sign the transaction, then broadcast to Tendermint
			msg := types.NewMsgCreateDocBatch(issuer.String(), holderId, proof, data)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func UpdateDocCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update [hodler id] [proof] [data]",
		Short: `Update the data of a document`,
		Long: `Update the data of a document.
eg: 
$ update uid-159654 c89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc6 https://sharering.network/id/463"
		`,
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			issuer := clientCtx.GetFromAddress()

			holderId := args[0]
			proof := args[1]
			data := args[2]

			// build and sign the transaction, then broadcast to Tendermint
			msg := types.NewMsgUpdateDoc(issuer.String(), holderId, proof, data)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func RevokeDocCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "revoke [holder id] [proof]",
		Short: `Revoke a document.`,
		Long: `Revoke a document.
eg: 
$ revoke uid-159654 c89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc6"
		`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			issuer := clientCtx.GetFromAddress()

			holderId := args[0]
			proof := args[1]

			// build and sign the transaction, then broadcast to Tendermint
			msg := types.NewMsgRevokeDoc(issuer.String(), holderId, proof)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
