package app

import (
	"fmt"
	"os"

	"github.com/sharering/shareledger/x/pos"
	abci "github.com/tendermint/abci/types"
	cmn "github.com/tendermint/tmlibs/common"
	dbm "github.com/tendermint/tmlibs/db"
	"github.com/tendermint/tmlibs/log"

	bapp "bitbucket.org/shareringvn/cosmos-sdk/baseapp"
	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
	"bitbucket.org/shareringvn/cosmos-sdk/wire"

	"github.com/sharering/shareledger/constants"
	"github.com/sharering/shareledger/types"

	"github.com/sharering/shareledger/x/asset"
	"github.com/sharering/shareledger/x/auth"
	"github.com/sharering/shareledger/x/bank"
	"github.com/sharering/shareledger/x/booking"
)

const (
	appName = "ShareLedger_v0.0.1"
)

var (
	DefaultCLIHome = os.ExpandEnv("$HOME/.shareledgercli")
)

type ShareLedgerApp struct {
	*bapp.BaseApp
	cdc *wire.Codec

	// keys to access the substores
	assetKey   *sdk.KVStoreKey
	bookingKey *sdk.KVStoreKey
	//accountKey *sdk.KVStoreKey

	// Manage getting and setting accounts
	accountMapper auth.AccountMapper
}

func NewShareLedgerApp(logger log.Logger, db dbm.DB) *ShareLedgerApp {

	cdc := MakeCodec()

	// Create the base application object.
	baseApp := bapp.NewBaseApp(appName, cdc, logger, db)

	assetKey := sdk.NewKVStoreKey(constants.STORE_ASSET)
	bookingKey := sdk.NewKVStoreKey(constants.STORE_BOOKING)
	//accountKey := sdk.NewKVStoreKey(constants.STORE_BANK)
	authKey := sdk.NewKVStoreKey(constants.STORE_AUTH)

	// Mount Store

	baseApp.MountStoresIAVL(authKey, assetKey, bookingKey)
	err := baseApp.LoadLatestVersion(authKey)
	if err != nil {
		cmn.Exit(err.Error())
	}

	// accountMapper for Auth Module storing and Bank module
	accountMapper := auth.NewAccountMapper(
		cdc,
		authKey,
		&auth.SHRAccount{},
	)

	//baseApp.SetInitChainer(InitChainer(cdc, accountMapper)) //working on it
	//baseApp.SetEndBlocker(EndBlocker) //working on it

	SetupAsset(baseApp, cdc, assetKey)
	SetupBank(baseApp, cdc, accountMapper)
	SetupBooking(baseApp, cdc, bookingKey, assetKey, accountMapper)

	// Determine how transactions are decoded.
	//baseApp.SetTxDecoder(types.GetTxDecoder(cdc))
	baseApp.SetTxDecoder(auth.GetTxDecoder(cdc))
	baseApp.SetAnteHandler(auth.NewAnteHandler(accountMapper))
	baseApp.Router().
		AddRoute(constants.MESSAGE_AUTH, auth.NewHandler(accountMapper))
	cdc = auth.RegisterCodec(cdc)

	// Register InitChain
	// logger.Info("Register Init Chainer")
	// baseApp.SetInitChainer(InitChainer(cdc, accountMapper))

	return &ShareLedgerApp{
		BaseApp:    baseApp,
		assetKey:   assetKey,
		bookingKey: bookingKey,
		//accountKey:    accountKey,
		accountMapper: accountMapper,
	}
}

// InitChainer will set initial balances for accounts as well as initial coin metadata

