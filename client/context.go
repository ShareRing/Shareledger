package client

import (
	"fmt"
	"strconv"
	"strings"

	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
	"bitbucket.org/shareringvn/cosmos-sdk/wire"
	cfg "github.com/tendermint/tendermint/config"
	"github.com/tendermint/tendermint/privval"
	rpcclient "github.com/tendermint/tendermint/rpc/client"
	tdmtypes "github.com/tendermint/tendermint/types"
	cmn "github.com/tendermint/tmlibs/common"

	"github.com/sharering/shareledger/app"
	"github.com/sharering/shareledger/types"
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
	Codec   *wire.Codec
}

func NewCoreContextFromConfig(config *cfg.Config) CoreContext {
	proto, addr := getRPCAddress(config, 0)

	addr = "127.0.0.1:46657"
	return CoreContext{
		Client:  rpcclient.NewHTTP(proto+"://"+addr, "/websocket"),
		PrivKey: getPrivKey(config),
		Codec:   getCodec(),
	}
}

func (c CoreContext) ConstructTransaction(msg sdk.Msg) (auth.AuthTx, error) {
	nonce, err := c.GetNonce()
	if err != nil {
		panic(err)
	}

	authTx := auth.GetAuthTx(c.PrivKey.PubKey(), c.PrivKey, msg, nonce + 1)
	return authTx, nil
}

// ConstructTendermintTransaction - encode a ShareLedger authTx in Amino and form a Tendermint tx
// before sending to ShareLedger
func (c CoreContext) ConstructTendermintTransaction(tx auth.AuthTx) (tdmtx tdmtypes.Tx, err error) {

	// amino encode
	encodedTx, err := c.Codec.MarshalBinary(tx)
	if err != nil {
		return tdmtx, err
	}

	return tdmtypes.Tx(encodedTx), nil

}

func (c CoreContext) GetNonce() (int64, error) {
	nonceMsg := auth.NewMsgNonce(c.PrivKey.PubKey().Address())
	queryTx := types.NewQueryTx(nonceMsg)

	encodedTx, err := c.Codec.MarshalBinary(queryTx)
	if err != nil {
		return -1, err
	}

	res, err := c.Client.ABCIQuery("app/query", encodedTx)
	if err != nil {
		return -1, err
	}

	nonce, err := strconv.ParseInt(res.Response.Log, 10, 64)
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

	fmt.Println("Construct Transaction\n")
	authTx, err := c.ConstructTransaction(msgCreateValidator)
	if err != nil {
		return err
	}

	fmt.Println("Construct Tendermint transaction\n")

	tdmTx, err := c.ConstructTendermintTransaction(authTx)
	if err != nil {
		return err
	}

	fmt.Println("Broadcast Tx")

	result, err := c.Client.BroadcastTxSync(tdmTx)
	if err != nil {
		return err
	}

	fmt.Printf("RegisterValidator: %v\n", result)

	return nil
}

func (c CoreContext) LoadBalance(amount int64, denom string) error {
	msgLoad := bmsg.NewMsgLoad(c.PrivKey.PubKey().Address(), types.NewCoin(denom, amount))

	fmt.Println("Construct Transaction\n")
	authTx, err := c.ConstructTransaction(msgLoad)
	if err != nil {
		return err
	}

	fmt.Println("Construct Tendermint transaction\n")
	tdmTx, err := c.ConstructTendermintTransaction(authTx)
	if err != nil {
		return err
	}

	fmt.Println("Broadcast Tx")
	result, err := c.Client.BroadcastTxSync(tdmTx)
	if err != nil {
		return err
	}

	fmt.Printf("LoadBalance: %v\n", result)
	return nil

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
	privValFile := config.PrivValidatorFile()

	if !cmn.FileExists(privValFile) {
		panic("Private Key file does not exist")
	}

	pv := privval.LoadFilePV(privValFile)
	privKey := types.ConvertToPrivKey(pv.PrivKey)
	return privKey

}

func getCodec() *wire.Codec {
	cdc := app.MakeCodec()
	cdc = auth.RegisterCodec(cdc)
	cdc = pos.RegisterCodec(cdc)
	cdc = bank.RegisterCodec(cdc)
	return cdc
}
