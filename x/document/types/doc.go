package types

import (
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// type Doc struct {
// 	Holder  string         `json:"holder"`
// 	Issuer  sdk.AccAddress `json:"issuer"`
// 	Proof   string         `json:"proof"`
// 	Data    string         `json:"data"`
// 	Version uint16         `json:"version"`
// }

// type DocDetailState struct {
// 	Data    string
// 	Version uint16
// }

// type DocBasicState struct {
// 	Holder string
// 	Issuer sdk.AccAddress
// }

func (d Document) GetDetailState() *DocDetailState {
	ds := DocDetailState{d.Data, d.Version}
	return &ds
}

// 0x1|<hodler>|<proof>|<issuer>
func (d Document) GetKeyDetailState() []byte {
	key := []byte{}
	key = append(key, []byte(StateKeySep)...)
	key = append(key, []byte(d.Holder)...)

	key = append(key, []byte(StateKeySep)...)
	key = append(key, []byte(d.Proof)...)

	key = append(key, []byte(StateKeySep)...)
	key = append(key, []byte(d.Issuer)...)

	key = append(DocDetailPrefix, key...)

	return key
}

func (d Document) GetKeyDetailOfHolder() []byte {
	key := []byte{}
	key = append(key, StateKeySep...)
	key = append(key, d.Holder...)

	key = append(DocDetailPrefix, key...)

	return key
}

// Marshal doc state
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

// 0x2|<proof>
func (d Document) GetKeyBasicState() []byte {
	key := []byte{}
	key = append(key, []byte(StateKeySep)...)
	key = append(key, []byte(d.Proof)...)

	key = append(DocBasicPrefix, key...)
	return key
}

// Marshal doc state
func MustMarshalDocBasicState(cdc codec.BinaryCodec, bs *DocBasicState) []byte {
	return cdc.MustMarshalLengthPrefixed(bs)
}

func MustUnmarshalDocBasicState(cdc codec.BinaryCodec, value []byte) DocBasicState {
	ds := DocBasicState{}

	err := cdc.UnmarshalLengthPrefixed(value, &ds)
	if err != nil {
		panic(err)
	}

	return ds
}

func MustMarshalFromDetailRawState(cdc codec.BinaryCodec, key, value []byte) Document {
	sKey := string(key)
	sKeyArr := strings.Split(sKey, StateKeySep)
	doc := Document{}

	issuer, err := sdk.AccAddressFromBech32(sKeyArr[3])
	if err != nil {
		panic(err)
	}

	doc.Holder = sKeyArr[1]
	doc.Proof = sKeyArr[2]
	doc.Issuer = issuer.String()

	ds := MustUnmarshalDocDetailState(cdc, value)

	doc.Data = ds.Data
	doc.Version = ds.Version
	return doc
}
