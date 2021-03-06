/*
 * Based on https://github.com/cosmos/gaia/blob/v2.0.12/cli_test/cli_test.go
 */
package tests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/tendermint/tendermint/crypto/secp256k1"
	"github.com/tendermint/tendermint/types"
	tmtypes "github.com/tendermint/tendermint/types"

	app "github.com/sharering/shareledger"
	apptypes "github.com/sharering/shareledger/types"

	clientkeys "github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/tests"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"

	idtypes "github.com/ShareRing/modules/id/types"
	shareringUtils "github.com/ShareRing/modules/utils"
)

const (
	denom = "shr"

	keyAuthority = "authority"
	keyTreasurer = "treasurer"
	keyValidator = "validator"
	keyIdSigner  = "id-signer"
	keyAccOp     = "acc-op"
	keyDocIssuer = "doc-issuer"
	keyUser1     = "user1"
	keyUser2     = "user2"
	keyUser3     = "user3"
	keyEmtyUser  = "emptyUser"

	fooDenom  = "footoken"
	feeDenom  = "shr"
	fee2Denom = "fee2token"

	keyBaz       = "baz"
	keyVesting   = "vesting"
	keyFooBarBaz = "foobarbaz"

	DefaultKeyPass = "12345678"
)

var (
	totalCoins = sdk.NewCoins(
		// sdk.NewCoin(fee2Denom, sdk.TokensFromConsensusPower(2000000)),
		// sdk.NewCoin(feeDenom, sdk.TokensFromConsensusPower(2000000)),
		// sdk.NewCoin(fooDenom, sdk.TokensFromConsensusPower(2000)),
		sdk.NewCoin(denom, sdk.TokensFromConsensusPower(300).Add(sdk.NewInt(12))), // add coins from inflation
	)

	startCoins = sdk.NewCoins(
		// sdk.NewCoin(fee2Denom, sdk.TokensFromConsensusPower(1000000)),
		// sdk.NewCoin(feeDenom, sdk.TokensFromConsensusPower(1000000)),
		// sdk.NewCoin(fooDenom, sdk.TokensFromConsensusPower(1000)),
		sdk.NewCoin(denom, shareringUtils.SHRDecimal.Mul(sdk.NewInt(int64(1000)))),
	)

	vestingCoins = sdk.NewCoins(
		sdk.NewCoin(feeDenom, sdk.TokensFromConsensusPower(500000)),
	)
)

//___________________________________________________________________________________
// Fixtures

// Fixtures is used to setup the testing environment
type Fixtures struct {
	BuildDir      string
	RootDir       string
	GaiadBinary   string
	GaiacliBinary string
	ChainID       string
	RPCAddr       string
	Port          string
	GaiadHome     string
	GaiacliHome   string
	P2PAddr       string
	T             *testing.T
}

