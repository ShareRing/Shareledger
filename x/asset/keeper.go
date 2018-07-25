package asset

import (
	"fmt"
	"errors"

	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
	"bitbucket.org/shareringvn/cosmos-sdk/wire"

	"github.com/sharering/shareledger/types"
	msg "github.com/sharering/shareledger/x/asset/messages"
	"encoding/json"
)

// Keeper data type
type Keeper struct {
	storeKey sdk.StoreKey // key used to access the store from the Context.
	cdc      *wire.Codec
}

// NewKeeper - Returns the Keeper
func NewKeeper(key sdk.StoreKey, cdc *wire.Codec) Keeper {
	return Keeper{
		storeKey: key,
		cdc:      cdc,
	}
}

//----------------------------------------------------------
func (k Keeper) CreateAsset(ctx sdk.Context, msg msg.MsgCreate) (types.Asset, error) {
	fmt.Println(ctx) //fake action for testing

	store := ctx.KVStore(k.storeKey)

	asset := types.NewAsset(msg.UUID, msg.Creator, msg.Hash, msg.Status, msg.Fee)


	assetBytes, err := json.Marshal(asset)

	if err != nil {
		return types.Asset{}, errors.New("Asset Encoding Error")

	}

	// Store to KVStore
	store.Set([]byte(msg.UUID), assetBytes)

	return asset, nil
}

func (k Keeper) RetrieveAsset(ctx sdk.Context, msg msg.MsgRetrieve) (types.Asset, error) {

	store := ctx.KVStore(k.storeKey)

	assetBytes := store.Get([]byte(msg.UUID))

	if assetBytes == nil {
		return types.Asset{}, errors.New("Asset is not found")
	}

	var asset types.Asset

	derr := json.Unmarshal(assetBytes, &asset)
	if derr != nil {
		return types.Asset{}, errors.New("Asset decoding error")
	}


	return asset, nil
}


func (k Keeper) UpdateAsset(ctx sdk.Context, msg msg.MsgUpdate) (types.Asset, error) {

	store := ctx.KVStore(k.storeKey)

	assetBytes := store.Get([]byte(msg.UUID))

	if assetBytes == nil {
		return types.Asset{}, errors.New("Asset is not found")
	}

	var asset types.Asset

	derr := json.Unmarshal(assetBytes, &asset)
	if derr != nil {
		return types.Asset{}, errors.New("Asset decoding error")
	}


	asset = types.NewAsset(msg.UUID, msg.Creator, msg.Hash, msg.Status, msg.Fee)

	nassetBytes, err := json.Marshal(asset)

	if err != nil {
		return types.NewAsset("", []byte(""), []byte(""), true, 0), errors.New("Asset Encoding Error")

	}

	store.Set([]byte(msg.UUID), nassetBytes)

	return asset, nil
}

func (k Keeper) DeleteAsset(ctx sdk.Context, msg msg.MsgDelete) (types.Asset, error) {

	store := ctx.KVStore(k.storeKey)

	assetBytes := store.Get([]byte(msg.UUID))

	if assetBytes == nil {
		return types.NewAsset("", []byte(""), []byte(""), true, 0), errors.New("Asset is not found")
	}

	// Unmarshall asset
	var asset types.Asset

	derr := json.Unmarshal(assetBytes, &asset)

	if derr != nil {
		return types.NewAsset("", []byte(""), []byte(""), true, 0), errors.New("Asset decoding error")
	}

	// Delete asset
	store.Delete([]byte(msg.UUID))

	return asset, nil
}
