package sub

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	myutils "github.com/sharering/shareledger/x/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type bechKeyOutFn func(keyInfo keyring.Info) (keyring.KeyOutput, error)

func ShowAddressCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-address-from-seed",
		Short: "show address from json seed file",
		RunE: func(cmd *cobra.Command, args []string) error {
			keySeed := viper.GetString(myutils.FlagKeySeed)
			seed, err := myutils.GetKeySeedFromFile(keySeed)
			if err != nil {
				return err
			}

			kb := keyring.NewInMemory()
			keyName := "elon_musk_deer"
			info, err := kb.NewAccount(keyName, seed, "", sdk.GetConfig().Seal().GetFullBIP44Path(), hd.Secp256k1)
			if err != nil {
				return err
			}

			bechOut, err := getBechKeyOut(viper.GetString(keys.FlagBechPrefix))
			if err != nil {
				return err
			}

			printKeyAddress(info, bechOut)

			return nil
		},
	}

	cmd.Flags().String(myutils.FlagKeySeed, "", myutils.KeySeedUsage)
	cmd.Flags().String(keys.FlagBechPrefix, "acc", "address format: acc|val|cons")

	return cmd
}

func getBechKeyOut(bechPrefix string) (bechKeyOutFn, error) {
	switch bechPrefix {
	case sdk.PrefixAccount:
		return keyring.MkAccKeyOutput, nil

	case sdk.PrefixValidator:
		return keyring.MkValKeyOutput, nil

	case sdk.PrefixConsensus:
		return keyring.MkConsKeyOutput, nil
	}

	return nil, fmt.Errorf("invalid Bech32 prefix encoding provided: %s", bechPrefix)
}

func printKeyAddress(info keyring.Info, bechKeyOut bechKeyOutFn) {
	keyOutput, err := bechKeyOut(info)
	if err != nil {
		panic(err)
	}

	fmt.Println(keyOutput.Address)
}
