package auth

import (
	"strconv"

	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
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
func (tx AuthTx) GetSignature() SHRSignature {
	return tx.Signature.(SHRSignature)
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
	fmt.Printf("SignBytes: %s\n", msg)
	return tx.Signature.Verify(msg)
}

// JSON decode MsgSend.
func GetTxDecoder(cdc *wire.Codec) func([]byte) (sdk.Tx, sdk.Error) {
	return func(txBytes []byte) (sdk.Tx, sdk.Error) {
		var tx = SHRTx{}

		//fmt.Println("TxDecoder:", txBytes)
		//err := json.Unmarshal(txBytes, &tx)
		err := cdc.UnmarshalJSON(txBytes, &tx)

		if err != nil {
			fmt.Println("Error in decoding")
			return nil, sdk.ErrTxDecode(err.Error())
		}
		fmt.Println("Decoded Tx:", tx)
		isVerified := tx.VerifySignature()
		if !isVerified {
			return nil, sdk.ErrTxDecode("InvalidSignature")
		}
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
	return fmt.Sprintf("AuthSig{%s, %s, %s}", sig.PubKey, sig.Signature, sig.nonce)
}

// Verify signature according to message
// Prefix message with a nonce
func (sig AuthSig) Verify(msg []byte) bool {
	// convert Nonce to byte
	nonceBytes := []byte(strconv.Itoa(sig.Nonce))

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
