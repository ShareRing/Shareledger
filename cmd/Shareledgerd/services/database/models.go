package database

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BatchStatus string

const (
	BatchStatusPending   BatchStatus = "pending"
	BatchStatusDone      BatchStatus = "done"
	BatchStatusSubmitted BatchStatus = "submitted"
	BatchStatusCancelled BatchStatus = "cancelled"
	BatchStatusFailed    BatchStatus = "failed"
)

type BatchType string

const (
	BatchTypeOut BatchType = "out"
	BatchTypeIn  BatchType = "in"
)

type IBatch interface {
	BatchType() BatchType
}

type Batch struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ShareledgerID uint64             `bson:"shareledgerID" json:"shareledgerID"`
	Status        BatchStatus        `bson:"status" json:"status"`
	Type          BatchType          `bson:"type"  json:"type"`
	TxHashes      []string           `bson:"txHashes" json:"txHashes"`
	Network       string             `bson:"network" json:"network"`
	Signer        string             `bson:"signer" json:"signer"`
	Synced        bool               `bson:"synced" json:"synced"`
}

func InitBatch(b Batch) IBatch {
	switch b.batchType() {
	case BatchTypeOut:
		return BatchOut{
			Batch: b,
		}
	case BatchTypeIn:
		return BatchIn{
			Batch: b,
		}
	}
	return nil
}

func (b Batch) batchType() BatchType {
	return b.Type
}

type BatchOut struct {
	Batch       `bson:",inline"`
	BlockNumber uint64 `bson:"blockNumber" json:"blockNumber"`
	Nonce       uint64 `bson:"nonce" json:"nonce"`
}

func (b BatchOut) BatchType() BatchType {
	return b.Batch.batchType()
}

type BatchIn struct {
	Batch      `bson:",inline"`
	BaseAmount string `bson:"baseAmount"`
	BaseFee    string `bson:"baseFee"`
}

func (b BatchIn) BatchType() BatchType {
	return b.Batch.batchType()
}

type RequestInStatus string

const (
	RequestInPending RequestInStatus = "pending"
	RequestInBatched RequestInStatus = "batched"
)

type RequestsIn struct {
	ID          primitive.ObjectID  `bson:"_id,omitempty" json:"id"`
	Status      RequestInStatus     `bson:"status" json:"status"`
	TxHash      string              `bson:"txHash"`
	Network     string              `bson:"network"`
	DestAddress string              `bson:"destAddress"`
	SrcAddress  string              `bson:"srcAddress"`
	BaseAmount  string              `bson:"baseAmount"`
	BatchID     *primitive.ObjectID `bson:"batchID,omitempty"`
}

//type BatchesInStatus string

//const (
//	BatchesInPending   BatchesInStatus = "pending"
//	BatchesInSubmitted BatchesInStatus = "submitted"
//	BatchesInDone      BatchesInStatus = "done"
//)

//type BatchIn struct {
//	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
//	Status        BatchesInStatus    `bson:"status"`
//	ShareledgerID uint64             `bson:"ShareledgerID"`

//	Submitter     string             `bson:"submitter"`
//	Network       string             `bson:"network"`
//}

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

type RelayerNetworkState struct {
	Network                      string            `bson:"network"`
	LastScannedEventBlockNumbers map[string]uint64 `bson:"lastScannedEventBlockNumbers"` //[Contract]uint
	LastScannedBatchID           uint64            `bson:"lastScannedBatchID"'`
}
