package tests

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/tests"
	sdk "github.com/cosmos/cosmos-sdk/types"
	app "github.com/sharering/shareledger"
	"github.com/sharering/shareledger/x/electoral"
	"github.com/sharering/shareledger/x/electoral/types"
	"github.com/stretchr/testify/require"
)

const ElectoralModuleName = electoral.ModuleName


func (f *Fixtures) ExecuteElectoralEnrollVoter(userKey string ,addressForEnroll sdk.Address ) (bool, string, string) {
	flag := []string{fmt.Sprintf("--key-seed ./%s_key_seed.json --yes --fees 1shr", userKey)}
	cmd := fmt.Sprintf("%s tx %v enroll %v %v", f.GaiacliBinary,ElectoralModuleName, addressForEnroll.String(), f.Flags())
	return executeWriteRetStdStreams(f.T,addFlags(cmd,flag), DefaultKeyPass)
}

func (f *Fixtures) ExecuteElectoralGetVoter(targetAddr sdk.Address ) types.Voter {
	cmd := fmt.Sprintf("%s query %v get %s %v", f.GaiacliBinary,ElectoralModuleName, targetAddr, f.Flags())
	out, _ := tests.ExecuteT(f.T, cmd, "")

	var voter types.Voter
	cdc := app.MakeCodec()
	err := cdc.UnmarshalJSON([]byte(out), &voter)
	require.NoError(f.T, err, "out %v\n, err %v", out, err)
	return voter
}


func (f *Fixtures) ExecuteElectoralRevokeVoter(userKey string ,addressForRevoke sdk.Address ) (bool, string, string) {
	flag := []string{fmt.Sprintf("--key-seed ./%s_key_seed.json --yes --fees 1shr", userKey)}

	cmd := fmt.Sprintf("%s tx %v revoke %v %v", f.GaiacliBinary,ElectoralModuleName, addressForRevoke.String(), f.Flags())
	return executeWriteRetStdStreams(f.T,addFlags(cmd,flag), DefaultKeyPass)
}