// NewFixtures creates a new instance of Fixtures with many vars set
func NewFixtures(t *testing.T) *Fixtures {
	tmpDir, err := ioutil.TempDir("", "shareledger_integration_"+t.Name()+"_")
	require.NoError(t, err)

	servAddr, port, err := server.FreeTCPAddr()
	require.NoError(t, err)

	p2pAddr, _, err := server.FreeTCPAddr()
	require.NoError(t, err)

	buildDir := os.Getenv("BUILDDIR")
	if buildDir == "" {
		buildDir, err = filepath.Abs("../build/")
		require.NoError(t, err)
	}
	// apptypes.ConfigureSDK()
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(apptypes.Bech32PrefixAccAddr, apptypes.Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(apptypes.Bech32PrefixValAddr, apptypes.Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(apptypes.Bech32PrefixConsAddr, apptypes.Bech32PrefixConsPub)

	return &Fixtures{
		T:             t,
		BuildDir:      buildDir,
		RootDir:       tmpDir,
		GaiadBinary:   filepath.Join(buildDir, "shareledger"),
		GaiacliBinary: filepath.Join(buildDir, "slcli"),
		GaiadHome:     filepath.Join(tmpDir, ".shareledger"),
		GaiacliHome:   filepath.Join(tmpDir, ".slcli"),
		RPCAddr:       servAddr,
		P2PAddr:       p2pAddr,
		Port:          port,
	}
}

// GenesisFile returns the path of the genesis file
func (f Fixtures) GenesisFile() string {
	return filepath.Join(f.GaiadHome, "config", "genesis.json")
}

// GenesisFile returns the application's genesis state
func (f Fixtures) GenesisState() simapp.GenesisState {
	cdc := codec.New()
	genDoc, err := tmtypes.GenesisDocFromFile(f.GenesisFile())
	require.NoError(f.T, err)

	var appState simapp.GenesisState
	require.NoError(f.T, cdc.UnmarshalJSON(genDoc.AppState, &appState))
	return appState
}

// InitFixtures is called at the beginning of a test  and initializes a chain
// with 1 validator.
func InitFixtures(t *testing.T) (f *Fixtures) {
	f = NewFixtures(t)

	// reset test state
	f.UnsafeResetAll()
	f.CLIConfig("keyring-backend", "test")

	// ensure keystore has foo and bar keys
	f.KeysDelete(keyAuthority)
	f.KeysDelete(keyTreasurer)
	f.KeysDelete(keyValidator)
	f.KeysDelete(keyIdSigner)
	f.KeysDelete(keyAccOp)
	f.KeysDelete(keyDocIssuer)

	f.KeysDelete(keyUser1)
	f.KeysDelete(keyUser2)
	f.KeysDelete(keyUser3)
	f.KeysDelete(keyEmtyUser)

	f.KeysAdd(keyAuthority)
	f.KeysAdd(keyTreasurer)
	f.KeysAdd(keyValidator)
	f.KeysAdd(keyIdSigner)

	f.KeysAdd(keyUser1)
	f.KeysAdd(keyUser2)
	f.KeysAdd(keyUser3)
	f.KeysAdd(keyEmtyUser)

	f.KeysAdd(keyAccOp)
	f.KeysAdd(keyDocIssuer)

	// ensure that CLI output is in JSON format
	f.CLIConfig("output", "json")

	// NOTE: GDInit sets the ChainID
	f.GDInit(keyValidator)
	f.CLIConfig("chain-id", f.ChainID)
	f.CLIConfig("broadcast-mode", "block")
	f.CLIConfig("trust-node", "true")

	f.KeysShow(keyValidator)

	// start an account with tokens
	f.AddGenesisAccount(f.KeyAddress(keyAuthority), startCoins)
	f.AddGenesisAccount(f.KeyAddress(keyTreasurer), startCoins)
	f.AddGenesisAccount(f.KeyAddress(keyValidator), startCoins)
	f.AddGenesisAccount(f.KeyAddress(keyIdSigner), startCoins)
	f.AddGenesisAccount(f.KeyAddress(keyDocIssuer), startCoins)

	f.AddGenesisAccount(f.KeyAddress(keyUser1), startCoins)
	f.AddGenesisAccount(f.KeyAddress(keyUser2), startCoins)
	f.AddGenesisAccount(f.KeyAddress(keyUser3), startCoins)

	f.AddGenesisAccount(f.KeyAddress(keyAccOp), startCoins)

	f.AddGenesisAuthorityAccount(f.KeyAddress(keyAuthority))
	f.AddGenesisTreasurerAccount(f.KeyAddress(keyTreasurer))
	f.AddGenesisValidatorAccount(f.KeyAddress(keyValidator))
	f.AddGenesisOperatorAccount(f.KeyAddress(keyAccOp))

	f.GenTx(keyValidator, "--keyring-backend test --amount=1000000shr")
	f.CollectGenTxs()

	return
}

// Cleanup is meant to be run at the end of a test to clean up an remaining test state
func (f *Fixtures) Cleanup(dirs ...string) {
	clean := append(dirs, f.RootDir)
	for _, d := range clean {
		require.NoError(f.T, os.RemoveAll(d))
	}
}

// Flags returns the flags necessary for making most CLI calls
func (f *Fixtures) Flags() string {
	return fmt.Sprintf("--home=%s --node=%s", f.GaiacliHome, f.RPCAddr)
}

//___________________________________________________________________________________
// id
func (f *Fixtures) CreateId(id string, backup, owner sdk.AccAddress, extraData string, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("%s tx id create %s %s %s %s %v", f.GaiacliBinary, id, backup, owner, extraData, f.Flags())
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), DefaultKeyPass)
}

func (f *Fixtures) CreateIdInBatch(id []string, backup, owner []sdk.AccAddress, extraData []string, flags ...string) (bool, string, string) {
	sep := ","

	backupArr := make([]string, 0, len(backup))
	ownerArr := make([]string, 0, len(backup))

	for i := 0; i < len(backup); i++ {
		backupArr = append(backupArr, backup[i].String())
		ownerArr = append(ownerArr, owner[i].String())
	}

	ids := strings.Join(id, sep)
	owners := strings.Join(ownerArr, sep)
	backups := strings.Join(backupArr, sep)
	extras := strings.Join(extraData, sep)

	cmd := fmt.Sprintf("%s tx id create-batch %s %s %s %s %v", f.GaiacliBinary, ids, backups, owners, extras, f.Flags())
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), DefaultKeyPass)
}

func (f *Fixtures) ReplaceIdOwner(id string, newOwner sdk.AccAddress, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("%s tx id replace %s %s %v", f.GaiacliBinary, id, newOwner.String(), f.Flags())
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), DefaultKeyPass)
}

func (f *Fixtures) UpdateId(id, extraData string, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("%s tx id update %s %s %v", f.GaiacliBinary, id, extraData, f.Flags())
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), DefaultKeyPass)
}

func (f *Fixtures) QueryIdById(id string, flags ...string) idtypes.ID {
	cmd := fmt.Sprintf("%s query id info id %s %v", f.GaiacliBinary, id, f.Flags())

	out, _ := tests.ExecuteT(f.T, addFlags(cmd, flags), "")
	var idRs idtypes.ID
	cdc := app.MakeCodec()
	err := cdc.UnmarshalJSON([]byte(out), &idRs)
	require.NoError(f.T, err, "out %v\n, err %v", out, err)
	return idRs
}

