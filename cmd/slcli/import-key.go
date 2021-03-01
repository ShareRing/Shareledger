package main

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	flagKeySeed = "key-seed"
)

func importKeyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "import-key",
		Short: "import validator key",
		Long:  "import validator key from mnemonic into keyring",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliHome := viper.GetString(flagHome)
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
			seedPath := viper.GetString(flagKeySeed)
			seeds, err := ioutil.ReadFile(seedPath)
			if err != nil {
				return err
			}
			var a map[string]string
			json.Unmarshal(seeds, &a)
			kb.CreateAccount(viper.GetString(flagKeyName), a["secret"], "", defaultKeyPass, sdk.GetConfig().GetFullFundraiserPath(), keys.Secp256k1)
			return nil
		},
	}
	cmd.Flags().String(flagHome, defaultHome, "home for the cli")
	cmd.Flags().String(flags.FlagKeyringBackend, "os", "keyring backend: os/pass/file")
	cmd.Flags().String(flagKeyName, "", "key name of validator key to be store in keybase")
	cmd.Flags().String(flagKeySeed, "", "path to key_seed.json")
	return cmd
}
