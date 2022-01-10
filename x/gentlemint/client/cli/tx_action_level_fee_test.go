package cli_test

import (
	"fmt"
    "strconv"
	"testing"

	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/sharering/shareledger/testutil/network"
	"github.com/sharering/shareledger/x/gentlemint/client/cli"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestCreateActionLevelFee(t *testing.T) {
	net := network.New(t)
	val := net.Validators[0]
	ctx := val.ClientCtx

    fields := []string{ "xyz", }
	for _, tc := range []struct {
		desc string
        idAction string
        
		args []string
		err  error
		code uint32
	}{
		{
            idAction: strconv.Itoa(0),
            
			desc: "valid",
			args: []string{
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(net.Config.BondDenom, sdk.NewInt(10))).String()),
			},
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
            args := []string{
                tc.idAction,
                
            }
            args = append(args, fields...)
            args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdCreateActionLevelFee(), args)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				var resp sdk.TxResponse
				require.NoError(t, ctx.JSONCodec.UnmarshalJSON(out.Bytes(), &resp))
				require.Equal(t, tc.code, resp.Code)
			}
		})
	}
}

func TestUpdateActionLevelFee(t *testing.T) {
	net := network.New(t)
	val := net.Validators[0]
	ctx := val.ClientCtx

    fields := []string{ "xyz", }
	common := []string{
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(net.Config.BondDenom, sdk.NewInt(10))).String()),
	}
    args := []string{
        "0",
        
    }
	args = append(args, fields...)
	args = append(args, common...)
	_, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdCreateActionLevelFee(), args)
	require.NoError(t, err)

	for _, tc := range []struct {
		desc string
		idAction string
        
		args []string
		code uint32
		err  error
	}{
		{
			desc: "valid",
			idAction: strconv.Itoa(0),
            
			args: common,
		},
		{
			desc: "key not found",
			idAction: strconv.Itoa(100000),
            
			args: common,
			code: sdkerrors.ErrKeyNotFound.ABCICode(),
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
            args := []string{
                tc.idAction,
                
            }
            args = append(args, fields...)
            args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdUpdateActionLevelFee(), args)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				var resp sdk.TxResponse
				require.NoError(t, ctx.JSONCodec.UnmarshalJSON(out.Bytes(), &resp))
				require.Equal(t, tc.code, resp.Code)
			}
		})
	}
}

func TestDeleteActionLevelFee(t *testing.T) {
	net := network.New(t)

	val := net.Validators[0]
	ctx := val.ClientCtx

	fields := []string{ "xyz", }
	common := []string{
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(net.Config.BondDenom, sdk.NewInt(10))).String()),
	}
    args := []string{
        "0",
        
    }
	args = append(args, fields...)
	args = append(args, common...)
	_, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdCreateActionLevelFee(), args)
	require.NoError(t, err)

	for _, tc := range []struct {
		desc string
		idAction string
        
		args []string
		code uint32
		err  error
	}{
		{
			desc: "valid",
			idAction: strconv.Itoa(0),
            
			args: common,
		},
		{
			desc: "key not found",
			idAction: strconv.Itoa(100000),
            
			args: common,
			code: sdkerrors.ErrKeyNotFound.ABCICode(),
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
		    args := []string{
                tc.idAction,
                
            }
            args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdDeleteActionLevelFee(), args)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				var resp sdk.TxResponse
				require.NoError(t, ctx.JSONCodec.UnmarshalJSON(out.Bytes(), &resp))
				require.Equal(t, tc.code, resp.Code)
			}
		})
	}
}
