package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"

	clkeys "github.com/cosmos/cosmos-sdk/client/keys"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authexported "github.com/cosmos/cosmos-sdk/x/auth/exported"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/cosmos/cosmos-sdk/x/supply"
	"github.com/sharering/shareledger/x/electoral"
	"github.com/sharering/shareledger/x/gentlemint"
	"github.com/sharering/shareledger/x/myutils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tendermint/go-amino"
	tmconfig "github.com/tendermint/tendermint/config"
	"github.com/tendermint/tendermint/crypto"
	tmos "github.com/tendermint/tendermint/libs/os"
	"github.com/tendermint/tendermint/types"
	tmtime "github.com/tendermint/tendermint/types/time"
)

const (
	flagOutputDir     = "output-dir"
	flagAuthDir       = "auth-dir"
	flagStartingIP    = "starting-ip"
	flagNodeNum       = "node-num"
	nodeDaemonHome    = "shareledger"
	nodeClientHome    = "slcli"
	nodeDirPerm       = 0755
	defaultKeyPass    = "defaultKeyPass"
	stakingTokenDenom = "shr"
)

var (
	stakeSupply = sdk.NewInt(0)
)

func testnetCmd(ctx *server.Context, cdc *codec.Codec, mbm module.BasicManager, genAccIterator genutiltypes.GenesisAccountsIterator) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "testnet [chain-id]",
		Args:  cobra.ExactArgs(1),
		Short: "",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			config := ctx.Config
			chainID := args[0]
			outputDir := viper.GetString(flagOutputDir)
			authDir := viper.GetString(flagAuthDir)
			nodeNum := viper.GetInt(flagNodeNum)
			startingIP := viper.GetString(flagStartingIP)
			return InitTestnet(cmd, cdc, mbm, config, chainID, outputDir, authDir, startingIP, nodeNum, genAccIterator)
		},
	}
	cmd.Flags().StringP(flagOutputDir, "o", "./testnet",
		"Directory to store initialization data for the testnet")
	cmd.Flags().StringP(flagAuthDir, "a", "./authority", "Directory containing auth keys")
	cmd.Flags().String(flags.FlagKeyringBackend, "os", "keyring backend: os/pass/file")
	cmd.Flags().IntP(flagNodeNum, "n", 4, "Number of nodes to initialize testnet with")
	cmd.Flags().String(flagStartingIP, "172.194.4.2",
		"Starting IP address (172.194.4.2 results in persistent peers list ID0@172.194.4.2:46656, ID1@192.168.0.2:46656, ...)")
	return cmd
}

