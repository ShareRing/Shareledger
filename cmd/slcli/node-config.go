package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tendermint/tendermint/config"
	"github.com/tendermint/tendermint/p2p"
)

const (
	flagMoniker       = "moniker"
	flagP2PPort       = "p2p-port"
	flagStartingIP    = "starting-ip"
	flagRPCPort       = "rpc-port"
	flagAllPeers      = "all-peers"
	persistentPeerNum = 3
)

type persistentPeer struct {
	ID   string `json:"id"`
	IP   string `json:"ip"`
	Port int    `json:"port"`
}

func (p persistentPeer) String() string {
	return fmt.Sprintf("%s@%s:%d", p.ID, p.IP, p.Port)
}

func createNodeConfig() *cobra.Command {
	cmd := &cobra.Command{
		Use: "create-node-config",
		RunE: func(cmd *cobra.Command, args []string) error {
			home := viper.GetString(flagHome)
			configPath := filepath.Join(home, daemonHome, "config")
			persistentPeers, err := choosePersistentPeers()
			if err != nil {
				return err
			}
			cfg := getBasicConfig(configPath)
			cfg.P2P.PersistentPeers = persistentPeers
			config.WriteConfigFile(filepath.Join(configPath, "config.toml"), cfg)
			return nil
		},
	}
	cmd.Flags().String(flagHome, defaultHome, "home for the daemon")
	cmd.Flags().String(flagAllPeers, "./all_peers.json", "path for persistent peers file")
	cmd.Flags().String(flagMoniker, "monica", "moniker name")
	cmd.Flags().String(flagRPCPort, "", "port for rpc listen addr")
	cmd.Flags().String(flagP2PPort, "", "port for p2p listen addr")
	return cmd
}

