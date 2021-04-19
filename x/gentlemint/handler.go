package gentlemint

import (
	"fmt"

	shareringUtils "github.com/ShareRing/modules/utils"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/x/gentlemint/types"
)

const (
	ShrpLoaderPrefix = "shrploader"
)

var (
	requiredSHRAmt = shareringUtils.SHRDecimal.Mul(sdk.NewInt(10))
)

func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		switch msg := msg.(type) {
		case MsgLoadSHR:
			return handleMsgLoadSHR(ctx, keeper, msg)
		case MsgLoadSHRP:
			return handleMsgLoadSHRP(ctx, keeper, msg)
		case MsgSendSHRP:
			return handleMsgSendSHRP(ctx, keeper, msg)
		case MsgSendSHR:
			return handleMsgSendSHR(ctx, keeper, msg)
		// case MsgBuyCent:
		// 	return handleMsgBuyCent(ctx, keeper, msg)
		case MsgBurnSHRP:
			return handleMsgBurnSHRP(ctx, keeper, msg)
		case MsgBurnSHR:
			return handleMsgBurnSHR(ctx, keeper, msg)
		case MsgEnrollSHRPLoader:
			return handleMsgEnrollSHRPLoader(ctx, keeper, msg)
		case MsgRevokeSHRPLoader:
			return handleMsgRevokeSHRPLoader(ctx, keeper, msg)
		case MsgBuySHR:
			return handleMsgBuySHR(ctx, keeper, msg)
		case MsgSetExchange:
			return handleMsgSetExchange(ctx, keeper, msg)
		case MsgEnrollIDSigners:
			return handleMsgEnrollIdSigners(ctx, keeper, msg)
		case MsgRevokeIDSigners:
			return handleMsgRevokeIdSigners(ctx, keeper, msg)
		case MsgEnrollDocIssuers:
			return handleMsgEnrollDocummentIssuer(ctx, keeper, msg)
		case MsgRevokeDocIssuers:
			return handleMsgRevokeDocumentIssuers(ctx, keeper, msg)
		case MsgEnrollAccOperators:
			return handleMsgEnrollAccountOperator(ctx, keeper, msg)
		case MsgRevokeAccOperators:
			return handleMsgRevokeAccountOperator(ctx, keeper, msg)

		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, fmt.Sprintf("Unrecognized gentlemint Msg type: %v", msg.Type()))
		}
	}
}

func handleMsgLoadSHR(ctx sdk.Context, keeper Keeper, msg MsgLoadSHR) (*sdk.Result, error) {
	if !IsAuthority(ctx, msg.GetSigners()[0], keeper) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Approver's Address is not authority")
	}

	amt, _ := sdk.NewIntFromString(msg.Amount)
	if !keeper.ShrMintPossible(ctx, amt) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "SHR possible mint exceeded")
	}
	coins := sdk.NewCoins(sdk.NewCoin("shr", amt))
	if err := keeper.LoadCoins(ctx, msg.Receiver, coins); err != nil {
		return nil, err
	}
	log := fmt.Sprintf("Successfully loaded shr {address: %s, amount %v}", msg.Receiver.String(), coins)
	return &sdk.Result{
		Log: log,
	}, nil
}

