package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/sharering/shareledger/x/swap/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
	listSeparator              = ","
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

	cmd.AddCommand(CmdIn())
	cmd.AddCommand(CmdOut())
	cmd.AddCommand(CmdCancel())
	cmd.AddCommand(CmdReject())
	cmd.AddCommand(CmdWithdraw())
	cmd.AddCommand(CmdDeposit())
	cmd.AddCommand(
		getApproveCmd(),
		getBatchCmd(),
		getSchemaCmd(),
	)

	// this line is used by starport scaffolding # 1

	return cmd
}

func getApproveCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "approve",
		Short: "Approve subcommand from swap module",
	}
	cmd.AddCommand(
		CmdApproveIn(),
		CmdApprove(),
	)
	return cmd
}

func getBatchCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "batch",
		Short: "Batch subcommand from swap module",
	}
	cmd.AddCommand(
		CmdCancelBatches(),
		CmdCompleteBatch(),
	)
	return cmd
}

func getSchemaCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "schema",
		Short: "Schema subcommand from swap module",
	}
	cmd.AddCommand(
		CmdCreateSchema(),
		CmdUpdateSchema(),
		CmdDeleteSchema(),
		CmdUpdateSwapFee(),
	)
	return cmd
}
