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
		Short: "Approve batch of swap out transactions. \n [txIDs] format ID1,ID2,ID3 \n [signer_name] approver key name for signing swapping batch\n[network] that applied data format from",
		Args:  cobra.ExactArgs(3),
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
			filter.Status = types.BatchStatusPending
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

			//var signFormatData apitypes.TypedData
			//if err := json.Unmarshal([]byte(formatRes.GetSchema().Schema), &signFormatData); err != nil {
			//	return err
			//}
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
	//signerData := apitypes.TypedData{
	//	Domain: apitypes.TypedDataDomain{
	//		Name:              "ShareRingSwap",
	//		Version:           "2.0",
	//		VerifyingContract: "0x0165878a594ca255338adfa4d48449f69242eb8f",
	//		ChainId:           math.NewHexOrDecimal256(31337),
	//	},
	//	Types: apitypes.Types{
	//		"EIP712Domain": {
	//			{Name: "name", Type: "string"},
	//			{Name: "version", Type: "string"},
	//			{Name: "chainId", Type: "uint256"},
	//			{Name: "verifyingContract", Type: "address"},
	//		},
	//		"Swap": {
	//			{
	//				Name: "ids",
	//				Type: "uint256[]",
	//			},
	//			{
	//				Name: "tos",
	//				Type: "address[]",
	//			},
	//			{
	//				Name: "amounts",
	//				Type: "uint256[]",
	//			},
	//		},
	//	},
	//	PrimaryType: "Swap",
	//}
	//fmt.Println(json.Marshal(signer))

	//signData, err := swap.buildTypedData(signFormatData, requests)
	//if err != nil {
	//	return "", err
	//}
	//signHash, err := crypto2.Keccak256HashEIP712(signData)
	//if err != nil {
	//	return "", err
	//}
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
