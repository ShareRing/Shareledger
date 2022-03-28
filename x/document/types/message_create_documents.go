package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateDocuments{}

func NewMsgCreateDocuments(data []string, holder []string, issuer string, proof []string) *MsgCreateDocuments {
	return &MsgCreateDocuments{
		Data:   data,
		Holder: holder,
		Issuer: issuer,
		Proof:  proof,
	}
}

func (msg *MsgCreateDocuments) Route() string {
	return RouterKey
}

func (msg *MsgCreateDocuments) Type() string {
	return TypeMsgCreateDocs
}

func (msg *MsgCreateDocuments) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Issuer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateDocuments) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateDocuments) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Issuer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	// check len
	maxLen := len(msg.Holder)
	if maxLen == 0 || maxLen > MAX_LEN_BATCH || maxLen != len(msg.Proof) || maxLen != len(msg.Data) {
		return ErrDocInvalidData
	}

	for i := 0; i < maxLen; i++ {
		if len(msg.Holder[i]) > MAX_LEN || len(msg.Data[i]) > MAX_LEN || len(msg.Proof[i]) > MAX_LEN {
			return ErrDocInvalidData
		}
	}

	return nil
}
