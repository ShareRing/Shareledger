package services

//
//import (
//	"github.com/cosmos/cosmos-sdk/client"
//	"github.com/cosmos/cosmos-sdk/client/tx"
//	sdk "github.com/cosmos/cosmos-sdk/types"
//	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
//	"github.com/spf13/pflag"
//)
//
//func (r *Relayer) AddPreRunBroadcastTx(tx tx.PreRunBroadcastTx) {
//	r.preRunBroadcastTxs = append(r.preRunBroadcastTxs, tx)
//}
//
//func (r *Relayer) BroadcastTx(clientCtx client.Context, flagSet *pflag.FlagSet, msgs ...sdk.Msg) error {
//	txf := tx.NewFactoryCLI(clientCtx, flagSet)
//	// GenerateOrBroadcastTxWithFactory(clientCtx client.Context, txf Factory, msgs ...sdk.Msg) error
//	for _, msg := range msgs {
//		if err := msg.ValidateBasic(); err != nil {
//			return err
//		}
//	}
//	// Check if there is any pre run funcs that need to be executed before broadcast
//	var err error
//	for i, f := range r.preRunBroadcastTxs {
//		clientCtx, txf, msgs, err = f(clientCtx, txf, msgs...)
//		if err != nil {
//			return sdkerrors.Wrapf(err, "pre run function at %v", i)
//		}
//	}
//	//BroadcastTx(clientCtx client.Context, txf Factory, msgs ...sdk.Msg) error
//	//txf, err := prepareFactory(clientCtx, txf)
//
//
//	if err != nil {
//		return err
//	}
//}
