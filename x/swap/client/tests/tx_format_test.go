package tests

import (
	"strconv"
)

// Prevent strconv unused error
var _ = strconv.IntSize

//func TestCreateFormat(t *testing.T) {
//	delete(NetConf.GenesisState, types.ModuleName)
//	net := network.New(t, NetConf)
//	defer net.Cleanup()
//	val := net.Validators[0]
//	ctx := val.ClientCtx
//
//	fields := []string{}
//	for _, tc := range []struct {
//		desc      string
//		idNetwork string
//
//		args []string
//		err  error
//		code uint32
//	}{
//		{
//			idNetwork: strconv.Itoa(0),
//			desc:      "valid",
//			args: []string{
//				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
//				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
//				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
//				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(denom.Shr, sdk.NewInt(10))).String()),
//			},
//		},
//	} {
//		tc := tc
//		t.Run(tc.desc, func(t *testing.T) {
//			args := []string{
//				tc.idNetwork,
//				"data",
//			}
//			args = append(args, fields...)
//			args = append(args, tc.args...)
//			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdCreateFormat(), args)
//			if tc.err != nil {
//				require.ErrorIs(t, err, tc.err)
//			} else {
//				require.NoError(t, err)
//				var resp sdk.TxResponse
//				require.NoError(t, ctx.Codec.UnmarshalJSON(out.Bytes(), &resp))
//				require.Equal(t, tc.code, resp.Code)
//			}
//		})
//	}
//}
//
//func TestUpdateFormat(t *testing.T) {
//	delete(NetConf.GenesisState, types.ModuleName)
//	net := network.New(t, NetConf)
//	defer net.Cleanup()
//	val := net.Validators[0]
//	ctx := val.ClientCtx
//
//	fields := []string{}
//	common := []string{
//		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
//		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
//		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
//		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(net.Config.BondDenom, sdk.NewInt(10))).String()),
//	}
//	args := []string{
//		"0",
//	}
//	args = append(args, fields...)
//	args = append(args, common...)
//	_, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdCreateFormat(), args)
//	require.NoError(t, err)
//
//	for _, tc := range []struct {
//		desc      string
//		idNetwork string
//
//		args []string
//		code uint32
//		err  error
//	}{
//		{
//			desc:      "valid",
//			idNetwork: strconv.Itoa(0),
//
//			args: common,
//		},
//		{
//			desc:      "key not found",
//			idNetwork: strconv.Itoa(100000),
//
//			args: common,
//			code: sdkerrors.ErrKeyNotFound.ABCICode(),
//		},
//	} {
//		tc := tc
//		t.Run(tc.desc, func(t *testing.T) {
//			args := []string{
//				tc.idNetwork,
//			}
//			args = append(args, fields...)
//			args = append(args, tc.args...)
//			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdUpdateFormat(), args)
//			if tc.err != nil {
//				require.ErrorIs(t, err, tc.err)
//			} else {
//				require.NoError(t, err)
//				var resp sdk.TxResponse
//				require.NoError(t, ctx.Codec.UnmarshalJSON(out.Bytes(), &resp))
//				require.Equal(t, tc.code, resp.Code)
//			}
//		})
//	}
//}
//
//func TestDeleteFormat(t *testing.T) {
//	delete(NetConf.GenesisState, types.ModuleName)
//	net := network.New(t, NetConf)
//	defer net.Cleanup()
//	val := net.Validators[0]
//	ctx := val.ClientCtx
//
//	fields := []string{}
//	common := []string{
//		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
//		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
//		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
//		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(denom.Shr, sdk.NewInt(10))).String()),
//	}
//	args := []string{
//		"0",
//		"data",
//	}
//	args = append(args, fields...)
//	args = append(args, common...)
//	_, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdCreateFormat(), args)
//	require.NoError(t, err)
//
//	for _, tc := range []struct {
//		desc      string
//		idNetwork string
//
//		args []string
//		code uint32
//		err  error
//	}{
//		{
//			desc:      "valid",
//			idNetwork: strconv.Itoa(0),
//
//			args: common,
//		},
//		{
//			desc:      "key not found",
//			idNetwork: strconv.Itoa(100000),
//
//			args: common,
//			code: sdkerrors.ErrKeyNotFound.ABCICode(),
//		},
//	} {
//		tc := tc
//		t.Run(tc.desc, func(t *testing.T) {
//			args := []string{
//				tc.idNetwork,
//			}
//			args = append(args, tc.args...)
//			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdDeleteFormat(), args)
//			if tc.err != nil {
//				require.ErrorIs(t, err, tc.err)
//			} else {
//				require.NoError(t, err)
//				var resp sdk.TxResponse
//				require.NoError(t, ctx.Codec.UnmarshalJSON(out.Bytes(), &resp))
//				require.Equal(t, tc.code, resp.Code)
//			}
//		})
//	}
//}
