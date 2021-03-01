/*
 * Based on https://github.com/cosmos/gaia/blob/v2.0.12/cli_test/cli_test.go
 */
package tests

import (
	"fmt"
	"strings"

	doctypes "github.com/ShareRing/modules/document/types"
	"github.com/cosmos/cosmos-sdk/tests"
	app "github.com/sharering/shareledger"
	"github.com/stretchr/testify/require"
)

const ()

var ()

func (f *Fixtures) CreateNewDoc(holder, proof, data string, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("%s tx doc create %s %s %s %v", f.GaiacliBinary, holder, proof, data, f.Flags())
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), DefaultKeyPass)
}

func (f *Fixtures) CreateNewDocInBatch(holder, proof, data []string, flags ...string) (bool, string, string) {
	sep := ","
	cmd := fmt.Sprintf("%s tx doc create-batch %s %s %s %v", f.GaiacliBinary, strings.Join(holder, sep), strings.Join(proof, sep), strings.Join(data, sep), f.Flags())
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), DefaultKeyPass)
}

func (f *Fixtures) UdpateDoc(holder, proof, data string, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("%s tx doc update %s %s %s %v", f.GaiacliBinary, holder, proof, data, f.Flags())
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), DefaultKeyPass)
}

func (f *Fixtures) RevokeDoc(holder, proof string, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("%s tx doc revoke %s %s %v", f.GaiacliBinary, holder, proof, f.Flags())
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), DefaultKeyPass)
}

func (f *Fixtures) QueryDocByProof(proof string, flags ...string) doctypes.Doc {
	cmd := fmt.Sprintf("%s query doc proof %s %v", f.GaiacliBinary, proof, f.Flags())

	out, _ := tests.ExecuteT(f.T, addFlags(cmd, flags), "")
	var doc doctypes.Doc
	if len(out) == 0 {
		fmt.Println("STDOut is empty")
		return doc
	}

	cdc := app.MakeCodec()
	err := cdc.UnmarshalJSON([]byte(out), &doc)
	require.NoError(f.T, err, "out %v\n, err %v", out, err)
	return doc
}

func (f *Fixtures) QueryDocByHolder(holder string, flags ...string) []doctypes.Doc {
	cmd := fmt.Sprintf("%s query doc holder %s %v", f.GaiacliBinary, holder, f.Flags())

	out, _ := tests.ExecuteT(f.T, addFlags(cmd, flags), "")
	var docs []doctypes.Doc
	if len(out) == 0 {
		fmt.Println("STDOut is empty")
		return docs
	}

	cdc := app.MakeCodec()
	err := cdc.UnmarshalJSON([]byte(out), &docs)
	require.NoError(f.T, err, "out %v\n, err %v", out, err)
	return docs
}
