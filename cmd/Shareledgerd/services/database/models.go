package database

import "go.mongodb.org/mongo-driver/bson/primitive"

type Status string
type Type string

type Batch struct {
	ID            primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	ShareledgerID string             `json:"shareledgerID,omitempty" validate:"required"`
	Status        Status             `json:"status,omitempty" validate:"required"`
	Type          Type               `json:"type,omitempty" validate:"required"`
	TxHash        string             `json:"txHash,omitempty" validate:"required"`
	Network       string             `json:"network,omitempty" validate:"required"`
	BlockNumber   uint32             `json:"blockNumber,omitempty" validate:"required"`
	Nonce         uint32             `json:"nonce,omitempty" validate:"required"`
}
