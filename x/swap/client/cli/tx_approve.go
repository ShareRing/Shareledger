package cli

import (
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	crypto2 "github.com/sharering/shareledger/pkg/crypto"
	"github.com/sharering/shareledger/x/swap/types"
	denom "github.com/sharering/shareledger/x/utils/demo"
	"github.com/spf13/cobra"
	"math/big"
	"strconv"
	"strings"
)

var _ = strconv.Itoa(0)

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

			var filter types.QuerySearchRequest
			filter.Status = types.BatchStatusPending
			filter.Ids = txIds[:]
			filter.Pagination = &query.PageRequest{
				Offset: 0,
				Limit:  uint64(len(txIds)),
			}
			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.Search(cmd.Context(), &filter)
			if err != nil {
				return err
			}
			if len(res.Request) != len(txIds) {
				return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "the swapping requests ids list does not match with current pending requests")
			}

			formatRes, err := queryClient.Format(cmd.Context(), &types.QueryGetFormatRequest{Network: networkName})
			if err != nil {
				return err
			}

			var signFormatData apitypes.TypedData
			if err := json.Unmarshal([]byte(formatRes.Format.DataFormat), &signFormatData); err != nil {
				return err
			}

			signedHash, err := signApprovedSwap(clientCtx, signer, res.Request, signFormatData)
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

func signApprovedSwap(ctx client.Context, signer string, requests []types.Request, signFormatData apitypes.TypedData) (string, error) {
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
	fmt.Println(json.Marshal(signer))

	txIds := make([]interface{}, 0, len(requests))
	destinations := make([]interface{}, 0, len(requests))
	amounts := make([]interface{}, 0, len(requests))
	for _, tx := range requests {
		txIds = append(txIds, (*math.HexOrDecimal256)(new(big.Int).SetUint64(tx.Id)))
		destinations = append(destinations, tx.DestAddr)
		bCoin, err := denom.NormalizeToBaseCoins(sdk.NewDecCoins(*tx.Amount), false)
		if err != nil {
			return "", err
		}
		amounts = append(amounts, (*math.HexOrDecimal256)(big.NewInt(bCoin.AmountOf(denom.Base).Int64())))
	}
	signFormatData.Message = apitypes.TypedDataMessage{
		"ids":     txIds,
		"tos":     destinations,
		"amounts": amounts,
	}
	signHash, err := crypto2.Keccak256HashEIP712(signFormatData)
	if err != nil {
		return "", err
	}

	kb := ctx.Keyring
	ks := keyring.NewKeyRingEIP712(kb)
	sig, npk, err := ks.Sign(signer, signHash.Bytes())
	if err != nil {
		return "", err
	}
	fmt.Println("digest", signHash.String())
	fmt.Println("Signed eip712", hexutil.Encode(sig))
	fmt.Println("Signed Address EIP712", hexutil.Encode(npk.Address().Bytes()))
	fmt.Println("Trying to verify", npk.VerifySignature(signHash.Bytes(), sig))
	return hexutil.Encode(sig), nil
}
