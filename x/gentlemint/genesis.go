package gentlemint

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/gentlemint/types"
)

type GenesisState struct {
	LoaderKeys   []string `json:"loader_keys"`
	ExchangeRate string   `json:"exchange_rate"`

	Authority        string           `json:"authority"`
	Treasurer        string           `json:"treasurer"`
	IdSigners        []types.AccState `json:"id_signers"`
	DocumentIssuer   []types.AccState `json:"document_issuer"`
	AccountOperators []types.AccState `json:"account_operators"`
}

func NewGenesisState() GenesisState {
	return GenesisState{}
}

func ValidateGenesis(data GenesisState) error {
	return nil
}

func DefaultGenesisState() GenesisState {
	return GenesisState{}
}

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) {
	for _, loaderKey := range data.LoaderKeys {
		keeper.SetSHRPLoaderStatus(ctx, loaderKey, types.StatusSHRPLoaderActived)
	}
	if data.ExchangeRate != "" {
		keeper.SetExchangeRate(ctx, data.ExchangeRate)
	}

	keeper.SetAuthorityAccount(ctx, data.Authority)
	keeper.SetTreasurerAccount(ctx, data.Treasurer)

	for _, acc := range data.IdSigners {
		keeper.SetIdSigner(ctx, acc)
	}

	for _, acc := range data.DocumentIssuer {
		keeper.SetDocIssuer(ctx, acc)
	}

	for _, acc := range data.AccountOperators {
		keeper.SetAccOp(ctx, acc)
	}
}

func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	var loaderKeys []string
	cb := func(loaderKey string, loader types.SHRPLoader) bool {
		if loader.Status == types.StatusSHRPLoaderActived {
			loaderKeys = append(loaderKeys, loaderKey)
		}
		return false
	}

	var idSigners []types.AccState
	idSignerCb := func(idSigner types.AccState) bool {
		idSigners = append(idSigners, idSigner)
		return false
	}

	var issuers []types.AccState
	issuerCb := func(issuer types.AccState) bool {
		issuers = append(issuers, issuer)
		return false
	}

	var operators []types.AccState
	operatorCb := func(operator types.AccState) bool {
		operators = append(operators, operator)
		return false
	}

	k.IterateSHRPLoaders(ctx, cb)
	k.IterateIdSigners(ctx, idSignerCb)
	k.IterateDocIssuers(ctx, issuerCb)
	k.IterateAccOps(ctx, operatorCb)

	exchangeRate := k.GetExchangeRate(ctx)
	authorityAcc := k.GetAuthorityAccount(ctx)
	treasurer := k.GetTreasurerAccount(ctx)

	return GenesisState{
		LoaderKeys:       loaderKeys,
		ExchangeRate:     exchangeRate,
		Authority:        authorityAcc,
		Treasurer:        treasurer,
		IdSigners:        idSigners,
		DocumentIssuer:   issuers,
		AccountOperators: operators,
	}
}

func GetGenesisStateFromAppState(cdc *codec.Codec, appState map[string]json.RawMessage) GenesisState {
	var genesisState GenesisState
	if appState[ModuleName] != nil {
		cdc.MustUnmarshalJSON(appState[ModuleName], &genesisState)
	}

	return genesisState
}
