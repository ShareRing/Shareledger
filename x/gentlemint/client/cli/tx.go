package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/sharering/shareledger/x/gentlemint/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Aliases:                    []string{types.ModuleNameAlias},
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdBuyShr())
	cmd.AddCommand(CmdSetExchange())
	cmd.AddCommand(CmdSetLevelFee())
	cmd.AddCommand(CmdDeleteLevelFee())
	cmd.AddCommand(CmdSetActionLevelFee())
	cmd.AddCommand(CmdDeleteActionLevelFee())
	cmd.AddCommand(CmdLoad())
	cmd.AddCommand(CmdSend())
	cmd.AddCommand(CmdBurn())
	// this line is used by starport scaffolding # 1

	return cmd
}
