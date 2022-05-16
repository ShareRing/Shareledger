package subscriber

import (
	"context"
	"math/big"
	"strings"

	eth "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sharering/shareledger/cmd/Shareledgerd/services/database"
	"github.com/sharering/shareledger/cmd/Shareledgerd/services/subscriber/erc20"
	"github.com/sharering/shareledger/cmd/Shareledgerd/services/subscriber/swapcontract"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
)

const (
	transferEvent = "Transfer"
	swapEvent     = "SwapCompleted"
)

type EventTransferInput struct {
	PegWalletAddress string
	TransferTopic    string
}

type EventTransferOutput struct {
	FromAddress string
	ToAddress   string
	Amount      decimal.Decimal
	TxHash      string
	BlockNumber uint64
}

type EventSwapCompleteInput struct {
	SwapContractAddress string
	SwapTopic           string
}

type EventSwapCompleteOutput struct {
	TxHash string
}

type Service struct {
	client               *ethclient.Client
	transferCurrentBlock *big.Int
	swapCurrentBlock     *big.Int
	DBClient             database.DBRelayer
}

type NewInput struct {
	ProviderURL          string
	TransferCurrentBlock *big.Int
	SwapCurrentBlock     *big.Int
	DBClient             database.DBRelayer
}

func init() {
	log.SetLevel(log.DebugLevel)
}

func New(input *NewInput) (*Service, error) {
	client, err := ethclient.Dial(input.ProviderURL)
	if err != nil {
		return nil, err
	}

	return &Service{
		client:               client,
		transferCurrentBlock: input.TransferCurrentBlock,
		swapCurrentBlock:     input.SwapCurrentBlock,
		DBClient:             input.DBClient,
	}, nil
}

func (s *Service) GetSwapCompleteEvent(ctx context.Context, input *EventSwapCompleteInput) (events []EventSwapCompleteOutput, err error) {
	swapAbi, err := abi.JSON(strings.NewReader(string(swapcontract.SwapcontractMetaData.ABI)))
	if err != nil {
		return nil, err
	}

	header, err := s.client.HeaderByNumber(ctx, nil)
	if err != nil {
		return events, err
	}

	if header.Number.Cmp(s.swapCurrentBlock) == 0 {
		log.Info("there is no new block")
		return events, nil
	}

	log.Debugf("Scanning from block %v to block %v", s.swapCurrentBlock, header.Number)
	any := []common.Hash{}
	query := eth.FilterQuery{
		FromBlock: s.swapCurrentBlock,
		ToBlock:   header.Number,
		Addresses: []common.Address{common.HexToAddress(input.SwapContractAddress)},
		Topics: [][]common.Hash{
			[]common.Hash{common.HexToHash(input.SwapTopic)},
			any,
		},
	}

	logs, err := s.client.FilterLogs(ctx, query)
	if err != nil {
		return events, err
	}

	events = make([]EventSwapCompleteOutput, 0, len(logs))

	for _, vLog := range logs {
		output := EventSwapCompleteOutput{}
		var event = struct {
			Value *big.Int // amount in erc20 contract
		}{}

		err := swapAbi.UnpackIntoInterface(&event, swapEvent, vLog.Data)
		if err != nil {
			log.Errorf("Event unpacking error: %s", err)
			continue
		}

		output.TxHash = vLog.TxHash.String()

		events = append(events, output)

	}

	// save last scanned block number to db
	err = s.DBClient.SetLastScannedBlockNumber(input.SwapContractAddress, header.Number.Int64())
	if err != nil {
		return nil, err
	}

	// set current block number = latest + 1 for next tick interval
	s.swapCurrentBlock.Add(header.Number, big.NewInt(1))
	return events, nil
}

func (s *Service) GetTransferEvent(ctx context.Context, input *EventTransferInput) (events []EventTransferOutput, err error) {
	erc20Abi, err := abi.JSON(strings.NewReader(string(erc20.Erc20MetaData.ABI)))
	if err != nil {
		return nil, err
	}
	header, err := s.client.HeaderByNumber(ctx, nil)
	// skip if header not found
	if err != nil {
		return events, err
	}

	// if head = current => skip
	if header.Number.Cmp(s.transferCurrentBlock) == 0 {
		log.Info("there is no new block")
		return events, nil
	}

	log.Debugf("Scanning from block %v to block %v", s.transferCurrentBlock, header.Number)
	any := []common.Hash{}
	query := eth.FilterQuery{
		FromBlock: s.transferCurrentBlock,
		ToBlock:   header.Number,
		Addresses: []common.Address{common.HexToAddress(input.PegWalletAddress)},
		Topics: [][]common.Hash{
			[]common.Hash{common.HexToHash(input.TransferTopic)},
			any,
			any,
		},
	}

	logs, err := s.client.FilterLogs(ctx, query)
	if err != nil {
		return events, err
	}

	events = make([]EventTransferOutput, 0, len(logs))

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
	// save last scanned block number to db
	err = s.DBClient.SetLastScannedBlockNumber(input.PegWalletAddress, header.Number.Int64())
	if err != nil {
		return nil, err
	}

	// set current block number = latest + 1 for next tick interval
	s.transferCurrentBlock.Add(header.Number, big.NewInt(1))
	return events, nil
}
