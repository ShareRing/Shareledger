package requests

import (
	"encoding/json"

	core_types "github.com/tendermint/tendermint/rpc/core/types"
	lib_types "github.com/tendermint/tendermint/rpc/lib/types"

	"github.com/sharering/shareledger/cmd/stress-test/utils"
	"github.com/sharering/shareledger/types"
	"github.com/sharering/shareledger/x/auth"
)

func decodeQueryResponse(res []byte) []byte {
	var rpcResponse lib_types.RPCResponse

	err := json.Unmarshal(res, &rpcResponse)

	utils.Check(err)

	var rpcResult core_types.ResultABCIQuery

	err = json.Unmarshal(rpcResponse.Result, &rpcResult)

	utils.Check(err)

	return []byte(rpcResult.Response.Log)
}

func decodeBalanceResponse(res []byte) types.Coins {

	// decode into Query Response
	res = decodeQueryResponse(res)

	var account auth.SHRAccount

	err := json.Unmarshal(res, &account)

	utils.Check(err)

	return account.Coins
}

func decodeNonceResponse(res []byte) int64 {
	// decode into Query Response
	res = decodeQueryResponse(res)

	var nonce int64

	err := json.Unmarshal(res, &nonce)

	utils.Check(err)

	return nonce
}
