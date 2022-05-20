package services

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/pkg/errors"
)

type cachedKeyRing struct {
	keyring.Keyring
	cachedInfo map[string]keyring.Info
}

func (ck cachedKeyRing) Key(uid string) (keyring.Info, error) {
	info, found := ck.cachedInfo[uid]
	if !found {
		return nil, errors.New(fmt.Sprintf("%v not found", uid))
	}
	return info, nil
}
func (ck *cachedKeyRing) InitCaches(uids []string) error {
	var err error
	ck.cachedInfo = make(map[string]keyring.Info)
	for _, uid := range uids {
		info, err := ck.Keyring.Key(uid)
		if err != nil {
			return err
		}
		ck.cachedInfo[uid] = info
	}
	return err
}