func (f *Fixtures) QueryIdByOwner(owner sdk.Address, flags ...string) idtypes.ID {
	cmd := fmt.Sprintf("%s query id info address %s %v", f.GaiacliBinary, owner.String(), f.Flags())

	out, _ := tests.ExecuteT(f.T, addFlags(cmd, flags), "")
	var idRs idtypes.ID
	cdc := app.MakeCodec()
	err := cdc.UnmarshalJSON([]byte(out), &idRs)
	require.NoError(f.T, err, "out %v\n, err %v", out, err)
	return idRs
}

//___________________________________________________________________________________
// gaiad

// UnsafeResetAll is gaiad unsafe-reset-all
func (f *Fixtures) UnsafeResetAll(flags ...string) {
	cmd := fmt.Sprintf("%s --home=%s unsafe-reset-all", f.GaiadBinary, f.GaiadHome)
	executeWrite(f.T, addFlags(cmd, flags))
	err := os.RemoveAll(filepath.Join(f.GaiadHome, "config", "gentx"))
	require.NoError(f.T, err)
}

// GDInit is gaiad init
// NOTE: GDInit sets the ChainID for the Fixtures instance
func (f *Fixtures) GDInit(moniker string, flags ...string) {
	cmd := fmt.Sprintf("%s init -o --home=%s %s", f.GaiadBinary, f.GaiadHome, moniker)
	_, stderr := tests.ExecuteT(f.T, addFlags(cmd, flags), DefaultKeyPass)

	var chainID string
	var initRes map[string]json.RawMessage

	err := json.Unmarshal([]byte(stderr), &initRes)
	require.NoError(f.T, err)

	err = json.Unmarshal(initRes["chain_id"], &chainID)
	require.NoError(f.T, err)

	f.ChainID = chainID
}

// AddGenesisAccount is gaiad add-genesis-account
func (f *Fixtures) AddGenesisAccount(address sdk.AccAddress, coins sdk.Coins, flags ...string) {
	cmd := fmt.Sprintf("%s add-genesis-account %s %s --home=%s", f.GaiadBinary, address, coins, f.GaiadHome)
	executeWriteCheckErr(f.T, addFlags(cmd, flags))
}

// GenTx is gaiad gentx
func (f *Fixtures) GenTx(name string, flags ...string) {
	cmd := fmt.Sprintf("%s gentx --name=%s --home=%s --home-client=%s", f.GaiadBinary, name, f.GaiadHome, f.GaiacliHome)
	executeWriteCheckErr(f.T, addFlags(cmd, flags), DefaultKeyPass)
}

// CollectGenTxs is gaiad collect-gentxs
func (f *Fixtures) CollectGenTxs(flags ...string) {
	cmd := fmt.Sprintf("%s collect-gentxs --home=%s", f.GaiadBinary, f.GaiadHome)
	executeWriteCheckErr(f.T, addFlags(cmd, flags), DefaultKeyPass)
}

// GDStart runs gaiad start with the appropriate flags and returns a process
func (f *Fixtures) GDStart(flags ...string) *tests.Process {
	cmd := fmt.Sprintf("%s start --home=%s --rpc.laddr=%v --p2p.laddr=%v", f.GaiadBinary, f.GaiadHome, f.RPCAddr, f.P2PAddr)
	proc := tests.GoExecuteTWithStdout(f.T, addFlags(cmd, flags))
	tests.WaitForTMStart(f.Port)
	tests.WaitForNextNBlocksTM(1, f.Port)
	return proc
}

// GDTendermint returns the results of gaiad tendermint [query]
func (f *Fixtures) GDTendermint(query string) string {
	cmd := fmt.Sprintf("%s tendermint %s --home=%s", f.GaiadBinary, query, f.GaiadHome)
	success, stdout, stderr := executeWriteRetStdStreams(f.T, cmd)
	require.Empty(f.T, stderr)
	require.True(f.T, success)
	return strings.TrimSpace(stdout)
}

// ValidateGenesis runs gaiad validate-genesis
func (f *Fixtures) ValidateGenesis() {
	cmd := fmt.Sprintf("%s validate-genesis --home=%s", f.GaiadBinary, f.GaiadHome)
	executeWriteCheckErr(f.T, cmd)
}

// gentlemint: Add default system accounts to genesis
func (f *Fixtures) AddGenesisAuthorityAccount(address sdk.AccAddress, flags ...string) {
	cmd := fmt.Sprintf("%s add-genesis-authority %s --home=%s", f.GaiadBinary, address, f.GaiadHome)
	executeWriteCheckErr(f.T, addFlags(cmd, flags))
}

func (f *Fixtures) AddGenesisTreasurerAccount(address sdk.AccAddress, flags ...string) {
	cmd := fmt.Sprintf("%s add-genesis-treasurer %s --home=%s", f.GaiadBinary, address, f.GaiadHome)
	executeWriteCheckErr(f.T, addFlags(cmd, flags))
}
func (f *Fixtures) AddGenesisValidatorAccount(address sdk.AccAddress, flags ...string) {
	cmd := fmt.Sprintf("%s add-genesis-validator %s --home=%s", f.GaiadBinary, address, f.GaiadHome)
	executeWriteCheckErr(f.T, addFlags(cmd, flags))
}

