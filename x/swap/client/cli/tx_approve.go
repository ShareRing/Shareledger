package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/sharering/shareledger/pkg/swap"
	"github.com/sharering/shareledger/x/swap/types"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
)

func CmdApprove() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "approve [txIDs] [signer_name] [network_name]",
		Short: "approve batch of swap out transactions. The approved requests will be batched and sent as a batch by relayer",
		Long: `
		 	[txIDs] format ID1,ID2,ID3 
			[signer_name] approver key name for signing swapping batch. The key name should be imported with HD-PATH corresponding to destination network
			[network] that applied data schema for singing EIP712. The network should exist in schema data`,
		Example: `approve 1,2,3 eth_signer eth`,
		Args:    cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argIDS := strings.Split(args[0], ",")
			txIds := make([]uint64, 0, len(argIDS))
			for _, str := range argIDS {
				id, err := strconv.ParseUint(str, 10, 64)
				if err != nil {
					return err
				}
				txIds = append(txIds, id)
			}
			signer := args[1]
			networkName := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			var filter types.QuerySwapRequest
			filter.Status = types.SwapStatusPending
			filter.Ids = txIds[:]
			filter.Pagination = &query.PageRequest{
				Offset: 0,
				Limit:  uint64(len(txIds)),
			}
			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.Swap(cmd.Context(), &filter)
			if err != nil {
				return err
			}
			if len(res.Swaps) != len(txIds) {
				return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "the swapping requests ids list does not match with current pending requests")
			}

			formatRes, err := queryClient.Schema(cmd.Context(), &types.QueryGetSchemaRequest{Network: networkName})
			if err != nil {
				return err
			}

			signDetail := swap.NewSignDetail(res.Swaps, formatRes.GetSchema())
			signedHash, err := signApprovedSwap(clientCtx, signer, signDetail)
			if err != nil {
				return err
			}
			msg := types.NewMsgApprove(
				clientCtx.GetFromAddress().String(),
				signedHash,
				txIds,
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

func signApprovedSwap(ctx client.Context, signer string, signData swap.SignDetail) (string, error) {

	digest, err := signData.Digest()
	if err != nil {
		return "", err
	}
	kb := ctx.Keyring
	ks := keyring.NewKeyRingETH(kb)
	sig, npk, err := ks.Sign(signer, digest.Bytes())
	if err != nil {
		return "", err
	}
	fmt.Println("digest", digest.String())
	fmt.Println("Signed eip712", hexutil.Encode(sig))
	fmt.Println("Signed Address EIP712", hexutil.Encode(npk.Address().Bytes()))
	fmt.Println("Trying to verify", npk.VerifySignature(digest.Bytes(), sig))
	return hexutil.Encode(sig), nil
}
