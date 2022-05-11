package swap

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	crypto2 "github.com/sharering/shareledger/pkg/crypto"
	swapmoduletypes "github.com/sharering/shareledger/x/swap/types"
)

type BatchDetail struct {
	Batch      swapmoduletypes.Batch
	Requests   []swapmoduletypes.Request
	SignSchema swapmoduletypes.SignSchema
}

func NewBatchDetail(batch swapmoduletypes.Batch, requests []swapmoduletypes.Request, signSchema swapmoduletypes.SignSchema) BatchDetail {
	return BatchDetail{
		Batch:      batch,
		Requests:   requests,
		SignSchema: signSchema,
	}
}

func (b BatchDetail) Validate() error {
	if len(b.Requests) == 0 {
		return fmt.Errorf("requests is empty")
	}
	if len(b.SignSchema.Schema) == 0 {
		return fmt.Errorf("schema is empty")
	}
	return nil
}

func (b BatchDetail) Digest() (common.Hash, error) {
	var hash common.Hash
	var signFormatData apitypes.TypedData
	if err := json.Unmarshal([]byte(b.SignSchema.Schema), &signFormatData); err != nil {
		return hash, err
	}
	data, err := BuildTypedData(signFormatData, b.Requests)
	if err != nil {
		return hash, err
	}
	hash, err = crypto2.Keccak256HashEIP712(data)
	return hash, err
}
