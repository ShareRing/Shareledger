package database

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"
	denom "github.com/sharering/shareledger/x/utils/demo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"

	"github.com/rs/zerolog/log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	*mongo.Client
	DBName string
}

func (c *DB) GetSubmittedBatchesIn(network string) ([]BatchIn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	bCol := c.GetCollection(c.DBName, BatchCollection)
	var batches []BatchIn
	cur, err := bCol.Find(ctx, bson.M{
		"type":    "in",
		"status":  "submitted",
		"network": network,
	})
	if err != nil {
		return nil, err
	}

	err = cur.All(context.Background(), &batches)
	if err != nil {
		return nil, err
	}
	return batches, nil
}

func (c *DB) GetBatchInByTxHashes(network string, txHashes []string) (*BatchIn, error) {
	rCol := c.GetCollection(c.DBName, BatchCollection)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	var request BatchIn
	result := rCol.FindOne(ctx, bson.M{"txHashes": txHashes, "network": network}, nil)
	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, result.Err()
	}
	err := result.Decode(&request)
	return &request, err
}

func (c *DB) UnBatchRequestIn(network string, submittedTxHashes []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	bCol := c.GetCollection(c.DBName, BatchCollection)
	cursor, err := bCol.Find(ctx, bson.M{
		"type":    BatchTypeIn,
		"network": network,
		"txHashes": bson.M{
			"$in": submittedTxHashes,
		},
	})
	if err != nil {
		return err
	}
	var batches []BatchIn
	if err := cursor.Decode(&batches); err != nil {
		return err
	}
	unBatchID := make([]primitive.ObjectID, 0, len(batches))
	for i := range batches {
		unBatchID = append(unBatchID, batches[i].ID)
	}
	_, err = bCol.UpdateMany(ctx, bson.M{
		"_id": bson.M{
			"$in": unBatchID,
		},
	}, bson.M{
		"$set": bson.M{"status": BatchStatusFailed},
	})
	if err != nil {
		return err
	}

	rCol := c.GetCollection(c.DBName, RequestInCollection)
	_, err = rCol.UpdateMany(ctx, bson.M{
		"network": network,
		"batchID": bson.M{
			"$in": unBatchID,
		},
		"txHash": bson.M{
			"$nin": submittedTxHashes,
		},
	}, bson.M{
		"$set": bson.M{
			"status":  RequestInPending,
			"batchID": nil,
		},
	})
	return err
}

func (c *DB) GetPendingBatchesIn(ctx context.Context, network string) ([]BatchIn, error) {
	rCol := c.GetCollection(c.DBName, BatchCollection)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	var requests []BatchIn
	r, err := rCol.Find(ctx, bson.M{
		"status":  BatchStatusPending,
		"type":    BatchTypeIn,
		"network": network,
	}, nil)
	if err != nil {
		return nil, err
	}
	err = r.All(ctx, &requests)
	return requests, err
}

func (c *DB) InsertRequestIn(request RequestsIn) error {
	rCol := c.GetCollection(c.DBName, RequestInCollection)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	_, err := rCol.InsertOne(ctx, request)
	return err
}

func (c *DB) GetRequestIn(network string, txHash string) (*RequestsIn, error) {
	rCol := c.GetCollection(c.DBName, RequestInCollection)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	var req RequestsIn
	err := rCol.FindOne(ctx, bson.M{
		"txHash":  txHash,
		"network": network,
	}).Decode(&req)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return nil, err
		}
		return nil, nil
	}
	return &req, nil
}

func (c *DB) GetPendingRequestsIn(network string, destAddress string) ([]RequestsIn, error) {
	rCol := c.GetCollection(c.DBName, RequestInCollection)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	var requests []RequestsIn
	r, err := rCol.Find(ctx, bson.M{"network": network, "status": RequestInPending, "destAddress": destAddress}, nil)
	if err != nil {
		return nil, err
	}
	err = r.All(ctx, &requests)
	return requests, err
}

