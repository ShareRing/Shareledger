package types

import (
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeLoadSHRMsg          = "load_shr"
	TypeLoadSHRPMsg         = "load_shrp"
	TypeBurnSHRPMsg         = "burn_shrp"
	TypeBurnSHRMsg          = "burn_shr"
	TypeBuyCent             = "buy_cent"
	TypeEnrollSHRPLoaderMsg = "enroll_shrp_loader"
	TypeRevokeSHRPLoaderMsg = "revoke_shrp_loader"
	TypeBuySHRMsg           = "buy_shr"
	TypeMsgSetExchange      = "set_exchange"
	TypeMsgSendSHR          = "send_shr"
	TypeMsgSendSHRP         = "send_shrp"
)

type MsgLoadSHR struct {
	Approver sdk.AccAddress `json:"approver"`
	Receiver sdk.AccAddress `json:"receiver"`
	Amount   string         `json:"amount"`
}

func NewMsgLoadSHR(approver, receiver sdk.AccAddress, amount string) MsgLoadSHR {
	return MsgLoadSHR{
		Approver: approver,
		Receiver: receiver,
		Amount:   amount,
	}
}

func (msg MsgLoadSHR) Route() string {
	return RouterKey
}

func (msg MsgLoadSHR) Type() string {
	return TypeLoadSHRMsg
}

func (msg MsgLoadSHR) ValidateBasic() error {
	if msg.Approver.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Approver.String())
	}
	if msg.Receiver.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Receiver.String())
	}
	// if msg.Amount <= 0 {
	// 	return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "Amount must be positive")
	// }
	amount, ok := sdk.NewIntFromString(msg.Amount)
	if !ok || amount.LTE(sdk.ZeroInt()) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, fmt.Sprintf("Invalid amount number %s", msg.Amount))
	}
	return nil
}

func (msg MsgLoadSHR) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgLoadSHR) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Approver}
}

type MsgBuyCent struct {
	Buyer  sdk.AccAddress `json:"buyer"`
	Amount int            `json:"amount"`
}

func NewMsgBuyCent(buyer sdk.AccAddress, amount int) MsgBuyCent {
	return MsgBuyCent{
		Buyer:  buyer,
		Amount: amount,
	}
}

func (msg MsgBuyCent) Route() string {
	return RouterKey
}

func (msg MsgBuyCent) Type() string {
	return TypeBuyCent
}

func (msg MsgBuyCent) ValidateBasic() error {
	if msg.Buyer.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Buyer.String())
	}

	if msg.Amount <= 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "Amount must be positive")
	}
	return nil
}

func (msg MsgBuyCent) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgBuyCent) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Buyer}
}

type MsgBuySHR struct {
	Buyer  sdk.AccAddress `json:"buyer"`
	Amount string         `json:"amount"`
}

func NewMsgBuySHR(buyer sdk.AccAddress, amount string) MsgBuySHR {
	return MsgBuySHR{
		Buyer:  buyer,
		Amount: amount,
	}
}

func (msg MsgBuySHR) Route() string {
	return RouterKey
}

func (msg MsgBuySHR) Type() string {
	return TypeBuySHRMsg
}

func (msg MsgBuySHR) ValidateBasic() error {
	if msg.Buyer.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Buyer.String())
	}

	amount, ok := sdk.NewIntFromString(msg.Amount)
	if !ok || amount.LTE(sdk.ZeroInt()) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Amount must be positive")
	}
	return nil
}

func (msg MsgBuySHR) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgBuySHR) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Buyer}
}

type MsgBurnSHRP struct {
	Approver sdk.AccAddress `json:"approver"`
	Amount   string         `json:"amount"`
}

func NewMsgBurnSHRP(approver sdk.AccAddress, amt string) MsgBurnSHRP {
	return MsgBurnSHRP{
		Approver: approver,
		Amount:   amt,
	}
}

func (msg MsgBurnSHRP) Route() string {
	return RouterKey
}

func (msg MsgBurnSHRP) Type() string {
	return TypeBurnSHRPMsg
}

