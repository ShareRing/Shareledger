package swap

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/sharering/shareledger/tests"
	"github.com/sharering/shareledger/testutil/network"
	"github.com/sharering/shareledger/x/swap/client/cli"
	swapTypes "github.com/sharering/shareledger/x/swap/types"
	"github.com/sharering/shareledger/x/utils/denom"
)

func (s *E2ETestQuerySuite) TestQueryRequest() {
	coin2 := sdk.NewCoin(denom.Base, sdk.NewInt(30*denom.ShrExponent))
	testCases := tests.TestCases{
		{
			Name:      "query the valid request",
			Args:      []string{"pending", "--dest_network=eth", "--dest_addr=0xx1234x"},
			ExpectErr: false,
			RespType:  &swapTypes.QuerySwapResponse{},
			Expected: &swapTypes.QuerySwapResponse{
				Swaps: []swapTypes.Request{
					{
						Id:          uint64(1),
						SrcAddr:     network.Accounts[network.KeyAccount1].String(),
						DestAddr:    "0xx1234x",
						SrcNetwork:  swapTypes.NetworkNameShareLedger,
						DestNetwork: swapTypes.NetworkNameEthereum,
						Amount:      sdk.NewCoin(denom.Base, sdk.NewInt(1000*denom.ShrExponent)),
						Fee:         coin2,
						Status:      swapTypes.SwapStatusPending,
						TxEvents:    []*swapTypes.TxEvent{},
					},
				},
				Pagination: &query.PageResponse{},
			},
		},
		{
			Name:      "query by ID",
			Args:      []string{"pending", "--ids=1"},
			ExpectErr: false,
			RespType:  &swapTypes.QuerySwapResponse{},
			Expected: &swapTypes.QuerySwapResponse{
				Swaps: []swapTypes.Request{
					{
						Id:          uint64(1),
						SrcAddr:     network.Accounts[network.KeyAccount1].String(),
						DestAddr:    "0xx1234x",
						SrcNetwork:  swapTypes.NetworkNameShareLedger,
						DestNetwork: swapTypes.NetworkNameEthereum,
						Amount:      sdk.NewCoin(denom.Base, sdk.NewInt(1000*denom.ShrExponent)),
						Fee:         coin2,
						Status:      swapTypes.SwapStatusPending,
						TxEvents:    []*swapTypes.TxEvent{},
					},
				},
				Pagination: &query.PageResponse{},
			},
		},
		{
			Name:      "query by ID not found",
			Args:      []string{"pending", "--ids=20"},
			ExpectErr: false,
			RespType:  &swapTypes.QuerySwapResponse{},
			Expected: &swapTypes.QuerySwapResponse{
				Swaps: []swapTypes.Request{}, Pagination: &query.PageResponse{}},
		},
	}

	tests.RunTestCases(&s.Suite, testCases, cli.CmdRequests(), s.network.Validators[0])
}

func (s *E2ETestQuerySuite) TestQuerySchema() {
	coin1 := sdk.NewCoin(denom.Base, sdk.NewInt(20*denom.ShrExponent))
	coin2 := sdk.NewCoin(denom.Base, sdk.NewInt(30*denom.ShrExponent))
	testCases := tests.TestCases{
		{
			Name:      "query the valid schema request",
			Args:      []string{"eth"},
			ExpectErr: false,
			RespType:  &swapTypes.QuerySchemaResponse{},
			Expected: &swapTypes.QuerySchemaResponse{
				Schema: swapTypes.Schema{
					Network: "eth",
					Fee: &swapTypes.Fee{
						In:  &coin1,
						Out: &coin2,
					},
				},
			},
		},
		{
			Name:      "query non exist schema",
			Args:      []string{"eth222222"},
			ExpectErr: true,
		}, {
			Name:      "query empty schema",
			Args:      []string{},
			ExpectErr: true,
		},
	}
	tests.RunTestCases(&s.Suite, testCases, cli.CmdShowSchema(), s.network.Validators[0])
}

func (s *E2ETestQuerySuite) TestQueryPastTransactionEvent() {

	testCases := tests.TestCases{
		{
			Name:      "query past transaction event by transaction hash",
			Args:      []string{"hashxxxxx"},
			ExpectErr: false,
			RespType:  &swapTypes.QueryPastTxEventsByTxHashResponse{},
			Expected: &swapTypes.QueryPastTxEventsByTxHashResponse{
				Events: []*swapTypes.PastTxEvent{
					{
						SrcAddr:  "src111",
						DestAddr: "dest111",
					},
					{
						SrcAddr:  "src222",
						DestAddr: "dest22",
					},
				},
			},
		},
		{
			Name:      "query query past transaction event by tx hash and log index ",
			Args:      []string{"hashxxxxx", "2"},
			ExpectErr: false,
			RespType:  &swapTypes.QueryPastTxEventResponse{},
			Expected: &swapTypes.QueryPastTxEventResponse{
				Event: &swapTypes.PastTxEvent{
					SrcAddr:  "src222",
					DestAddr: "dest22",
				},
			},
		},
	}
	tests.RunTestCases(&s.Suite, testCases, cli.CmdPastTxEvents(), s.network.Validators[0])
}
func (s *E2ETestQuerySuite) TestQueryBatch() {

	testCases := tests.TestCases{
		{
			Name:      "query batch by ID expect result correct",
			Args:      []string{"--ids", "1"},
			ExpectErr: false,
			RespType:  &swapTypes.QueryBatchesResponse{},
			Expected: &swapTypes.QueryBatchesResponse{
				Pagination: &query.PageResponse{},
				Batches: []swapTypes.Batch{
					{
						Id:         1,
						Signature:  "xx",
						RequestIds: []uint64{3},
						Status:     swapTypes.BatchStatusPending,
						Network:    swapTypes.NetworkNameBinanceSmartChain,
					},
				},
			},
		},
	}
	tests.RunTestCases(&s.Suite, testCases, cli.CmdBatches(), s.network.Validators[0])
}
