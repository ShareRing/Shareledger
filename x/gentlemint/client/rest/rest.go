package rest

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"

	"github.com/gorilla/mux"
)

func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, storeName string) {
	r.HandleFunc(fmt.Sprintf("/%s", storeName), nil).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/exchange", storeName), QueryExchangeRateHandleFn(cliCtx)).Methods("GET")
}
