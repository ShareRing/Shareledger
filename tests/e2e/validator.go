package e2e

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"

	sdkcrypto "github.com/cosmos/cosmos-sdk/crypto"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdktx "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	tmcfg "github.com/tendermint/tendermint/config"
	tmos "github.com/tendermint/tendermint/libs/os"
	"github.com/tendermint/tendermint/p2p"
	"github.com/tendermint/tendermint/privval"

	"github.com/sharering/shareledger/app"
)

type validator struct {
	chain            *chain
	index            int
	moniker          string
	mnemonic         string
	privateKey       cryptotypes.PrivKey
	keyInfo          *keyring.Record
	consensusKey     privval.FilePVKey
	consensusPrivKey cryptotypes.PrivKey
	nodeKey          p2p.NodeKey
}

type account struct {
	moniker    string //nolint:unused
	mnemonic   string
	keyInfo    *keyring.Record
	privateKey cryptotypes.PrivKey
}

func (v *validator) instanceName() string {
	return fmt.Sprintf("%s%d", v.moniker, v.index)
}

func (v *validator) configDir() string {
	return fmt.Sprintf("%s/%s", v.chain.configDir(), v.instanceName())
}

func (v *validator) createConfig() error {
	p := path.Join(v.configDir(), "config")
	return os.MkdirAll(p, 0o755)
}

func (v *validator) init() error {
	if err := v.createConfig(); err != nil {
		return err
	}

	serverCtx := server.NewDefaultContext()
	config := serverCtx.Config

	config.SetRoot(v.configDir())
	config.Moniker = v.moniker

	genDoc, err := getGenDoc(v.configDir())
	if err != nil {
		return err
	}

	appState, err := json.MarshalIndent(app.NewDefaultGenesisState(), "", " ")
	if err != nil {
		return fmt.Errorf("failed to JSON encode app genesis state: %w", err)
	}

	genDoc.ChainID = v.chain.id
	genDoc.Validators = nil
	genDoc.AppState = appState

	if err = genutil.ExportGenesisFile(genDoc, config.GenesisFile()); err != nil {
		return fmt.Errorf("failed to export app genesis state: %w", err)
	}

	tmcfg.WriteConfigFile(filepath.Join(config.RootDir, "config", "config.toml"), config)
	return nil
}

func (v *validator) createNodeKey() error {
	serverCtx := server.NewDefaultContext()
	config := serverCtx.Config

	config.SetRoot(v.configDir())
	config.Moniker = v.moniker

	nodeKey, err := p2p.LoadOrGenNodeKey(config.NodeKeyFile())
	if err != nil {
		return err
	}

	v.nodeKey = *nodeKey
	return nil
}

func (v *validator) createConsensusKey() error {
	serverCtx := server.NewDefaultContext()
	config := serverCtx.Config

	config.SetRoot(v.configDir())
	config.Moniker = v.moniker

	pvKeyFile := config.PrivValidatorKeyFile()
	if err := tmos.EnsureDir(filepath.Dir(pvKeyFile), 0o777); err != nil {
		return err
	}

	pvStateFile := config.PrivValidatorStateFile()
	if err := tmos.EnsureDir(filepath.Dir(pvStateFile), 0o777); err != nil {
		return err
	}

	filePV := privval.LoadOrGenFilePV(pvKeyFile, pvStateFile)
	v.consensusKey = filePV.Key

	return nil
}

func (v *validator) createKeyFromMnemonic(name, mnemonic string) error {
	dir := v.configDir()
	kb, err := keyring.New(keyringAppName, keyring.BackendTest, dir, nil, cdc)
	if err != nil {
		return err
	}

	keyringAlgos, _ := kb.SupportedAlgorithms()
	algo, err := keyring.NewSigningAlgoFromString(string(hd.Secp256k1Type), keyringAlgos)
	if err != nil {
		return err
	}

	info, err := kb.NewAccount(name, mnemonic, "", sdk.FullFundraiserPath, algo)
	if err != nil {
		return err
	}

	privKeyArmor, err := kb.ExportPrivKeyArmor(name, keyringPassphrase)
	if err != nil {
		return err
	}

	privKey, _, err := sdkcrypto.UnarmorDecryptPrivKey(privKeyArmor, keyringPassphrase)
	if err != nil {
		return err
	}

	v.keyInfo = info
	v.mnemonic = mnemonic
	v.privateKey = privKey

	return nil
}

func (c *chain) addAccountFromMnemonic(counts int) error {
	val0ConfigDir := c.validators[0].configDir()
	kb, err := keyring.New(keyringAppName, keyring.BackendTest, val0ConfigDir, nil, cdc)
	if err != nil {
		return err
	}

	keyringAlgos, _ := kb.SupportedAlgorithms()
	algo, err := keyring.NewSigningAlgoFromString(string(hd.Secp256k1Type), keyringAlgos)
	if err != nil {
		return err
	}

	for i := 0; i < counts; i++ {
		name := fmt.Sprintf("acct-%d", i)
		mnemonic := createMnemonic()
		info, err := kb.NewAccount(name, mnemonic, "", sdk.FullFundraiserPath, algo)
		if err != nil {
			return err
		}

		privKeyArmor, err := kb.ExportPrivKeyArmor(name, keyringPassphrase)
		if err != nil {
			return err
		}

		privKey, _, err := sdkcrypto.UnarmorDecryptPrivKey(privKeyArmor, keyringPassphrase)
		if err != nil {
			return err
		}
		acct := account{}
		acct.keyInfo = info
		acct.mnemonic = mnemonic
		acct.privateKey = privKey
		c.genesisAccounts = append(c.genesisAccounts, &acct)
	}

	return nil
}

func (v *validator) createKey(name string) error {
	return v.createKeyFromMnemonic(name, createMnemonic())
}

func (v *validator) buildCreateValidatorMsg(amount sdk.Coin) (sdk.Msg, error) {
	description := stakingtypes.NewDescription(v.moniker, "", "", "", "")
	commissionRates := stakingtypes.CommissionRates{
		Rate:          sdk.MustNewDecFromStr("0.1"),
		MaxRate:       sdk.MustNewDecFromStr("0.2"),
		MaxChangeRate: sdk.MustNewDecFromStr("0.01"),
	}

	// get the initial validator min self delegation
	minSelfDelegation := sdk.OneInt()

	valPubKey, err := cryptocodec.FromTmPubKeyInterface(v.consensusKey.PubKey)
	if err != nil {
		return nil, err
	}

	valAddr, err := v.keyInfo.GetAddress()
	if err != nil {
		return nil, err
	}

	return stakingtypes.NewMsgCreateValidator(
		sdk.ValAddress(valAddr),
		valPubKey,
		amount,
		description,
		commissionRates,
		minSelfDelegation,
	)
}

func (v *validator) signMsg(msgs ...sdk.Msg) (*sdktx.Tx, error) {
	memo := fmt.Sprintf("%s@%s:26656", v.nodeKey.ID(), v.instanceName())
	return signWithPrivKey(v.privateKey, v.chain.id, memo, msgs...)
}
