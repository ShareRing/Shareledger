package messages

import (
	"encoding/json"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sharering/shareledger/constants"
	"github.com/sharering/shareledger/types"
)

var _ sdk.Msg = MsgUpdate{}

type MsgUpdate struct {
	FromDenom string    `json:"from_denom"`
	ToDenom   string    `json:"to_denom"`
	Rate      types.Dec `json:"rate"`
}

func NewMsgUpdate(
	from string,
	to string,
	rate types.Dec,
) MsgUpdate {
	return MsgUpdate{
		FromDenom: from,
		ToDenom:   to,
		Rate:      rate,
	}
}

// Type type of this message
func (msg MsgUpdate) Type() string {
	return constants.MESSAGE_EXCHANGE_RATE
}

func (msg MsgUpdate) Route() string { return constants.MESSAGE_EXCHANGE_RATE }

func (msg MsgUpdate) ValidateBasic() sdk.Error {
	if msg.FromDenom == msg.ToDenom {
		return sdk.ErrInternal(fmt.Sprintf(constants.EXC_SAME_DENOM, msg.FromDenom))
	}

	if !types.IsValidDenom(msg.FromDenom) || !types.IsValidDenom(msg.ToDenom) {
		return sdk.ErrInternal(fmt.Sprintf(constants.EXC_INVALID_DENOM,
			strings.Join(constants.ALL_DENOMS, ","),
			strings.Join([]string{msg.FromDenom, msg.ToDenom}, ",")))
	}

	if msg.Rate.IsZero() {
		return sdk.ErrInternal(fmt.Sprintf(constants.EXC_INVALID_RATE, msg.Rate.String()))
	}

	return nil
}

func (msg MsgUpdate) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return b
}

func (msg MsgUpdate) String() string {
	return fmt.Sprintf("ExchangeRate/MsgUpdate{%s}", msg.GetSignBytes())
}

func (msg MsgUpdate) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{}
}
