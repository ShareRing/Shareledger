package v3

import (
	"github.com/cosmos/cosmos-sdk/store/types"

	"github.com/sharering/shareledger/app/upgrades"
)

const (
	UpgradeName = "v3"
)

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades:        types.StoreUpgrades{}, // no added/renamed/deleted stores
}