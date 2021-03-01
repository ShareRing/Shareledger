package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	StatusVoterEnrolled = "enrolled"
	StatusVoterRevoked  = "revoked"
	DefaultVoterStatus  = StatusVoterRevoked
)

type Voter struct {
	Address sdk.AccAddress `json:"address"`
	Status  string         `json:"status"`
}

func NewVoter() Voter {
	return Voter{
		Status: DefaultVoterStatus,
	}
}

func (v Voter) String() string {
	return fmt.Sprintf("{Voter: %s, Status: %s}", v.Address.String(), v.Status)
}
