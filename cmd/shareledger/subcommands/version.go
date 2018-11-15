package subcommands

import (
	"fmt"

	"github.com/spf13/cobra"

	tmversion "github.com/tendermint/tendermint/version"
	"github.com/sharering/shareledger/version"
)

// VersionCmd ...
var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version info",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("shareledger-"+version.Version, "tendermint-"+tmversion.Version)
	},
}




