package sub

import (
	"bufio"
	"os"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	myutils "github.com/sharering/shareledger/x/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	nodeDirPerm = 0755
	flagKeyName = "key-name"
)

func NewImportKeyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "import-key",
		Short: "import validator key",
		Long:  "import validator key from mnemonic into keyring",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			seedPath := viper.GetString(myutils.FlagKeySeed)
			keySeed, err := myutils.GetKeySeedFromFile(seedPath)
			if err != nil {
				return err
			}

			keyName := viper.GetString(flagKeyName)
			cliHome := viper.GetString(flags.FlagHome)
			if err := os.MkdirAll(cliHome, nodeDirPerm); err != nil {
				return err
			}

			buf := bufio.NewReader(clientCtx.Input)
			kb, _ := keyring.New(sdk.KeyringServiceName(), viper.GetString(flags.FlagKeyringBackend), cliHome, buf)

			_, err = kb.NewAccount(keyName, keySeed, "", sdk.GetConfig().GetFullBIP44Path(), hd.Secp256k1)
			if err != nil {
				return err
			}

			return nil
		},
	}
	cmd.Flags().String(flags.FlagKeyringBackend, "os", "keyring backend: os/pass/file")
	cmd.Flags().String(flagKeyName, "", "key name of validator key to be store in keybase")
	cmd.Flags().String(myutils.FlagKeySeed, "", myutils.KeySeedUsage)

	return cmd
}
