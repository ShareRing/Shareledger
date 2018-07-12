package main

import (
	"fmt"
	"os"

	"github.com/tendermint/abci/server"
	"github.com/tendermint/tmlibs/cli"
	cmn "github.com/tendermint/tmlibs/common"
	dbm "github.com/tendermint/tmlibs/db"
	"github.com/tendermint/tmlibs/log"

	"github.com/sharering/shareledger/app"
	//"github.com/cosmos/cosmos-sdk/server"
	"github.com/spf13/viper"
	"path/filepath"
)

func main() {
	fi, err := os.Open("./temp.txt")
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()

	logger := log.NewTMLogger(log.NewSyncWriter(fi)).With("module", "main")
	//logger := log.NewTMLogger(log.NewSyncWriter(os.Stdout)).With("module", "main")

	rootDir := viper.GetString(cli.HomeFlag)
	db, err := dbm.NewGoLevelDB("shareledgerd", filepath.Join(rootDir, "data"))

	shareledgerApp := app.NewShareLedgerApp(logger, db)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Start the ABCI server
	srv, err := server.NewServer("0.0.0.0:46658", "socket", shareledgerApp)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = srv.Start()
	if err != nil {
		cmn.Exit(err.Error())
	}

	// Wait forever
	cmn.TrapSignal(func() {
		// Cleanup
		err = srv.Stop()
		if err != nil {
			cmn.Exit(err.Error())
		}
	})
	return
}
