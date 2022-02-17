package types

//func GetCostForBaseDenom(currentShrp sdk.Coins, baseAmount sdk.Int, rate sdk.Dec) (cost sdk.Coin, err error) {
//	neededDecShrp := denom.ToDecShrPCoin(sdk.NewDecCoins(sdk.NewDecCoin(denom.Base, baseAmount)), rate)
//	neededShrp := denom.ShrpDecToCoins(sdk.NewDecCoins(neededDecShrp))
//	if err != nil {
//		return
//	}
//	newBalance, err := denom.SubShrpCoins(currentShrp, neededShrp)
//	if err != nil {
//		return
//	}
//	cost = AdjustmentCoins{
//		Sub: sdk.NewCoins(),
//		Add: sdk.NewCoins(),
//	}
//	zeroI := sdk.NewInt(0)
//	if v := currentShrp.AmountOf(denom.ShrP).Sub(newBalance.AmountOf(denom.ShrP)); v.GT(zeroI) {
//		cost.Sub = cost.Sub.Add(sdk.NewCoin(denom.ShrP, v))
//	}
//	if v := currentShrp.AmountOf(denom.BaseUSD).Sub(newBalance.AmountOf(denom.BaseUSD)); !v.Equal(zeroI) {
//		if v.LT(zeroI) {
//			cost.Add = cost.Add.Add(sdk.NewCoin(denom.BaseUSD, v.Abs()))
//		} else {
//			cost.Sub = cost.Sub.Add(sdk.NewCoin(denom.BaseUSD, v))
//		}
//	}
//	return
//}

//func AddShrpCoins(currentCoins sdk.Coins, addedCoins sdk.Coins) (ac AdjustmentCoins, err error) {
//	if err = currentCoins.Validate(); err != nil {
//		return
//	}
//	if err = addedCoins.Validate(); err != nil {
//		return
//	}
//
//	oldCents := currentCoins.AmountOf(denom.BaseUSD)
//	addedCents := addedCoins.AmountOf(denom.BaseUSD)
//	totalCents := oldCents.Add(addedCents)
//
//	ac.Add = sdk.NewCoins()
//	ac.Sub = sdk.NewCoins()
//	// convert cent to shrp
//	ac.Add = ac.Add.Add(sdk.NewCoin(denom.ShrP, addedCoins.AmountOf(denom.ShrP)))
//	ac.Add = ac.Add.Add(sdk.NewCoin(denom.ShrP, sdk.NewInt(totalCents.Int64()/100)))
//
//	newCent := sdk.NewInt(totalCents.Int64() % 100)
//	if oldCents.GT(newCent) {
//		ac.Sub = ac.Sub.Add(sdk.NewCoin(denom.BaseUSD, oldCents.Sub(newCent)))
//	} else {
//		ac.Add = ac.Add.Add(sdk.NewCoin(denom.BaseUSD, newCent.Sub(oldCents)))
//	}
//
//	return
//}
