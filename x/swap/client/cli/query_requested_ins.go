package cli

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/sharering/shareledger/x/swap/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdRequestedIns() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "requested-ins [address]",
		Short: "get requested swaps in information",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqAddress := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			if _, err := sdk.AccAddressFromBech32(reqAddress); err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			params := &types.QueryRequestedInsRequest{

				Address: reqAddress,
			}

			res, err := queryClient.RequestedIns(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
