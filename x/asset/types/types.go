package types

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Asset struct {
	UUID    string         `json:"uuid"`
	Hash    []byte         `json:"hash"`
	Creator sdk.AccAddress `json:"creator"`
	Status  bool           `json:"status"`
	Rate    int64          `json:""`
}

func NewAsset() Asset {
	return Asset{}
}

func NewAssetFromMsgCreate(msg MsgCreate) Asset {
	asset := NewAsset()
	asset.Creator = msg.Creator
	asset.Rate = msg.Rate
	asset.Hash = msg.Hash
	asset.Status = msg.Status
	asset.UUID = msg.UUID
	return asset
}

func NewAssetFromMsgUpdate(msg MsgUpdate) Asset {
	asset := NewAsset()
	asset.Creator = msg.Creator
	asset.Rate = msg.Rate
	asset.Hash = msg.Hash
	asset.Status = msg.Status
	asset.UUID = msg.UUID
	return asset
}

func (a Asset) GetString() (string, error) {
	js, err := json.Marshal(a)
	if err != nil {
		return "", err
	}
	return string(js), nil
}
