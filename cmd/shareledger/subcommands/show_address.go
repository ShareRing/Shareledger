package subcommands

import (
	"fmt"

	"github.com/spf13/cobra"

	// "github.com/tendermint/tendermint/p2p"
	"github.com/tendermint/tendermint/privval"
	cmn "github.com/tendermint/tmlibs/common"

	"github.com/sharering/shareledger/types"
)


// ShowPrivKeyCmd dumps node's Private Key to the standard output.
var ShowAddressCmd = &cobra.Command{
	Use:   "show_address",
	Short: "Show this node's address",
	RunE:  showAddress,
}


func showAddress(cmd *cobra.Command, args []string) error {
	privValFile := config.PrivValidatorFile()

	var pv *privval.FilePV

	if cmn.FileExists(privValFile) {
		pv = privval.LoadFilePV(privValFile)
		privateKey := types.ConvertToPrivKey(pv.PrivKey)

		fmt.Printf("%x\n", privateKey.PubKey().Address()[:])
		return nil

	}

	fmt.Printf("Private Validator File not found.")
	return nil
}
