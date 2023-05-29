package testutil

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authcli "github.com/cosmos/cosmos-sdk/x/auth/client/cli"
	"github.com/spf13/cobra"
)

func PtrOf[T any](i T) *T {
	return &i
}

func ExecTestCLICmdBlocked(clientCtx client.Context, cmd *cobra.Command, extraArgs []string) (*sdk.TxResponse, error) {
	out, err := cli.ExecTestCLICmd(clientCtx, cmd, extraArgs)
	if err != nil {
		return nil, err
	}
	resp := &sdk.TxResponse{}
	err = clientCtx.Codec.UnmarshalJSON(out.Bytes(), resp)
	if err != nil {
		return nil, err
	}

	queryCmd := authcli.QueryTxCmd()
	retry := func() error {
		time.Sleep(time.Second)
		out, err = cli.ExecTestCLICmd(clientCtx, queryCmd, []string{resp.TxHash, fmt.Sprintf("--%s=json", flags.FlagOutput)})
		if err != nil {
			return err
		}
		err = clientCtx.Codec.UnmarshalJSON(out.Bytes(), resp)
		fmt.Printf("resp: %v\n", resp)
		return err
	}
	// poll response 30 times
	for i := 0; i < 30; i++ {
		err = retry()
		if err == nil {
			return resp, nil
		}
	}
	return resp, err
}
