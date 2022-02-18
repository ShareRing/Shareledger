package types

//
//func TestAddShrpCoins(t *testing.T) {
//	type testCase struct {
//		ic sdk.Coins
//		ia sdk.Coins
//		oa AdjustmentCoins
//		oe error
//		d  string
//	}
//	testCases := []testCase{
//		{
//			ic: sdk.NewCoins(sdk.NewCoin(denom.ShrP, sdk.NewInt(1)), sdk.NewCoin(denom.BaseUSD, sdk.NewInt(0))),
//			ia: sdk.NewCoins(sdk.NewCoin(denom.ShrP, sdk.NewInt(1)), sdk.NewCoin(denom.BaseUSD, sdk.NewInt(0))),
//			oa: AdjustmentCoins{
//				Add: sdk.NewCoins(sdk.NewCoin(denom.ShrP, sdk.NewInt(1))),
//				Sub: sdk.NewCoins(sdk.NewCoin(denom.BaseUSD, sdk.NewInt(0))),
//			},
//			d: "1 shrp + 1 shrp = 2 shrp -> add: 1 shrp, sub 0 cent",
//		},
//		{
//			ic: sdk.NewCoins(sdk.NewCoin(denom.ShrP, sdk.NewInt(1)), sdk.NewCoin(denom.BaseUSD, sdk.NewInt(50))),
//			ia: sdk.NewCoins(sdk.NewCoin(denom.ShrP, sdk.NewInt(1)), sdk.NewCoin(denom.BaseUSD, sdk.NewInt(50))),
//			oa: AdjustmentCoins{
//				Add: sdk.NewCoins(sdk.NewCoin(denom.ShrP, sdk.NewInt(2))),
//				Sub: sdk.NewCoins(sdk.NewCoin(denom.BaseUSD, sdk.NewInt(50))),
//			},
//			d: "1.5 shrp + 1.5 shrp = 3 shrp -> add: 2 shrp, sub 50 cent",
//		},
//		{
//			ic: sdk.NewCoins(sdk.NewCoin(denom.ShrP, sdk.NewInt(1)), sdk.NewCoin(denom.BaseUSD, sdk.NewInt(50))),
//			ia: sdk.NewCoins(sdk.NewCoin(denom.ShrP, sdk.NewInt(1)), sdk.NewCoin(denom.BaseUSD, sdk.NewInt(70))),
//			oa: AdjustmentCoins{
//				Add: sdk.NewCoins(sdk.NewCoin(denom.ShrP, sdk.NewInt(2))),
//				Sub: sdk.NewCoins(sdk.NewCoin(denom.BaseUSD, sdk.NewInt(30))),
//			},
//			d: "1.5 shrp + 1.7 shrp = 3.2 shrp -> add: 1 shrp, sub 30 cent",
//		},
//		{
//			ic: sdk.NewCoins(sdk.NewCoin(denom.ShrP, sdk.NewInt(1)), sdk.NewCoin(denom.BaseUSD, sdk.NewInt(50))),
//			ia: sdk.NewCoins(sdk.NewCoin(denom.BaseUSD, sdk.NewInt(50))),
//			oa: AdjustmentCoins{
//				Add: sdk.NewCoins(sdk.NewCoin(denom.ShrP, sdk.NewInt(1))),
//				Sub: sdk.NewCoins(sdk.NewCoin(denom.BaseUSD, sdk.NewInt(50))),
//			},
//			d: "1.5 shrp + 0.5 shrp = 2 shrp -> add: 1 shrp, sub 5 cent",
//		},
//		{
//			ic: sdk.NewCoins(sdk.NewCoin(denom.ShrP, sdk.NewInt(1)), sdk.NewCoin(denom.BaseUSD, sdk.NewInt(50))),
//			ia: sdk.NewCoins(sdk.NewCoin(denom.BaseUSD, sdk.NewInt(70))),
//			oa: AdjustmentCoins{
//				Add: sdk.NewCoins(sdk.NewCoin(denom.ShrP, sdk.NewInt(1))),
//				Sub: sdk.NewCoins(sdk.NewCoin(denom.BaseUSD, sdk.NewInt(30))),
//			},
//			d: "1.5 shrp + 0.7 shrp = 2.2 -> add: 1 shrp, sub: 3 cent",
//		},
//		{
//			ic: sdk.NewCoins(sdk.NewCoin(denom.ShrP, sdk.NewInt(1)), sdk.NewCoin(denom.BaseUSD, sdk.NewInt(550))),
//			ia: sdk.NewCoins(sdk.NewCoin(denom.BaseUSD, sdk.NewInt(70))),
//			oa: AdjustmentCoins{
//				Add: sdk.NewCoins(sdk.NewCoin(denom.ShrP, sdk.NewInt(6))),
//				Sub: sdk.NewCoins(sdk.NewCoin(denom.BaseUSD, sdk.NewInt(530))),
//			},
//			d: "1.550 shrp + 0.7 shrp = 7.2 -> add: 6 shrp, sub: 530 cent",
//		},
//	}
//
//	for _, tc := range testCases {
//		r, e := AddShrpCoins(tc.ic, tc.ia)
//		if tc.oe != nil {
//			require.NotNil(t, e, tc.d)
//			require.True(t, sdkerrors.IsOf(e, tc.oe), tc.d)
//		} else {
//			require.Equal(t, tc.oa, r, tc.d)
//		}
//	}
//}
//
//func TestGetCostShrpForShr(t *testing.T) {
//	rate := sdk.NewDec(200)
//	type testCase struct {
//		iCurrent sdk.Coins
//		iNeed    sdk.Int
//		iRate    sdk.Dec
//		oCost    AdjustmentCoins
//		oError   error
//		d        string
//	}
//	testCases := []testCase{
//		{
//			iCurrent: sdk.NewCoins(
//				sdk.NewCoin(denom.ShrP, sdk.NewInt(1)),
//			),
//			iRate: rate,
//			iNeed: sdk.NewInt(200),
//			oCost: AdjustmentCoins{
//				Sub: sdk.NewCoins(
//					sdk.NewCoin(denom.ShrP, sdk.NewInt(1)),
//				),
//				Add: sdk.NewCoins(),
//			},
//			oError: nil,
//			d:      "1 shrp buy 200 base (1 shrp) -> cost -1 shrp",
//		},
//		{
//			iCurrent: sdk.NewCoins(
//				sdk.NewCoin(denom.ShrP, sdk.NewInt(1)),
//			),
//			iRate: rate,
//			iNeed: sdk.NewInt(200),
//			oCost: AdjustmentCoins{
//				Sub: sdk.NewCoins(
//					sdk.NewCoin(denom.ShrP, sdk.NewInt(1)),
//				),
//				Add: sdk.NewCoins(),
//			},
//			oError: nil,
//			d:      "1.1 shrp buy 200 base (1 shrp) -> cost: -1 shrp",
//		},
//		{
//			iCurrent: sdk.NewCoins(
//				sdk.NewCoin(denom.ShrP, sdk.NewInt(2)),
//				sdk.NewCoin(denom.BaseUSD, sdk.NewInt(10)),
//			),
//			iRate: rate,
//			iNeed: sdk.NewInt(300),
//			oCost: AdjustmentCoins{
//				Sub: sdk.NewCoins(
//					sdk.NewCoin(denom.ShrP, sdk.NewInt(2)),
//				),
//				Add: sdk.NewCoins(
//					sdk.NewCoin(denom.BaseUSD, sdk.NewInt(50)),
//				),
//			},
//			oError: nil,
//			d:      "2.10 shrp buy 300 base (1.5 shrp)-> new: 0.6 shrp -> cost: - 2 shrp +50 cent",
//		},
//		{
//			iCurrent: sdk.NewCoins(
//				sdk.NewCoin(denom.ShrP, sdk.NewInt(2)),
//				sdk.NewCoin(denom.BaseUSD, sdk.NewInt(10)),
//			),
//			iRate: rate,
//			iNeed: sdk.NewInt(30000000),
//			oCost: AdjustmentCoins{
//				Sub: nil,
//				Add: nil,
//			},
//			oError: sdkerrors.ErrInsufficientFunds,
//			d:      "insufficient funds",
//		},
//	}
//	for _, tc := range testCases {
//		o, e := GetCostForBaseDenom(tc.iCurrent, tc.iNeed, tc.iRate)
//		require.Equal(t, tc.oCost.Sub, o.Sub, tc.d)
//		require.Equal(t, tc.oCost.Add, o.Add, tc.d)
//		if tc.oError != nil {
//			require.NotNil(t, e, tc.d)
//			require.True(t, sdkerrors.IsOf(e, tc.oError), tc.d)
//		} else {
//			require.Nil(t, e)
//		}
//	}
//
//}
//
