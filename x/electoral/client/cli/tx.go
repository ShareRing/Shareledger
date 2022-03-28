package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/sharering/shareledger/x/electoral/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdEnrollVoter())
	cmd.AddCommand(CmdRevokeVoter())
	cmd.AddCommand(CmdEnrollLoaders())
	cmd.AddCommand(CmdEnrollLoadersFromFile())
	cmd.AddCommand(CmdRevokeLoaders())
	cmd.AddCommand(CmdRevokeLoadersFromFile())
	cmd.AddCommand(CmdEnrollIdSigners())
	cmd.AddCommand(CmdEnrollIdSignerFromFile())
	cmd.AddCommand(CmdRevokeIdSigners())
	cmd.AddCommand(CmdEnrollDocIssuers())
	cmd.AddCommand(CmdRevokeDocIssuers())
	cmd.AddCommand(CmdEnrollAccountOperators())
	cmd.AddCommand(CmdRevokeAccountOperators())
	// this line is used by starport scaffolding # 1

	return cmd
}
