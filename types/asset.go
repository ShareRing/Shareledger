package types

import (
	"encoding/hex"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strconv"
)

// Asset asset infomation
type Asset struct {
	UUID    string      `json:"uuid"`
	Hash    []byte      `json:"hash"`
	Creator sdk.Address `json:"creator"`
	Status  bool        `json:"status"`
	Fee     int64       `json:"fee"`
}

func (a Asset) String() string {
	return fmt.Sprintf("Asset{UUID:%s, Hash:%s, Creator:%s, Status:%s, Fee:%d}",
		                a.UUID,
		                hex.EncodeToString(a.Hash),
		                a.Creator.String(),
						strconv.FormatBool(a.Status),
						a.Fee)
}

//--------------------------------------------------------
//--------------------------------------------------------

//Asset defines basic information of Assets in ShareRing Platform
func NewAsset(uuid string, creator sdk.Address, hash []byte, status bool, fee int64) Asset {
	return Asset{
		UUID:    uuid,
		Creator: creator,
		Hash:    hash,
		Status:  status,
		Fee:     fee,
	}
}
