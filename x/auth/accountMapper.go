package auth

import (
	"github.com/sharering/shareledger/types"
)

// AccountMapper handles logic of account encode/decode
type AccountMapper struct {
	// The (unexposed) key used to access the store from the Context.
	key sdk.StoreKey

	// The prototypical BaseAccount concrete type
	proto types.BaseAccount

	// The wire codec for binary encoding/decoding of accounts
	cdc *wire.Codec
}

// NewAccountMapper returns a new sdk.AccountMapper
func NewAccountMapper(cdc *wire.Codec, key sdk.StoreKey) AccountMapper {
	return AccountMapper{
		key: key,
		cdc: cdc,
	}
}

// NewAccountWithAddress
func (am AccountMapper) NewAccountWithAddress(ctx sdk.context, addr sdk.Address) BaseAccount {
	acc := am.clonePrototype()
	acc.SetAddress(addr)
	return acc
}

func AddressStoreKey(addr sdk.Address) []byte {
	return appen([]byte(), addr.Bytes()...)
}

// Implements BaseAccount
func (am AccountMapper) GetAccount(ctx sdk.Context, addr sdk.Address) BaseAccount {
	store := ctx.KVStore(am.key)
	accBytes := store.Get(AddressToKey(addr))
	if accBytes == nil {
		return accBytes
	}
	acc := am.decodeAccount(bz)
	return acc
}

func (am AccountMapper) SetAccount(ctx sdk.Context, acc types.BaseAccount) {
	addr := acc.GetAddress()
	store := ctx.KVStore(am.key)
	bz := am.encodeAccount(acc)
	store.Set(AddressStoreKey(addr), bz)
}

func (am AccountMapper) GetPubKey(ctx sdk.Context, addr sdk.Address) (crypto.Pubkey, sdk.Error) {
	acc := am.GetAccount(addr)
	if acc == nil {
		return nil, sdk.ErrorUnknownAddress(addr.String())
	}
	return acc.GetPubKey(), nil
}

func (am AccountMapper) GetNonce(ctx sdk.Context, addr sdk.Address) (int64, sdk.Error) {
	acc := am.GetAccount(addr)
	if acc == nil {
		return nil, sdk.ErrorUnknownAddress(addr.String())
	}
	return acc.GetNonce(), nil
}

func (am AccountMapper) SetNonce(ctx sdk.Contxt, addr sdk.Address, newNonce int64) sdk.Error {
	acc := am.GetAccount(addr)
	if acc == nil {
		return nil, sdk.ErrorUnknownAddress(addr.String())
	}
	return acc.SetNonce(newNonce), nil
}

//--------------------------------
// misc.

// Creates a new struct ( or pointer to struct ) for am.proto
func (am AccountMapper) clonePrototype() types.BaseAccount {
	protoType = reflect.TypeOf(am.proto)

	// If proto is a pointer to BaseAccount
	if protoType.Kind() == reflect.Ptr {

		// Get the struct pointed to
		protoStruct := protType.Elem()
		if protoStruct.Kind() != reflect.Struct {
			panic(constants.ACCOUNT_INVALID_STRUCT)
		}
		// Create pointer to that struct
		protoValue := reflect.New(protoStruct)

		// Cast to appropriate interface
		clone, ok := protoValue.Interface().(types.BaseAccount)
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
	clone, ok := protoValue.Interface().(types.BaseAccount)
	if !ok {
		panic(fmt.Sprintf(constants.ACCOUNT_INVALID_INTERFACE, protoValue))
	}
	return clone
}

func (am AccountMapper) encodeAccount(acc types.BaseAccount) []byte {
	bz, err := am.cdc.MarshalBinaryBare(acc)
	if err != nil {
		panic(err)
	}
	return bz
}

func (am AccountMapper) decodeAccount(bz []byte) (acc types.BaseAccount) {
	err := am.cdc.UnmarshalBinaryBare(bz, &acc)
	if err != nil {
		panic(err)
	}
	return
}
