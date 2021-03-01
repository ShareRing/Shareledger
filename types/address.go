package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	PrefixValidator = "val"

	PrefixConsensus = "cons"

	PrefixPublic = "pub"

	PrefixOperator = "oper"

	// Bech32 prefix

	Bech32MainPrefix = "shareledger"

	Bech32PrefixAccAddr = Bech32MainPrefix

	Bech32PrefixAccPub = Bech32MainPrefix + PrefixPublic

	Bech32PrefixValAddr = Bech32MainPrefix + PrefixValidator + PrefixOperator

	Bech32PrefixValPub = Bech32MainPrefix + PrefixValidator + PrefixOperator + PrefixPublic

	Bech32PrefixConsAddr = Bech32MainPrefix + PrefixValidator + PrefixConsensus

	Bech32PrefixConsPub = Bech32MainPrefix + PrefixValidator + PrefixConsensus + PrefixPublic
)

func ConfigureSDK() {
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(Bech32PrefixAccAddr, Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(Bech32PrefixValAddr, Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(Bech32PrefixConsAddr, Bech32PrefixConsPub)
	config.Seal()
}
