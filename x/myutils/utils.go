package myutils

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client/context"
	clkeys "github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/tendermint/go-amino"
	tmos "github.com/tendermint/tendermint/libs/os"
)

const (
	FlagKeySeed = "key-seed"
)

func GetFeeFromShrp(cdc *codec.Codec, cliCtx context.CLIContext, shrpFee float64) (string, error) {
	res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/gentlemint/exchange"), nil)
	if err != nil {
		return "", err
	}
	var exRateStr string
	cdc.MustUnmarshalJSON(res, &exRateStr)
	exRate, err := strconv.ParseFloat(exRateStr, 64)
	if err != nil {
		return "", err
	}
	shrAmt := int(exRate*shrpFee) + 1
	return strconv.Itoa(shrAmt) + "shr", nil
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

func GetKeeySeedFromFile(seedPath string) (string, error) {
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

func GetPrivKeyFromSeed(seed string) (string, error) {
	priv, err := keys.StdDeriveKey(seed, "", sdk.GetConfig().Seal().GetFullFundraiserPath(), keys.Secp256k1)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(priv[:]), nil
}

func GetPrivKeysFromMnemonic(mnemonic string, amount uint32) ([]string, []sdk.AccAddress, error) {
	account := uint32(0)
	privs := make([]string, 0, amount)
	addrs := make([]sdk.AccAddress, 0, amount)
	for i := uint32(0); i < amount; i++ {
		hdPath := keys.CreateHDPath(account, i).String()
		priv, err := keys.StdDeriveKey(mnemonic, "", hdPath, keys.Secp256k1)
		key := hex.EncodeToString(priv[:])
		// fmt.Printf("%s : %v \n", key, priv)

		privs = append(privs, key)

		if err != nil {
			return privs, addrs, err
		}

		kb := clkeys.NewInMemoryKeyBase()
		keyName := "elon_musk_deer"
		info, err := kb.CreateAccount(keyName, mnemonic, "", clkeys.DefaultKeyPass, hdPath, keys.Secp256k1)
		if err != nil {
			return privs, addrs, err
		}

		addrs = append(addrs, info.GetAddress())

	}

	return privs, addrs, nil
}

func GetKeysFromDir(dir string) ([]string, error) {
	var privs []string
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		if strings.Contains(f.Name(), "key_seed") {
			path := filepath.Join(dir, f.Name())
			seed, err := GetKeeySeedFromFile(path)
			if err != nil {
				continue
			}
			priv, err := GetPrivKeyFromSeed(seed)
			if err != nil {
				continue
			}
			privs = append(privs, priv)
		}
	}
	return privs, nil
}

func GetTxBldrAndCtxFromSeed(inBuf io.Reader, cdc *amino.Codec, seed string) (context.CLIContext, auth.TxBuilder, error) {
	kb := clkeys.NewInMemoryKeyBase()
	keyName := "elon_musk_deer"
	info, err := kb.CreateAccount(keyName, seed, "", clkeys.DefaultKeyPass, sdk.GetConfig().Seal().GetFullFundraiserPath(), keys.Secp256k1)
	if err != nil {
		return context.CLIContext{}, auth.TxBuilder{}, err
	}
	cliCtx := context.NewCLIContext().WithCodec(cdc).WithFrom(keyName).WithFromName(info.GetName()).WithFromAddress(info.GetAddress())
	txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc)).WithKeybase(kb)
	return cliCtx, txBldr, nil
}

func CreateKeySeed(home, keyName string) (sdk.AccAddress, error) {
	addr, secret, err := server.GenerateCoinKey()
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
