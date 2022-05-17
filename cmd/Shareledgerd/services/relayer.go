package services

import (
	"context"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"math/big"
	"os"
	"os/signal"
	"path/filepath"
	"sort"
	"sync"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"gopkg.in/yaml.v3"

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
	"github.com/sharering/shareledger/cmd/Shareledgerd/services/database"
	event "github.com/sharering/shareledger/cmd/Shareledgerd/services/subscriber"
	swaputil "github.com/sharering/shareledger/pkg/swap"
	"github.com/sharering/shareledger/pkg/swap/abi/swap"
	swapmoduletypes "github.com/sharering/shareledger/x/swap/types"
	denom "github.com/sharering/shareledger/x/utils/demo"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/mongo"
)

var log *zap.SugaredLogger

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
			logger, _ := zap.NewProduction()
			defer logger.Sync()
			log = logger.Sugar()

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

			ctx, cancel := context.WithCancel(context.Background())
			timeoutContext, cancelTimeOut := context.WithTimeout(ctx, time.Second*10)
			defer cancelTimeOut()
			relayerClient, err := initRelayer(clientTx, cfg, mgClient)
			if err != nil {
				return err
			}
			if err := mgClient.ConnectDB(timeoutContext); err != nil {
				cancel()
				return err
			}

			defer func() {
				if err := mgClient.Disconnect(ctx); err != nil {
					log.Info("Disconnected from DB")
				}
			}()

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
							log.Errorw(fmt.Sprintf("process with network %s", process.Network), "error", process.Err)
						}
						if numberProcessing == 0 {
							log.Info("all process were quited. Exiting")
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
						}{
							Network: network,
							Err:     relayerClient.startProcess(ctx, relayerClient.processIn, network),
						}
					}(network)
				}

			case "out":
				for network := range cfg.Network {
					go func(network string) {
						result := struct {
							Network string
							Err     error
						}{
							Network: network,
							Err:     relayerClient.startProcess(ctx, relayerClient.processOut, network),
						}
						processChan <- result
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

func initRelayer(client client.Context, cfg RelayerConfig, db database.DBRelayer) (*Relayer, error) {
	mClient := swapmoduletypes.NewMsgClient(client)
	qClient := swapmoduletypes.NewQueryClient(client)

	events := make(map[string]event.Service)
	for network, cfg := range cfg.Network {
		e, err := event.New(&event.NewInput{
			ProviderURL:          cfg.Url,
			TransferCurrentBlock: big.NewInt(cfg.LastScannedTransferEventBlockNumber), // config.yaml pre-define before running process
			SwapCurrentBlock:     big.NewInt(cfg.LastScannedSwapEventBlockNumber),
			PegWalletAddress:     cfg.TokenContract,
			TransferTopic:        cfg.TransferTopic,
			SwapContractAddress:  cfg.SwapContract,
			SwapTopic:            cfg.SwapTopic,
			DBClient:             db,
		})
		if err != nil {
			return nil, err
		}
		events[network] = *e
	}

	return &Relayer{
		Config:    cfg,
		Client:    client,
		db:        db,
		events:    events,
		qClient:   qClient,
		msgClient: mClient,
	}, nil
}

func (r *Relayer) startProcess(ctx context.Context, f processFunc, network string) error {
	doneChan := make(chan error)
	initInterval := time.Millisecond
	ticker := time.NewTicker(initInterval)
	defer func() {
		ticker.Stop()
	}()
	go func() {
		for {
			select {
			case <-ticker.C:
				ticker.Stop()
				err := f(ctx, network)
				if err == nil {
					ticker.Reset(r.Config.ScanInterval)
				}
				if err != nil {
					doneChan <- err
					return
				}
			case <-ctx.Done():
				log.Info("context is done. out process is exiting")
				doneChan <- nil
				return
			}
		}
	}()
	err := <-doneChan
	return err
}

func (r *Relayer) setLog(id uint64, msg string) error {
	return r.db.SetLog(id, msg)
}

func (r *Relayer) trackSubmittedBatch(ctx context.Context, batch database.Batch, timeout time.Duration) (database.Status, error) {
	deadTime := time.Now().Add(timeout).Add(time.Millisecond)
	for {
		select {
		case <-time.Tick(timeout / 2):
			if time.Now().After(deadTime) {
				return database.Submitted, nil
			}
			log.Info("checking receipt, network, %s, txHash, %s", batch.Network, batch.TxHash)
			fmt.Println("checking receipt", batch.Network, batch.TxHash)
			receipt, err := r.checkTxHash(ctx, batch.Network, common.HexToHash(batch.TxHash))
			if err != nil && err.Error() != "not found" {
				if err.Error() == "not found" { //transaction is still on mem pool
					continue
				}
				if IsErrBatchProcessed(err) {
					return database.Done, r.db.UpdateBatchesOut([]uint64{batch.ShareledgerID}, database.Done)
				}
				if e := r.setLog(batch.ShareledgerID, err.Error()); e != nil {
					log.Errorw("set log error", "original error", err, "log error", e)
				}
				return database.Failed, r.db.UpdateBatchesOut([]uint64{batch.ShareledgerID}, database.Failed)
			}
			if receipt != nil {
				switch receipt.Status {
				case 1:
					return database.Done, r.db.UpdateBatchesOut([]uint64{batch.ShareledgerID}, database.Done)
				case 0:
					msgLog, _ := json.MarshalIndent(receipt, "", "  ")
					if e := r.setLog(batch.ShareledgerID, string(msgLog)); e != nil {
						log.Errorw("set log msgLog", "msgLog", msgLog, "log error", e, "raw log", receipt)
					}
					return database.Failed, r.db.UpdateBatchesOut([]uint64{batch.ShareledgerID}, database.Failed)
				default:
					return database.Submitted, nil
				}
			}
		case <-ctx.Done():
			return database.Submitted, nil
		}
	}
}

func (r *Relayer) syncEventSuccessfulBatches(ctx context.Context, network string) error {
	service, found := r.events[network]
	if !found {
		return fmt.Errorf("%s network does not have swap event subscrber", network)
	}
	err := service.HandlerSwapCompleteEvent(ctx, func(hashes []common.Hash) error {
		for _, hash := range hashes {
			batch, err := r.db.GetBatchByTxHash(hash.String())
			if err != nil {
				return errors.Wrapf(err, "get batch by tx hash, %v", hash.String())
			}
			batch.Status = database.Done
			nonce := batch.Nonce
			if err := r.db.UpdateBatchesOut([]uint64{batch.ShareledgerID}, database.Done); err != nil {
				return errors.Wrapf(err, "update batch out. shareledger id %v", batch.ShareledgerID)
			}
			if err := r.db.SetBatchesOutFailed(nonce); err != nil {
				return errors.Wrapf(err, "set batches out failed. nonce %v", nonce)
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *Relayer) syncNewBatchesOut(ctx context.Context, network string) error {
	lastScannedBatchId, err := r.db.GetLastScannedBatch(network)
	if err != nil {
		return err
	}
	maxBatchId := lastScannedBatchId
	newBatches := make([]database.Batch, 0)

	res, err := r.qClient.Batches(ctx, &swapmoduletypes.QueryBatchesRequest{
		Status:  swapmoduletypes.BatchStatusPending,
		Network: network,
	})
	if err != nil {
		return err
	}
	for _, b := range res.Batches {
		if b.Id > lastScannedBatchId {
			newBatches = append(newBatches, database.Batch{
				ShareledgerID: b.Id,
				Status:        database.Pending,
				Type:          database.Out,
				TxHash:        common.Hash{}.String(),
				Network:       b.Network,
			})
		}
		if maxBatchId < b.Id {
			maxBatchId = b.Id
		}
	}
	if len(newBatches) > 0 {
		if err := r.db.InsertBatches(newBatches); err != nil {
			return errors.Wrapf(err, "new batches %+v", newBatches)
		}
		if err := r.db.UpdateLatestScannedBatchId(maxBatchId, network); err != nil {
			return errors.Wrapf(err, "update latest scanned batch id %+v", maxBatchId)
		}
	}
	return nil
}

func (r *Relayer) syncFailedBatches(ctx context.Context, network string) error {
	failedBatches, err := r.db.SearchBatchByStatus(network, database.Failed)
	if err != nil {
		return err
	}
	if len(failedBatches) == 0 {
		return nil
	}

	ids := make([]uint64, 0, len(failedBatches))
	for _, f := range failedBatches {
		ids = append(ids, f.ShareledgerID)
	}
	_, err = r.msgClient.CancelBatches(ctx, &swapmoduletypes.MsgCancelBatches{
		Creator: r.Client.GetFromAddress().String(),
		Ids:     ids,
	})
	if err != nil {
		return err
	}
	for i := range failedBatches {
		failedBatches[i].Status = database.Cancelled
	}
	return r.db.SetBatches(failedBatches)
}

func (r *Relayer) processNextPendingBatchesOut(ctx context.Context, network string) error {
	var offset int64
	for {
		batch, err := r.db.GetNextPendingBatchOut(network, offset)
		offset++
		if err != nil {
			return err
		}
		if batch == nil {
			log.Infow("pending batches list is empty", "network", network)
			return nil
		}

		b, err := r.getBatch(batch.ShareledgerID)
		if err != nil || b == nil {
			return err
		}
		batchDetail, err := r.getBatchDetail(ctx, *b)
		if err != nil {
			return err
		}
		total := sdk.NewCoin(denom.Base, sdk.NewInt(0))
		for _, r := range batchDetail.Requests {
			total = total.Add(*r.Amount)
		}

		currentBalance, err := r.getBalance(ctx, network)
		if err != nil {
			return err
		}
		if currentBalance.IsLT(total) {
			log.Warnw("total balance of current contract is less than swap total", "network", network, "batch", batch.ShareledgerID, "contract_total", currentBalance.String(), "swap_total", total.String())
			continue
		}

		retry := r.Config.Network[network].Retry
		var currentPrice *big.Int
		numberRetry := 0
		for numberRetry < retry.MaxRetry {
			numberRetry++
			if currentPrice != nil {
				nextPrice := retry.RetryPercentage*float64(currentPrice.Int64())/100 + float64(currentPrice.Int64())
				currentPrice, _ = big.NewFloat(nextPrice).Int(nil)
			}
			tx, err := r.submitBatch(ctx, network, batchDetail, currentPrice)
			if err != nil {
				if IsErrBatchProcessed(err) {
					batch.Status = database.Done
				} else {
					batch.Status = database.Failed
				}
				r.setLog(batch.ShareledgerID, err.Error())
			}
			if tx != nil {
				batch.Nonce = tx.Nonce()
				batch.Status = database.Submitted
				batch.TxHash = tx.Hash().String()
				currentPrice = tx.GasPrice()
			}
			if err := r.db.SetBatch(*batch); err != nil {
				return err
			}
			if batch.Status == database.Done || batch.Status == database.Failed {
				break
			}
			if batch.Status == database.Submitted {
				status, err := r.trackSubmittedBatch(ctx, *batch, retry.IntervalRetry)
				if err != nil {
					return err
				}
				if status != database.Submitted {
					// status will be failed and done which do not need to keep tracking
					break
				}
				continue
			}
		}
	}
	return nil
}

func (r *Relayer) processOut(ctx context.Context, network string) error {
	if err := r.syncEventSuccessfulBatches(ctx, network); err != nil {
		return err
	}
	if err := r.syncFailedBatches(ctx, network); err != nil {
		return err
	}
	if err := r.syncNewBatchesOut(ctx, network); err != nil {
		return err
	}
	if err := r.processNextPendingBatchesOut(ctx, network); err != nil {
		return err
	}
	return nil
}

func (r *Relayer) getBalance(ctx context.Context, network string) (sdk.Coin, error) {
	conn, networkConfig, err := r.initConn(network)
	swapClient, err := swap.NewSwap(common.HexToAddress(networkConfig.SwapContract), conn)
	if err != nil {
		return sdk.Coin{}, err
	}
	value, err := swapClient.TokensAvailable(&bind.CallOpts{
		Pending: false,
		Context: ctx,
	})
	return denom.ExponentToBase(sdk.NewInt(value.Int64()), r.Config.Network[network].Exponent), err
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

func (r *Relayer) getBatch(batchId uint64) (*swapmoduletypes.Batch, error) {
	qClient := swapmoduletypes.NewQueryClient(r.Client)
	pendingQuery := &swapmoduletypes.QueryBatchesRequest{
		Ids: []uint64{batchId},
	}

	batchesRes, err := qClient.Batches(context.Background(), pendingQuery)
	batches := batchesRes.GetBatches()
	sort.Sort(BatchSortByIDAscending(batches))
	if err != nil {
		return nil, err
	}
	if len(batches) == 0 {
		return nil, nil
	}
	return &batches[0], err
}

func (r *Relayer) getBatchDetail(ctx context.Context, batch swapmoduletypes.Batch) (detail swaputil.BatchDetail, err error) {
	qClient := swapmoduletypes.NewQueryClient(r.Client)
	// only approved swap requests have batch
	batchesRes, err := qClient.Swap(ctx, &swapmoduletypes.QuerySwapRequest{Ids: batch.TxIds, Status: swapmoduletypes.SwapStatusApproved})
	if err != nil {
		return detail, sdkerrors.Wrapf(err, "get list swap")
	}
	schema, err := qClient.Schema(ctx, &swapmoduletypes.QueryGetSchemaRequest{Network: batch.Network})
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

func (r *Relayer) updateBatchStatus(batchId uint64, status string) (b swapmoduletypes.Batch, err error) {
	updateMsg := &swapmoduletypes.MsgUpdateBatch{
		Creator: r.Client.GetFromAddress().String(),
		BatchId: batchId,
		Status:  status,
	}
	return r.updateBatch(updateMsg)
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
	swapClient, err := swap.NewSwap(common.HexToAddress(networkConfig.SwapContract), conn)
	if err != nil {
		return false, sdkerrors.Wrapf(sdkerrors.ErrLogic, err.Error())
	}
	res, err := swapClient.Batch(nil, digest)
	if len(res.Signature) > 0 {
		done = true
	}
	return done, err

}

func (r *Relayer) submitBatch(ctx context.Context, network string, batchDetail swaputil.BatchDetail, price *big.Int) (tx *types.Transaction, err error) {
	if err := batchDetail.Validate(); err != nil {
		return tx, err
	}
	conn, networkConfig, err := r.initConn(network)
	if err != nil {
		return tx, err
	}
	defer func() {
		conn.Close()
	}()

	swapClient, err := swap.NewSwap(common.HexToAddress(networkConfig.SwapContract), conn)
	if err != nil {
		return tx, sdkerrors.Wrapf(sdkerrors.ErrLogic, err.Error())
	}

	info, err := r.Client.Keyring.Key(networkConfig.Signer)
	if err != nil {
		return tx, err
	}
	pubkey := keyring.PubKeyETH{
		PubKey: info.GetPubKey(),
	}
	commonAdd := common.BytesToAddress(pubkey.Address().Bytes())

	//it should override pending nonce
	currentNonce, err := conn.NonceAt(ctx, commonAdd, nil)
	if err != nil {
		return tx, err
	}
	opts, err := keyring.NewKeyedTransactorWithChainID(r.Client.Keyring, networkConfig.Signer, big.NewInt(networkConfig.ChainId))
	if err != nil {
		return tx, err
	}
	if price != nil {
		opts.GasPrice = price
	}
	opts.Nonce = big.NewInt(int64(currentNonce))
	sig, err := hexutil.Decode(batchDetail.Batch.Signature)

	if err != nil {
		return tx, err
	}
	params, err := batchDetail.GetContractParams()
	if err != nil {
		return tx, err
	}
	tx, err = swapClient.Swap(opts, params.TransactionIds, params.DestAddrs, params.Amounts, sig)
	return tx, err
}

func (r *Relayer) processIn(ctx context.Context, network string) error {
	eventService, found := r.events[network]
	if !found {
		return fmt.Errorf("%s does not have event subcriber", network)
	}
	err := eventService.HandlerTransferEvent(ctx, func(events []event.EventTransferOutput) error {
		// check if these event are handle or not in db
		for _, e := range events {
			batch, err := r.db.GetBatchByTxHash(e.TxHash)
			if err != nil {
				// batch existed, skip processing
				if batch != (database.Batch{}) && err != mongo.ErrNoDocuments {
					continue
				}
				log.Errorw("get batch by tx hash", "err", err, "txHash", e.TxHash)
			}

			// get slp3 address from db
			slp3, err := r.db.GetSLP3Address(e.ToAddress, network)
			if err != nil {
				// log error and skip process this request
				log.Errorw("get slp3 address", "err", err)
				continue
			}

			// call shareledger transaction
			err = r.initSwapInRequest(
				ctx,
				e.ToAddress,
				slp3,
				network,
				e.TxHash,
				e.BlockNumber,
				e.Amount.BigInt().Uint64(),
				15,
			)
			if err != nil {
				log.Errorw("init swap in request", "err", err, "network", network, "txHash", e.TxHash)
			}
		}
		return nil
	})
	return err
}

func (r *Relayer) SubmitSwapIn(ctx context.Context, swap swapmoduletypes.MsgRequestIn) error {
	return nil
}

func (r *Relayer) initSwapInRequest(
	ctx context.Context,
	srcAddr, destAddr, network, txHash string,
	blockNumber, amount, fee uint64) error {
	swapAmount := sdk.NewDecCoin(denom.Shr, sdk.NewIntFromUint64(amount))
	swapFee := sdk.NewDecCoin(denom.Shr, sdk.NewIntFromUint64(fee))
	inMsg := swapmoduletypes.NewMsgRequestIn(
		r.Client.GetFromAddress().String(),
		srcAddr,
		destAddr,
		network,
		swapAmount,
		swapFee,
	)
	response, err := r.msgClient.RequestIn(ctx, inMsg)
	if err != nil {
		return err
	}

	// approve in msg

	newBatch := database.Batch{
		ShareledgerID: response.Id,
		Status:        database.Done,
		Type:          database.In,
		Network:       network,
		TxHash:        txHash,
		BlockNumber:   blockNumber,
	}
	err = r.db.SetBatch(newBatch)
	if err != nil {
		return err
	}

	return nil
}
