package swap

import (
	"encoding/json"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	crypto2 "github.com/sharering/shareledger/pkg/crypto"
	swapmoduletypes "github.com/sharering/shareledger/x/swap/types"
	"github.com/sharering/shareledger/x/utils/denom"
	"math/big"
)

type BatchDetail struct {
	Batch swapmoduletypes.Batch
	SignDetail
}

func NewBatchDetail(batch swapmoduletypes.Batch, requests []swapmoduletypes.Request, signSchema swapmoduletypes.Schema) BatchDetail {
	return BatchDetail{
		Batch:      batch,
		SignDetail: NewSignDetail(requests, signSchema),
	}
}

type SignDetail struct {
	Requests   []swapmoduletypes.Request
	SignSchema swapmoduletypes.Schema
}

func NewSignDetail(requests []swapmoduletypes.Request, signSchema swapmoduletypes.Schema) SignDetail {
	return SignDetail{
		Requests:   requests,
		SignSchema: signSchema,
	}
}

func (b SignDetail) Validate() error {
	if len(b.Requests) == 0 {
		return fmt.Errorf("requests is empty")
	}
	if len(b.SignSchema.Schema) == 0 {
		return fmt.Errorf("schema is empty")
	}
	return nil
}

// Digest Keccak256HashEIP712 param coins will be converted from shareledger base denom to contract base denom
func (b SignDetail) Digest() (common.Hash, error) {
	var hash common.Hash
	var signFormatData apitypes.TypedData
	if err := json.Unmarshal([]byte(b.SignSchema.Schema), &signFormatData); err != nil {
		return hash, err
	}
	params, err := b.GetContractParams()
	if err != nil {
		return hash, err
	}
	data, err := buildTypedData(signFormatData, params)

	if err != nil {
		return hash, err
	}
	hash, err = crypto2.Keccak256HashEIP712(data)
	return hash, err
}

type SwapParams struct {
	TransactionIds []*big.Int
	DestAddrs      []common.Address
	Amounts        []*big.Int
}

// GetContractParams return params for interact with contract
// The Amounts will be converted from base coins to exponent that configured in SignSchema ContractExponent
func (b SignDetail) GetContractParams() (params SwapParams, err error) {
	if b.SignSchema.ContractExponent == 0 {
		return SwapParams{}, fmt.Errorf("contract exponent is required")
	}
	params = SwapParams{
		TransactionIds: make([]*big.Int, 0, len(b.Requests)),
		DestAddrs:      make([]common.Address, 0, len(b.Requests)),
		Amounts:        make([]*big.Int, 0, len(b.Requests)),
	}
	for _, r := range b.Requests {
		if err != nil {
			return params, err
		}
		params.TransactionIds = append(params.TransactionIds, big.NewInt(int64(r.Id)))
		params.DestAddrs = append(params.DestAddrs, common.HexToAddress(r.DestAddr))
		total, err := denom.ShrCoinsToExponent(sdk.NewDecCoinsFromCoins(r.Amount), int(b.SignSchema.ContractExponent), false)
		if err != nil {
			return SwapParams{}, err
		}
		params.Amounts = append(params.Amounts, total)
	}
	return
}
