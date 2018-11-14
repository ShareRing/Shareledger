package subcommands

import (
	"encoding/json"
	"os"
	"time"

	"github.com/spf13/cobra"

	"github.com/tendermint/go-crypto"
	cfg "github.com/tendermint/tendermint/config"
	"github.com/tendermint/tendermint/p2p"
	"github.com/tendermint/tendermint/privval"
	tmtypes "github.com/tendermint/tendermint/types"
	cmn "github.com/tendermint/tmlibs/common"
	"github.com/tendermint/tmlibs/log"

	"github.com/sharering/shareledger/app"
	"github.com/sharering/shareledger/types"
)

var (
	configDir = "/.shareledger"
	rootDir = os.Getenv("HOME") + configDir
	config = cfg.DefaultConfig().SetRoot(rootDir)
	logger = log.NewTMLogger(log.NewSyncWriter(os.Stdout))
)

// InitFilesCmd initialises a fresh Tendermint Core instance.
var InitFilesCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize essential files",
	RunE:  initFiles,
}

func init() {
	cfg.EnsureRoot(rootDir)
}


func initFiles(cmd *cobra.Command, args []string) error {
	return initFilesWithConfig(config)
}

func initFilesWithConfig(config *cfg.Config) error {
	// private validator
	privValFile := config.PrivValidatorFile()
	logger.Info("PrivateValidator File:", "filePath", privValFile)
	var pv *privval.FilePV
	if cmn.FileExists(privValFile) {
		pv = privval.LoadFilePV(privValFile)
		logger.Info("Found private validator", "path", privValFile)
	} else {
		pv = privval.GenFilePV(privValFile)
		newPrivKey := crypto.GenPrivKeySecp256k1()
		pv.PrivKey = newPrivKey
		pv.PubKey = pv.PrivKey.PubKey()
		pv.Address = pv.PubKey.Address()

		pv.Save()
		logger.Info("Generated private validator", "path", privValFile)
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
		genDoc := tmtypes.GenesisDoc{
			ChainID:     cmn.Fmt("test-chain-%v", cmn.RandStr(6)),
			GenesisTime: time.Now(),
		}
		genDoc.Validators = []tmtypes.GenesisValidator{{
			PubKey: pubKey,
			Power:  10,
		}}

		genDoc.AppStateJSON = json.RawMessage(genesisState.ToJSON())

		if err := genDoc.SaveAs(genFile); err != nil {
			return err
		}
		logger.Info("Generated genesis file", "path", genFile)
	}

	return nil
}

func genGenesisState(pv *privval.FilePV) (app.GenesisState, crypto.PubKey) {
	// Change Ed25519 to Secp256k1
		// save new priv_validator.json
	pv.Save()

	privK, ok := pv.PrivKey.(crypto.PrivKeySecp256k1)

	if !ok {
		panic(ok)
	}

	// privKey in ShareLedger PrivKeySecp256k1
	privKey := types.NewPrivKeySecp256k1(privK[:])
	pubKey := privKey.PubKey()

	gs := app.GenerateGenesisState(pubKey)
	return gs, pv.PubKey
}