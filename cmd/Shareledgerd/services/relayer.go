package services

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	slog "github.com/sharering/shareledger/cmd/Shareledgerd/services/log"
	"go.uber.org/zap"
	"math/big"
	"os"
	"os/signal"
	"path/filepath"
	"sort"
	"syscall"
	"time"

	"gopkg.in/yaml.v3"

	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/sharering/shareledger/cmd/Shareledgerd/services/database"
	event "github.com/sharering/shareledger/cmd/Shareledgerd/services/subscriber"
	swaputil "github.com/sharering/shareledger/pkg/swap"
	swapmoduletypes "github.com/sharering/shareledger/x/swap/types"
	denom "github.com/sharering/shareledger/x/utils/demo"
	"github.com/spf13/cobra"
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
var log *zap.SugaredLogger

func NewStartCommands(defaultNodeHome string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "start Relayer process",
		RunE: func(cmd *cobra.Command, args []string) error {
			slog.Init()
			defer slog.Log.Sync()
			log = slog.Log

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
						if process.Err != nil {
							log.Errorw(fmt.Sprintf("process with network %s", process.Network), "error", process.Err)
						}
						log.Infof("Application exiting")
						cancel()
						return
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
					if cfg.AutoApprove {
						go func(network string) {
							processChan <- struct {
								Network string
								Err     error
							}{
								Network: network,
								Err:     relayerClient.startProcess(ctx, relayerClient.processApprovingIn, network),
							}
						}(network)
					}
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
			case "approver-in":
				for network := range cfg.Network {
					go func(network string) {
						processChan <- struct {
							Network string
							Err     error
						}{
							Network: network,
							Err:     relayerClient.startProcess(ctx, relayerClient.processApprovingIn, network),
						}
					}(network)
				}
			default:
				return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "Relayer type is required either in or out")
			}

			from, _ := cmd.Flags().GetString(flags.FlagFrom)
			log.With(
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

	txClientCtx = txClientCtx.
		WithSkipConfirmation(true).
		WithBroadcastMode("block").
		WithKeyring(cKeyRing).
		WithOutputFormat("json")

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

	relayer := Relayer{
		Config:   cfg,
		clientTx: txClientCtx,
		db:       db,
		events:   events,
		cmd:      cmd,
		qClient:  qClient,
	}
	if err := relayer.Validate(); err != nil {
		return nil, err
	}
	return &relayer, nil
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
				log.Info("context is done. processes are exiting")
				doneChan <- nil
				return
			}
		}
	}()
	err := <-doneChan
	return err
}

