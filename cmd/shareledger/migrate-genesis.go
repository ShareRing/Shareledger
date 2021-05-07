package main

import (
	"fmt"

	"github.com/ShareRing/modules/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	auth "github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/cosmos/cosmos-sdk/x/supply"
	gentlemint "github.com/sharering/shareledger/x/gentlemint"
	"github.com/sharering/shareledger/x/identity"
	oldId "github.com/sharering/shareledger/x/identity"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tendermint/tendermint/libs/cli"
)

/**
 * FromV0_0_1 copies all idsigners from identity modules to gentlemint module
 */
func FromV0_0_1(oldGenesisFile, newGenesisFile, mergeGenesisFile string, cdc *codec.Codec) error {
	oldAppState, _, err := genutil.GenesisStateFromGenFile(cdc, oldGenesisFile)
	if err != nil {
		return fmt.Errorf("failed to unmarshal old genesis state: %w", err)
	}

	newAppState, newGenDoc, err := genutil.GenesisStateFromGenFile(cdc, newGenesisFile)
	if err != nil {
		return fmt.Errorf("failed to unmarshal new genesis state: %w", err)
	}

	var oldIdState oldId.GenesisState
	if oldAppState[oldId.ModuleName] != nil {
		cdc.MustUnmarshalJSON(oldAppState[oldId.ModuleName], &oldIdState)
	}

	var gentlemintState gentlemint.GenesisState
	if newAppState[gentlemint.ModuleName] != nil {
		cdc.MustUnmarshalJSON(newAppState[gentlemint.ModuleName], &gentlemintState)
	}

	// Copy id signers
	for _, idSignerAddr := range oldIdState.IDSigners {
		addr, err := sdk.AccAddressFromBech32(idSignerAddr[8:])
		if err != nil {
			return err
		}
		idSigner := gentlemint.AccState{Address: addr, Status: gentlemint.ActiveStatus}
		gentlemintState.IdSigners = append(gentlemintState.IdSigners, idSigner)
	}

	gentlemintStateBz, err := cdc.MarshalJSON(gentlemintState)

	if err != nil {
		return fmt.Errorf("failed to marshal auth genesis state: %w", err)
	}

	newAppState[gentlemint.ModuleName] = gentlemintStateBz

	appStateJSON, err := cdc.MarshalJSON(newAppState)
	if err != nil {
		return fmt.Errorf("failed to marshal application genesis state: %w", err)
	}

	newGenDoc.AppState = appStateJSON
	return genutil.ExportGenesisFile(newGenDoc, mergeGenesisFile)

	// Copy id data: Id data is totally different, can not copy
}

