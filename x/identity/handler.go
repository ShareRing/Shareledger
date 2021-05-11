package identity

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/x/identity/types"
)

func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, fmt.Sprintf("Unrecognized identity Msg type: %v [Deprecated: Old module]", msg.Type()))
		// switch msg := msg.(type) {
		// case MsgCreateId:
		// 	return handleMsgCreateId(ctx, keeper, msg)
		// case MsgUpdateId:
		// 	return handleMsgUpdateId(ctx, keeper, msg)
		// case MsgDeleteId:
		// 	return handleMsgDeleteId(ctx, keeper, msg)
		// case MsgEnrollIDSigners:
		// 	return handleMsgEnrollIdSigners(ctx, keeper, msg)
		// case MsgRevokeIDSigners:
		// 	return handleMsgRevokeIdSigners(ctx, keeper, msg)
		// default:
		// 	return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, fmt.Sprintf("Unrecognized identity Msg type: %v", msg.Type()))
		// }
	}
}

func handleMsgCreateId(ctx sdk.Context, keeper Keeper, msg MsgCreateId) (*sdk.Result, error) {
	if !IsIdSigner(ctx, msg.GetSigners()[0], keeper) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Approver's Address is not an Enrolled Id signer")
	}

	addr := IdPrefix + msg.Owner.String()
	if keeper.IsIdPresent(ctx, addr) {
		return nil, types.ErrIdAlreadyExists
	}

	keeper.SetId(ctx, addr, msg.Hash)
	event := sdk.NewEvent(
		EventTypeIDCreate,
		sdk.NewAttribute(AttributeAddress, msg.Owner.String()),
		sdk.NewAttribute(AttributeHash, msg.Hash),
	)
	ctx.EventManager().EmitEvent(event)
	return &sdk.Result{
		Log:    msg.String(),
		Events: ctx.EventManager().Events(),
	}, nil
}

func handleMsgUpdateId(ctx sdk.Context, keeper Keeper, msg MsgUpdateId) (*sdk.Result, error) {
	if !IsIdSigner(ctx, msg.GetSigners()[0], keeper) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Approver's Address is not an Enrolled Id signer")
	}

	addr := IdPrefix + msg.Owner.String()
	if !keeper.IsIdPresent(ctx, addr) {
		return nil, types.ErrIdNotExists
	}

	event := sdk.NewEvent(
		EventTypeIDUpdate,
		sdk.NewAttribute(AttributeAddress, msg.Owner.String()),
		sdk.NewAttribute(AttributeHash, msg.Hash),
	)
	ctx.EventManager().EmitEvent(event)
	keeper.SetId(ctx, addr, msg.Hash)
	return &sdk.Result{
		Log:    msg.String(),
		Events: ctx.EventManager().Events(),
	}, nil
}

func handleMsgDeleteId(ctx sdk.Context, keeper Keeper, msg MsgDeleteId) (*sdk.Result, error) {
	if !IsIdSigner(ctx, msg.GetSigners()[0], keeper) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Approver's Address is not an Enrolled Id signer")
	}

	addr := IdPrefix + msg.Owner.String()
	if !keeper.IsIdPresent(ctx, addr) {
		return nil, types.ErrIdNotExists
	}
	keeper.DeleteId(ctx, addr)
	event := sdk.NewEvent(
		EventTypeIDDelete,
		sdk.NewAttribute(AttributeAddress, msg.Owner.String()),
	)
	ctx.EventManager().EmitEvent(event)
	log := fmt.Sprintf("delete id for address %s", msg.Owner.String())
	return &sdk.Result{
		Log:    log,
		Events: ctx.EventManager().Events(),
	}, nil
}

func handleMsgEnrollIdSigners(ctx sdk.Context, keeper Keeper, msg MsgEnrollIDSigners) (*sdk.Result, error) {
	if !keeper.IsAuthority(ctx, msg.GetSigners()[0]) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Approver's Address is not authority")
	}
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	log := "signer addresses: "
	signerAllowance := sdk.NewCoins(sdk.NewCoin("shr", sdk.NewInt(int64(20))))

	for _, addr := range msg.IDSigners {
		log = log + "," + addr.String()
		signerKey := fmt.Sprintf("%s%s", IdSignerPrefix, addr.String())
		keeper.SetIdSignerStatus(ctx, signerKey, types.IdSignerActive)
		if err := keeper.LoadCoins(ctx, addr, signerAllowance); err != nil {
			return nil, err
		}
	}
	log = "Successfully enrolled" + log
	return &sdk.Result{
		Log: log,
	}, nil
}

func handleMsgRevokeIdSigners(ctx sdk.Context, keeper Keeper, msg MsgRevokeIDSigners) (*sdk.Result, error) {
	if !keeper.IsAuthority(ctx, msg.GetSigners()[0]) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Approver's Address is not authority")
	}
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	log := "signer addresses: "
	for _, addr := range msg.IDSigners {
		log = log + "," + addr.String()
		signerKey := fmt.Sprintf("%s%s", IdSignerPrefix, addr.String())
		keeper.DeleteIdSigner(ctx, signerKey)
	}
	log = "Successfully revoked" + log

	return &sdk.Result{
		Log: log,
	}, nil
}

func IsIdSigner(ctx sdk.Context, address sdk.AccAddress, k Keeper) bool {
	addr := IdSignerPrefix + address.String()
	status := k.GetIdSignerStatus(ctx, addr)
	if status == types.IdSignerActive {
		return true
	}
	return false
}
