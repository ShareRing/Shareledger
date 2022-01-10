package keeper

import "github.com/cosmos/cosmos-sdk/types"

type ActionsTable interface {
	GetActionKey(types.Msg) string
	GetListActionKeys() []string
	Have(actionKey string) bool
}