func (c *DB) TryToBatchPendingSwapIn(network string, destAddress string, minFee sdk.Coin) error {
	pendingRequests, err := c.GetPendingRequestsIn(network, destAddress)
	if err != nil {
		return err
	}
	if len(pendingRequests) == 0 {
		return nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	totalSwapIn := sdk.NewCoin(denom.Base, sdk.NewInt(0))
	ids := make([]primitive.ObjectID, 0, len(pendingRequests))
	txHashes := make([]string, 0, len(pendingRequests))
	for _, pr := range pendingRequests {
		coin, err := sdk.ParseCoinNormalized(pr.BaseAmount)
		if err != nil {
			return errors.New(fmt.Sprintf("pending request, txHash, %x, does not have correct amount value, %s", pr.TxHash, pr.BaseAmount))
		}
		totalSwapIn = totalSwapIn.Add(coin)
		ids = append(ids, pr.ID)
		txHashes = append(txHashes, pr.TxHash)
	}
	if totalSwapIn.IsLT(minFee) {
		// skip this destAddr for next time.
		return nil
	}
	bCol := c.GetCollection(c.DBName, BatchCollection)
	ires, err := bCol.InsertOne(ctx, BatchIn{
		Batch: Batch{
			Status:   BatchStatusPending,
			Type:     BatchTypeIn,
			Network:  network,
			TxHashes: txHashes,
		},
		BaseAmount: totalSwapIn.Sub(minFee).String(),
		BaseFee:    minFee.String(),
		DestAddr:   destAddress,
	})
	if err != nil {
		return err
	}

	inCol := c.GetCollection(c.DBName, RequestInCollection)
	uFilter := bson.M{
		"_id": bson.M{
			"$in": ids,
		},
	}
	_, err = inCol.UpdateMany(ctx, uFilter, bson.M{
		"$set": bson.M{
			"status":  RequestInBatched,
			"batchID": ires.InsertedID,
		},
	})
	return err
}

func (c *DB) SetBatchesOutFailed(network string, nonceNumber uint64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := c.GetCollection(c.DBName, BatchCollection)

	_, err := collection.UpdateMany(
		ctx,
		bson.M{
			"nonce": bson.M{
				"$lte": nonceNumber,
			},
			"type":    "out",
			"status":  BatchStatusSubmitted,
			"network": network,
		},
		bson.D{
			{Key: "$set", Value: bson.D{{Key: "status", Value: BatchStatusFailed}}},
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func (c *DB) InsertBatchesOut(batches []BatchOut) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	collection := c.GetCollection(c.DBName, BatchCollection)
	var newDocs []interface{}
	for _, v := range batches {
		newDocs = append(newDocs, v)
	}

	_, err := collection.InsertMany(ctx, newDocs)
	if err != nil {
		return err
	}

	return nil
}

func (c *DB) UpdateLatestScannedBatchId(id uint64, network string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	collection := c.GetCollection(c.DBName, StateCollection)
	upsert := true
	_, err := collection.UpdateMany(
		ctx,
		bson.M{
			"network": network,
		},
		bson.D{
			{
				Key: "$set",
				Value: bson.D{
					{
						Key:   "lastScannedBatchID",
						Value: id,
					},
				},
			},
		}, &options.UpdateOptions{Upsert: &upsert})

	if err != nil {
		return err
	}

	return nil
}

func (c *DB) GetLastScannedBlockNumber(network string, contractAddr string) (uint64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	collection := c.GetCollection(c.DBName, StateCollection)

	var state RelayerNetworkState
	err := collection.FindOne(ctx, bson.M{
		"network": network,
	}).Decode(&state)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return 0, nil
		}
		return 0, err
	}

	return state.LastScannedEventBlockNumbers[contractAddr], nil
}

func (c *DB) GetLastScannedBatch(network string) (uint64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	collection := c.GetCollection(c.DBName, StateCollection)

	var state RelayerNetworkState
	err := collection.FindOne(ctx, bson.M{
		"network": network,
	}).Decode(&state)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return 0, nil
		}
		return 0, err
	}

	if err != nil {
		return 0, err
	}

	return state.LastScannedBatchID, nil
}

