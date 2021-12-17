package gentlemint_test

//func TestGenesis(t *testing.T) {
//	genesisState := types.GenesisState{
//		ExchangeRate: &types.ExchangeRate{
//			Rate: 200,
//		},
//		// this line is used by starport scaffolding # genesis/test/state
//	}
//
//	k, ctx := keepertest.GentlemintKeeper(t)
//	gentlemint.InitGenesis(ctx, *k, genesisState)
//	got := gentlemint.ExportGenesis(ctx, *k)
//	require.NotNil(t, got)
//
//	require.Len(t, got.AccStateList, len(genesisState.AccStateList))
//	require.Subset(t, genesisState.AccStateList, got.AccStateList)
//	require.Equal(t, genesisState.Authority, got.Authority)
//	require.Equal(t, genesisState.Treasurer, got.Treasurer)
//	require.Equal(t, genesisState.ExchangeRate, got.ExchangeRate)
//	// this line is used by starport scaffolding # genesis/test/assert
//}