func InitTestnet(cmd *cobra.Command, cdc *codec.Codec, mbm module.BasicManager, config *tmconfig.Config, chainID, outputDir, authDir, startingIP string, nodeNum int, genAccIterator genutiltypes.GenesisAccountsIterator) error {
	if chainID == "" {
		chainID = "chain-unspecified"
	}
	monikers := make([]string, nodeNum)
	nodeIDs := make([]string, nodeNum)
	valPubKeys := make([]crypto.PubKey, nodeNum)

	var (
		genFilePaths   []string
		genAccounts    []authexported.GenesisAccount
		voterAddresses []sdk.AccAddress
	)
	inBuf := bufio.NewReader(cmd.InOrStdin())

	// Create special accounts and keys
	specialKeysDir := "special-accounts"

	// Authority account
	authorityAcc, _, _, err := createKeyAndGenesisAccount(nil, specialKeysDir, "authority", inBuf, 0)
	if err != nil {
		return err
	}
	genAccounts = append(genAccounts, authorityAcc)

	// Treasurer account
	treasuryAcc, _, _, err := createKeyAndGenesisAccount(nil, specialKeysDir, "treasury", inBuf, 0)
	if err != nil {
		return err
	}
	genAccounts = append(genAccounts, treasuryAcc)
	cmd.PrintErrf("The keys are in: %s\n", specialKeysDir)
	// Account operator account
	accountOperatorAcc, _, _, err := createKeyAndGenesisAccount(nil, specialKeysDir, "account_operator", inBuf, 0)
	if err != nil {
		return err
	}
	genAccounts = append(genAccounts, accountOperatorAcc)

	// Validator accounts
	for i := 0; i < nodeNum; i++ {
		// setup nodeName, path for slcli, shareledger
		nodeName := fmt.Sprintf("node%d", i)
		daemonHome := filepath.Join(outputDir, nodeName, nodeDaemonHome)
		clientHome := filepath.Join(outputDir, nodeName, nodeClientHome)

		// setup config
		config.SetRoot(daemonHome)
		config.BaseConfig.DBBackend = "cleveldb"
		config.RPC.ListenAddress = "tcp://0.0.0.0:26657"
		config.RPC.CORSAllowedOrigins = []string{"*"}
		if err := os.MkdirAll(filepath.Join(daemonHome, "config"), nodeDirPerm); err != nil {
			os.RemoveAll(outputDir)
			return err
		}
		monikers = append(monikers, nodeName)
		config.Moniker = nodeName

		// path to write genesis.json for each node
		genFilePaths = append(genFilePaths, config.GenesisFile())

		// setup node files
		var err error
		nodeIDs[i], valPubKeys[i], err = genutil.InitializeNodeValidatorFiles(config)
		if err != nil {
			os.RemoveAll(outputDir)
		}

		// generate node ip and peer address
		ip, err := getIP(i, startingIP)
		if err != nil {
			os.RemoveAll(outputDir)
			return err
		}
		peerAddr := fmt.Sprintf("%s@%s:26656", nodeIDs[i], ip)

		// create genesis accounts
		genAcc, addr, seedPath, err := createKeyAndGenesisAccount(nil, clientHome, nodeName, inBuf, 1)
		if err != nil {
			os.RemoveAll(outputDir)
		}
		genAccounts = append(genAccounts, genAcc)
		voterAddresses = append(voterAddresses, genAcc.GetAddress())

		// create genTxs and write it down for each node
		gentxsDir := filepath.Join(outputDir, "gentxs")
		if err := writeStakingGenTx(
			gentxsDir,
			1,
			addr,
			valPubKeys[i],
			nodeName,
			seedPath,
			peerAddr,
			inBuf,
			chainID,
			cdc,
		); err != nil {
			os.RemoveAll(outputDir)
			return err
		}
	}
	if err := initGenFiles(cdc, mbm, chainID, genAccounts, voterAddresses, genFilePaths, nodeNum, authorityAcc.GetAddress(), treasuryAcc.GetAddress(), accountOperatorAcc.GetAddress()); err != nil {
		return err
	}

	if err := collectGenFiles(cdc, config, chainID, monikers, nodeIDs, valPubKeys, nodeNum, outputDir, genAccIterator); err != nil {
		return err
	}
	cmd.PrintErrf("Successfully initialized %d node directories\n", nodeNum)
	data, err := json.Marshal(nodeIDs)
	if err != nil {
		return err
	}
	nodeIDPath := filepath.Join(outputDir, "node_id.json")
	err = ioutil.WriteFile(nodeIDPath, data, 0600)
	if err != nil {
		return err
	}
	return nil
}

func getIP(i int, startingIP string) (ip string, err error) {
	if len(startingIP) == 0 {
		ip, err = server.ExternalIP()
		if err != nil {
			return "", err
		}
		return ip, nil
	}
	return calculateIP(startingIP, i)
}

func calculateIP(ip string, i int) (string, error) {
	ipv4 := net.ParseIP(ip).To4()
	if ipv4 == nil {
		return "", fmt.Errorf("%v: can not parse ipv4 address", ip)
	}
	for j := 0; j < i; j++ {
		ipv4[3]++
	}

	return ipv4.String(), nil
}

