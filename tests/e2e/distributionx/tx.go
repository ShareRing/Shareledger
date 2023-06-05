package distributionx

import (
	"encoding/hex"
	"fmt"
	"time"

	sdkmath "cosmossdk.io/math"
	wasmcli "github.com/CosmWasm/wasmd/x/wasm/client/cli"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankcli "github.com/cosmos/cosmos-sdk/x/bank/client/cli"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/sharering/shareledger/tests"
	"github.com/sharering/shareledger/testutil/network"
	"github.com/sharering/shareledger/x/distributionx/client/cli"
	"github.com/sharering/shareledger/x/distributionx/types"
	"github.com/sharering/shareledger/x/utils/denom"
)

func (s *E2ETestSuite) TestWithdrawReward() {
	val := s.network.Validators[0]
	acc1Address := network.MustAddressFormKeyring(val.ClientCtx.Keyring, network.KeyAccount1).String()
	reward := s.getReward(acc1Address)
	before := s.getBalance(acc1Address)
	_, err := tests.RunCmdBlock(
		&s.Suite,
		cli.CmdWithdrawReward(),
		val,
		[]string{
			network.MakeByAccount(network.KeyAccount1),
		},
	)
	after := s.getBalance(acc1Address)
	s.NoError(err)
	// check balances increase (and substract tx fee = 10SHR)
	s.Equal(reward.Sub(sdkmath.NewInt(10_000_000_000)), after.Sub(before))
}

func (s *E2ETestSuite) TestRewardNormalTx() {
	val := s.network.Validators[0]
	acc1 := network.MustAddressFormKeyring(val.ClientCtx.Keyring, network.KeyAccount1)

	before := s.getReward(devPoolAccount)

	_, err := tests.RunCmdBlock(
		&s.Suite,
		bankcli.NewSendTxCmd(),
		val,
		[]string{
			acc1.String(),
			"shareledger18pf3zdwqjntd9wkvfcjvmdc7hua6c0q2eck5h5",
			"1nshr",
			network.MakeByAccount(network.KeyAccount1),
		})
	s.NoError(err)
	s.NoError(s.network.WaitForNextBlock())

	after := s.getReward(devPoolAccount)
	// check dev pool account have reward of 5SHR
	s.Equal(sdkmath.NewInt(5_000_000_000), after.Sub(before))
}

func (s *E2ETestSuite) TestRewardWasmTx() {
	val := s.network.Validators[0]
	contractAddress := s.setupContract()
	acc1Address := network.MustAddressFormKeyring(val.ClientCtx.Keyring, network.KeyAccount1).String()
	s.waitNextWindow()
	for i := 0; i < 3; i++ {
		s.mintToken(contractAddress, acc1Address)
	}
	s.waitNextWindow()
	// make sure contractAddress in builder list
	bl := s.getBuiderList()
	s.Len(bl, 1)
	s.Equal(bl[0].ContractAddress, contractAddress)

	beforeAcc1 := s.getReward(acc1Address)
	before := s.getReward(devPoolAccount)
	s.mintToken(contractAddress, acc1Address)
	s.NoError(s.network.WaitForNextBlock())

	afterAcc1 := s.getReward(acc1Address)
	after := s.getReward(devPoolAccount)
	// check acc1 account have reward of 2.5SHR
	s.Equal(sdkmath.NewInt(2_500_000_000), afterAcc1.Sub(beforeAcc1))
	// check dev pool account have reward of 2.5SHR
	s.Equal(sdkmath.NewInt(2_500_000_000), after.Sub(before))
}

// setupContract return address of deployed contract
func (s *E2ETestSuite) setupContract() string {
	val := s.network.Validators[0]
	// store the code
	resp, err := tests.RunCmdBlock(
		&s.Suite,
		wasmcli.StoreCodeCmd(),
		val,
		[]string{
			"testdata/cw721_base.wasm",
			network.MakeByAccount(network.KeyAccount1),
		},
	)
	s.NoError(err)
	msgData, err := decodeMsg(s.cfg.Codec, resp.Data)
	s.NoError(err)
	storeCodeResp := &wasmtypes.MsgStoreCodeResponse{}
	s.cfg.Codec.MustUnmarshal(msgData.MsgResponses[0].Value, storeCodeResp)

	// instantiate contract
	minter := network.MustAddressFormKeyring(val.ClientCtx.Keyring, network.KeyAccount1)
	initMsg := InstantiateMsg{
		Name:   "NFT Contract",
		Symbol: "NFT",
		Minter: minter.String(),
	}
	resp, err = tests.RunCmdBlock(
		&s.Suite,
		wasmcli.InstantiateContractCmd(),
		val, []string{
			fmt.Sprint(storeCodeResp.CodeID),
			initMsg.MustToJSON(),
			"--label", "NFT contract", "--no-admin",
			network.MakeByAccount(network.KeyAccount1),
		},
	)
	s.NoError(err)
	msgData, err = decodeMsg(s.cfg.Codec, resp.Data)
	s.NoError(err)
	initContractResp := &wasmtypes.MsgInstantiateContractResponse{}
	s.cfg.Codec.MustUnmarshal(msgData.MsgResponses[0].Value, initContractResp)
	return initContractResp.Address
}

func (s *E2ETestSuite) waitNextWindow() {
	builderWindows := params.BuilderWindows
	for {
		height, err := s.network.LatestHeight()
		if err == nil && height%int64(builderWindows) == 0 {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func (s *E2ETestSuite) getReward(address string) sdkmath.Int {
	resp := &types.QueryGetRewardResponse{}
	err := tests.RunQueryCmd(s.network.Validators[0], cli.CmdShowReward(), []string{address}, resp)
	if err != nil {
		return sdkmath.NewInt(0)
	}
	return resp.Reward.Amount.AmountOf(denom.Base)
}

func (s *E2ETestSuite) getBalance(address string) sdkmath.Int {
	resp := &banktypes.QueryAllBalancesResponse{}
	err := tests.RunQueryCmd(s.network.Validators[0], bankcli.GetBalancesCmd(), []string{address}, resp)
	if err != nil {
		return sdkmath.NewInt(0)
	}
	return resp.Balances.AmountOf(denom.Base)
}

func (s *E2ETestSuite) getBuiderList() []types.BuilderList {
	resp := &types.QueryAllBuilderListResponse{}
	err := tests.RunQueryCmd(s.network.Validators[0], cli.CmdListBuilderList(), nil, resp)
	if err != nil {
		panic(err)
	}
	return resp.BuilderList
}

func (s *E2ETestSuite) mintToken(contractAddress, receiver string) {
	val := s.network.Validators[0]
	msg := MintMsg{
		Owner: receiver,
	}
	_, err := tests.RunCmdBlock(
		&s.Suite,
		wasmcli.ExecuteContractCmd(),
		val,
		[]string{
			contractAddress,
			msg.MustToJSON(),
			network.MakeByAccount(network.KeyAccount1),
		},
	)
	s.NoError(err)
}

func decodeMsg(cdc codec.Codec, data string) (*sdk.TxMsgData, error) {
	b, err := hex.DecodeString(data)
	if err != nil {
		return nil, err
	}
	txMsg := &sdk.TxMsgData{}
	err = cdc.Unmarshal(b, txMsg)
	return txMsg, err
}
