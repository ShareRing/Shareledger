// this is sorely use for cw-721 contract
// and for testing purpose only
package distributionx

import "encoding/json"

type InstantiateMsg struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
	Minter string `json:"minter"`
}

func (m *InstantiateMsg) MustToJSON() string {
	bz, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	return string(bz)
}

type MintMsg struct {
	Owner   string `json:"owner"`
	TokenID string `json:"token_id"`
}

func (m *MintMsg) MustToJSON() string {
	bz, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	return string(bz)
}
