package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ShareRing/Shareledger/x/gentlemint/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group gentlemint queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdListAccState())
	cmd.AddCommand(CmdShowAccState())
	cmd.AddCommand(CmdShowAuthority())
	cmd.AddCommand(CmdShowTreasurer())
	cmd.AddCommand(CmdShowExchangeRate())
	cmd.AddCommand(CmdGetLoader())

	// this line is used by starport scaffolding # 1

	return cmd
}
