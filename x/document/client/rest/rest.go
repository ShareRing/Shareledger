package rest

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/gorilla/mux"
)

// RegisterRoutes - Central function to define routes that get registered by the main application
func RegisterRoutes(cliCtx client.Context, r *mux.Router, storeKey string) {
	r.HandleFunc(
		fmt.Sprintf("/%s/proof/{proof}", storeKey),
		QueryDocumentByProofRequestHandlerFn(cliCtx),
	).Methods("GET")

	r.HandleFunc(
		fmt.Sprintf("/%s/holder/{address}", storeKey),
		QueryDocumentByHolderIdRequestHandlerFn(cliCtx),
	).Methods("GET")
}
