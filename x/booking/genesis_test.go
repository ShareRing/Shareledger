package booking_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/sharering/shareledger/testutil/keeper"
	"github.com/sharering/shareledger/x/booking"
	"github.com/sharering/shareledger/x/booking/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BookingKeeper(t)
	booking.InitGenesis(ctx, *k, genesisState)
	got := booking.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	// this line is used by starport scaffolding # genesis/test/assert
}
