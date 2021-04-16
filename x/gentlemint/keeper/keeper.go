package keeper

import (
	"fmt"

	myutil "github.com/ShareRing/modules/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/x/gentlemint/types"

	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/supply"
)

const (
	ExchangeRateKey = "exchange_shrp_to_shr"

	ShrpToCentRate = 100
	AuthorityKey   = "A"
	TreasurerKey   = "T"
	IdSignerKey    = "IDS"
	DocIssuerKey   = "DOCIS"
	AccOpKey       = "ACCOP"
)

var (
	DefaultExchangeRate = sdk.NewInt(200)
	MaxSHRSupply        = sdk.NewInt(4396000000).Mul(myutil.SHRDecimal)
)

type Keeper struct {
	accountKeeper auth.AccountKeeper
	supplyKeeper  supply.Keeper
	bankKeeper    bank.Keeper
	storeKey      sdk.StoreKey
	cdc           *codec.Codec
}

// NewKeeper creates new instances of the gentlemint Keeper
func NewKeeper(cdc *codec.Codec, storeKey sdk.StoreKey, ak auth.AccountKeeper, sk supply.Keeper, bk bank.Keeper) Keeper {
	return Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		accountKeeper: ak,
		supplyKeeper:  sk,
		bankKeeper:    bk,
	}
}

func (keeper Keeper) LoadCoins(ctx sdk.Context, toAddr sdk.AccAddress, amt sdk.Coins) error {
	if !amt.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, amt.String())
	}
	_, err := keeper.bankKeeper.AddCoins(ctx, toAddr, amt)
	if err != nil {
		return err
	}

	return keeper.supplyKeeper.MintCoins(ctx, types.ModuleName, amt)
}

func (keeper Keeper) BurnCoins(ctx sdk.Context, addr sdk.AccAddress, amt sdk.Coins) error {
	if !amt.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, amt.String())
	}
	_, err := keeper.bankKeeper.SubtractCoins(ctx, addr, amt)
	if err != nil {
		return err
	}
	return keeper.supplyKeeper.BurnCoins(ctx, types.ModuleName, amt)
}

func (keeper Keeper) GetCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins {
	acc := keeper.accountKeeper.GetAccount(ctx, addr)
	if acc != nil {
		return acc.GetCoins()
	}
	return sdk.NewCoins()
}

func (keeper Keeper) AddCoins(ctx sdk.Context, toAddr sdk.AccAddress, amt sdk.Coins) (sdk.Coins, error) {
	return keeper.bankKeeper.AddCoins(ctx, toAddr, amt)
}

func (keeper Keeper) SubtractCoins(ctx sdk.Context, addr sdk.AccAddress, amt sdk.Coins) (sdk.Coins, error) {
	return keeper.bankKeeper.SubtractCoins(ctx, addr, amt)
}

func (keeper Keeper) SendCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error {
	return keeper.bankKeeper.SendCoins(ctx, fromAddr, toAddr, amt)
}

func (keeper Keeper) SupplyMintCoins(ctx sdk.Context, amt sdk.Coins) error {
	return keeper.supplyKeeper.MintCoins(ctx, types.ModuleName, amt)
}

func (keeper Keeper) SupplyBurnCoins(ctx sdk.Context, amt sdk.Coins) error {
	return keeper.supplyKeeper.BurnCoins(ctx, types.ModuleName, amt)
}

func (k Keeper) GetSHRPLoader(ctx sdk.Context, address string) types.SHRPLoader {
	if !k.IsSHRPLoaderPresent(ctx, address) {
		return types.NewSHRPLoader()
	}
	store := ctx.KVStore(k.storeKey)
	bz := store.Get([]byte(address))
	var loader types.SHRPLoader
	k.cdc.MustUnmarshalBinaryBare(bz, &loader)
	return loader
}

func (k Keeper) SetSHRPLoader(ctx sdk.Context, address string, loader types.SHRPLoader) {
	if loader.Status == "" {
		return
	}
	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(address), k.cdc.MustMarshalBinaryBare(loader))
}

