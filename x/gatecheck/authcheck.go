package gatecheck

import (
	"bitbucket.org/shareringvietnam/shareledger-fix/x/electoral"
	"bitbucket.org/shareringvietnam/shareledger-fix/x/gentlemint"
	"bitbucket.org/shareringvietnam/shareledger-modules/document"
	"bitbucket.org/shareringvietnam/shareledger-modules/id"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	NotIdSignerErr      = "Transacation's Signer is not IdSigner"
	NotDocIssuerErr     = "Transacation's Signer is not document issuer"
	NotAuthorityErr     = "Transacation's Signer is not Authority"
	NotBackupAccountErr = "Transacation's Signer is not the backup account"
)

type CheckAuthDecorator struct {
	gk gentlemint.Keeper
	ik id.Keeper
}

func NewCheckAuthDecorator(gk gentlemint.Keeper, ik id.Keeper) CheckAuthDecorator {
	return CheckAuthDecorator{
		gk: gk,
		ik: ik,
	}
}

func (cfd CheckAuthDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	msgs := tx.GetMsgs()
	msg := msgs[0]

	switch msg.Type() {
	case gentlemint.TypeLoadSHRMsg, gentlemint.TypeEnrollSHRPLoaderMsg, gentlemint.TypeRevokeSHRPLoaderMsg, electoral.TypeEnrollVoter, electoral.TypeRevokeVoter:
		if msg.GetSigners()[0].String() != cfd.gk.GetAuthorityAccount(ctx) {
			return ctx, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, NotAuthorityErr)
		}
	case gentlemint.TypeLoadSHRPMsg:
		if !gentlemint.IsSHRPLoader(ctx, msg.GetSigners()[0], cfd.gk) {
			return ctx, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "Signer is not approved SHRP loader")
		}
	case gentlemint.TypeBurnSHRMsg, gentlemint.TypeBurnSHRPMsg:
		if msg.GetSigners()[0].String() != cfd.gk.GetTreasurerAccount(ctx) {
			return ctx, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Signer is not treasurer")
		}
	case id.TypeMsgCreateID, id.TypeMsgCreateIDBatch, id.TypeMsgUpdateID:
		idSigner := cfd.gk.GetIdSigner(ctx, msg.GetSigners()[0])
		if !idSigner.IsActive() {
			return ctx, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, NotIdSignerErr)
		}
	case id.TypeMsgReplaceIdOwner:
		rMsg, _ := msg.(id.MsgReplaceIdOwner)
		id := cfd.ik.GetIDByIdString(ctx, rMsg.Id)
		if !id.BackupAddr.Equals(rMsg.BackupAddr) {
			return ctx, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, NotBackupAccountErr)
		}
	case document.TypeMsgCreateDoc, document.TypeMsgCreateDocInBatch, document.TypeMsgUpdateDoc, document.TypeMsgRevokeDoc:
		issuer := cfd.gk.GetDocIssuer(ctx, msg.GetSigners()[0])
		if !issuer.IsActive() {
			return ctx, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, NotDocIssuerErr)
		}
	// case gentlemint.TypeEnrollIDSignersMsg, gentlemint.TypeRevokeIDSignersMsg:

	default:
	}
	return next(ctx, tx, simulate)
}