type Collection struct {
	*mongo.Collection
}

const (
	RequestInCollection = "requestsIn"

	BatchCollection   = "batches"
	AddressCollection = "addresses"
	StateCollection   = "states"
	LogsCollection    = "logs"
	timeout           = 100 * time.Second
)

func (c *DB) GetRequestsInByBatchID(batchID primitive.ObjectID) ([]RequestsIn, error) {
	rCol := c.GetCollection(c.DBName, RequestInCollection)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	var result []RequestsIn
	cursor, err := rCol.Find(ctx, bson.M{
		"batchID": batchID,
	})
	if err != nil {
		return nil, err
	}
	if cursor.Err() == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err := cursor.All(ctx, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *DB) GetSLP3Address(erc20Addr, network string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := c.GetCollection(c.DBName, AddressCollection)

	var queryResult Address
	result := collection.FindOne(ctx, bson.M{
		"result":  erc20Addr,
		"network": network,
	})
	if err := result.Err(); err != nil {
		if err != mongo.ErrNoDocuments {
			return "", err
		}
		return "", err
	}
	err := result.Decode(&queryResult)
	if err != nil {
		return "", err
	}
	return queryResult.ShareledgerAddress, nil
}

func (c *DB) SetLastScannedBlockNumber(network string, contractAddress string, lastScannedBlockNumber int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := c.GetCollection(c.DBName, StateCollection)
	upsert := true
	_, err := collection.UpdateMany(
		ctx,
		bson.M{
			"network": network,
		},
		bson.D{
			{
				Key: "$set",
				Value: bson.D{
					{
						Key:   fmt.Sprintf("lastScannedEventBlockNumbers.%s", contractAddress),
						Value: lastScannedBlockNumber,
					},
				},
			},
		}, &options.UpdateOptions{Upsert: &upsert})

	if err != nil {
		return err
	}

	return nil
}

func (c *DB) UpdateBatchesOut(shareledgerIDs []uint64, status BatchStatus) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := c.GetCollection(c.DBName, BatchCollection)

	_, err := collection.UpdateMany(
		ctx,
		bson.M{
			"shareledgerID": bson.M{
				"$in": shareledgerIDs,
			},
			"type": "out",
		},
		bson.D{
			{Key: "$set", Value: bson.D{{Key: "status", Value: status}}},
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func (c *DB) GetNextUnfinishedBatchOut(network string, offset int64) (*BatchOut, error) {
	// submitted is preferred to process first
	submittedBatch, err := c.getOneBatchStatus(network, BatchStatusSubmitted, BatchTypeOut, &offset)
	if err != nil {
		return nil, err
	}
	if submittedBatch != nil {
		return submittedBatch, nil
	}
	return c.getOneBatchStatus(network, BatchStatusPending, BatchTypeOut, &offset)
}

func (c *DB) getOneBatchStatus(network string, status BatchStatus, bType BatchType, offset *int64) (*BatchOut, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := c.GetCollection(c.DBName, BatchCollection)

	var batch BatchOut
	err := collection.FindOne(ctx, bson.M{
		"status":  status,
		"network": network,
		"type":    bType,
	}, &options.FindOneOptions{
		Sort: bson.M{
			"shareledgerID": 1,
		},
		Skip: offset,
	}).Decode(&batch)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, errors.Wrapf(err, "get one batch by status from mongodb fail")
	}

	return &batch, err
}

func (c *DB) SearchBatchByType(shareledgerID uint64, requestType BatchType) (*BatchOut, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := c.GetCollection(c.DBName, BatchCollection)

	var queryResult BatchOut

	err := collection.FindOne(ctx, bson.M{
		"shareledgerID": shareledgerID,
		"type":          requestType,
	}).Decode(&queryResult)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return &BatchOut{}, errors.Wrapf(err, "search batch by batch type from mongodb fail")
	}

	return &queryResult, nil
}
func (c *DB) SearchBatchByStatus(network string, status BatchStatus) ([]BatchOut, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := c.GetCollection(c.DBName, BatchCollection)

	var queryResult []BatchOut

	cursor, err := collection.Find(ctx, bson.M{
		"network": network,
		"status":  status,
	}, nil)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, errors.Wrapf(err, "search batch by status from mongodb fail")
	}

	if err = cursor.All(ctx, &queryResult); err != nil {
		return nil, errors.Wrapf(err, "decoding query result to struct fail")
	}

	return queryResult, nil
}
func (c *DB) SearchUnSyncedBatchOutByStatus(network string, status BatchStatus) ([]BatchOut, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := c.GetCollection(c.DBName, BatchCollection)

	var queryResult []BatchOut

	cursor, err := collection.Find(ctx, bson.M{
		"network": network,
		"status":  status,
		"synced":  false,
		"type":    BatchTypeOut,
	}, nil)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, errors.Wrapf(err, "search batch by status from mongodb fail")
	}

	if err = cursor.All(ctx, &queryResult); err != nil {
		return nil, errors.Wrapf(err, "decoding query result to struct fail")
	}

	return queryResult, nil
}
func (c *DB) GetBatchOutByTxHash(network string, txHash string) (*BatchOut, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := c.GetCollection(c.DBName, BatchCollection)

	var queryResult BatchOut

	err := collection.FindOne(ctx, bson.M{
		"txHashes": txHash,
		"type":     BatchTypeOut,
		"network":  network,
	}).Decode(&queryResult)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return nil, errors.Wrapf(err, "decoding query result to struct fail")
		}
		return nil, nil
	}

	return &queryResult, nil
}