func (f *Fixtures) AddGenesisOperatorAccount(address sdk.AccAddress, flags ...string) {
	cmd := fmt.Sprintf("%s add-genesis-account-operator %s --home=%s", f.GaiadBinary, address, f.GaiadHome)
	executeWriteCheckErr(f.T, addFlags(cmd, flags))
}

//___________________________________________________________________________________
// gaiacli keys

// KeysDelete is gaiacli keys delete
func (f *Fixtures) KeysDelete(name string, flags ...string) {
	cmd := fmt.Sprintf("%s keys delete --home=%s %s", f.GaiacliBinary, f.GaiacliHome, name)
	executeWrite(f.T, addFlags(cmd, append(append(flags, "-y"), "-f")))
}

// KeysAdd is gaiacli keys add
func (f *Fixtures) KeysAdd(name string, flags ...string) {
	cmd := fmt.Sprintf("%s keys add --home=%s %s", f.GaiacliBinary, f.GaiacliHome, name)
	executeWriteCheckErr(f.T, addFlags(cmd, flags), DefaultKeyPass)
}

// KeysAddRecover prepares gaiacli keys add --recover
func (f *Fixtures) KeysAddRecover(name, mnemonic string, flags ...string) (exitSuccess bool, stdout, stderr string) {
	cmd := fmt.Sprintf("%s keys add --home=%s --recover %s", f.GaiacliBinary, f.GaiacliHome, name)
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), DefaultKeyPass, mnemonic)
}

// KeysAddRecoverHDPath prepares gaiacli keys add --recover --account --index
func (f *Fixtures) KeysAddRecoverHDPath(name, mnemonic string, account uint32, index uint32, flags ...string) {
	cmd := fmt.Sprintf("%s keys add --home=%s --recover %s --account %d --index %d", f.GaiacliBinary, f.GaiacliHome, name, account, index)
	executeWriteCheckErr(f.T, addFlags(cmd, flags), DefaultKeyPass, mnemonic)
}

// KeysShow is gaiacli keys show
func (f *Fixtures) KeysShow(name string, flags ...string) keys.KeyOutput {
	cmd := fmt.Sprintf("%s keys show --home=%s %s", f.GaiacliBinary, f.GaiacliHome, name)
	out, _ := tests.ExecuteT(f.T, addFlags(cmd, flags), "")
	var ko keys.KeyOutput
	err := clientkeys.UnmarshalJSON([]byte(out), &ko)
	require.NoError(f.T, err)
	return ko
}

// KeyAddress returns the SDK account address from the key
func (f *Fixtures) KeyAddress(name string) sdk.AccAddress {
	ko := f.KeysShow(name)
	accAddr, err := sdk.AccAddressFromBech32(ko.Address)
	require.NoError(f.T, err)
	return accAddr
}

//___________________________________________________________________________________
// gaiacli config

// CLIConfig is gaiacli config
func (f *Fixtures) CLIConfig(key, value string, flags ...string) {
	cmd := fmt.Sprintf("%s config --home=%s %s %s", f.GaiacliBinary, f.GaiacliHome, key, value)
	executeWriteCheckErr(f.T, addFlags(cmd, flags))
}

//___________________________________________________________________________________
// gaiacli tx send/sign/broadcast

// TxSend is gaiacli tx send
func (f *Fixtures) TxSend(from string, to sdk.AccAddress, amount sdk.Coin, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("%s tx send %s %s %s %v", f.GaiacliBinary, from, to, amount, f.Flags())
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), DefaultKeyPass)
}

// TxSign is gaiacli tx sign
func (f *Fixtures) TxSign(signer, fileName string, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("%s tx sign %v --from=%s %v", f.GaiacliBinary, f.Flags(), signer, fileName)
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), DefaultKeyPass)
}

// TxBroadcast is gaiacli tx broadcast
func (f *Fixtures) TxBroadcast(fileName string, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("%s tx broadcast %v %v", f.GaiacliBinary, f.Flags(), fileName)
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), DefaultKeyPass)
}

// TxEncode is gaiacli tx encode
func (f *Fixtures) TxEncode(fileName string, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("%s tx encode %v %v", f.GaiacliBinary, f.Flags(), fileName)
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), DefaultKeyPass)
}

// TxMultisign is gaiacli tx multisign
func (f *Fixtures) TxMultisign(fileName, name string, signaturesFiles []string,
	flags ...string) (bool, string, string) {

	cmd := fmt.Sprintf("%s tx multisign %v %s %s %s", f.GaiacliBinary, f.Flags(),
		fileName, name, strings.Join(signaturesFiles, " "),
	)
	return executeWriteRetStdStreams(f.T, cmd)
}

//___________________________________________________________________________________
// gaiacli tx staking

