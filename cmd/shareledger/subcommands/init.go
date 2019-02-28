package subcommands

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"strconv"
	"time"

	"github.com/spf13/cobra"

	cfg "github.com/tendermint/tendermint/config"
	crypto "github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/secp256k1"
	cmn "github.com/tendermint/tendermint/libs/common"
	"github.com/tendermint/tendermint/p2p"
	"github.com/tendermint/tendermint/privval"
	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/sharering/shareledger/app"
	"github.com/sharering/shareledger/types"
)

// InitFilesCmd initialises a fresh Tendermint Core instance.
var InitFilesCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize essential files",
	RunE:  initFiles,
}

const (
	listenAddress string = "tcp://0.0.0.0:"
)

var (
	persistentPeers string
	privKeyParam    string
)

func init() {
	InitFilesCmd.Flags().StringVar(&persistentPeers, "persistent-peers", "",
		"List of persistent peers.")
	InitFilesCmd.Flags().IntVar(&rpcPort, "rpc-port", 46657, "RPC listening port.")
	InitFilesCmd.Flags().IntVar(&p2pPort, "p2p-port", 46656, "P2P listening port.")
	InitFilesCmd.Flags().StringVar(&moniker, "moniker", "default-moniker",
		"Unique name of your Masternode.")
	InitFilesCmd.Flags().StringVar(&privKeyParam, "privKey", "", "Generate using an existing private key.")
}

func initFiles(cmd *cobra.Command, args []string) error {
	return initFilesWithConfig(config)
}

func initFilesWithConfig(config *cfg.Config) error {
	// private validator
	privValKeyFile := config.PrivValidatorKeyFile()
	logger.Info("PrivateValidator KeyFile:", "filePath", privValKeyFile)

	privValStateFile := config.PrivValidatorStateFile()
	logger.Info("PrivateValidator StateFile:", "filePath", privValStateFile)

	var pv *privval.FilePV
	if cmn.FileExists(privValKeyFile) {
		pv = privval.LoadFilePV(privValKeyFile, privValStateFile)
		logger.Info("Found private validator", "path", privValKeyFile)
	} else {
		pv = privval.GenFilePV(privValKeyFile, privValStateFile)
		var newPrivKey secp256k1.PrivKeySecp256k1

		if privKeyParam == "" {
			// newPrivKey = secp256k1.GenPrivKeySecp256k1([]byte{})
			newPrivKey = secp256k1.GenPrivKey()
		} else {
			newPrivKey = types.GetCryptoPrivKey(privKeyParam)
		}
		pv.Key.PrivKey = newPrivKey
		pv.Key.PubKey = pv.Key.PrivKey.PubKey()
		pv.Key.Address = pv.Key.PubKey.Address()

		pv.Save()
		logger.Info("Generated private validator", "path", privValKeyFile)
	}

	nodeKeyFile := config.NodeKeyFile()
	if cmn.FileExists(nodeKeyFile) {
		logger.Info("Found node key", "path", nodeKeyFile)
	} else {
		if _, err := p2p.LoadOrGenNodeKey(nodeKeyFile); err != nil {
			return err
		}
		logger.Info("Generated node key", "path", nodeKeyFile)
	}

	//Replace ed25519 with secp256k1
	// return corresponding app.GenesisState
	genesisState, pubKey := genGenesisState(pv)

	// genesis file
	genFile := config.GenesisFile()
	if cmn.FileExists(genFile) {
		logger.Info("Found genesis file", "path", genFile)
	} else {
		consensusParams := tmtypes.DefaultConsensusParams()
		consensusParams.Validator = tmtypes.ValidatorParams{[]string{tmtypes.ABCIPubKeyTypeSecp256k1}}

		genDoc := tmtypes.GenesisDoc{
			ChainID:         fmt.Sprintf("test-chain-%v", cmn.RandStr(6)),
			GenesisTime:     time.Now(),
			ConsensusParams:	 consensusParams,
		}
		genDoc.Validators = []tmtypes.GenesisValidator{{
			Address: pubKey.Address(),
			PubKey:  pubKey,
			Power:   genesisState.StakeData.Validators[0].ABCIValidator().Power,
			Name:    genesisState.StakeData.Validators[0].Description.Moniker,
		}}

		genDoc.AppState = json.RawMessage(genesisState.ToJSON())

		if err := genDoc.SaveAs(genFile); err != nil {
			return err
		}
		logger.Info("Generated genesis file", "path", genFile)
	}

	// Update RP2 & P2P port
	logger.Info("PersistentPeers", "peers", persistentPeers)
	config.BaseConfig.Moniker = moniker
	config.RPC.ListenAddress = listenAddress + strconv.Itoa(rpcPort)
	config.P2P.ListenAddress = listenAddress + strconv.Itoa(p2pPort)
	config.P2P.PersistentPeers = persistentPeers

	// Rewrite config file
	path := filepath.Join(config.BaseConfig.RootDir, ConfigDir, RootFile)
	cfg.WriteConfigFile(path, config)

	logger.Info("Generated config file", "path", path)

	return nil
}

func genGenesisState(pv *privval.FilePV) (app.GenesisState, crypto.PubKey) {
	// Change Ed25519 to Secp256k1
	// save new priv_validator.json
	pv.Save()

	privK, ok := pv.Key.PrivKey.(secp256k1.PrivKeySecp256k1)

	if !ok {
		panic(ok)
	}

	// privKey in ShareLedger PrivKeySecp256k1
	privKey := types.NewPrivKeySecp256k1(privK[:])
	pubKey := privKey.PubKey()

	tpk, ok := pv.Key.PubKey.(secp256k1.PubKeySecp256k1)
	if !ok {
		panic(ok)
	}

	// tpk1 := types.ConvertToPubKey(tpk[:])

	// fmt.Printf("Shareledger PubKey: %x\n", pubKey[:])
	// fmt.Printf("Tendermint PubKey : %x\n", tpk1[:])

	// fmt.Printf("Shareledger address: %x\n", pubKey.Address()[:])
	// fmt.Printf("Bech32 address     : %s\n", pubKey.Address().String())
	// fmt.Printf("Tendermint address : %x\n", pv.Key.PubKey.Address()[:])

	// fmt.Printf("Our TDM address    : %x\n", types.ConvertToTDMPubKey(pubKey).Address()[:])

	gs := app.GenerateGenesisState(pubKey)
	return gs, pv.Key.PubKey
}