func (msg MsgBurnSHRP) ValidateBasic() error {
	if msg.Approver.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Approver.String())
	}

	// if _, _, err := ParseCoinStr(msg.Amount); err != nil {
	// 	return err
	// }
	amount, ok := sdk.NewIntFromString(msg.Amount)
	if !ok || amount.LTE(sdk.ZeroInt()) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Amount must be positive")
	}
	return nil
}

func (msg MsgBurnSHRP) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgBurnSHRP) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Approver}
}

type MsgBurnSHR struct {
	Approver sdk.AccAddress `json:"approver"`
	Amount   string         `json:"amount"`
}

func NewMsgBurnSHR(approver sdk.AccAddress, amt string) MsgBurnSHR {
	return MsgBurnSHR{
		Approver: approver,
		Amount:   amt,
	}
}

func (msg MsgBurnSHR) Route() string {
	return RouterKey
}

func (msg MsgBurnSHR) Type() string {
	return TypeBurnSHRMsg
}

func (msg MsgBurnSHR) ValidateBasic() error {
	if msg.Approver.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Approver.String())
	}
	amount, ok := sdk.NewIntFromString(msg.Amount)
	if !ok || amount.LTE(sdk.ZeroInt()) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Amount must be positive")
	}
	return nil
}

func (msg MsgBurnSHR) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgBurnSHR) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Approver}
}

type MsgLoadSHRP struct {
	Approver sdk.AccAddress `json:"approver"`
	Receiver sdk.AccAddress `json:"receiver"`
	Amount   string         `json:"amount"`
}

func NewMsgLoadSHRP(approver, receiver sdk.AccAddress, amount string) MsgLoadSHRP {
	return MsgLoadSHRP{
		Approver: approver,
		Receiver: receiver,
		Amount:   amount,
	}
}

func (msg MsgLoadSHRP) Route() string {
	return RouterKey
}

func (msg MsgLoadSHRP) Type() string {
	return TypeLoadSHRPMsg
}

func (msg MsgLoadSHRP) ValidateBasic() error {
	if msg.Approver.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Approver.String())
	}
	if msg.Receiver.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Receiver.String())
	}
	// if _, _, err := ParseCoinStr(msg.Amount); err != nil {
	// 	return err
	// }
	amount, ok := sdk.NewIntFromString(msg.Amount)
	if !ok || amount.LTE(sdk.ZeroInt()) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, fmt.Sprintf("Invalid amount number %s", msg.Amount))
	}
	return nil
}

func (msg MsgLoadSHRP) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgLoadSHRP) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Approver}
}

type MsgSendSHRP struct {
	Sender   sdk.AccAddress `json:"sender"`
	Receiver sdk.AccAddress `json:"receiver"`
	Amount   string         `json:"amount"`
}

func NewMsgSendSHRP(sender, receiver sdk.AccAddress, amount string) MsgSendSHRP {
	return MsgSendSHRP{
		Sender:   sender,
		Receiver: receiver,
		Amount:   amount,
	}
}

func (msg MsgSendSHRP) Route() string {
	return RouterKey
}

func (msg MsgSendSHRP) Type() string {
	return TypeMsgSendSHRP
}

func (msg MsgSendSHRP) ValidateBasic() error {
	if msg.Sender.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender.String())
	}
	if msg.Receiver.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Receiver.String())
	}
	amount, ok := sdk.NewIntFromString(msg.Amount)
	if !ok || amount.LTE(sdk.NewInt(0)) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Amount must be positive")
	}

	// if _, _, err := ParseCoinStr(msg.Amount); err != nil {
	// 	return err
	// }
	return nil
}

func (msg MsgSendSHRP) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgSendSHRP) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

type MsgSendSHR struct {
	Sender   sdk.AccAddress `json:"sender"`
	Receiver sdk.AccAddress `json:"receiver"`
	Amount   string         `json:"amount"`
}

func NewMsgSendSHR(sender, receiver sdk.AccAddress, amount string) MsgSendSHR {
	return MsgSendSHR{
		Sender:   sender,
		Receiver: receiver,
		Amount:   amount,
	}
}

func (msg MsgSendSHR) Route() string {
	return RouterKey
}

func (msg MsgSendSHR) Type() string {
	return TypeMsgSendSHR
}

