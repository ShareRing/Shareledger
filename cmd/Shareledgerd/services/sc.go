package services

import (
	"context"
	"fmt"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	swaputil "github.com/sharering/shareledger/pkg/swap"
	"github.com/sharering/shareledger/pkg/swap/abi/sharetoken"
	"github.com/sharering/shareledger/pkg/swap/abi/swap"
	denom "github.com/sharering/shareledger/x/utils/demo"
	"math/big"
	"strings"
)

func (r *Relayer) initConn(network string) (*ethclient.Client, Network, error) {
	networkConfig, found := r.Config.Network[network]
	if !found {
		return nil, networkConfig, sdkerrors.Wrapf(sdkerrors.ErrLogic, "network, %s, is not supported", network)
	}
	conn, err := ethclient.Dial(networkConfig.Url)
	return conn, networkConfig, err
}

func (r *Relayer) isBatchDoneOnSC(network string, digest common.Hash) (done bool, err error) {
	conn, networkConfig, err := r.initConn(network)
	if err != nil {
		return false, err
	}
	defer func() {
		conn.Close()
	}()
	swapClient, err := swap.NewSwap(common.HexToAddress(networkConfig.SwapContract), conn)
	if err != nil {
		return false, sdkerrors.Wrapf(sdkerrors.ErrLogic, err.Error())
	}
	res, err := swapClient.Batch(nil, digest)
	if len(res.Signature) > 0 {
		done = true
	}
	return done, err
}

func (r *Relayer) submitBatch(ctx context.Context, network string, batchDetail swaputil.BatchDetail, tip *big.Int, gasPrice *big.Int, noSend bool) (tx *types.Transaction, signerAddr string, err error) {
	if tip != nil && gasPrice != nil {
		return nil, "", errors.New("tip and gas price should not have value at same time")
	}
	if err := batchDetail.Validate(); err != nil {
		return tx, signerAddr, err
	}
	conn, networkConfig, err := r.initConn(network)
	if err != nil {
		return tx, signerAddr, errors.Wrapf(err, "init eth connection")
	}
	defer func() {
		conn.Close()
	}()

	swapClient, err := swap.NewSwap(common.HexToAddress(networkConfig.SwapContract), conn)
	if err != nil {
		return tx, signerAddr, sdkerrors.Wrapf(sdkerrors.ErrLogic, err.Error())
	}

	info, err := r.clientTx.Keyring.Key(networkConfig.Signer)
	if err != nil {
		return tx, signerAddr, errors.Wrapf(err, "get keyring instant fail signer=%s", networkConfig.Signer)
	}
	pubKey := keyring.PubKeyETH{
		PubKey: info.GetPubKey(),
	}
	signerAddr = pubKey.Address().String()
	commonAdd := common.BytesToAddress(pubKey.Address().Bytes())

	//it should override pending nonce
	currentNonce, err := conn.NonceAt(ctx, commonAdd, nil)
	if err != nil {
		return tx, signerAddr, errors.Wrapf(err, "can't overide pending nonce for address %s", commonAdd.String())
	}
	opts, err := keyring.NewKeyedTransactorWithChainID(r.clientTx.Keyring, networkConfig.Signer, big.NewInt(networkConfig.ChainId))
	if err != nil {
		return tx, signerAddr, errors.Wrapf(err, "get eth connection options fail")
	}
	opts.GasTipCap = tip
	opts.GasPrice = gasPrice
	opts.NoSend = noSend

	opts.Nonce = big.NewInt(int64(currentNonce))
	sig, err := hexutil.Decode(batchDetail.Batch.Signature)

	if err != nil {
		return tx, signerAddr, errors.Wrapf(err, "decoding singature fail")
	}
	params, err := batchDetail.GetContractParams()
	if err != nil {
		return tx, signerAddr, err
	}
	tx, err = swapClient.Swap(opts, params.TransactionIds, params.DestAddrs, params.Amounts, sig)
	return tx, signerAddr, errors.Wrapf(err, "swapping at smart contract fail")
}