// Add 8 zero to all coin
func FromV1_1_0(inputFilePath, outputFilePath string, cdc *codec.Codec) error {
	appState, genDoc, err := genutil.GenesisStateFromGenFile(cdc, inputFilePath)
	if err != nil {
		return fmt.Errorf("failed to unmarshal old genesis state: %w", err)
	}

	// Auth module
	var authState auth.GenesisState
	if appState[auth.ModuleName] != nil {
		cdc.MustUnmarshalJSON(appState[auth.ModuleName], &authState)
	}
	for i := 0; i < len(authState.Accounts); i++ {
		acc := authState.Accounts[i]
		coins := acc.GetCoins()
		coins = coins.Sort()
		newCoins := sdk.Coins{}

		// SHR
		shrAmount := coins.AmountOf("shr")
		newShr := sdk.NewCoin("shr", shrAmount.Mul(utils.SHRDecimal))
		newCoins = newCoins.Add(newShr)

		// SHRP & cent
		shrpAmount := coins.AmountOf("shrp")
		centAmount := coins.AmountOf("cent")

		shrpFromCent := centAmount.Mul(utils.SHRPDecimal).Quo(sdk.NewInt(100))
		shrpAmount = shrpAmount.Mul(utils.SHRPDecimal)
		shrpAmount = shrpAmount.Add(shrpFromCent)

		newShrp := sdk.NewCoin("shrp", shrpAmount)
		newCoins = newCoins.Add(newShrp)

		acc.SetCoins(newCoins)
	}

	authStateBz, err := cdc.MarshalJSON(authState)
	if err != nil {
		return fmt.Errorf("failed to marshal auth genesis state: %w", err)
	}
	appState[auth.ModuleName] = authStateBz

	// Staking module
	var stakingState staking.GenesisState
	if appState[staking.ModuleName] != nil {
		cdc.MustUnmarshalJSON(appState[staking.ModuleName], &stakingState)
	}

	// The default power = token/10^6 and we apply 10^8 decimal
	// so the power will be mutiplied by 10^2 = 100
	powerMultiplier := int64(100)
	stakingState.LastTotalPower = stakingState.LastTotalPower.Mul(sdk.NewInt(powerMultiplier))
	for i := 0; i < len(stakingState.LastValidatorPowers); i++ {
		stakingState.LastValidatorPowers[i].Power = stakingState.LastValidatorPowers[i].Power * powerMultiplier
	}

	for i := 0; i < len(stakingState.Validators); i++ {
		stakingState.Validators[i].DelegatorShares = stakingState.Validators[i].DelegatorShares.MulInt(utils.SHRDecimal)
		stakingState.Validators[i].Tokens = stakingState.Validators[i].Tokens.Mul(utils.SHRDecimal)
		stakingState.Validators[i].UnbondingHeight = 0
	}

	for i := 0; i < len(stakingState.Delegations); i++ {
		stakingState.Delegations[i].Shares = stakingState.Delegations[i].Shares.MulInt(utils.SHRDecimal)
	}

	for i := 0; i < len(stakingState.UnbondingDelegations); i++ {
		for j := 0; j < len(stakingState.UnbondingDelegations[i].Entries); j++ {
			stakingState.UnbondingDelegations[i].Entries[j].InitialBalance = stakingState.UnbondingDelegations[i].Entries[j].InitialBalance.Mul(utils.SHRDecimal)
			stakingState.UnbondingDelegations[i].Entries[j].Balance = stakingState.UnbondingDelegations[i].Entries[j].Balance.Mul(utils.SHRDecimal)
			stakingState.UnbondingDelegations[i].Entries[j].CreationHeight = 0
		}
	}

	for i := 0; i < len(stakingState.Redelegations); i++ {
		for j := 0; j < len(stakingState.Redelegations[i].Entries); j++ {
			stakingState.Redelegations[i].Entries[j].InitialBalance = stakingState.Redelegations[i].Entries[j].InitialBalance.Mul(utils.SHRDecimal)
			stakingState.Redelegations[i].Entries[j].SharesDst = stakingState.Redelegations[i].Entries[j].SharesDst.MulInt(utils.SHRDecimal)
			stakingState.Redelegations[i].Entries[j].CreationHeight = 0
		}
	}

	stakingStateBz, err := cdc.MarshalJSON(stakingState)
	if err != nil {
		return fmt.Errorf("failed to marshal auth genesis state: %w", err)
	}
	appState[staking.ModuleName] = stakingStateBz

	// Distribution module
	var distState distribution.GenesisState
	if appState[distribution.ModuleName] != nil {
		cdc.MustUnmarshalJSON(appState[distribution.ModuleName], &distState)
	}

	feePool := distState.FeePool
	feePool.CommunityPool = decimalize(feePool.CommunityPool)
	distState.FeePool = feePool

	for i := 0; i < len(distState.OutstandingRewards); i++ {
		distState.OutstandingRewards[i].OutstandingRewards = decimalize(distState.OutstandingRewards[i].OutstandingRewards)
	}

	for i := 0; i < len(distState.ValidatorAccumulatedCommissions); i++ {
		distState.ValidatorAccumulatedCommissions[i].Accumulated = decimalize(distState.ValidatorAccumulatedCommissions[i].Accumulated)
	}

	for i := 0; i < len(distState.ValidatorHistoricalRewards); i++ {
		distState.ValidatorHistoricalRewards[i].Rewards.CumulativeRewardRatio = decimalize(distState.ValidatorHistoricalRewards[i].Rewards.CumulativeRewardRatio)
	}

	for i := 0; i < len(distState.ValidatorCurrentRewards); i++ {
		distState.ValidatorCurrentRewards[i].Rewards.Rewards = decimalize(distState.ValidatorCurrentRewards[i].Rewards.Rewards)
	}

	for i := 0; i < len(distState.DelegatorStartingInfos); i++ {
		distState.DelegatorStartingInfos[i].StartingInfo.Stake = distState.DelegatorStartingInfos[i].StartingInfo.Stake.MulInt(utils.SHRPDecimal)
		distState.DelegatorStartingInfos[i].StartingInfo.Height = 0
	}

	distStateBz, err := cdc.MarshalJSON(distState)
	if err != nil {
		return fmt.Errorf("failed to marshal auth genesis state: %w", err)
	}
	appState[distribution.ModuleName] = distStateBz

	// Slashing module
	var slashingState slashing.GenesisState
	if appState[slashing.ModuleName] != nil {
		cdc.MustUnmarshalJSON(appState[slashing.ModuleName], &slashingState)
	}

	for k, v := range slashingState.SigningInfos {
		v.StartHeight = 0
		slashingState.SigningInfos[k] = v
	}

	slashingStateBz, err := cdc.MarshalJSON(slashingState)
	if err != nil {
		return fmt.Errorf("failed to marshal auth genesis state: %w", err)
	}
	appState[slashing.ModuleName] = slashingStateBz

	// Supply module
	var supplyState supply.GenesisState
	if appState[supply.ModuleName] != nil {
		cdc.MustUnmarshalJSON(appState[supply.ModuleName], &supplyState)
	}

	var supplyCoins sdk.Coins
	for i := 0; i < supplyState.Supply.Len(); i++ {
		coin := supplyState.Supply[i]
		var newCoin sdk.Coin
		if coin.Denom == "shr" {
			newCoin = sdk.NewCoin(coin.Denom, coin.Amount.Mul(utils.SHRDecimal))
		} else if coin.Denom == "shrp" {
			newCoin = sdk.NewCoin(coin.Denom, coin.Amount.Mul(utils.SHRPDecimal))
		}
		supplyCoins = supplyCoins.Add(newCoin)
	}
	supplyState.Supply = supplyCoins

	supplyStateBz, err := cdc.MarshalJSON(supplyState)
	if err != nil {
		return fmt.Errorf("failed to marshal auth genesis state: %w", err)
	}
	appState[supply.ModuleName] = supplyStateBz

	// Reset old id module
	var identityState identity.GenesisState
	identityStateBz, err := cdc.MarshalJSON(identityState)
	if err != nil {
		return fmt.Errorf("failed to marshal auth genesis state: %w", err)
	}
	appState[oldId.ModuleName] = identityStateBz

	appStateJSON, err := cdc.MarshalJSON(appState)
	if err != nil {
		return fmt.Errorf("failed to marshal application genesis state: %w", err)
	}

	genDoc.AppState = appStateJSON

	// Update gendoc
	// Validators info
	for i := 0; i < len(genDoc.Validators); i++ {
		genDoc.Validators[i].Power = genDoc.Validators[i].Power * powerMultiplier
	}

	return genutil.ExportGenesisFile(genDoc, outputFilePath)
}

