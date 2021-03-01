package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	StatusSHRPLoaderActived   = "actived"
	StatusSHRPLoaderInactived = "inactived"
	defaultSHRPLoaderStatus   = StatusSHRPLoaderInactived
	SHRP                      = "shrp"
	CENT                      = "cent"
	SHR                       = "shr"
)

type SHRPLoader struct {
	Status string `json:"status"`
}

func NewSHRPLoader() SHRPLoader {
	return SHRPLoader{
		Status: defaultSHRPLoaderStatus,
	}
}

func (sl SHRPLoader) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Status %s`, sl.Status))
}

// Status
const (
	Active   = "active"
	Inactive = "inactive"
)

type AccState struct {
	Address sdk.AccAddress `json:"address"`
	Status  string         `json:"status"`
}

func NewAccState(addr sdk.AccAddress, status string) AccState {
	if status != Active && status != Inactive {
		panic("Status is wrong")
	}
	return AccState{
		Address: addr,
		Status:  status,
	}
}

func (ids AccState) String() string {
	return strings.TrimSpace(fmt.Sprintf(`%s:%s`, ids.Address.String(), ids.Status))
}

func (ids *AccState) Activate() {
	ids.Status = Active
}

func (ids AccState) IsEmpty() bool {
	if len(ids.Address.Bytes()) == 0 || len(ids.Status) == 0 {
		return true
	}
	return false
}

func (ids AccState) IsActive() bool {
	if ids.IsEmpty() {
		return false
	}
	return ids.Status == Active
}
