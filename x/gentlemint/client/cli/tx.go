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

	cmd.AddCommand(CmdLoadShr())
	cmd.AddCommand(CmdLoadShrp())
	cmd.AddCommand(CmdBuyShr())
	cmd.AddCommand(CmdSendShr())
	cmd.AddCommand(CmdBuyCent())
	cmd.AddCommand(CmdBurnShrp())
	cmd.AddCommand(CmdSendShrp())
	cmd.AddCommand(CmdBurnShr())
	cmd.AddCommand(CmdSetExchange())
	// this line is used by starport scaffolding # 1

	return cmd
}
