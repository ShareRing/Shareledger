package exchange

import (
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sharering/shareledger/constants"
	"github.com/sharering/shareledger/types"
	"github.com/sharering/shareledger/x/bank"
	"github.com/sharering/shareledger/x/exchange/messages"
	etypes "github.com/sharering/shareledger/x/exchange/types"
)

// Keeper to store ExchangeRate
type Keeper struct {
	storeKey   sdk.StoreKey // key used to access the store from Context
	bankKeeper bank.Keeper  // bank keeper to swap tokens
}

// NewKeeper - Return a new keeper
func NewKeeper(key sdk.StoreKey, bk bank.Keeper) Keeper {
	return Keeper{
		storeKey:   key,
		bankKeeper: bk,
	}
}

//-----------------------------------------------------------

// GetStoreKey return keys to be used as key for store
func GetStoreKey(fromDenom string, toDenom string) []byte {
	return append([]byte(fromDenom), []byte(toDenom)...)
}

// StoreExchangeRate - store exchangeRate
func (k Keeper) Store(ctx sdk.Context, e etypes.ExchangeRate) error {
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
) (e etypes.ExchangeRate, err error) {

	store := ctx.KVStore(k.storeKey)

	key := GetStoreKey(fromDenom, toDenom)

	jsonBytes := store.Get(key)

	if jsonBytes == nil {
		return e, sdk.ErrInternal(fmt.Sprintf(constants.EXC_EXCHANGE_RATE_NOT_FOUND, fromDenom, toDenom))
	}

	err = json.Unmarshal(jsonBytes, &e)
	if err != nil {
		return e, sdk.ErrInternal(fmt.Sprintf(constants.EXC_JSON_MARSHAL, err.Error()))
	}

	return e, nil
}

func (k Keeper) Delete(
	ctx sdk.Context,
	fromDenom string,
	toDenom string,
) (e etypes.ExchangeRate, err error) {
	e, err = k.Get(ctx, fromDenom, toDenom)
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
	msg messages.MsgCreate,
) (ex etypes.ExchangeRate, err error) {

	_, err = k.Get(ctx, msg.FromDenom, msg.ToDenom)

	// Already exist an exchange
	if err == nil {
		return ex, fmt.Errorf(constants.EXC_ALREADY_EXIST, msg.FromDenom, msg.ToDenom)
	}

	ex = etypes.NewExchangeRate(msg.FromDenom, msg.ToDenom, msg.Rate)

	err = k.Store(ctx, ex)

	return ex, err
}

func (k Keeper) RetrieveExchangeRate(
	ctx sdk.Context,
	fromDenom string,
	toDenom string,
) (ex etypes.ExchangeRate, err error) {
	return k.Get(ctx, fromDenom, toDenom)
}

func (k Keeper) UpdateExchangeRate(
	ctx sdk.Context,
	msg messages.MsgUpdate,
) (ex etypes.ExchangeRate, err error) {
	ex = etypes.NewExchangeRate(msg.FromDenom, msg.ToDenom, msg.Rate)

	err = k.Store(ctx, ex)

	return ex, err
}

func (k Keeper) DeleteExchangeRate(
	ctx sdk.Context,
	msg messages.MsgDelete,
) (ex etypes.ExchangeRate, err error) {
	return k.Delete(ctx, msg.FromDenom, msg.ToDenom)
}

func (k Keeper) SellCoin(
	ctx sdk.Context,
	account sdk.AccAddress,
	reserveAddress sdk.AccAddress,
	fromDenom string,
	toDenom string,
	sellingAmount types.Dec,
) (err error) {
	exr, err := k.RetrieveExchangeRate(ctx, fromDenom, toDenom)

	if err != nil {
		return err
	}
	// Get balance
	fromAcc := k.bankKeeper.GetCoins(ctx, account)

	// Already check validity with the message
	reserve := etypes.NewReserve(reserveAddress)

	reserveAcc := reserve.GetCoins(ctx, k.bankKeeper)

	sellingCoin := types.NewCoinFromDec(fromDenom, sellingAmount)

	buyingCoin := exr.Convert(sellingCoin)

	if fromAcc.LT(sellingCoin) || reserveAcc.LT(buyingCoin) {
		return fmt.Errorf(constants.EXC_INSUFFICIENT_BALANCE,
			fromAcc.String(),
			sellingCoin.String(),
			reserveAcc.String(),
			buyingCoin.String())
	}

	// Transfer selling currencies from FromAcc to ReserveAcc
	newFromAcc := fromAcc.Minus(sellingCoin)
	newReserveAcc := reserveAcc.Plus(sellingCoin)

	// Transfer Buying currencies from ReserveAcc to FromAcc
	newReserveAcc = newReserveAcc.Minus(buyingCoin)
	newFromAcc = newFromAcc.Plus(buyingCoin)

	// Save to store

	sdkErr := k.bankKeeper.SetCoins(ctx, account, newFromAcc)

	if sdkErr != nil {
		return fmt.Errorf(sdkErr.Error())
	}

	sdkErr = reserve.SetCoins(ctx, k.bankKeeper, newReserveAcc)

	if sdkErr != nil {
		return fmt.Errorf(sdkErr.Error())
	}

	return nil
}

func (k Keeper) BuyCoin(
	ctx sdk.Context,
	account sdk.AccAddress,
	reserveAddress sdk.AccAddress,
	fromDenom string,
	toDenom string,
	buyingAmount types.Dec,
) (err error) {
	// fmt.Printf("BUY COIN\n")
	exr, err := k.RetrieveExchangeRate(ctx, fromDenom, toDenom)

	if err != nil {
		return err
	}
	// Get balance
	fromAcc := k.bankKeeper.GetCoins(ctx, account)

	// Already check validity with the message
	reserve := etypes.NewReserve(reserveAddress)

	reserveAcc := reserve.GetCoins(ctx, k.bankKeeper)

	buyingCoin := types.NewCoinFromDec(toDenom, buyingAmount)

	sellingCoin := exr.Obtain(buyingCoin)

	if fromAcc.LT(sellingCoin) || reserveAcc.LT(buyingCoin) {
		return fmt.Errorf(constants.EXC_INSUFFICIENT_BALANCE,
			fromAcc.String(),
			sellingCoin.String(),
			reserveAcc.String(),
			buyingCoin.String())
	}

	// Transfer selling currencies from FromAcc to ReserveAcc
	newFromAcc := fromAcc.Minus(sellingCoin)
	newReserveAcc := reserveAcc.Plus(sellingCoin)

	// Transfer Buying currencies from ReserveAcc to FromAcc
	newReserveAcc = newReserveAcc.Minus(buyingCoin)
	newFromAcc = newFromAcc.Plus(buyingCoin)

	// Save to store

	sdkErr := k.bankKeeper.SetCoins(ctx, account, newFromAcc)

	if sdkErr != nil {
		return fmt.Errorf(sdkErr.Error())
	}

	sdkErr = reserve.SetCoins(ctx, k.bankKeeper, newReserveAcc)

	if sdkErr != nil {
		return fmt.Errorf(sdkErr.Error())
	}

	return nil

}