// Load SHRP function is used to load the given amount of SHRP to the given recepient
// - Tx fee is 0.01 SHRP
// - Automatically buy 10SHR for the recepient
// - Send 1SHR from recepient to loader as the loading fee
func handleMsgLoadSHRP(ctx sdk.Context, keeper Keeper, msg MsgLoadSHRP) (*sdk.Result, error) {
	if !IsSHRPLoader(ctx, msg.GetSigners()[0], keeper) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Approver's Address is not an Enrolled SHRP Loader")
	}
	// i, d, err := types.ParseCoinStr(msg.Amount)
	// if err != nil {
	// 	return nil, err
	// }
	// shrpAmt := sdk.NewInt(i)
	// centAmt := sdk.NewInt(d)

	// amt := sdk.NewCoins(sdk.NewCoin("cent", centAmt), sdk.NewCoin("shrp", shrpAmt))
	// if err := keeper.LoadCoins(ctx, msg.Receiver, amt); err != nil {
	// 	return nil, err
	// }

	shrpAmt, _ := sdk.NewIntFromString(msg.Amount)
	amt := sdk.NewCoins(sdk.NewCoin("shrp", shrpAmt))
	if err := keeper.LoadCoins(ctx, msg.Receiver, amt); err != nil {
		return nil, err
	}

	oldCoins := keeper.GetCoins(ctx, msg.Receiver)
	oldShr := oldCoins.AmountOf("shr")
	// if there is less that 10 shr in the wallet, buy 10 shr
	if oldShr.LT(requiredSHRAmt) {
		err := keeper.BuyShr(ctx, requiredSHRAmt, msg.Receiver)
		if err != nil {
			return nil, sdkerrors.Wrap(err, "Can not buy SHR")
		}
	}
	// return 1 SHR fee spent by the loader
	reimbursed := sdk.NewCoin("shr", shareringUtils.OneSHR)

	if err := keeper.SendCoins(ctx, msg.Receiver, msg.Approver, sdk.NewCoins(reimbursed)); err != nil {
		return nil, err
	}
	log := fmt.Sprintf("Successfully load SHRP {amount %s, address: %s}", msg.Amount, msg.Receiver.String())
	return &sdk.Result{
		Log: log,
	}, nil
}

func handleMsgSendSHRP(ctx sdk.Context, keeper Keeper, msg MsgSendSHRP) (*sdk.Result, error) {
	// if err := msg.ValidateBasic(); err != nil {
	// 	return nil, err
	// }
	// i, d, err := types.ParseCoinStr(msg.Amount)
	// if err != nil {
	// 	return nil, err
	// }
	// shrpAmt := sdk.NewInt(i)
	// centAmt := sdk.NewInt(d)
	// oldCoins := keeper.GetCoins(ctx, msg.Sender)
	// if oldCoins.AmountOf("cent").LT(centAmt) {
	// 	if oldCoins.AmountOf("shrp").LTE(shrpAmt) {
	// 		return nil, sdkerrors.ErrInsufficientFunds
	// 	}
	// 	if _, err := keeper.SubtractCoins(ctx, msg.Sender, sdk.NewCoins(sdk.NewCoin("shrp", sdk.NewInt(int64(1))))); err != nil {
	// 		return nil, err
	// 	}
	// 	if _, err := keeper.AddCoins(ctx, msg.Sender, sdk.NewCoins(sdk.NewCoin("cent", sdk.NewInt(int64(100))))); err != nil {
	// 		return nil, err
	// 	}
	// 	if err := keeper.SupplyBurnCoins(ctx, sdk.NewCoins(sdk.NewCoin("shrp", sdk.NewInt(int64(1))))); err != nil {
	// 		return nil, err
	// 	}
	// 	if err := keeper.SupplyMintCoins(ctx, sdk.NewCoins(sdk.NewCoin("cent", sdk.NewInt(int64(100))))); err != nil {
	// 		return nil, err
	// 	}
	// }
	shrpAmt, _ := sdk.NewIntFromString(msg.Amount)
	amt := sdk.NewCoins(sdk.NewCoin("shrp", shrpAmt))
	if err := keeper.SendCoins(ctx, msg.Sender, msg.Receiver, amt); err != nil {
		return nil, err
	}
	log := fmt.Sprintf("Successfully Send SHRP {amount %s, from: %s, to: %s}", msg.Amount, msg.Sender.String(), msg.Receiver.String())
	return &sdk.Result{
		Log: log,
	}, nil
}

