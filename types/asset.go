package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"fmt"
	"encoding/hex"
)

// Asset asset infomation
type Asset struct {
	UUID    string       `json:"uuid"`
	Hash    []byte `json:"hash"`
	Creator sdk.Address  `json:"creator"`
}


func (a Asset) String() string {
	return fmt.Sprintf("Asset{UUID:%s, Hash:%s, Creator:%s}", a.UUID, hex.EncodeToString(a.Hash), a.Creator.String())
}

//--------------------------------------------------------
//--------------------------------------------------------

//Asset defines basic information of Assets in ShareRing Platform
func NewAsset(uuid string, creator sdk.Address, hash []byte) Asset {
	return Asset{
		UUID:    uuid,
		Creator: creator,
		Hash: hash,
	}
}
