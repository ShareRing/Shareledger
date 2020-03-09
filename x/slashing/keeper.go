package slashing

import (
	"bytes"
	"encoding/gob"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	pKeeper "github.com/sharering/shareledger/x/pos/keeper"
	"github.com/sharering/shareledger/x/slashing/types"

	"github.com/tendermint/tendermint/crypto"
	db "github.com/tendermint/tm-db"
)

type Keeper struct {
	storeKey *sdk.KVStoreKey
	pk       pKeeper.Keeper
	database db.DB
}

const (
	dbKeyLatestByVal = "%v.latestBlocks"
)

func NewKeeper(key *sdk.KVStoreKey, pk pKeeper.Keeper, database db.DB) Keeper {
	return Keeper{
		storeKey: key,
		pk:       pk,
		database: database,
	}
}

func (k Keeper) HandleValidatorSignature(ctx sdk.Context, batch db.Batch, addr crypto.Address, power int64, signed bool) {
	consAddr := sdk.AccAddress(addr)
	validator, found := k.pk.GetValidatorByConsAddr(ctx, consAddr)
	if !found {
		return
	}
	if validator.Revoked {
		return
	}
	latestBlocks := k.getLatestBlocksForValidator(consAddr)
	latestBlocks = addLatestBlocks(latestBlocks, signed)
	k.setLatestBlocksForValidator(batch, consAddr, latestBlocks)
	missedBlocksCount := countLatestMissedBlocks(latestBlocks)
	if missedBlocksCount > 81 {
		k.pk.Revoke(ctx, consAddr)
		k.setLatestBlocksForValidator(batch, consAddr, []bool{})
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeSlash,
				sdk.NewAttribute(types.AttributeKeyAddress, consAddr.String()),
				sdk.NewAttribute(types.AttributeKeyPower, fmt.Sprintf("%d", power)),
				sdk.NewAttribute(types.AttributeKeyReason, types.AttributeValueMissingSignature),
				sdk.NewAttribute(types.AttributeKeyRevoked, consAddr.String()),
			),
		)
	}
}

func (k Keeper) getLatestBlocksForValidator(address sdk.AccAddress) []bool {
	key := []byte(fmt.Sprintf(dbKeyLatestByVal, address.String()))
	bz := k.database.Get(key)
	if len(bz) == 0 {
		return nil
	}

	b := bytes.NewBuffer(bz)
	dec := gob.NewDecoder(b)

	res := make([]bool, 0)
	err := dec.Decode(&res)
	if err != nil {
		panic(err)
	}

	return res
}

func (k Keeper) setLatestBlocksForValidator(batch db.Batch, address sdk.AccAddress, latestBlocks []bool) {
	bz := new(bytes.Buffer)
	enc := gob.NewEncoder(bz)
	err := enc.Encode(latestBlocks)
	if err != nil {
		panic(err)
	}

	key := []byte(fmt.Sprintf(dbKeyLatestByVal, address.String()))
	batch.Set(key, bz.Bytes())
}

func addLatestBlocks(latestBlocks []bool, signed bool) []bool {
	if len(latestBlocks) > 99 {
		return append(latestBlocks[1:], signed)
	}
	return append(latestBlocks, signed)
}

func countLatestMissedBlocks(latestBlocks []bool) int {
	missedCount := 0
	for _, signed := range latestBlocks {
		if !signed {
			missedCount++
		}
	}
	return missedCount
}