func decimalize(input sdk.DecCoins) sdk.DecCoins {
	var newCoins sdk.DecCoins
	for i := 0; i < len(input); i++ {
		coin := input[i]
		var newCoin sdk.DecCoin
		if coin.Denom == "shr" {
			newCoin = sdk.NewDecCoinFromDec(coin.Denom, coin.Amount.MulInt(utils.SHRDecimal))
		} else if coin.Denom == "shrp" {
			newCoin = sdk.NewDecCoinFromDec(coin.Denom, coin.Amount.MulInt(utils.SHRPDecimal))
		}
		newCoins = newCoins.Add(newCoin)
	}
	return newCoins
}

func AddGenesisCustomMigrate(
	ctx *server.Context, cdc *codec.Codec, defaultNodeHome, defaultClientHome string,
) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "custom-migrate <version> <old file> <new file> <output file>",
		Short: "Migrate genesis file.",
		Long: `Copy data from old genesis file the new one. This command generates merge file.
version 0.0.1: custom-migrate 0.0.1 <old file> <new file> <output file>
version 1.1.0: custom-migrate 1.1.0 <old file> <output file>
		`,
		// Args:  cobra.ExactArgs(1),
		RunE: func(_ *cobra.Command, args []string) error {
			config := ctx.Config
			config.SetRoot(viper.GetString(cli.HomeFlag))

			switch v := args[0]; v {
			case "0.0.1":
				return FromV0_0_1(args[1], args[2], args[3], cdc)
			case "1.1.0":
				return FromV1_1_0(args[1], args[2], cdc)
			default:
				fmt.Println("No match version " + v)
			}
			return nil
		},
	}

	return cmd
}