func (r *Relayer) trackSubmittedBatch(ctx context.Context, batch database.BatchOut, timeout time.Duration) (database.BatchStatus, error) {
	tickerTimeout := time.NewTicker(timeout)
	scanPeriod := time.NewTicker(timeout / 5)
	defer func() {
		tickerTimeout.Stop()
		scanPeriod.Stop()
	}()
	for {
		select {
		case <-tickerTimeout.C:
			return database.BatchStatusSubmitted, nil
		case <-scanPeriod.C:
			for _, hash := range batch.TxHashes {
				receipt, err := r.checkTxHash(ctx, batch.Network, common.HexToHash(hash))
				if err != nil {
					if IsErrNotFound(err) { //transaction is still on mem pool
						continue
					}
					if IsErrBatchProcessed(err) {
						return database.BatchStatusDone, r.db.UpdateBatchesOut([]uint64{batch.ShareledgerID}, database.BatchStatusDone)
					}

					if e := r.db.SetLog(batch, err.Error()); e != nil {
						log.Errorw("set log error", "original error", err, "log error", e)
					}
					return database.BatchStatusFailed, r.db.UpdateBatchesOut([]uint64{batch.ShareledgerID}, database.BatchStatusFailed)
				}
				if receipt != nil {
					switch receipt.Status {
					case 1:
						return database.BatchStatusDone, r.db.UpdateBatchesOut([]uint64{batch.ShareledgerID}, database.BatchStatusDone)
					case 0:
						msgLog, _ := json.MarshalIndent(receipt, "", "  ")
						if e := r.db.SetLog(batch, string(msgLog)); e != nil {
							log.Errorw("set log msgLog", "msgLog", msgLog, "log error", e, "raw log", receipt)
						}
						return database.BatchStatusFailed, r.db.UpdateBatchesOut([]uint64{batch.ShareledgerID}, database.BatchStatusFailed)
					default:
						fmt.Println("default:")
						return database.BatchStatusSubmitted, nil
					}
				}
			}
		case <-ctx.Done():
			return database.BatchStatusSubmitted, nil
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
			batch, err := r.db.GetBatchOutByTxHash(network, hash.String())
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
			if batch.Status == database.BatchStatusDone {
				// already processed done in other process. Skip
				logData = append(logData, "msg", "already done")
				continue
			}
			batch.Status = database.BatchStatusDone
			if err := r.db.UpdateBatchesOut([]uint64{batch.ShareledgerID}, database.BatchStatusDone); err != nil {
				logData = append(logData, "err", err)
				return errors.Wrapf(err, "update batch out. shareledger id %v", batch.ShareledgerID)

			}
			if err := r.db.SetBatchesOutFailed(network, nonce); err != nil {
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
	newBatches := make([]database.BatchOut, 0)

	res, err := r.qClient.Batches(ctx, &swapmoduletypes.QueryBatchesRequest{
		Network: network,
	})
	if err != nil {
		return err
	}
	for _, b := range res.Batches {
		if b.Id > lastScannedBatchId {
			newBatches = append(newBatches, database.BatchOut{
				Batch: database.Batch{
					ShareledgerID: b.Id,
					Status:        database.BatchStatusPending,
					Type:          database.BatchTypeOut,
					TxHashes:      []string{},
					Network:       b.Network,
				},
			})
			logData = append(logData, "batch_id", b.Id)
		}
		if maxBatchId < b.Id {
			maxBatchId = b.Id
		}
	}
	if len(newBatches) > 0 {
		if err := r.db.InsertBatchesOut(newBatches); err != nil {
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
	batches, err := r.db.SearchUnSyncedBatchOutByStatus(network, database.BatchStatusDone)
	if err != nil {
		return errors.Wrapf(err, "search batches by status %s fail", database.BatchStatusDone)
	}
	doneID := make([]uint64, 0, len(batches))
	for i := range batches {
		if batches[i].ShareledgerID > 0 {
			setDone := &swapmoduletypes.MsgCompleteBatch{
				BatchId: batches[i].ShareledgerID,
			}
			// mark sync done to avoid it runs multiple times when there is any wrong with sync process
			doneID = append(doneID, setDone.BatchId)
			cBatches, err := r.qClient.Batches(context.Background(), &swapmoduletypes.QueryBatchesRequest{Ids: []uint64{setDone.BatchId}})
			if err != nil || cBatches == nil || len(cBatches.Batches) == 0 {
				log.Errorw("get batch detail",
					"err", errors.Wrapf(err, "qClient"),
					"is_not_found", cBatches == nil || len(cBatches.Batches) == 0,
					"batch", batches[i])
				continue
			}
			if err := r.txUpdateBatch(setDone); err != nil {
				return errors.Wrapf(err, "update batchID=%d to status done", batches[i].ShareledgerID)
			}
		}
	}
	if len(doneID) > 0 {
		err = r.db.MarkBatchToSynced(doneID)
		if err != nil {
			return errors.Wrapf(err, "fail to update batch out to synced")
		}
	}
	return err
}

func (r *Relayer) syncFailedBatches(ctx context.Context, network string) error {
	failedBatches, err := r.db.SearchUnSyncedBatchOutByStatus(network, database.BatchStatusFailed)
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
	var ifailed []database.IBatch
	for i := range failedBatches {
		failedBatches[i].Status = database.BatchStatusCancelled
		ifailed = append(ifailed, failedBatches[i])
	}

	err = r.db.SetBatches(ifailed)
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
			offset++
			log.Errorw("get submitted batch", "err", err, "batch", batch)
			continue
		}
		batchDetail, err := r.getBatchDetail(ctx, *b)
		if err != nil {
			return errors.Wrap(err, "get batch detail fail")
		}

		total := sdk.NewCoin(denom.Base, sdk.NewInt(0))
		for _, r := range batchDetail.Requests {
			total = total.Add(r.Amount)
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

func (r *Relayer) submitAndTrack(ctx context.Context, batch database.BatchOut, detail swaputil.BatchDetail) error {
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
				batch.Status = database.BatchStatusDone
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
				batch.Status = database.BatchStatusFailed
			}
			_ = r.db.SetLog(batch, err.Error())
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
			batch.Status = database.BatchStatusSubmitted
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
		if batch.Status == database.BatchStatusDone || batch.Status == database.BatchStatusFailed {
			return nil
		}
		if batch.Status == database.BatchStatusSubmitted {
			status, err := r.trackSubmittedBatch(ctx, batch, rConfig.IntervalRetry)
			logSubmit = append(logSubmit,
				"track_submit_status", status,
				"track_time", time.Now().Format("2006-01-02 15:04:05"),
			)
			if err != nil {
				return errors.Wrapf(err, "tracking summited batch at smartcontract fail")
			}
			if status != database.BatchStatusSubmitted {
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

func (r *Relayer) getBatch(batchId uint64) (*swapmoduletypes.Batch, error) {
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
