package document

import (
	"fmt"

	"github.com/sharering/shareledger/tests"
	"github.com/sharering/shareledger/x/document/types"
)

func (s *E2ETestSuite) TestGRPCQueryDocumentByHolderId() {
	val := s.network.Validators[0]
	baseURL := val.APIAddress
	testCases := tests.TestCasesGrpc{
		{
			Name:      "gRPC document by holder id ok",
			URL:       fmt.Sprintf("%s/shareledger/document/%s", baseURL, firstDoc.Holder),
			Headers:   map[string]string{},
			ExpectErr: false,
			RespType:  &types.QueryDocumentByHolderIdResponse{},
			Expected: &types.QueryDocumentByHolderIdResponse{
				Documents: []*types.Document{&firstDoc},
			},
		},
		{
			Name:      "gRPC document by empty holder id",
			URL:       fmt.Sprintf("%s/shareledger/document/", baseURL),
			Headers:   map[string]string{},
			ExpectErr: true,
			RespType:  nil,
			Expected:  nil,
		},
		{
			Name:      "gRPC id by holder id long exceeds threshold length",
			URL:       fmt.Sprintf("%s/shareledger/document/%s", baseURL, "KFhnI60lCMlwL1gtu2nwZKyNsCzd42eXodt9hmRrBsf4Y1L4fNasGSRibI7geMzcX"),
			Headers:   map[string]string{},
			ExpectErr: true,
			RespType:  nil,
			Expected:  nil,
		},
		{
			Name:      "gRPC id by holder id long exceeds threshold length",
			URL:       fmt.Sprintf("%s/shareledger/document/%s", baseURL, "invalid_id"),
			Headers:   map[string]string{},
			ExpectErr: true,
			RespType:  nil,
			Expected:  nil,
		},
		{
			Name:      "gRPC document by proof ok",
			URL:       fmt.Sprintf("%s/shareledger/document/proof/%s", baseURL, firstDoc.Proof),
			Headers:   map[string]string{},
			ExpectErr: false,
			RespType:  &types.QueryDocumentByProofResponse{},
			Expected: &types.QueryDocumentByProofResponse{
				Document: &firstDoc,
			},
		},
		{
			Name:      "gRPC document by empty proof",
			URL:       fmt.Sprintf("%s/shareledger/document/proof", baseURL),
			Headers:   map[string]string{},
			ExpectErr: true,
			RespType:  nil,
			Expected:  nil,
		},
		{
			Name:      "gRPC document by invalid proof",
			URL:       fmt.Sprintf("%s/shareledger/document/proof/%s", baseURL, "invalid_id"),
			Headers:   map[string]string{},
			ExpectErr: true,
			RespType:  nil,
			Expected:  nil,
		},
		{
			Name:      "gRPC document of holder by issuer ok",
			URL:       fmt.Sprintf("%s/shareledger/document/%s/%s", baseURL, firstDoc.Holder, firstDoc.Issuer),
			Headers:   map[string]string{},
			ExpectErr: false,
			RespType:  &types.QueryDocumentOfHolderByIssuerResponse{},
			Expected: &types.QueryDocumentOfHolderByIssuerResponse{
				Documents: []*types.Document{&firstDoc},
			},
		},
		{
			Name:      "gRPC document with empty request",
			URL:       fmt.Sprintf("%s/shareledger/document/%s/%s", baseURL, "", ""),
			Headers:   map[string]string{},
			ExpectErr: true,
			RespType:  nil,
			Expected:  nil,
		},
		{
			Name:      "gRPC document with empty holder",
			URL:       fmt.Sprintf("%s/shareledger/document/%s/%s", baseURL, firstDoc.Holder, ""),
			Headers:   map[string]string{},
			ExpectErr: true,
			RespType:  nil,
			Expected:  nil,
		},
		{
			Name:      "gRPC document with empty issuer",
			URL:       fmt.Sprintf("%s/shareledger/document/%s/%s", baseURL, "", firstDoc.Issuer),
			Headers:   map[string]string{},
			ExpectErr: true,
			RespType:  nil,
			Expected:  nil,
		},
		{
			Name:      "gRPC document with invalid holder",
			URL:       fmt.Sprintf("%s/shareledger/document/proof/%s/%s", baseURL, "invalid_id", firstDoc.Issuer),
			Headers:   map[string]string{},
			ExpectErr: true,
			RespType:  nil,
			Expected:  nil,
		},
		{
			Name:      "gRPC document with invalid issuer",
			URL:       fmt.Sprintf("%s/shareledger/document/proof/%s/%s", baseURL, firstDoc.Holder, "invalid_id"),
			Headers:   map[string]string{},
			ExpectErr: true,
			RespType:  nil,
			Expected:  nil,
		},
	}
	tests.RunTestCasesGrpc(&s.Suite, testCases, val)
}
