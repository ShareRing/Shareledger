package subcommands

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"

	cfg "github.com/tendermint/tendermint/config"
	"github.com/tendermint/tendermint/p2p"
	"github.com/tendermint/tendermint/privval"
	"github.com/tendermint/tendermint/types"
	cmn "github.com/tendermint/tmlibs/common"

	"github.com/sharering/shareledger/app"
	posTypes "github.com/sharering/shareledger/x/pos/type"
)

var (
	nValidators    int
	nNonValidators int
	outputDir      string
	nodeDirPrefix  string

	populatePersistentPeers bool
	hostnamePrefix          string
	startingIPAddress       string
	p2pPort                 int
	rpcPort                 int
	listIPAddress           string
	listMonikers            string
)

const (
	nodeDirPerm = 0755
)

func init() {
	TestnetFilesCmd.Flags().IntVar(&nValidators, "v", 4,
		"Number of validators to initialize the testnet with")
	TestnetFilesCmd.Flags().IntVar(&nNonValidators, "n", 0,
		"Number of non-validators to initialize the testnet with")
	TestnetFilesCmd.Flags().StringVar(&outputDir, "o", "./mytestnet",
		"Directory to store initialization data for the testnet")
	TestnetFilesCmd.Flags().StringVar(&nodeDirPrefix, "node-dir-prefix", "node",
		"Prefix the directory name for each node with (node results in node0, node1, ...)")

	TestnetFilesCmd.Flags().BoolVar(&populatePersistentPeers, "populate-persistent-peers", true,
		"Update config of each node with the list of persistent peers build using either hostname-prefix or starting-ip-address")
	TestnetFilesCmd.Flags().StringVar(&hostnamePrefix, "hostname-prefix", "node",
		"Hostname prefix (node results in persistent peers list ID0@node0:46656, ID1@node1:46656, ...)")
	TestnetFilesCmd.Flags().StringVar(&startingIPAddress, "starting-ip-address", "",
		"Starting IP address (192.168.0.1 results in persistent peers list ID0@192.168.0.1:46656, ID1@192.168.0.2:46656, ...)")
	TestnetFilesCmd.Flags().StringVar(&listIPAddress, "list-ip-address", "",
		"List of IP addresses ( Ex: 192.168.0.1:46656,192.168.1.3:46656). Number of addreses equals to number of nodes other error throws. This overrides starting-ip-address and p2p-port")
	TestnetFilesCmd.Flags().IntVar(&p2pPort, "p2p-port", 46656, "P2P Port")
	// TestnetFilesCmd.Flags().IntVar(&rpcPort, "rpc-port", 46657, "RPC Port")
	TestnetFilesCmd.Flags().StringVar(&listMonikers, "monikers", "",
		"List of monikers separated by comma( Ex: moniker1,moniker1,moniker3)")
}

// TestnetFilesCmd allows initialisation of files for a Tendermint testnet.
var TestnetFilesCmd = &cobra.Command{
	Use:   "testnet",
	Short: "Initialize files for a Tendermint testnet",
	Long: `testnet will create "v" + "n" number of directories and populate each with
necessary files (private validator, genesis, config, etc.).

Note, strict routability for addresses is turned off in the config file.

Optionally, it will fill in persistent_peers list in config file using either hostnames or IPs.

Example:

	shareledger testnet --v 4 --o ./output --populate-persistent-peers --starting-ip-address 192.168.10.2
	`,
	RunE: testnetFiles,
}

