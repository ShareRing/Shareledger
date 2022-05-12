package database

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Client struct {
	*mongo.Client
}

func (c *Client) UpdateRequests(shareledgerIDs []uint64, status Status) error {
	//TODO implement me
	panic("implement me")
}

func (c *Client) GetRequestByType(shareledgerID, requestType string) (Request, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) GetRequestByTxHash(txHash string) (Request, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) SetRequest(request Request) {
	//TODO implement me
	panic("implement me")
}

type RequestCollection struct {
	*mongo.Collection
}

func NewMongo(mongoURI string) (DBRelayer, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	return &Client{
		Client: client,
	}, err
}

func ConnectDB(mongoURI string) (*Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) // timeout should be in config
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	log.Info().Msg("Connected to MongoDB")
	return &Client{
		Client: client,
	}, nil
}

//getting database collections
func (c *Client) GetCollection(dbName, collectionName string) *RequestCollection {
	collection := c.Database(dbName).Collection(collectionName)
	return &RequestCollection{
		Collection: collection,
	}
}

func (rc *RequestCollection) GetRequestByType(shareledgerID, requestType string) (Request, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var queryResult Request

	filter := bson.M{
		"shareledgerID": shareledgerID,
		"type":          requestType,
	}
	err := rc.FindOne(ctx, filter).Decode(&queryResult)
	if err != nil {
		return Request{}, err
	}

	return queryResult, nil
}

func (rc *RequestCollection) GetRequestByTxHash(txHash string) (Request, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var queryResult Request

	filter := bson.M{
		"txHash": txHash,
	}
	err := rc.FindOne(ctx, filter).Decode(&queryResult)
	if err != nil {
		return Request{}, err
	}

	return queryResult, nil
}

func (rc *RequestCollection) SetRequest(request Request) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id, err := rc.InsertOne(ctx, bson.M{
		"shareledgerID": request.ShareledgerID,
		"status":        request.Status,
		"txHash":        request.TxHash,
		"type":          request.Type,
		"network":       request.Network,
		"blockNumber":   request.BlockNumber,
		"nonce":         request.Nonce,
	})
	if err != nil {
		return "", err
	}

	return id.InsertedID, err
}
