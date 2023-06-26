package swap

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/sharering/shareledger/tests"
	"github.com/sharering/shareledger/testutil/network"
	swaptypes "github.com/sharering/shareledger/x/swap/types"
	"github.com/sharering/shareledger/x/utils/denom"
	"github.com/stretchr/testify/suite"
)

type E2ETestQuerySuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
}

func NewE2ETestQuerySuite(cfg network.Config) *E2ETestQuerySuite {
	return &E2ETestQuerySuite{cfg: cfg}
}

func (s *E2ETestQuerySuite) TestGRPCQueryRequest() {
	val := s.network.Validators[0]

	coin2 := sdk.NewCoin(denom.Base, sdk.NewInt(30*denom.ShrExponent))
	testCases := tests.TestCasesGrpc{
		{
			Name:      "Query swap pending request",
			URL:       fmt.Sprintf("%s/sharering/shareledger/swap/requests/pending/1/%s/%s/%s/%s", val.APIAddress, network.Accounts[network.KeyAccount1].String(), "0xx1234x", swaptypes.NetworkNameShareLedger, swaptypes.NetworkNameEthereum),
			Headers:   map[string]string{},
			ExpectErr: false,
			RespType:  &swaptypes.QuerySwapResponse{},
			Expected: &swaptypes.QuerySwapResponse{
				Swaps: []swaptypes.Request{
					{
						Id:          uint64(1),
						SrcAddr:     network.Accounts[network.KeyAccount1].String(),
						DestAddr:    "0xx1234x",
						SrcNetwork:  swaptypes.NetworkNameShareLedger,
						DestNetwork: swaptypes.NetworkNameEthereum,
						Amount:      sdk.NewCoin(denom.Base, sdk.NewInt(1000*denom.ShrExponent)),
						Fee:         coin2,
						Status:      swaptypes.SwapStatusPending,
						TxEvents:    []*swaptypes.TxEvent{},
					},
				},
				Pagination: &query.PageResponse{Total: 1},
			},
		},
	}
	tests.RunTestCasesGrpc(&s.Suite, testCases, val)
}

func (s *E2ETestQuerySuite) TestGRPCQuerySwapBalance() {
	val := s.network.Validators[0]

	coin := sdk.NewCoin(denom.Base, sdk.NewInt(10000000*denom.ShrExponent))
	testCases := tests.TestCasesGrpc{
		{
			Name:      "Query swap balance",
			URL:       fmt.Sprintf("%s/sharering/shareledger/swap/balance", val.APIAddress),
			Headers:   map[string]string{},
			ExpectErr: false,
			RespType:  &swaptypes.QueryBalanceResponse{},
			Expected: &swaptypes.QueryBalanceResponse{
				Balance: &coin,
			},
		},
	}
	tests.RunTestCasesGrpc(&s.Suite, testCases, val)
}

func (s *E2ETestQuerySuite) TestGRPCQuerySchema() {
	val := s.network.Validators[0]
	coin1 := sdk.NewCoin(denom.Base, sdk.NewInt(20*denom.ShrExponent))
	coin2 := sdk.NewCoin(denom.Base, sdk.NewInt(30*denom.ShrExponent))
	testCases := tests.TestCasesGrpc{
		{
			Name:      "Query swap schema eth should be success",
			URL:       fmt.Sprintf("%s/sharering/shareledger/swap/schemas/%s", val.APIAddress, "eth"),
			Headers:   map[string]string{},
			ExpectErr: false,
			RespType:  &swaptypes.QuerySchemaResponse{},
			Expected: &swaptypes.QuerySchemaResponse{
				Schema: swaptypes.Schema{
					Network: "eth",
					Fee: &swaptypes.Fee{
						In:  &coin1,
						Out: &coin2,
					},
				},
			},
		},
		{
			Name:      "Query swap schema not found should be empty",
			URL:       fmt.Sprintf("%s/sharering/shareledger/swap/schemas/%s", val.APIAddress, "not_found"),
			Headers:   map[string]string{},
			ExpectErr: true,
			RespType:  &swaptypes.QuerySchemaResponse{},
			Expected:  &swaptypes.QuerySchemaResponse{},
		},
	}
	tests.RunTestCasesGrpc(&s.Suite, testCases, val)
}

