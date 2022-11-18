package main

import (
	"os"
	"sync"

	"github.com/cosmos/cosmos-sdk/server"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/sharering/shareledger/app"
	"github.com/sharering/shareledger/cmd/Shareledgerd/cli"
	"github.com/sharering/shareledger/cmd/Shareledgerd/cmd"
	"github.com/sharering/shareledger/cmd/Shareledgerd/tools"
	"github.com/spf13/cobra"
)

var runOnce sync.Once

func init() {
	runOnce.Do(func() {
		cli.InitMiddleWare()
	})
}

func main() {
	rootCmd, _ := cmd.NewRootCmd()

	if err := svrcmd.Execute(rootCmd, "", app.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)
		default:
			os.Exit(1)
		}
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
