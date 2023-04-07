package simulation

import (
	"encoding/json"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/sharering/shareledger/testutil"
	"github.com/sharering/shareledger/x/swap/types"
	"github.com/sharering/shareledger/x/utils/denom"
	"github.com/thanhpk/randstr"
	"math/rand"
)

const (
	MaximumBatchCount   = 3
	MaximumRequestCount = 20
)

func MustGenRandGenesis(simState *module.SimulationState) {

	batchCount := simState.Rand.Int63n(MaximumBatchCount-0) + MaximumBatchCount
	requestCount := simState.Rand.Int63n(MaximumRequestCount-0) + MaximumRequestCount
	schemas := MustRandSchema(simState.Rand, simState.Accounts, 4)

	swapGenesis := &types.GenesisState{
		Schemas: schemas,
	}

	batches, request := MustGenRandBatchAndRequest(simState.Rand, simState.Accounts, schemas, int(requestCount), 21, int(batchCount))

	swapGenesis.BatchCount = uint64(len(batches)) + uint64(batchCount)
	swapGenesis.RequestCount = uint64(len(request)) + uint64(requestCount)
	swapGenesis.Requests = request
	swapGenesis.Batches = batches

	if err := swapGenesis.Validate(); err != nil {
		panic(fmt.Sprintf("swap genesis generation invalid err %s", err.Error()))
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(swapGenesis)
	return
}

func MustGenRandBatchAndRequest(rand *rand.Rand, simAcc []simulation.Account, schemas []types.Schema, reqCount, numReq, batchCount int) ([]types.Batch, []types.Request) {
	var requests = make([]types.Request, 0, numReq)
	var approvedRequest = make([]types.Request, 0, numReq)
	for i := reqCount; i < numReq+reqCount; i++ {
		if testutil.RandBool(rand) {
			//RequestIN
			requests = append(requests, mustGenRandRequestIn(rand, simAcc, i, schemas))
			continue
		}
		requestOut := mustGenRandRequestOut(rand, simAcc, i, schemas)
		if testutil.RandBool(rand) {
			requests = append(requests, requestOut)
			continue
		}

		requestOut.Status = types.SwapStatusApproved
		approvedRequest = append(approvedRequest, requestOut)
	}
	//random picking request for batch

	ethBatch := MustGenBatch(rand, batchCount, types.NetworkNameEthereum)
	bscBatch := MustGenBatch(rand, batchCount+1, types.NetworkNameBinanceSmartChain)

	for i := range approvedRequest {
		if approvedRequest[i].DestNetwork == types.NetworkNameEthereum {
			ethBatch.RequestIds = append(ethBatch.RequestIds, approvedRequest[i].Id)
			approvedRequest[i].BatchId = ethBatch.Id
		}
		if approvedRequest[i].DestNetwork == types.NetworkNameBinanceSmartChain {
			bscBatch.RequestIds = append(bscBatch.RequestIds, approvedRequest[i].Id)
			approvedRequest[i].BatchId = bscBatch.Id
		}
	}

	requests = append(requests, approvedRequest...)
	batches := []types.Batch{ethBatch, bscBatch}

	return batches, requests
}

func mustGenRandRequestOut(r *rand.Rand, simAcc []simulation.Account, id int, schemas []types.Schema) types.Request {
	acc, _ := simulation.RandomAcc(r, simAcc)

	randEthAddr := fmt.Sprintf("0x%s", randstr.Hex(40))
	shrRand := rand.Int63n(10000000000000-1000000000) + 10000000000000

	amount := sdk.NewCoin(denom.Base, sdk.NewInt(shrRand))
	network := testutil.RandNetwork(r)

	request := types.Request{
		Id:          uint64(id),
		SrcAddr:     acc.Address.String(),
		DestAddr:    randEthAddr,
		SrcNetwork:  types.NetworkNameShareLedger,
		DestNetwork: network,
		Amount:      amount,
		Status:      "pending",
	}

	for _, s := range schemas {
		if network == s.Network {
			request.Fee = *s.Fee.Out
		}
	}
	return request
}
func mustGenRandRequestIn(r *rand.Rand, simAcc []simulation.Account, id int, schemas []types.Schema) types.Request {
	acc, _ := simulation.RandomAcc(r, simAcc)
	randEthAddr := fmt.Sprintf("0x%s", randstr.Hex(40))

	shrRand := rand.Int63n(10000000000000-1000000000) + 10000000000000

	amount := sdk.NewCoin(denom.Base, sdk.NewInt(shrRand))

	network := testutil.RandNetwork(r)
	request := types.Request{
		Id:          uint64(id),
		SrcAddr:     randEthAddr,
		DestAddr:    acc.Address.String(),
		SrcNetwork:  network,
		DestNetwork: types.NetworkNameShareLedger,
		Amount:      amount,

		Status: "pending",
	}

	for _, s := range schemas {
		if network == s.Network {
			request.Fee = *s.Fee.In
		}
	}

	return request
}

func MustRandSchema(r *rand.Rand, simAcc []simulation.Account, numSchema int) []types.Schema {
	schemas := make([]types.Schema, numSchema)
	if numSchema < 2 {
		panic("at least 2 schema is needed in simulation")
	}
	acc, _ := simulation.RandomAcc(r, simAcc)

	shrRandIn := rand.Int63n(10000000000000-1000000000) + 10000000000000
	amountIn := sdk.NewCoin(denom.Base, sdk.NewInt(shrRandIn))

	shrRandOut := rand.Int63n(10000000000000-1000000000) + 10000000000000
	amountOut := sdk.NewCoin(denom.Base, sdk.NewInt(shrRandOut))

	ethFee := &types.Fee{
		In:  testutil.PtrOf(amountIn),
		Out: testutil.PtrOf(amountOut),
	}

	ethS := apitypes.TypedData{
		Domain: apitypes.TypedDataDomain{
			VerifyingContract: fmt.Sprintf("0x%s", randstr.Hex(40)),
		},
	}

	bz, err := json.Marshal(ethS)
	if err != nil {
		panic(err)
	}
	ethSchema := types.Schema{
		Network:          "eth",
		Creator:          acc.Address.String(),
		Schema:           string(bz),
		ContractExponent: int32(simulation.RandIntBetween(r, 9, 18)),
		Fee:              ethFee,
	}
	shrRandIn = rand.Int63n(10000000000000-1000000000) + 10000000000000
	amountIn = sdk.NewCoin(denom.Base, sdk.NewInt(shrRandIn))

	shrRandOut = rand.Int63n(10000000000000-1000000000) + 10000000000000
	amountOut = sdk.NewCoin(denom.Base, sdk.NewInt(shrRandOut))

	bscFee := &types.Fee{
		In:  testutil.PtrOf(amountIn),
		Out: testutil.PtrOf(amountOut),
	}

	bscS := apitypes.TypedData{
		Domain: apitypes.TypedDataDomain{
			VerifyingContract: fmt.Sprintf("0x%s", randstr.Hex(40)),
		},
	}

	bz, err = json.Marshal(bscS)
	if err != nil {
		panic(err)
	}

	acc, _ = simulation.RandomAcc(r, simAcc)
	bscSchema := types.Schema{
		Network:          "bsc",
		Creator:          acc.Address.String(),
		Schema:           string(bz),
		ContractExponent: int32(simulation.RandIntBetween(r, 9, 18)),
		Fee:              bscFee,
	}

	schemas[0] = ethSchema
	schemas[1] = bscSchema

	for i := 2; i < numSchema; i++ {

		shrRandIn = rand.Int63n(10000000000000-1000000000) + 10000000000000
		amountIn = sdk.NewCoin(denom.Base, sdk.NewInt(shrRandIn))

		shrRandOut = rand.Int63n(10000000000000-1000000000) + 10000000000000
		amountOut = sdk.NewCoin(denom.Base, sdk.NewInt(shrRandOut))

		randFee := &types.Fee{
			In:  testutil.PtrOf(amountIn),
			Out: testutil.PtrOf(amountOut),
		}

		randSchemaS := apitypes.TypedData{
			Domain: apitypes.TypedDataDomain{
				VerifyingContract: fmt.Sprintf("0x%s", randstr.Hex(40)),
			},
		}
		bz, err = json.Marshal(randSchemaS)
		if err != nil {
			panic(err)
		}
		randSchema := types.Schema{
			Network:          randstr.String(4),
			Creator:          acc.Address.String(),
			Schema:           string(bz),
			ContractExponent: int32(simulation.RandIntBetween(r, 9, 18)),
			Fee:              randFee,
		}
		schemas[i] = randSchema
	}
	return schemas
}

func MustGenBatch(_ *rand.Rand, batchId int, network string) types.Batch {
	return types.Batch{
		Id:        uint64(batchId),
		Signature: randstr.Hex(50),
		Status:    "approved",
		Network:   network,
	}
}
