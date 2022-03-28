package fee

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	feegrantmoduletypes "github.com/cosmos/cosmos-sdk/x/feegrant"
	govmoduletypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	stakingmoduletypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	assettypes "github.com/sharering/shareledger/x/asset/types"
	bookingtypes "github.com/sharering/shareledger/x/booking/types"
	"github.com/sharering/shareledger/x/constant"
	documenttypes "github.com/sharering/shareledger/x/document/types"
	electoralmoduletypes "github.com/sharering/shareledger/x/electoral/types"
	gentlemintmoduletypes "github.com/sharering/shareledger/x/gentlemint/types"
	idmoduletypes "github.com/sharering/shareledger/x/id/types"
	"reflect"
	"sync"
)

func init() {
	// init map action string _ msg type
	doOnce.Do(func() {
		mapMsg = make(map[string]reflect.Type)
		// We check fee of transactions based on mapActions' values. So it should not be duplicated.
		for msg, v := range mapActions {
			if _, found := mapMsg[v]; found {
				panic(fmt.Sprintf("map actions has duplicate values, %v.", v))
			}
			mapMsg[v] = msg
		}
	})
}

var mapMsg map[string]reflect.Type
var doOnce sync.Once

var mapActions = map[reflect.Type]string{
	reflect.ValueOf(&assettypes.MsgCreateAsset{}).Type(): "asset_create",
	reflect.ValueOf(&assettypes.MsgDeleteAsset{}).Type(): "asset_delete",
	reflect.ValueOf(&assettypes.MsgUpdateAsset{}).Type(): "asset_update",

	reflect.ValueOf(&banktypes.MsgSend{}).Type(): "bank_send",

	reflect.ValueOf(&bookingtypes.MsgCreateBooking{}).Type():   "booking_create",
	reflect.ValueOf(&bookingtypes.MsgCompleteBooking{}).Type(): "booking_complete",

	reflect.ValueOf(&distributiontypes.MsgWithdrawDelegatorReward{}).Type():     "distribution_withdraw-delegator-reward",
	reflect.ValueOf(&distributiontypes.MsgWithdrawValidatorCommission{}).Type(): "distribution_withdraw-validator-commission",
	reflect.ValueOf(&distributiontypes.MsgSetWithdrawAddress{}).Type():          "distribution_set-withdraw-address",
	reflect.ValueOf(&distributiontypes.MsgFundCommunityPool{}).Type():           "distribution_fund-community-pool",

	reflect.ValueOf(&documenttypes.MsgCreateDocument{}).Type():  "document_create",
	reflect.ValueOf(&documenttypes.MsgCreateDocuments{}).Type(): "documents_create",
	reflect.ValueOf(&documenttypes.MsgRevokeDocument{}).Type():  "document_revoke",
	reflect.ValueOf(&documenttypes.MsgUpdateDocument{}).Type():  "document_update",

	reflect.ValueOf(&electoralmoduletypes.MsgEnrollVoter{}).Type():            "electoral_enroll-voter",
	reflect.ValueOf(&electoralmoduletypes.MsgRevokeVoter{}).Type():            "electoral_revoke-voter",
	reflect.ValueOf(&electoralmoduletypes.MsgEnrollLoaders{}).Type():          "electoral_enroll-loaders",
	reflect.ValueOf(&electoralmoduletypes.MsgRevokeLoaders{}).Type():          "electoral_revoke-loaders",
	reflect.ValueOf(&electoralmoduletypes.MsgEnrollIdSigners{}).Type():        "electoral_enroll-id-signers",
	reflect.ValueOf(&electoralmoduletypes.MsgRevokeIdSigners{}).Type():        "electoral_revoke-id-signers",
	reflect.ValueOf(&electoralmoduletypes.MsgEnrollDocIssuers{}).Type():       "electoral_enroll-doc-issuers",
	reflect.ValueOf(&electoralmoduletypes.MsgRevokeDocIssuers{}).Type():       "electoral_revoke-doc-issuers",
	reflect.ValueOf(&electoralmoduletypes.MsgEnrollAccountOperators{}).Type(): "electoral_enroll-account-operators",
	reflect.ValueOf(&electoralmoduletypes.MsgRevokeAccountOperators{}).Type(): "electoral_revoke-account-operators",

	//feegrant
	reflect.ValueOf(&feegrantmoduletypes.MsgGrantAllowance{}).Type():  "feegrant_grant",
	reflect.ValueOf(&feegrantmoduletypes.MsgRevokeAllowance{}).Type(): "feegrant_revoke",

	//gentlemint
	reflect.ValueOf(&gentlemintmoduletypes.MsgLoad{}).Type():        "gentlemint_load",
	reflect.ValueOf(&gentlemintmoduletypes.MsgBuyShr{}).Type():      "gentlemint_buy-shr",
	reflect.ValueOf(&gentlemintmoduletypes.MsgSend{}).Type():        "gentlemint_send",
	reflect.ValueOf(&gentlemintmoduletypes.MsgBurn{}).Type():        "gentlemint_burn",
	reflect.ValueOf(&gentlemintmoduletypes.MsgSetExchange{}).Type(): "gentlemint_set-exchange",
	reflect.ValueOf(&gentlemintmoduletypes.MsgLoadFee{}).Type():     "gentlemint_load-fee",

	//gov
	reflect.ValueOf(&govmoduletypes.MsgDeposit{}).Type():        "gov_deposit",
	reflect.ValueOf(&govmoduletypes.MsgSubmitProposal{}).Type(): "gov_submit-proposal",
	reflect.ValueOf(&govmoduletypes.MsgVote{}).Type():           "gov_vote",
	reflect.ValueOf(&govmoduletypes.MsgVoteWeighted{}).Type():   "gov_weighted-vote",

	//ids
	reflect.ValueOf(&idmoduletypes.MsgCreateId{}).Type():       "id_create",
	reflect.ValueOf(&idmoduletypes.MsgUpdateId{}).Type():       "id_update",
	reflect.ValueOf(&idmoduletypes.MsgReplaceIdOwner{}).Type(): "id_replace",
	reflect.ValueOf(&idmoduletypes.MsgCreateIds{}).Type():      "id_create-ids",

	//staking
	reflect.ValueOf(&stakingmoduletypes.MsgCreateValidator{}).Type(): "staking_create-validator",
	reflect.ValueOf(&stakingmoduletypes.MsgDelegate{}).Type():        "staking_delegate",
	reflect.ValueOf(&stakingmoduletypes.MsgEditValidator{}).Type():   "staking_edit-validator",
	reflect.ValueOf(&stakingmoduletypes.MsgBeginRedelegate{}).Type(): "staking_redelegate",
	reflect.ValueOf(&stakingmoduletypes.MsgUndelegate{}).Type():      "staking_unbond",
}

func GetActionKey(msg sdk.Msg) string {
	k := mapActions[reflect.TypeOf(msg)]
	return k
}

func GetListActions() []string {
	actions := make([]string, 0, len(mapActions))
	for _, v := range mapActions {
		actions = append(actions, v)
	}
	return actions
}

func GetListActionsWithDefaultLevel() map[string]string {
	m := make(map[string]string)
	for _, v := range mapActions {
		m[v] = string(constant.MinFee)
	}
	return m
}

func IsSpecialActionKey(actionKey string) bool {
	return actionKey == "staking_create-validator" ||
		actionKey == "gentlemint_load-fee"
}

func HaveActionKey(actionKey string) bool {
	_, found := mapMsg[actionKey]
	return found
}
