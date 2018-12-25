package types

import (
	// "encoding/hex"
    "encoding/json"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	// "strconv"
)

// Asset asset infomation
type Asset struct {
	UUID    string      `json:"uuid"`
	Hash    []byte      `json:"hash"`
	Creator sdk.AccAddress `json:"creator"`
	Status  bool        `json:"status"`
	Fee     int64       `json:"fee"`
}

func (a Asset) String() string {
    b, err := json.Marshal(a)
    if err != nil {
        panic(err)
    }
    return fmt.Sprintf("%s", b)
}

//--------------------------------------------------------
//--------------------------------------------------------

//Asset defines basic information of Assets in ShareRing Platform
func NewAsset(uuid string, creator sdk.AccAddress, hash []byte, status bool, fee int64) Asset {
	return Asset{
		UUID:    uuid,
		Creator: creator,
		Hash:    hash,
		Status:  status,
		Fee:     fee,
	}
}
