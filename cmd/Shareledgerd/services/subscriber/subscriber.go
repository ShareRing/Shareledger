package subscriber

import (
	"context"
	"math/big"
	"strings"

	eth "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
)

const (
	eventName = "Transfer"
)

type EventInput struct {
	ContractAddress string
	Topic           string
	// OutputChannel   chan EventOutput
}

type EventOutput struct {
	FromAddress string
	ToAddress   string
	Amount      decimal.Decimal
	TxHash      string
}

type Service struct {
	client       *ethclient.Client
	currentBlock *big.Int
}

type NewInput struct {
	ProviderURL  string
	CurrentBlock *big.Int
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
		client:       client,
		currentBlock: input.CurrentBlock,
	}, nil
}

func (s *Service) GetEvents(ctx context.Context, input *EventInput) (events []EventOutput, err error) {

	erc20Abi, err := abi.JSON(strings.NewReader(string(AMetaData.ABI)))
	if err != nil {
		return nil, err
	}
	header, err := s.client.HeaderByNumber(ctx, nil)
	// skip if header not found
	if err != nil {
		return events, err
	}

	// if head = current => skip
	if header.Number.Cmp(s.currentBlock) == 0 {
		log.Info("there is no new")
		return events, nil
	}

	log.Debugf("Scanning from block %v to block %v", s.currentBlock, header.Number)
	any := []common.Hash{}
	query := eth.FilterQuery{
		FromBlock: s.currentBlock,
		ToBlock:   header.Number,
		Addresses: []common.Address{common.HexToAddress(input.ContractAddress)},
		Topics: [][]common.Hash{
			[]common.Hash{common.HexToHash(input.Topic)},
			any,
			any,
		},
	}

	logs, err := s.client.FilterLogs(ctx, query)
	if err != nil {
		return events, err
	}

	events = make([]EventOutput, 0, len(logs))

	for _, vLog := range logs {
		output := EventOutput{}
		var event = struct {
			Value *big.Int // amount in erc20 contract
		}{}

		err := erc20Abi.UnpackIntoInterface(&event, eventName, vLog.Data)
		if err != nil {
			log.Errorf("Event unpacking error: %s", err)
			continue
		}

		output.Amount = decimal.NewFromBigInt(event.Value, 0)
		output.FromAddress = common.BytesToAddress(vLog.Topics[1].Bytes()).String()
		output.ToAddress = common.BytesToAddress(vLog.Topics[2].Bytes()).String()
		output.TxHash = vLog.TxHash.String()

		events = append(events, output)

	}
	// set current block number = latest + 1 for next tick interval
	s.currentBlock.Add(header.Number, big.NewInt(1))
	return events, nil
}
