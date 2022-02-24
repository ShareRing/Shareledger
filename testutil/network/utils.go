package network

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/sharering/shareledger/app"
	denom "github.com/sharering/shareledger/x/utils/demo"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spm/cosmoscmd"
	"testing"
)

const (
	ErrorCodeNotFound = uint32(41)
	ErrorCodeExisted  = uint32(42)
	ErrorNotAvailable = uint32(45)

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

	ShareLedgerSuccessCode             = uint32(0)
	ShareLedgerErrorCodeUnauthorized   = uint32(4)
	ShareLedgerErrorCodeInvalidCoin    = uint32(10)
	ShareLedgerErrorCodeInvalidRequest = uint32(18)
	ShareLedgerErrorCodeMaxSupply      = uint32(41)

	ShareLedgerErrorCodeAssetNotExisted     = ErrorCodeNotFound
	ShareLedgerErrorCodeAssetAlreadyExisted = ErrorCodeExisted

	ShareLedgerErrorCodeDocumentAlreadyExisted = ErrorCodeExisted

	ShareLedgerBookingAssetAlreadyBooked = uint32(43)
	ShareLedgerBookingBookerIsNotOwner   = uint32(46)

	ShareLedgerDocumentDuplicated = uint32(3)
	ShareLedgerDocumentNotFound   = ErrorCodeNotFound

	ShareLedgerErrorCodeIDAddressOwnerID = ErrorNotAvailable
	ShareLedgerCoinInsufficient          = uint32(5)
)