func (k Keeper) GetSHRPLoaderStatus(ctx sdk.Context, address string) string {
	loader := k.GetSHRPLoader(ctx, address)
	return loader.Status
}

func (k Keeper) SetSHRPLoaderStatus(ctx sdk.Context, address string, status string) {
	loader := k.GetSHRPLoader(ctx, address)
	loader.Status = status
	k.SetSHRPLoader(ctx, address, loader)
}

func (k Keeper) DeleteSHRPLoader(ctx sdk.Context, address string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(address))
}

func (k Keeper) GetSHRPLoadersIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, []byte("shrploader"))
}

func (k Keeper) IterateSHRPLoaders(ctx sdk.Context, cb func(loaderKey string, loader types.SHRPLoader) (stop bool)) {
	iterator := k.GetSHRPLoadersIterator(ctx)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var loader types.SHRPLoader
		loaderKey := string(iterator.Key())
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &loader)
		if cb(loaderKey, loader) {
			break
		}
	}
}

func (k Keeper) IsSHRPLoaderPresent(ctx sdk.Context, address string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(address))
}

func (k Keeper) ShrMintPossible(ctx sdk.Context, amt sdk.Int) bool {
	total := k.supplyKeeper.GetSupply(ctx).GetTotal()
	newAmt := total.AmountOf("shr").Add(amt)
	if newAmt.GTE(MaxSHRSupply) {
		return false
	}
	return true
}

func (k Keeper) GetExchangeRate(ctx sdk.Context) sdk.Int {
	if !k.IsExchangeRatePresent(ctx) {
		return DefaultExchangeRate
	}
	store := ctx.KVStore((k.storeKey))
	bz := store.Get([]byte(ExchangeRateKey))
	var rate sdk.Int
	k.cdc.MustUnmarshalBinaryBare(bz, &rate)
	return rate
}

func (k Keeper) SetExchangeRate(ctx sdk.Context, rate string) {
	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(ExchangeRateKey), k.cdc.MustMarshalBinaryBare(rate))
}

func (k Keeper) SetAuthorityAccount(ctx sdk.Context, authority string) {
	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(AuthorityKey), k.cdc.MustMarshalBinaryBare(authority))
}
func (k Keeper) GetAuthorityAccount(ctx sdk.Context) string {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get([]byte(AuthorityKey))
	if len(bz) == 0 {
		return ""
	}
	var authority string
	k.cdc.MustUnmarshalBinaryBare(bz, &authority)
	return authority
}

func (k Keeper) SetTreasurerAccount(ctx sdk.Context, treasurer string) {
	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(TreasurerKey), k.cdc.MustMarshalBinaryBare(treasurer))
}

func (k Keeper) GetTreasurerAccount(ctx sdk.Context) string {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get([]byte(TreasurerKey))
	if len(bz) == 0 {
		return ""
	}
	var treasurer string
	k.cdc.MustUnmarshalBinaryBare(bz, &treasurer)
	return treasurer
}

func (k Keeper) BuyShr(ctx sdk.Context, shrAmt sdk.Int, addr sdk.AccAddress) error {

	if !k.ShrMintPossible(ctx, shrAmt) {
		return types.ErrSHRSupplyExceeded
	}

	rate := k.GetExchangeRate(ctx)
	// rate, err := strconv.ParseFloat(rateStr, 64)
	// if err != nil {
	// 	return err
	// }

	// conv := float64(shrAmt.Int64()) / rate

	oldCoins := k.GetCoins(ctx, addr)
	addedCoins := sdk.NewCoins(sdk.NewCoin("shr", shrAmt))
	removedCoins := sdk.NewCoins()

	// i, d, err := ParseCoinFloat(conv)
	// if err != nil {
	// 	return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	// }
	// shrpAmt := sdk.NewInt(i)
	// centAmt := sdk.NewInt(d)
	shrpAmt := shrAmt.Mul(myutil.RateDecimal).Quo(rate)

	if !oldCoins.AmountOf("shrp").GT(shrpAmt) {
		return sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, fmt.Sprintf("Need %sSHRP, balance %sSHRP", shrpAmt.String(), oldCoins.AmountOf("shrp").String()))
	}
	// if !oldCoins.AmountOf("cent").GT(centAmt) {
	// 	shrpAmt = shrpAmt.AddRaw(int64(1))
	// 	addedCoins = addedCoins.Add(sdk.NewCoin("cent", sdk.NewInt(int64(ShrpToCentRate))))
	// }
	fmt.Println("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx 1")
	removedCoins = removedCoins.Add(sdk.NewCoin("shrp", shrpAmt))

	if _, err := k.AddCoins(ctx, addr, addedCoins); err != nil {
		return err
	}
	fmt.Println("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx 11")
	if _, err := k.SubtractCoins(ctx, addr, removedCoins); err != nil {
		return err
	}
	fmt.Println("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx 12")
	if err := k.SupplyMintCoins(ctx, addedCoins); err != nil {
		return err
	}
	fmt.Println("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx 13")
	if err := k.SupplyBurnCoins(ctx, removedCoins); err != nil {
		return err
	}
	fmt.Println("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx 2")
	return nil
}

