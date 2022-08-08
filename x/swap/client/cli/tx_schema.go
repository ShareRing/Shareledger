package cli

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/x/swap/types"
	"github.com/spf13/cobra"
)

const (
	FlagsFeeIn            = "fee-in"
	FlagsFeeOut           = "fee-out"
	FlagsContractExponent = "exp"
	FlagsSchema           = "schema"
)

func CmdCreateSchema() *cobra.Command { //...
	cmd := &cobra.Command{
		Use:   "create [network] [data] [feeIn] [feeOut] [contractExponent]",
		Short: "Create Schema for signing EIP712 to external chain and set fee to swap actions between Shareledger and external chains",
		Long: `
			[network]: corresponding external network name for this schema
			[data]: json string for eip712 schema signing with smart contract on the external network
			[feeIn]: fee for swap IN from this external network to Shareledger
			[feeOut]: fee for swap OUT from Shareledfe to the external network
			[contractExponent]: base coin's exponent - supported decimal number of deployed token contract on the external network 
			`,
		Example: `create eth '{"types":{"EIP712Domain":[{"name":"name","type":"string"},{"name":"version","type":"string"},{"name":"chainId","type":"uint256"},{"name":"verifyingContract","type":"address"}],"Swap":[{"name":"ids","type":"uint256[]"},{"name":"tos","type":"address[]"},{"name":"amounts","type":"uint256[]"}]},"primaryType":"Swap","domain":{"name":"ShareRingSwap","version":"2.0","chainId":"0x7a69","verifyingContract":"0x0165878a594ca255338adfa4d48449f69242eb8f","salt":""}}' 50shr 100shr 2`,
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

			if feeIn == nil || feeOut == nil || contractExp == 0 {
				return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "create schema requires fee-in fee-out and exponent")
			}

			msg := types.NewMsgCreateSchema(
				clientCtx.GetFromAddress().String(),
				indexNetwork,
				dataStr,
				*feeIn,
				*feeOut,
				contractExp,
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
		Use:     "update [network]",
		Short:   "Update a [network] data schema",
		Example: fmt.Sprintf(`update eth --%s '{"types":{"EIP712Domain":[{"name":"name","type":"string"},{"name":"version","type":"string"},{"name":"chainId","type":"uint256"},{"name":"verifyingContract","type":"address"}],"Swap":[{"name":"ids","type":"uint256[]"},{"name":"tos","type":"address[]"},{"name":"amounts","type":"uint256[]"}]},"primaryType":"Swap","domain":{"name":"ShareRingSwap","version":"2.0","chainId":"0x7a69","verifyingContract":"0x0165878a594ca255338adfa4d48449f69242eb8f","salt":""}}' --%s 10shr --%s 200shr --%s 9`, FlagsSchema, FlagsFeeIn, FlagsFeeOut, FlagsContractExponent),
		Args:    cobra.ExactArgs(1),

		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexNetwork := args[0]

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
	cmd.Flags().String(FlagsSchema, "", "schema data - json string for eip712 schema signing with smart contract on the external network")
	cmd.Flags().String(FlagsFeeIn, "", "swapping fee in")
	cmd.Flags().String(FlagsFeeOut, "", "swapping fee out")
	cmd.Flags().String(FlagsContractExponent, "", "base coin's exponent - supported decimal number of deployed token contract on the external network")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteSchema() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete [network]",
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
	inf, outf, _, err := parseCoinArgs(in, out, "")
	if err != nil {
		return
	}
	return inf, outf, nil
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
	if in != "" {
		i, errP := sdk.ParseDecCoin(in)
		if errP != nil {
			err = errP
			return
		}
		inFee = &i
	}
	if out != "" {
		o, errP := sdk.ParseDecCoin(out)
		if errP != nil {
			err = errP
			return
		}
		outFee = &o
	}

	if ce != "" {
		c, errP := strconv.ParseInt(ce, 10, 32)
		if errP != nil {
			err = errP
			return
		}
		cExp = int32(c)
	}

	return
}