var (
	defaultInitSHR    = 10000 * denom.ShrExponent
	defaultInitSHRP   = 100 * denom.USDExponent
	becauseImRichSHR  = 1000000 * denom.ShrExponent //10 0 million shr and shrp
	becauseImRichSHRP = 1000000 * denom.USDExponent //100 million shr and shrp
	defaultCoins      = sdk.NewCoins(sdk.NewCoin(denom.Base, sdk.NewInt(defaultInitSHR)), sdk.NewCoin(denom.BaseUSD, sdk.NewInt(defaultInitSHRP)))
	becauseImRich     = sdk.NewCoins(sdk.NewCoin(denom.Base, sdk.NewInt(becauseImRichSHR)), sdk.NewCoin(denom.BaseUSD, sdk.NewInt(becauseImRichSHRP)))
	poorMen           = sdk.NewCoins(sdk.NewCoin(denom.Base, sdk.NewInt(0)), sdk.NewCoin(denom.BaseUSD, sdk.NewInt(0)))
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

//func startInProcess(cfg network.Config, val *network.Validator) error {
//	logger := val.Ctx.Logger
//	tmCfg := val.Ctx.Config
//	tmCfg.Instrumentation.Prometheus = false
//
//	if err := val.AppConfig.ValidateBasic(); err != nil {
//		return err
//	}
//
//	nodeKey, err := p2p.LoadOrGenNodeKey(tmCfg.NodeKeyFile())
//	if err != nil {
//		return err
//	}
//
//	app := cfg.AppConstructor(*val)
//
//	genDocProvider := node.DefaultGenesisDocProviderFunc(tmCfg)
//	tmNode, err := node.NewNode(
//		tmCfg,
//		pvm.LoadOrGenFilePV(tmCfg.PrivValidatorKeyFile(), tmCfg.PrivValidatorStateFile()),
//		nodeKey,
//		proxy.NewLocalClientCreator(app),
//		genDocProvider,
//		node.DefaultDBProvider,
//		node.DefaultMetricsProvider(tmCfg.Instrumentation),
//		logger.With("module", val.Moniker),
//	)
//	if err != nil {
//		return err
//	}
//
//	if err := tmNode.Start(); err != nil {
//		return err
//	}
//
//	val.tmNode = tmNode
//
//	if val.RPCAddress != "" {
//		val.RPCClient = local.New(tmNode)
//	}
//
//	// We'll need a RPC client if the validator exposes a gRPC or REST endpoint.
//	if val.APIAddress != "" || val.AppConfig.GRPC.Enable {
//		val.ClientCtx = val.ClientCtx.
//			WithClient(val.RPCClient)
//
//		// Add the tx service in the gRPC router.
//		app.RegisterTxService(val.ClientCtx)
//
//		// Add the tendermint queries service in the gRPC router.
//		app.RegisterTendermintService(val.ClientCtx)
//	}
//
//	if val.APIAddress != "" {
//		apiSrv := api.New(val.ClientCtx, logger.With("module", "api-server"))
//		app.RegisterAPIRoutes(apiSrv, val.AppConfig.API)
//
//		errCh := make(chan error)
//
//		go func() {
//			if err := apiSrv.Start(*val.AppConfig); err != nil {
//				errCh <- err
//			}
//		}()
//
//		select {
//		case err := <-errCh:
//			return err
//		case <-time.After(srvtypes.ServerStartTime): // assume server started successfully
//		}
//
//		val.api = apiSrv
//	}
//
//	if val.AppConfig.GRPC.Enable {
//		grpcSrv, err := servergrpc.StartGRPCServer(val.ClientCtx, app, val.AppConfig.GRPC.Address)
//		if err != nil {
//			return err
//		}
//
//		val.grpc = grpcSrv
//
//		if val.AppConfig.GRPCWeb.Enable {
//			val.grpcWeb, err = servergrpc.StartGRPCWeb(grpcSrv, *val.AppConfig)
//			if err != nil {
//				return err
//			}
//		}
//	}
//
//	return nil
//}
//
//func collectGenFiles(cfg network.Config, vals []*Validator, outputDir string) error {
//	genTime := tmtime.Now()
//
//	for i := 0; i < cfg.NumValidators; i++ {
//		tmCfg := vals[i].Ctx.Config
//
//		nodeDir := filepath.Join(outputDir, vals[i].Moniker, "shareledgerd")
//		gentxsDir := filepath.Join(outputDir, "gentxs")
//
//		tmCfg.Moniker = vals[i].Moniker
//		tmCfg.SetRoot(nodeDir)
//
//		initCfg := genutiltypes.NewInitConfig(cfg.ChainID, gentxsDir, vals[i].NodeID, vals[i].PubKey)
//
//		genFile := tmCfg.GenesisFile()
//		genDoc, err := types.GenesisDocFromFile(genFile)
//		if err != nil {
//			return err
//		}
//
//		appState, err := genutil.GenAppStateFromConfig(cfg.Codec, cfg.TxConfig,
//			tmCfg, initCfg, *genDoc, banktypes.GenesisBalancesIterator{})
//		if err != nil {
//			return err
//		}
//
//		// overwrite each validator's genesis file to have a canonical genesis time
//		if err := genutil.ExportGenesisFileWithTime(genFile, cfg.ChainID, nil, appState, genTime); err != nil {
//			return err
//		}
//	}
//
//	return nil
//}
//
//func initGenFiles(cfg network.Config, genAccounts []authtypes.GenesisAccount, genBalances []banktypes.Balance, electoralGen electoraltypes.GenesisState, genFiles []string) error {
//
//	// set the Accounts in the genesis state
//	var authGenState authtypes.GenesisState
//	cfg.Codec.MustUnmarshalJSON(cfg.GenesisState[authtypes.ModuleName], &authGenState)
//
//	accounts, err := authtypes.PackAccounts(genAccounts)
//	if err != nil {
//		return err
//	}
//
//	authGenState.Accounts = append(authGenState.Accounts, accounts...)
//	cfg.GenesisState[authtypes.ModuleName] = cfg.Codec.MustMarshalJSON(&authGenState)
//
//	// set the balances in the genesis state
//	var bankGenState banktypes.GenesisState
//	cfg.Codec.MustUnmarshalJSON(cfg.GenesisState[banktypes.ModuleName], &bankGenState)
//
//	bankGenState.DenomMetadata = GetShareLedgerTestMetadata()
//	bankGenState.Balances = append(bankGenState.Balances, genBalances...)
//	//bankGenState.Supply = sdk.NewCoins(sdk.NewCoin("shr", sdk.NewInt(100000000)), sdk.NewCoin("shrp", sdk.NewInt(100000000)), sdk.NewCoin("node0token,", sdk.NewInt(1000000000)),sdk.NewCoin("node0token,", sdk.NewInt(1000000000)))
//	cfg.GenesisState[banktypes.ModuleName] = cfg.Codec.MustMarshalJSON(&bankGenState)
//
//	cfg.Codec.MustUnmarshalJSON(cfg.GenesisState[electoraltypes.ModuleName], &electoralGen)
//	cfg.GenesisState[electoraltypes.ModuleName] = cfg.Codec.MustMarshalJSON(&electoralGen)
//
//	appGenStateJSON, err := json.MarshalIndent(cfg.GenesisState, "", "  ")
//	if err != nil {
//		return err
//	}
//
//	genDoc := types.GenesisDoc{
//		ChainID:    cfg.ChainID,
//		AppState:   appGenStateJSON,
//		Validators: nil,
//	}
//
//	// generate empty genesis files for each validator and save
//	for i := 0; i < cfg.NumValidators; i++ {
//		if err := genDoc.SaveAs(genFiles[i]); err != nil {
//			return err
//		}
//	}
//
//	return nil
//}
//
//func writeFile(name string, dir string, contents []byte) error {
//	writePath := filepath.Join(dir)
//	file := filepath.Join(writePath, name)
//
//	err := tmos.EnsureDir(writePath, 0755)
//	if err != nil {
//		return err
//	}
//
//	err = tmos.WriteFile(file, contents, 0644)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}

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
