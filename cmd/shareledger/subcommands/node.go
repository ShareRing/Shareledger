package subcommands

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	abci "github.com/tendermint/tendermint/abci/types"
	cfg "github.com/tendermint/tendermint/config"
	cmn "github.com/tendermint/tendermint/libs/common"
	dbm "github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"
	"github.com/tendermint/tendermint/node"
	"github.com/tendermint/tendermint/p2p"
	pvm "github.com/tendermint/tendermint/privval"
	"github.com/tendermint/tendermint/proxy"

	"github.com/sharering/shareledger/app"
	"github.com/sharering/shareledger/constants"
)

var NodeCmd = constructCommand()

func constructCommand() *cobra.Command {

	nodeCmd := &cobra.Command{
		Use:   "node",
		Short: "ShareLedger - a custom-designed distributed blockchain",

		RunE: startCombo,
	}

	// nodeCmd.Flags().StringP(constants.LogLevel, constants.LogLevelS, "info", "Logging level: info, debug, warn, error")

	// bind flags
	nodeCmd.Flags().String("moniker", config.Moniker, "Node Name")

	// priv val flags
	// nodeCmd.Flags().String("priv_validator_laddr", config.PrivValidatorListenAddr, "Socket address to listen on for connections from external priv_validator process")

	// node flags
	// nodeCmd.Flags().Bool("fast_sync", config.FastSync, "Fast blockchain syncing")

	// abci flags
	// nodeCmd.Flags().String("proxy_app", config.ProxyApp, "Proxy app address, or 'nilapp' or 'kvstore' for local testing.")
	// nodeCmd.Flags().String("abci", config.ABCI, "Specify abci transport (socket | grpc)")

	// rpc flags
	nodeCmd.Flags().String("rpc.laddr", RPCListenAddress, "RPC listen address. Port required")
	// nodeCmd.Flags().String("rpc.grpc_laddr", config.RPC.GRPCListenAddress, "GRPC listen address (BroadcastTx only). Port required")
	// nodeCmd.Flags().Bool("rpc.unsafe", config.RPC.Unsafe, "Enabled unsafe rpc methods")

	// p2p flags
	nodeCmd.Flags().String("p2p.laddr", P2PListenAddress, "Node listen address. (0.0.0.0:0 means any interface, any port)")
	// nodeCmd.Flags().String("p2p.seeds", config.P2P.Seeds, "Comma-delimited ID@host:port seed nodes")
	// nodeCmd.Flags().String("p2p.persistent_peers", config.P2P.PersistentPeers, "Comma-delimited ID@host:port persistent peers")
	// nodeCmd.Flags().Bool("p2p.skip_upnp", config.P2P.SkipUPNP, "Skip UPNP configuration")
	// nodeCmd.Flags().Bool("p2p.pex", config.P2P.PexReactor, "Enable/disable Peer-Exchange")
	// nodeCmd.Flags().Bool("p2p.seed_mode", config.P2P.SeedMode, "Enable/disable seed mode")
	// nodeCmd.Flags().String("p2p.private_peer_ids", config.P2P.PrivatePeerIDs, "Comma-delimited private peer IDs")

	// consensus flags
	// nodeCmd.Flags().Bool("consensus.create_empty_blocks", config.Consensus.CreateEmptyBlocks, "Set this to false to only produce blocks when there are txs or when the AppHash changes")

	return nodeCmd
}

func initLogger() log.Logger {
	logger := log.NewTMLogger(log.NewSyncWriter(os.Stdout)).With("module", "shareledger")

	// Only allow some level
	// opt, err := log.AllowLevel(strings.ToLower(lvl))

	// if err != nil {
	// 	panic(err)
	// }

	// logger = log.NewFilter(logger, opt)

	// assign global variable LOGGER to log all input
	constants.LOGGER = logger

	return logger
}

func newApp(logger log.Logger, db dbm.DB) abci.Application {
	return app.NewShareLedgerApp(logger, db)
}

func startCombo(cmd *cobra.Command, args []string) error {

	logger := initLogger()

	// Adjust logger with each module filter
	logger, err := ParseLogLevel(config.LogLevel, logger, cfg.DefaultLogLevel())
	if err != nil {
		return err
	}

	rootDir := viper.GetString(HomeFlag)

	db, err := dbm.NewGoLevelDB("shareledgerd", filepath.Join(rootDir, "data"))

	nodeKey, err := p2p.LoadOrGenNodeKey(config.NodeKeyFile())
	if err != nil {
		return err
	}

	shareledgerApp := app.NewShareLedgerApp(logger, db)

	n, err := node.NewNode(
		config,
		pvm.LoadOrGenFilePV(config.PrivValidatorKeyFile() , config.PrivValidatorStateFile()),
		nodeKey,
		proxy.NewLocalClientCreator(shareledgerApp),
		node.DefaultGenesisDocProviderFunc(config),
		node.DefaultDBProvider,
		node.DefaultMetricsProvider(config.Instrumentation),
		logger,
	)

	if err != nil {
		return err
	}

	err = n.Start()
	if err != nil {
		return err
	}

	cmn.TrapSignal(func() {
		if n.IsRunning() {
			_ = n.Stop()
		}
	})

	// run forever (the node will not be returned)
	select {}
}

// ParseLogLevel parses complex log level - comma-separated
// list of module:level pairs with an optional *:level pair (* means
// all other modules).
//
// Example:
//		ParseLogLevel("consensus:debug,mempool:debug,*:error", log.NewTMLogger(os.Stdout), "info")
func ParseLogLevel(lvl string, logger log.Logger, defaultLogLevelValue string) (log.Logger, error) {
	if lvl == "" {
		return nil, errors.New("Empty log level")
	}

	l := lvl

	// prefix simple one word levels (e.g. "info") with "*"
	if !strings.Contains(l, ":") {
		l = defaultLogLevelKey + ":" + l
	}

	options := make([]log.Option, 0)

	isDefaultLogLevelSet := false
	var option log.Option
	var err error

	list := strings.Split(l, ",")
	for _, item := range list {
		moduleAndLevel := strings.Split(item, ":")

		if len(moduleAndLevel) != 2 {
			return nil, fmt.Errorf("Expected list in a form of \"module:level\" pairs, given pair %s, list %s", item, list)
		}

		module := moduleAndLevel[0]
		level := moduleAndLevel[1]

		if module == defaultLogLevelKey {
			option, err = log.AllowLevel(level)
			if err != nil {
				return nil, errors.Wrap(err, fmt.Sprintf("Failed to parse default log level (pair %s, list %s)", item, l))
			}
			options = append(options, option)
			isDefaultLogLevelSet = true
		} else {
			switch level {
			case "debug":
				option = log.AllowDebugWith("module", module)
			case "info":
				option = log.AllowInfoWith("module", module)
			case "error":
				option = log.AllowErrorWith("module", module)
			case "none":
				option = log.AllowNoneWith("module", module)
			default:
				return nil, fmt.Errorf("Expected either \"info\", \"debug\", \"error\" or \"none\" log level, given %s (pair %s, list %s)", level, item, list)
			}
			options = append(options, option)

		}
	}

	// if "*" is not provided, set default global level
	if !isDefaultLogLevelSet {
		option, err = log.AllowLevel(defaultLogLevelValue)
		if err != nil {
			return nil, err
		}
		options = append(options, option)
	}

	return log.NewFilter(logger, options...), nil
}
