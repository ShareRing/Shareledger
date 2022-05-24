package subscriber

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"math/big"
	"strings"

	eth "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sharering/shareledger/cmd/Shareledgerd/services/database"
	"github.com/sharering/shareledger/pkg/swap/abi/erc20"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
)

const (
	transferEvent = "Transfer"
	swapEvent     = "SwapCompleted"
)

//type EventTransferInput struct {
//	PegWalletAddress string
//	TransferTopic    string
//}

type EventTransferOutput struct {
	FromAddress string
	ToAddress   string
	Amount      decimal.Decimal
	TxHash      string
	BlockNumber uint64
}

//type EventSwapCompleteInput struct {
//	SwapContractAddress string
//	SwapTopic           string
//}

type Service struct {
	client               *ethclient.Client
	transferCurrentBlock *big.Int
	swapCurrentBlock     *big.Int
	DBClient             database.DBRelayer

	pegWalletAddress    string
	transferTopic       string
	swapContractAddress string
	swapTopic           string
	network             string
}

type NewInput struct {
	Network              string
	ProviderURL          string
	TransferCurrentBlock *big.Int
	SwapCurrentBlock     *big.Int

	PegWalletAddress    string
	TransferTopic       string
	SwapContractAddress string
	SwapTopic           string

	DBClient database.DBRelayer
}

func init() {
	log.SetLevel(log.DebugLevel)
}

func New(input *NewInput) (*Service, error) {
	client, err := ethclient.Dial(input.ProviderURL)
	if err != nil {
		return nil, errors.Wrap(err, "dial to eth fail")
	}

	return &Service{
		client:               client,
		transferCurrentBlock: input.TransferCurrentBlock,
		swapCurrentBlock:     input.SwapCurrentBlock,
		DBClient:             input.DBClient,

		pegWalletAddress:    input.PegWalletAddress,
		transferTopic:       input.TransferTopic,
		swapContractAddress: input.SwapContractAddress,
		swapTopic:           input.SwapTopic,
		network:             input.Network,
	}, nil
}

type handlerSwapEvent func(events []common.Hash) error

func (s *Service) HandlerSwapCompleteEvent(ctx context.Context, fn handlerSwapEvent) (err error) {
	header, err := s.client.HeaderByNumber(ctx, nil)
	if err != nil {
		return errors.Wrapf(err, "get block head fail")
	}
	if header.Number.Cmp(s.swapCurrentBlock) < 0 {
		log.Info("there is no new block")
		return errors.Wrapf(err, "no block")
	}

	for header.Number.Cmp(s.swapCurrentBlock) > 0 {
		currentHeaderNumber := header.Number

		toBlock := big.NewInt(s.swapCurrentBlock.Int64() + 4500)
		if toBlock.Cmp(currentHeaderNumber) > 0 {
			toBlock = big.NewInt(currentHeaderNumber.Int64())
		}

		log.Debugf("Scanning from block %v to block %v, network, %s", s.swapCurrentBlock, toBlock, s.network)
		any := []common.Hash{}
		query := eth.FilterQuery{
			FromBlock: s.swapCurrentBlock,
			ToBlock:   toBlock,
			Addresses: []common.Address{common.HexToAddress(s.swapContractAddress)},
			Topics: [][]common.Hash{
				[]common.Hash{common.HexToHash(s.swapTopic)},
				any,
			},
		}

		logs, err := s.client.FilterLogs(ctx, query)
		if err != nil {
			return errors.Wrapf(err, "filter event log by %+v  fail", query)
		}

		events := make([]common.Hash, 0, len(logs))

		for _, vLog := range logs {

			events = append(events, vLog.TxHash)
		}
		if err := fn(events); err != nil {
			return errors.Wrapf(err, "handle event fail")
		}

		// save last scanned block number to db
		err = s.DBClient.SetLastScannedBlockNumber(s.network, s.swapContractAddress, toBlock.Int64())
		if err != nil {
			return errors.Wrapf(err, "set the last scanned block into db fail")
		}

		// set current block number = latest + 1 for next tick interval
		s.swapCurrentBlock.Add(toBlock, big.NewInt(1))
	}

	return nil
}

type handlerTransferEvent func(events []EventTransferOutput) error

func (s *Service) HandlerTransferEvent(ctx context.Context, fn handlerTransferEvent) (err error) {
	erc20Abi, err := abi.JSON(strings.NewReader(string(erc20.Erc20MetaData.ABI)))
	if err != nil {
		return errors.Wrapf(err, "unmarshal swap abi code fail")
	}

	// skip if header not found
	header, err := s.client.HeaderByNumber(ctx, nil)
	if err != nil {
		return errors.Wrapf(err, "get block head fail")
	}

	if header.Number.Cmp(s.transferCurrentBlock) < 0 {
		log.Info("there is no new block")
		return errors.Wrapf(err, "no block")
	}

	for header.Number.Cmp(s.transferCurrentBlock) > 0 {
		currentHeaderNumber := header.Number

		toBlock := big.NewInt(s.swapCurrentBlock.Int64() + 4500)
		if toBlock.Cmp(currentHeaderNumber) > 0 {
			toBlock = big.NewInt(currentHeaderNumber.Int64())
		}

		log.Debugf("Scanning from block %v to block %v", s.transferCurrentBlock, toBlock.Int64())
		any := []common.Hash{}
		query := eth.FilterQuery{
			FromBlock: s.transferCurrentBlock,
			ToBlock:   toBlock,
			Addresses: []common.Address{common.HexToAddress(s.pegWalletAddress)},
			Topics: [][]common.Hash{
				[]common.Hash{common.HexToHash(s.transferTopic)},
				any,
				any,
			},
		}

		logs, err := s.client.FilterLogs(ctx, query)
		if err != nil {
			return errors.Wrapf(err, "filter the logs response fail by query %v+", query)
		}

		events := make([]EventTransferOutput, 0, len(logs))

		for _, vLog := range logs {
			output := EventTransferOutput{}
			var event = struct {
				Value *big.Int // amount in erc20 contract
			}{}

			err := erc20Abi.UnpackIntoInterface(&event, transferEvent, vLog.Data)
			if err != nil {
				log.Errorf("Event unpacking error: %s", err)
				continue
			}

			output.Amount = decimal.NewFromBigInt(event.Value, 0)
			output.FromAddress = common.BytesToAddress(vLog.Topics[1].Bytes()).String()
			output.ToAddress = common.BytesToAddress(vLog.Topics[2].Bytes()).String()
			output.TxHash = vLog.TxHash.String()
			output.BlockNumber = vLog.BlockNumber

			events = append(events, output)
		}
		if err := fn(events); err != nil {
			return errors.Wrapf(err, "handle the transfer event from ETH smartcontract fail")
		}

		// save last scanned block number to db
		err = s.DBClient.SetLastScannedBlockNumber(s.network, s.pegWalletAddress, toBlock.Int64())
		if err != nil {
			return errors.Wrapf(err, "set last scanned block number fail")
		}

		// set current block number = latest + 1 for next tick interval
		s.transferCurrentBlock.Add(toBlock, big.NewInt(1))
	}

	return nil
}
