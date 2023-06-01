package booking

import (
	"fmt"

	"github.com/sharering/shareledger/tests"
	"github.com/sharering/shareledger/x/booking/types"
)

func (s *E2ETestSuite) TestGRPCQueryBooking() {
	val := s.network.Validators[0]
	getURL := func(bookingID string) string {
		return fmt.Sprintf("%s/shareledger/booking/%s", val.APIAddress, bookingID)
	}

	testCases := tests.TestCasesGrpc{
		{
			Name:      "gRPC get booking by id",
			URL:       getURL(booking1.BookID),
			ExpectErr: false,
			RespType:  &types.QueryBookingResponse{},
			Expected: &types.QueryBookingResponse{
				Booking: booking1,
			},
		},
		{
			Name:      "gRPC get booking by id empty",
			URL:       getURL(""),
			ExpectErr: true,
		},
		{
			Name:      "gRPC get booking by id empty not exists",
			URL:       getURL("noneExistsBookingID"),
			ExpectErr: true,
		},
	}
	tests.RunTestCasesGrpc(&s.Suite, testCases, val)
}
