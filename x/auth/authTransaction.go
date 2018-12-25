package auth

import (
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/go-amino"

	"github.com/sharering/shareledger/constants"
	"github.com/sharering/shareledger/types"
)

// AuthTx is of interface SHRTx
var _ types.SHRTx = AuthTx{}

type AuthTx struct {
	sdk.Msg   `json:"message"`
	Signature AuthSig `json:"signature"`
}

func NewAuthTx(msg sdk.Msg, sig AuthSig) AuthTx {
	return AuthTx{
		Msg:       msg,
		Signature: sig,
	}
}

// GetMsgs returns multiple messages
func (tx AuthTx) GetMsgs() []sdk.Msg {
	return []sdk.Msg{tx.Msg}
}

// GetMsg returns the message of this transaction
func (tx AuthTx) GetMsg() sdk.Msg {
	return tx.Msg
}

// GetSignature returns the signature with this transaction
func (tx AuthTx) GetSignature() types.SHRSignature {
	return tx.Signature
}

// GetNonce returns Nonce sent with the signature
func (tx AuthTx) GetNonce() int64 {
	return tx.Signature.GetNonce()
}

// GetSignBytes returns Bytes to be signed
func (tx AuthTx) GetSignBytes() []byte {
	return tx.Msg.GetSignBytes()
}

// VerifySignature to verify signature
func (tx AuthTx) VerifySignature() bool {
	msg := tx.GetSignBytes()
	constants.LOGGER.Info("SignBytes", "signBytes", msg)
	return tx.Signature.Verify(msg)
}

// JSON decode MsgSend.
func GetTxDecoder(cdc *amino.Codec) func([]byte) (sdk.Tx, sdk.Error) {
	return func(txBytes []byte) (sdk.Tx, sdk.Error) {
		var tx types.SHRTx

		//fmt.Println("TxDecoder:", txBytes)
		//err := json.Unmarshal(txBytes, &tx)
		//err := cdc.UnmarshalJSON(txBytes, &tx)
		err := cdc.UnmarshalBinaryLengthPrefixed(txBytes, &tx)

		if err != nil {
			constants.LOGGER.Error("Error in decoding Tx", "err", err.Error())
			return nil, sdk.ErrTxDecode(err.Error())
		}
		(constants.LOGGER).Info("Decoded Tx", "tx", tx)
		//isVerified := tx.VerifySignature()
		//if !isVerified {
		//return nil, sdk.ErrTxDecode("InvalidSignature")
		//}
		return tx, nil
	}
}

//-------------------------------------------------------------------
// AuthSig

var _ types.SHRSignature = AuthSig{}

type AuthSig struct {
	types.PubKey    `json:"pub_key"`
	types.Signature `json:"signature"`
	Nonce           int64 `json:"nonce"`
}

func NewAuthSig(key types.PubKey, sig types.Signature, nonce int64) AuthSig {
	return AuthSig{
		PubKey:    key,
		Signature: sig,
		Nonce:     nonce,
	}
}

func (sig AuthSig) String() string {
	return fmt.Sprintf("AuthSig{%s, %s, %d}", sig.PubKey, sig.Signature, sig.Nonce)
}

// Verify signature according to message
// Prefix message with a nonce
func (sig AuthSig) Verify(msg []byte) bool {
	// convert Nonce to byte
	nonceBytes := []byte(strconv.Itoa(int(sig.Nonce)))

	// Prefix msg with Nonce
	msg = append(nonceBytes, msg...)

	return sig.PubKey.VerifyBytes(msg, sig.Signature)
}

func (sig AuthSig) GetPubKey() types.PubKey {
	return sig.PubKey
}

// GetNonce returns Nonce from Signature
func (sig AuthSig) GetNonce() int64 {
	return sig.Nonce
}

//------------------------------------------------------------
// Function for testing

// GetAuthTx - create an AuthTx message
func GetAuthTx(pubKey types.PubKey, privKey types.PrivKey, msg sdk.Msg, nonce int64) AuthTx {

	sig := privKey.SignWithNonce(msg, nonce)

	authSig := NewAuthSig(pubKey, sig, nonce)

	return NewAuthTx(msg, authSig)
}
