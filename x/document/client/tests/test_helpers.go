package tests

import (
	"encoding/json"
	"github.com/ShareRing/Shareledger/testutil/network"
	"github.com/ShareRing/Shareledger/x/document/client/cli"
	"github.com/ShareRing/Shareledger/x/document/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/testutil"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"testing"
)

func CmdExCreateDocument(clientCtx client.Context,t *testing.T,holderID,docProof,extraData string,extraFlags ...string)testutil.BufferWriter{
	args := []string{holderID,docProof,extraData}
	args = append(args,extraFlags...)
	out,err := clitestutil.ExecTestCLICmd(clientCtx,cli.NewDocumentTxCmd(),args)
	if err!= nil{
		t.Errorf("fail create document: %v",err)
	}
	return out
}

func CmdExCreateDocumentInBatch(clientCtx client.Context,t *testing.T,holderID,docProof,extraData string,extraFlags ...string)testutil.BufferWriter{
	args := []string{holderID,docProof,extraData}
	args = append(args,extraFlags...)
	out,err := clitestutil.ExecTestCLICmd(clientCtx,cli.NewDocumentInBatchTxCmd(),args)
	if err!= nil{
		t.Errorf("fail create document: %v",err)
	}
	return out
}

func CmdExUpdateDocument(clientCtx client.Context,t *testing.T,holderID,docProof,extraData string,extraFlags ...string)testutil.BufferWriter{
	args := []string{holderID,docProof,extraData}
	args = append(args,extraFlags...)
	out,err := clitestutil.ExecTestCLICmd(clientCtx,cli.UpdateDocCmd(),args)
	if err!= nil{
		t.Errorf("fail update document: %v",err)
	}
	return out
}

func CmdExRevokeDocument(clientCtx client.Context,t *testing.T,holderID,docProof string,extraFlags ...string)testutil.BufferWriter{
	args := []string{holderID,docProof}
	args = append(args,extraFlags...)
	out,err := clitestutil.ExecTestCLICmd(clientCtx,cli.RevokeDocCmd(),args)
	if err!= nil{
		t.Errorf("fail revoke document: %v",err)
	}
	return out
}

func CmdExGetDocByHolderID(clientCtx client.Context,t *testing.T,holderID string,extraFlags ...string)types.QueryDocumentByHolderIdResponse{
	args := []string{holderID}
	args = append(args,extraFlags...)
	args = append(args,network.JSONFlag())
	out,err := clitestutil.ExecTestCLICmd(clientCtx,cli.GetDocByHolderCmd(""),args)
	if err!= nil{
		t.Errorf("fail get document by holderID: %v",err)
	}

	return DocumentsResponseUnMarshal(t,out.Bytes())
}
func CmdExGetDocByProof(clientCtx client.Context,t *testing.T,proof string,extraFlags ...string)types.QueryDocumentByProofResponse{
	args := []string{proof}
	args = append(args,extraFlags...)
	args = append(args,network.JSONFlag())
	out,err := clitestutil.ExecTestCLICmd(clientCtx,cli.GetDocByProofCmd(""),args)
	if err!= nil{
		t.Errorf("fail get document by proof: %v",err)
	}

	return DocumentResponseUnMarshal(t,out.Bytes())
}


func DocumentResponseUnMarshal(t *testing.T,bs []byte)(res types.QueryDocumentByProofResponse){
	err :=json.Unmarshal(bs,&res)
	if err!= nil{
		t.Errorf("unmarshall fail %v",err)
	}
	return
}
func DocumentsResponseUnMarshal(t *testing.T,bs []byte)(res types.QueryDocumentByHolderIdResponse){
	err :=json.Unmarshal(bs,&res)
	if err!= nil{
		t.Errorf("unmarshall fail %v",err)
	}
	return
}