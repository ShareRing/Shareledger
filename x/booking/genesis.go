package booking

import (
	"bitbucket.org/shareringvietnam/shareledger-fix/x/booking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type GenesisState struct {
	Bookings []types.Booking
}

func NewGenesisState() GenesisState {
	return GenesisState{}
}

func ValidateGenesis(data GenesisState) error {
	return nil
}

func DefaultGenesisState() GenesisState {
	return GenesisState{}
}

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) {
	for _, b := range data.Bookings {
		keeper.SetBooking(ctx, b.BookID, b)
	}
}

func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	var bookings []types.Booking
	cb := func(b types.Booking) bool {
		bookings = append(bookings, b)
		return false
	}
	k.IterateBookings(ctx, cb)
	return GenesisState{
		Bookings: bookings,
	}
}
