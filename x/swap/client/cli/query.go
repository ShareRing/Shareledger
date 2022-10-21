package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sharering/shareledger/x/swap/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group swap queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryParams())
	cmd.AddCommand(CmdRequests())

	cmd.AddCommand(CmdBalance())

	cmd.AddCommand(CmdListSchema())
	cmd.AddCommand(CmdShowSchema())
	cmd.AddCommand(CmdPastTxEvents())

	cmd.AddCommand(CmdNextRequestId())

	cmd.AddCommand(CmdNextBatchId())

	// this line is used by starport scaffolding # 1
	cmd.AddCommand(CmdBatches())

	// this line is used by starport scaffolding # 1

	return cmd
}
