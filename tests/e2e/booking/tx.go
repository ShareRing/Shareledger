package booking

import (
	errorsmod "cosmossdk.io/errors"
	"github.com/sharering/shareledger/tests"
	"github.com/sharering/shareledger/testutil/network"
	"github.com/sharering/shareledger/x/booking/client/cli"
	"github.com/sharering/shareledger/x/booking/keeper"
	"github.com/sharering/shareledger/x/booking/types"
)

func (s *E2ETestSuite) TestCreatBooking() {
	testCases := tests.TestCasesTx{
		{
			Name:         "create new booking",
			Args:         []string{asset1.UUID, "1", network.MakeByAccount(network.KeyAccount1)},
			ExpectErr:    false,
			ExpectedCode: errorsmod.SuccessABCICode,
		},
		{
			Name:      "create new booking not enough args",
			Args:      []string{"", network.MakeByAccount(network.KeyAccount1)},
			ExpectErr: true,
		},
		{
			Name:      "create new booking invalid argument 1",
			Args:      []string{"newUUID"},
			ExpectErr: true,
		},
	}
	tests.RunTestCasesTx(&s.Suite, testCases, cli.CmdBook(), s.network.Validators[0])
}

func (s *E2ETestSuite) TestCompleteBooking() {
	bookID := s.createNewBooking(asset2.UUID)
	testCases := tests.TestCasesTx{
		{
			Name:      "complete booking not enough args",
			Args:      []string{},
			ExpectErr: true,
		},
		{
			Name: "complete booking",
			Args: []string{
				bookID,
				network.MakeByAccount(network.KeyAccount1),
			},
			ExpectErr:    false,
			ExpectedCode: errorsmod.SuccessABCICode,
		},
		{
			Name: "complete booking not exists bookID",
			Args: []string{
				"notExistsBookingID",
				network.MakeByAccount(network.KeyAccount1),
			},
			ExpectErr:    false,
			ExpectedCode: types.ErrBookingDoesNotExist.ABCICode(),
		},
	}
	tests.RunTestCasesTx(&s.Suite, testCases, cli.CmdComplete(), s.network.Validators[0])
}

func (s *E2ETestSuite) createNewBooking(uuid string) string {
	val := s.network.Validators[0]
	_, err := tests.RunCmdWithRetry(&s.Suite,
		cli.CmdBook(),
		val,
		[]string{uuid, "1", network.MakeByAccount(network.KeyAccount1)},
		100,
	)
	s.NoError(err)
	booker := network.MustAddressFormKeyring(val.ClientCtx.Keyring, network.KeyAccount1)
	bookID, err := keeper.GenBookID(&types.MsgCreateBooking{
		Booker:   booker.String(),
		UUID:     uuid,
		Duration: 1,
	})
	s.NoError(err)
	return bookID
}
