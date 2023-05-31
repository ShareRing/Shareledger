package v2

import (
	"github.com/cosmos/cosmos-sdk/store/types"

	consensustypes "github.com/cosmos/cosmos-sdk/x/consensus/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	"github.com/cosmos/cosmos-sdk/x/group"
	icacontrollertypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/controller/types"
	"github.com/sharering/shareledger/app/upgrades"
	distributionxType "github.com/sharering/shareledger/x/distributionx/types"
)

const (
	UpgradeName = "v2"
)

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: types.StoreUpgrades{
		Added: []string{
			distributionxType.StoreKey,
			group.StoreKey,
			icacontrollertypes.StoreKey,
			consensustypes.StoreKey,
			crisistypes.StoreKey,
		},
	},
}
