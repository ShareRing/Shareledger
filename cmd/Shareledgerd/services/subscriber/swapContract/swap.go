// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package swapcontract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// SwapcontractMetaData contains all meta data concerning the Swapcontract contract.
var SwapcontractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_relayer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_approver\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"SwapCompleted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_approvers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"addApprover\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"addApprovers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"approver\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_digest\",\"type\":\"bytes32\"}],\"name\":\"batch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"tos\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_ids\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"_tos\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"digest\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_digest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"relayer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"removeApprover\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"removeApprovers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"request\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setRelayer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_ids\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"_tos\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"swap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokensAvailable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unsetRelayer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_ids\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"_tos\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SwapcontractABI is the input ABI used to generate the binding from.
// Deprecated: Use SwapcontractMetaData.ABI instead.
var SwapcontractABI = SwapcontractMetaData.ABI

// Swapcontract is an auto generated Go binding around an Ethereum contract.
type Swapcontract struct {
	SwapcontractCaller     // Read-only binding to the contract
	SwapcontractTransactor // Write-only binding to the contract
	SwapcontractFilterer   // Log filterer for contract events
}

// SwapcontractCaller is an auto generated read-only Go binding around an Ethereum contract.
type SwapcontractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapcontractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SwapcontractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapcontractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SwapcontractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapcontractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SwapcontractSession struct {
	Contract     *Swapcontract     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwapcontractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SwapcontractCallerSession struct {
	Contract *SwapcontractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// SwapcontractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SwapcontractTransactorSession struct {
	Contract     *SwapcontractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// SwapcontractRaw is an auto generated low-level Go binding around an Ethereum contract.
type SwapcontractRaw struct {
	Contract *Swapcontract // Generic contract binding to access the raw methods on
}

// SwapcontractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SwapcontractCallerRaw struct {
	Contract *SwapcontractCaller // Generic read-only contract binding to access the raw methods on
}

// SwapcontractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SwapcontractTransactorRaw struct {
	Contract *SwapcontractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSwapcontract creates a new instance of Swapcontract, bound to a specific deployed contract.
func NewSwapcontract(address common.Address, backend bind.ContractBackend) (*Swapcontract, error) {
	contract, err := bindSwapcontract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Swapcontract{SwapcontractCaller: SwapcontractCaller{contract: contract}, SwapcontractTransactor: SwapcontractTransactor{contract: contract}, SwapcontractFilterer: SwapcontractFilterer{contract: contract}}, nil
}

// NewSwapcontractCaller creates a new read-only instance of Swapcontract, bound to a specific deployed contract.
func NewSwapcontractCaller(address common.Address, caller bind.ContractCaller) (*SwapcontractCaller, error) {
	contract, err := bindSwapcontract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SwapcontractCaller{contract: contract}, nil
}

// NewSwapcontractTransactor creates a new write-only instance of Swapcontract, bound to a specific deployed contract.
func NewSwapcontractTransactor(address common.Address, transactor bind.ContractTransactor) (*SwapcontractTransactor, error) {
	contract, err := bindSwapcontract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SwapcontractTransactor{contract: contract}, nil
}

// NewSwapcontractFilterer creates a new log filterer instance of Swapcontract, bound to a specific deployed contract.
func NewSwapcontractFilterer(address common.Address, filterer bind.ContractFilterer) (*SwapcontractFilterer, error) {
	contract, err := bindSwapcontract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SwapcontractFilterer{contract: contract}, nil
}

// bindSwapcontract binds a generic wrapper to an already deployed contract.
func bindSwapcontract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SwapcontractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Swapcontract *SwapcontractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Swapcontract.Contract.SwapcontractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Swapcontract *SwapcontractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swapcontract.Contract.SwapcontractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Swapcontract *SwapcontractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Swapcontract.Contract.SwapcontractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Swapcontract *SwapcontractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Swapcontract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Swapcontract *SwapcontractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swapcontract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Swapcontract *SwapcontractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Swapcontract.Contract.contract.Transact(opts, method, params...)
}

