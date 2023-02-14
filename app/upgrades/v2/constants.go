package v2

import (
	"github.com/cosmos/cosmos-sdk/store/types"

	"github.com/cosmos/cosmos-sdk/x/group"
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
		Added: []string{distributionxType.StoreKey, group.StoreKey}},
}
