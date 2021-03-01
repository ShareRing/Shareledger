/**
 * Based on https://github.com/cosmos/cosmos-sdk/blob/master/x/distribution/keeper
 */

package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authexported "github.com/cosmos/cosmos-sdk/x/auth/exported"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/cosmos/cosmos-sdk/x/supply"
)

func MakeTestCodec() *codec.Codec {
	var cdc = codec.New()

	// Register Msgs
	cdc.RegisterInterface((*sdk.Msg)(nil), nil)
	cdc.RegisterConcrete(bank.MsgSend{}, "test/staking/Send", nil)
	cdc.RegisterConcrete(types.MsgCreateValidator{}, "test/staking/CreateValidator", nil)
	cdc.RegisterConcrete(types.MsgEditValidator{}, "test/staking/EditValidator", nil)
	cdc.RegisterConcrete(types.MsgUndelegate{}, "test/staking/Undelegate", nil)
	cdc.RegisterConcrete(types.MsgBeginRedelegate{}, "test/staking/BeginRedelegate", nil)

	// Register AppAccount
	cdc.RegisterInterface((*authexported.Account)(nil), nil)
	cdc.RegisterConcrete(&auth.BaseAccount{}, "test/staking/BaseAccount", nil)
	supply.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)

	return cdc
}