// TxStakingCreateValidator is gaiacli tx staking create-validator
func (f *Fixtures) TxStakingCreateValidator(from, consPubKey string, amount sdk.Coin, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("%s tx staking create-validator %v --from=%s --pubkey=%s", f.GaiacliBinary, f.Flags(), from, consPubKey)
	cmd += fmt.Sprintf(" --amount=%v --moniker=%v --commission-rate=%v", amount, from, "0.05")
	cmd += fmt.Sprintf(" --commission-max-rate=%v --commission-max-change-rate=%v", "0.20", "0.10")
	cmd += fmt.Sprintf(" --min-self-delegation=%v", "1")
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), DefaultKeyPass)
}

// TxStakingUnbond is gaiacli tx staking unbond
func (f *Fixtures) TxStakingUnbond(from, shares string, validator sdk.ValAddress, flags ...string) bool {
	cmd := fmt.Sprintf("%s tx staking unbond %s %v --from=%s %v", f.GaiacliBinary, validator, shares, from, f.Flags())
	return executeWrite(f.T, addFlags(cmd, flags), DefaultKeyPass)
}

//___________________________________________________________________________________
// gaiacli tx gov

// TxGovSubmitProposal is gaiacli tx gov submit-proposal
func (f *Fixtures) TxGovSubmitProposal(from, typ, title, description string, deposit sdk.Coin, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("%s tx gov submit-proposal %v --from=%s --type=%s", f.GaiacliBinary, f.Flags(), from, typ)
	cmd += fmt.Sprintf(" --title=%s --description=%s --deposit=%s", title, description, deposit)
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), DefaultKeyPass)
}

// TxGovDeposit is gaiacli tx gov deposit
func (f *Fixtures) TxGovDeposit(proposalID int, from string, amount sdk.Coin, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("%s tx gov deposit %d %s --from=%s %v", f.GaiacliBinary, proposalID, amount, from, f.Flags())
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), DefaultKeyPass)
}

// TxGovVote is gaiacli tx gov vote
func (f *Fixtures) TxGovVote(proposalID int, option gov.VoteOption, from string, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("%s tx gov vote %d %s --from=%s %v", f.GaiacliBinary, proposalID, option, from, f.Flags())
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), DefaultKeyPass)
}

// TxGovSubmitParamChangeProposal executes a CLI parameter change proposal
// submission.
func (f *Fixtures) TxGovSubmitParamChangeProposal(
	from, proposalPath string, deposit sdk.Coin, flags ...string,
) (bool, string, string) {

	cmd := fmt.Sprintf(
		"%s tx gov submit-proposal param-change %s --from=%s %v",
		f.GaiacliBinary, proposalPath, from, f.Flags(),
	)

	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), DefaultKeyPass)
}

// TxGovSubmitCommunityPoolSpendProposal executes a CLI community pool spend proposal
// submission.
func (f *Fixtures) TxGovSubmitCommunityPoolSpendProposal(
	from, proposalPath string, deposit sdk.Coin, flags ...string,
) (bool, string, string) {

	cmd := fmt.Sprintf(
		"%s tx gov submit-proposal community-pool-spend %s --from=%s %v",
		f.GaiacliBinary, proposalPath, from, f.Flags(),
	)

	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), DefaultKeyPass)
}

//___________________________________________________________________________________
// gaiacli query account

// QueryAccount is gaiacli query account
func (f *Fixtures) QueryAccount(address sdk.AccAddress, flags ...string) auth.BaseAccount {
	cmd := fmt.Sprintf("%s query account %s %v", f.GaiacliBinary, address, f.Flags())
	out, _ := tests.ExecuteT(f.T, addFlags(cmd, flags), "")
	var initRes map[string]json.RawMessage
	err := json.Unmarshal([]byte(out), &initRes)
	require.NoError(f.T, err, "out %v, err %v", out, err)
	value := initRes["value"]
	var acc auth.BaseAccount
	cdc := codec.New()
	codec.RegisterCrypto(cdc)
	err = cdc.UnmarshalJSON(value, &acc)
	require.NoError(f.T, err, "value %v, err %v", string(value), err)
	return acc
}

//___________________________________________________________________________________
// gaiacli query txs

// QueryTxs is gaiacli query txs
func (f *Fixtures) QueryTxs(page, limit int, events ...string) *sdk.SearchTxsResult {
	cmd := fmt.Sprintf("%s query txs --page=%d --limit=%d --events='%s' %v", f.GaiacliBinary, page, limit, queryEvents(events), f.Flags())
	out, _ := tests.ExecuteT(f.T, cmd, "")
	var result sdk.SearchTxsResult
	cdc := app.MakeCodec()
	err := cdc.UnmarshalJSON([]byte(out), &result)
	require.NoError(f.T, err, "out %v\n, err %v", out, err)
	return &result
}

// QueryTxsInvalid query txs with wrong parameters and compare expected error
func (f *Fixtures) QueryTxsInvalid(expectedErr error, page, limit int, tags ...string) {
	cmd := fmt.Sprintf("%s query txs --page=%d --limit=%d --events='%s' %v", f.GaiacliBinary, page, limit, queryEvents(tags), f.Flags())
	_, err := tests.ExecuteT(f.T, cmd, "")
	require.EqualError(f.T, expectedErr, err)
}

