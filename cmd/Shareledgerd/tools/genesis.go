package tools

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/gogo/protobuf/proto"
	"github.com/spf13/cobra"
)

func unmarshalGenesisState(cmd *cobra.Command, homeDir string, moduleName string, ptr proto.Message) error {
	serverCtx := server.NewDefaultContext()
	config := serverCtx.Config
	config = config.SetRoot(homeDir)

	clientCtx, err := client.GetClientQueryContext(cmd)
	if err != nil {
		return err
	}
	depCdc := clientCtx.Codec
	cdc := depCdc.(codec.Codec)

	genFile := config.GenesisFile()
	appState, _, err := genutiltypes.GenesisStateFromGenFile(genFile)
	if err != nil {
		return err
	}
	return cdc.UnmarshalJSON(appState[moduleName], ptr)
}

func exportGenesisFile(cmd *cobra.Command, homeDir string, moduleName string, ptr proto.Message) error {
	serverCtx := server.NewDefaultContext()
	config := serverCtx.Config
	config = config.SetRoot(homeDir)

	clientCtx, err := client.GetClientQueryContext(cmd)
	if err != nil {
		return err
	}
	depCdc := clientCtx.Codec
	cdc := depCdc.(codec.Codec)

	genFile := config.GenesisFile()
	appState, genDoc, err := genutiltypes.GenesisStateFromGenFile(genFile)
	if err != nil {
		return err
	}
	jsonData, err := cdc.MarshalJSON(ptr)
	if err != nil {
		return err
	}
	appState[moduleName] = jsonData
	appStateJSON, err := json.Marshal(appState)
	if err != nil {
		return err
	}
	genDoc.AppState = appStateJSON

	return genutil.ExportGenesisFile(genDoc, genFile)
}

func getAddr(cmd *cobra.Command, homeDir string, args []string) (sdk.AccAddress, error) {
	addr, err := sdk.AccAddressFromBech32(args[0])

	// try to convert from key if param is not address
	if err != nil {
		inBuf := bufio.NewReader(cmd.InOrStdin())
		keyringBackend, _ := cmd.Flags().GetString(flags.FlagKeyringBackend)

		// attempt to lookup address from Keybase if no address was provided
		kb, err := keyring.New(sdk.KeyringServiceName(), keyringBackend, homeDir, inBuf)
		if err != nil {
			return nil, err
		}

		info, err := kb.Key(args[0])
		if err != nil {
			return nil, fmt.Errorf("failed to get address from Keybase: %w", err)
		}

		addr = info.GetAddress()
	}
	return addr, nil
}