func InitChainer(cdc *wire.Codec, accountMapper auth.AccountMapper) sdk.InitChainer {
	return func(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
		stateJSON := req.AppStateBytes
		fmt.Printf("RequestInitChain.Time: %v\n", req.Time)
		fmt.Printf("RequestInitChain.ChainId: %v\n", req.ChainId)
		fmt.Printf("RequestInitChain.ConsensusParams: %v\n", req.ConsensusParams)
		fmt.Printf("RequestInitChain.Validators: %v\n", req.Validators)
		fmt.Printf("RequestInitChain.AppStateBytes: %v\n", req.AppStateBytes)

		var genesisState GenesisState
		fmt.Printf("stateJSON=%s\n", stateJSON)

		err := cdc.UnmarshalJSON(stateJSON, &genesisState)
		fmt.Printf("req=%v\n", genesisState)
		if err != nil {
			panic(err)
		}

		// load the accounts - TODO

		// load the initial POS information
		abciVals, err := pos.InitGenesis(ctx, pos.Keeper{}, genesisState.StakeData)
		if err != nil {
			panic(err)
		}
		for _, abciVal := range abciVals {
			fmt.Printf("abciVal=%v\n", abciVal)
		}
		return abci.ResponseInitChain{
			Validators: abciVals, //use the validator defined in stake
		}

	}
}

// application updates every end block

func EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {

	validatorUpdates := pos.EndBlocker(ctx, pos.Keeper{})
	// Add these new validators to the addr -> pubkey map.

	return abci.ResponseEndBlock{
		ValidatorUpdates: validatorUpdates,
	}
}

func MakeCodec() *wire.Codec {
	cdc := wire.NewCodec()
	cdc.RegisterInterface((*types.SHRTx)(nil), nil)
	cdc.RegisterConcrete(types.BasicTx{}, "shareledger/BasicTx", nil)
	cdc.RegisterConcrete(auth.AuthTx{}, "shareledger/AuthTx", nil)
	cdc.RegisterConcrete(types.QueryTx{}, "shareledger/QueryTx", nil)

	cdc.RegisterInterface((*types.SHRSignature)(nil), nil)
	cdc.RegisterConcrete(types.BasicSig{}, "shareledger/BasicSig", nil)
	cdc.RegisterConcrete(auth.AuthSig{}, "shareledger/AuthSig", nil)

	cdc.RegisterInterface((*auth.BaseAccount)(nil), nil)
	cdc.RegisterConcrete(auth.SHRAccount{}, "shareledger/SHRAccount", nil)

	cdc.RegisterInterface((*types.PubKey)(nil), nil)
	cdc.RegisterConcrete(types.PubKeySecp256k1{}, "shareledger/PubSecp256k1", nil)

	cdc.RegisterInterface((*types.Signature)(nil), nil)
	cdc.RegisterConcrete(types.SignatureSecp256k1{}, "shareledger/SigSecp256k1", nil)

	cdc.RegisterInterface((*sdk.Msg)(nil), nil)
	return cdc
}

func SetupBank(app *bapp.BaseApp, cdc *wire.Codec, am auth.AccountMapper) {
	// Bank module
	// Create a key for accessing the account store.
	cdc = bank.RegisterCodec(cdc)

	// Register message routes.
	// Note the handler gets access to the account store.
	app.Router().
		AddRoute("bank", bank.NewHandler(am))

}

func SetupAsset(app *bapp.BaseApp, cdc *wire.Codec, assetKey *sdk.KVStoreKey) {

	keeper := asset.NewKeeper(assetKey, cdc)

	cdc = asset.RegisterCodec(cdc)

	app.Router().
		AddRoute("asset", asset.NewHandler(keeper))

	// app.MountStoresIAVL(assetKey)
}

func SetupBooking(app *bapp.BaseApp, cdc *wire.Codec, bookingKey *sdk.KVStoreKey,
	assetKey *sdk.KVStoreKey, am auth.AccountMapper) {

	cdc = booking.RegisterCodec(cdc)

	k := booking.NewKeeper(bookingKey,
		assetKey,
		am,
		cdc)

	app.Router().
		AddRoute("booking", booking.NewHandler(k))

	// app.MountStoresIAVL(bookingKey)

}
