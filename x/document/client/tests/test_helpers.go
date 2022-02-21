package tests

import (
	"encoding/json"
	"testing"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/testutil"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"

	"github.com/sharering/shareledger/testutil/network"
	"github.com/sharering/shareledger/x/document/client/cli"
	"github.com/sharering/shareledger/x/document/types"
)

func CmdExCreateDocument(clientCtx client.Context, t *testing.T, holderID, docProof, extraData string, extraFlags ...string) testutil.BufferWriter {
	args := []string{holderID, docProof, extraData}
	args = append(args, extraFlags...)
	out, err := clitestutil.ExecTestCLICmd(clientCtx, cli.CmdCreateDocument(), args)
	if err != nil {
		t.Errorf("fail create document: %v", err)
	}
	return out
}

func CmdExCreateDocumentInBatch(clientCtx client.Context, t *testing.T, holderID, docProof, extraData string, extraFlags ...string) testutil.BufferWriter {
	args := []string{holderID, docProof, extraData}
	args = append(args, extraFlags...)
	out, err := clitestutil.ExecTestCLICmd(clientCtx, cli.CmdCreateDocuments(), args)
	if err != nil {
		t.Errorf("fail create document: %v", err)
	}
	return out
}

func CmdExUpdateDocument(clientCtx client.Context, t *testing.T, holderID, docProof, extraData string, extraFlags ...string) testutil.BufferWriter {
	args := []string{holderID, docProof, extraData}
	args = append(args, extraFlags...)
	out, err := clitestutil.ExecTestCLICmd(clientCtx, cli.CmdUpdateDocument(), args)
	if err != nil {
		t.Errorf("fail update document: %v", err)
	}
	return out
}

func CmdExRevokeDocument(clientCtx client.Context, t *testing.T, holderID, docProof string, extraFlags ...string) testutil.BufferWriter {
	args := []string{holderID, docProof}
	args = append(args, extraFlags...)
	out, err := clitestutil.ExecTestCLICmd(clientCtx, cli.CmdRevokeDocument(), args)
	if err != nil {
		t.Errorf("fail revoke document: %v", err)
	}
	return out
}

func CmdExGetDocByHolderID(clientCtx client.Context, t *testing.T, holderID string, extraFlags ...string) types.QueryDocumentByHolderIdResponse {
	args := []string{holderID}
	args = append(args, extraFlags...)
	args = append(args, network.JSONFlag())
	out, err := clitestutil.ExecTestCLICmd(clientCtx, cli.CmdDocumentByHolderId(), args)
	if err != nil {
		t.Errorf("fail get document by holderID: %v", err)
	}

	return DocumentsResponseUnMarshal(t, out.Bytes())
}
func CmdExGetDocByProof(clientCtx client.Context, t *testing.T, proof string, extraFlags ...string) types.QueryDocumentByProofResponse {
	args := []string{proof}
	args = append(args, extraFlags...)
	args = append(args, network.JSONFlag())
	out, err := clitestutil.ExecTestCLICmd(clientCtx, cli.CmdDocumentByProof(), args)
	if err != nil {
		t.Errorf("fail get document by proof: %v", err)
	}

	return DocumentResponseUnMarshal(t, out.Bytes())
}

func DocumentResponseUnMarshal(t *testing.T, bs []byte) (res types.QueryDocumentByProofResponse) {
	err := json.Unmarshal(bs, &res)
	if err != nil {
		t.Errorf("unmarshall fail %v", err)
	}
	return
}
func DocumentsResponseUnMarshal(t *testing.T, bs []byte) (res types.QueryDocumentByHolderIdResponse) {
	err := json.Unmarshal(bs, &res)
	if err != nil {
		t.Errorf("unmarshall fail %v", err)
	}
	return
}
