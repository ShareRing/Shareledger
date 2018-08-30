package auth

import (
	"fmt"
	"reflect"

	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"

	"bitbucket.org/shareringvn/cosmos-sdk/wire"

	"github.com/sharering/shareledger/constants"
	"github.com/sharering/shareledger/types"
)

// AccountMapper handles logic of account encode/decode
type AccountMapper struct {
	// The (unexposed) key used to access the store from the Context.
	key sdk.StoreKey

	// The prototypical BaseAccount concrete type
	proto BaseAccount

	// The wire codec for binary encoding/decoding of accounts
	cdc *wire.Codec
}

// NewAccountMapper returns a new sdk.AccountMapper
func NewAccountMapper(cdc *wire.Codec, key sdk.StoreKey, proto BaseAccount) AccountMapper {
	return AccountMapper{
		key:   key,
		cdc:   cdc,
		proto: proto,
	}
}

// NewAccountWithAddress
func (am AccountMapper) NewAccountWithAddress(ctx sdk.Context, addr sdk.Address) BaseAccount {
	acc := am.clonePrototype()
	acc.SetAddress(addr)
	return acc
}

func AddressToKey(addr sdk.Address) []byte {
	return append([]byte(constants.PREFIX_ADDRESS), addr.Bytes()...)
}

// Implements BaseAccount
func (am AccountMapper) GetAccount(ctx sdk.Context, addr sdk.Address) BaseAccount {
	store := ctx.KVStore(am.key)
	accBytes := store.Get(AddressToKey(addr))
	if accBytes == nil {
		return nil
	}
	acc := am.decodeAccount(accBytes)
	return acc
}

func (am AccountMapper) SetAccount(ctx sdk.Context, acc BaseAccount) {
	addr := acc.GetAddress()
	store := ctx.KVStore(am.key)
	bz := am.encodeAccount(acc)
	store.Set(AddressToKey(addr), bz)
}

func (am AccountMapper) GetPubKey(ctx sdk.Context, addr sdk.Address) (types.PubKey, sdk.Error) {
	acc := am.GetAccount(ctx, addr)
	if acc == nil {
		return nil, sdk.ErrUnknownAddress(addr.String())
	}
	return acc.GetPubKey(), nil
}

func (am AccountMapper) GetNonce(ctx sdk.Context, addr sdk.Address) (int64, sdk.Error) {
	acc := am.GetAccount(ctx, addr)
	if acc == nil {
		return 0, nil
		//return -1, sdk.ErrUnknownAddress(addr.String())
	}
	return acc.GetNonce(), nil
}

func (am AccountMapper) SetNonce(ctx sdk.Context, addr sdk.Address, newNonce int64) sdk.Error {
	acc := am.GetAccount(ctx, addr)
	if acc == nil {
		return sdk.ErrUnknownAddress(addr.String())
	}
	acc.SetNonce(newNonce)
	am.SetAccount(ctx, acc)
	return nil
}

//--------------------------------
// misc.

// Creates a new struct ( or pointer to struct ) for am.proto
func (am AccountMapper) clonePrototype() BaseAccount {
	protoType := reflect.TypeOf(am.proto)

	// If proto is a pointer to BaseAccount
	if protoType.Kind() == reflect.Ptr {

		// Get the struct pointed to
		protoStruct := protoType.Elem()
		if protoStruct.Kind() != reflect.Struct {
			panic(constants.ACCOUNT_INVALID_STRUCT)
		}
		// Create pointer to that struct
		protoValue := reflect.New(protoStruct)

		// Cast to appropriate interface
		clone, ok := protoValue.Interface().(BaseAccount)
		if !ok {
			panic(fmt.Sprintf(constants.ACCOUNT_INVALID_INTERFACE, protoValue))
		}
		return clone
	}

	// if proto is a struct of BaseAccount Interface
	protoValuePtr := reflect.New(protoType)

	// Get struct
	protoValue := protoValuePtr.Elem()

	// Cast to appropriate interface
	clone, ok := protoValue.Interface().(BaseAccount)
	if !ok {
		panic(fmt.Sprintf(constants.ACCOUNT_INVALID_INTERFACE, protoValue))
	}
	return clone
}

func (am AccountMapper) encodeAccount(acc BaseAccount) []byte {
	bz, err := am.cdc.MarshalBinaryBare(acc)
	if err != nil {
		panic(err)
	}
	return bz
}

func (am AccountMapper) decodeAccount(bz []byte) (acc BaseAccount) {
	var acc1 *SHRAccount
	err := am.cdc.UnmarshalBinaryBare(bz, &acc1)
	if err != nil {
		panic(err)
	}
	return acc1
}
