package services

import (
	"context"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/server"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/pkg/errors"
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
			clientTx, err := client.GetClientTxContext(cmd)
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
				kb := clientTx.Keyring
				ks := keyring.NewKeyRingEIP712(kb)
				if _, err := ks.Key(keyName); err != nil {
					return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "%v key name has error %+v", keyName, err)
				}
				mapNetworkSigners[networkName] = keyName
			}
			swapType, _ := cmd.Flags().GetString(flagType)

			ctx, cancel := context.WithCancel(context.Background())
			go func() {
				server.WaitForQuitSignals()
				cancel()
			}()
			relayerClient := initRelayer(clientTx, mapNetworkSigners, "https://eth-ropsten.alchemyapi.io/v2/0M8yP6-iyIof8dFJN0Jph59jJlSKqmbW")
			time.Now().UTC().Unix()
			switch swapType {
			case "in":
				return relayerClient.startInProcess(ctx)
			case "out":
				return relayerClient.startOutProcess(ctx)
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

func initRelayer(client client.Context, signers map[string]string, ethUrl string) *relayer {
	return &relayer{
		Signers:      signers,
		IntervalScan: time.Second * 10, //TODO: interval scan
		EthUrl:       ethUrl,
		Client:       client,
	}
}

type relayer struct {
	Signers      map[string]string
	IntervalScan time.Duration
	EthUrl       string
	Client       client.Context
}

func (r *relayer) startOutProcess(ctx context.Context) error {
	// TODO: Get pending batches
	//var batches []swapmoduletypes.Batch
	// batches := getPendingBatches()
	// processingBatch(batches[0].id)
	// submitBatch(...)
	// TODO: Change batch to pending
	// TODO: fw the current batch to sc

	return nil
}

func (r *relayer) getPendingBatches() []swapmoduletypes.Batch {
	qClient := swapmoduletypes.NewQueryClient(r.Client)
	pendingQuery := &swapmoduletypes.QuerySearchBatchesRequest{
		Status: swapmoduletypes.BatchStatusPending,
	}

	batchesRes, err := qClient.SearchBatches(context.Background(), pendingQuery)
	//TODO consider logging
	if err != nil {
		return nil
	}
	return batchesRes.GetBatchs()

}

func (r *relayer) processingBatch(batchId uint64) (swapmoduletypes.Batch, error) {
	mClient := swapmoduletypes.NewMsgClient(r.Client)
	qClient := swapmoduletypes.NewQueryClient(r.Client)
	updateMsg := &swapmoduletypes.MsgUpdateBatch{
		Creator: r.Client.GetFromAddress().String(),
		BatchId: batchId,
		Status:  swapmoduletypes.BatchStatusProcessing,
	}
	_, err := mClient.UpdateBatch(context.Background(), updateMsg)
	if err != nil {
		return swapmoduletypes.Batch{}, errors.Wrapf(err, "update batch id %d to processing fail", batchId)
	}

	batchIdReq := &swapmoduletypes.QueryBatchesRequest{
		Ids: []uint64{batchId},
	}

	batchesRes, err := qClient.Batches(context.Background(), batchIdReq)

	if err != nil {
		return swapmoduletypes.Batch{}, errors.Wrapf(err, "geting batch id %d fail", batchId)
	}
	if len(batchesRes.GetBatches()) != 0 {
		return swapmoduletypes.Batch{}, fmt.Errorf("batches reponse is empty")
	}
	return batchesRes.GetBatches()[0], nil
}

func (r *relayer) submitBatch(ctx context.Context, batches swapmoduletypes.Batch) (txHash string, err error) {
	//conn, err := ethclient.Dial(r.EthUrl)
	//if err != nil {
	//	return "", sdkerrors.Wrapf(sdkerrors.ErrLogic, err.Error())
	//}
	//swapClient, err := swap.NewSwap(common.HexToAddress("0xC5eAdD9b5ea60A991a65888ECC8F26FbDdc7Dbf4"), conn)
	//if err != nil {
	//	return "", sdkerrors.Wrapf(sdkerrors.ErrLogic, err.Error())
	//}
	//info, err := r.Client.Keyring.Key(r.Signers["erc20"])
	//if err != nil {
	//	return "", err
	//}
	//var priv types.PrivKey
	//
	//switch i := info.(type) {
	//case info:
	//	if i.PrivKeyArmor == "" {
	//		return "nil, nil", fmt.Errorf("private key not available")
	//	}
	//	priv, err = legacy.PrivKeyFromBytes([]byte(i.PrivKeyArmor))
	//	if err != nil {
	//		return nil, nil, err
	//	}
	//default:
	//	return nil, nil, fmt.Errorf("eip712 currently supports for local key")
	//}
	//secp256k1Priv, ok := priv.(*secp256k1.PrivKey)
	//
	//bind.NewKeyedTransactorWithChainID()
	//
	//swapClient.Swap(bind.TransactOpts{
	//	From:      common.Address{},
	//	Nonce:     nil,
	//	Signer:    nil,
	//	Value:     nil,
	//	GasPrice:  nil,
	//	GasFeeCap: nil,
	//	GasTipCap: nil,
	//	GasLimit:  0,
	//	Context:   nil,
	//	NoSend:    false,
	//})
	panic("implement me")
}

func (r *relayer) startInProcess(ctx context.Context) error {
	return nil
}
