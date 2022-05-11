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
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	swaputil "github.com/sharering/shareledger/pkg/swap"
	"github.com/sharering/shareledger/pkg/swap/abi/swap"
	swapmoduletypes "github.com/sharering/shareledger/x/swap/types"
	denom "github.com/sharering/shareledger/x/utils/demo"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"math/big"
	"os"
	"path/filepath"
	"sort"
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
				_ = server.WaitForQuitSignals()
				cancel()
			}()

			switch cfg.Type {
			case "in":
				return relayerClient.startInProcess(ctx)
			case "out":
				return relayerClient.startProcess(ctx, relayerClient.processOut, "erc20")
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

func initRelayer(client client.Context, cfg RelayerConfig) *Relayer {

	mClient := swapmoduletypes.NewMsgClient(client)
	qClient := swapmoduletypes.NewQueryClient(client)
	return &Relayer{
		Config: cfg,
		Client: client,

		qClient:   qClient,
		msgClient: mClient,
	}
}

func (r *Relayer) startProcess(ctx context.Context, f processFunc, network string) error {
	doneChan := make(chan error)
	initInterval := time.Millisecond
	ticker := time.NewTicker(initInterval)
	defer func() {
		ticker.Stop()
	}()
	firstRun := true
	go func() {
		for {
			select {
			case <-ticker.C:
				if firstRun {
					ticker.Reset(r.Config.ScanInterval)
					firstRun = false
				}
				err := f(ctx, network)
				if err != nil {
					doneChan <- err
					return
				}
			case <-ctx.Done():
				log.Info().Msg("context is done. out process is exiting")
				doneChan <- nil
				return
			}
		}
	}()

	return <-doneChan
}

func (r *Relayer) processOut(ctx context.Context, network string) error {
	batch, err := r.getNextPendingBatch(network)
	if err != nil {
		log.Err(err).Msg("get pending batches")
	}
	if batch == nil {
		log.Info().Msg("pending batches list is empty")
		return nil
	}
	batchDetail, err := r.getBatchDetail(ctx, *batch)
	if err != nil {
		return err
	}
	digest, err := batchDetail.Digest()
	if err != nil {
		return err
	}
	if done, err := r.isBatchDoneOnSC(network, digest); err != nil || done {
		if err != nil {
			return err
		}
		if done {
			_, err = r.markDone(batchDetail.Batch.Id)
			if err != nil {
				return err
			}
			return nil
		}
	}
	// There is a case that there is already pending transaction.
	// But the sc wil double check and return fee if the request batch already processed.
	txHash, err := r.submitBatch(ctx, network, batchDetail)
	fmt.Println(txHash, err)
	// TODO: Hoai job check requently
	if err != nil {
		//TODO: handle error??
	}
	return err
}

func (r *Relayer) getNextPendingBatch(network string) (*swapmoduletypes.Batch, error) {
	qClient := swapmoduletypes.NewQueryClient(r.Client)
	pendingQuery := &swapmoduletypes.QuerySearchBatchesRequest{
		Status:  swapmoduletypes.BatchStatusPending,
		Network: network,
	}

	batchesRes, err := qClient.SearchBatches(context.Background(), pendingQuery)
	batches := batchesRes.GetBatchs()
	sort.Sort(BatchSortByIDAscending(batches))

	if len(batches) == 0 {
		return nil, fmt.Errorf("pending batchs is empty")
	}

	idQ := &swapmoduletypes.QueryBatchesRequest{Ids: []uint64{batches[0].Id}}
	batch, err := r.qClient.Batches(context.Background(), idQ)
	if err != nil || len(batch.GetBatches()) == 0 {
		return nil, fmt.Errorf("batch not found")
	}

	return &batch.GetBatches()[0], err
}

func (r *Relayer) getBatchDetail(ctx context.Context, batch swapmoduletypes.Batch) (detail swaputil.BatchDetail, err error) {
	qClient := swapmoduletypes.NewQueryClient(r.Client)
	// only approved swap requests have batch
	batchesRes, err := qClient.Swap(ctx, &swapmoduletypes.QuerySwapRequest{Ids: batch.TxIds, Status: swapmoduletypes.SwapStatusApproved})
	if err != nil {
		return detail, sdkerrors.Wrapf(err, "get list swap")
	}
	schema, err := qClient.SignSchema(ctx, &swapmoduletypes.QueryGetSignSchemaRequest{Network: batch.Network})
	if err != nil {
		return detail, err
	}
	return swaputil.NewBatchDetail(batch, batchesRes.Swaps, schema.Schema), nil
}

func (r *Relayer) updateBatch(msg *swapmoduletypes.MsgUpdateBatch) (swapmoduletypes.Batch, error) {
	_, err := r.msgClient.UpdateBatch(context.Background(), msg)
	if err != nil {
		return swapmoduletypes.Batch{}, errors.Wrapf(err, "update batch id %d to processing fail", msg.GetBatchId())
	}
	batchIdReq := &swapmoduletypes.QueryBatchesRequest{
		Ids: []uint64{msg.GetBatchId()},
	}

	batchesRes, err := r.qClient.Batches(context.Background(), batchIdReq)

	if err != nil {
		return swapmoduletypes.Batch{}, errors.Wrapf(err, "geting batch id %d fail", msg.GetBatchId())
	}
	if len(batchesRes.GetBatches()) != 0 {
		return swapmoduletypes.Batch{}, fmt.Errorf("batches response is empty")
	}
	return batchesRes.GetBatches()[0], nil
}

func (r *Relayer) markDone(batchId uint64) (b swapmoduletypes.Batch, err error) {

	updateMsg := &swapmoduletypes.MsgUpdateBatch{
		Creator: r.Client.GetFromAddress().String(),
		BatchId: batchId,
		Status:  swapmoduletypes.BatchStatusDone,
	}
	return r.updateBatch(updateMsg)
}

func (r *Relayer) markBatchProcessing(batchId uint64) (swapmoduletypes.Batch, error) {
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

func (r *Relayer) isBatchDoneOnSC(network string, digest common.Hash) (done bool, err error) {
	networkConfig, found := r.Config.Network[network]
	if !found {
		return false, sdkerrors.Wrapf(sdkerrors.ErrLogic, "network, %s, is not supported", "erc20")
	}
	conn, err := ethclient.Dial(networkConfig.Url)
	defer func() {
		conn.Close()
	}()
	swapClient, err := swap.NewSwap(common.HexToAddress(networkConfig.Contract), conn)
	if err != nil {
		return false, sdkerrors.Wrapf(sdkerrors.ErrLogic, err.Error())
	}
	res, err := swapClient.Swaps(nil, digest)
	if len(res) > 0 {
		done = true
	}
	return done, err

}

func (r *Relayer) submitBatch(ctx context.Context, network string, batchDetail swaputil.BatchDetail) (txHash string, err error) {
	if err := batchDetail.Validate(); err != nil {
		return "", err
	}
	networkConfig, found := r.Config.Network[network]
	uid := networkConfig.Signer
	if !found {
		return "", sdkerrors.Wrapf(sdkerrors.ErrLogic, "network, %s, is not supported", network)
	}
	conn, err := ethclient.Dial(networkConfig.Url)
	defer func() {
		conn.Close()
	}()
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

	transactionIds := make([]*big.Int, 0, len(batchDetail.Requests))
	destAddr := make([]common.Address, 0, len(batchDetail.Requests))
	amounts := make([]*big.Int, 0, len(batchDetail.Requests))

	for _, r := range batchDetail.Requests {
		coins, err := denom.NormalizeToBaseCoins(sdk.NewDecCoins(*r.Amount), false)
		if err != nil {
			return "", err
		}
		transactionIds = append(transactionIds, big.NewInt(int64(r.Id)))
		destAddr = append(destAddr, common.HexToAddress(r.DestAddr))
		amounts = append(amounts, big.NewInt(coins[0].Amount.Int64()))
	}
	sig, err := hexutil.Decode(batchDetail.Batch.Signature)
	if err != nil {
		return "", err
	}
	d, _ := batchDetail.Digest()
	fmt.Println("digest", d.Hex())
	fmt.Println("sig", batchDetail.Batch.Signature)
	tx, err := swapClient.Swap(opts, transactionIds, destAddr, amounts, sig)

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
