package types

import sdk "github.com/cosmos/cosmos-sdk/types"

// query endpoints supported by the Id Querier
const (
	QueryInfo = "info"
)

func NewQueryIdByAddressParams(ownerAddr sdk.AccAddress) *QueryIdByAddressRequest {
	return &QueryIdByAddressRequest{Address: ownerAddr.String()}
}

func NewQueryIdByIdParams(id string) *QueryIdByIdRequest {
	return &QueryIdByIdRequest{Id: id}
}
