package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sharering/shareledger/x/gentlemint/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group gentlemint queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Aliases:                    []string{types.ModuleNameAlias},
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdShowExchangeRate())

	cmd.AddCommand(CmdListLevelFee())
	cmd.AddCommand(CmdShowLevelFee())
	cmd.AddCommand(CmdListActionLevelFee())
	cmd.AddCommand(CmdShowActionLevelFee())
	cmd.AddCommand(CmdCheckFees())

	cmd.AddCommand(CmdBalances())

	// this line is used by starport scaffolding # 1

	return cmd
}
