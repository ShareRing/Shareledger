package database

import "go.mongodb.org/mongo-driver/bson/primitive"

type Status string

const (
	Pending Status = "pending"
	Done    Status = "done"
	Failed  Status = "failed"
)

type Type string

const (
	In  Type = "in"
	Out Type = "out"
)

type Batch struct {
	ID            primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	ShareledgerID uint64             `json:"shareledgerID,omitempty" validate:"required"`
	Status        Status             `json:"status,omitempty" validate:"required"`
	Type          Type               `json:"type,omitempty" validate:"required"`
	TxHash        string             `json:"txHash,omitempty" validate:"required"`
	Network       string             `json:"network,omitempty" validate:"required"`
	BlockNumber   uint64             `json:"blockNumber,omitempty" validate:"required"`
}

type Network string

type Address struct {
	ShareledgerAddress string  `json:"shareledgerAddress,omitempty" validate:"required"`
	AccIndex           uint32  `json:"accIndex,omitempty" validate:"required"`
	MnemonicHash       string  `json:"mnemonicHash,omitempty" validate:"required"`
	Network            Network `json:"network,omitempty" validate:"required"`
	Result             string  `json:"result,omitempty" validate:"required"`
}
