// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ierc20metadata

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

// Ierc20metadataMetaData contains all meta data concerning the Ierc20metadata contract.
var Ierc20metadataMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// Ierc20metadataABI is the input ABI used to generate the binding from.
// Deprecated: Use Ierc20metadataMetaData.ABI instead.
var Ierc20metadataABI = Ierc20metadataMetaData.ABI

// Ierc20metadata is an auto generated Go binding around an Ethereum contract.
type Ierc20metadata struct {
	Ierc20metadataCaller     // Read-only binding to the contract
	Ierc20metadataTransactor // Write-only binding to the contract
	Ierc20metadataFilterer   // Log filterer for contract events
}

// Ierc20metadataCaller is an auto generated read-only Go binding around an Ethereum contract.
type Ierc20metadataCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ierc20metadataTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Ierc20metadataTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ierc20metadataFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Ierc20metadataFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ierc20metadataSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Ierc20metadataSession struct {
	Contract     *Ierc20metadata   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Ierc20metadataCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Ierc20metadataCallerSession struct {
	Contract *Ierc20metadataCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// Ierc20metadataTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Ierc20metadataTransactorSession struct {
	Contract     *Ierc20metadataTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// Ierc20metadataRaw is an auto generated low-level Go binding around an Ethereum contract.
type Ierc20metadataRaw struct {
	Contract *Ierc20metadata // Generic contract binding to access the raw methods on
}

// Ierc20metadataCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Ierc20metadataCallerRaw struct {
	Contract *Ierc20metadataCaller // Generic read-only contract binding to access the raw methods on
}

// Ierc20metadataTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Ierc20metadataTransactorRaw struct {
	Contract *Ierc20metadataTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIerc20metadata creates a new instance of Ierc20metadata, bound to a specific deployed contract.
func NewIerc20metadata(address common.Address, backend bind.ContractBackend) (*Ierc20metadata, error) {
	contract, err := bindIerc20metadata(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ierc20metadata{Ierc20metadataCaller: Ierc20metadataCaller{contract: contract}, Ierc20metadataTransactor: Ierc20metadataTransactor{contract: contract}, Ierc20metadataFilterer: Ierc20metadataFilterer{contract: contract}}, nil
}

// NewIerc20metadataCaller creates a new read-only instance of Ierc20metadata, bound to a specific deployed contract.
func NewIerc20metadataCaller(address common.Address, caller bind.ContractCaller) (*Ierc20metadataCaller, error) {
	contract, err := bindIerc20metadata(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Ierc20metadataCaller{contract: contract}, nil
}

// NewIerc20metadataTransactor creates a new write-only instance of Ierc20metadata, bound to a specific deployed contract.
func NewIerc20metadataTransactor(address common.Address, transactor bind.ContractTransactor) (*Ierc20metadataTransactor, error) {
	contract, err := bindIerc20metadata(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Ierc20metadataTransactor{contract: contract}, nil
}

// NewIerc20metadataFilterer creates a new log filterer instance of Ierc20metadata, bound to a specific deployed contract.
func NewIerc20metadataFilterer(address common.Address, filterer bind.ContractFilterer) (*Ierc20metadataFilterer, error) {
	contract, err := bindIerc20metadata(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Ierc20metadataFilterer{contract: contract}, nil
}

// bindIerc20metadata binds a generic wrapper to an already deployed contract.
func bindIerc20metadata(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Ierc20metadataABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ierc20metadata *Ierc20metadataRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ierc20metadata.Contract.Ierc20metadataCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ierc20metadata *Ierc20metadataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ierc20metadata.Contract.Ierc20metadataTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ierc20metadata *Ierc20metadataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ierc20metadata.Contract.Ierc20metadataTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ierc20metadata *Ierc20metadataCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ierc20metadata.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ierc20metadata *Ierc20metadataTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ierc20metadata.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ierc20metadata *Ierc20metadataTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ierc20metadata.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Ierc20metadata *Ierc20metadataCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Ierc20metadata.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Ierc20metadata *Ierc20metadataSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Ierc20metadata.Contract.Allowance(&_Ierc20metadata.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Ierc20metadata *Ierc20metadataCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Ierc20metadata.Contract.Allowance(&_Ierc20metadata.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Ierc20metadata *Ierc20metadataCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Ierc20metadata.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Ierc20metadata *Ierc20metadataSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Ierc20metadata.Contract.BalanceOf(&_Ierc20metadata.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Ierc20metadata *Ierc20metadataCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Ierc20metadata.Contract.BalanceOf(&_Ierc20metadata.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Ierc20metadata *Ierc20metadataCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Ierc20metadata.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Ierc20metadata *Ierc20metadataSession) Decimals() (uint8, error) {
	return _Ierc20metadata.Contract.Decimals(&_Ierc20metadata.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Ierc20metadata *Ierc20metadataCallerSession) Decimals() (uint8, error) {
	return _Ierc20metadata.Contract.Decimals(&_Ierc20metadata.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Ierc20metadata *Ierc20metadataCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Ierc20metadata.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Ierc20metadata *Ierc20metadataSession) Name() (string, error) {
	return _Ierc20metadata.Contract.Name(&_Ierc20metadata.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Ierc20metadata *Ierc20metadataCallerSession) Name() (string, error) {
	return _Ierc20metadata.Contract.Name(&_Ierc20metadata.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Ierc20metadata *Ierc20metadataCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Ierc20metadata.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Ierc20metadata *Ierc20metadataSession) Symbol() (string, error) {
	return _Ierc20metadata.Contract.Symbol(&_Ierc20metadata.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Ierc20metadata *Ierc20metadataCallerSession) Symbol() (string, error) {
	return _Ierc20metadata.Contract.Symbol(&_Ierc20metadata.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Ierc20metadata *Ierc20metadataCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ierc20metadata.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Ierc20metadata *Ierc20metadataSession) TotalSupply() (*big.Int, error) {
	return _Ierc20metadata.Contract.TotalSupply(&_Ierc20metadata.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Ierc20metadata *Ierc20metadataCallerSession) TotalSupply() (*big.Int, error) {
	return _Ierc20metadata.Contract.TotalSupply(&_Ierc20metadata.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Ierc20metadata *Ierc20metadataTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ierc20metadata.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Ierc20metadata *Ierc20metadataSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ierc20metadata.Contract.Approve(&_Ierc20metadata.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Ierc20metadata *Ierc20metadataTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ierc20metadata.Contract.Approve(&_Ierc20metadata.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Ierc20metadata *Ierc20metadataTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ierc20metadata.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Ierc20metadata *Ierc20metadataSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ierc20metadata.Contract.Transfer(&_Ierc20metadata.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Ierc20metadata *Ierc20metadataTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ierc20metadata.Contract.Transfer(&_Ierc20metadata.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Ierc20metadata *Ierc20metadataTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ierc20metadata.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Ierc20metadata *Ierc20metadataSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ierc20metadata.Contract.TransferFrom(&_Ierc20metadata.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Ierc20metadata *Ierc20metadataTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ierc20metadata.Contract.TransferFrom(&_Ierc20metadata.TransactOpts, from, to, amount)
}

// Ierc20metadataApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Ierc20metadata contract.
type Ierc20metadataApprovalIterator struct {
	Event *Ierc20metadataApproval // Event containing the contract specifics and raw log

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
func (it *Ierc20metadataApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Ierc20metadataApproval)
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
		it.Event = new(Ierc20metadataApproval)
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
func (it *Ierc20metadataApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Ierc20metadataApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Ierc20metadataApproval represents a Approval event raised by the Ierc20metadata contract.
type Ierc20metadataApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Ierc20metadata *Ierc20metadataFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*Ierc20metadataApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Ierc20metadata.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &Ierc20metadataApprovalIterator{contract: _Ierc20metadata.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Ierc20metadata *Ierc20metadataFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Ierc20metadataApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Ierc20metadata.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Ierc20metadataApproval)
				if err := _Ierc20metadata.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Ierc20metadata *Ierc20metadataFilterer) ParseApproval(log types.Log) (*Ierc20metadataApproval, error) {
	event := new(Ierc20metadataApproval)
	if err := _Ierc20metadata.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Ierc20metadataTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Ierc20metadata contract.
type Ierc20metadataTransferIterator struct {
	Event *Ierc20metadataTransfer // Event containing the contract specifics and raw log

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
func (it *Ierc20metadataTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Ierc20metadataTransfer)
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
		it.Event = new(Ierc20metadataTransfer)
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
func (it *Ierc20metadataTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Ierc20metadataTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Ierc20metadataTransfer represents a Transfer event raised by the Ierc20metadata contract.
type Ierc20metadataTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Ierc20metadata *Ierc20metadataFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Ierc20metadataTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Ierc20metadata.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Ierc20metadataTransferIterator{contract: _Ierc20metadata.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Ierc20metadata *Ierc20metadataFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Ierc20metadataTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Ierc20metadata.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Ierc20metadataTransfer)
				if err := _Ierc20metadata.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Ierc20metadata *Ierc20metadataFilterer) ParseTransfer(log types.Log) (*Ierc20metadataTransfer, error) {
	event := new(Ierc20metadataTransfer)
	if err := _Ierc20metadata.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