// Automatically buy SHR using SHRP if the SHR balane is not enough
func handleMsgSendSHR(ctx sdk.Context, keeper Keeper, msg MsgSendSHR) (*sdk.Result, error) {
	// if err := msg.ValidateBasic(); err != nil {
	// 	return nil, err
	// }
	// shrAmt := sdk.NewInt(int64(msg.Amount))
	shrAmt, _ := sdk.NewIntFromString(msg.Amount)
	oldCoins := keeper.GetCoins(ctx, msg.Sender)
	if oldCoins.AmountOf("shr").LT(shrAmt) {
		shrToBuy := oldCoins.AmountOf("shr").Sub(shrAmt)
		if err := keeper.BuyShr(ctx, shrToBuy, msg.Sender); err != nil {
			return nil, err
		}
	}
	amt := sdk.NewCoins(sdk.NewCoin("shr", shrAmt))
	if err := keeper.SendCoins(ctx, msg.Sender, msg.Receiver, amt); err != nil {
		return nil, err
	}
	log := fmt.Sprintf("Successfully Send SHR {amount %d, from: %s, to: %s}", msg.Amount, msg.Sender.String(), msg.Receiver.String())
	return &sdk.Result{
		Log: log,
	}, nil
}

func handleMsgBuyCent(ctx sdk.Context, keeper Keeper, msg MsgBuyCent) (*sdk.Result, error) {
	err := msg.ValidateBasic()
	if err != nil {
		return nil, err
	}
	shrpAmt := sdk.NewInt(int64(msg.Amount))
	shrpCoins := sdk.NewCoins(sdk.NewCoin("shrp", shrpAmt))
	centAmt := sdk.NewInt(int64(msg.Amount * 100))
	centCoins := sdk.NewCoins(sdk.NewCoin("cent", centAmt))
	_, err = keeper.SubtractCoins(ctx, msg.Buyer, shrpCoins)
	if err != nil {
		return nil, err
	}
	_, err = keeper.AddCoins(ctx, msg.Buyer, centCoins)
	if err != nil {
		return nil, err
	}
	err = keeper.SupplyBurnCoins(ctx, shrpCoins)
	if err != nil {
		return nil, err
	}
	err = keeper.SupplyMintCoins(ctx, centCoins)
	if err != nil {
		return nil, err
	}
	log := fmt.Sprintf("Successfull exchange %d shrp to cent for address %s", msg.Amount, msg.Buyer.String())
	return &sdk.Result{
		Log: log,
	}, nil
}

func handleMsgBuySHR(ctx sdk.Context, keeper Keeper, msg MsgBuySHR) (*sdk.Result, error) {
	shrAmt, _ := sdk.NewIntFromString(msg.Amount)

	if err := keeper.BuyShr(ctx, shrAmt, msg.Buyer); err != nil {
		return nil, err
	}

	log := fmt.Sprintf("Successfull buy %d shr for address %s", msg.Amount, msg.Buyer.String())
	return &sdk.Result{
		Log: log,
	}, nil
}

func handleMsgBurnSHRP(ctx sdk.Context, keeper Keeper, msg MsgBurnSHRP) (*sdk.Result, error) {
	if !IsTreasurer(ctx, msg.GetSigners()[0], keeper) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Approver's Address is not Treasurer")
	}
	// i, d, err := ParseCoinStr(msg.Amount)
	// if err != nil {
	// 	return nil, err
	// }

	// shrpAmt := sdk.NewInt(i)
	// centAmt := sdk.NewInt(d)
	// amt := sdk.NewCoins(sdk.NewCoin("cent", centAmt), sdk.NewCoin("shrp", shrpAmt))
	shrpAmt, _ := sdk.NewIntFromString(msg.Amount)
	amt := sdk.NewCoins(sdk.NewCoin("shrp", shrpAmt))
	if err := keeper.BurnCoins(ctx, msg.Approver, amt); err != nil {
		return nil, err
	}
	log := fmt.Sprintf("Successfully burn coins %s", msg.Amount)
	return &sdk.Result{
		Log: log,
	}, nil
}

