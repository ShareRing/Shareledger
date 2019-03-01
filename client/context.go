package client

import (
	// "encoding/json"
	"fmt"
	"strconv"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/go-amino"
	cfg "github.com/tendermint/tendermint/config"
	cmn "github.com/tendermint/tendermint/libs/common"
	"github.com/tendermint/tendermint/privval"
	rpcclient "github.com/tendermint/tendermint/rpc/client"
	rpcTypes "github.com/tendermint/tendermint/rpc/core/types"
	tdmtypes "github.com/tendermint/tendermint/types"

	"github.com/sharering/shareledger/app"
	"github.com/sharering/shareledger/types"
	"github.com/sharering/shareledger/utils"
	"github.com/sharering/shareledger/x/auth"
	"github.com/sharering/shareledger/x/bank"
	bmsg "github.com/sharering/shareledger/x/bank/messages"
	"github.com/sharering/shareledger/x/pos"
	pmsg "github.com/sharering/shareledger/x/pos/message"
	posTypes "github.com/sharering/shareledger/x/pos/type"
)

type CoreContext struct {
	Client  rpcclient.Client
	PrivKey types.PrivKeySecp256k1
	Codec   *amino.Codec
}

type SHRAccount1 struct {
	Address sdk.AccAddress `json:"address"`
	Coins   types.Coins    `json:"coins"`
	PubKey  []byte         `json:"pub_key"`
	Nonce   int64          `json:"nonce"`
}

func NewCoreContextFromConfig(config *cfg.Config) CoreContext {
	proto, addr := getRPCAddress(config, 0)

	return CoreContext{
		Client:  rpcclient.NewHTTP(proto+"://"+addr, "/websocket"),
		PrivKey: getPrivKey(config),
		Codec:   getCodec(),
	}
}

func NewCoreContextFromConfigWithClient(config *cfg.Config, client string) CoreContext {

	return CoreContext{
		Client:  rpcclient.NewHTTP(client, "/websocket"),
		PrivKey: getPrivKey(config),
		Codec:   getCodec(),
	}
}

func (c CoreContext) ConstructTransaction(msg sdk.Msg) (auth.AuthTx, error) {
	nonce, err := c.GetNonce()
	if err != nil {
		panic(err)
	}

	authTx := auth.GetAuthTx(c.PrivKey.PubKey(), c.PrivKey, msg, nonce+1)
	return authTx, nil
}

// ConstructTendermintTransaction - encode a ShareLedger authTx in Amino and form a Tendermint tx
// before sending to ShareLedger
func (c CoreContext) ConstructTendermintTransaction(tx auth.AuthTx) (tdmtx tdmtypes.Tx, err error) {

	// amino encode
	encodedTx, err := c.Codec.MarshalBinaryLengthPrefixed(tx)
	// fmt.Printf("Tx: %x\n", encodedTx)
	if err != nil {
		return tdmtx, err
	}

	return tdmtypes.Tx(encodedTx), nil

}

func (c CoreContext) GetNonce() (int64, error) {
	nonceQuery := auth.QueryNonceParams{
		Address: c.PrivKey.PubKey().Address(),
	}

	// queryTx := types.NewQueryTx(nonceMsg)

	encodedTx, err := c.Codec.MarshalBinaryLengthPrefixed(nonceQuery)
	if err != nil {
		return -1, err
	}

	res, err := c.Client.ABCIQuery("custom/auth/nonce", encodedTx)

	if err != nil {
		return -1, err
	}

	nonce, err := strconv.ParseInt(string(res.Response.Value), 10, 64)
	if err != nil {
		return -1, err
	}
	return nonce, nil
}

func (c CoreContext) RegisterValidator(
	amount int64, // Amount of tokens to be staked
	moniker string, // name
	identity string, // optional, default to ""
	website string, // optional, default to "sharering.network"
	details string, // optional default to ""
) error {

	description := posTypes.NewDescription(moniker, identity, website, details)
	delAddr := c.PrivKey.PubKey().Address()
	pubKey := c.PrivKey.PubKey()
	delegation := types.NewPOSCoin(amount)

	msgCreateValidator := pmsg.MsgCreateValidator{
		Description:   description,
		DelegatorAddr: delAddr,
		ValidatorAddr: delAddr,
		PubKey:        pubKey,
		Delegation:    delegation,
	}

	authTx, err := c.ConstructTransaction(msgCreateValidator)
	if err != nil {
		return err
	}

	tdmTx, err := c.ConstructTendermintTransaction(authTx)
	if err != nil {
		return err
	}

	r, err := c.Client.BroadcastTxCommit(tdmTx)

	if err != nil {
		return err
	}

	err, _ = processTDMResponse(r)

	return err
}

func (c CoreContext) LoadBalance(amount int64, denom string) error {
	msgLoad := bmsg.NewMsgLoad(c.PrivKey.PubKey().Address(), types.NewCoin(denom, amount))

	authTx, err := c.ConstructTransaction(msgLoad)
	if err != nil {
		return err
	}

	tdmTx, err := c.ConstructTendermintTransaction(authTx)
	if err != nil {
		return err
	}

	r, err := c.Client.BroadcastTxCommit(tdmTx)
	if err != nil {
		return err
	}

	err, _ = processTDMResponse(r)

	return err

}

