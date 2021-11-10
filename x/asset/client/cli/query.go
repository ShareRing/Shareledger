package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ShareRing/Shareledger/x/asset/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group asset queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		CmdGetAsset(types.StoreKey),
	)

	return cmd
}

func CmdGetAsset(queryRoute string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get [uuid]",
		Short: "Get asset by uuid",
		Long: `Get asset by uuid
eg: asset get uuid1
		`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			uuid := args[0]
			queryClient := types.NewQueryClient(cliCtx)

			params := types.QueryAssetByUUIDRequest{Uuid: uuid}
			res, err := queryClient.AssetByUUID(cmd.Context(), &params)

			if err != nil {
				fmt.Printf("could not get asset - %s, %v \n", uuid, err)
				return err
			}
			return cliCtx.PrintProto(res)

		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
