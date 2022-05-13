package services

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"syscall"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/sharering/shareledger/cmd/Shareledgerd/services/database"
	event "github.com/sharering/shareledger/cmd/Shareledgerd/services/subscriber"
	swaputil "github.com/sharering/shareledger/pkg/swap"
	"github.com/sharering/shareledger/pkg/swap/abi/swap"
	swapmoduletypes "github.com/sharering/shareledger/x/swap/types"
	denom "github.com/sharering/shareledger/x/utils/demo"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
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
			mgClient, err := database.NewMongo(cfg.MongoURI)
			if err != nil {
				return err
			}
			relayerClient := initRelayer(clientTx, cfg, mgClient)

			ctx, cancel := context.WithCancel(context.Background())

			numberProcessing := len(cfg.Network)
			processChan := make(chan struct {
				Network string
				Err     error
			})
			doneChan := make(chan struct{})
			sigs := make(chan os.Signal, 1)
			go func() {
				signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
			}()

			go func() {
				defer func() {
					doneChan <- struct{}{}
				}()
				for {
					select {
					case <-sigs:
						cancel()
						return
					case process := <-processChan:
						numberProcessing--
						if process.Err != nil {
							log.Err(process.Err).Stack().Msg(fmt.Sprintf("process with network %s", process.Network))
						}
						if numberProcessing == 0 {
							log.Info().Msg("all process were quited. Exiting")
							cancel()
							return
						}
					}
				}
			}()

			switch cfg.Type {
			case "in":
				for network := range cfg.Network {
					go func(network string) {
						processChan <- struct {
							Network string
							Err     error
						}{Network: network, Err: relayerClient.startProcess(ctx, relayerClient.processIn, network)}
					}(network)
				}

			case "out":
				for network := range cfg.Network {
					go func(network string) {
						processChan <- struct {
							Network string
							Err     error
						}{Network: network, Err: relayerClient.startProcess(ctx, relayerClient.processOut, network)}
					}(network)
				}
			default:
				return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "Relayer type is required either in or out")
			}
			<-doneChan
			return nil
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