func (c CoreContext) CheckBalance() error {
	balanceQuery := bank.QueryBalanceParams{
		Address: c.PrivKey.PubKey().Address(),
	}

	encodedTx, err := c.Codec.MarshalBinaryLengthPrefixed(balanceQuery)
	if err != nil {
		return err
	}

	res, err := c.Client.ABCIQuery("custom/bank/balance", encodedTx)

	if err != nil {
		return err
	}

	fmt.Printf(string(res.Response.Value))

	if err != nil {
		return err
	}

	return nil
}

func (c CoreContext) CheckValidatorDistInfo() error {
	queryValidatorDist := pos.QueryValidatorDistParams{
		ValidatorAddr: c.PrivKey.PubKey().Address(),
	}

	req, err := c.Codec.MarshalBinaryLengthPrefixed(queryValidatorDist)
	if err != nil {
		return err
	}

	result, err := c.Client.ABCIQuery("custom/pos/validatorDistInfo", req)
	if err != nil {
		return err
	}

	var vdi posTypes.ValidatorDistInfo

	err = c.Codec.UnmarshalJSON(result.Response.Value, &vdi)

	if err != nil {
		return err
	}

	fmt.Printf("%v\n", vdi.RewardAccum)

	return nil

}

func (c CoreContext) WithdrawBlockReward() error {

	address := c.PrivKey.PubKey().Address()

	msgWithdraw := pmsg.NewMsgWithdraw(address, address)

	authTx, err := c.ConstructTransaction(msgWithdraw)
	if err != nil {
		return err
	}

	tdmTx, err := c.ConstructTendermintTransaction(authTx)
	if err != nil {
		return err
	}

	r, err := c.Client.BroadcastTxCommit(tdmTx)

	if err != nil {
		return err
	}

	err, _ = processTDMResponse(r)

	return err

}

func (c CoreContext) BeginUnbonding(amount int64) error {

	address := c.PrivKey.PubKey().Address()

	msgBeginUnbonding := pmsg.NewMsgBeginUnbonding(
		address,
		address,
		types.NewDec(amount),
	)

	authTx, err := c.ConstructTransaction(msgBeginUnbonding)
	if err != nil {
		return err
	}

	tdmTx, err := c.ConstructTendermintTransaction(authTx)
	if err != nil {
		return err
	}

	result, err := c.Client.BroadcastTxCommit(tdmTx)
	if err != nil {
		return err
	}

	err, _ = processTDMResponse(result)

	return err

}

func (c CoreContext) CompleteUnbonding() error {
	address := c.PrivKey.PubKey().Address()

	msgCompleteUnbonding := pmsg.NewMsgCompleteUnbonding(
		address,
		address,
	)

	authTx, err := c.ConstructTransaction(msgCompleteUnbonding)
	if err != nil {
		return err
	}

	tdmTx, err := c.ConstructTendermintTransaction(authTx)
	if err != nil {
		return err
	}

	result, err := c.Client.BroadcastTxCommit(tdmTx)
	if err != nil {
		return err
	}

	err, _ = processTDMResponse(result)

	return err

}

//----------------------------------------------------------
// Utilities

func getRPCAddress(config *cfg.Config, index int) (string, string) {

	rpcError := fmt.Sprintf("Incorrect RPC ListenAddress from config.toml: %s", config.RPC.ListenAddress)
	p2pError := fmt.Sprintf("Incorrect P2P PersistentPeers from config.toml: %s", config.P2P.PersistentPeers)

	parts := strings.SplitN(config.RPC.ListenAddress, "://", 2)

	if len(parts) != 2 {
		panic(rpcError)
	}

	protocol := parts[0]
	ipAddress := parts[1]

	elems := strings.Split(ipAddress, ":")

	if len(elems) != 2 {
		panic(rpcError)
	}

	port := elems[1]

	peers := strings.Split(config.P2P.PersistentPeers, ",")

	if len(peers) <= 0 {
		panic(p2pError)
	}

	parts = strings.Split(peers[index], "@")

	if len(parts) < 2 {
		panic(p2pError)
	}

	address := strings.Split(parts[1], ":")

	if len(address) < 2 {
		panic(p2pError)
	}

	ip := address[0]

	return protocol, ip + ":" + port

}

func getPrivKey(config *cfg.Config) types.PrivKeySecp256k1 {
	privValFile := config.PrivValidatorKeyFile()

	if !cmn.FileExists(privValFile) {
		panic("Private Key file does not exist")
	}

	pv := privval.LoadFilePV(privValFile, config.PrivValidatorStateFile())
	privKey := types.ConvertToPrivKey(pv.Key.PrivKey)
	return privKey

}

func getCodec() *amino.Codec {
	cdc := app.MakeCodec()
	cdc = auth.RegisterCodec(cdc)
	cdc = pos.RegisterCodec(cdc)
	cdc = bank.RegisterCodec(cdc)
	return cdc
}

func processTDMResponse(resp *rpcTypes.ResultBroadcastTxCommit) (err error, output string) {

	if resp.DeliverTx.GetCode() != 0 {
		return fmt.Errorf(utils.CleanupTDMLog(resp.DeliverTx.GetLog())), output
	}

	if resp.CheckTx.GetCode() != 0 {
		return fmt.Errorf(utils.CleanupTDMLog(resp.CheckTx.GetLog())), output
	}

	return nil, utils.CleanupTDMLog(resp.DeliverTx.GetLog())
}
