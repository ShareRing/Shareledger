package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/tendermint/abci/server"
	"github.com/tendermint/tmlibs/cli"
	cmn "github.com/tendermint/tmlibs/common"
	dbm "github.com/tendermint/tmlibs/db"
	"github.com/tendermint/tmlibs/log"

	"github.com/sharering/shareledger/app"
	"github.com/sharering/shareledger/cmd/shareledger/subcommands"
	"github.com/sharering/shareledger/constants"
)

func main() {
	cmd := constructCommand()
	cmd.Execute()
}

func constructCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "shareledger",
		Short: "Shareledger is distributed blockchain for sharing services",
		RunE:  startServer,
	}

	rootCmd.Flags().StringP(constants.LogLevel, constants.LogLevelS, "info", "Logging level: info, debug, warn, error")
	rootCmd.Flags().StringP(constants.HostFlag, constants.HostFlagS, constants.HOST, "Host")
	rootCmd.Flags().IntP(constants.PortFlag, constants.PortFlagS, constants.PORT, "Port")

	rootCmd.AddCommand(
		subcommands.InitFilesCmd,
		subcommands.ShowNodeIDCmd,
		subcommands.ResetAllCmd,
	)

	return rootCmd
}

func initLogger(lvl string) log.Logger {
	logger := log.NewTMLogger(log.NewSyncWriter(os.Stdout)).With("module", "main")

	// Only allow some level
	opt, err := log.AllowLevel(strings.ToLower(lvl))

	if err != nil {
		panic(err)
	}

	logger = log.NewFilter(logger, opt)

	// assign global variable LOGGER to log all input
	constants.LOGGER = logger

	return logger
}

func startServer(cmd *cobra.Command, args []string) error {

	logLvl, _ := cmd.Flags().GetString(constants.LogLevel)

	logger := initLogger(logLvl)

	rootDir := viper.GetString(cli.HomeFlag)

	db, err := dbm.NewGoLevelDB("shareledgerd", filepath.Join(rootDir, "data"))

	shareledgerApp := app.NewShareLedgerApp(logger, db)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	logger.Info("Start ShareLedger", "Host", constants.HOST, "Port", constants.PORT)

	// Start the ABCI server
	srv, err := server.NewServer(strings.Join([]string{
		constants.HOST,
		":",
		strconv.Itoa(constants.PORT)}, ""),
		"socket",
		shareledgerApp)

	if err != nil {
		logger.Error(fmt.Sprintf("Start server failing. Error: %s", err))
		os.Exit(1)
	}

	constants.LOGGER.Info("Testing")

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
	return nil
}
