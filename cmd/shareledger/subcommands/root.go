package subcommands

import (
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
}

// Bind all flags and read the config into viper
func bindFlagsLoadViper(cmd *cobra.Command, args []string) error {
	// cmd.Flags() includes flags from this command and all persistent flags from the parent
	if err := viper.BindPFlags(cmd.Flags()); err != nil {
		return err
	}

	homeDir := viper.GetString(HomeFlag)
	viper.Set(HomeFlag, homeDir)
	viper.SetConfigName("config")                         // name of config file (without extension)
	viper.AddConfigPath(homeDir)                          // search root directory
	viper.AddConfigPath(filepath.Join(homeDir, "config")) // search root directory /config

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
	conf := cfg.DefaultConfig()

	err := viper.Unmarshal(conf)

	if err != nil {
		return nil, err
	}

	conf.SetRoot(conf.RootDir)
	cfg.EnsureRoot(conf.RootDir)
	return conf, err
}
