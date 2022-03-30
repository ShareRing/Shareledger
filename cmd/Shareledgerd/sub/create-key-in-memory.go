package sub

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/cosmos/cosmos-sdk/client/flags"
	myutils "github.com/sharering/shareledger/x/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func CreateKeyInMemCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "create-key-in-memory",
		Short: "create validator key seed. Only key seed return, key is not stored in the keybase",
		Long: `
Example: shareledger create-key-in-memory --key-name alice --home ./alice_home
Note: --home flag cannot be ommited 
		`,
		RunE: func(cmd *cobra.Command, args []string) error {
			keyName := viper.GetString(flagKeyName)
			cliHome := viper.GetString(flags.FlagHome)

			if err := os.MkdirAll(cliHome, nodeDirPerm); err != nil {
				return err
			}

			addr, err := myutils.CreateKeySeed(cliHome, keyName)
			if err != nil {
				return err
			}

			curr, _ := os.Getwd()
			abs, _ := filepath.Abs(curr)

			if cliHome != "." {
				abs = filepath.Join(abs, cliHome)
			}
			fmt.Println("key save at: ", abs)
			fmt.Println("address: ", addr)
			fmt.Println("bech32 address: ", addr.String())
			return nil
		},
	}

	cmd.Flags().String(flags.FlagHome, ".", "location to save node key")
	cmd.Flags().String(flagKeyName, "", "key name of validator key to be store in keybase")
	return cmd
}
