package types

import "github.com/cosmos/cosmos-sdk/codec"

// MustMarshalBaseID encodes data for storage
func MustMarshalBaseID(cdc codec.BinaryCodec, id *BaseID) []byte {
	return cdc.MustMarshal(id)
}

// MustUnmarshalBaseID decodes data from storage value. Throw exception when there is error
func MustUnmarshalBaseID(cdc codec.BinaryCodec, value []byte) BaseID {
	var id BaseID
	cdc.MustUnmarshal(value, &id)

	return id
}

func (id *Id) ToBaseID() BaseID {
	return BaseID{
		IssuerAddress: id.Data.IssuerAddress,
		BackupAddress: id.Data.BackupAddress,
		OwnerAddress:  id.Data.OwnerAddress,
		ExtraData:     id.Data.ExtraData,
	}
}

func NewIDFromBaseID(id string, bid *BaseID) Id {
	return Id{
		Id:   id,
		Data: bid,
	}
}
