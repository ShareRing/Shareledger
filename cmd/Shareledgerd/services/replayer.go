package services

import (
	"context"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/server"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	swapmoduletypes "github.com/sharering/shareledger/x/swap/types"
	"github.com/spf13/cobra"
	"strings"
	"time"
)

func GetRelayerCommands(defaultNodeHome string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "relayer",
		Short: "relayer application commands",
	}
	cmd.AddCommand(
		NewStartCommands(defaultNodeHome),
	)
	return cmd
}

const flagType = "type" // in/out
const flagSignerKeyName = "network-signer"

var supportedTypes = map[string]struct{}{
	"in":  {},
	"out": {},
}

func NewStartCommands(defaultNodeHome string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "start relayer process",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			signerStr, _ := cmd.Flags().GetString(flagSignerKeyName)
			networkSignerPairs := strings.Split(signerStr, ",")
			lenSigners := len(networkSignerPairs)
			if lenSigners == 0 || lenSigners%2 != 0 {
				return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, fmt.Sprintf("%v flag is required and should be in pairs format <network-name>,<signer-key>..."))
			}
			mapNetworkSigners := make(map[string]string)
			for i := 0; i < lenSigners-1; i += 2 {
				networkName := networkSignerPairs[i]
				keyName := networkSignerPairs[i+1]
				kb := ctx.Keyring
				ks := keyring.NewKeyRingEIP712(kb)
				if _, err := ks.Key(keyName); err != nil {
					return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "%v key name has error %+v", keyName, err)
				}
				mapNetworkSigners[networkName] = keyName
			}
			swapType, _ := cmd.Flags().GetString(flagType)

			cancelContext, cancel := context.WithCancel(context.Background())
			go func() {
				server.WaitForQuitSignals()
				cancel()
			}()

			time.Now().UTC().Unix()
			switch swapType {
			case "in":
				return startInProcess(cancelContext, ctx, mapNetworkSigners)
			case "out":
				return startOutProcess(cancelContext, ctx, mapNetworkSigners)
			default:
				return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "relayer type is required either in or out")
			}
			//serverCtx := server.NewDefaultContext()
			//
			//config := serverCtx.Config
			//homeDir, _ := cmd.Flags().GetString(cli.HomeFlag)
			//config = config.SetRoot(homeDir)
			//
			//clientCtx, err := client.GetClientQueryContext(cmd)
			//addr, err := getAddr(cmd, clientCtx.HomeDir, args)
			//if err != nil {
			//	return err
			//}
			//var electoralGenesis electoralmoduletypes.GenesisState
			//if err := unmarshalGenesisState(cmd, homeDir, electoralmoduletypes.ModuleName, &electoralGenesis); err != nil {
			//	return errors.Wrap(err, "unmarshal genesis state electoral module types")
			//}
			//electoralGenesis.Authority = &electoralmoduletypes.Authority{
			//	Address: addr.String(),
			//}
			//if err := exportGenesisFile(cmd, homeDir, electoralmoduletypes.ModuleName, &electoralGenesis); err != nil {
			//	return errors.Wrap(err, "export genesis file ")
			//}
			//return nil
		},
	}

	cmd.Flags().String(flagType, "", "swap type(in/out) that the relayer will process")
	cmd.Flags().String(flagSignerKeyName, "", "network name and respectively key name that will be used to sign for a network. \n Format <network0>,<keyname0,<network1>,<keyname1>...")
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

type relayer struct {
	Signers      map[string]string
	IntervalScan time.Duration
}

func startOutProcess(ctx context.Context, client client.Context, signers map[string]string) error {
	// TODO: Get pending batches
	//var batches []swapmoduletypes.Batch
	// batches := getPendingBatches()
	// processingBatch(batches[0].id)
	// submitBatch(...)
	// TODO: Change batch to pending
	// TODO: fw the current batch to sc

	return nil
}

func getPendingBatches() []swapmoduletypes.Batch {
	panic("implement me")
}

func processingBatch(batchId uint64) (swapmoduletypes.Batch, error) {
	panic("implement me")
}

func submitBatch(ctx context.Context, client client.Context, batches swapmoduletypes.Batch) (txHash string, err error) {
	panic("implement me")
}

func startInProcess(ctx context.Context, client client.Context, signers map[string]string) error {
	return nil
}