func (s *E2ETestQuerySuite) TestGRPCQuerySchemas() {
	val := s.network.Validators[0]
	coin1 := sdk.NewCoin(denom.Base, sdk.NewInt(20*denom.ShrExponent))
	coin2 := sdk.NewCoin(denom.Base, sdk.NewInt(30*denom.ShrExponent))
	testCases := tests.TestCasesGrpc{
		{
			Name:      "Query list of swap schema should be success",
			URL:       fmt.Sprintf("%s/sharering/shareledger/swap/schemas", val.APIAddress),
			Headers:   map[string]string{},
			ExpectErr: false,
			RespType:  &swaptypes.QuerySchemasResponse{},
			Expected: &swaptypes.QuerySchemasResponse{
				Schemas: []swaptypes.Schema{
					{
						Network: "eth",
						Fee: &swaptypes.Fee{
							In:  &coin1,
							Out: &coin2,
						},
					},
					{
						Network: "hero",
						Fee: &swaptypes.Fee{
							In:  &coin1,
							Out: &coin2,
						},
						Schema: "{}",
					},
					{
						Network: "hero1",
						Fee: &swaptypes.Fee{
							In:  &coin1,
							Out: &coin2,
						},
						Schema: "{}",
					},
					{
						Network: "schemaWithoutFee",
						Fee:     nil,
					},
				},
				Pagination: &query.PageResponse{
					Total: 4,
				},
			},
		},
	}
	tests.RunTestCasesGrpc(&s.Suite, testCases, val)
}

func (s *E2ETestQuerySuite) TestGRPCQueryNextRequestID() {
	val := s.network.Validators[0]

	testCases := tests.TestCasesGrpc{
		{
			Name:      "Query query net request ID data should be success",
			URL:       fmt.Sprintf("%s/sharering/shareledger/swap/next_request_id", val.APIAddress),
			Headers:   map[string]string{},
			ExpectErr: false,
			RespType:  &swaptypes.QueryNextRequestIdResponse{},
			Expected: &swaptypes.QueryNextRequestIdResponse{
				NextCount: 7,
			},
		},
	}
	tests.RunTestCasesGrpc(&s.Suite, testCases, val)
}

func (s *E2ETestQuerySuite) TestGRPCQueryNextBatchId() {
	val := s.network.Validators[0]

	testCases := tests.TestCasesGrpc{
		{
			Name:      "Query next batch id should be success",
			URL:       fmt.Sprintf("%s/sharering/shareledger/swap/next_batch_id", val.APIAddress),
			Headers:   map[string]string{},
			ExpectErr: false,
			RespType:  &swaptypes.QueryNextBatchIdResponse{},
			Expected: &swaptypes.QueryNextBatchIdResponse{
				NextCount: 4,
			},
		},
	}
	tests.RunTestCasesGrpc(&s.Suite, testCases, val)
}

func (s *E2ETestQuerySuite) TestGRPCQueryPastTxEvent() {
	val := s.network.Validators[0]

	testCases := tests.TestCasesGrpc{
		{
			Name:      "Query list transaction event approved swap in should be success",
			URL:       fmt.Sprintf("%s/sharering/shareledger/swap/past_tx_events", val.APIAddress),
			Headers:   map[string]string{},
			ExpectErr: false,
			RespType:  &swaptypes.QueryPastTxEventsResponse{},
			Expected: &swaptypes.QueryPastTxEventsResponse{
				Events: []*swaptypes.PastTxEvent{
					{
						SrcAddr:  "src111",
						DestAddr: "dest111",
					},
					{
						SrcAddr:  "src222",
						DestAddr: "dest22",
					},
				},
				Pagination: &query.PageResponse{Total: 2},
			},
		},
	}
	tests.RunTestCasesGrpc(&s.Suite, testCases, val)
}

func (s *E2ETestQuerySuite) TestGRPCQueryPastTxEventByTxHash() {
	val := s.network.Validators[0]

	testCases := tests.TestCasesGrpc{
		{
			Name:      "Query past transaction event by tx hash should be success",
			URL:       fmt.Sprintf("%s/sharering/shareledger/swap/past_tx_events/%s", val.APIAddress, "hashxxxxx"),
			Headers:   map[string]string{},
			ExpectErr: false,
			RespType:  &swaptypes.QueryPastTxEventsByTxHashResponse{},
			Expected: &swaptypes.QueryPastTxEventsByTxHashResponse{
				Events: []*swaptypes.PastTxEvent{
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
	}
	tests.RunTestCasesGrpc(&s.Suite, testCases, val)
}

func (s *E2ETestQuerySuite) TestGRPCQueryPastTxEventByTxHashLogIndex() {
	val := s.network.Validators[0]

	testCases := tests.TestCasesGrpc{
		{
			Name:      "Query past transaction event by log index and tx hash should be success",
			URL:       fmt.Sprintf("%s/sharering/shareledger/swap/past_tx_events/%s/%d", val.APIAddress, "hashxxxxx", 2),
			Headers:   map[string]string{},
			ExpectErr: false,
			RespType:  &swaptypes.QueryPastTxEventResponse{},
			Expected: &swaptypes.QueryPastTxEventResponse{
				Event: &swaptypes.PastTxEvent{
					SrcAddr:  "src222",
					DestAddr: "dest22",
				},
			},
		},
	}
	tests.RunTestCasesGrpc(&s.Suite, testCases, val)
}
