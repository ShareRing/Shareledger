package gentlemint

import (
	"fmt"

	"github.com/sharering/shareledger/tests"
	"github.com/sharering/shareledger/x/gentlemint/types"
)

func (s *E2ETestSuite) TestGRPCQueryDocumentByHolderId() {
	val := s.network.Validators[0]
	baseURL := val.APIAddress
	testCases := tests.TestCasesGrpc{
		{
			Name:      "gRPC exchange rate ok",
			URL:       fmt.Sprintf("%s/shareledger/gentlemint/exchangeRate", baseURL),
			Headers:   map[string]string{},
			ExpectErr: false,
			RespType:  &types.QueryExchangeRateResponse{},
			Expected: &types.QueryExchangeRateResponse{
				Rate: exchangeRate.Rate,
			},
		},
		{
			Name:      "gRPC level fee ok",
			URL:       fmt.Sprintf("%s/shareledger/gentlemint/levelFee/%s", baseURL, "low"),
			Headers:   map[string]string{},
			ExpectErr: false,
			RespType:  &types.QueryLevelFeeResponse{},
			Expected: &types.QueryLevelFeeResponse{
				LevelFee: types.LevelFeeDetail{
					Level:        defaultLevelFees[0].Level,
					Creator:      defaultLevelFees[0].Creator,
					OriginalFee:  defaultLevelFees[0].Fee,
					ConvertedFee: &lowConvertedFee,
				},
			},
		},
		{
			Name:      "gRPC invalid level fee",
			URL:       fmt.Sprintf("%s/shareledger/gentlemint/levelFee/%s", baseURL, "test"),
			Headers:   map[string]string{},
			ExpectErr: true,
		},
		{
			Name:      "gRPC list level fee ok",
			URL:       fmt.Sprintf("%s/shareledger/gentlemint/levelFee", baseURL),
			Headers:   map[string]string{},
			ExpectErr: false,
			RespType:  &types.QueryLevelFeesResponse{},
			Expected: &types.QueryLevelFeesResponse{
				LevelFees: []types.LevelFeeDetail{
					{
						Level:        defaultLevelFees[0].Level,
						Creator:      defaultLevelFees[0].Creator,
						OriginalFee:  defaultLevelFees[0].Fee,
						ConvertedFee: &lowConvertedFee,
					},
					{
						Level:        defaultLevelFees[1].Level,
						Creator:      defaultLevelFees[1].Creator,
						OriginalFee:  defaultLevelFees[1].Fee,
						ConvertedFee: &mediumConvertedFee,
					},
					{
						Level:        defaultLevelFees[2].Level,
						Creator:      defaultLevelFees[2].Creator,
						OriginalFee:  defaultLevelFees[2].Fee,
						ConvertedFee: &highConvertedFee,
					},
					{
						Level:        levelZero.Level,
						OriginalFee:  levelZero.Fee,
						ConvertedFee: &zeroConvertedFee,
					},
					{
						Level:        levelMin.Level,
						OriginalFee:  levelMin.Fee,
						ConvertedFee: &minConvertedFee,
					},
				},
			},
		},
		{
			Name:      "gRPC action level fee ok",
			URL:       fmt.Sprintf("%s/shareledger/gentlemint/actionLevelFee/%s", baseURL, "gentlemint_load"),
			Headers:   map[string]string{},
			ExpectErr: false,
			RespType:  &types.QueryActionLevelFeeResponse{},
			Expected: &types.QueryActionLevelFeeResponse{
				Action: defaultActionLevelFees[0].Action,
				Level:  defaultActionLevelFees[0].Level,
				Fee:    feeLevelLow,
			},
		},
		{
			Name:      "gRPC invalid action level fee",
			URL:       fmt.Sprintf("%s/shareledger/gentlemint/actionLevelFee/%s", baseURL, "test"),
			Headers:   map[string]string{},
			ExpectErr: true,
		},
		{
			Name:         "gRPC list action level fee ok",
			URL:          fmt.Sprintf("%s/shareledger/gentlemint/actionLevelFee", baseURL),
			Headers:      map[string]string{},
			ExpectErr:    false,
			CheckContain: true,
			RespType:     &types.QueryActionLevelFeesResponse{},
			Expected: &types.QueryActionLevelFeesResponse{
				// ignore the last item that used to test delete
				ActionLevelFee: defaultActionLevelFees[:2],
			},
		},
		{
			// no data will be test here due to invalid request
			Name:      "gRPC checkfees ok",
			URL:       fmt.Sprintf("%s/shareledger/gentlemint/checkFees", baseURL),
			Headers:   map[string]string{},
			ExpectErr: true,
		},
		{
			// no data will be test here due to invalid request
			Name:      "gRPC list balances",
			URL:       fmt.Sprintf("%s/shareledger/gentlemint/balances", baseURL),
			Headers:   map[string]string{},
			ExpectErr: true,
		},
	}
	tests.RunTestCasesGrpc(&s.Suite, testCases, val)
}
