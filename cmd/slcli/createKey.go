package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"bitbucket.org/shareringvietnam/shareledger-fix/x/myutils"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	cliHome        = "slcli"
	flagKeyName    = "key-name"
	flagPrivFile   = "priv-file"
	defaultKeyPass = "12345678"
	flagAmount     = "amount"
	flagMnemoic    = "mnemonic"
)

func createKeyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-key",
		Short: "create validator key",
		Long:  "validator private key, store in cli home and keybase",
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
				os.RemoveAll(cliHome)
				return err
			}
			addr, secret, err := server.GenerateSaveCoinKey(kb, viper.GetString(flagKeyName), defaultKeyPass, true)
			if err != nil {
				os.RemoveAll(cliHome)
				return err
			}
			info := map[string]string{"secret": secret}
			cliPrint, err := json.Marshal(info)
			if err != nil {
				os.RemoveAll(cliHome)
				return err
			}
			if err := myutils.WriteFile("key_seed.json", cliHome, cliPrint); err != nil {
				os.RemoveAll(cliHome)
				return err
			}
			fmt.Printf("\n%s address: %v\n\n", viper.GetString(flagKeyName), addr)

			return nil
		},
	}
	cmd.Flags().String(flagHome, defaultHome, "home for the cli")
	cmd.Flags().String(flags.FlagKeyringBackend, "pass", "keyring backend: os/pass/file")
	cmd.Flags().String(flagKeyName, "", "key name of validator key to be store in keybase")
	return cmd
}

func createKeyInMemCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-key-in-memory",
		Short: "create validator key seed",
		Long:  "only key seed return, key is not stored in the keybase",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliHome := viper.GetString(flagHome)
			if err := os.MkdirAll(cliHome, nodeDirPerm); err != nil {
				return err
			}
			keyName := viper.GetString(flagKeyName)
			addr, err := myutils.CreateKeySeed(cliHome, keyName)
			if err != nil {
				return err
			}
			fmt.Printf("\n%s address: %v\n\n", keyName, addr)
			return nil
		},
	}
	cmd.Flags().String(flagHome, defaultHome, "home for the cli")
	cmd.Flags().String(flagKeyName, "", "key name of validator key to be store in keybase")
	return cmd
}

func getKeyFromSeedCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-key-from-seed",
		Short: "get private key from seed",
		RunE: func(cmd *cobra.Command, args []string) error {
			keySeed := viper.GetString(myutils.FlagKeySeed)
			privFile := viper.GetString(flagPrivFile)
			seed, err := myutils.GetKeeySeedFromFile(keySeed)
			if err != nil {
				return err
			}
			priv, err := myutils.GetPrivKeyFromSeed(seed)
			if err != nil {
				return err
			}
			data, err := json.Marshal(priv)
			if err != nil {
				return err
			}
			return myutils.WriteFile(privFile, "./", data)
		},
	}
	cmd.Flags().String(myutils.FlagKeySeed, "", "path to key_seed.json file")
	cmd.Flags().String(flagPrivFile, "priv.json", "name of private key file")
	return cmd
}
func getKeyFromMnemonicCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-key-from-mnemonic",
		Short: "get private key from mnemoic (word phrases)",
		// Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			mnemonic := viper.GetString("mnemonic")
			index := viper.GetInt32(flagAmount)

			privs, addrs, err := myutils.GetPrivKeysFromMnemonic(mnemonic, uint32(index))
			if err != nil {
				return err
			}

			for i := 0; i < len(privs); i++ {
				fmt.Printf("%s %s\n", privs[i], addrs[i].String())
			}
			return nil
		},
	}

	cmd.Flags().Uint32(flagAmount, 1, "Amount of private keys")
	cmd.Flags().String(flagMnemoic, "", "Mnemonic(word phrases) for HD derivation")
	return cmd
}
func createKeyBatchCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-key-batch [number of keys]",
		Short: "create a large number of keys",
		Long:  "validator private key, store in cli home and keybase",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliHome := viper.GetString(flagHome)
			if err := os.MkdirAll(cliHome, nodeDirPerm); err != nil {
				return err
			}
			num, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}
			var addresses []sdk.AccAddress
			for i := 0; i < num; i++ {
				addr, secret, err := server.GenerateCoinKey()
				addresses = append(addresses, addr)
				if err != nil {
					return err
				}
				info := map[string]string{"secret": secret}
				cliPrint, err := json.Marshal(info)
				if err != nil {
					os.RemoveAll(cliHome)
					return err
				}
				if err := myutils.WriteFile(fmt.Sprintf("key_seed_%d.json", i), cliHome, cliPrint); err != nil {
					os.RemoveAll(cliHome)
					return err
				}
			}

			data, err := json.Marshal(addresses)
			if err != nil {
				os.RemoveAll(cliHome)
				return err
			}
			if err := myutils.WriteFile("address_list.json", cliHome, data); err != nil {
				os.RemoveAll(cliHome)
				return err
			}

			return nil
		},
	}
	cmd.Flags().String(flagHome, defaultHome, "home for the cli")
	return cmd
}

func getKeysFromDirCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-keys-from-dir [dir]",
		Args:  cobra.ExactArgs(1),
		Short: "get private keys from folder and export them to a file",
		RunE: func(cmd *cobra.Command, args []string) error {
			privFile := viper.GetString(flagPrivFile)
			privs, err := myutils.GetKeysFromDir(args[0])
			if err != nil {
				return err
			}
			data, err := json.Marshal(privs)
			if err != nil {
				return err
			}
			return myutils.WriteFile(privFile, "./", data)
		},
	}
	cmd.Flags().String(myutils.FlagKeySeed, "", "path to key_seed.json file")
	cmd.Flags().String(flagPrivFile, "priv.json", "name of private key file")
	return cmd
}
