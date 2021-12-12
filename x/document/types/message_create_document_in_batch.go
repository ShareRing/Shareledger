package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateDocumentInBatch{}

func NewMsgCreateDocumentInBatch(creator string, data string, holder string, issuer string, proof string) *MsgCreateDocumentInBatch {
	return &MsgCreateDocumentInBatch{
		Creator: creator,
		Data:    data,
		Holder:  holder,
		Issuer:  issuer,
		Proof:   proof,
	}
}

func (msg *MsgCreateDocumentInBatch) Route() string {
	return RouterKey
}

func (msg *MsgCreateDocumentInBatch) Type() string {
	return "CreateDocumentInBatch"
}

func (msg *MsgCreateDocumentInBatch) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateDocumentInBatch) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateDocumentInBatch) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
