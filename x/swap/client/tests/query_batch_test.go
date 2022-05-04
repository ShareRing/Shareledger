package tests

//func networkWithBatchObjects(t *testing.T, n int) (*network.Network, []types.Batch) {
//	t.Helper()
//	state := types.GenesisState{}
//	require.NoError(t, NetConf.Codec.UnmarshalJSON(NetConf.GenesisState[types.ModuleName], &state))
//
//	for i := 0; i < n; i++ {
//		batch := types.Batch{
//			Id: uint64(i),
//		}
//		nullify.Fill(&batch)
//		state.BatchList = append(state.BatchList, batch)
//	}
//	buf, err := NetConf.Codec.MarshalJSON(&state)
//	require.NoError(t, err)
//	delete(NetConf.GenesisState, types.ModuleName)
//	NetConf.GenesisState[types.ModuleName] = buf
//	return network.New(t, NetConf), state.BatchList
//}
//
//func TestShowBatch(t *testing.T) {
//	net, objs := networkWithBatchObjects(t, 2)
//	defer net.Cleanup()
//	ctx := net.Validators[0].ClientCtx
//	common := []string{
//		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
//	}
//	for _, tc := range []struct {
//		desc string
//		id   string
//		args []string
//		err  error
//		obj  types.Batch
//	}{
//		{
//			desc: "found",
//			id:   fmt.Sprintf("%d", objs[0].Id),
//			args: common,
//			obj:  objs[0],
//		},
//		{
//			desc: "not found",
//			id:   "not_found",
//			args: common,
//			err:  status.Error(codes.NotFound, "not found"),
//		},
//	} {
//		tc := tc
//		t.Run(tc.desc, func(t *testing.T) {
//			args := []string{tc.id}
//			args = append(args, tc.args...)
//			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowBatch(), args)
//			if tc.err != nil {
//				stat, ok := status.FromError(tc.err)
//				require.True(t, ok)
//				require.ErrorIs(t, stat.Err(), tc.err)
//			} else {
//				require.NoError(t, err)
//				var resp types.QueryGetBatchResponse
//				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
//				require.NotNil(t, resp.Batch)
//				require.Equal(t,
//					nullify.Fill(&tc.obj),
//					nullify.Fill(&resp.Batch),
//				)
//			}
//		})
//	}
//}
