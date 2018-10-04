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

func generateGenesisState(filePath string) app.GenesisState {
	pv := privval.LoadFilePV(filePath)

	privK, ok := pv.PrivKey.(crypto.PrivKeySecp256k1)

	if !ok {
		panic(ok)
	}

	privKey := types.NewPrivKeySecp256k1(privK[:])
	pubKey := privKey.PubKey()

	gs := app.GenerateGenesisState(pubKey)
	return gs
}

func generateGenesisFile(filePath string, genesisState []byte) {
	genesisDoc, err := tmtypes.GenesisDocFromFile(filePath)
	if err != nil {
		panic(err)
	}

	genesisDoc.AppStateJSON = json.RawMessage(genesisState)

	genesisDoc.SaveAs(filePath)
}

func main() {
	home := os.Getenv("HOME")
	genesisState := generateGenesisState(home + "/.tendermint/config/priv_validator.json")
	generateGenesisFile(home+"/.tendermint/config/genesis.json", genesisState.ToJSON())
}
