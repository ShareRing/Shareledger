package types

import (
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"
)

func (d Document) GetDetailState() *DocDetailState {
	ds := DocDetailState{d.Data, d.Version}
	return &ds
}

// holder/issuer/proof
func (d Document) GetKeyDetailState() []byte {
	key := []byte{}

	key = append(key, []byte(Separator)...)
	key = append(key, []byte(d.Issuer)...)

	key = append(key, []byte(Separator)...)
	key = append(key, []byte(d.Proof)...)

	key = append([]byte(d.Holder), key...)

	return key
}

// holder
func (d Document) GetKeyDetailOfHolder() []byte {
	key := d.Holder

	return KeyPrefix(key)
}

func (d Document) GetKeyDetailHolderAndIssuer() []byte {
	key := []byte{}

	key = append(key, []byte(Separator)...)
	key = append(key, []byte(d.Issuer)...)

	key = append([]byte(d.Holder), key...)

	return key
}

func MustMarshalDocDetailState(cdc codec.BinaryCodec, ds *DocDetailState) []byte {
	return cdc.MustMarshalLengthPrefixed(ds)
}

func MustUnmarshalDocDetailState(cdc codec.BinaryCodec, value []byte) *DocDetailState {
	ds := DocDetailState{}

	err := cdc.UnmarshalLengthPrefixed(value, &ds)
	if err != nil {
		panic(err)
	}

	return &ds
}

func (d Document) GetBasicState() *DocBasicState {
	ds := DocBasicState{Holder: d.Holder, Issuer: d.Issuer}
	return &ds
}

// proof
func (d Document) GetKeyBasicState() []byte {
	key := d.Proof

	return KeyPrefix(key)
}

func MustMarshalDocBasicState(cdc codec.BinaryCodec, bs *DocBasicState) []byte {
	return cdc.MustMarshalLengthPrefixed(bs)
}

func MustUnmarshalDocBasicState(cdc codec.BinaryCodec, value []byte) *DocBasicState {
	bs := DocBasicState{}

	err := cdc.UnmarshalLengthPrefixed(value, &bs)
	if err != nil {
		panic(err)
	}

	return &bs
}

// holder/issuer/proof
func MustMarshalFromDetailRawState(cdc codec.BinaryCodec, key, value []byte) *Document {
	sKey := string(key)
	sKeyArr := strings.Split(sKey, Separator)

	doc := Document{}

	// issuer := sdk.AccAddress(sKeyArr[1])
	// // if err != nil {
	// // 	panic(err)
	// // }

	doc.Holder = sKeyArr[0]
	doc.Proof = sKeyArr[2]
	doc.Issuer = sKeyArr[1]
	// doc.Issuer = issuer.String()

	ds := MustUnmarshalDocDetailState(cdc, value)
	doc.Data = ds.Data
	doc.Version = ds.Version

	return &doc
}