//___________________________________________________________________________________
// gaiacli query staking

// QueryStakingValidator is gaiacli query staking validator
func (f *Fixtures) QueryStakingValidator(valAddr sdk.ValAddress, flags ...string) staking.Validator {
	cmd := fmt.Sprintf("%s query staking validator %s %v", f.GaiacliBinary, valAddr, f.Flags())
	out, _ := tests.ExecuteT(f.T, addFlags(cmd, flags), "")
	var validator staking.Validator
	cdc := app.MakeCodec()
	err := cdc.UnmarshalJSON([]byte(out), &validator)
	require.NoError(f.T, err, "out %v\n, err %v", out, err)
	return validator
}

// QueryStakingUnbondingDelegationsFrom is gaiacli query staking unbonding-delegations-from
func (f *Fixtures) QueryStakingUnbondingDelegationsFrom(valAddr sdk.ValAddress, flags ...string) []staking.UnbondingDelegation {
	cmd := fmt.Sprintf("%s query staking unbonding-delegations-from %s %v", f.GaiacliBinary, valAddr, f.Flags())
	out, _ := tests.ExecuteT(f.T, addFlags(cmd, flags), "")
	var ubds []staking.UnbondingDelegation
	cdc := app.MakeCodec()
	err := cdc.UnmarshalJSON([]byte(out), &ubds)
	require.NoError(f.T, err, "out %v\n, err %v", out, err)
	return ubds
}

// QueryStakingDelegationsTo is gaiacli query staking delegations-to
func (f *Fixtures) QueryStakingDelegationsTo(valAddr sdk.ValAddress, flags ...string) []staking.Delegation {
	cmd := fmt.Sprintf("%s query staking delegations-to %s %v", f.GaiacliBinary, valAddr, f.Flags())
	out, _ := tests.ExecuteT(f.T, addFlags(cmd, flags), "")
	var delegations []staking.Delegation
	cdc := app.MakeCodec()
	err := cdc.UnmarshalJSON([]byte(out), &delegations)
	require.NoError(f.T, err, "out %v\n, err %v", out, err)
	return delegations
}

// QueryStakingPool is gaiacli query staking pool
func (f *Fixtures) QueryStakingPool(flags ...string) staking.Pool {
	cmd := fmt.Sprintf("%s query staking pool %v", f.GaiacliBinary, f.Flags())
	out, _ := tests.ExecuteT(f.T, addFlags(cmd, flags), "")
	var pool staking.Pool
	cdc := app.MakeCodec()
	err := cdc.UnmarshalJSON([]byte(out), &pool)
	require.NoError(f.T, err, "out %v\n, err %v", out, err)
	return pool
}

// QueryStakingParameters is gaiacli query staking parameters
func (f *Fixtures) QueryStakingParameters(flags ...string) staking.Params {
	cmd := fmt.Sprintf("%s query staking params %v", f.GaiacliBinary, f.Flags())
	out, _ := tests.ExecuteT(f.T, addFlags(cmd, flags), "")
	var params staking.Params
	cdc := app.MakeCodec()
	err := cdc.UnmarshalJSON([]byte(out), &params)
	require.NoError(f.T, err, "out %v\n, err %v", out, err)
	return params
}

//___________________________________________________________________________________
// gaiacli query gov

// QueryGovParamDeposit is gaiacli query gov param deposit
func (f *Fixtures) QueryGovParamDeposit() gov.DepositParams {
	cmd := fmt.Sprintf("%s query gov param deposit %s", f.GaiacliBinary, f.Flags())
	out, _ := tests.ExecuteT(f.T, cmd, "")
	var depositParam gov.DepositParams
	cdc := app.MakeCodec()
	err := cdc.UnmarshalJSON([]byte(out), &depositParam)
	require.NoError(f.T, err, "out %v\n, err %v", out, err)
	return depositParam
}

// QueryGovParamVoting is gaiacli query gov param voting
func (f *Fixtures) QueryGovParamVoting() gov.VotingParams {
	cmd := fmt.Sprintf("%s query gov param voting %s", f.GaiacliBinary, f.Flags())
	out, _ := tests.ExecuteT(f.T, cmd, "")
	var votingParam gov.VotingParams
	cdc := app.MakeCodec()
	err := cdc.UnmarshalJSON([]byte(out), &votingParam)
	require.NoError(f.T, err, "out %v\n, err %v", out, err)
	return votingParam
}

// QueryGovParamTallying is gaiacli query gov param tallying
func (f *Fixtures) QueryGovParamTallying() gov.TallyParams {
	cmd := fmt.Sprintf("%s query gov param tallying %s", f.GaiacliBinary, f.Flags())
	out, _ := tests.ExecuteT(f.T, cmd, "")
	var tallyingParam gov.TallyParams
	cdc := app.MakeCodec()
	err := cdc.UnmarshalJSON([]byte(out), &tallyingParam)
	require.NoError(f.T, err, "out %v\n, err %v", out, err)
	return tallyingParam
}

