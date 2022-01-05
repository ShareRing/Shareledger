package types

import "encoding/json"

func NewAssetFromMsgCreate(msg MsgCreateAsset) Asset {
	asset := Asset{}
	asset.Creator = msg.Creator
	asset.Rate = msg.Rate
	asset.Hash = msg.Hash
	asset.Status = msg.Status
	asset.UUID = msg.UUID
	return asset
}

func NewAssetFromMsgUpdate(msg MsgUpdateAsset) Asset {
	asset := Asset{}
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
