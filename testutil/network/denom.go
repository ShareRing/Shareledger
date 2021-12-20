package network

import "github.com/cosmos/cosmos-sdk/x/bank/types"

func GetShareLedgerTestMetadata() []types.Metadata {
	return []types.Metadata{
		{
			Name:        "Sharering token",
			Symbol:      "SHR",
			Description: "The native token in shareLedger",
			DenomUnits: []*types.DenomUnit{
				{"shr", uint32(6), nil},
			},
			Base:    "shr",
			Display: "shr",
		},
		{
			Name:        "Sharering token stable coin",
			Symbol:      "SHRP",
			Description: "The stable token in shareLedger",
			DenomUnits: []*types.DenomUnit{
				{"shrp", uint32(6), nil},
			},
			Base:    "shrp",
			Display: "shrp",
		},
	}
}
