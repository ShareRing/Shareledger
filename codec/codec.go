package codec

import (
	"bytes"
	"encoding/json"

	amino "github.com/tendermint/go-amino"
	"github.com/tendermint/go-crypto"
)

// amino codec to marshal/unmarshal
type Codec = amino.Codec

func New() *Codec {
	cdc := amino.NewCodec()
	return cdc
}

// Register the go-crypto to the codec
func RegisterCrypto(cdc *Codec) {
	registerAmino(cdc)
}

func registerAmino(cdc *Codec) {
	cdc.RegisterInterface((*crypto.PubKey)(nil), nil)
	cdc.RegisterConcrete(crypto.PubKeyEd25519{},
		"tendermint/PubKeyEd25519", nil)
	cdc.RegisterConcrete(crypto.PubKeySecp256k1{},
		"tendermint/PubKeySecp256k1", nil)

	cdc.RegisterInterface((*crypto.PrivKey)(nil), nil)
	cdc.RegisterConcrete(crypto.PrivKeyEd25519{},
		"tendermint/PrivKeyEd25519", nil)
	cdc.RegisterConcrete(crypto.PrivKeySecp256k1{},
		"tendermint/PrivKeySecp256k1", nil)

}

// attempt to make some pretty json
func MarshalJSONIndent(cdc *Codec, obj interface{}) ([]byte, error) {
	bz, err := cdc.MarshalJSON(obj)
	if err != nil {
		return nil, err
	}

	var out bytes.Buffer
	err = json.Indent(&out, bz, "", "  ")
	if err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}

//__________________________________________________________________

// generic sealed codec to be used throughout sdk
var Cdc *Codec

func init() {
	cdc := New()
	// RegisterCrypto(cdc)
	// Cdc = cdc.Seal()
	Cdc = cdc
}
