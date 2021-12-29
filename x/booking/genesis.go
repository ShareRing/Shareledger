package booking

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/booking/keeper"
	"github.com/sharering/shareledger/x/booking/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	for _, b := range genState.Bookings {
		k.SetBooking(ctx, b.GetBookID(), *b)
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// this line is used by starport scaffolding # genesis/module/export
	cb := func(b types.Booking) bool {
		genesis.Bookings = append(genesis.Bookings, &b)
		return false
	}

	k.IterateBookings(ctx, cb)

	return genesis
}
