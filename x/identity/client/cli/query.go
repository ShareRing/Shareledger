package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/identity/types"
	"github.com/spf13/cobra"
)

func QueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	identityQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the identity module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	identityQueryCmd.AddCommand(flags.GetCommands(
		CmdGetId(storeKey, cdc),
		CmdGetIdSigner(storeKey, cdc),
	)...)

	return identityQueryCmd
}

func CmdGetId(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "id [address]",
		Short: "get id from address",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			addr, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}
			idKey := "Id" + addr.String()
			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/id/%s", queryRoute, idKey), nil)
			if err != nil {
				return err
			}
			var out string
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

func CmdGetIdSigner(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "signer [address]",
		Short: "get id signer from address",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			addr, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}
			idKey := "IdSigner" + addr.String()
			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/signer/%s", queryRoute, idKey), nil)
			if err != nil {
				return err
			}
			var out types.IdSigner
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
