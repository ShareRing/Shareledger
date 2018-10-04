package main

import (
	"encoding/json"
	"os"

	"github.com/tendermint/go-crypto"
	"github.com/tendermint/tendermint/privval"
	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/sharering/shareledger/app"
	"github.com/sharering/shareledger/types"
)

const (
	ConfigDir            string = "/.tendermint/config/"
	PrivateValidatorFile string = "priv_validator.json"
	GenesisFile          string = "genesis.json"
)

// Read priv_validator.jso
// Replace ed25519 with secp256k1
// return corresponding app.GenesisState
func generateGenesisState(filePath string) (app.GenesisState, crypto.PubKey) {
	pv := privval.LoadFilePV(filePath)

	// Change Ed25519 to Secp256k1
	newPrivKey := crypto.GenPrivKeySecp256k1()
	pv.PrivKey = newPrivKey
	pv.PubKey = pv.PrivKey.PubKey()
	pv.Address = pv.PubKey.Address()

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

// Update GenesisFile with new Key
// Update GenesisFile with AppState
// Save new GenesisFile
func generateGenesisFile(filePath string, genesisState []byte, pubKey crypto.PubKey) {
	genesisDoc, err := tmtypes.GenesisDocFromFile(filePath)
	if err != nil {
		panic(err)
	}

	// Update genesisDoc with new validator
	genesisDoc.Validators = []tmtypes.GenesisValidator{{
		PubKey: pubKey,
		Power:  10,
	}}

	genesisDoc.AppStateJSON = json.RawMessage(genesisState)

	genesisDoc.SaveAs(filePath)
}

func main() {
	homeDir := os.Getenv("HOME")
	// reading priv_validator
	genesisState, pubKey := generateGenesisState(homeDir + ConfigDir + PrivateValidatorFile)

	// update genesis accordingly
	generateGenesisFile(homeDir+ConfigDir+GenesisFile, genesisState.ToJSON(), pubKey)
}
