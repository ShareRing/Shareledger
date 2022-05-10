package services

import (
	"context"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/sharering/shareledger/pkg/swap/abi/swap"
	swapmoduletypes "github.com/sharering/shareledger/x/swap/types"
	denom "github.com/sharering/shareledger/x/utils/demo"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"math/big"
	"os"
	"path/filepath"
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

//const flagType = "type" // in/out
//const flagSignerKeyName = "network-signers"
const flagConfigPath = "config"

var supportedTypes = map[string]struct{}{
	"in":  {},
	"out": {},
}

type Network struct {
	Signer   string `yaml:"signer"`
	Url      string `yaml:"url"`
	ChainId  int64  `yaml:"chainId"`
	Contract string `yaml:"contract"`
}

type RelayerConfig struct {
	Network      map[string]Network `yaml:"networks"`
	Type         string             `yaml:"type"`
	ScanInterval time.Duration      `yaml:"scanInterval"`
}

func parseConfig(filePath string) (RelayerConfig, error) {
	var cfg RelayerConfig
	filePath, err := filepath.Abs(filePath)
	if err != nil {
		return cfg, err
	}
	f, err := os.ReadFile(filePath)
	if err != nil {
		return cfg, err
	}
	err = yaml.Unmarshal(f, &cfg)
	return cfg, err
}

func NewStartCommands(defaultNodeHome string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "start Relayer process",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientTx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			configPath, _ := cmd.Flags().GetString(flagConfigPath)
			cfg, err := parseConfig(configPath)
			if err != nil {
				return err
			}
			relayerClient := initRelayer(clientTx, cfg)

			ctx, cancel := context.WithCancel(context.Background())
			go func() {
				server.WaitForQuitSignals()
				cancel()
			}()

			switch cfg.Type {
			case "in":
				return relayerClient.startInProcess(ctx)
			case "out":
				return relayerClient.startOutProcess(ctx)
			default:
				return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "Relayer type is required either in or out")
			}

			//hash, err := relayerClient.submitBatch(context.Background(), swapmoduletypes.Batch{})
			//fmt.Println(hash, err)
			//return err
			//return err
			//
			//signerStr, _ := cmd.Flags().GetString(flagSignerKeyName)
			//networkSignerPairs := strings.Split(signerStr, ",")
			//lenSigners := len(networkSignerPairs)
			//if lenSigners == 0 || lenSigners%2 != 0 {
			//	return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, fmt.Sprintf("%v flag is required and should be in pairs format <network-name>,<signer-key>..."))
			//}
			//mapNetworkSigners := make(map[string]string)
			//for i := 0; i < lenSigners-1; i += 2 {
			//	networkName := networkSignerPairs[i]
			//	keyName := networkSignerPairs[i+1]
			//	kb := clientTx.Keyring
			//	ks := keyring.NewKeyRingETH(kb)
			//	if _, err := ks.Key(keyName); err != nil {
			//		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "%v key name has error %+v", keyName, err)
			//	}
			//	mapNetworkSigners[networkName] = keyName
			//}
			//swapType, _ := cmd.Flags().GetString(flagType)
			//

			//relayerClient := initRelayer(clientTx, mapNetworkSigners, "https://eth-ropsten.alchemyapi.io/v2/0M8yP6-iyIof8dFJN0Jph59jJlSKqmbW")
			//time.Now().UTC().Unix()
			//switch swapType {
			//case "in":
			//	return relayerClient.startInProcess(ctx)
			//case "out":
			//	return relayerClient.startOutProcess(ctx)
			//default:
			//	return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "Relayer type is required either in or out")
			//}
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

	cmd.Flags().String(flagConfigPath, "./cmd/Shareledgerd/services/config.yml", "config path for Relayer")
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func initRelayer(client client.Context, cfg RelayerConfig) *Relayer {
	return &Relayer{
		Config: cfg,
		Client: client,
	}
}

type Relayer struct {
	Config RelayerConfig
	Client client.Context
}

func (r *Relayer) startOutProcess(ctx context.Context) error {
	// TODO: Get pending batches
	//var batches []swapmoduletypes.Batch
	// batches := getPendingBatches()
	// processingBatch(batches[0].id)
	// TODO: Change batch to pending

	// TODO: fw the current batch to sc -- pending
	// submitBatch(...)
	return nil
}

func (r *Relayer) getPendingBatches() []swapmoduletypes.Batch {
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

func (r *Relayer) updateBatch(msg *swapmoduletypes.MsgUpdateBatch) (swapmoduletypes.Batch, error) {
	mClient := swapmoduletypes.NewMsgClient(r.Client)
	qClient := swapmoduletypes.NewQueryClient(r.Client)

	_, err := mClient.UpdateBatch(context.Background(), msg)
	if err != nil {
		return swapmoduletypes.Batch{}, errors.Wrapf(err, "update batch id %d to processing fail", msg.GetBatchId())
	}
	batchIdReq := &swapmoduletypes.QueryBatchesRequest{
		Ids: []uint64{msg.GetBatchId()},
	}

	batchesRes, err := qClient.Batches(context.Background(), batchIdReq)

	if err != nil {
		return swapmoduletypes.Batch{}, errors.Wrapf(err, "geting batch id %d fail", msg.GetBatchId())
	}
	if len(batchesRes.GetBatches()) != 0 {
		return swapmoduletypes.Batch{}, fmt.Errorf("batches response is empty")
	}
	return batchesRes.GetBatches()[0], nil
}

func (r *Relayer) processingBatch(batchId uint64) (swapmoduletypes.Batch, error) {

	updateMsg := &swapmoduletypes.MsgUpdateBatch{
		Creator: r.Client.GetFromAddress().String(),
		BatchId: batchId,
		Status:  swapmoduletypes.BatchStatusProcessing,
	}

	batch, err := r.updateBatch(updateMsg)
	if err != nil {
		return swapmoduletypes.Batch{}, nil
	}

	return batch, nil
}

func (r *Relayer) submitBatch(ctx context.Context, batch swapmoduletypes.Batch) (txHash string, err error) {
	batch.Network = "erc20"
	networkConfig, found := r.Config.Network[batch.Network]
	uid := networkConfig.Signer
	if !found {
		return "", sdkerrors.Wrapf(sdkerrors.ErrLogic, "network, %s, is not supported", batch.Network)
	}
	conn, err := ethclient.Dial(networkConfig.Url)
	if err != nil {
		return "", sdkerrors.Wrapf(sdkerrors.ErrLogic, err.Error())
	}
	swapClient, err := swap.NewSwap(common.HexToAddress(networkConfig.Contract), conn)
	if err != nil {
		return "", sdkerrors.Wrapf(sdkerrors.ErrLogic, err.Error())
	}

	info, err := r.Client.Keyring.Key(uid)
	if err != nil {
		return "", err
	}
	pubkey := keyring.PubKeyETH{
		PubKey: info.GetPubKey(),
	}
	commonAdd := common.BytesToAddress(pubkey.Address().Bytes())
	currentNonce, err := conn.PendingNonceAt(ctx, commonAdd)
	if err != nil {
		return "", err
	}
	opts, err := keyring.NewKeyedTransactorWithChainID(r.Client.Keyring, uid, big.NewInt(networkConfig.ChainId))
	if err != nil {
		return "", err
	}
	opts.Nonce = big.NewInt(int64(currentNonce))
	tx, err := swapClient.Swap(opts, []*big.Int{}, []common.Address{}, []*big.Int{}, []byte{})
	if err != nil {
		return "", err
	}
	return tx.Hash().Hex(), err
}

func (r *Relayer) startInProcess(ctx context.Context) error {
	//TODO concurrency handle here
	_, err := r.listenETHSwap(ctx)
	if err != nil {
		return err
	}
	//Mocking the input
	return r.initSwapInRequest(ctx, "0xxa", "shareledger1w4l5fchs69d9avlgvdehq9ypvdh4xyev3p490g", "erc20", 100, 15)
}

func (r *Relayer) listenETHSwap(ctx context.Context) (interface{}, error) {
	return nil, nil
}

func (r *Relayer) initSwapInRequest(
	ctx context.Context,
	destAddr, slp3Addr, srcNet string,
	amount, fee uint64) error {
	swapAmount := sdk.NewDecCoin(denom.Shr, sdk.NewIntFromUint64(amount))
	swapFee := sdk.NewDecCoin(denom.Shr, sdk.NewIntFromUint64(fee))

	msgClient := swapmoduletypes.NewMsgClient(r.Client)
	inMsg := swapmoduletypes.NewMsgRequestIn(
		r.Client.GetFromAddress().String(),
		slp3Addr,
		destAddr,
		srcNet,
		swapAmount,
		swapFee,
	)
	_, err := msgClient.RequestIn(ctx, inMsg)
	if err != nil {
		return err
	}
	return nil
}
