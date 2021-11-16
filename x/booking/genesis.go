package booking

import (
	"github.com/ShareRing/Shareledger/x/booking/keeper"
	"github.com/ShareRing/Shareledger/x/booking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	for _, b := range genState.Bookings {
		k.SetBooking(ctx, b.BookID, *b)
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	var bookings []*types.Booking

	cb := func(b types.Booking) bool {
		bookings = append(bookings, &b)
		return false
	}

	k.IterateBookings(ctx, cb)

	return &types.GenesisState{
		Bookings: bookings,
	}
}

// type types.GenesisState struct {
// 	Bookings []types.Booking
// }

func NewGenesisState() types.GenesisState {
	return types.GenesisState{}
}

func ValidateGenesis(data types.GenesisState) error {
	return nil
}

func DefaultGenesisState() types.GenesisState {
	return types.GenesisState{}
}