func (c *DB) SetBatches(batches []IBatch) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := c.GetCollection(c.DBName, BatchCollection)

	var operations []mongo.WriteModel
	for _, b := range batches {
		operations = append(operations,
			b.SetOperator(bson.M{
				"_id": b.GetID(),
			}, true),
		)
	}
	_, err := collection.BulkWrite(
		ctx, operations,
	)
	return err
}

func (c *DB) SetBatch(batch IBatch) error {
	if err := batch.Validate(); err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := c.GetCollection(c.DBName, BatchCollection)
	operation := batch.SetOperator(bson.M{
		"_id": batch.GetID(),
	}, true)

	_, err := collection.BulkWrite(
		ctx,
		[]mongo.WriteModel{
			operation,
		},
	)
	return err
}

func (c *DB) SetLog(batch IBatch, msg string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := c.GetCollection(c.DBName, LogsCollection)
	_, err := collection.InsertOne(ctx, Logs{
		Batch:   batch,
		Message: msg,
	})
	return errors.Wrapf(err, "insert log data into mongodb")
}

func NewMongo(mongoURI string, dbName string) (DBRelayer, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	return &DB{
		Client: client,
		DBName: dbName,
	}, errors.Wrapf(err, "fail to connect to mongodb")
}

func (c *DB) ConnectDB(ctx context.Context) error {
	err := c.Client.Connect(ctx)
	if err != nil {
		return err
	}

	//ping the database
	err = c.Client.Ping(ctx, nil)
	if err != nil {
		return errors.Wrapf(err, "testing mongodb connection fail")
	}

	log.Info().Msg("Connected to MongoDB")
	return nil
}

func (c *DB) Disconnect(ctx context.Context) error {
	return c.Client.Disconnect(ctx)
}

//getting database collections
func (c *DB) GetCollection(dbName, collectionName string) *Collection {
	collection := c.Database(dbName).Collection(collectionName)
	return &Collection{
		Collection: collection,
	}
}

//MarkBatchToSynced make the batch synced to true
func (c *DB) MarkBatchToSynced(sIDs []uint64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := c.GetCollection(c.DBName, BatchCollection)

	_, err := collection.UpdateMany(ctx, bson.M{
		"shareledgerID": bson.M{
			"$in": sIDs,
		},
	}, bson.M{
		"$set": bson.M{
			"synced": true,
		},
	})
	if err != nil {
		return errors.Wrapf(err, "update batch fail")
	}
	return nil
}
