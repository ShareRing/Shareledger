package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ShareRing/Shareledger/x/electoral/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group electoral queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdGetVoter())

	cmd.AddCommand(CmdShowAuthority())
	cmd.AddCommand(CmdShowTreasurer())
	cmd.AddCommand(CmdGetLoader())
	cmd.AddCommand(CmdIdSigner())
	cmd.AddCommand(CmdIdSigners())
	cmd.AddCommand(CmdAccountOperator())
	cmd.AddCommand(CmdAccountOperators())
	cmd.AddCommand(CmdDocumentIssuer())
	cmd.AddCommand(CmdDocumentIssuers())

	// this line is used by starport scaffolding # 1

	return cmd
}
