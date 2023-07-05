package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"

	"github.com/sharering/shareledger/x/distributionx/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group distributionx queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryParams())
	cmd.AddCommand(CmdListReward())
	cmd.AddCommand(CmdShowReward())
	cmd.AddCommand(CmdListBuilderCount())
	cmd.AddCommand(CmdShowBuilderCount())
	cmd.AddCommand(CmdListBuilderList())
	cmd.AddCommand(CmdShowBuilderList())
	// this line is used by starport scaffolding # 1

	return cmd
}
