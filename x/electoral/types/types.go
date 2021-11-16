package types

const (
	StatusVoterEnrolled = "enrolled"
	StatusVoterRevoked  = "revoked"
	DefaultVoterStatus  = StatusVoterRevoked
)

func NewVoter() Voter {
	return Voter{
		Status: DefaultVoterStatus,
	}
}
