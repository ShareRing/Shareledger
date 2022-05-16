package database

import (
	"context"
	"fmt"
	"time"

	"github.com/rs/zerolog/log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	*mongo.Client
}

func (c *DB) InsertBatches(batches []Batch) error {
	//TODO implement me
	panic("implement me")
}

func (c *DB) UpdateLatestScannedBatchId(id uint64) error {
	//TODO implement me
	panic("implement me")
}

func (c *DB) GetLastScannedBatch() (uint64, error) {
	//TODO implement me
	panic("implement me")
}

type Collection struct {
	*mongo.Collection
}

const (
	ShareRing         = "sharering"
	RequestCollection = "requests"
	BatchCollection   = "batches"
	SettingCollection = "settings"
	AddressCollection = "addresses"
	LogsCollection    = "logs"
)

func (c *DB) GetSLP3Address(erc20Addr, network string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := c.GetCollection(ShareRing, AddressCollection)

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

	collection := c.GetCollection(ShareRing, SettingCollection)

	_, err := collection.UpdateMany(
		ctx,
		bson.M{},
		bson.D{
			{
				Key: "$set",
				Value: bson.D{
					{
						Key:   fmt.Sprintf("lastScannedBlockNumber.%s", contractAddress),
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

	collection := c.GetCollection(ShareRing, BatchCollection)

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

func (c *DB) GetNextPendingBatchOut(network string) (*Batch, error) {
	return c.getOneBatchStatus(network, Pending)
}

func (c *DB) getOneBatchStatus(network string, status Status) (*Batch, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := c.GetCollection(ShareRing, BatchCollection)

	var batch Batch
	err := collection.FindOne(ctx, bson.M{
		"status":  status,
		"network": network,
	}, &options.FindOneOptions{
		Sort: bson.M{
			"$sort": bson.M{
				"shareledgerID": 1,
			},
		},
	}).Decode(&batch)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &batch, err
}

func (c *DB) SearchBatchByType(shareledgerID uint64, requestType Type) (*Batch, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := c.GetCollection(ShareRing, BatchCollection)

	var queryResult Batch

	err := collection.FindOne(ctx, bson.M{
		"shareledgerID": shareledgerID,
		"type":          requestType,
	}).Decode(&queryResult)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return &Batch{}, err
	}

	return &queryResult, nil
}
func (c *DB) SearchBatchByStatus(network string, status Status, nonce uint64) ([]Batch, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := c.GetCollection(ShareRing, BatchCollection)

	var queryResult []Batch

	cursor, err := collection.Find(ctx, bson.M{
		"network": network,
		"status":  status,
		"nonce": bson.M{
			"$lt": nonce,
		},
	}, nil)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	if err = cursor.All(ctx, &queryResult); err != nil {
		return nil, err
	}

	return queryResult, nil
}
func (c *DB) GetBatchByTxHash(txHash string) (Batch, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := c.GetCollection(ShareRing, BatchCollection)

	var queryResult Batch

	err := collection.FindOne(ctx, bson.M{
		"txHash": txHash,
	}).Decode(&queryResult)
	if err != nil {
		return Batch{}, err
	}

	return queryResult, nil
}

func (c *DB) SetBatch(request Batch) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := c.GetCollection(ShareRing, BatchCollection)
	upsert := true
	_, err := collection.UpdateOne(ctx, bson.M{
		"shareledgerID": request.ShareledgerID,
		"type":          request.Type,
	}, request, &options.UpdateOptions{Upsert: &upsert})
	if err != nil {
		return err
	}

	return nil
}

func (c *DB) SetLog(batchId uint64, msg string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := c.GetCollection(ShareRing, LogsCollection)
	_, err := collection.InsertOne(ctx, Logs{
		BathID:  batchId,
		Message: msg,
	})
	return err
}

func NewMongo(mongoURI string) (DBRelayer, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	return &DB{
		Client: client,
	}, err
}

func (c *DB) ConnectDB(ctx context.Context) error {
	err := c.Client.Connect(ctx)
	if err != nil {
		return err
	}

	//ping the database
	err = c.Client.Ping(ctx, nil)
	if err != nil {
		return err
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
