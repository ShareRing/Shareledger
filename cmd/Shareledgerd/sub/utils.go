package sub

import (
	"errors"
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	stakingcli "github.com/cosmos/cosmos-sdk/x/staking/client/cli"
	"github.com/cosmos/cosmos-sdk/x/staking/types"
	flag "github.com/spf13/pflag"
)

type newGenerateOrBroadcastFunc func(client.Context, *flag.FlagSet, ...sdk.Msg) error

func newBuildCreateValidatorMsg(clientCtx client.Context, txf tx.Factory, fs *flag.FlagSet) (tx.Factory, *types.MsgCreateValidator, error) {
	fAmount, _ := fs.GetString(stakingcli.FlagAmount)
	amount, err := sdk.ParseCoinNormalized(fAmount)
	if err != nil {
		return txf, nil, err
	}

	valAddr := clientCtx.GetFromAddress()
	pkStr, err := fs.GetString(stakingcli.FlagPubKey)
	if err != nil {
		return txf, nil, err
	}

	var pk cryptotypes.PubKey
	if err := clientCtx.Codec.UnmarshalInterfaceJSON([]byte(pkStr), &pk); err != nil {
		return txf, nil, err
	}

	moniker, _ := fs.GetString(stakingcli.FlagMoniker)
	identity, _ := fs.GetString(stakingcli.FlagIdentity)
	website, _ := fs.GetString(stakingcli.FlagWebsite)
	security, _ := fs.GetString(stakingcli.FlagSecurityContact)
	details, _ := fs.GetString(stakingcli.FlagDetails)
	description := types.NewDescription(
		moniker,
		identity,
		website,
		security,
		details,
	)

	// get the initial validator commission parameters
	rateStr, _ := fs.GetString(stakingcli.FlagCommissionRate)
	maxRateStr, _ := fs.GetString(stakingcli.FlagCommissionMaxRate)
	maxChangeRateStr, _ := fs.GetString(stakingcli.FlagCommissionMaxChangeRate)

	commissionRates, err := buildCommissionRates(rateStr, maxRateStr, maxChangeRateStr)
	if err != nil {
		return txf, nil, err
	}

	// get the initial validator min self delegation
	msbStr, _ := fs.GetString(stakingcli.FlagMinSelfDelegation)

	minSelfDelegation, ok := sdk.NewIntFromString(msbStr)
	if !ok {
		return txf, nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "minimum self delegation must be a positive integer")
	}

	msg, err := types.NewMsgCreateValidator(
		sdk.ValAddress(valAddr), pk, amount, description, commissionRates, minSelfDelegation,
	)
	if err != nil {
		return txf, nil, err
	}
	if err := msg.ValidateBasic(); err != nil {
		return txf, nil, err
	}

	genOnly, _ := fs.GetBool(flags.FlagGenerateOnly)
	if genOnly {
		ip, _ := fs.GetString(stakingcli.FlagIP)
		nodeID, _ := fs.GetString(stakingcli.FlagNodeID)

		if nodeID != "" && ip != "" {
			txf = txf.WithMemo(fmt.Sprintf("%s@%s:26656", nodeID, ip))
		}
	}

	return txf, msg, nil
}

func buildCommissionRates(rateStr, maxRateStr, maxChangeRateStr string) (commission types.CommissionRates, err error) {
	if rateStr == "" || maxRateStr == "" || maxChangeRateStr == "" {
		return commission, errors.New("must specify all validator commission parameters")
	}

	rate, err := sdk.NewDecFromStr(rateStr)
	if err != nil {
		return commission, err
	}

	maxRate, err := sdk.NewDecFromStr(maxRateStr)
	if err != nil {
		return commission, err
	}

	maxChangeRate, err := sdk.NewDecFromStr(maxChangeRateStr)
	if err != nil {
		return commission, err
	}

	commission = types.NewCommissionRates(rate, maxRate, maxChangeRate)

	return commission, nil
}

func flagSetDescriptionEdit() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)

	fs.String(stakingcli.FlagMoniker, types.DoNotModifyDesc, "The validator's name")
	fs.String(stakingcli.FlagIdentity, types.DoNotModifyDesc, "The (optional) identity signature (ex. UPort or Keybase)")
	fs.String(stakingcli.FlagWebsite, types.DoNotModifyDesc, "The validator's (optional) website")
	fs.String(stakingcli.FlagSecurityContact, types.DoNotModifyDesc, "The validator's (optional) security contact email")
	fs.String(stakingcli.FlagDetails, types.DoNotModifyDesc, "The validator's (optional) details")

	return fs
}

func flagSetCommissionUpdate() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)

	fs.String(stakingcli.FlagCommissionRate, "", "The new commission rate percentage")

	return fs
}

func flagSetDescriptionCreate() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)

	fs.String(stakingcli.FlagMoniker, "", "The validator's name")
	fs.String(stakingcli.FlagIdentity, "", "The optional identity signature (ex. UPort or Keybase)")
	fs.String(stakingcli.FlagWebsite, "", "The validator's (optional) website")
	fs.String(stakingcli.FlagSecurityContact, "", "The validator's (optional) security contact email")
	fs.String(stakingcli.FlagDetails, "", "The validator's (optional) details")

	return fs
}

func newSplitAndApply(
	genOrBroadcastFn newGenerateOrBroadcastFunc, clientCtx client.Context,
	fs *flag.FlagSet, msgs []sdk.Msg, chunkSize int,
) error {

	if chunkSize == 0 {
		return genOrBroadcastFn(clientCtx, fs, msgs...)
	}

	// split messages into slices of length chunkSize
	totalMessages := len(msgs)
	for i := 0; i < len(msgs); i += chunkSize {

		sliceEnd := i + chunkSize
		if sliceEnd > totalMessages {
			sliceEnd = totalMessages
		}

		msgChunk := msgs[i:sliceEnd]
		if err := genOrBroadcastFn(clientCtx, fs, msgChunk...); err != nil {
			return err
		}
	}

	return nil
}

func getPeriodReset(duration int64) time.Time {
	return time.Now().Add(getPeriod(duration))
}

func getPeriod(duration int64) time.Duration {
	return time.Duration(duration) * time.Second
}
