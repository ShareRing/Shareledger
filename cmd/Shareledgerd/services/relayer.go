package services

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"os/signal"
	"path/filepath"
	"sort"
	"sync"
	"syscall"
	"time"

	"github.com/rs/zerolog/log"
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

			ctx, cancel := context.WithCancel(context.Background())
			timeoutContext, cancelTimeOut := context.WithTimeout(ctx, time.Second*10)
			defer cancelTimeOut()
			relayerClient := initRelayer(clientTx, cfg, mgClient)
			if err := mgClient.ConnectDB(timeoutContext); err != nil {
				cancel()
				return err
			}

			defer func() {
				if err := mgClient.Disconnect(ctx); err != nil {
					log.Info().Msg("Disconnected from DB")
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
							log.Error().Stack().Err(process.Err).Msg(fmt.Sprintf("process with network %s", process.Network))
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
		Config: cfg,
		Client: client,
		db:     db,

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

//func (r *Relayer) getRequestStatus(ctx context.Context, detail swaputil.BatchDetail, network string) (status database.Status, txHash string, err error) {
//	processedBatch, err := r.db.SearchBatchByType(detail.Batch.Id, "in")
//	if err != nil || processedBatch == nil {
//		return "pending", err
//	}
//
//	// the final state is in db. so there is no need to check if it's done or not when the status done or failed.
//	if processedBatch.Status == database.Done || processedBatch.Status == database.Failed {
//		return processedBatch.Status, nil
//	}
//
//	// pending in db -> need to check on sc
//
//	hash, err := detail.Digest()
//	if err != nil {
//		return "", err
//	}
//	processedBatch.
//	isSCDone, err := r.isBatchDoneOnSC(network, hash)
//	if err != nil {
//		return false, err
//	}
//	done = isSCDone
//	if done {
//		if _, err = r.markDone(batchId); err != nil {
//			return false, err
//		}
//	}
//	return done, nil
//}

// setBatchOutFail update on db the fail batch.
func (r *Relayer) setBatchOutFail(ctx context.Context, batch database.Batch) error {
	if batch.Nonce == 0 {
		return fmt.Errorf("failed batch nonce is required")
	}
	batch.Status = database.Failed
	return r.db.SetBatch(batch)
}

func (r *Relayer) syncBatchOutFailed(ctx context.Context, id uint64, latestNonce uint64) error {
	return nil
}

// syncBatchOutDone update on db and shareledger
func (r *Relayer) syncBatchOutDone(ctx context.Context, id uint64) error {
	if _, err := r.updateBatchStatus(id, swapmoduletypes.BatchStatusDone); err != nil {
		return err
	}
	return r.db.UpdateBatchesOut([]uint64{id}, database.Done)
}

func (r *Relayer) syncBatchStatus(ctx context.Context, batch database.Batch) error {
	for {
		log.Info().Msgf("checking receipt, network, %s, txHash, %s", batch.Network, batch.TxHash)
		fmt.Println("checking receipt", batch.Network, batch.TxHash)
		receipt, err := r.checkTxHash(ctx, batch.Network, common.HexToHash(batch.TxHash))
		if err != nil && err.Error() != "not found" {
			if err.Error() == "not found" { //transaction is still on mem pool
				continue
			}
			if IsErrBatchProcessed(err) {
				return r.syncBatchOutDone(ctx, batch.ShareledgerID)
			}
			if IsErrRequestProcessed(err) {
				//return r.set
			}
		}
		if receipt != nil {
			switch receipt.Status {
			case 1:
				//if _, err = r.updateBatchStatus(batch.Id, swapmoduletypes.BatchStatusDone); err != nil {
				//	return err
				//}
			case 0:
				//receipt.
			}
			break
		}
	}
	return nil
}

func (r *Relayer) processOut(ctx context.Context, network string) error {
	batch, err := r.getNextPendingBatch(network)
	//batch, err := r.db.GetNextPendingBatchOut(network)
	if err != nil {
		return err
	}
	if batch == nil {
		log.Info().Msg("pending batches list is empty")
		return nil
	}
	//
	if len(batch.TxHash) == 0 {
		//err := r.syncBatchStatus(ctx, batch)
	} else {
		//err := r.submitBatch(context, batch)
	}

	batchDetail, err := r.getBatchDetail(ctx, *batch)
	if err != nil {
		return errors.Wrap(err, "get batch detail")
	}
	var txHash common.Hash
	processedBatch, err := r.db.SearchBatchByType(batch.Id, database.Out)
	if err != nil {
		return err
	}
	if processedBatch != nil {
		switch processedBatch.Status {
		case database.Failed:
			_, err = r.updateBatchStatus(batch.Id, swapmoduletypes.BatchStatusFail)
			return err
		case database.Done:
			_, err = r.updateBatchStatus(batch.Id, swapmoduletypes.BatchStatusDone)
			return err
		default:
			if len(processedBatch.TxHash) > 0 {
				txHash = common.HexToHash(processedBatch.TxHash)
			}
		}
	}
	var zeroHash common.Hash
	// in case this batch was not processed
	if txHash == zeroHash {
		txHash, err = r.submitBatch(ctx, network, batchDetail)
		if err != nil {
			return err
		}
		if err = r.db.SetBatch(database.Batch{
			ShareledgerID: batchDetail.Batch.Id,
			Status:        database.Pending,
			Type:          database.In,
			TxHash:        txHash.Hex(),
			Network:       network,
			BlockNumber:   0,
		}); err != nil {
			return err
		}
	}

	//for {
	//	time.Sleep(time.Second * 5)
	//	fmt.Println("checking receipt", network, txHash)
	//	receipt, err := r.checkTxHash(ctx, network, txHash)
	//	if err != nil && err.Error() != "not found" {
	//		return err
	//	}
	//	if receipt != nil {
	//		switch receipt.Status {
	//		case 1:
	//			if _, err = r.updateBatchStatus(batch.Id, swapmoduletypes.BatchStatusDone); err != nil {
	//				return err
	//			}
	//		case 0: //FAIL TODO: -> check many case fail. processed -> update done, wrong -> failed
	//			fmt.Println("wrong")
	//		}
	//		break
	//	}
	//}
	fmt.Println(txHash, err)
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
	pendingQuery := &swapmoduletypes.QueryBatchesRequest{
		Status:  swapmoduletypes.BatchStatusPending,
		Network: network,
	}

	batchesRes, err := qClient.Batches(context.Background(), pendingQuery)
	batches := batchesRes.GetBatches()
	sort.Sort(BatchSortByIDAscending(batches))

	if len(batches) == 0 {
		return nil, fmt.Errorf("pending batchs is empty")
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
	swapClient, err := swap.NewSwap(common.HexToAddress(networkConfig.Contract), conn)
	if err != nil {
		return false, sdkerrors.Wrapf(sdkerrors.ErrLogic, err.Error())
	}
	res, err := swapClient.Batch(nil, digest)
	if len(res.Signature) > 0 {
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

	//it should override pending nonce
	currentNonce, err := conn.NonceAt(ctx, commonAdd, nil)
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
		ProviderURL:          r.Config.Network[network].Url,
		TransferCurrentBlock: big.NewInt(r.Config.Network[network].LastScannedTransferEventBlockNumber), // config.yaml pre-define before running process
		SwapCurrentBlock:     big.NewInt(r.Config.Network[network].LastScannedSwapEventBlockNumber),
		DBClient:             r.db,
	})
	if err != nil {
		return err
	}

	events, err := s.GetTransferEvent(ctx, &event.EventTransferInput{
		PegWalletAddress: r.Config.Network[network].PegWallet,
		TransferTopic:    r.Config.Network[network].TransferTopic,
	})

	// check if these event are handle or not in db
	for _, e := range events {
		batch, err := r.db.GetBatchByTxHash(e.TxHash)
		if err != nil {
			// batch existed, skip processing
			if batch != (database.Batch{}) && err != mongo.ErrNoDocuments {
				continue
			}

			log.Err(err)
		}

		// get slp3 address from db
		slp3, err := r.db.GetSLP3Address(e.ToAddress, network)
		if err != nil {
			// log error and skip process this request
			log.Err(err)
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
			log.Err(err)
		}
	}

	return nil
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

	msgClient := swapmoduletypes.NewMsgClient(r.Client)
	inMsg := swapmoduletypes.NewMsgRequestIn(
		r.Client.GetFromAddress().String(),
		srcAddr,
		destAddr,
		network,
		swapAmount,
		swapFee,
	)
	response, err := msgClient.RequestIn(ctx, inMsg)
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
