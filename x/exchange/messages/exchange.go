package messages

import (
	"encoding/json"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sharering/shareledger/constants"
	"github.com/sharering/shareledger/types"
	etypes "github.com/sharering/shareledger/x/exchange/types"
)

type MsgExchange struct {
	FromDenom string         `json:"from_denom"`
	ToDenom   string         `json:"to_denom"`
	Amount    types.Dec      `json:"amount"`
	Reserve   sdk.AccAddress `json:"reserve"`
}

var _ sdk.Msg = MsgExchange{}

func NewMsgExchange(
	from string,
	to string,
	amount types.Dec,
	reserve sdk.AccAddress,
) MsgExchange {
	return MsgExchange{
		FromDenom: from,
		ToDenom:   to,
		Amount:    amount,
		Reserve:   reserve,
	}
}

// Type type of this message
func (msg MsgExchange) Type() string {
	return constants.MESSAGE_EXCHANGE_RATE
}

func (msg MsgExchange) Route() string { return constants.MESSAGE_EXCHANGE_RATE }

func (msg MsgExchange) ValidateBasic() sdk.Error {
	if msg.FromDenom == msg.ToDenom {
		return sdk.ErrInternal(fmt.Sprintf(constants.EXC_SAME_DENOM, msg.FromDenom))
	}

	if !types.IsValidDenom(msg.FromDenom) || !types.IsValidDenom(msg.ToDenom) {
		return sdk.ErrInternal(fmt.Sprintf(constants.EXC_INVALID_DENOM,
			strings.Join(constants.ALL_DENOMS, ","),
			strings.Join([]string{msg.FromDenom, msg.ToDenom}, ",")))
	}

	if msg.Amount.IsZero() {
		return sdk.ErrInternal(fmt.Sprintf(constants.EXC_INVALID_AMOUNT, msg.Amount.String()))
	}

	reserve := etypes.NewReserve(msg.Reserve)

	if !reserve.IsValid() {
		return sdk.ErrInternal(fmt.Sprintf("constants.EXC_INVALID_RESERVE, msg.Reserve.String()"))
	}

	return nil
}

func (msg MsgExchange) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return b
}

func (msg MsgExchange) String() string {
	return fmt.Sprintf("ExchangeRate/MsgExchange{%s}", msg.GetSignBytes())
}

func (msg MsgExchange) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{}
}

func (msg MsgExchange) Tags() sdk.Tags {
	return sdk.NewTags("msg.module", "exchangerate").
		AppendTag("fromDenom", msg.FromDenom).
		AppendTag("toDenom", msg.ToDenom).
		AppendTag("amount", msg.Amount.String())
}
