package messages

import (
	"encoding/json"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sharering/shareledger/constants"
	"github.com/sharering/shareledger/types"
)

type MsgRetrieve struct {
	FromDenom string `json:"from_denom"`
	ToDenom   string `json:"to_denom"`
}

var _ sdk.Msg = MsgRetrieve{}

func NewMsgRetrieve(
	from string,
	to string,
) MsgRetrieve {
	return MsgRetrieve{
		FromDenom: from,
		ToDenom:   to,
	}
}

// Type type of this message
func (msg MsgRetrieve) Type() string {
	return constants.MESSAGE_EXCHANGE_RATE
}

func (msg MsgRetrieve) Route() string { return constants.MESSAGE_EXCHANGE_RATE }

func (msg MsgRetrieve) ValidateBasic() sdk.Error {
	if msg.FromDenom == msg.ToDenom {
		return sdk.ErrInternal(fmt.Sprintf(constants.EXC_SAME_DENOM, msg.FromDenom))
	}

	if !types.IsValidDenom(msg.FromDenom) || !types.IsValidDenom(msg.ToDenom) {
		return sdk.ErrInternal(fmt.Sprintf(constants.EXC_INVALID_DENOM,
			strings.Join(constants.ALL_DENOMS, ","),
			strings.Join([]string{msg.FromDenom, msg.ToDenom}, ",")))
	}

	return nil
}

func (msg MsgRetrieve) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return b
}

func (msg MsgRetrieve) String() string {
	return fmt.Sprintf("ExchangeRate/MsgRetrieve{%s}", msg.GetSignBytes())
}

func (msg MsgRetrieve) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{}
}

func (msg MsgRetrieve) Tags() sdk.Tags {
	return sdk.NewTags("msg.module", "exchangerate").
		AppendTag("fromDenom", msg.FromDenom).
		AppendTag("toDenom", msg.ToDenom)
}
