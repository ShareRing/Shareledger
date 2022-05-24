package services

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
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
			logger, _ := zap.NewDevelopment()
			defer logger.Sync()
			log = logger.Sugar()

			configPath, _ := cmd.Flags().GetString(flagConfigPath)
			cfg, err := parseConfig(configPath)
			if err != nil {
				return errors.Wrapf(err, "parse config")
			}
			mgClient, err := database.NewMongo(cfg.MongoURI, cfg.DbName)
			if err != nil {
				return errors.Wrapf(err, "new mongodb")
			}

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			relayerClient, err := initRelayer(ctx, cmd, cfg, mgClient)
			if err != nil {
				return errors.Wrapf(err, "init relayer")
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
							log.Infof("all process were quited. Exiting")
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

			from, _ := cmd.Flags().GetString(flags.FlagFrom)
			logger.With(
				zap.String("started_at", time.Now().String()),
				zap.String("running_cosmos_key", from),
				zap.String("db_uri", cfg.MongoURI),
			).
				Info("Shareledger RELAYER processor is stared successfully")
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

func initRelayer(ctx context.Context, cmd *cobra.Command, cfg RelayerConfig, db database.DBRelayer) (*Relayer, error) {
	qClientCtx, err := client.GetClientQueryContext(cmd)
	if err != nil {
		return nil, err
	}

	// init keyrings for relayer's auto sign
	txClientCtx, err := client.GetClientTxContext(cmd)

	cKeyRing := cachedKeyRing{
		Keyring: txClientCtx.Keyring,
	}
	uids := make([]string, 0, len(cfg.Network)+1)
	uids = append(uids, txClientCtx.GetFromName())
	for _, n := range cfg.Network {
		uids = append(uids, n.Signer)
	}
	_ = cKeyRing.InitCaches(uids)

	txClientCtx = txClientCtx.WithSkipConfirmation(true).WithBroadcastMode("block").WithKeyring(cKeyRing)

	timeoutContext, cancelTimeOut := context.WithTimeout(ctx, time.Second*10)
	defer cancelTimeOut()

	if err := db.ConnectDB(timeoutContext); err != nil {
		return nil, err
	}

	qClient := swapmoduletypes.NewQueryClient(qClientCtx)
	events := make(map[string]event.Service)
	for network, cfg := range cfg.Network {
		eventTransferCurrentBlock, err := db.GetLastScannedBlockNumber(network, cfg.TokenContract)
		if err != nil {
			return nil, err
		}
		eventSwapCurrentBlock, err := db.GetLastScannedBlockNumber(network, cfg.SwapContract)
		if err != nil {
			return nil, err
		}

		e, err := event.New(&event.NewInput{
			ProviderURL:          cfg.Url,
			TransferCurrentBlock: big.NewInt(int64(eventTransferCurrentBlock)),
			SwapCurrentBlock:     big.NewInt(int64(eventSwapCurrentBlock)),
			PegWalletAddress:     cfg.TokenContract,
			TransferTopic:        cfg.TransferTopic,
			SwapContractAddress:  cfg.SwapContract,
			SwapTopic:            cfg.SwapTopic,
			DBClient:             db,
			Network:              network,
		})
		if err != nil {
			return nil, err
		}
		events[network] = *e
	}

	return &Relayer{
		Config:   cfg,
		clientTx: txClientCtx,
		db:       db,
		events:   events,
		cmd:      cmd,
		qClient:  qClient,
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
	tickerTimeout := time.NewTicker(timeout)
	scanPeriod := time.NewTicker(timeout / 5)
	defer func() {
		tickerTimeout.Stop()
		scanPeriod.Stop()
	}()
	for {
		select {
		case <-tickerTimeout.C:
			return database.Submitted, nil
		case <-scanPeriod.C:
			for _, hash := range batch.TxHashes {
				receipt, err := r.checkTxHash(ctx, batch.Network, common.HexToHash(hash))
				if err != nil {
					if IsErrNotFound(err) { //transaction is still on mem pool
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
						fmt.Println("default:")
						return database.Submitted, nil
					}
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
		var logDatas [][]interface{}
		defer func() {
			for _, l := range logDatas {
				log.Infow("handle swap complete event", l...)
			}
		}()
		var logData []interface{}
		for _, hash := range hashes {
			if len(logData) > 0 {
				logDatas = append(logDatas, logData)
				logData = []interface{}{}
			}
			logData = append(logData, "txHash", hash.String())
			batch, err := r.db.GetBatchByTxHash(hash.String())
			if err != nil {
				return errors.Wrapf(err, "get batch by tx hash, %v", hash.String())
			}
			if batch == nil {
				logData = append(logData, "batch", nil)
				log.Infof("get batch by tx hash, %v, is empty", hash.String())
				continue
			}

			nonce := batch.Nonce
			logData = append(logData, "batch_id", batch.ShareledgerID)
			logData = append(logData, "nonce", batch.Nonce)
			logData = append(logData, "batch_current_status", batch.Status)
			if batch.Status == database.Done {
				// already processed done in other process. Skip
				logData = append(logData, "msg", "already done")
				continue
			}
			batch.Status = database.Done
			if err := r.db.UpdateBatchesOut([]uint64{batch.ShareledgerID}, database.Done); err != nil {
				logData = append(logData, "err", err)
				return errors.Wrapf(err, "update batch out. shareledger id %v", batch.ShareledgerID)

			}
			if err := r.db.SetBatchesOutFailed(nonce); err != nil {
				logData = append(logData, "err", err)
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
	var logData []interface{}
	defer func() {
		log.Infow("sync new batches out", logData...)
	}()
	logData = append(logData, "network", network)
	lastScannedBatchId, err := r.db.GetLastScannedBatch(network)
	logData = append(logData, "network", network, "lastScannedBatchId", lastScannedBatchId)
	if err != nil {
		logData = append(logData, "err", err)
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
				TxHashes:      []string{},
				Network:       b.Network,
			})
			logData = append(logData, "batch_id", b.Id)
		}
		if maxBatchId < b.Id {
			maxBatchId = b.Id
		}
	}
	if len(newBatches) > 0 {
		if err := r.db.InsertBatches(newBatches); err != nil {
			logData = append(logData, "err_insert_batches", err)
			return errors.Wrapf(err, "new batches %+v", newBatches)
		}
		if err := r.db.UpdateLatestScannedBatchId(maxBatchId, network); err != nil {
			logData = append(logData, "err_update_latest_scanned_batch", err)
			return errors.Wrapf(err, "update latest scanned batch id %+v", maxBatchId)
		}
	}
	return nil
}

func (r *Relayer) syncFinishedBatches(ctx context.Context, network string) error {
	if err := r.syncDoneBatches(ctx, network); err != nil {
		return err
	}
	if err := r.syncFailedBatches(ctx, network); err != nil {
		return err
	}
	return nil
}

func (r *Relayer) syncDoneBatches(ctx context.Context, network string) error {
	batches, err := r.db.SearchUnSyncedBatchByStatus(network, database.Done)
	if err != nil {
		return errors.Wrapf(err, "search batches by status %s fail", database.Done)
	}
	sID := make([]uint64, 0, len(batches))
	for i := range batches {
		updateMsg := &swapmoduletypes.MsgUpdateBatch{
			BatchId: batches[i].ShareledgerID,
			Status:  swapmoduletypes.BatchStatusDone,
			Network: network,
		}

		if err := r.txUpdateBatch(updateMsg); err != nil {
			return errors.Wrapf(err, "update batchID=%d to status done fail", batches[i].ShareledgerID)
		}
		sID = append(sID, batches[i].ShareledgerID)
	}
	if len(sID) > 0 {
		err = r.db.MarkBatchToSynced(sID)
		if err != nil {
			return errors.Wrapf(err, "fail to update batch out to synced")
		}
	}
	return err
}

func (r *Relayer) syncFailedBatches(ctx context.Context, network string) error {
	failedBatches, err := r.db.SearchUnSyncedBatchByStatus(network, database.Failed)
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
	err = r.txCancelBatches(ids)
	if err != nil {
		return errors.Wrapf(err, "can't cancel batches")
	}
	for i := range failedBatches {
		failedBatches[i].Status = database.Cancelled
	}
	err = r.db.SetBatches(failedBatches)
	if err != nil {
		return errors.Wrapf(err, "set batches fail")
	}
	if len(ids) > 0 {
		err = r.db.MarkBatchToSynced(ids)
		if err != nil {
			return errors.Wrapf(err, "fail to update batch out to synced")
		}
	}

	return nil
}

func printOutLog(msg string, logData []interface{}) {
	if len(logData) > 0 {
		log.Infow(msg, logData...)
	}
}

func (r *Relayer) processNextPendingBatchesOut(ctx context.Context, network string) error {
	var offset int64
	var logData []interface{}
	defer func() {
		printOutLog("process next pending batches out", logData)
	}()
	for {
		printOutLog("process next pending batches out", logData)
		logData = []interface{}{}
		batch, err := r.db.GetNextUnfinishedBatchOut(network, offset)
		if err != nil {
			return err
		}
		if batch == nil {
			log.Infow("pending batches list is empty", "network", network)
			return nil
		}
		logData = append(logData, "batch_id", batch.ShareledgerID, "network", batch.Network)
		b, err := r.getBatch(batch.ShareledgerID)
		if err != nil || b == nil {
			return errors.Wrapf(err, "can't get next batch by ID %d", batch.ShareledgerID)
		}
		batchDetail, err := r.getBatchDetail(ctx, *b)
		if err != nil {
			return errors.Wrap(err, "get batch detail fail")
		}

		total := sdk.NewCoin(denom.Base, sdk.NewInt(0))
		for _, r := range batchDetail.Requests {
			total = total.Add(*r.Amount)
		}

		logData = append(logData, "batch_total", total.String())
		currentBalance, err := r.getBalance(ctx, network)
		if err != nil {
			return errors.Wrap(err, "can't get current balance of swap module")
		}
		logData = append(logData, "swap_contract_balance", currentBalance.String())
		if currentBalance.IsLT(total) {
			log.Warnw("total balance of current contract is less than swap total", "network", network, "batch", batch.ShareledgerID, "contract_total", currentBalance.String(), "swap_total", total.String())
			// try to the next pending batch
			offset++
			continue
		}
		if err := r.submitAndTrack(ctx, *batch, batchDetail); err != nil {
			logData = append(logData, "submit_and_track_err", err)
		}

		retry := r.Config.Network[network].Retry
		logData = append(logData, "retry_config", retry)
	}
}

func (r *Relayer) submitAndTrack(ctx context.Context, batch database.Batch, detail swaputil.BatchDetail) error {
	var logSubmit []interface{}
	var currentTip *big.Int
	var currentPrice *big.Int
	rConfig := r.Config.Network[batch.Network].Retry
	defer func() {
		logSubmit = append(logSubmit, "batch_id", batch.ShareledgerID, "network", batch.Network, "batch_status", batch.Status)
		printOutLog("submit and track end", logSubmit)
	}()
	for {
		logSubmit = append(logSubmit, "batch_id", batch.ShareledgerID)
		currentTip = CalculatePercentage(currentTip, rConfig.RetryPercentage)
		currentPrice = CalculatePercentage(currentPrice, rConfig.RetryPercentage)

		txRes, signerAddr, err := r.submitBatch(ctx, batch.Network, detail, currentTip, currentPrice, false)

		if err != nil {
			logSubmit = append(logSubmit,
				"err_submit_batch", err.Error(),
			)
			if IsErrBatchProcessed(err) {
				batch.Status = database.Done
			} else if IsErrUnderPrice(err) || IsErrAlreadyKnown(err) {
				// retry with higher tip.
				if currentTip == nil && currentPrice == nil {
					legacy, err := r.isLegacy(ctx, batch.Network, detail)
					if err != nil {
						return errors.Wrapf(err, "check tx legacy")
					}
					if legacy {
						currentPrice, err = r.SuggestGasPrice(ctx, batch.Network)
					} else {
						currentTip, err = r.SuggestGasTip(ctx, batch.Network)
					}
					if err != nil {
						return err
					}
				}
				continue
			} else {
				batch.Status = database.Failed
			}
			_ = r.setLog(batch.ShareledgerID, err.Error())
		}
		if txRes != nil {
			logSubmit = append(logSubmit,
				"tx_nonce", txRes.Nonce(),
				"tx_hashes", txRes.Hash(),
				"tx_tip", txRes.GasTipCap(),
				"tx_base_price", txRes.GasFeeCap(),
			)
			batch.Signer = signerAddr
			batch.Nonce = txRes.Nonce()
			batch.Status = database.Submitted
			batch.TxHashes = append(batch.TxHashes, txRes.Hash().String())

			currentTip = txRes.GasTipCap()
			currentPrice = txRes.GasPrice()

			if currentPrice.Cmp(currentTip) == 0 {
				// this is a legacy transaction. We use gas price.
				currentTip = nil
			} else {
				currentPrice = nil
			}
		}
		logSubmit = append(logSubmit, "network", batch.Network)
		printOutLog("submitted batch", logSubmit)
		logSubmit = []interface{}{}
		if err := r.db.SetBatch(batch); err != nil {
			return errors.Wrapf(err, "insert batch into database %v", batch)
		}
		if batch.Status == database.Done || batch.Status == database.Failed {
			return nil
		}
		if batch.Status == database.Submitted {
			status, err := r.trackSubmittedBatch(ctx, batch, rConfig.IntervalRetry)
			logSubmit = append(logSubmit,
				"track_submit_status", status,
				"track_time", time.Now().Format("2006-01-02 15:04:05"),
			)
			if err != nil {
				return errors.Wrapf(err, "tracking summited batch at smartcontract fail")
			}
			if status != database.Submitted {
				// status will be failed and done which do not need to keep tracking
				// trying to re-submit with higher tip
				return nil
			}
		}
	}
}

func (r *Relayer) processOut(ctx context.Context, network string) error {
	if err := r.syncEventSuccessfulBatches(ctx, network); err != nil {
		return errors.Wrapf(err, "sync success event batches fail network %s", network)
	}
	if err := r.syncFinishedBatches(ctx, network); err != nil {
		return errors.Wrapf(err, "syncFinishedBatches network %s", network)
	}
	if err := r.syncNewBatchesOut(ctx, network); err != nil {
		return errors.Wrapf(err, "syncNewBatchesOut network=%s", network)
	}
	if err := r.processNextPendingBatchesOut(ctx, network); err != nil {
		return errors.Wrapf(err, "process pending batch swap out fail network %s", network)
	}
	return nil
}

func (r *Relayer) getBalance(ctx context.Context, network string) (sdk.Coin, error) {
	conn, networkConfig, err := r.initConn(network)
	swapClient, err := swap.NewSwap(common.HexToAddress(networkConfig.SwapContract), conn)
	if err != nil {
		return sdk.Coin{}, errors.Wrapf(err, "fail to int Swap smartcontract client")
	}
	value, err := swapClient.TokensAvailable(&bind.CallOpts{
		Pending: false,
		Context: ctx,
	})
	return denom.ExponentToBase(sdk.NewIntFromBigInt(value), r.Config.Network[network].Exponent), err
}

func (r *Relayer) checkTxHash(ctx context.Context, network string, txHash common.Hash) (*types.Receipt, error) {
	conn, _, err := r.initConn(network)
	if err != nil {
		return nil, errors.Wrapf(err, "fail to int ETH network conection")
	}
	defer func() {
		conn.Close()
	}()
	return conn.TransactionReceipt(ctx, txHash)
}

func (r *Relayer) getBatch(batchId uint64) (*swapmoduletypes.Batch, error) {
	//qClient := swapmoduletypes.NewQueryClient(r.Client)
	pendingQuery := &swapmoduletypes.QueryBatchesRequest{
		Ids: []uint64{batchId},
	}

	batchesRes, err := r.qClient.Batches(context.Background(), pendingQuery)
	if err != nil {
		return nil, errors.Wrapf(err, "getting batches from blockchain fail")
	}
	batches := batchesRes.GetBatches()
	if len(batches) == 0 {
		return nil, errors.New("batches is empty")
	}

	sort.Sort(BatchSortByIDAscending(batches))
	return &batches[0], nil
}

func (r *Relayer) getBatchDetail(ctx context.Context, batch swapmoduletypes.Batch) (detail swaputil.BatchDetail, err error) {
	//qClient := swapmoduletypes.NewQueryClient(r.Client)
	// only approved swap requests have batches
	batchesRes, err := r.qClient.Swap(ctx, &swapmoduletypes.QuerySwapRequest{Ids: batch.TxIds, Status: swapmoduletypes.SwapStatusApproved})
	if err != nil {
		return detail, errors.Wrapf(err, "get list swap fail")
	}
	schema, err := r.qClient.Schema(ctx, &swapmoduletypes.QueryGetSchemaRequest{Network: batch.Network})
	if err != nil {
		return detail, errors.Wrapf(err, "can't get schema")
	}
	return swaputil.NewBatchDetail(batch, batchesRes.Swaps, schema.Schema), nil
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

func (r *Relayer) SuggestGasTip(ctx context.Context, network string) (*big.Int, error) {
	conn, _, err := r.initConn(network)
	if err != nil {
		return nil, errors.Wrapf(err, "init eth connection")
	}
	defer func() {
		conn.Close()
	}()
	return conn.SuggestGasTipCap(ctx)
}
func (r *Relayer) SuggestGasPrice(ctx context.Context, network string) (*big.Int, error) {
	conn, _, err := r.initConn(network)
	if err != nil {
		return nil, errors.Wrapf(err, "init eth connection")
	}
	defer func() {
		conn.Close()
	}()
	return conn.SuggestGasPrice(ctx)
}

// isLegacy check tx is support dynamic or legacy
// legacy transaction will return tip from gas price property.
// dynamic transaction, eip-1559, will return tip from tip and gas price from GasFeeCap.
// since the internal tx does not hav public field to determine types of transactions, we need to cmp to work around this.
func (r *Relayer) isLegacy(ctx context.Context, network string, batchDetail swaputil.BatchDetail) (bool, error) {
	tx, _, err := r.submitBatch(ctx, network, batchDetail, nil, nil, true)
	if err != nil {
		return false, err
	}
	return tx.GasTipCap().Cmp(tx.GasPrice()) != 0, nil
}

func (r *Relayer) submitBatch(ctx context.Context, network string, batchDetail swaputil.BatchDetail, tip *big.Int, gasPrice *big.Int, noSend bool) (tx *types.Transaction, signerAddr string, err error) {
	if tip != nil && gasPrice != nil {
		return nil, "", errors.New("tip and gas price should not have value at same time")
	}
	if err := batchDetail.Validate(); err != nil {
		return tx, signerAddr, err
	}
	conn, networkConfig, err := r.initConn(network)
	if err != nil {
		return tx, signerAddr, errors.Wrapf(err, "init eth connection")
	}
	defer func() {
		conn.Close()
	}()

	swapClient, err := swap.NewSwap(common.HexToAddress(networkConfig.SwapContract), conn)
	if err != nil {
		return tx, signerAddr, sdkerrors.Wrapf(sdkerrors.ErrLogic, err.Error())
	}

	info, err := r.clientTx.Keyring.Key(networkConfig.Signer)
	if err != nil {
		return tx, signerAddr, errors.Wrapf(err, "get keyring instant fail signer=%s", networkConfig.Signer)
	}
	pubKey := keyring.PubKeyETH{
		PubKey: info.GetPubKey(),
	}
	signerAddr = pubKey.Address().String()
	commonAdd := common.BytesToAddress(pubKey.Address().Bytes())

	//it should override pending nonce
	currentNonce, err := conn.NonceAt(ctx, commonAdd, nil)
	if err != nil {
		return tx, signerAddr, errors.Wrapf(err, "can't overide pending nonce for address %s", commonAdd.String())
	}
	opts, err := keyring.NewKeyedTransactorWithChainID(r.clientTx.Keyring, networkConfig.Signer, big.NewInt(networkConfig.ChainId))
	if err != nil {
		return tx, signerAddr, errors.Wrapf(err, "get eth connection options fail")
	}
	opts.GasTipCap = tip
	opts.GasPrice = gasPrice
	opts.NoSend = noSend

	opts.Nonce = big.NewInt(int64(currentNonce))
	sig, err := hexutil.Decode(batchDetail.Batch.Signature)

	if err != nil {
		return tx, signerAddr, errors.Wrapf(err, "decoding singature fail")
	}
	params, err := batchDetail.GetContractParams()
	if err != nil {
		return tx, signerAddr, err
	}
	tx, err = swapClient.Swap(opts, params.TransactionIds, params.DestAddrs, params.Amounts, sig)
	return tx, signerAddr, errors.Wrapf(err, "swapping at smart contract fail")
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
				log.Errorw("get batch by tx hash", "err", err, "txHash", e.TxHash)
			}
			if batch == nil {
				log.Infof("get batch by tx has is empty")
				continue
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
	txLock.Lock()
	defer txLock.Unlock()
	swapAmount := sdk.NewDecCoin(denom.Shr, sdk.NewIntFromUint64(amount))
	swapFee := sdk.NewDecCoin(denom.Shr, sdk.NewIntFromUint64(fee))

	inMsg := swapmoduletypes.NewMsgRequestIn(
		r.clientTx.GetFromAddress().String(),
		srcAddr,
		destAddr,
		network,
		swapAmount,
		swapFee,
	)
	if err := inMsg.ValidateBasic(); err != nil {
		return err
	}
	err := tx.GenerateOrBroadcastTxCLI(r.clientTx, r.cmd.Flags(), inMsg)
	if err != nil {
		return errors.Wrapf(err, "can't request swap in msg:%s", inMsg.String())
	}

	swapsRes, err := r.qClient.Swap(ctx, &swapmoduletypes.QuerySwapRequest{
		Status:      swapmoduletypes.SwapStatusPending,
		SrcAddr:     srcAddr,
		DestNetwork: network,
	})
	if err != nil {
		return errors.Wrapf(err, "fail to query the swap in request pending")
	}
	var rInIds = make([]uint64, 0, len(swapsRes.GetSwaps()))
	for _, rq := range swapsRes.Swaps {
		rInIds = append(rInIds, rq.GetId())
	}

	approveMsg := swapmoduletypes.NewMsgApproveIn(r.clientTx.GetFromAddress().String(), rInIds)
	if err := approveMsg.ValidateBasic(); err != nil {
		return errors.Wrap(err, "message approve in is invalid")
	}
	err = tx.GenerateOrBroadcastTxCLI(r.clientTx, r.cmd.Flags(), approveMsg)
	if err != nil {
		return errors.Wrap(err, "approve swap fail")
	}
	return nil
}

//#region shareledger-bc
var txLock sync.Mutex

func (r *Relayer) txCancelBatches(ids []uint64) error {
	txLock.Lock()
	defer txLock.Unlock()

	clientCtx, err := client.GetClientTxContext(r.cmd)
	if err != nil {
		return err
	}
	clientCtx = clientCtx.WithSkipConfirmation(true).WithBroadcastMode(flags.BroadcastBlock)
	msg := &swapmoduletypes.MsgCancelBatches{
		Creator: r.clientTx.GetFromAddress().String(),
		Ids:     ids,
	}
	if err := msg.ValidateBasic(); err != nil {
		return err
	}
	return tx.GenerateOrBroadcastTxCLI(r.clientTx, r.cmd.Flags(), msg)
}

// txUpdateBatch thread safe to avoid running in multiple go routine for multiple network
func (r *Relayer) txUpdateBatch(msg *swapmoduletypes.MsgUpdateBatch) error {
	txLock.Lock()
	defer txLock.Unlock()

	msg.Creator = r.clientTx.GetFromAddress().String()
	if err := msg.ValidateBasic(); err != nil {
		return err
	}
	err := tx.GenerateOrBroadcastTxCLI(r.clientTx, r.cmd.Flags(), msg)

	batchRes, err := r.qClient.Batches(context.Background(), &swapmoduletypes.QueryBatchesRequest{Ids: []uint64{msg.GetBatchId()}})
	if err != nil || len(batchRes.GetBatches()) == 0 {
		return errors.Wrapf(err, "recheck the batch id %d fail", msg.GetBatchId())
	}

	if batchRes.GetBatches()[0].GetStatus() != msg.GetStatus() {
		return errors.New("update the batch status fail")
	}

	return nil
}

//#endregion
