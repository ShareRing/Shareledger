package database

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
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

func (c *DB) SetBatchesOutFailed(nonceNumber uint64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := c.GetCollection(c.DBName, BatchCollection)

	_, err := collection.UpdateMany(
		ctx,
		bson.M{
			"nonce": bson.M{
				"$lte": nonceNumber,
			},
			"type":   "out",
			"status": Submitted,
		},
		bson.D{
			{Key: "$set", Value: bson.D{{Key: "status", Value: Failed}}},
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func (c *DB) InsertBatches(batches []Batch) error {
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

	collection := c.GetCollection(c.DBName, SettingCollection)

	_, err := collection.UpdateMany(
		ctx,
		bson.M{},
		bson.D{
			{
				Key: "$set",
				Value: bson.D{
					{
						Key:   fmt.Sprintf("settings.lastScannedBatchID.%s", network),
						Value: id,
					},
				},
			},
		})

	if err != nil {
		return err
	}

	return nil
}

func (c *DB) GetLastScannedBlockNumber(contractAddr string) (uint64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	collection := c.GetCollection(c.DBName, SettingCollection)

	var queryResult bson.M
	var setting Setting
	_ = collection.FindOne(ctx, bson.M{}).Decode(&queryResult)
	doc, err := bson.Marshal(queryResult["settings"])
	if err != nil {
		return 0, err
	}

	err = bson.Unmarshal(doc, &setting)
	if err != nil {
		return 0, err
	}

	return setting.LastScannedBlockNumber[contractAddr], nil
}

func (c *DB) GetLastScannedBatch(network string) (uint64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	collection := c.GetCollection(c.DBName, SettingCollection)

	var queryResult bson.M
	var setting Setting
	_ = collection.FindOne(ctx, bson.M{}).Decode(&queryResult)
	doc, err := bson.Marshal(queryResult["settings"])
	if err != nil {
		return 0, err
	}

	err = bson.Unmarshal(doc, &setting)
	if err != nil {
		return 0, err
	}

	return setting.LastScannedBatchID[Network(network)], nil
}

type Collection struct {
	*mongo.Collection
}

const (
	RequestCollection = "requests"
	BatchCollection   = "batches"
	SettingCollection = "settings"
	AddressCollection = "addresses"
	LogsCollection    = "logs"
	timeout           = 10 * time.Second
)

func (c *DB) GetSLP3Address(erc20Addr, network string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := c.GetCollection(c.DBName, AddressCollection)

	var queryResult Address
	err := collection.FindOne(ctx, bson.M{
		"result":  erc20Addr,
		"network": network,
	}).Decode(&queryResult)
	if err != nil {
		return "", err
	}

	return queryResult.ShareledgerAddress, nil
}

func (c *DB) SetLastScannedBlockNumber(contractAddress string, lastScannedBlockNumber int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := c.GetCollection(c.DBName, SettingCollection)

	_, err := collection.UpdateMany(
		ctx,
		bson.M{},
		bson.D{
			{
				Key: "$set",
				Value: bson.D{
					{
						Key:   fmt.Sprintf("settings.lastScannedBlockNumber.%s", contractAddress),
						Value: lastScannedBlockNumber,
					},
				},
			},
		})

	if err != nil {
		return err
	}

	return nil
}

func (c *DB) UpdateBatchesOut(shareledgerIDs []uint64, status Status) error {
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

func (c *DB) GetNextPendingBatchOut(network string, offset int64) (*Batch, error) {
	return c.getOneBatchStatus(network, Pending, &offset)
}

func (c *DB) getOneBatchStatus(network string, status Status, offset *int64) (*Batch, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := c.GetCollection(c.DBName, BatchCollection)

	var batch Batch
	err := collection.FindOne(ctx, bson.M{
		"status":  status,
		"network": network,
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

func (c *DB) SearchBatchByType(shareledgerID uint64, requestType Type) (*Batch, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := c.GetCollection(c.DBName, BatchCollection)

	var queryResult Batch

	err := collection.FindOne(ctx, bson.M{
		"shareledgerID": shareledgerID,
		"type":          requestType,
	}).Decode(&queryResult)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return &Batch{}, errors.Wrapf(err, "search batch by batch type from mongodb fail")
	}

	return &queryResult, nil
}
func (c *DB) SearchBatchByStatus(network string, status Status) ([]Batch, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := c.GetCollection(c.DBName, BatchCollection)

	var queryResult []Batch

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
func (c *DB) GetBatchByTxHash(txHash string) (Batch, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := c.GetCollection(c.DBName, BatchCollection)

	var queryResult Batch

	err := collection.FindOne(ctx, bson.M{
		"txHash": txHash,
	}).Decode(&queryResult)
	if err != nil {
		return Batch{}, errors.Wrapf(err, "decoding query result to struct fail")
	}

	return queryResult, nil
}

func (c *DB) SetBatches(batches []Batch) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := c.GetCollection(c.DBName, BatchCollection)

	var operations []mongo.WriteModel
	for _, b := range batches {
		operations = append(operations, c.buildSetOperations(bson.M{
			"shareledgerID": bson.M{
				"$eq": b.ShareledgerID,
			},
			"type": bson.M{
				"$eq": b.Type,
			},
		}, b, true,
		))
	}
	_, err := collection.BulkWrite(
		ctx, operations,
	)
	return err
}

func (c *DB) buildSetOperations(filter bson.M, bach Batch, upsert bool) *mongo.UpdateOneModel {
	operation := mongo.NewUpdateOneModel()
	operation.SetFilter(filter)
	operation.Upsert = &upsert
	operation.SetUpdate(bson.M{"$set": bson.M{
		"shareledgerID": bach.ShareledgerID,
		"type":          bach.Type,
		"status":        bach.Status,
		"txHash":        bach.TxHash,
		"network":       bach.Network,
		"blockNumber":   bach.BlockNumber,
		"nonce":         bach.Nonce,
		"signer":        bach.Signer,
	}})
	return operation
}

func (c *DB) SetBatch(request Batch) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := c.GetCollection(c.DBName, BatchCollection)
	operation := c.buildSetOperations(bson.M{
		"shareledgerID": bson.M{
			"$eq": request.ShareledgerID,
		},
		"type": bson.M{
			"$eq": request.Type,
		},
	}, request, true)
	_, err := collection.BulkWrite(
		ctx,
		[]mongo.WriteModel{
			operation,
		},
	)
	return err
}

func (c *DB) SetLog(batchId uint64, msg string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := c.GetCollection(c.DBName, LogsCollection)
	_, err := collection.InsertOne(ctx, Logs{
		BathID:  batchId,
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
