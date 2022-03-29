package main

import (
	"github.com/sharering/shareledger/cmd/Shareledgerd/cli"
	"os"
	"sync"

	"github.com/sharering/shareledger/cmd/Shareledgerd/sub"
	"github.com/sharering/shareledger/cmd/Shareledgerd/tools"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/sharering/shareledger/app"
	"github.com/spf13/cobra"
	"github.com/tendermint/spm/cosmoscmd"
)

var runOnce sync.Once

func init() {
	runOnce.Do(func() {
		cli.InitMiddleWare()
	})
}

func main() {
	rootCmd, _ := cosmoscmd.NewRootCmd(
		app.Name,
		app.AccountAddressPrefix,
		app.DefaultNodeHome,
		app.Name,
		app.ModuleBasics,
		app.New,
		cosmoscmd.AddSubCmd(
			sub.NewImportKeyCmd(),
			getGenesisCmd(app.DefaultNodeHome),
		),
	)

	rootCmd.Short = "ShareRing-VoyagerNet"
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}

func getGenesisCmd(defaultNodeHome string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "genesis",
		Short: "genesis file subcommands",
	}
	cmd.AddCommand(
		tools.NewGenesisAddAuthorityAccountCmd(defaultNodeHome),
		tools.NewGenesisAddTreasureAccountCmd(defaultNodeHome),
		tools.NewGenesisAddValidatorAccountCmd(defaultNodeHome),
		tools.NewGenesisAddAccountOperatorCmd(defaultNodeHome),
	)
	return cmd
}
