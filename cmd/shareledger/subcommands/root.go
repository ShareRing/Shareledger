package subcommands

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	cfg "github.com/tendermint/tendermint/config"
	"github.com/tendermint/tmlibs/log"
)

var (
	configDir  = "/.shareledger"
	defaultDir = os.Getenv("HOME") + configDir
	config     = cfg.DefaultConfig()
	logger     = log.NewTMLogger(log.NewSyncWriter(os.Stdout))
)

const (
	defaultLogLevelKey = "*"
	HomeFlag           = "home"
	RPCListenAddress   = "tcp://0.0.0.0:46657"
	P2PListenAddress   = "tcp://0.0.0.0:46656"
	BaseConfigProxyApp = "tcp://127.0.0.1:46658"
	ConfigDir          = "config"
	RootFile           = "config.toml"
)

var RootCmd = &cobra.Command{
	Use:   "shareledger",
	Short: "Shareledger is a distributed blockchain for sharing services",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
		// Bind Flags and read config from files
		// Command line args override file config
		bindFlagsLoadViper(cmd, args)

		// Unmarshal viper to config struct
		config, err = ParseConfig()

		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	RootCmd.PersistentFlags().String("log_level", "shareledger:info,"+config.LogLevel, "Log Level")
	RootCmd.PersistentFlags().String(HomeFlag, defaultDir, "Home directory where to store all configuration files and data")
	config.P2P.ListenAddress = P2PListenAddress
	config.RPC.ListenAddress = RPCListenAddress
	config.BaseConfig.ProxyApp = BaseConfigProxyApp
}

// Bind all flags and read the config into viper
func bindFlagsLoadViper(cmd *cobra.Command, args []string) error {
	// cmd.Flags() includes flags from this command and all persistent flags from the parent
	if err := viper.BindPFlags(cmd.Flags()); err != nil {
		return err
	}

	homeDir := viper.GetString(HomeFlag)
	viper.Set(HomeFlag, homeDir)
	viper.SetConfigName("config")                          // name of config file (without extension)
	viper.AddConfigPath(homeDir)                           // search root directory
	viper.AddConfigPath(filepath.Join(homeDir, ConfigDir)) // search root directory /config

	// Ensure Root
	config.RootDir = homeDir
	config.SetRoot(homeDir)
	cfg.EnsureRoot(homeDir)

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {

		// stderr, so if we redirect output to json file, this doesn't appear
		// fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	} else if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
		// ignore not found error, return other errors
		return err
	}

	return nil
}

func ParseConfig() (*cfg.Config, error) {
	err := viper.Unmarshal(config)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil, err
	}

	return config, err
}