func (r *Relayer) getConfirmedTXTransfer(ctx context.Context, network string, txHash common.Hash) (toAddr common.Address, amount *sdk.Coin, err error) {
	erc20Abi, err := abi.JSON(strings.NewReader(sharetoken.SharetokenMetaData.ABI))
	if err != nil {
		return [20]byte{}, nil, errors.Wrapf(err, "unmarshal swap abi code fail")
	}
	conn, _, err := r.initConn(network)
	if err != nil {
		return [20]byte{}, nil, errors.Wrapf(err, "initConn")
	}

	tx, pending, err := conn.TransactionByHash(ctx, txHash)
	if pending {
		return [20]byte{}, nil, errors.New("the transaction is still in pending")
	}

	data, err := decodeTxParams(erc20Abi, tx.Data())
	if err != nil {
		return [20]byte{}, nil, errors.Wrapf(err, "decode tx, %s", tx.Hash())
	}
	if len(data) < 2 {
		return [20]byte{}, nil, errors.New(fmt.Sprintf("data len is not exected, %v", data))
	}
	av, ok := data["amount"].(*big.Int)
	if !ok {
		return [20]byte{}, nil, errors.New(fmt.Sprintf("%v is not in correct format of big.Int", data["amount"]))
	}
	ba := denom.ExponentToBase(sdk.NewDecFromBigInt(av), r.Config.Network[network].Exponent)
	amount = &ba
	toAddr, ok = data["to"].(common.Address)
	if !ok {
		return [20]byte{}, nil, errors.New(fmt.Sprintf("%v is not in correct format of common.Address", data["to"]))
	}
	return toAddr, amount, nil
}

func decodeTxParams(abi abi.ABI, data []byte) (map[string]interface{}, error) {
	v := make(map[string]interface{})
	m, err := abi.MethodById(data[:4])
	if err != nil {
		return map[string]interface{}{}, err
	}
	if err := m.Inputs.UnpackIntoMap(v, data[4:]); err != nil {
		return map[string]interface{}{}, err
	}
	return v, nil
}

func (r *Relayer) SuggestGasTip(ctx context.Context, network string) (*big.Int, error) {
	conn, _, err := r.initConn(network)
	if err != nil {
		return nil, errors.Wrapf(err, "init eth connection")
	}
	defer func() {
		conn.Close()
	}()
	return conn.SuggestGasTipCap(ctx)
}
func (r *Relayer) SuggestGasPrice(ctx context.Context, network string) (*big.Int, error) {
	conn, _, err := r.initConn(network)
	if err != nil {
		return nil, errors.Wrapf(err, "init eth connection")
	}
	defer func() {
		conn.Close()
	}()
	return conn.SuggestGasPrice(ctx)
}

// isLegacy check tx is support dynamic or legacy
// legacy transaction will return tip from gas price property.
// dynamic transaction, eip-1559, will return tip from tip and gas price from GasFeeCap.
// since the internal tx does not hav public field to determine types of transactions, we need to cmp to work around this.
func (r *Relayer) isLegacy(ctx context.Context, network string, batchDetail swaputil.BatchDetail) (bool, error) {
	tx, _, err := r.submitBatch(ctx, network, batchDetail, nil, nil, true)
	if err != nil {
		return false, err
	}
	return tx.GasTipCap().Cmp(tx.GasPrice()) != 0, nil
}

func (r *Relayer) getBalance(ctx context.Context, network string) (sdk.Coin, error) {
	conn, networkConfig, err := r.initConn(network)
	swapClient, err := swap.NewSwap(common.HexToAddress(networkConfig.SwapContract), conn)
	if err != nil {
		return sdk.Coin{}, errors.Wrapf(err, "fail to int Swap smartcontract client")
	}
	value, err := swapClient.TokensAvailable(&bind.CallOpts{
		Pending: false,
		Context: ctx,
	})
	return denom.ExponentToBase(sdk.NewDecFromBigInt(value), r.Config.Network[network].Exponent), err
}

func (r *Relayer) checkTxHash(ctx context.Context, network string, txHash common.Hash) (*types.Receipt, error) {
	conn, _, err := r.initConn(network)
	if err != nil {
		return nil, errors.Wrapf(err, "fail to int ETH network conection")
	}
	defer func() {
		conn.Close()
	}()
	return conn.TransactionReceipt(ctx, txHash)
}