func testnetFiles(cmd *cobra.Command, args []string) error {
	config := cfg.DefaultConfig()

	config.P2P.ListenAddress = P2PListenAddress
	config.RPC.ListenAddress = RPCListenAddress
	config.BaseConfig.ProxyApp = BaseConfigProxyApp

	genVals := make([]types.GenesisValidator, nValidators)
	stateVals := make([]posTypes.Validator, nValidators)
	var appGenesisState app.GenesisState

	for i := 0; i < nValidators; i++ {
		// nodeDirName := cmn.Fmt("%s%d", nodeDirPrefix, i)
		nodeDirName := getMoniker(i)
		nodeDir := filepath.Join(outputDir, nodeDirName)
		config.SetRoot(nodeDir)

		// create config dir
		err := os.MkdirAll(filepath.Join(nodeDir, "config"), nodeDirPerm)
		if err != nil {
			_ = os.RemoveAll(outputDir)
			return err
		}

		// create data dir
		err = os.MkdirAll(filepath.Join(nodeDir, "data"), nodeDirPerm)
		if err != nil {
			_ = os.RemoveAll(outputDir)
			return err
		}

		initFilesWithConfig(config)

		pvKeyFile := filepath.Join(nodeDir, config.BaseConfig.PrivValidatorKey)
		pvStateFile := filepath.Join(nodeDir, config.BaseConfig.PrivValidatorState)

		pv := privval.LoadFilePV(pvKeyFile, pvStateFile)
		genesisState, pubKey := genGenesisState(pv)

		// Update Moniker for Validator tendermint
		genVals[i] = types.GenesisValidator{
			PubKey: pubKey,
			Power:  genesisState.StakeData.Validators[0].ABCIValidator().Power,
			Name:   getMoniker(i),
		}

		stateVals[i] = genesisState.StakeData.Validators[0]
		// Upate Moniker for genesisState Validator
		stateVals[i].Description.Moniker = getMoniker(i)

		appGenesisState = genesisState
	}

	for i := 0; i < nNonValidators; i++ {
		nodeDir := filepath.Join(outputDir, cmn.Fmt("%s%d", nodeDirPrefix, i+nValidators))
		config.SetRoot(nodeDir)

		err := os.MkdirAll(filepath.Join(nodeDir, "config"), nodeDirPerm)
		if err != nil {
			_ = os.RemoveAll(outputDir)
			return err
		}

		initFilesWithConfig(config)
	}

	// Generate genesis doc from generated validators
	genDoc := &types.GenesisDoc{
		GenesisTime: time.Now(),
		ChainID:     "chain-" + cmn.RandStr(6),
		Validators:  genVals,
	}

	appGenesisState.StakeData.Validators = stateVals

	genDoc.AppState = json.RawMessage(appGenesisState.ToJSON())

	// Write genesis file.
	for i := 0; i < nValidators+nNonValidators; i++ {
		var nodeDir string
		if i < nValidators {
			nodeDir = filepath.Join(outputDir, getMoniker(i))
		} else {
			nodeDir = filepath.Join(outputDir, cmn.Fmt("%s%d", nodeDirPrefix, i))
		}

		if err := genDoc.SaveAs(filepath.Join(nodeDir, config.BaseConfig.Genesis)); err != nil {
			_ = os.RemoveAll(outputDir)
			return err
		}
	}

	if populatePersistentPeers {
		err := populatePersistentPeersInConfigAndWriteIt(config)
		if err != nil {
			_ = os.RemoveAll(outputDir)
			return err
		}
	}

	fmt.Printf("Successfully initialized %v node directories\n", nValidators+nNonValidators)
	return nil
}

func hostnameOrIP(i int) (ip string, port int) {
	if listIPAddress != "" {
		ips := strings.Split(listIPAddress, ",")

		// i is index, so ips has at least i+1 elements
		if len(ips) <= i {
			fmt.Printf("%v: list-ip-address doesn't have enough items\n", listIPAddress)
			os.Exit(1)
		}

		selectedIP := ips[i]

		elems := strings.Split(selectedIP, ":")

		if len(elems) != 2 {
			fmt.Printf("%v: is not in format IP:PORT\n", selectedIP)
			os.Exit(1)
		}

		ipN := net.ParseIP(elems[0])
		ipN = ipN.To4()

		if ipN == nil {
			fmt.Printf("%v: non ipv4 address\n", elems[0])
			os.Exit(1)
		}

		port, err := strconv.Atoi(elems[1])
		if err != nil {
			fmt.Printf("%v: is not a valid integer port\n", elems[1])
			os.Exit(1)
		}

		return ipN.String(), port

	} else if startingIPAddress != "" {
		ip := net.ParseIP(startingIPAddress)
		ip = ip.To4()
		if ip == nil {
			fmt.Printf("%v: non ipv4 address\n", startingIPAddress)
			os.Exit(1)
		}

		for j := 0; j < i; j++ {
			ip[3]++
		}
		return ip.String(), p2pPort
	}

	return fmt.Sprintf("%s%d", hostnamePrefix, i), p2pPort
}

func getMoniker(i int) string {
	if listMonikers != "" {
		monikers := strings.Split(listMonikers, ",")
		if len(monikers) <= i {
			fmt.Printf("%v: doesn't have enough monikers.", listMonikers)
			os.Exit(1)
		}

		return monikers[i]
	}
	return cmn.Fmt("%s%d", nodeDirPrefix, i)

}

func populatePersistentPeersInConfigAndWriteIt(config *cfg.Config) error {
	persistentPeers := make([]string, nValidators+nNonValidators)
	for i := 0; i < nValidators+nNonValidators; i++ {
		var nodeDir string
		if i < nValidators {
			nodeDir = filepath.Join(outputDir, getMoniker(i))
		} else {
			nodeDir = filepath.Join(outputDir, cmn.Fmt("%s%d", nodeDirPrefix, i))
		}

		config.SetRoot(nodeDir)
		nodeKey, err := p2p.LoadNodeKey(config.NodeKeyFile())
		if err != nil {
			return err
		}
		ip, port := hostnameOrIP(i)
		persistentPeers[i] = p2p.IDAddressString(nodeKey.ID(), fmt.Sprintf("%s:%d", ip, port))
	}
	persistentPeersList := strings.Join(persistentPeers, ",")

	for i := 0; i < nValidators+nNonValidators; i++ {
		var nodeDir string
		if i < nValidators {
			nodeDir = filepath.Join(outputDir, getMoniker(i))
		} else {
			nodeDir = filepath.Join(outputDir, cmn.Fmt("%s%d", nodeDirPrefix, i))
		}

		config.SetRoot(nodeDir)
		config.P2P.PersistentPeers = persistentPeersList
		config.P2P.AddrBookStrict = false

		moniker := getMoniker(i)
		if moniker != "" {
			config.BaseConfig.Moniker = moniker
		}

		// overwrite default config
		cfg.WriteConfigFile(filepath.Join(nodeDir, "config", "config.toml"), config)
	}

	return nil
}
