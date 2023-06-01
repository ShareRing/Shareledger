package asset

import (
	errorsmod "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/tests"
	"github.com/sharering/shareledger/testutil/network"
	"github.com/sharering/shareledger/x/asset/client/cli"
	"github.com/sharering/shareledger/x/asset/types"
)

func (s *E2ETestSuite) TestCreateAsset() {
	testCases := tests.TestCasesTx{
		{
			Name: "create new asset",
			Args: []string{
				"hash",
				"new_uuid",
				"true",
				"10",
				network.MakeByAccount(network.KeyAccount1),
			},
			ExpectErr:    false,
			ExpectedCode: errorsmod.SuccessABCICode,
		},
		{
			Name: "create new asset invalid arg 2",
			Args: []string{
				"hash",
				"new_uuid",
				"invalid_bool_value",
				"10",
				network.MakeByAccount(network.KeyAccount1),
			},
			ExpectErr: true,
		},
		{
			Name: "create new asset invalid arg 3",
			Args: []string{
				"hash",
				"new_uuid",
				"true",
				"invalid_int",
				network.MakeByAccount(network.KeyAccount1),
			},
			ExpectErr: true,
		},
	}
	tests.RunTestCasesTx(&s.Suite, testCases, cli.CmdCreate(), s.network.Validators[0])
}

func (s *E2ETestSuite) TestDeleteAsset() {
	s.T().Log("TestDeleteAsset")

	deleteUUID := "TestDeleteAsset"
	testCases := tests.TestCasesTx{
		{
			Name: "delete asset unauthorize",
			Args: []string{
				deleteUUID,
				network.MakeByAccount(network.KeyAccount2),
			},
			ExpectErr:    false,
			ExpectedCode: sdkerrors.ErrUnauthorized.ABCICode(),
		},
		{
			Name: "delete asset",
			Args: []string{
				deleteUUID,
				network.MakeByAccount(network.KeyAccount1),
			},
			ExpectErr:    false,
			ExpectedCode: errorsmod.SuccessABCICode,
		},
		{
			Name: "delete asset not exists",
			Args: []string{
				deleteUUID,
				network.MakeByAccount(network.KeyAccount1),
			},
			ExpectErr:    false,
			ExpectedCode: sdkerrors.ErrNotFound.ABCICode(),
		},
	}

	// create asset for delete
	s.createNewAsset(deleteUUID)
	tests.RunTestCasesTx(&s.Suite, testCases, cli.CmdDelete(), s.network.Validators[0])
}

func (s *E2ETestSuite) TestUpdateAsset() {
	uuid := "TestUpdateAsset"
	testCases := tests.TestCasesTx{
		{
			Name: "update asset unauthorize",
			Args: []string{
				"new_hash",
				uuid,
				"false",
				"1000",
				network.MakeByAccount(network.KeyAccount2),
			},
			ExpectErr:    false,
			ExpectedCode: sdkerrors.ErrUnauthorized.ABCICode(),
		},
		{
			Name: "update asset",
			Args: []string{
				"new_hash",
				uuid,
				"false",
				"1000",
				network.MakeByAccount(network.KeyAccount1),
			},
			ExpectErr:    false,
			ExpectedCode: errorsmod.SuccessABCICode,
		},
		{
			Name: "update asset not exists",
			Args: []string{
				"new_hash",
				"notExistsUUID",
				"false",
				"1000",
				network.MakeByAccount(network.KeyAccount1),
			},
			ExpectErr:    false,
			ExpectedCode: types.ErrNameDoesNotExist.ABCICode(),
		},
	}
	// create asset for update
	s.createNewAsset(uuid)
	tests.RunTestCasesTx(&s.Suite, testCases, cli.CmdUpdate(), s.network.Validators[0])
}

func (s *E2ETestSuite) createNewAsset(uuid string) {
	_, err := tests.RunCmdWithRetry(&s.Suite,
		cli.CmdCreate(),
		s.network.Validators[0],
		[]string{"hash", uuid, "true", "10", network.MakeByAccount(network.KeyAccount1)},
		100,
	)
	s.NoError(err)
}
