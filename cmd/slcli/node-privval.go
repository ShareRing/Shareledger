package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tendermint/tendermint/p2p"
	"github.com/tendermint/tendermint/privval"
)

const (
	flagHome     = "home"
	flagNodeFile = "node-file"
	defaultHome  = "./"
	daemonHome   = "shareledger"
	nodeDirPerm  = 0755
)

func createNodeKeysCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-node-keys",
		Short: "generate keys and files for nodes and validators",
		Long:  "for node: public key, private key, nodeID, for validator: private key and store it in keybase",
		RunE: func(cmd *cobra.Command, args []string) error {
			homeDir := viper.GetString(flagHome)
			daemonHome := filepath.Join(homeDir, daemonHome)
			ctx := server.NewDefaultContext()
			config := ctx.Config
			config.SetRoot(daemonHome)
			if err := os.MkdirAll(filepath.Join(daemonHome, "config"), nodeDirPerm); err != nil {
				os.RemoveAll(daemonHome)
				return err
			}
			nodeID, valPubKey, err := genutil.InitializeNodeValidatorFiles(config)
			if err != nil {
				os.RemoveAll(daemonHome)
				return err
			}
			fmt.Println("\nnode_ID is", nodeID)
			bech32PubKey, err := sdk.Bech32ifyPubKey(sdk.Bech32PubKeyTypeConsPub, valPubKey)
			if err != nil {
				os.RemoveAll(daemonHome)
				return err
			}
			fmt.Println("\nprivval_pubkey:", bech32PubKey)
			return nil
		},
	}
	cmd.Flags().String(flagHome, defaultHome, "home for the cli")
	return cmd
}

func getNodeID() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-node-id",
		Short: "generate keys and files for nodes and validators",
		Long:  "for node: public key, private key, nodeID, for validator: private key and store it in keybase",
		RunE: func(cmd *cobra.Command, args []string) error {
			homeDir := viper.GetString(flagHome)
			daemonHome := filepath.Join(homeDir, daemonHome)
			ctx := server.NewDefaultContext()
			config := ctx.Config
			config.SetRoot(daemonHome)
			nodeKey, err := p2p.LoadNodeKey(config.NodeKeyFile())
			if err != nil {
				return err
			}
			fmt.Println("node Id is: ", nodeKey.ID())
			pubKey := nodeKey.PubKey()
			bech32PubKey, err := sdk.Bech32ifyPubKey(sdk.Bech32PubKeyTypeConsPub, pubKey)
			if err != nil {
				return err
			}
			fmt.Println("node pubkey is", bech32PubKey)
			return nil
		},
	}
	cmd.Flags().String(flagHome, defaultHome, "home for the cli")
	return cmd
}

func getNodeIDFromFile() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-node-id-from-file",
		Short: "generate keys and files for nodes and validators",
		Long:  "for node: public key, private key, nodeID, for validator: private key and store it in keybase",
		RunE: func(cmd *cobra.Command, args []string) error {
			nodeFile := viper.GetString(flagNodeFile)
			nodeKey, err := p2p.LoadNodeKey(nodeFile)
			if err != nil {
				return err
			}
			fmt.Println("node Id is: ", nodeKey.ID())
			pubKey := nodeKey.PubKey()
			bech32PubKey, err := sdk.Bech32ifyPubKey(sdk.Bech32PubKeyTypeConsPub, pubKey)
			if err != nil {
				return err
			}
			fmt.Println("node pubkey is", bech32PubKey)
			return nil
		},
	}
	cmd.Flags().String(flagNodeFile, "", "path to node key file")
	return cmd
}

func getPrivvalPubKey() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-privval-pubkey",
		Short: "get private validator public key from files in node config",
		RunE: func(cmd *cobra.Command, args []string) error {
			homeDir := viper.GetString(flagHome)
			daemonHome := filepath.Join(homeDir, daemonHome)
			ctx := server.NewDefaultContext()
			config := ctx.Config
			config.SetRoot(daemonHome)
			pv := privval.LoadFilePV(config.PrivValidatorKeyFile(), config.PrivValidatorStateFile())

			pubKey := pv.GetPubKey()
			bech32PubKey, err := sdk.Bech32ifyPubKey(sdk.Bech32PubKeyTypeConsPub, pubKey)
			if err != nil {
				return err
			}
			fmt.Println(bech32PubKey)
			return nil
		},
	}
	cmd.Flags().String(flagHome, defaultHome, "home for the cli")
	return cmd
}
