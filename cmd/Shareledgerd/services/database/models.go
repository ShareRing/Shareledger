package database

import "go.mongodb.org/mongo-driver/bson/primitive"

type Status string

const (
	Pending   Status = "pending"
	Done      Status = "done"
	Submitted Status = "submitted"
	Cancelled Status = "cancelled"
	Failed    Status = "failed"
)

type Type string

const (
	In  Type = "in"
	Out Type = "out"
)

type Batch struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ShareledgerID uint64             `bson:"shareledgerID" json:"shareledgerID"`
	Status        Status             `bson:"status" json:"status"`
	Type          Type               `bson:"type"  json:"type"`
	TxHash        string             `bson:"txHash" json:"txHash"`
	Network       string             `bson:"network" json:"network"`
	BlockNumber   uint64             `bson:"blockNumber" json:"blockNumber"`
	Nonce         uint64             `bson:"nonce" json:"nonce"`
	Signer        string             `bson:"signer" json:"signer"`
}

type Logs struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	BathID  uint64             `bson:"batchID" json:"batchID"`
	Message string             `bson:"message" json:"message"`
}

type Network string

type Address struct {
	ShareledgerAddress string  `bson:"shareledgerAddress" validate:"required"`
	AccIndex           uint32  `bson:"accIndex" validate:"required"`
	MnemonicHash       string  `bson:"mnemonicHash" validate:"required"`
	Network            Network `bson:"network" validate:"required"`
	Result             string  `bson:"result" validate:"required"`
}

type Setting struct {
	LastScannedBatchID     map[Network]uint64 `bson:"lastScannedBatchID"`
	LastScannedBlockNumber map[string]uint64  `bson:"lastScannedBlockNumber"`
}
