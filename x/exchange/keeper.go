package exchange

import (
	"encoding/json"
	"fmt"

	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
	"bitbucket.org/shareringvn/cosmos-sdk/wire"

	"github.com/sharering/shareledger/x/exchange/types"
)

// Keeper to store ExchangeRate
type Keeper struct {
	storeKey sdk.StoreKey // key used to access the store from Context
	cdc      *wire.Codec
}

// NewKeeper - Return a new keeper
func NewKeeper(key sdk.StoreKey, cdc *wire.Codec) Keeper {
	return Keeper{
		storeKey: key,
		cdc:      cdc,
	}
}

//-----------------------------------------------------------

// GetStoreKey return keys to be used as key for store
func GetStoreKey(fromDemon string, toDenom string) []byte {
	return append([]byte(fromDenom), []byte(toDenom)...)
}

// StoreExchangeRate - store exchangeRate
func (k Keeper) Store(ctx sdk.Context, e types.ExchangeRate) error {
	store := ctx.KVStore(k.storeKey)

	key := GetStoreKey(e.FromDenom, e.ToDenom)

	jsonBytes, err := json.Marshal(e)

	if err != nil {
		return sdk.ErrInternal(fmt.Sprintf(constants.EXC_JSON_MARSHAL, err.Error()))
	}

	store.Set(key, jsonBytes)
	return nil
}

// GetExchangeRate - get exchangeRate from a store
func (k Keeper) Get(
	ctx sdk.Context,
	fromDenom string,
	toDenom string,
) (e types.ExchangeRate, err error) {

	store := ctx.KVStore(k.storeKey)

	key := GetStoreKey(fromDenom, toDenom)

	jsonBytes := store.Get(key)

	if jsonBytes == nil {
		return e, sdk.ErrInternal(fmt.Sprintf(constants.EXCHANGE_RATE_NOT_FOUND, fromDenom, toDenom))
	}

	err := json.Unmarshal(jsonBytes, &e)
	if err != nil {
		return e, sdk.ErrInternal(fmt.Sprintf(constants.EXC_JSON_MARSHAL, err.Error()))
	}

	return e, nil
}

func (k Keeper) Delete(
	ctx sdk.Context,
	fromDenom string,
	toDenom string,
) (e types.ExchangeRate, err error) {
	e, err = Get(ctx, fromDenom, toDenom)
	if err != nil {
		return e, err
	}

	store := ctx.KVStore(k.storeKey)

	store.Delete(GetStoreKey(fromDenom, toDenom))

	return e, nil
}

//-----------------------------------------------------------
// 		API Create, Retrieve, Update, Deletion

func (k Keeper) CreateExchangeRate(
	ctx sdk.Context,
	msg msg.MsgCreate,
) (ex types.ExchangeRate, err error) {

	ex = types.NewExchangeRate(msg.FromDenom, msg.ToDenom, msg.Rate)

	err = k.Store(ctx, ex)

	return ex, err
}

func (k Keeper) RetrieveExchangeRate(
	ctx sdk.Context,
	msg msg.MsgRetrieve,
) (ex types.ExchangeRate, err error) {
	return k.Get(ctx, msg.FromDenom, msg.ToDenom)
}

func (k Keeper) UpdateExchangeRate(
	ctx sdk.Context,
	msg msg.MsgUpdate,
) (ex types.ExchangeRate, err error) {
	ex = types.NewExchangeRate(msg.FromDenom, msg.ToDenom, msg.Rate)

	err = k.Store(ctx, ex)

	return ex, err
}

func (k Keeper) DeleteExchangeRate(
	ctx sdk.Context,
	msg msg.MsgDelete,
) (ex types.ExchangeRate, err error) {
	return k.Delete(ctx)
}