// QueryGovProposals is gaiacli query gov proposals
func (f *Fixtures) QueryGovProposals(flags ...string) gov.Proposals {
	cmd := fmt.Sprintf("%s query gov proposals %v", f.GaiacliBinary, f.Flags())
	stdout, stderr := tests.ExecuteT(f.T, addFlags(cmd, flags), "")
	if strings.Contains(stderr, "No matching proposals found") {
		return gov.Proposals{}
	}
	require.Empty(f.T, stderr)
	var out gov.Proposals
	cdc := app.MakeCodec()
	err := cdc.UnmarshalJSON([]byte(stdout), &out)
	require.NoError(f.T, err)
	return out
}

// QueryGovProposal is gaiacli query gov proposal
func (f *Fixtures) QueryGovProposal(proposalID int, flags ...string) gov.Proposal {
	cmd := fmt.Sprintf("%s query gov proposal %d %v", f.GaiacliBinary, proposalID, f.Flags())
	out, _ := tests.ExecuteT(f.T, addFlags(cmd, flags), "")
	var proposal gov.Proposal
	cdc := app.MakeCodec()
	err := cdc.UnmarshalJSON([]byte(out), &proposal)
	require.NoError(f.T, err, "out %v\n, err %v", out, err)
	return proposal
}

// QueryGovVote is gaiacli query gov vote
func (f *Fixtures) QueryGovVote(proposalID int, voter sdk.AccAddress, flags ...string) gov.Vote {
	cmd := fmt.Sprintf("%s query gov vote %d %s %v", f.GaiacliBinary, proposalID, voter, f.Flags())
	out, _ := tests.ExecuteT(f.T, addFlags(cmd, flags), "")
	var vote gov.Vote
	cdc := app.MakeCodec()
	err := cdc.UnmarshalJSON([]byte(out), &vote)
	require.NoError(f.T, err, "out %v\n, err %v", out, err)
	return vote
}

// QueryGovVotes is gaiacli query gov votes
func (f *Fixtures) QueryGovVotes(proposalID int, flags ...string) []gov.Vote {
	cmd := fmt.Sprintf("%s query gov votes %d %v", f.GaiacliBinary, proposalID, f.Flags())
	out, _ := tests.ExecuteT(f.T, addFlags(cmd, flags), "")
	var votes []gov.Vote
	cdc := app.MakeCodec()
	err := cdc.UnmarshalJSON([]byte(out), &votes)
	require.NoError(f.T, err, "out %v\n, err %v", out, err)
	return votes
}

// QueryGovDeposit is gaiacli query gov deposit
func (f *Fixtures) QueryGovDeposit(proposalID int, depositor sdk.AccAddress, flags ...string) gov.Deposit {
	cmd := fmt.Sprintf("%s query gov deposit %d %s %v", f.GaiacliBinary, proposalID, depositor, f.Flags())
	out, _ := tests.ExecuteT(f.T, addFlags(cmd, flags), "")
	var deposit gov.Deposit
	cdc := app.MakeCodec()
	err := cdc.UnmarshalJSON([]byte(out), &deposit)
	require.NoError(f.T, err, "out %v\n, err %v", out, err)
	return deposit
}

// QueryGovDeposits is gaiacli query gov deposits
func (f *Fixtures) QueryGovDeposits(propsalID int, flags ...string) []gov.Deposit {
	cmd := fmt.Sprintf("%s query gov deposits %d %v", f.GaiacliBinary, propsalID, f.Flags())
	out, _ := tests.ExecuteT(f.T, addFlags(cmd, flags), "")
	var deposits []gov.Deposit
	cdc := app.MakeCodec()
	err := cdc.UnmarshalJSON([]byte(out), &deposits)
	require.NoError(f.T, err, "out %v\n, err %v", out, err)
	return deposits
}

//___________________________________________________________________________________
// query slashing

// QuerySigningInfo returns the signing info for a validator
func (f *Fixtures) QuerySigningInfo(val string) slashing.ValidatorSigningInfo {
	cmd := fmt.Sprintf("%s query slashing signing-info %s %s", f.GaiacliBinary, val, f.Flags())
	res, errStr := tests.ExecuteT(f.T, cmd, "")
	require.Empty(f.T, errStr)
	cdc := app.MakeCodec()
	var sinfo slashing.ValidatorSigningInfo
	err := cdc.UnmarshalJSON([]byte(res), &sinfo)
	require.NoError(f.T, err)
	return sinfo
}

// QuerySlashingParams is gaiacli query slashing params
func (f *Fixtures) QuerySlashingParams() slashing.Params {
	cmd := fmt.Sprintf("%s query slashing params %s", f.GaiacliBinary, f.Flags())
	res, errStr := tests.ExecuteT(f.T, cmd, "")
	require.Empty(f.T, errStr)
	cdc := app.MakeCodec()
	var params slashing.Params
	err := cdc.UnmarshalJSON([]byte(res), &params)
	require.NoError(f.T, err)
	return params
}

//___________________________________________________________________________________
// query distribution

