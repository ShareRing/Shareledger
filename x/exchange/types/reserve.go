package types

import (
	"encoding/hex"
	"fmt"

	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
	"github.com/sharering/shareledger/constants"
	"github.com/sharering/shareledger/types"
	"github.com/sharering/shareledger/utils"
	"github.com/sharering/shareledger/x/bank"
)

type Reserve struct {
	Address sdk.Address
}

func NewReserve(addr sdk.Address) Reserve {
	return Reserve{
		Address: addr,
	}
}

func (res Reserve) IsValid() bool {
	return utils.IsValidReserve(res.Address)
}

func (res Reserve) String() string {
	return fmt.Sprintf("shareledger/Reserve{%s}", res.Address.String())
}

func (res Reserve) GetCoins(
	ctx sdk.Context,
	bankKeeper bank.Keeper,
) types.Coins {
	return bankKeeper.GetCoins(ctx, res.Address)
}

func (res Reserve) SetCoins(
	ctx sdk.Context,
	bankKeeper bank.Keeper,
	newCoins types.Coins,
) sdk.Error {
	return bankKeeper.SetCoins(ctx, res.Address, newCoins)
}

//---------------------------------------------------------------
func GetAllReserve() []Reserve {
	var allRes []Reserve

	for _, resStr := range constants.RESERVE_ACCOUNTS {
		decoded, err := hex.DecodeString(resStr)
		if err != nil {
			panic(err)
		}

		var addr sdk.Address
		copy(addr[:], decoded)

		allRes = append(allRes, NewReserve(addr))
	}
	return allRes
}

//---------------------------------------------------------------
