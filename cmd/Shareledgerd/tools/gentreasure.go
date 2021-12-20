package tools

import (
	electoralmoduletypes "github.com/ShareRing/Shareledger/x/electoral/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/libs/cli"
)

func NewGenesisAddValidatorAccountCmd(defaultNodeHome string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-treasure [address_or_key_name]",
		Short: "Add a treasure account to genesis.json",
		Long:  "Add a treasure account to genesis.json",
		RunE: func(cmd *cobra.Command, args []string) error {
			serverCtx := server.NewDefaultContext()

			config := serverCtx.Config
			homeDir, _ := cmd.Flags().GetString(cli.HomeFlag)
			config = config.SetRoot(homeDir)

			clientCtx, err := client.GetClientQueryContext(cmd)
			addr, err := getAddr(cmd, clientCtx.HomeDir, args)
			if err != nil {
				return err
			}
			var electoralGenesis electoralmoduletypes.GenesisState
			if err := unmarshalGenesisState(cmd, homeDir, electoralmoduletypes.ModuleName, &electoralGenesis); err != nil {
				return errors.Wrap(err, "unmarshal genesis state electoral module types")
			}
			electoralGenesis.Treasurer = &electoralmoduletypes.Treasurer{
				Address: addr.String(),
			}
			if err := exportGenesisFile(cmd, homeDir, electoralmoduletypes.ModuleName, &electoralGenesis); err != nil {
				return errors.Wrap(err, "export genesis file ")
			}
			return nil
		},
	}
	cmd.Flags().String(flags.FlagKeyringBackend, flags.DefaultKeyringBackend, "Select keyring's backend (os|file|kwallet|pass|test)")
	cmd.Flags().String(cli.HomeFlag, defaultNodeHome, "node's home directory")
	return cmd
}
