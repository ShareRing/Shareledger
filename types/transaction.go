package types

import (
	//"encoding/json"
	"fmt"

	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
    "bitbucket.org/shareringvn/cosmos-sdk/wire"


)

//------------------------------------------------------------------
// Tx

// Simple tx to wrap the Msg.
type SHRTx struct {
	sdk.Msg   `json:"message"`
	Signature SHRSignature `json:"signature"`
}


func NewSHRTx(msg sdk.Msg, sig SHRSignature) SHRTx {
    return SHRTx{
        Msg: msg,
        Signature: sig,
    }
}

// GetMsgs returns multiple messages
func (tx SHRTx) GetMsgs() []sdk.Msg {
	return []sdk.Msg{tx.Msg}
}

// GetMsg returns the message of this transaction
func (tx SHRTx) GetMsg() sdk.Msg {
	return tx.Msg
}

// GetSignature returns the signature with this transaction
func (tx SHRTx) GetSignature() SHRSignature {
	return tx.Signature
}

// VerifySignature to verify signature
func (tx SHRTx) VerifySignature() bool {
    msg := tx.Msg.GetSignBytes()
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

//------------------------------------------------------------------
// Signature

type SHRSignature struct {
	PubKey `json:"pub_key"`
	Signature `json:"signature"`
}

func NewSHRSignature(key PubKey, sig Signature) SHRSignature {
    return SHRSignature{
        PubKey: key,
        Signature: sig,
    }
}


func (sig SHRSignature) String() string {
    return fmt.Sprintf("SHRSignature{%s, %s}", sig.PubKey, sig.Signature)
}

func (sig SHRSignature) Verify(msg []byte) bool {
    return sig.PubKey.VerifyBytes(msg, sig.Signature)
}
