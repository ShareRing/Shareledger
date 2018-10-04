package app

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"bitbucket.org/shareringvn/cosmos-sdk/wire"
	"github.com/tendermint/go-crypto"
	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/sharering/shareledger/types"
)

func TestMakeGenesisFile(t *testing.T) {
	pubKey, _ := types.GenerateKeyPair()
	abciPubKey := pubKey.ToABCIPubKey()

	cdc := MakeCodec()
	registerAmino(cdc)

	// These are all written here instead of
	pubKeyJSON, err := cdc.MarshalJSON(abciPubKey)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Json: %s\n", pubKeyJSON)

	pubKeyJSON, err = cdc.MarshalJSON(pubKey)
	if err != nil {
		panic(err)
	}

	fmt.Printf("JSON: %s\n", pubKeyJSON)

	owner := pubKey.Address()

	appState := json.RawMessage(`{
		"stake":{ 
		 "pool": {
		   "loose_tokens": "100",
		   "bonded_tokens": "50",    
		   "date_last_commission_reset": 0,
		   "prev_bonded_shares": "0"
		 }, 
		 "params": {       
		   "goal_bonded": "67/100",
		   "unbonding_time": 10000,
		   "max_validators": 200,
		   "bond_denom": "SHR"
		 },
		   "validators":[           
			  {  
			   "owner":"405C725BC461DCA455B8AA84769E8ACE6B3763F4", 
			   "pub_key":{  
					"type":"88B6D5D73C58D0",
					"value":"BEKZrPS2oJw28meokkVZtZ+gbF0+Kl38BOg4sBVGxhIwKnzhATQeSI4vVyzZcYMUdZsX4i92C4yyxw2d5WnEwaE="
			},
			 
				 "revoked":false,
				 "delegator_shares":"100",
				 "description":{  
					"moniker":"TanDo",
					"identity":"",
					"website":"",
					"details":""
				 },
				 "bond_height":0,
				 "bond_intra_tx_counter":0
			  }
		   ]
		}
	 }`)

	genesisDoc := tmtypes.GenesisDoc{
		GenesisTime:     time.Now(),
		ChainID:         "chain-trang-test",
		ConsensusParams: tmtypes.DefaultConsensusParams(),
		Validators: []tmtypes.GenesisValidator{
			tmtypes.GenesisValidator{
				PubKey: abciPubKey,
				Power:  10,
				Name:   "TestValidator",
			},
		},
		AppHash:      []byte(""),
		AppStateJSON: appState,
	}

	fmt.Printf("%v\n", genesisDoc)

}
