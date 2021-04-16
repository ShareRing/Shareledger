package cli

import (
	"fmt"

	myutils "github.com/ShareRing/modules/utils"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/sharering/shareledger/x/gentlemint/types"

	"github.com/spf13/cobra"
)

func QueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	gentlemintQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the gentlemint module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	gentlemintQueryCmd.AddCommand(flags.GetCommands(
		CmdGetLoader(storeKey, cdc),
		CmdGetLoadersFromFile(storeKey, cdc),
		CmdGetExchange(storeKey, cdc),
		CmdGetIdSigner(storeKey, cdc),
		CmdGetAllIdSigner(storeKey, cdc),
		CmdGetAccountOperator(storeKey, cdc),
		CmdGetAllAccountOperator(storeKey, cdc),
		CmdGetDocumentIssuer(storeKey, cdc),
		CmdGetAllDocumentIssuer(storeKey, cdc),
	)...)

	return gentlemintQueryCmd
}

func CmdGetLoader(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "get-loader [address]",
		Short: "get shrp loader from address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			addr := "shrploader" + args[0]
			queryPath := fmt.Sprintf("custom/%s/loader/%s", queryRoute, addr)
			res, _, err := cliCtx.QueryWithData(queryPath, nil)
			if err != nil {
				fmt.Printf("could not get loader - %s \n", addr)
				return nil
			}

			var out types.SHRPLoader
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

func CmdGetLoadersFromFile(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "get-loaders-from-file [file-path]",
		Short: "get shrp loaders from json file of addresses",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			addrList, err := myutils.GetAddressFromFile(args[0])
			if err != nil {
				return err
			}
			for i, addr := range addrList {
				fmt.Println("loader", i, addr)
				addr = "shrploader" + addr
				queryPath := fmt.Sprintf("custom/%s/loader/%s", queryRoute, addr)
				res, _, err := cliCtx.QueryWithData(queryPath, nil)
				if err != nil {
					fmt.Printf("could not get loader - %s \n", addr)
					return nil
				}

				var out types.SHRPLoader
				cdc.MustUnmarshalJSON(res, &out)
				if err := cliCtx.PrintOutput(out); err != nil {
					return err
				}
			}
			return nil
		},
	}
}

func CmdGetExchange(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "get-exchange",
		Short: "get shrp to shr exchange rate",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			queryPath := fmt.Sprintf("custom/%s/exchange", queryRoute)
			res, _, err := cliCtx.QueryWithData(queryPath, nil)
			if err != nil {
				fmt.Println("could not get exchange", err)
				return nil
			}

			var out string
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
