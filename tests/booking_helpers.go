package tests

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/tests"
	app "github.com/sharering/shareledger"
	"github.com/sharering/shareledger/x/booking"
	"github.com/stretchr/testify/require"
)

const (
	BookingModuleName = booking.ModuleName
)

func (f *Fixtures) ExecuteBookingBook(bookingUUID string,bookDuration int ,bookerKey string ) (bool, string, string) {
	flag := []string{fmt.Sprintf("--key-seed ./%s_key_seed.json --yes --fees 1shr", bookerKey)}
	cmd := fmt.Sprintf("%s tx %v book %v %v %v", f.GaiacliBinary,BookingModuleName,bookingUUID ,bookDuration, f.Flags())
	return executeWriteRetStdStreams(f.T,addFlags(cmd,flag), DefaultKeyPass)
}

func (f *Fixtures) ExecuteBookingComplete(bookingID string,bookerKey string ) (bool, string, string) {
	flag := []string{fmt.Sprintf("--key-seed ./%s_key_seed.json --yes --fees 1shr", bookerKey)}
	cmd := fmt.Sprintf("%s tx %v complete %v %v", f.GaiacliBinary,BookingModuleName,bookingID , f.Flags())
	return executeWriteRetStdStreams(f.T,addFlags(cmd,flag), DefaultKeyPass)
}


func (f *Fixtures) ExecuteBookingGetBook(bookingID string )  booking.Booking {
	cmd := fmt.Sprintf("%s query %v get %s %v", f.GaiacliBinary,BookingModuleName, bookingID, f.Flags())
	out, _ := tests.ExecuteT(f.T, cmd, "")
	var voter booking.Booking
	cdc := app.MakeCodec()
	err := cdc.UnmarshalJSON([]byte(out), &voter)
	require.NoError(f.T, err, "out %v\n, err %v", out, err)
	return voter
}