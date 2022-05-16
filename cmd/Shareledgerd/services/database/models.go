package database

import "go.mongodb.org/mongo-driver/bson/primitive"

type Status string

const (
	Pending   Status = "pending"
	Done      Status = "done"
	Cancelled Status = "cancelled"
	Failed    Status = "failed"
)

type Type string

const (
	In  Type = "in"
	Out Type = "out"
)

type Batch struct {
	ID            primitive.ObjectID `bson:"_id" json:"id"`
	ShareledgerID uint64             `bson:"shareledgerID" json:"shareledgerID"`
	Status        Status             `bson:"status" json:"status"`
	Type          Type               `bson:"type"  json:"type"`
	TxHash        string             `bson:"txHash" json:"txHash"`
	Network       string             `bson:"network" json:"network"`
	BlockNumber   uint64             `bson:"blockNumber" json:"blockNumber"`
	Nonce         uint64             `bson:"nonce" json:"nonce"`
}

type Logs struct {
	ID      primitive.ObjectID `bson:"_id" json:"id"`
	BathID  uint64             `bson:"batchID" json:"batchID"`
	Message string             `bson:"message" json:"message"`
}

type Network string

type Address struct {
	ShareledgerAddress string  `json:"shareledgerAddress,omitempty" validate:"required"`
	AccIndex           uint32  `json:"accIndex,omitempty" validate:"required"`
	MnemonicHash       string  `json:"mnemonicHash,omitempty" validate:"required"`
	Network            Network `json:"network,omitempty" validate:"required"`
	Result             string  `json:"result,omitempty" validate:"required"`
}

type Setting struct {
	LastScannedBatchID     uint64            `bson:"lastScannedBatchID"`
	LastScannedBlockNumber map[string]uint64 `bson:"lastScannedBlockNumber"`
}
