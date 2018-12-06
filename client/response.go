package client

import (
	"fmt"

	tdmrpctypes "github.com/tendermint/tendermint/rpc/core/types"
)

// Response - response from any transaction to Shareledger
type Response struct {
	Code uint32 `json:"code"`
	Data []byte `json:"data"`
	Log  string `json:"log"`
	Hash []byte `json:"hash"`
}

func (r Response) String() string {
	return fmt.Sprintf("Code: %d\nData: %v\nLog: %s\nHash: %x\n", r.Code, r.Data, r.Log, r.Hash)
}

func convertBroadcastResult(result *tdmrpctypes.ResultBroadcastTx) Response {
	return Response{
		Code: result.Code,
		Data: []byte(result.Data),
		Log:  result.Log,
		Hash: []byte(result.Hash),
	}
}
