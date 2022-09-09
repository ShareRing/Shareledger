// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package approver

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

// ApproverMetaData contains all meta data concerning the Approver contract.
var ApproverMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_approvers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"addApprover\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"addApprovers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"approver\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"removeApprover\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"removeApprovers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ApproverABI is the input ABI used to generate the binding from.
// Deprecated: Use ApproverMetaData.ABI instead.
var ApproverABI = ApproverMetaData.ABI

// Approver is an auto generated Go binding around an Ethereum contract.
type Approver struct {
	ApproverCaller     // Read-only binding to the contract
	ApproverTransactor // Write-only binding to the contract
	ApproverFilterer   // Log filterer for contract events
}

// ApproverCaller is an auto generated read-only Go binding around an Ethereum contract.
type ApproverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApproverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ApproverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApproverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ApproverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApproverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ApproverSession struct {
	Contract     *Approver         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApproverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ApproverCallerSession struct {
	Contract *ApproverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ApproverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ApproverTransactorSession struct {
	Contract     *ApproverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ApproverRaw is an auto generated low-level Go binding around an Ethereum contract.
type ApproverRaw struct {
	Contract *Approver // Generic contract binding to access the raw methods on
}

// ApproverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ApproverCallerRaw struct {
	Contract *ApproverCaller // Generic read-only contract binding to access the raw methods on
}

// ApproverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ApproverTransactorRaw struct {
	Contract *ApproverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewApprover creates a new instance of Approver, bound to a specific deployed contract.
func NewApprover(address common.Address, backend bind.ContractBackend) (*Approver, error) {
	contract, err := bindApprover(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Approver{ApproverCaller: ApproverCaller{contract: contract}, ApproverTransactor: ApproverTransactor{contract: contract}, ApproverFilterer: ApproverFilterer{contract: contract}}, nil
}

// NewApproverCaller creates a new read-only instance of Approver, bound to a specific deployed contract.
func NewApproverCaller(address common.Address, caller bind.ContractCaller) (*ApproverCaller, error) {
	contract, err := bindApprover(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ApproverCaller{contract: contract}, nil
}

// NewApproverTransactor creates a new write-only instance of Approver, bound to a specific deployed contract.
func NewApproverTransactor(address common.Address, transactor bind.ContractTransactor) (*ApproverTransactor, error) {
	contract, err := bindApprover(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ApproverTransactor{contract: contract}, nil
}

// NewApproverFilterer creates a new log filterer instance of Approver, bound to a specific deployed contract.
func NewApproverFilterer(address common.Address, filterer bind.ContractFilterer) (*ApproverFilterer, error) {
	contract, err := bindApprover(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ApproverFilterer{contract: contract}, nil
}

// bindApprover binds a generic wrapper to an already deployed contract.
func bindApprover(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ApproverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Approver *ApproverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Approver.Contract.ApproverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Approver *ApproverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Approver.Contract.ApproverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Approver *ApproverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Approver.Contract.ApproverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Approver *ApproverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Approver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Approver *ApproverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Approver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Approver *ApproverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Approver.Contract.contract.Transact(opts, method, params...)
}

// Approvers is a free data retrieval call binding the contract method 0x3ab970e0.
//
// Solidity: function _approvers(address ) view returns(bool)
func (_Approver *ApproverCaller) Approvers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Approver.contract.Call(opts, &out, "_approvers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Approvers is a free data retrieval call binding the contract method 0x3ab970e0.
//
// Solidity: function _approvers(address ) view returns(bool)
func (_Approver *ApproverSession) Approvers(arg0 common.Address) (bool, error) {
	return _Approver.Contract.Approvers(&_Approver.CallOpts, arg0)
}

// Approvers is a free data retrieval call binding the contract method 0x3ab970e0.
//
// Solidity: function _approvers(address ) view returns(bool)
func (_Approver *ApproverCallerSession) Approvers(arg0 common.Address) (bool, error) {
	return _Approver.Contract.Approvers(&_Approver.CallOpts, arg0)
}

// Approver is a free data retrieval call binding the contract method 0x3f131d9f.
//
// Solidity: function approver(address _address) view returns(bool)
func (_Approver *ApproverCaller) Approver(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var out []interface{}
	err := _Approver.contract.Call(opts, &out, "approver", _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Approver is a free data retrieval call binding the contract method 0x3f131d9f.
//
// Solidity: function approver(address _address) view returns(bool)
func (_Approver *ApproverSession) Approver(_address common.Address) (bool, error) {
	return _Approver.Contract.Approver(&_Approver.CallOpts, _address)
}

// Approver is a free data retrieval call binding the contract method 0x3f131d9f.
//
// Solidity: function approver(address _address) view returns(bool)
func (_Approver *ApproverCallerSession) Approver(_address common.Address) (bool, error) {
	return _Approver.Contract.Approver(&_Approver.CallOpts, _address)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Approver *ApproverCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Approver.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Approver *ApproverSession) Owner() (common.Address, error) {
	return _Approver.Contract.Owner(&_Approver.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Approver *ApproverCallerSession) Owner() (common.Address, error) {
	return _Approver.Contract.Owner(&_Approver.CallOpts)
}

// AddApprover is a paid mutator transaction binding the contract method 0xb646c194.
//
// Solidity: function addApprover(address _address) returns()
func (_Approver *ApproverTransactor) AddApprover(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Approver.contract.Transact(opts, "addApprover", _address)
}

// AddApprover is a paid mutator transaction binding the contract method 0xb646c194.
//
// Solidity: function addApprover(address _address) returns()
func (_Approver *ApproverSession) AddApprover(_address common.Address) (*types.Transaction, error) {
	return _Approver.Contract.AddApprover(&_Approver.TransactOpts, _address)
}

// AddApprover is a paid mutator transaction binding the contract method 0xb646c194.
//
// Solidity: function addApprover(address _address) returns()
func (_Approver *ApproverTransactorSession) AddApprover(_address common.Address) (*types.Transaction, error) {
	return _Approver.Contract.AddApprover(&_Approver.TransactOpts, _address)
}

// AddApprovers is a paid mutator transaction binding the contract method 0x6a882d51.
//
// Solidity: function addApprovers(address[] _addresses) returns()
func (_Approver *ApproverTransactor) AddApprovers(opts *bind.TransactOpts, _addresses []common.Address) (*types.Transaction, error) {
	return _Approver.contract.Transact(opts, "addApprovers", _addresses)
}

// AddApprovers is a paid mutator transaction binding the contract method 0x6a882d51.
//
// Solidity: function addApprovers(address[] _addresses) returns()
func (_Approver *ApproverSession) AddApprovers(_addresses []common.Address) (*types.Transaction, error) {
	return _Approver.Contract.AddApprovers(&_Approver.TransactOpts, _addresses)
}

// AddApprovers is a paid mutator transaction binding the contract method 0x6a882d51.
//
// Solidity: function addApprovers(address[] _addresses) returns()
func (_Approver *ApproverTransactorSession) AddApprovers(_addresses []common.Address) (*types.Transaction, error) {
	return _Approver.Contract.AddApprovers(&_Approver.TransactOpts, _addresses)
}

// RemoveApprover is a paid mutator transaction binding the contract method 0x6cf4c88f.
//
// Solidity: function removeApprover(address _address) returns()
func (_Approver *ApproverTransactor) RemoveApprover(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Approver.contract.Transact(opts, "removeApprover", _address)
}

// RemoveApprover is a paid mutator transaction binding the contract method 0x6cf4c88f.
//
// Solidity: function removeApprover(address _address) returns()
func (_Approver *ApproverSession) RemoveApprover(_address common.Address) (*types.Transaction, error) {
	return _Approver.Contract.RemoveApprover(&_Approver.TransactOpts, _address)
}

// RemoveApprover is a paid mutator transaction binding the contract method 0x6cf4c88f.
//
// Solidity: function removeApprover(address _address) returns()
func (_Approver *ApproverTransactorSession) RemoveApprover(_address common.Address) (*types.Transaction, error) {
	return _Approver.Contract.RemoveApprover(&_Approver.TransactOpts, _address)
}

// RemoveApprovers is a paid mutator transaction binding the contract method 0x7569d66f.
//
// Solidity: function removeApprovers(address[] _addresses) returns()
func (_Approver *ApproverTransactor) RemoveApprovers(opts *bind.TransactOpts, _addresses []common.Address) (*types.Transaction, error) {
	return _Approver.contract.Transact(opts, "removeApprovers", _addresses)
}

// RemoveApprovers is a paid mutator transaction binding the contract method 0x7569d66f.
//
// Solidity: function removeApprovers(address[] _addresses) returns()
func (_Approver *ApproverSession) RemoveApprovers(_addresses []common.Address) (*types.Transaction, error) {
	return _Approver.Contract.RemoveApprovers(&_Approver.TransactOpts, _addresses)
}

// RemoveApprovers is a paid mutator transaction binding the contract method 0x7569d66f.
//
// Solidity: function removeApprovers(address[] _addresses) returns()
func (_Approver *ApproverTransactorSession) RemoveApprovers(_addresses []common.Address) (*types.Transaction, error) {
	return _Approver.Contract.RemoveApprovers(&_Approver.TransactOpts, _addresses)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Approver *ApproverTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Approver.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Approver *ApproverSession) RenounceOwnership() (*types.Transaction, error) {
	return _Approver.Contract.RenounceOwnership(&_Approver.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Approver *ApproverTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Approver.Contract.RenounceOwnership(&_Approver.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Approver *ApproverTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Approver.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Approver *ApproverSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Approver.Contract.TransferOwnership(&_Approver.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Approver *ApproverTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Approver.Contract.TransferOwnership(&_Approver.TransactOpts, newOwner)
}

// ApproverOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Approver contract.
type ApproverOwnershipTransferredIterator struct {
	Event *ApproverOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ApproverOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApproverOwnershipTransferred)
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
		it.Event = new(ApproverOwnershipTransferred)
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
func (it *ApproverOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApproverOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApproverOwnershipTransferred represents a OwnershipTransferred event raised by the Approver contract.
type ApproverOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Approver *ApproverFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ApproverOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Approver.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ApproverOwnershipTransferredIterator{contract: _Approver.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Approver *ApproverFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ApproverOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Approver.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApproverOwnershipTransferred)
				if err := _Approver.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Approver *ApproverFilterer) ParseOwnershipTransferred(log types.Log) (*ApproverOwnershipTransferred, error) {
	event := new(ApproverOwnershipTransferred)
	if err := _Approver.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
