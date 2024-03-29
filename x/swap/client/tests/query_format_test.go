package tests

import (
	"strconv"
)

// Prevent strconv unused error
var _ = strconv.IntSize

//func networkWithFormatObjects(t *testing.T, n int) (*network.Network, []types.Format) {
//	t.Helper()
//	state := types.GenesisState{}
//	require.NoError(t, NetConf.Codec.UnmarshalJSON(NetConf.GenesisState[types.ModuleName], &state))
//
//	for i := 0; i < n; i++ {
//		format := types.Format{
//			Network: strconv.Itoa(i),
//		}
//		nullify.Fill(&format)
//		state.FormatList = append(state.FormatList, format)
//	}
//	buf, err := NetConf.Codec.MarshalJSON(&state)
//	require.NoError(t, err)
//	delete(NetConf.GenesisState, types.ModuleName)
//	NetConf.GenesisState[types.ModuleName] = buf
//	return network.New(t, NetConf), state.FormatList
//}
//
//func TestShowFormat(t *testing.T) {
//	net, objs := networkWithFormatObjects(t, 2)
//	defer net.Cleanup()
//	ctx := net.Validators[0].ClientCtx
//	common := []string{
//		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
//	}
//	for _, tc := range []struct {
//		desc      string
//		idNetwork string
//
//		args []string
//		err  error
//		obj  types.Format
//	}{
//		{
//			desc:      "found",
//			idNetwork: objs[0].Network,
//
//			args: common,
//			obj:  objs[0],
//		},
//		{
//			desc:      "not found",
//			idNetwork: strconv.Itoa(100000),
//
//			args: common,
//			err:  status.Error(codes.NotFound, "not found"),
//		},
//	} {
//		tc := tc
//		t.Run(tc.desc, func(t *testing.T) {
//			args := []string{
//				tc.idNetwork,
//			}
//			args = append(args, tc.args...)
//			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowFormat(), args)
//			if tc.err != nil {
//				stat, ok := status.FromError(tc.err)
//				require.True(t, ok)
//				require.ErrorIs(t, stat.Err(), tc.err)
//			} else {
//				require.NoError(t, err)
//				var resp types.QueryGetFormatResponse
//				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
//				require.NotNil(t, resp.Format)
//				require.Equal(t,
//					nullify.Fill(&tc.obj),
//					nullify.Fill(&resp.Format),
//				)
//			}
//		})
//	}
//}
//
//func TestListFormat(t *testing.T) {
//	net, objs := networkWithFormatObjects(t, 5)
//	defer net.Cleanup()
//	ctx := net.Validators[0].ClientCtx
//	request := func(next []byte, offset, limit uint64, total bool) []string {
//		args := []string{
//			fmt.Sprintf("--%s=json", tmcli.OutputFlag),
//		}
//		if next == nil {
//			args = append(args, fmt.Sprintf("--%s=%d", flags.FlagOffset, offset))
//		} else {
//			args = append(args, fmt.Sprintf("--%s=%s", flags.FlagPageKey, next))
//		}
//		args = append(args, fmt.Sprintf("--%s=%d", flags.FlagLimit, limit))
//		if total {
//			args = append(args, fmt.Sprintf("--%s", flags.FlagCountTotal))
//		}
//		return args
//	}
//	t.Run("ByOffset", func(t *testing.T) {
//		step := 2
//		for i := 0; i < len(objs); i += step {
//			args := request(nil, uint64(i), uint64(step), false)
//			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListFormat(), args)
//			require.NoError(t, err)
//			var resp types.QueryAllFormatResponse
//			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
//			require.LessOrEqual(t, len(resp.Format), step)
//			require.Subset(t,
//				nullify.Fill(objs),
//				nullify.Fill(resp.Format),
//			)
//		}
//	})
//	t.Run("ByKey", func(t *testing.T) {
//		step := 2
//		var next []byte
//		for i := 0; i < len(objs); i += step {
//			args := request(next, 0, uint64(step), false)
//			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListFormat(), args)
//			require.NoError(t, err)
//			var resp types.QueryAllFormatResponse
//			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
//			require.LessOrEqual(t, len(resp.Format), step)
//			require.Subset(t,
//				nullify.Fill(objs),
//				nullify.Fill(resp.Format),
//			)
//			next = resp.Pagination.NextKey
//		}
//	})
//	t.Run("Total", func(t *testing.T) {
//		args := request(nil, 0, uint64(len(objs)), true)
//		out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListFormat(), args)
//		require.NoError(t, err)
//		var resp types.QueryAllFormatResponse
//		require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
//		require.NoError(t, err)
//		require.Equal(t, len(objs), int(resp.Pagination.Total))
//		require.ElementsMatch(t,
//			nullify.Fill(objs),
//			nullify.Fill(resp.Format),
//		)
//	})
//}
