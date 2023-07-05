package crypto

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

func Keccak256HashEIP712(signData apitypes.TypedData) (common.Hash, error) {
	bSignData, err := encodeForEIP712Signing(signData)
	if err != nil {
		return common.Hash{}, err
	}
	return crypto.Keccak256Hash(bSignData), nil
}

func encodeForEIP712Signing(signData apitypes.TypedData) ([]byte, error) {
	domainSeparator, err := signData.HashStruct("EIP712Domain", signData.Domain.Map())
	if err != nil {
		return nil, err
	}
	typedDataHash, err := signData.HashStruct(signData.PrimaryType, signData.Message)
	if err != nil {
		return nil, err
	}
	return []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator), string(typedDataHash))), nil
}