// Approvers is a free data retrieval call binding the contract method 0x3ab970e0.
//
// Solidity: function _approvers(address ) view returns(bool)
func (_Swapcontract *SwapcontractCaller) Approvers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Swapcontract.contract.Call(opts, &out, "_approvers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Approvers is a free data retrieval call binding the contract method 0x3ab970e0.
//
// Solidity: function _approvers(address ) view returns(bool)
func (_Swapcontract *SwapcontractSession) Approvers(arg0 common.Address) (bool, error) {
	return _Swapcontract.Contract.Approvers(&_Swapcontract.CallOpts, arg0)
}

// Approvers is a free data retrieval call binding the contract method 0x3ab970e0.
//
// Solidity: function _approvers(address ) view returns(bool)
func (_Swapcontract *SwapcontractCallerSession) Approvers(arg0 common.Address) (bool, error) {
	return _Swapcontract.Contract.Approvers(&_Swapcontract.CallOpts, arg0)
}

// Approver is a free data retrieval call binding the contract method 0x3f131d9f.
//
// Solidity: function approver(address _address) view returns(bool)
func (_Swapcontract *SwapcontractCaller) Approver(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var out []interface{}
	err := _Swapcontract.contract.Call(opts, &out, "approver", _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Approver is a free data retrieval call binding the contract method 0x3f131d9f.
//
// Solidity: function approver(address _address) view returns(bool)
func (_Swapcontract *SwapcontractSession) Approver(_address common.Address) (bool, error) {
	return _Swapcontract.Contract.Approver(&_Swapcontract.CallOpts, _address)
}

// Approver is a free data retrieval call binding the contract method 0x3f131d9f.
//
// Solidity: function approver(address _address) view returns(bool)
func (_Swapcontract *SwapcontractCallerSession) Approver(_address common.Address) (bool, error) {
	return _Swapcontract.Contract.Approver(&_Swapcontract.CallOpts, _address)
}

// Batch is a free data retrieval call binding the contract method 0xfddaa065.
//
// Solidity: function batch(bytes32 _digest) view returns(uint256[] ids, address[] tos, uint256[] amounts, bytes signature)
func (_Swapcontract *SwapcontractCaller) Batch(opts *bind.CallOpts, _digest [32]byte) (struct {
	Ids       []*big.Int
	Tos       []common.Address
	Amounts   []*big.Int
	Signature []byte
}, error) {
	var out []interface{}
	err := _Swapcontract.contract.Call(opts, &out, "batch", _digest)

	outstruct := new(struct {
		Ids       []*big.Int
		Tos       []common.Address
		Amounts   []*big.Int
		Signature []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Ids = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.Tos = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	outstruct.Amounts = *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)
	outstruct.Signature = *abi.ConvertType(out[3], new([]byte)).(*[]byte)

	return *outstruct, err

}

// Batch is a free data retrieval call binding the contract method 0xfddaa065.
//
// Solidity: function batch(bytes32 _digest) view returns(uint256[] ids, address[] tos, uint256[] amounts, bytes signature)
func (_Swapcontract *SwapcontractSession) Batch(_digest [32]byte) (struct {
	Ids       []*big.Int
	Tos       []common.Address
	Amounts   []*big.Int
	Signature []byte
}, error) {
	return _Swapcontract.Contract.Batch(&_Swapcontract.CallOpts, _digest)
}

// Batch is a free data retrieval call binding the contract method 0xfddaa065.
//
// Solidity: function batch(bytes32 _digest) view returns(uint256[] ids, address[] tos, uint256[] amounts, bytes signature)
func (_Swapcontract *SwapcontractCallerSession) Batch(_digest [32]byte) (struct {
	Ids       []*big.Int
	Tos       []common.Address
	Amounts   []*big.Int
	Signature []byte
}, error) {
	return _Swapcontract.Contract.Batch(&_Swapcontract.CallOpts, _digest)
}

// Digest is a free data retrieval call binding the contract method 0xf6708f96.
//
// Solidity: function digest(uint256[] _ids, address[] _tos, uint256[] _amounts) view returns(bytes32 _digest)
func (_Swapcontract *SwapcontractCaller) Digest(opts *bind.CallOpts, _ids []*big.Int, _tos []common.Address, _amounts []*big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Swapcontract.contract.Call(opts, &out, "digest", _ids, _tos, _amounts)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Digest is a free data retrieval call binding the contract method 0xf6708f96.
//
// Solidity: function digest(uint256[] _ids, address[] _tos, uint256[] _amounts) view returns(bytes32 _digest)
func (_Swapcontract *SwapcontractSession) Digest(_ids []*big.Int, _tos []common.Address, _amounts []*big.Int) ([32]byte, error) {
	return _Swapcontract.Contract.Digest(&_Swapcontract.CallOpts, _ids, _tos, _amounts)
}

// Digest is a free data retrieval call binding the contract method 0xf6708f96.
//
// Solidity: function digest(uint256[] _ids, address[] _tos, uint256[] _amounts) view returns(bytes32 _digest)
func (_Swapcontract *SwapcontractCallerSession) Digest(_ids []*big.Int, _tos []common.Address, _amounts []*big.Int) ([32]byte, error) {
	return _Swapcontract.Contract.Digest(&_Swapcontract.CallOpts, _ids, _tos, _amounts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Swapcontract *SwapcontractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Swapcontract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Swapcontract *SwapcontractSession) Owner() (common.Address, error) {
	return _Swapcontract.Contract.Owner(&_Swapcontract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Swapcontract *SwapcontractCallerSession) Owner() (common.Address, error) {
	return _Swapcontract.Contract.Owner(&_Swapcontract.CallOpts)
}

// Relayer is a free data retrieval call binding the contract method 0x8406c079.
//
// Solidity: function relayer() view returns(address)
func (_Swapcontract *SwapcontractCaller) Relayer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Swapcontract.contract.Call(opts, &out, "relayer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Relayer is a free data retrieval call binding the contract method 0x8406c079.
//
// Solidity: function relayer() view returns(address)
func (_Swapcontract *SwapcontractSession) Relayer() (common.Address, error) {
	return _Swapcontract.Contract.Relayer(&_Swapcontract.CallOpts)
}

// Relayer is a free data retrieval call binding the contract method 0x8406c079.
//
// Solidity: function relayer() view returns(address)
func (_Swapcontract *SwapcontractCallerSession) Relayer() (common.Address, error) {
	return _Swapcontract.Contract.Relayer(&_Swapcontract.CallOpts)
}

// Request is a free data retrieval call binding the contract method 0xd845a4b3.
//
// Solidity: function request(uint256 _id) view returns(bytes32)
func (_Swapcontract *SwapcontractCaller) Request(opts *bind.CallOpts, _id *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Swapcontract.contract.Call(opts, &out, "request", _id)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Request is a free data retrieval call binding the contract method 0xd845a4b3.
//
// Solidity: function request(uint256 _id) view returns(bytes32)
func (_Swapcontract *SwapcontractSession) Request(_id *big.Int) ([32]byte, error) {
	return _Swapcontract.Contract.Request(&_Swapcontract.CallOpts, _id)
}

// Request is a free data retrieval call binding the contract method 0xd845a4b3.
//
// Solidity: function request(uint256 _id) view returns(bytes32)
func (_Swapcontract *SwapcontractCallerSession) Request(_id *big.Int) ([32]byte, error) {
	return _Swapcontract.Contract.Request(&_Swapcontract.CallOpts, _id)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Swapcontract *SwapcontractCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Swapcontract.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Swapcontract *SwapcontractSession) Token() (common.Address, error) {
	return _Swapcontract.Contract.Token(&_Swapcontract.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Swapcontract *SwapcontractCallerSession) Token() (common.Address, error) {
	return _Swapcontract.Contract.Token(&_Swapcontract.CallOpts)
}

// TokensAvailable is a free data retrieval call binding the contract method 0x60659a92.
//
// Solidity: function tokensAvailable() view returns(uint256)
func (_Swapcontract *SwapcontractCaller) TokensAvailable(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Swapcontract.contract.Call(opts, &out, "tokensAvailable")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokensAvailable is a free data retrieval call binding the contract method 0x60659a92.
//
// Solidity: function tokensAvailable() view returns(uint256)
func (_Swapcontract *SwapcontractSession) TokensAvailable() (*big.Int, error) {
	return _Swapcontract.Contract.TokensAvailable(&_Swapcontract.CallOpts)
}

// TokensAvailable is a free data retrieval call binding the contract method 0x60659a92.
//
// Solidity: function tokensAvailable() view returns(uint256)
func (_Swapcontract *SwapcontractCallerSession) TokensAvailable() (*big.Int, error) {
	return _Swapcontract.Contract.TokensAvailable(&_Swapcontract.CallOpts)
}

// Verify is a free data retrieval call binding the contract method 0x1d3d89b6.
//
// Solidity: function verify(uint256[] _ids, address[] _tos, uint256[] _amounts, bytes signature) view returns(address _address)
func (_Swapcontract *SwapcontractCaller) Verify(opts *bind.CallOpts, _ids []*big.Int, _tos []common.Address, _amounts []*big.Int, signature []byte) (common.Address, error) {
	var out []interface{}
	err := _Swapcontract.contract.Call(opts, &out, "verify", _ids, _tos, _amounts, signature)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Verify is a free data retrieval call binding the contract method 0x1d3d89b6.
//
// Solidity: function verify(uint256[] _ids, address[] _tos, uint256[] _amounts, bytes signature) view returns(address _address)
func (_Swapcontract *SwapcontractSession) Verify(_ids []*big.Int, _tos []common.Address, _amounts []*big.Int, signature []byte) (common.Address, error) {
	return _Swapcontract.Contract.Verify(&_Swapcontract.CallOpts, _ids, _tos, _amounts, signature)
}

// Verify is a free data retrieval call binding the contract method 0x1d3d89b6.
//
// Solidity: function verify(uint256[] _ids, address[] _tos, uint256[] _amounts, bytes signature) view returns(address _address)
func (_Swapcontract *SwapcontractCallerSession) Verify(_ids []*big.Int, _tos []common.Address, _amounts []*big.Int, signature []byte) (common.Address, error) {
	return _Swapcontract.Contract.Verify(&_Swapcontract.CallOpts, _ids, _tos, _amounts, signature)
}

// AddApprover is a paid mutator transaction binding the contract method 0xb646c194.
//
// Solidity: function addApprover(address _address) returns()
func (_Swapcontract *SwapcontractTransactor) AddApprover(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Swapcontract.contract.Transact(opts, "addApprover", _address)
}

// AddApprover is a paid mutator transaction binding the contract method 0xb646c194.
//
// Solidity: function addApprover(address _address) returns()
func (_Swapcontract *SwapcontractSession) AddApprover(_address common.Address) (*types.Transaction, error) {
	return _Swapcontract.Contract.AddApprover(&_Swapcontract.TransactOpts, _address)
}

// AddApprover is a paid mutator transaction binding the contract method 0xb646c194.
//
// Solidity: function addApprover(address _address) returns()
func (_Swapcontract *SwapcontractTransactorSession) AddApprover(_address common.Address) (*types.Transaction, error) {
	return _Swapcontract.Contract.AddApprover(&_Swapcontract.TransactOpts, _address)
}

// AddApprovers is a paid mutator transaction binding the contract method 0x6a882d51.
//
// Solidity: function addApprovers(address[] _addresses) returns()
func (_Swapcontract *SwapcontractTransactor) AddApprovers(opts *bind.TransactOpts, _addresses []common.Address) (*types.Transaction, error) {
	return _Swapcontract.contract.Transact(opts, "addApprovers", _addresses)
}

// AddApprovers is a paid mutator transaction binding the contract method 0x6a882d51.
//
// Solidity: function addApprovers(address[] _addresses) returns()
func (_Swapcontract *SwapcontractSession) AddApprovers(_addresses []common.Address) (*types.Transaction, error) {
	return _Swapcontract.Contract.AddApprovers(&_Swapcontract.TransactOpts, _addresses)
}

// AddApprovers is a paid mutator transaction binding the contract method 0x6a882d51.
//
// Solidity: function addApprovers(address[] _addresses) returns()
func (_Swapcontract *SwapcontractTransactorSession) AddApprovers(_addresses []common.Address) (*types.Transaction, error) {
	return _Swapcontract.Contract.AddApprovers(&_Swapcontract.TransactOpts, _addresses)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_Swapcontract *SwapcontractTransactor) Deposit(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Swapcontract.contract.Transact(opts, "deposit", _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_Swapcontract *SwapcontractSession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _Swapcontract.Contract.Deposit(&_Swapcontract.TransactOpts, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_Swapcontract *SwapcontractTransactorSession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _Swapcontract.Contract.Deposit(&_Swapcontract.TransactOpts, _amount)
}

// RemoveApprover is a paid mutator transaction binding the contract method 0x6cf4c88f.
//
// Solidity: function removeApprover(address _address) returns()
func (_Swapcontract *SwapcontractTransactor) RemoveApprover(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Swapcontract.contract.Transact(opts, "removeApprover", _address)
}

// RemoveApprover is a paid mutator transaction binding the contract method 0x6cf4c88f.
//
// Solidity: function removeApprover(address _address) returns()
func (_Swapcontract *SwapcontractSession) RemoveApprover(_address common.Address) (*types.Transaction, error) {
	return _Swapcontract.Contract.RemoveApprover(&_Swapcontract.TransactOpts, _address)
}

// RemoveApprover is a paid mutator transaction binding the contract method 0x6cf4c88f.
//
// Solidity: function removeApprover(address _address) returns()
func (_Swapcontract *SwapcontractTransactorSession) RemoveApprover(_address common.Address) (*types.Transaction, error) {
	return _Swapcontract.Contract.RemoveApprover(&_Swapcontract.TransactOpts, _address)
}

// RemoveApprovers is a paid mutator transaction binding the contract method 0x7569d66f.
//
// Solidity: function removeApprovers(address[] _addresses) returns()
func (_Swapcontract *SwapcontractTransactor) RemoveApprovers(opts *bind.TransactOpts, _addresses []common.Address) (*types.Transaction, error) {
	return _Swapcontract.contract.Transact(opts, "removeApprovers", _addresses)
}

// RemoveApprovers is a paid mutator transaction binding the contract method 0x7569d66f.
//
// Solidity: function removeApprovers(address[] _addresses) returns()
func (_Swapcontract *SwapcontractSession) RemoveApprovers(_addresses []common.Address) (*types.Transaction, error) {
	return _Swapcontract.Contract.RemoveApprovers(&_Swapcontract.TransactOpts, _addresses)
}

// RemoveApprovers is a paid mutator transaction binding the contract method 0x7569d66f.
//
// Solidity: function removeApprovers(address[] _addresses) returns()
func (_Swapcontract *SwapcontractTransactorSession) RemoveApprovers(_addresses []common.Address) (*types.Transaction, error) {
	return _Swapcontract.Contract.RemoveApprovers(&_Swapcontract.TransactOpts, _addresses)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Swapcontract *SwapcontractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swapcontract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Swapcontract *SwapcontractSession) RenounceOwnership() (*types.Transaction, error) {
	return _Swapcontract.Contract.RenounceOwnership(&_Swapcontract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Swapcontract *SwapcontractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Swapcontract.Contract.RenounceOwnership(&_Swapcontract.TransactOpts)
}

// SetRelayer is a paid mutator transaction binding the contract method 0x6548e9bc.
//
// Solidity: function setRelayer(address _address) returns()
func (_Swapcontract *SwapcontractTransactor) SetRelayer(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Swapcontract.contract.Transact(opts, "setRelayer", _address)
}

// SetRelayer is a paid mutator transaction binding the contract method 0x6548e9bc.
//
// Solidity: function setRelayer(address _address) returns()
func (_Swapcontract *SwapcontractSession) SetRelayer(_address common.Address) (*types.Transaction, error) {
	return _Swapcontract.Contract.SetRelayer(&_Swapcontract.TransactOpts, _address)
}

// SetRelayer is a paid mutator transaction binding the contract method 0x6548e9bc.
//
// Solidity: function setRelayer(address _address) returns()
func (_Swapcontract *SwapcontractTransactorSession) SetRelayer(_address common.Address) (*types.Transaction, error) {
	return _Swapcontract.Contract.SetRelayer(&_Swapcontract.TransactOpts, _address)
}

// Swap is a paid mutator transaction binding the contract method 0x58d937ce.
//
// Solidity: function swap(uint256[] _ids, address[] _tos, uint256[] _amounts, bytes _signature) returns()
func (_Swapcontract *SwapcontractTransactor) Swap(opts *bind.TransactOpts, _ids []*big.Int, _tos []common.Address, _amounts []*big.Int, _signature []byte) (*types.Transaction, error) {
	return _Swapcontract.contract.Transact(opts, "swap", _ids, _tos, _amounts, _signature)
}

// Swap is a paid mutator transaction binding the contract method 0x58d937ce.
//
// Solidity: function swap(uint256[] _ids, address[] _tos, uint256[] _amounts, bytes _signature) returns()
func (_Swapcontract *SwapcontractSession) Swap(_ids []*big.Int, _tos []common.Address, _amounts []*big.Int, _signature []byte) (*types.Transaction, error) {
	return _Swapcontract.Contract.Swap(&_Swapcontract.TransactOpts, _ids, _tos, _amounts, _signature)
}

// Swap is a paid mutator transaction binding the contract method 0x58d937ce.
//
// Solidity: function swap(uint256[] _ids, address[] _tos, uint256[] _amounts, bytes _signature) returns()
func (_Swapcontract *SwapcontractTransactorSession) Swap(_ids []*big.Int, _tos []common.Address, _amounts []*big.Int, _signature []byte) (*types.Transaction, error) {
	return _Swapcontract.Contract.Swap(&_Swapcontract.TransactOpts, _ids, _tos, _amounts, _signature)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Swapcontract *SwapcontractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Swapcontract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Swapcontract *SwapcontractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Swapcontract.Contract.TransferOwnership(&_Swapcontract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Swapcontract *SwapcontractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Swapcontract.Contract.TransferOwnership(&_Swapcontract.TransactOpts, newOwner)
}

// UnsetRelayer is a paid mutator transaction binding the contract method 0xa521d4df.
//
// Solidity: function unsetRelayer() returns()
func (_Swapcontract *SwapcontractTransactor) UnsetRelayer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swapcontract.contract.Transact(opts, "unsetRelayer")
}

// UnsetRelayer is a paid mutator transaction binding the contract method 0xa521d4df.
//
// Solidity: function unsetRelayer() returns()
func (_Swapcontract *SwapcontractSession) UnsetRelayer() (*types.Transaction, error) {
	return _Swapcontract.Contract.UnsetRelayer(&_Swapcontract.TransactOpts)
}

// UnsetRelayer is a paid mutator transaction binding the contract method 0xa521d4df.
//
// Solidity: function unsetRelayer() returns()
func (_Swapcontract *SwapcontractTransactorSession) UnsetRelayer() (*types.Transaction, error) {
	return _Swapcontract.Contract.UnsetRelayer(&_Swapcontract.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Swapcontract *SwapcontractTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swapcontract.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Swapcontract *SwapcontractSession) Withdraw() (*types.Transaction, error) {
	return _Swapcontract.Contract.Withdraw(&_Swapcontract.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Swapcontract *SwapcontractTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Swapcontract.Contract.Withdraw(&_Swapcontract.TransactOpts)
}

// SwapcontractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Swapcontract contract.
type SwapcontractOwnershipTransferredIterator struct {
	Event *SwapcontractOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwapcontractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapcontractOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwapcontractOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwapcontractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapcontractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapcontractOwnershipTransferred represents a OwnershipTransferred event raised by the Swapcontract contract.
type SwapcontractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Swapcontract *SwapcontractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SwapcontractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Swapcontract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SwapcontractOwnershipTransferredIterator{contract: _Swapcontract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Swapcontract *SwapcontractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SwapcontractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Swapcontract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapcontractOwnershipTransferred)
				if err := _Swapcontract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Swapcontract *SwapcontractFilterer) ParseOwnershipTransferred(log types.Log) (*SwapcontractOwnershipTransferred, error) {
	event := new(SwapcontractOwnershipTransferred)
	if err := _Swapcontract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapcontractSwapCompletedIterator is returned from FilterSwapCompleted and is used to iterate over the raw logs and unpacked data for SwapCompleted events raised by the Swapcontract contract.
type SwapcontractSwapCompletedIterator struct {
	Event *SwapcontractSwapCompleted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwapcontractSwapCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapcontractSwapCompleted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwapcontractSwapCompleted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwapcontractSwapCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapcontractSwapCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapcontractSwapCompleted represents a SwapCompleted event raised by the Swapcontract contract.
type SwapcontractSwapCompleted struct {
	Ids []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSwapCompleted is a free log retrieval operation binding the contract event 0x796a6fb73c9c09afe863a5d1bc7040da846e5aeb2ad3cb42ee36e08a0c0a3e71.
//
// Solidity: event SwapCompleted(uint256[] indexed ids)
func (_Swapcontract *SwapcontractFilterer) FilterSwapCompleted(opts *bind.FilterOpts, ids [][]*big.Int) (*SwapcontractSwapCompletedIterator, error) {

	var idsRule []interface{}
	for _, idsItem := range ids {
		idsRule = append(idsRule, idsItem)
	}

	logs, sub, err := _Swapcontract.contract.FilterLogs(opts, "SwapCompleted", idsRule)
	if err != nil {
		return nil, err
	}
	return &SwapcontractSwapCompletedIterator{contract: _Swapcontract.contract, event: "SwapCompleted", logs: logs, sub: sub}, nil
}

// WatchSwapCompleted is a free log subscription operation binding the contract event 0x796a6fb73c9c09afe863a5d1bc7040da846e5aeb2ad3cb42ee36e08a0c0a3e71.
//
// Solidity: event SwapCompleted(uint256[] indexed ids)
func (_Swapcontract *SwapcontractFilterer) WatchSwapCompleted(opts *bind.WatchOpts, sink chan<- *SwapcontractSwapCompleted, ids [][]*big.Int) (event.Subscription, error) {

	var idsRule []interface{}
	for _, idsItem := range ids {
		idsRule = append(idsRule, idsItem)
	}

	logs, sub, err := _Swapcontract.contract.WatchLogs(opts, "SwapCompleted", idsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapcontractSwapCompleted)
				if err := _Swapcontract.contract.UnpackLog(event, "SwapCompleted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSwapCompleted is a log parse operation binding the contract event 0x796a6fb73c9c09afe863a5d1bc7040da846e5aeb2ad3cb42ee36e08a0c0a3e71.
//
// Solidity: event SwapCompleted(uint256[] indexed ids)
func (_Swapcontract *SwapcontractFilterer) ParseSwapCompleted(log types.Log) (*SwapcontractSwapCompleted, error) {
	event := new(SwapcontractSwapCompleted)
	if err := _Swapcontract.contract.UnpackLog(event, "SwapCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
