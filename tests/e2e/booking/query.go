package booking

import (
	"github.com/sharering/shareledger/tests"
	"github.com/sharering/shareledger/x/booking/client/cli"
	"github.com/sharering/shareledger/x/booking/types"
)

func (s *E2ETestSuite) TestQueryBooking() {
	testCases := tests.TestCases{
		{
			Name:      "query booking by bookID",
			Args:      []string{booking1.BookID},
			ExpectErr: false,
			RespType:  &types.QueryBookingResponse{},
			Expected: &types.QueryBookingResponse{
				Booking: booking1,
			},
		},
		{
			Name:      "query booking by bookID empty",
			Args:      []string{},
			ExpectErr: true,
		},
		{
			Name: "query booking by bookID not exists",
			Args: []string{
				"notExistsBookID",
			},
			ExpectErr: true,
		},
	}
	tests.RunTestCases(&s.Suite, testCases, cli.CmdBooking(), s.network.Validators[0])
}
