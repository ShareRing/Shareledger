package utils

import (
	"encoding/json"
	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
)

func Store(store sdk.KVStore, key []byte, value interface{}) error {
	// Encoding
	res, err := json.Marshal(value)

	if err != nil {
		return err
	}

	store.Set(key, res)
	return nil
}


func Retrieve(store sdk.KVStore, key []byte, value interface{})  error{
	// Get value
	valueBytes := store.Get(key)

	if valueBytes != nil {
		return nil
	}

	err := json.Unmarshal(valueBytes, &value)
	if err != nil {
		return err
	}

	return nil
}


