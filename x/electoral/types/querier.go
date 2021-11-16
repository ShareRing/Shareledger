package types

func NewQueryVoterRequest(voter string) *QueryVoterRequest {
	return &QueryVoterRequest{Address: voter}
}
