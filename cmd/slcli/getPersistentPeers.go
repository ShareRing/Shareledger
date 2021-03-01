package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

const (
	flagPeerFile = "peer-file"
)

func getPersistentPeer() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-peers [nodeId file] [ip] [first port] [peer file]",
		Short: "create persistent peers file from nodeId file",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			data, err := ioutil.ReadFile(args[0])
			if err != nil {
				return err
			}
			var nodeIDs []string
			err = json.Unmarshal(data, &nodeIDs)
			port0, err := strconv.Atoi(args[2])
			if err != nil {
				return err
			}
			var peerAddresses []string
			for i, id := range nodeIDs {
				port := port0 + i*2
				peerAddr := fmt.Sprintf("%s@%s:%d", id, args[1], port)
				peerAddresses = append(peerAddresses, peerAddr)
			}
			peerStr := strings.Join(peerAddresses, ",")
			peerdata, err := json.Marshal(peerStr)
			err = ioutil.WriteFile(args[3], peerdata, 0600)
			if err != nil {
				return err
			}
			var myStr string
			mydata, err := ioutil.ReadFile(args[3])
			if err != nil {
				return err
			}
			if err := json.Unmarshal(mydata, &myStr); err != nil {
				return err
			}
			return nil
		},
	}
	return cmd
}
