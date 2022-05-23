package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/swap/types"
	"github.com/spf13/cobra"
	"strconv"
)

const (
	FlagsFeeIn            = "fee-in"
	FlagsFeeOut           = "fee-out"
	FlagsContractExponent = "exp"
	FlagsSchema           = "schema"
)

func CmdCreateSchema() *cobra.Command { //...
	cmd := &cobra.Command{
		Use:     "schema [network] [data] [feeIn] [feeOut] [contractExponent]",
		Short:   "Create a new schema following eip712 format.\n\t\t[network] network name\n\t\t[data] json format for eip712 data \n\t\t the in and out fee are swapping fee the address must be pay for swapping.",
		Example: `schema eth '{"types":{"EIP712Domain":[{"name":"name","type":"string"},{"name":"version","type":"string"},{"name":"chainId","type":"uint256"},{"name":"verifyingContract","type":"address"}],"Swap":[{"name":"ids","type":"uint256[]"},{"name":"tos","type":"address[]"},{"name":"amounts","type":"uint256[]"}]},"primaryType":"Swap","domain":{"name":"ShareRingSwap","version":"2.0","chainId":"0x7a69","verifyingContract":"0x0165878a594ca255338adfa4d48449f69242eb8f","salt":""}}' 50shr 100shr 2`,
		Args:    cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexNetwork := args[0]

			// Get value arguments
			dataStr := args[1]
			feeInStr := args[2]
			feeOutStr := args[3]
			contractExponent := args[4]

			if err != nil {
				return err
			}
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			feeIn, feeOut, contractExp, err := parseCoinArgs(feeInStr, feeOutStr, contractExponent)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateSchema(
				clientCtx.GetFromAddress().String(),
				indexNetwork,
				dataStr,
				feeIn, feeOut, contractExp,
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

func CmdUpdateSchema() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "update-schema [network]",
		Short:   "Update a [network] data schema",
		Example: fmt.Sprintf("update_schema eth --%s {\\\"types\\\":{\\\"EIP712Domain\\\":[{\\\"name\\\":\\\"name\\\",\\\"type\\\":\\\"string\\\"},{\\\"name\\\":\\\"version\\\",\\\"type\\\":\\\"string\\\"},{\\\"name\\\":\\\"chainId\\\",\\\"type\\\":\\\"uint256\\\"},{\\\"name\\\":\\\"verifyingContract\\\",\\\"type\\\":\\\"address\\\"}],\\\"Swap\\\":[{\\\"name\\\":\\\"ids\\\",\\\"type\\\":\\\"uint256[]\\\"},{\\\"name\\\":\\\"tos\\\",\\\"type\\\":\\\"address[]\\\"},{\\\"name\\\":\\\"amounts\\\",\\\"type\\\":\\\"uint256[]\\\"}]},\\\"primaryType\\\":\\\"Swap\\\",\\\"domain\\\":{\\\"name\\\":\\\"ShareRingSwap\\\",\\\"version\\\":\\\"2.0\\\",\\\"chainId\\\":\\\"0x7a69\\\",\\\"verifyingContract\\\":\\\"0x0165878a594ca255338adfa4d48449f69242eb8f\\\",\\\"salt\\\":\\\"\\\"}}' --%s 10shr --%s 200shr --%s 9 ", FlagsSchema, FlagsFeeIn, FlagsFeeOut, FlagsContractExponent),
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexNetwork := args[0]

			// Get value arguments

			in, out, exp, dataStr, err := parseSchemaArgsFromCmd(cmd)
			if err != nil {
				return err
			}
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateSchema(
				clientCtx.GetFromAddress().String(),
				indexNetwork,
				dataStr,
				in, out, exp,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	cmd.Flags().String(FlagsSchema, "", "schema data")
	cmd.Flags().String(FlagsFeeIn, "", "swapping fee in")
	cmd.Flags().String(FlagsFeeOut, "", "swapping fee out")
	cmd.Flags().String(FlagsContractExponent, "", "the contract exponent of the contract token")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteSchema() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-schema [network]",
		Short: "Delete a schema",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			indexNetwork := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteSchema(
				clientCtx.GetFromAddress().String(),
				indexNetwork,
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
func parseSchemaArgsFromCmd(cmd *cobra.Command) (inFee, outFee *sdk.DecCoin, cExp int32, schema string, err error) {

	cExpStr, err := cmd.Flags().GetString(FlagsContractExponent)
	if err != nil {
		return
	}
	in, out, err := getInOutStrFromFlag(cmd)
	if err != nil {
		return
	}
	inFee, outFee, cExp, err = parseCoinArgs(in, out, cExpStr)
	if err != nil {
		return
	}
	schema, err = cmd.Flags().GetString(FlagsSchema)
	return
}

func parseInOutFeeFromCmd(cmd *cobra.Command) (inFee, outFee *sdk.DecCoin, err error) {
	in, out, err := getInOutStrFromFlag(cmd)
	inFee, outFee, _, err = parseCoinArgs(in, out, "")
	if err != nil {
		return
	}
	return
}

func getInOutStrFromFlag(cmd *cobra.Command) (i, o string, err error) {
	i, err = cmd.Flags().GetString(FlagsFeeIn)
	if err != nil {
		return
	}
	o, err = cmd.Flags().GetString(FlagsFeeOut)
	if err != nil {
		return
	}
	return

}
func parseCoinArgs(in, out, ce string) (inFee, outFee *sdk.DecCoin, cExp int32, err error) {
	inFee = new(sdk.DecCoin)
	outFee = new(sdk.DecCoin)
	if in != "" {
		*inFee, err = sdk.ParseDecCoin(in)
		if err != nil {
			return
		}
	}
	if out != "" {
		*outFee, err = sdk.ParseDecCoin(out)
		if err != nil {
			return
		}
	}
	var c int64
	if ce != "" {
		c, err = strconv.ParseInt(ce, 10, 32)
		if err != nil {
			return
		}
		cExp = int32(c)
	}

	return
}