func handleMsgBurnSHR(ctx sdk.Context, keeper Keeper, msg MsgBurnSHR) (*sdk.Result, error) {
	if !IsTreasurer(ctx, msg.GetSigners()[0], keeper) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Approver's Address is not Treasurer")
	}
	// if err := msg.ValidateBasic(); err != nil {
	// 	return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	// }
	shrAmt, _ := sdk.NewIntFromString(msg.Amount)
	amt := sdk.NewCoins(sdk.NewCoin("shr", shrAmt))
	if err := keeper.BurnCoins(ctx, msg.Approver, amt); err != nil {
		return nil, err
	}
	log := fmt.Sprintf("Successfully burn %d shr", msg.Amount)
	return &sdk.Result{
		Log: log,
	}, nil
}

func handleMsgEnrollSHRPLoader(ctx sdk.Context, keeper Keeper, msg MsgEnrollSHRPLoader) (*sdk.Result, error) {
	if !IsAuthority(ctx, msg.GetSigners()[0], keeper) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Approver's Address is not authority")
	}
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	log := "SHRP loaders' addresses: "
	loaderAllowance := sdk.NewCoins(sdk.NewCoin("shr", sdk.NewInt(int64(20))))
	for _, addr := range msg.SHRPLoaders {
		log = log + "," + addr.String()
		loaderKey := fmt.Sprintf("%s%s", ShrpLoaderPrefix, addr.String())
		keeper.SetSHRPLoaderStatus(ctx, loaderKey, types.StatusSHRPLoaderActived)
		if err := keeper.LoadCoins(ctx, addr, loaderAllowance); err != nil {
			return nil, err
		}
	}
	log = fmt.Sprintf("Successfully enroll SHRP loader %s", log)
	return &sdk.Result{
		Log: log,
	}, nil
}

func handleMsgRevokeSHRPLoader(ctx sdk.Context, keeper Keeper, msg MsgRevokeSHRPLoader) (*sdk.Result, error) {
	if !IsAuthority(ctx, msg.GetSigners()[0], keeper) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Approver's Address is not authority")
	}
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	log := "SHRP loaders' addresses: "
	for _, addr := range msg.SHRPLoaders {
		log = log + "," + addr.String()
		loaderKey := fmt.Sprintf("%s%s", ShrpLoaderPrefix, addr.String())
		keeper.SetSHRPLoaderStatus(ctx, loaderKey, types.StatusSHRPLoaderInactived)
	}
	log = fmt.Sprintf("Successfully revoke SHRP loader %s", log)
	return &sdk.Result{
		Log: log,
	}, nil
}

func handleMsgSetExchange(ctx sdk.Context, k Keeper, msg MsgSetExchange) (*sdk.Result, error) {
	if !IsTreasurer(ctx, msg.Approver, k) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Only treasurer can set exchange rate")
	}
	k.SetExchangeRate(ctx, msg.Rate)
	return &sdk.Result{
		Log: fmt.Sprintf("Successfully set exchange rate as %s", msg.Rate),
	}, nil
}

func IsAuthority(ctx sdk.Context, address sdk.AccAddress, k Keeper) bool {
	authority := k.GetAuthorityAccount(ctx)
	if authority == address.String() {
		return true
	}

	return false
}

func IsTreasurer(ctx sdk.Context, address sdk.AccAddress, k Keeper) bool {
	treasurer := k.GetTreasurerAccount(ctx)

	if treasurer == address.String() {
		return true
	}
	return false
}

func IsSHRPLoader(ctx sdk.Context, address sdk.AccAddress, k Keeper) bool {
	addr := ShrpLoaderPrefix + address.String()
	status := k.GetSHRPLoaderStatus(ctx, addr)
	return status == types.StatusSHRPLoaderActived
}
