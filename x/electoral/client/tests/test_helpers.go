package tests

import (
	"github.com/ShareRing/Shareledger/x/electoral/client/cli"
	"testing"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/testutil"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
)

func ExCmdEnrollDocIssuer(clientCtx client.Context, t *testing.T,address string,additionalFlags ...string) testutil.BufferWriter {
	args := []string{address}
	args = append(args, additionalFlags...)
	out,err :=clitestutil.ExecTestCLICmd(clientCtx,cli.CmdEnrollDocIssuer() , args)
	if err!= nil{
		t.Errorf("fail enroll doc_issuer %v",err)
	}
	return out
}
func ExCmdRevokeDocIssuer(clientCtx client.Context, t *testing.T,address string,additionalFlags ...string) testutil.BufferWriter {
	args := []string{address}
	args = append(args, additionalFlags...)
	out,err :=clitestutil.ExecTestCLICmd(clientCtx,cli.CmdRevokeDocIssuer() , args)
	if err!= nil{
		t.Errorf("fail revoke doc_issuer %v",err)
	}
	return out
}
func ExCmdGetDocIssuers(clientCtx client.Context, t *testing.T,additionalFlags ...string) testutil.BufferWriter {
	var args []string
	args = append(args, additionalFlags...)
	out,err :=clitestutil.ExecTestCLICmd(clientCtx,cli.CmdDocumentIssuers() , args)
	if err!= nil{
		t.Errorf("get list doc_issuer %v",err)
	}
	return out
}


func ExCmdEnrollAccountOperator(clientCtx client.Context,t *testing.T,address string,additionalFlags ...string) testutil.BufferWriter {
	args := []string{address}
	args = append(args, additionalFlags...)
	out,err :=clitestutil.ExecTestCLICmd(clientCtx,cli.CmdEnrollAccountOperator() , args)
	if err!= nil{
		t.Errorf("fail enroll account_operator %v",err)
	}
	return out
}
func ExCmdRevokeAccountOperator(clientCtx client.Context, t *testing.T,address string,additionalFlags ...string) testutil.BufferWriter {
	args := []string{address}
	args = append(args, additionalFlags...)
	out,err :=clitestutil.ExecTestCLICmd(clientCtx,cli.CmdRevokeAccountOperator() , args)
	if err!= nil{
		t.Errorf("fail revoke AccountOperator %v",err)
	}
	return out
}

func ExCmdEnrollIdSigner(clientCtx client.Context,t  *testing.T,address string,additionalFlags ...string) testutil.BufferWriter {
	args := []string{address}
	args = append(args, additionalFlags...)
	out,err :=clitestutil.ExecTestCLICmd(clientCtx,cli.CmdEnrollIdSigner() , args)
	if err!= nil{
		t.Errorf("fail enroll id singer %v",err)
	}
	return out
}
func ExCmdRevokeIdSigner(clientCtx client.Context,t  *testing.T,address string,additionalFlags ...string) testutil.BufferWriter {
	args := []string{address}
	args = append(args, additionalFlags...)
	out,err :=clitestutil.ExecTestCLICmd(clientCtx,cli.CmdRevokeIdSigner() , args)
	if err!= nil{
		t.Errorf("fail revoke id singer %v",err)
	}
	return out
}

func ExCmdEnrollLoader(clientCtx client.Context,t  *testing.T,address string,additionalFlags ...string) testutil.BufferWriter {
	args := []string{address}
	args = append(args, additionalFlags...)
	out,err :=clitestutil.ExecTestCLICmd(clientCtx,cli.CmdEnrollLoaders() , args)
	if err!= nil{
		t.Errorf("fail enroll loader singer %v",err)
	}
	return out
}
func ExCmdRevokeLoader(clientCtx client.Context,t  *testing.T,address string,additionalFlags ...string) testutil.BufferWriter {
	args := []string{address}
	args = append(args, additionalFlags...)
	out,err :=clitestutil.ExecTestCLICmd(clientCtx,cli.CmdRevokeLoaders() , args)
	if err!= nil{
		t.Errorf("fail revoke loader singer %v",err)
	}
	return out
}

func ExCmdEnrollVoter(clientCtx client.Context,t  *testing.T,address string,additionalFlags ...string) testutil.BufferWriter {
	args := []string{address}
	args = append(args, additionalFlags...)
	out,err :=clitestutil.ExecTestCLICmd(clientCtx,cli.CmdEnrollVoter() , args)
	if err!= nil{
		t.Errorf("fail enroll loader singer %v",err)
	}
	return out
}
func ExCmdRevokeVoter(clientCtx client.Context,t  *testing.T,address string,additionalFlags ...string) testutil.BufferWriter {
	args := []string{address}
	args = append(args, additionalFlags...)
	out,err :=clitestutil.ExecTestCLICmd(clientCtx,cli.CmdRevokeVoter() , args)
	if err!= nil{
		t.Errorf("fail revoke voter singer %v",err)
	}
	return out
}