package database

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	*mongo.Client
}

type Collection struct {
	*mongo.Collection
}

const (
	ShareRing         = "Sharering"
	RequestCollection = "Requests"
	BatchCollection   = "Batches"
	SettingCollection = "Settings"
)

func (c *DB) SetLastScannedBlockNumber(lastScannedBlockNumer uint32) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := c.GetCollection(ShareRing, SettingCollection)

	_, err := collection.UpdateMany(ctx, bson.M{}, bson.D{
		{Key: "$set", Value: bson.D{{Key: "lastScannedBlockNumber", Value: lastScannedBlockNumer}}},
	})

	if err != nil {
		return err
	}

	return nil
}

func (c *DB) UpdateBatches(shareledgerIDs []uint64, status Status) error {
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

func (c *DB) GetBatchByType(shareledgerID, requestType string) (Batch, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := c.GetCollection(ShareRing, BatchCollection)

	var queryResult Batch

	filter := bson.M{
		"shareledgerID": shareledgerID,
		"type":          requestType,
	}
	err := collection.FindOne(ctx, filter).Decode(&queryResult)
	if err != nil {
		return Batch{}, err
	}

	return queryResult, nil
}

func (c *DB) GetBatchByTxHash(txHash string) (Batch, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := c.GetCollection(ShareRing, BatchCollection)

	var queryResult Batch

	filter := bson.M{
		"txHash": txHash,
	}
	err := collection.FindOne(ctx, filter).Decode(&queryResult)
	if err != nil {
		return Batch{}, err
	}

	return queryResult, nil
}

func (c *DB) SetBatch(request Batch) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := c.GetCollection(ShareRing, BatchCollection)

	id, err := collection.InsertOne(ctx, bson.M{
		"shareledgerID": request.ShareledgerID,
		"status":        request.Status,
		"txHash":        request.TxHash,
		"type":          request.Type,
		"network":       request.Network,
		"blockNumber":   request.BlockNumber,
		"nonce":         request.Nonce,
	})
	if err != nil {
		return "", nil
	}

	return id.InsertedID, nil
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
