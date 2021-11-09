package types

// query endpoints supported by the Id Querier
const (
	QueryInfo = "info"
)

// QueryIdByAddressParams defines the params for querying an account id information.
type QueryDocByProofParams struct {
	Proof string
}

func NewQueryDocByProofParams(proof string) QueryDocByProofParams {
	return QueryDocByProofParams{Proof: proof}
}

// Query doc by holder id
type QueryDocByHolderParams struct {
	Holder string
}

func NewQueryDocByHolderParams(holderId string) QueryDocByHolderParams {
	return QueryDocByHolderParams{Holder: holderId}
}
