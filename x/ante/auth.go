package ante

import (
	documenttypes "github.com/ShareRing/Shareledger/x/document/types"
	electoraltypes "github.com/ShareRing/Shareledger/x/electoral/types"
	gentleminttypes "github.com/ShareRing/Shareledger/x/gentlemint/types"
	idtypes "github.com/ShareRing/Shareledger/x/id/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type Auth struct {
	rk RoleKeeper
	ik IDKeeper
}

const (
	ErrMsgNotIdSigner        = "Transaction's Signer is not ID signer"
	ErrMsgNotSHRPLoader      = "Transaction's Signer is not SHRP loader"
	ErrMsgNotDocIssuer       = "Transaction's Signer is not document issuer"
	ErrMsgNotAuthority       = "Transaction's Signer is not authority"
	ErrMsgNotBackupAccount   = "Transaction's Signer is not the backup account"
	ErrMsgNotTreasureAccount = "Transaction's Signer is not treasure account"
	ErrMsgNotOperatorAccount = "Transaction's Signer is not operator account"
)

func NewAuthDecorator(rk RoleKeeper, ik IDKeeper) Auth {
	return Auth{
		rk: rk,
		ik: ik,
	}
}

func (a Auth) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	for _, msgI := range tx.GetMsgs() {
		signer := msgI.GetSigners()[0]
		switch msg := msgI.(type) {
		case // Authority
			*gentleminttypes.MsgLoadShr,
			*electoraltypes.MsgEnrollLoaders,
			*electoraltypes.MsgRevokeLoaders,
			*electoraltypes.MsgEnrollAccountOperator,
			*electoraltypes.MsgRevokeAccountOperator,
			*electoraltypes.MsgEnrollVoter,
			*electoraltypes.MsgRevokeVoter:
			if !a.rk.IsAuthority(ctx, signer) {
				return ctx, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, ErrMsgNotAuthority)
			}
		case // SHRP Loaders
			*gentleminttypes.MsgLoadShrp:
			if !a.rk.IsSHRPLoader(ctx, signer) {
				return ctx, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, ErrMsgNotSHRPLoader)
			}
		case // Treasure account permission
			*gentleminttypes.MsgBurnShrp,
			*gentleminttypes.MsgBurnShr,
			*gentleminttypes.MsgSetExchange:
			if !a.rk.IsTreasurer(ctx, signer) {
				return ctx, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, ErrMsgNotTreasureAccount)
			}
		case // ID Signer permission
			*idtypes.MsgCreateId,
			*idtypes.MsgCreateIdBatch,
			*idtypes.MsgUpdateId:
			if !a.rk.IsIDSigner(ctx, signer) {
				return ctx, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, ErrMsgNotIdSigner)
			}
		case // Backup account permission
			*idtypes.MsgReplaceIdOwner:
			id, _ := a.ik.GetFullIDByIDString(ctx, msg.Id)
			if id == nil || id.Data == nil || id.Data.BackupAddress != msg.BackupAddress {
				return ctx, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, ErrMsgNotBackupAccount)
			}
		case //Doc Issuer
			*documenttypes.MsgCreateDocument,
			*documenttypes.MsgCreateDocumentInBatch,
			*documenttypes.MsgUpdateDocument,
			*documenttypes.MsgRevokeDocument:
			if !a.rk.IsDocIssuer(ctx, signer) {
				return ctx, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, ErrMsgNotDocIssuer)
			}
		case // Account Operator
			*electoraltypes.MsgEnrollDocIssuer,
			*electoraltypes.MsgEnrollIdSigner,
			*electoraltypes.MsgRevokeDocIssuer,
			*electoraltypes.MsgRevokeIdSigner:
			if !a.rk.IsAccountOperator(ctx, signer) {
				return ctx, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, ErrMsgNotOperatorAccount)
			}
		}
	}

	return next(ctx, tx, simulate)
}