func initRelayer(client client.Context, cfg RelayerConfig, db database.DBRelayer) *Relayer {
	mClient := swapmoduletypes.NewMsgClient(client)
	qClient := swapmoduletypes.NewQueryClient(client)

	return &Relayer{
		Config:   cfg,
		Client:   client,
		DBClient: db,

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

func (r *Relayer) checkAndMarkDone(ctx context.Context, batchId uint64, network string, digest common.Hash) (done bool, err error) {
	isSCDone, err := r.isBatchDoneOnSC(network, digest)
	if err != nil {
		return false, err
	}
	done = isSCDone
	if done {
		if _, err = r.markDone(batchId); err != nil {
			return false, err
		}
	}
	return done, nil
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
	/*digest, err := batchDetail.Digest()
	if err != nil {
		return err
	}*/
	//done, err := r.checkAndMarkDone(ctx, batchDetail.Batch.Id, network, digest)
	//if err != nil || done {
	//	return err // err nil when done is true
	//}

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

func (r *Relayer) getBalance(ctx context.Context, network string) (sdk.Coin, error) {
	conn, networkConfig, err := r.initConn(network)
	_, err = swap.NewSwap(common.HexToAddress(networkConfig.Contract), conn)
	//if err != nil {
	return sdk.Coin{}, err
	//}
	//swapClient.TokensAvailable()
}

func (r *Relayer) checkTxHash(ctx context.Context, network string, txHash common.Hash) (*types.Receipt, error) {
	conn, _, err := r.initConn(network)
	if err != nil {
		return nil, err
	}
	defer func() {
		conn.Close()
	}()
	return conn.TransactionReceipt(ctx, txHash)
}

func (r *Relayer) getNextPendingBatch(network string) (*swapmoduletypes.Batch, error) {
	qClient := swapmoduletypes.NewQueryClient(r.Client)
	pendingQuery := &swapmoduletypes.QuerySearchBatchesRequest{
		Status:  swapmoduletypes.BatchStatusPending,
		Network: network,
	}

	batchesRes, err := qClient.SearchBatches(context.Background(), pendingQuery)
	batches := batchesRes.GetBatchs()
	return &batches[0], err
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

var lock sync.Mutex

// updateBatch thread safe to avoid running in multiple go routine for multiple network
func (r *Relayer) updateBatch(msg *swapmoduletypes.MsgUpdateBatch) (swapmoduletypes.Batch, error) {
	lock.Lock()
	defer lock.Unlock()

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

func (r *Relayer) initConn(network string) (*ethclient.Client, Network, error) {
	networkConfig, found := r.Config.Network[network]
	if !found {
		return nil, networkConfig, sdkerrors.Wrapf(sdkerrors.ErrLogic, "network, %s, is not supported", network)
	}
	conn, err := ethclient.Dial(networkConfig.Url)
	return conn, networkConfig, err
}

func (r *Relayer) isBatchDoneOnSC(network string, digest common.Hash) (done bool, err error) {
	conn, networkConfig, err := r.initConn(network)
	if err != nil {
		return false, err
	}
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

func (r *Relayer) submitBatch(ctx context.Context, network string, batchDetail swaputil.BatchDetail) (txHash common.Hash, err error) {
	if err := batchDetail.Validate(); err != nil {
		return txHash, err
	}
	conn, networkConfig, err := r.initConn(network)
	if err != nil {
		return txHash, err
	}
	defer func() {
		conn.Close()
	}()

	swapClient, err := swap.NewSwap(common.HexToAddress(networkConfig.Contract), conn)
	if err != nil {
		return txHash, sdkerrors.Wrapf(sdkerrors.ErrLogic, err.Error())
	}

	info, err := r.Client.Keyring.Key(networkConfig.Signer)
	if err != nil {
		return txHash, err
	}
	pubkey := keyring.PubKeyETH{
		PubKey: info.GetPubKey(),
	}
	commonAdd := common.BytesToAddress(pubkey.Address().Bytes())
	currentNonce, err := conn.PendingNonceAt(ctx, commonAdd)
	if err != nil {
		return txHash, err
	}

	opts, err := keyring.NewKeyedTransactorWithChainID(r.Client.Keyring, networkConfig.Signer, big.NewInt(networkConfig.ChainId))
	if err != nil {
		return txHash, err
	}
	opts.Nonce = big.NewInt(int64(currentNonce))
	sig, err := hexutil.Decode(batchDetail.Batch.Signature)
	if err != nil {
		return txHash, err
	}
	params, err := batchDetail.GetContractParams()
	if err != nil {
		return txHash, err
	}
	tx, err := swapClient.Swap(opts, params.TransactionIds, params.DestAddrs, params.Amounts, sig)
	if err != nil {
		return txHash, err
	}
	txHash = tx.Hash()
	return txHash, err
}

func (r *Relayer) processIn(ctx context.Context, network string) error {
	s, err := event.New(&event.NewInput{
		ProviderURL:  r.Config.Network[network].Url,
		CurrentBlock: big.NewInt(0),
	})
	if err != nil {
		return err
	}

	//TODO concurrency handle here
	topic := "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef" // this should be in config.yaml
	events, err := s.GetEvents(ctx, &event.EventInput{
		ContractAddress: r.Config.Network[network].Contract,
		Topic:           topic,
	})

	fmt.Println(events) // unused error skip
	//Mocking the input
	r.initSwapInRequest(ctx, "0xxa", "shareledger1w4l5fchs69d9avlgvdehq9ypvdh4xyev3p490g", "erc20", 100, 15)
	return nil
}

func (r *Relayer) SubmitSwapIn(ctx context.Context, swap swapmoduletypes.MsgRequestIn) error {
	return nil
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
