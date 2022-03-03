package network

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/sharering/shareledger/app"
	assettypes "github.com/sharering/shareledger/x/asset/types"
	bookingtypes "github.com/sharering/shareledger/x/booking/types"
	documenttypes "github.com/sharering/shareledger/x/document/types"
	gentleminttypes "github.com/sharering/shareledger/x/gentlemint/types"
	idtypes "github.com/sharering/shareledger/x/id/types"
	denom "github.com/sharering/shareledger/x/utils/demo"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spm/cosmoscmd"
	tmos "github.com/tendermint/tendermint/libs/os"
	"path/filepath"
	"testing"
)

const (
	KeyAuthority   string = "authority"
	KeyTreasurer   string = "treasurer"
	KeyOperator    string = "operator"
	KeyLoader      string = "loader"
	KeyMillionaire string = "millionaire"

	KeyDocIssuer string = "doc_issuer"
	KeyIDSigner  string = "id_signer"

	KeyEmpty1 string = "empty1"
	KeyEmpty2 string = "empty2"
	KeyEmpty3 string = "empty3"
	KeyEmpty4 string = "empty4"
	KeyEmpty5 string = "empty5"

	KeyAccount1 string = "acc1"
	KeyAccount2 string = "acc2"
	KeyAccount3 string = "acc3"
	KeyAccount4 string = "acc4"

	KeyAccount5 string = "acc5" //Use this account if you want more
	KeyAccount6 string = "acc6" //Use this account if you want more
	KeyAccount7 string = "acc7" //Use this account if you want more
	KeyAccount8 string = "acc8" //Use this account if you want more

	ShareLedgerSuccessCode = uint32(0)
)

var (
	oneThousandSHR = 10000 * denom.ShrExponent
	oneHundredSHRP = 100 * denom.USDExponent
	oneMillionSHR  = 1000000 * denom.ShrExponent //1 million shr and shrp
	oneMillionSHRP = 1000000 * denom.USDExponent //1 million shr and shrp

	defaultCoins  = sdk.NewCoins(sdk.NewCoin(denom.Base, sdk.NewInt(oneThousandSHR)), sdk.NewCoin(denom.BaseUSD, sdk.NewInt(oneHundredSHRP)))
	becauseImRich = sdk.NewCoins(sdk.NewCoin(denom.Base, sdk.NewInt(oneMillionSHR)), sdk.NewCoin(denom.BaseUSD, sdk.NewInt(oneMillionSHRP)))
	poorMen       = sdk.NewCoins(sdk.NewCoin(denom.Base, sdk.NewInt(0)), sdk.NewCoin(denom.BaseUSD, sdk.NewInt(0)))

	ShareLedgerErrorCodeUnauthorized = sdkerrors.ErrUnauthorized.ABCICode()
	ShareLedgerErrorCodeMaxSupply    = gentleminttypes.ErrBaseSupplyExceeded.ABCICode()

	ShareLedgerErrorCodeAssetNotExisted = assettypes.ErrNameDoesNotExist.ABCICode()

	ShareLedgerErrorCodeDocumentAlreadyExisted = documenttypes.ErrDocExisted.ABCICode()
	ShareLedgerDocumentNotFound                = documenttypes.ErrDocNotExisted.ABCICode()

	ShareLedgerCoinInsufficient = sdkerrors.ErrInsufficientFunds.ABCICode()

	ShareLedgerErrorCodeIDAddressOwnerID = idtypes.ErrWrongBackupAddr.ABCICode()

	ShareLedgerBookingAssetAlreadyBooked = bookingtypes.ErrAssetAlreadyBooked.ABCICode()
	ShareLedgerBookingBookerIsNotOwner   = bookingtypes.ErrNotBookerOfAsset.ABCICode()
	//ShareLedgerInsufficientTXNFee        = sdkerrors.ErrInsufficientFee.ABCICode()
)

type (
	CosmosLogs []CosmosLog

	CosmosLog struct {
		MgsIndex int    `json:"mgs_index"`
		Events   Events `json:"events"`
	}
	Event struct {
		Type       string      `json:"type"`
		Attributes []Attribute `json:"attributes"`
	}

	Attribute struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}
	Events     []Event
	Attributes []Attribute
)

func (e Events) GetEventByType(t *testing.T, eType string) Attributes {
	for _, ev := range e {
		if ev.Type == eType {
			return ev.Attributes
		}
	}
	t.Log("event type not found")
	t.Fail()

	return nil
}

func (as Attributes) Get(t *testing.T, key string) Attribute {
	for _, a := range as {
		if a.Key == key {
			return a
		}
	}
	t.Log("attribute key not found")
	t.Fail()
	return Attribute{}
}

//Use later
func writeFile(name string, dir string, contents []byte) error {
	writePath := filepath.Join(dir)
	file := filepath.Join(writePath, name)

	err := tmos.EnsureDir(writePath, 0755)
	if err != nil {
		return err
	}

	err = tmos.WriteFile(file, contents, 0644)
	if err != nil {
		return err
	}

	return nil
}

func ParseStdOut(t *testing.T, stdOut []byte) sdk.TxResponse {
	txResponse := sdk.TxResponse{}

	encCfg := cosmoscmd.MakeEncodingConfig(app.ModuleBasics)
	err := encCfg.Marshaler.UnmarshalJSON(stdOut, &txResponse)
	require.NoError(t, err)
	return txResponse
}

func BalanceJsonUnmarshal(t *testing.T, data []byte) banktypes.QueryAllBalancesResponse {
	var b banktypes.QueryAllBalancesResponse
	encCfg := cosmoscmd.MakeEncodingConfig(app.ModuleBasics)
	err := encCfg.Marshaler.UnmarshalJSON(data, &b)
	require.NoError(t, err)
	return b

}

func ParseRawLogGetEvent(t *testing.T, logString string) CosmosLogs {
	var logs CosmosLogs
	err := json.Unmarshal([]byte(logString), &logs)
	require.NoError(t, err, "fail to get the log information form stdout")
	l := len(logs)
	require.Greater(t, l, 0, "empty logs")
	return logs
}
