package types

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func parseShrpCoinStr(s string) (shrp, cent int64, err error) {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return
	}
	if f < 0 {
		err = sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "Negative Coins are not accepted")
		return
	}

	shrp = int64(f)
	cent = int64(f*100 - float64(shrp*100))

	return
}

func ParseShrpCoinsStr(s string) (coins sdk.Coins, err error) {
	shrp, cent, err := parseShrpCoinStr(s)
	if err != nil {
		return nil, err
	}
	return sdk.Coins{
		sdk.NewCoin(DenomSHRP, sdk.NewInt(shrp)),
		sdk.NewCoin(DenomCent, sdk.NewInt(cent)),
	}, nil
}
