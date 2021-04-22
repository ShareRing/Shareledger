package main

import (
	"fmt"

	"github.com/ShareRing/modules/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	auth "github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/genutil"
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

// Add 8 zero to all coin
func FromV1_1_0(inputFilePath, outputFilePath string, cdc *codec.Codec) error {
	appState, genDoc, err := genutil.GenesisStateFromGenFile(cdc, inputFilePath)
	if err != nil {
		return fmt.Errorf("failed to unmarshal old genesis state: %w", err)
	}

	var authState auth.GenesisState
	if appState[auth.ModuleName] != nil {
		cdc.MustUnmarshalJSON(appState[auth.ModuleName], &authState)
	}

	// Update balance
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