// QueryRewards returns the rewards of a delegator
func (f *Fixtures) QueryRewards(delAddr sdk.AccAddress, flags ...string) distribution.QueryDelegatorTotalRewardsResponse {
	cmd := fmt.Sprintf("%s query distribution rewards %s %s", f.GaiacliBinary, delAddr, f.Flags())
	res, errStr := tests.ExecuteT(f.T, cmd, "")
	require.Empty(f.T, errStr)
	cdc := app.MakeCodec()
	var rewards distribution.QueryDelegatorTotalRewardsResponse
	err := cdc.UnmarshalJSON([]byte(res), &rewards)
	require.NoError(f.T, err)
	return rewards
}

//___________________________________________________________________________________
// query supply

// QueryTotalSupply returns the total supply of coins
func (f *Fixtures) QueryTotalSupply(flags ...string) (totalSupply sdk.Coins) {
	cmd := fmt.Sprintf("%s query supply total %s", f.GaiacliBinary, f.Flags())
	res, errStr := tests.ExecuteT(f.T, cmd, "")
	require.Empty(f.T, errStr)
	cdc := app.MakeCodec()
	err := cdc.UnmarshalJSON([]byte(res), &totalSupply)
	require.NoError(f.T, err)
	return totalSupply
}

// QueryTotalSupplyOf returns the total supply of a given coin denom
func (f *Fixtures) QueryTotalSupplyOf(denom string, flags ...string) sdk.Int {
	cmd := fmt.Sprintf("%s query supply total %s %s", f.GaiacliBinary, denom, f.Flags())
	res, errStr := tests.ExecuteT(f.T, cmd, "")
	require.Empty(f.T, errStr)
	cdc := app.MakeCodec()
	var supplyOf sdk.Int
	err := cdc.UnmarshalJSON([]byte(res), &supplyOf)
	require.NoError(f.T, err)
	return supplyOf
}

//___________________________________________________________________________________
//
func (f *Fixtures) ExportGenesis(flags ...string) (*tmtypes.GenesisDoc, error) {
	cmd := fmt.Sprintf("%s export --home=%s", f.GaiadBinary, f.GaiadHome)
	_, res, _ := executeWriteRetStdStreams(f.T, addFlags(cmd, flags), DefaultKeyPass)

	return types.GenesisDocFromJSON([]byte(res))
}

//___________________________________________________________________________________
// executors

func executeWriteCheckErr(t *testing.T, cmdStr string, writes ...string) {
	require.True(t, executeWrite(t, cmdStr, writes...))
}

func executeWrite(t *testing.T, cmdStr string, writes ...string) (exitSuccess bool) {
	exitSuccess, _, _ = executeWriteRetStdStreams(t, cmdStr, writes...)
	return
}

func executeWriteRetStdStreams(t *testing.T, cmdStr string, writes ...string) (bool, string, string) {
	proc := tests.GoExecuteT(t, cmdStr)

	// Enables use of interactive commands
	for _, write := range writes {
		_, err := proc.StdinPipe.Write([]byte(write + "\n"))
		require.NoError(t, err)
	}

	// Read both stdout and stderr from the process
	stdout, stderr, err := proc.ReadAll()
	if err != nil {
		fmt.Println("Err on proc.ReadAll()", err, cmdStr)
	}

	// Log output.
	if len(stdout) > 0 {
		t.Log("Stdout:", string(stdout))
	}
	if len(stderr) > 0 {
		t.Log("Stderr:", string(stderr))
	}

	// Wait for process to exit
	proc.Wait()

	// Return succes, stdout, stderr
	return proc.ExitState.Success(), string(stdout), string(stderr)
}

//___________________________________________________________________________________
// utils

func addFlags(cmd string, flags []string) string {
	for _, f := range flags {
		cmd += " " + f
	}
	return strings.TrimSpace(cmd)
}

func queryEvents(events []string) (out string) {
	for _, event := range events {
		out += event + "&"
	}
	return strings.TrimSuffix(out, "&")
}

// Write the given string to a new temporary file
func WriteToNewTempFile(t *testing.T, s string) *os.File {
	fp, err := ioutil.TempFile(os.TempDir(), "cosmos_cli_test_")
	require.Nil(t, err)
	_, err = fp.WriteString(s)
	require.Nil(t, err)
	return fp
}

func marshalStdTx(t *testing.T, stdTx auth.StdTx) []byte {
	cdc := app.MakeCodec()
	bz, err := cdc.MarshalBinaryBare(stdTx)
	require.NoError(t, err)
	return bz
}

func unmarshalStdTx(t *testing.T, s string) (stdTx auth.StdTx) {
	cdc := app.MakeCodec()
	require.Nil(t, cdc.UnmarshalJSON([]byte(s), &stdTx))
	return
}

func createRandomAddr(amount int) (prvs []secp256k1.PrivKeySecp256k1, addrs []sdk.AccAddress) {
	for i := 0; i < amount; i++ {
		prv := secp256k1.GenPrivKey()
		addr := sdk.AccAddress(prv.PubKey().Address())

		addrs = append(addrs, addr)
		prvs = append(prvs, prv)
	}
	return
}