func (msg MsgSendSHR) ValidateBasic() error {
	if msg.Sender.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender.String())
	}
	if msg.Receiver.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Receiver.String())
	}

	amount, ok := sdk.NewIntFromString(msg.Amount)
	if !ok || amount.LTE(sdk.NewInt(0)) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Amount must be positive")
	}
	// if !(msg.Amount > 0) {
	// 	return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Amount must be positive")
	// }
	return nil
}

func (msg MsgSendSHR) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgSendSHR) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

type MsgEnrollSHRPLoaders struct {
	Approver    sdk.AccAddress   `json:"approver"`
	SHRPLoaders []sdk.AccAddress `json:"shrploaders"`
}

func NewMsgEnrollSHRPLoaders(approver sdk.AccAddress, loaders []sdk.AccAddress) MsgEnrollSHRPLoaders {
	return MsgEnrollSHRPLoaders{
		Approver:    approver,
		SHRPLoaders: loaders,
	}
}

func (msg MsgEnrollSHRPLoaders) Route() string {
	return RouterKey
}

func (msg MsgEnrollSHRPLoaders) Type() string {
	return TypeEnrollSHRPLoaderMsg
}

func (msg MsgEnrollSHRPLoaders) ValidateBasic() error {
	if msg.Approver.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Approver.String())
	}
	if len(msg.SHRPLoaders) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "The signers list are empty")
	}
	for _, addr := range msg.SHRPLoaders {
		if addr.Empty() {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Signer address empty")
		}
	}
	return nil
}

func (msg MsgEnrollSHRPLoaders) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgEnrollSHRPLoaders) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Approver}
}

type MsgRevokeSHRPLoaders struct {
	Approver    sdk.AccAddress   `json:"approver"`
	SHRPLoaders []sdk.AccAddress `json:"receiver"`
}

func NewMsgRevokeSHRPLoaders(approver sdk.AccAddress, loaders []sdk.AccAddress) MsgRevokeSHRPLoaders {
	return MsgRevokeSHRPLoaders{
		Approver:    approver,
		SHRPLoaders: loaders,
	}
}

func (msg MsgRevokeSHRPLoaders) Route() string {
	return RouterKey
}

func (msg MsgRevokeSHRPLoaders) Type() string {
	return TypeRevokeSHRPLoaderMsg
}

func (msg MsgRevokeSHRPLoaders) ValidateBasic() error {
	if msg.Approver.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Approver.String())
	}
	if len(msg.SHRPLoaders) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "The signers list are empty")
	}
	for _, addr := range msg.SHRPLoaders {
		if addr.Empty() {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Signer address empty")
		}
	}
	return nil
}

func (msg MsgRevokeSHRPLoaders) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgRevokeSHRPLoaders) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Approver}
}

type MsgSetExchange struct {
	Approver sdk.AccAddress `json:"approver"`
	Rate     string         `json:"rate"`
}

func NewMsgSetExchange(approver sdk.AccAddress, rate string) MsgSetExchange {
	return MsgSetExchange{
		Approver: approver,
		Rate:     rate,
	}
}

func (msg MsgSetExchange) Route() string {
	return RouterKey
}

func (msg MsgSetExchange) Type() string {
	return TypeMsgSetExchange
}

func (msg MsgSetExchange) ValidateBasic() error {
	if msg.Approver.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Approver.String())
	}
	rate, ok := sdk.NewIntFromString(msg.Rate)
	if !ok || rate.LTE(sdk.NewInt(0)) {
		return ErrInvalidExchangeRate
	}
	return nil
}

func (msg MsgSetExchange) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgSetExchange) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Approver}
}

func ParseCoinStr(s string) (i, d int64, err error) {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return i, d, err
	}
	if f < 0 {
		return i, d, sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "Negative Coins are not accepted")
	}
	i = int64(f)

	d = int64(f*100 - float64(i*100))
	return
}

func ParseCoinFloat(f float64) (i, d int64, err error) {

	if f < 0 {
		return i, d, sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "Negative Coins are not accepted")
	}
	i = int64(f)
	d = int64(f*100 - float64(i*100) + 1) // make sure always round it up
	return
}
