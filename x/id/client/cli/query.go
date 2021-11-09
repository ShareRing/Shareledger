package cli

import (
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/version"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ShareRing/Shareledger/x/id/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group id queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		GetIdByAddressCmd(types.QuerierRoute),
	)

	return cmd
}

func GetIdByAddressCmd(queryRoute string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "info <address>|<id> [address,[id]]",
		Short: "Query for id information",
		Long: strings.TrimSpace(fmt.Sprintf(`
Query id information of an account by owner address or the id.
Example:
$ %s query %s info address shareledger1s432u6zv95wpluxhf4qru2ewy58kc3w4tkzm3v
$ %s query %s info id 123e4567-e89b-12d3-a456-426655440000`, version.Name, types.ModuleName, version.Name, types.ModuleName)),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			// cliCtx := context.NewCLIContext().WithCodec(cdc)
			cliCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(cliCtx)

			if args[0] == types.QueryPathAddress {
				addr, addrErr := sdk.AccAddressFromBech32(args[1])
				if addrErr != nil {
					return addrErr
				}
				params := types.QueryIdByAddressRequest{Address: addr.String()}
				res, err := queryClient.IdByAddress(cmd.Context(), &params)
				if err != nil {
					return err
				}
				return cliCtx.PrintProto(res)

			}
			// else if args[0] == types.QueryPathId {
			// 	params := types.QueryIdByAddressRequest{Address: addr.String()}
			// 	res, err := queryClient.GetIdByAddress(cmd.Context(), &params)
			// }

			return errors.New("unknow command: " + args[0])

		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
