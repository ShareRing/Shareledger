package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	tmos "github.com/tendermint/tendermint/libs/os"
)

const (
	FlagKeySeed  = "key-seed"
	KeySeedUsage = "path to key_seed.json"
)

func GetKeySeedFromFile(seedPath string) (string, error) {
	seeds, err := ioutil.ReadFile(seedPath)
	if err != nil {
		return "", err
	}
	var a map[string]string
	if err := json.Unmarshal(seeds, &a); err != nil {
		return "", err
	}
	return a["secret"], nil
}

func GetAddressFromFile(filepath string) ([]string, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	var addrList []string
	json.Unmarshal(data, &addrList)
	return addrList, nil
}

func CreateContextWithKeyBase(seed string, clientCtx client.Context) (client.Context, error) {
	kb := keyring.NewInMemory()
	keyName := "elon_musk_deer"
	info, err := kb.NewAccount(keyName, seed, "", sdk.GetConfig().Seal().GetFullBIP44Path(), hd.Secp256k1)
	if err != nil {
		return client.Context{}, err
	}

	clientCtx = clientCtx.WithFrom(keyName).WithFromName(info.GetName()).WithFromAddress(info.GetAddress()).WithKeyring(kb)

	return clientCtx, nil
}

func CreateContextFromSeed(seedFile string, clientCtx client.Context) (client.Context, error) {
	seed, err := GetKeySeedFromFile(seedFile)
	if err != nil {
		return client.Context{}, err
	}

	clientCtx, err = CreateContextWithKeyBase(seed, clientCtx)
	if err != nil {
		return client.Context{}, err
	}

	return clientCtx, nil
}

func CreateKeySeed(home, keyName string) (sdk.AccAddress, error) {
	addr, secret, err := server.GenerateCoinKey(hd.Secp256k1)

	if err != nil {
		return nil, err
	}

	info := map[string]string{"secret": secret}

	cliPrint, err := json.Marshal(info)
	if err != nil {
		return nil, err
	}

	if err := WriteFile(fmt.Sprintf("%s_key_seed.json", keyName), home, cliPrint); err != nil {
		return nil, err
	}

	return addr, nil
}

func WriteFile(name string, dir string, contents []byte) error {
	writePath := filepath.Join(dir)
	file := filepath.Join(writePath, name)

	err := tmos.EnsureDir(writePath, 0700)
	if err != nil {
		return err
	}

	err = tmos.WriteFile(file, contents, 0600)
	if err != nil {
		return err
	}

	return nil
}
