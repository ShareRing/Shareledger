package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// type BaseID struct {
// 	IssuerAddr sdk.AccAddress `json:"issuer_address"`
// 	BackupAddr sdk.AccAddress `json:"backup_address"`
// 	OwnerAddr  sdk.AccAddress `json:"owner_address"`
// 	ExtraData  string         `json:"extra_data"`
// }

// type ID struct {
// 	Id     string `json:"id"`
// 	BaseID `json:"data"`
// }

// NewBaseID creates new BaseId instance
func NewBaseID(issuerAddr, backupAddr, ownerAddr, extraData string) BaseID {
	return BaseID{
		IssuerAddress: issuerAddr,
		BackupAddress: backupAddr,
		OwnerAddress:  ownerAddr,
		ExtraData:     extraData,
	}
}

// MustMarshalBaseID encodes data for storage
func MustMarshalBaseID(cdc codec.BinaryCodec, id *BaseID) []byte {
	return cdc.MustMarshalLengthPrefixed(id)
}

// MustUnmarshalBaseID decodes data from storage value. Throw exception when there is error
func MustUnmarshalBaseID(cdc codec.BinaryCodec, value []byte) BaseID {
	id, err := UnmarshalBaseID(cdc, value)
	if err != nil {
		panic(err)
	}
	return id
}

// UnmarshalBaseID decodes data from store value
func UnmarshalBaseID(cdc codec.BinaryCodec, value []byte) (id BaseID, err error) {
	err = cdc.UnmarshalLengthPrefixed(value, &id)

	return id, err
}

func (id ID) ToBaseID() BaseID {
	return NewBaseID(id.Data.IssuerAddress, id.Data.BackupAddress, id.Data.OwnerAddress, id.Data.ExtraData)
}

func NewIDFromBaseID(id string, ids *BaseID) ID {
	return ID{
		Id:   id,
		Data: ids,
	}
}

func NewID(id string, issuerAddr, backupAddr, ownerAddr, extraData string) ID {
	data := NewBaseID(issuerAddr, backupAddr, ownerAddr, extraData)
	return ID{Id: id, Data: &data}
}

func (id *ID) IsEmpty() bool {
	if id == nil {
		return true
	}

	if len(id.Data.IssuerAddress) == 0 {
		return true
	}

	return false
}