func updateConfig() *cobra.Command {
	cmd := &cobra.Command{
		Use: "update-config",
		RunE: func(cmd *cobra.Command, args []string) error {
			persistentPeers, err := choosePersistentPeers()
			if err != nil {
				return err
			}
			// writing peers to config
			configPath := filepath.Join(viper.GetString(flagHome), daemonHome, "config")
			cfg := getBasicConfig(configPath)
			cfg.P2P.PersistentPeers = persistentPeers
			config.WriteConfigFile(filepath.Join(configPath, "config.toml"), cfg)
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.Flags().String(flagHome, defaultHome, "node's home")
	cmd.Flags().String(flagAllPeers, "./all_peers.json", "path for list of all peers file")
	cmd.Flags().String(flagMoniker, "monica", "moniker name")
	cmd.Flags().String(flagRPCPort, "", "port for rpc listen addr")
	cmd.Flags().String(flagP2PPort, "", "port for p2p listen addr")
	return cmd
}

func getBasicConfig(configPath string) *config.Config {
	cfg := config.DefaultConfig()
	cfg.SetRoot(configPath)
	cfg.BaseConfig.Moniker = viper.GetString(flagMoniker)
	cfg.BaseConfig.DBBackend = "cleveldb"
	cfg.RPC.ListenAddress = "tcp://0.0.0.0:26657"
	rpcPort := viper.GetString(flagRPCPort)
	if rpcPort != "" {
		cfg.RPC.ListenAddress = "tcp://0.0.0.0:" + rpcPort

	}
	p2pPort := viper.GetString(flagP2PPort)
	if p2pPort != "" {
		cfg.P2P.ListenAddress = "tcp://0.0.0.0:" + p2pPort
	}

	cfg.ProfListenAddress = "localhost:6060"
	cfg.P2P.RecvRate = 5120000
	cfg.P2P.SendRate = 5120000
	cfg.TxIndex.IndexAllKeys = true
	cfg.Consensus.TimeoutCommit = 5 * time.Second
	// cfg.RPC.CORSAllowedOrigins = []string{"*"}
	return cfg
}

func choosePersistentPeers() (string, error) {
	nodeKeyPath := filepath.Join(viper.GetString(flagHome), daemonHome, "config/node_key.json")
	nodeKey, err := p2p.LoadNodeKey(nodeKeyPath)
	if err != nil {
		return "", err
	}
	nodeID := fmt.Sprintf("%s", nodeKey.ID())

	// reading persitent peers from file
	allPeers := []persistentPeer{}
	excludedIDs := map[string]struct{}{nodeID: {}}

	data, err := ioutil.ReadFile(viper.GetString(flagAllPeers))
	if err != nil {
		return "", nil
	}

	if err := json.Unmarshal(data, &allPeers); err != nil {
		return "", nil
	}
	if len(allPeers) < persistentPeerNum {
		return "", nil
	}

	// get 3 unique persitent peers excluding the node itself
	mypeers := []persistentPeer{}
	for i := 0; i < persistentPeerNum; {
		rand.Seed(time.Now().UTC().UnixNano())
		myrand := rand.Intn(len(allPeers))
		id := allPeers[myrand].ID
		if _, ok := excludedIDs[id]; !ok {
			excludedIDs[id] = struct{}{}
			i++
			mypeers = append(mypeers, allPeers[myrand])
		} else {
		}
	}

	peerList := []string{}
	for _, peer := range mypeers {
		peerList = append(peerList, peer.String())
	}

	return strings.Join(peerList, ","), nil
}

func getPeersFromNodekeys() *cobra.Command {
	cmd := &cobra.Command{
		Use: "get-peers-from-node-keys [node-key-1] [node-key-2] [node-key-n]",
		RunE: func(cmd *cobra.Command, args []string) error {
			var peers []persistentPeer
			port, err := strconv.Atoi(viper.GetString(flagP2PPort))
			if err != nil {
				return err
			}
			for i, a := range args {
				nodeKey, err := p2p.LoadNodeKey(a)
				if err != nil {
					return err
				}
				nodeID := fmt.Sprintf("%s", nodeKey.ID())
				ipv4 := net.ParseIP(viper.GetString(flagStartingIP)).To4()
				for j := 0; j < i; j++ {
					ipv4[3]++
				}
				peers = append(peers, persistentPeer{nodeID, ipv4.String(), port})
			}
			data, err := json.Marshal(peers)
			if err != nil {
				return err
			}
			err = ioutil.WriteFile(viper.GetString(flagAllPeers), data, 0600)
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.Flags().String(flagAllPeers, "./all_peers_2.json", "path for list of all peers file")
	cmd.Flags().String(flagP2PPort, "26656", "port for p2p listen addr")
	cmd.Flags().String(flagStartingIP, "172.194.4.2", "starting ip")
	return cmd
}

func getPeerFromConfig() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "get-peer-from-config [config1] [config2]",
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			peerMap := make(map[string]string)
			var peers []persistentPeer
			cfgFile := viper.New()
			for _, a := range args {
				cfgFile.SetConfigFile(a)
				err := cfgFile.ReadInConfig()
				if err != nil {
					return err
				}
				err = addPersitentPeers(peerMap, cfgFile)
				if err != nil {
					return err
				}
			}
			for k, v := range peerMap {
				s := strings.Split(v, ":")
				port, err := strconv.Atoi(s[1])
				if err != nil {
					return err
				}
				peer := persistentPeer{k, s[0], port}
				peers = append(peers, peer)
			}
			data, err := json.Marshal(peers)
			if err != nil {
				return err
			}
			err = ioutil.WriteFile("./all_peers_2.json", data, 0600)
			if err != nil {
				return err
			}
			return nil
		},
	}
	return cmd
}

func addPersitentPeers(peers map[string]string, configFile *viper.Viper) error {
	s, err := getPersistentPeers(configFile)
	if err != nil {
		return err
	}
	for _, peer := range s {
		p := strings.Split(peer, "@")
		peers[p[0]] = p[1]
	}
	return nil
}

func getPersistentPeers(configFile *viper.Viper) ([]string, error) {
	s := configFile.Get("p2p.persistent_peers")
	r, ok := s.(string)
	if !ok {
		return nil, fmt.Errorf("Persistent_peers is not string %v", s)
	}
	return strings.Split(r, ","), nil
}
