package main

import (
	"fmt"

	gentlemint "bitbucket.org/shareringvietnam/shareledger-fix/x/gentlemint"
	oldId "bitbucket.org/shareringvietnam/shareledger-fix/x/identity"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/genutil"
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

func AddGenesisCustomMigrate(
	ctx *server.Context, cdc *codec.Codec, defaultNodeHome, defaultClientHome string,
) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "custom-migrate <version> <old file> <new file> <merge file>",
		Short: "Copy data from old genesis file the new one.",
		Long:  `Copy data from old genesis file the new one. This command generates merge file.`,
		// Args:  cobra.ExactArgs(1),
		RunE: func(_ *cobra.Command, args []string) error {
			config := ctx.Config
			config.SetRoot(viper.GetString(cli.HomeFlag))

			switch v := args[0]; v {
			case "0.0.1":
				return FromV0_0_1(args[1], args[2], args[3], cdc)
			default:
				fmt.Println("No match version " + v)
			}
			return nil
		},
	}

	return cmd
}
