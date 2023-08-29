package simapp

import (
	tmdb "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/libs/log"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"

	"github.com/sharering/shareledger/app"
)

// New creates application instance with in-memory database and disabled logging.
func New(appOpts servertypes.AppOptions) servertypes.Application {
	db := tmdb.NewMemDB()
	logger := log.NewNopLogger()

	a := app.New(logger, db, nil, true, appOpts)
	// InitChain updates deliverState which is required when app.NewContext is called
	return a
}

//var defaultConsensusParams = &abci.ConsensusParams{
//	Block: &abci.BlockParams{
//		MaxBytes: 200000,
//		MaxGas:   2000000,
//	},
//	Evidence: &tmproto.EvidenceParams{
//		MaxAgeNumBlocks: 302400,
//		MaxAgeDuration:  504 * time.Hour, // 3 weeks is the max duration
//		MaxBytes:        10000,
//	},
//	Validator: &tmproto.ValidatorParams{
//		PubKeyTypes: []string{
//			tmtypes.ABCIPubKeyTypeEd25519,
//		},
//	},
//}