func (k Keeper) NotEnoughShr(ctx sdk.Context, amt sdk.Int, addr sdk.AccAddress) (sdk.Int, bool) {
	oldCoins := k.GetCoins(ctx, addr)
	oldShr := oldCoins.AmountOf("shr")
	if !oldShr.GT(amt) {
		return amt.Sub(oldShr), true
	}
	return sdk.ZeroInt(), false
}

func (k Keeper) IsExchangeRatePresent(ctx sdk.Context) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(ExchangeRateKey))
}

func ParseCoinFloat(f float64) (i, d int64, err error) {

	if f < 0 {
		return i, d, sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "Negative Coins are not accepted")
	}
	i = int64(f)
	d = int64(f*100 - float64(i*100) + 1) // make sure always round it up
	return
}

// func (k Keeper) GetIdSigner(ctx sdk.Context, address string) types.IdSigner {
// 	if !k.IsIdSignerPresent(ctx, address) {
// 		return types.NewIdSigner()
// 	}
// 	store := ctx.KVStore(k.storeKey)
// 	bz := store.Get([]byte(address))
// 	var signer types.IdSigner
// 	k.cdc.MustUnmarshalBinaryBare(bz, &signer)
// 	return signer
// }

func (k Keeper) SetIdSigner(ctx sdk.Context, signer types.AccState) {
	store := ctx.KVStore(k.storeKey)
	store.Set(k.createIdSignerKey(signer.Address), k.cdc.MustMarshalBinaryBare(signer.Status))
}

func (k Keeper) GetIdSigner(ctx sdk.Context, signerAddr sdk.AccAddress) types.AccState {
	store := ctx.KVStore((k.storeKey))
	bz := store.Get(k.createIdSignerKey(signerAddr))
	if len(bz) == 0 {
		return types.AccState{}
	}
	var status string
	k.cdc.MustUnmarshalBinaryBare(bz, &status)
	idSigner := types.NewAccState(signerAddr, status)
	return idSigner
}

func (k Keeper) DeleteIdSigner(ctx sdk.Context, addr sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(k.createIdSignerKey(addr))
}

func (k Keeper) DetactivateIdSigner(ctx sdk.Context, addr sdk.AccAddress) {
	signer := k.GetIdSigner(ctx, addr)
	signer.Status = types.Inactive

	k.SetIdSigner(ctx, signer)
}

func (k Keeper) createIdSignerKey(addr sdk.AccAddress) []byte {
	key := fmt.Sprintf("%s%s", IdSignerKey, addr)
	return []byte(key)
}

func (k Keeper) IterateIdSigners(ctx sdk.Context, cb func(loader types.AccState) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(IdSignerKey))

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {

		key := string(iterator.Key())
		var status string
		sAddr, err := sdk.AccAddressFromBech32(key[len(IdSignerKey):])

		if err != nil {
			panic(err)
		}
		err = k.cdc.UnmarshalBinaryBare(iterator.Value(), &status)

		signer := types.AccState{Address: sAddr, Status: status}
		if cb(signer) {
			break
		}
	}
}
