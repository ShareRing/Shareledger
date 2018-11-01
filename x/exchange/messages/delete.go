package messages

import (
	"encoding/json"
	"fmt"
	"strings"

	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"

	"github.com/sharering/shareledger/constants"
	"github.com/sharering/shareledger/types"
)

type MsgDelete struct {
	FromDenom string `json:"from_denom"`
	ToDenom   string `json:"to_denom"`
}

var _ sdk.Msg = MsgDelete{}

func NewMsgDelete(
	from string,
	to string,
) MsgDelete {
	return MsgDelete{
		FromDenom: from,
		ToDenom:   to,
	}
}

// Type type of this message
func (msg MsgDelete) Type() string {
	return constants.MESSAGE_EXCHANGE_RATE
}

func (msg MsgDelete) ValidateBasic() sdk.Error {
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

func (msg MsgDelete) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return b
}

func (msg MsgDelete) String() string {
	return fmt.Sprintf("ExchangeRate/MsgDelete{%s}", msg.GetSignBytes())
}

func (msg MsgDelete) GetSigners() []sdk.Address {
	return []sdk.Address{}
}

func (msg MsgDelete) Tags() sdk.Tags {
	return sdk.NewTags("msg.module", []byte("exchangerate")).
		AppendTag("fromDenom", []byte(msg.FromDenom)).
		AppendTag("toDenom", []byte(msg.ToDenom))
}
