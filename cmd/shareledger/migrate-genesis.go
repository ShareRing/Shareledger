package main

import (
	"fmt"
	"math/big"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"
	gentlemint "github.com/sharering/shareledger/x/gentlemint"
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

func FromV1_1_0(inputFilePath, outputFilePath string, cdc *codec.Codec) error {
	appState, genDoc, err := genutil.GenesisStateFromGenFile(cdc, inputFilePath)
	if err != nil {
		return fmt.Errorf("failed to unmarshal old genesis state: %w", err)
	}

	const voting_period time.Duration = 432000000000000 // 5 days
	max_deposit_period := gov.DefaultPeriod             // 2 days
	// Add gov module
	govState := gov.NewGenesisState(
		1,
		gov.NewDepositParams(
			sdk.NewCoins(
				sdk.Coin{
					Denom:  "shr",
					Amount: sdk.NewInt(20000000),
				},
			),
			max_deposit_period,
		),
		gov.NewVotingParams(voting_period),
		gov.NewTallyParams(
			sdk.Dec{Int: big.NewInt(334000000000000000)},
			sdk.Dec{Int: big.NewInt(500000000000000000)},
			sdk.Dec{Int: big.NewInt(334000000000000000)}),
	)
	govStatebz, err := cdc.MarshalJSON(govState)
	if err != nil {
		return fmt.Errorf("failed to marshal gov genesis state: %w", err)
	}
	appState[gov.ModuleName] = govStatebz

	// Add upgrade module

	// Staking module
	var stakingState staking.GenesisState
	if appState[staking.ModuleName] != nil {
		cdc.MustUnmarshalJSON(appState[staking.ModuleName], &stakingState)
	}

	for i := 0; i < len(stakingState.Validators); i++ {
		stakingState.Validators[i].UnbondingHeight = 0
	}

	for i := 0; i < len(stakingState.UnbondingDelegations); i++ {
		for j := 0; j < len(stakingState.UnbondingDelegations[i].Entries); j++ {
			stakingState.UnbondingDelegations[i].Entries[j].CreationHeight = 0
		}
	}

	for i := 0; i < len(stakingState.Redelegations); i++ {
		for j := 0; j < len(stakingState.Redelegations[i].Entries); j++ {
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

	for i := 0; i < len(distState.DelegatorStartingInfos); i++ {
		distState.DelegatorStartingInfos[i].StartingInfo.Height = 0
	}

	for i := 0; i < len(distState.ValidatorSlashEvents); i++ {
		distState.ValidatorSlashEvents[i].Height = 0
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

	// Reset old id module
	// var identityState identity.GenesisState
	// identityStateBz, err := cdc.MarshalJSON(identityState)
	// if err != nil {
	// 	return fmt.Errorf("failed to marshal auth genesis state: %w", err)
	// }
	// appState[oldId.ModuleName] = identityStateBz

	appStateJSON, err := cdc.MarshalJSON(appState)
	if err != nil {
		return fmt.Errorf("failed to marshal application genesis state: %w", err)
	}

	genDoc.AppState = appStateJSON

	return genutil.ExportGenesisFile(genDoc, outputFilePath)
}

func AddGenesisCustomMigrate(
	ctx *server.Context, cdc *codec.Codec, defaultNodeHome, defaultClientHome string,
) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "custom-migrate <version> <old file> <new file> <merge file>",
		Short: "Copy data from old genesis file the new one.",
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
