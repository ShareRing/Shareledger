package cli

import (
	"fmt"

	"bitbucket.org/shareringvietnam/shareledger-fix/x/asset/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
)

func QueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	assetQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the identity module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	assetQueryCmd.AddCommand(flags.GetCommands(
		CmdGetAsset(storeKey, cdc),
	)...)

	return assetQueryCmd
}

func CmdGetAsset(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "get [uuid]",
		Short: "resolve name",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			uuid := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/asset/%s", queryRoute, uuid), nil)
			if err != nil {
				fmt.Printf("could not get asset - %s \n", uuid)
				return nil
			}

			var out types.Asset
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