func writeFile(name string, dir string, contents []byte) error {
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

func initGenFiles(
	cdc *codec.Codec,
	mbm module.BasicManager,
	chainID string,
	genAccounts []authexported.GenesisAccount,
	voterAddresses []sdk.AccAddress,
	genFiles []string,
	nodeNum int,
	authorityAcc, treasuryAcc, accountOperatorAcc sdk.AccAddress,
) error {
	appGenState := mbm.DefaultGenesis()

	appGenState["supply"] = createSupplyGenesisState(cdc)
	appGenState["electoral"] = createElectoralGenesis(cdc, voterAddresses)
	var authGenState auth.GenesisState
	cdc.MustUnmarshalJSON(appGenState[auth.ModuleName], &authGenState)

	authGenState.Accounts = genAccounts
	appGenState[auth.ModuleName] = cdc.MustMarshalJSON(authGenState)

	// Add gentlemint
	gentlemintGenState := gentlemint.GetGenesisStateFromAppState(cdc, appGenState)
	gentlemintGenState.Authority = authorityAcc.String()
	gentlemintGenState.Treasurer = treasuryAcc.String()
	gentlemintGenState.AccountOperators = append(gentlemintGenState.AccountOperators, gentlemint.AccState{Address: accountOperatorAcc, Status: gentlemint.ActiveStatus})

	gentlemintGenStateBz, err := cdc.MarshalJSON(gentlemintGenState)
	appGenState[gentlemint.ModuleName] = gentlemintGenStateBz

	appGenStateJSON, err := codec.MarshalJSONIndent(cdc, appGenState)
	if err != nil {
		return err
	}

	genDoc := types.GenesisDoc{
		ChainID:    chainID,
		AppState:   appGenStateJSON,
		Validators: nil,
	}

	for i := 0; i < nodeNum; i++ {
		if err := genDoc.SaveAs(genFiles[i]); err != nil {
			return err
		}
	}
	return nil
}

func collectGenFiles(
	cdc *codec.Codec, config *tmconfig.Config, chainID string,
	monikers, nodeIDs []string, valPubkeys []crypto.PubKey,
	nodeNum int, outputDir string, genAccIterator genutiltypes.GenesisAccountsIterator,
) error {
	var appState json.RawMessage
	genTime := tmtime.Now()
	for i := 0; i < nodeNum; i++ {
		nodeName := fmt.Sprintf("node%d", i)
		nodeDir := filepath.Join(outputDir, nodeName, nodeDaemonHome)
		gentxsDir := filepath.Join(outputDir, "gentxs")
		moniker := monikers[i]
		config.Moniker = nodeName
		config.SetRoot(nodeDir)

		nodeID, valPubkey := nodeIDs[i], valPubkeys[i]
		initCfg := genutil.NewInitConfig(chainID, gentxsDir, moniker, nodeID, valPubkey)
		genDoc, err := types.GenesisDocFromFile((config.GenesisFile()))
		if err != nil {
			return err
		}
		nodeAppState, err := genutil.GenAppStateFromConfig(cdc, config, initCfg, *genDoc, genAccIterator)
		if err != nil {
			return err
		}
		if appState == nil {
			appState = nodeAppState
		}

		genFile := config.GenesisFile()
		if err := genutil.ExportGenesisFileWithTime(genFile, chainID, nil, appState, genTime); err != nil {
			return err
		}
	}
	return nil
}
func createKeyAndGenesisAccount(addr sdk.AccAddress, keyHome, keyName string, inBuf *bufio.Reader, power int64) (authexported.GenesisAccount, sdk.AccAddress, string, error) {
	var err error
	if addr == nil {
		addr, err = myutils.CreateKeySeed(keyHome, keyName)
		if err != nil {
			return nil, nil, "", err
		}
	}
	seedPath := filepath.Join(keyHome, fmt.Sprintf("%s_key_seed.json", keyName))
	accStakingTokens := sdk.TokensFromConsensusPower(power).AddRaw(int64(100))

	stakeSupply = stakeSupply.Add(accStakingTokens)
	coins := sdk.Coins{
		sdk.NewCoin(stakingTokenDenom, accStakingTokens),
	}
	return auth.NewBaseAccount(addr, coins, nil, 0, 0), addr, seedPath, nil
}

func writeStakingGenTx(gentxsDir string, power int64, addr sdk.AccAddress, valPubKey crypto.PubKey, nodeName string, seedPath string, peerAddr string, inBuf *bufio.Reader, chainID string, cdc *amino.Codec) error {
	valTokens := sdk.TokensFromConsensusPower(power)
	msg := staking.NewMsgCreateValidator(
		sdk.ValAddress(addr),
		valPubKey,
		sdk.NewCoin(stakingTokenDenom, valTokens),
		staking.NewDescription(nodeName, "", "", "", ""),
		staking.NewCommissionRates(sdk.OneDec(), sdk.OneDec(), sdk.OneDec()),
		sdk.OneInt(),
	)
	tx := auth.NewStdTx([]sdk.Msg{msg}, auth.StdFee{}, []auth.StdSignature{}, peerAddr)
	seed, err := myutils.GetKeeySeedFromFile(seedPath)
	if err != nil {
		return err
	}
	_, txBldr, err := myutils.GetTxBldrAndCtxFromSeed(inBuf, cdc, seed)
	if err != nil {
		return err
	}
	txBldr = txBldr.WithChainID(chainID).WithMemo(peerAddr)
	signedTx, err := txBldr.SignStdTx("elon_musk_deer", clkeys.DefaultKeyPass, tx, false)
	if err != nil {
		return err
	}
	txBytes, err := cdc.MarshalJSON(signedTx)
	if err != nil {
		return err
	}
	if err := writeFile(fmt.Sprintf("%v.json", nodeName), gentxsDir, txBytes); err != nil {
		return err
	}
	return nil
}

func createSupplyGenesisState(cdc *amino.Codec) json.RawMessage {
	totalSupply := sdk.NewCoins(
		sdk.NewCoin(stakingTokenDenom, stakeSupply),
	)
	gen := supply.NewGenesisState(totalSupply)
	return cdc.MustMarshalJSON(gen)
}

func createElectoralGenesis(cdc *amino.Codec, addrs []sdk.AccAddress) json.RawMessage {
	gen := electoral.NewGenesisState(addrs)
	return cdc.MustMarshalJSON(gen)
}
