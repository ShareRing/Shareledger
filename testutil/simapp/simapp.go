package simapp

import (
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/sharering/shareledger/app"
	"github.com/tendermint/tendermint/libs/log"
	tmdb "github.com/tendermint/tm-db"
)

// New creates application instance with in-memory database and disabled logging.
func New(dir string, appOpts servertypes.AppOptions) servertypes.Application {
	db := tmdb.NewMemDB()
	logger := log.NewNopLogger()

	encoding := app.MakeTestEncodingConfig()
	a := app.New(logger, db, nil, true, make(map[int64]bool), dir, 0, encoding, appOpts)
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
