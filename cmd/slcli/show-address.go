package main

import (
	"fmt"

	"bitbucket.org/shareringvietnam/shareledger-fix/x/myutils"
	clkeys "github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type bechKeyOutFn func(keyInfo keys.Info) (keys.KeyOutput, error)

const (
	FlagBechPrefix = "bech"
)

func ShowAddressCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-address-from-seed",
		Short: "show address from seed file",
		RunE:  showAddress,
	}
	cmd.Flags().String(myutils.FlagKeySeed, "", "path to key_seed.json")
	cmd.Flags().String(FlagBechPrefix, "acc", "address format account/validator")
	return cmd
}

func showAddress(cmd *cobra.Command, args []string) error {
	keySeed := viper.GetString(myutils.FlagKeySeed)
	seed, err := myutils.GetKeeySeedFromFile(keySeed)
	if err != nil {
		return err
	}
	kb := clkeys.NewInMemoryKeyBase()
	keyName := "elon_musk_deer"
	info, err := kb.CreateAccount(keyName, seed, "", clkeys.DefaultKeyPass, sdk.GetConfig().Seal().GetFullFundraiserPath(), keys.Secp256k1)
	if err != nil {
		return err
	}
	bechOut, err := getBechKeyOut(viper.GetString(FlagBechPrefix))
	if err != nil {
		return err
	}
	printKeyAddress(info, bechOut)
	return nil
}

func getBechKeyOut(bechPrefix string) (bechKeyOutFn, error) {
	switch bechPrefix {
	case sdk.PrefixAccount:
		return keys.Bech32KeyOutput, nil
	case sdk.PrefixValidator:
		return keys.Bech32ValKeyOutput, nil
	case sdk.PrefixConsensus:
		return keys.Bech32ConsKeyOutput, nil
	}

	return nil, fmt.Errorf("invalid Bech32 prefix encoding provided: %s", bechPrefix)
}

func printKeyAddress(info keys.Info, bechKeyOut bechKeyOutFn) {
	ko, err := bechKeyOut(info)
	if err != nil {
		panic(err)
	}

	fmt.Println(ko.Address)
}

func showBech32FromAddress() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "show-bech32-from-address",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			addr, err := sdk.AccAddressFromHex(args[0])
			if err != nil {
				return err
			}
			fmt.Println(addr.String())
			return nil
		},
	}
	return cmd
}

func showAddressFromBech32() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "show-address-from-bech32",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			addr, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}
			fmt.Println(addr)
			return nil
		},
	}
	return cmd
}

func showValOperFromAddress() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "show-valoper-from-address",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			valAddr, err := sdk.ValAddressFromHex(args[0])
			if err != nil {
				return err
			}
			fmt.Println(valAddr.String())
			return nil
		},
	}
	return cmd
}
