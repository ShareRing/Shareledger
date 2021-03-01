package main

import (
	"bufio"
	"os"
	"path/filepath"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func deleteKeyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-key",
		Short: "delete validator key",
		Long:  "delete validator key from keyring",
		RunE: func(cmd *cobra.Command, args []string) error {
			homeDir := viper.GetString(flagHome)
			cliHome := filepath.Join(homeDir, cliHome)
			if err := os.MkdirAll(cliHome, nodeDirPerm); err != nil {
				return err
			}
			kb, err := keys.NewKeyring(
				sdk.KeyringServiceName(),
				viper.GetString(flags.FlagKeyringBackend),
				cliHome,
				bufio.NewReader(cmd.InOrStdin()),
			)
			if err != nil {
				return err
			}
			keyName := viper.GetString(flagKeyName)
			kb.Delete(keyName, defaultKeyPass, true)
			return nil
		},
	}
	cmd.Flags().String(flagHome, defaultHome, "home for the cli")
	cmd.Flags().String(flags.FlagKeyringBackend, "os", "keyring backend: os/pass/file")
	cmd.Flags().String(flagKeyName, "", "key name of validator key to be store in keybase")
	cmd.Flags().String(flagKeySeed, "", "path to key_seed.json")
	return cmd
}
