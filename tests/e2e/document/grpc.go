package document

import (
	"fmt"

	"github.com/sharering/shareledger/tests"
	"github.com/sharering/shareledger/x/document/types"
)

const (
	BaseURIFormat     = "%s/shareledger/document/%s"
	ByIssuerURIFormat = "%s/shareledger/document/%s/%s"
	ByProofURIFormat  = "%s/shareledger/document/proof/%s"
)

func (s *E2ETestSuite) TestGRPCQueryDocumentByHolderId() {
	val := s.network.Validators[0]
	baseURL := val.APIAddress

	var documentByHolder = func(holder string) string {
		return fmt.Sprintf(BaseURIFormat, baseURL, holder)
	}
	var documentByProof = func(proof string) string {
		return fmt.Sprintf(ByProofURIFormat, baseURL, proof)
	}

	var documentByIssuer = func(holder, issuer string) string {
		return fmt.Sprintf(ByIssuerURIFormat, baseURL, holder, issuer)
	}

	testCases := tests.TestCasesGrpc{
		{
			Name: "gRPC document by holder id ok",
			//URL:       fmt.Sprintf("%s/shareledger/document/%s", baseURL, firstDoc.Holder),
			URL:       documentByHolder(firstDoc.Holder),
			Headers:   map[string]string{},
			ExpectErr: false,
			RespType:  &types.QueryDocumentByHolderIdResponse{},
			Expected: &types.QueryDocumentByHolderIdResponse{
				Documents: []*types.Document{&firstDoc},
			},
		},
		{
			Name: "gRPC document by empty holder id",
			//URL:       fmt.Sprintf("%s/shareledger/document/", baseURL),
			URL:       documentByHolder(""),
			Headers:   map[string]string{},
			ExpectErr: true,
			RespType:  nil,
			Expected:  nil,
		},
		{
			Name: "gRPC id by holder id long exceeds threshold length",
			//URL:       fmt.Sprintf("%s/shareledger/document/%s", baseURL, "KFhnI60lCMlwL1gtu2nwZKyNsCzd42eXodt9hmRrBsf4Y1L4fNasGSRibI7geMzcX"),
			URL:       documentByHolder("KFhnI60lCMlwL1gtu2nwZKyNsCzd42eXodt9hmRrBsf4Y1L4fNasGSRibI7geMzcX"),
			Headers:   map[string]string{},
			ExpectErr: true,
			RespType:  nil,
			Expected:  nil,
		},
		{
			Name: "gRPC id by holder id long exceeds threshold length",
			//URL:       fmt.Sprintf("%s/shareledger/document/%s", baseURL, "invalid_id"),
			URL:       documentByHolder("invalid_id"),
			Headers:   map[string]string{},
			ExpectErr: true,
			RespType:  nil,
			Expected:  nil,
		},
		{
			Name:      "gRPC document by proof ok",
			URL:       documentByProof(firstDoc.Proof),
			Headers:   map[string]string{},
			ExpectErr: false,
			RespType:  &types.QueryDocumentByProofResponse{},
			Expected: &types.QueryDocumentByProofResponse{
				Document: &firstDoc,
			},
		},
		{
			Name:      "gRPC document by empty proof",
			URL:       documentByProof(""),
			Headers:   map[string]string{},
			ExpectErr: true,
			RespType:  nil,
			Expected:  nil,
		},
		{
			Name: "gRPC document by invalid proof",
			//URL:       fmt.Sprintf("%s/shareledger/document/proof/%s", baseURL, "invalid_id"),
			URL:       documentByProof("invalid_id"),
			Headers:   map[string]string{},
			ExpectErr: true,
			RespType:  nil,
			Expected:  nil,
		},
		{
			Name: "gRPC document of holder by issuer ok",
			//URL:       fmt.Sprintf("%s/shareledger/document/%s/%s", baseURL, firstDoc.Holder, firstDoc.Issuer),
			URL:       documentByIssuer(firstDoc.Holder, firstDoc.Issuer),
			Headers:   map[string]string{},
			ExpectErr: false,
			RespType:  &types.QueryDocumentOfHolderByIssuerResponse{},
			Expected: &types.QueryDocumentOfHolderByIssuerResponse{
				Documents: []*types.Document{&firstDoc},
			},
		},
		{
			Name: "gRPC document with empty request",
			//URL:       fmt.Sprintf("%s/shareledger/document/%s/%s", baseURL, "", ""),
			URL:       documentByIssuer("", ""),
			Headers:   map[string]string{},
			ExpectErr: true,
			RespType:  nil,
			Expected:  nil,
		},
		{
			Name: "gRPC document with empty issuer",
			//URL:       fmt.Sprintf("%s/shareledger/document/%s/%s", baseURL, firstDoc.Holder, ""),
			URL:       documentByIssuer(firstDoc.Holder, ""),
			Headers:   map[string]string{},
			ExpectErr: true,
			RespType:  nil,
			Expected:  nil,
		},
		{
			Name: "gRPC document with empty holder",
			//URL:       fmt.Sprintf("%s/shareledger/document/%s/%s", baseURL, "", firstDoc.Issuer),
			URL:       documentByIssuer("", firstDoc.Issuer),
			Headers:   map[string]string{},
			ExpectErr: true,
			RespType:  nil,
			Expected:  nil,
		},
	}
	tests.RunTestCasesGrpc(&s.Suite, testCases, val)
}
