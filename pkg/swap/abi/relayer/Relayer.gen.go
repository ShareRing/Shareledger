// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package relayer

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

// RelayerMetaData contains all meta data concerning the Relayer contract.
var RelayerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"relayer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setRelayer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unsetRelayer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// RelayerABI is the input ABI used to generate the binding from.
// Deprecated: Use RelayerMetaData.ABI instead.
var RelayerABI = RelayerMetaData.ABI

// Relayer is an auto generated Go binding around an Ethereum contract.
type Relayer struct {
	RelayerCaller     // Read-only binding to the contract
	RelayerTransactor // Write-only binding to the contract
	RelayerFilterer   // Log filterer for contract events
}

// RelayerCaller is an auto generated read-only Go binding around an Ethereum contract.
type RelayerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RelayerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RelayerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RelayerSession struct {
	Contract     *Relayer          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RelayerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RelayerCallerSession struct {
	Contract *RelayerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// RelayerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RelayerTransactorSession struct {
	Contract     *RelayerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// RelayerRaw is an auto generated low-level Go binding around an Ethereum contract.
type RelayerRaw struct {
	Contract *Relayer // Generic contract binding to access the raw methods on
}

// RelayerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RelayerCallerRaw struct {
	Contract *RelayerCaller // Generic read-only contract binding to access the raw methods on
}

// RelayerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RelayerTransactorRaw struct {
	Contract *RelayerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRelayer creates a new instance of Relayer, bound to a specific deployed contract.
func NewRelayer(address common.Address, backend bind.ContractBackend) (*Relayer, error) {
	contract, err := bindRelayer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Relayer{RelayerCaller: RelayerCaller{contract: contract}, RelayerTransactor: RelayerTransactor{contract: contract}, RelayerFilterer: RelayerFilterer{contract: contract}}, nil
}

// NewRelayerCaller creates a new read-only instance of Relayer, bound to a specific deployed contract.
func NewRelayerCaller(address common.Address, caller bind.ContractCaller) (*RelayerCaller, error) {
	contract, err := bindRelayer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RelayerCaller{contract: contract}, nil
}

// NewRelayerTransactor creates a new write-only instance of Relayer, bound to a specific deployed contract.
func NewRelayerTransactor(address common.Address, transactor bind.ContractTransactor) (*RelayerTransactor, error) {
	contract, err := bindRelayer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RelayerTransactor{contract: contract}, nil
}

// NewRelayerFilterer creates a new log filterer instance of Relayer, bound to a specific deployed contract.
func NewRelayerFilterer(address common.Address, filterer bind.ContractFilterer) (*RelayerFilterer, error) {
	contract, err := bindRelayer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RelayerFilterer{contract: contract}, nil
}

// bindRelayer binds a generic wrapper to an already deployed contract.
func bindRelayer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RelayerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Relayer *RelayerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Relayer.Contract.RelayerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Relayer *RelayerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Relayer.Contract.RelayerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Relayer *RelayerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Relayer.Contract.RelayerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Relayer *RelayerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Relayer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Relayer *RelayerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Relayer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Relayer *RelayerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Relayer.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Relayer *RelayerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Relayer.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Relayer *RelayerSession) Owner() (common.Address, error) {
	return _Relayer.Contract.Owner(&_Relayer.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Relayer *RelayerCallerSession) Owner() (common.Address, error) {
	return _Relayer.Contract.Owner(&_Relayer.CallOpts)
}

// Relayer is a free data retrieval call binding the contract method 0x8406c079.
//
// Solidity: function relayer() view returns(address)
func (_Relayer *RelayerCaller) Relayer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Relayer.contract.Call(opts, &out, "relayer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Relayer is a free data retrieval call binding the contract method 0x8406c079.
//
// Solidity: function relayer() view returns(address)
func (_Relayer *RelayerSession) Relayer() (common.Address, error) {
	return _Relayer.Contract.Relayer(&_Relayer.CallOpts)
}

// Relayer is a free data retrieval call binding the contract method 0x8406c079.
//
// Solidity: function relayer() view returns(address)
func (_Relayer *RelayerCallerSession) Relayer() (common.Address, error) {
	return _Relayer.Contract.Relayer(&_Relayer.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Relayer *RelayerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Relayer.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Relayer *RelayerSession) RenounceOwnership() (*types.Transaction, error) {
	return _Relayer.Contract.RenounceOwnership(&_Relayer.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Relayer *RelayerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Relayer.Contract.RenounceOwnership(&_Relayer.TransactOpts)
}

// SetRelayer is a paid mutator transaction binding the contract method 0x6548e9bc.
//
// Solidity: function setRelayer(address _address) returns()
func (_Relayer *RelayerTransactor) SetRelayer(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Relayer.contract.Transact(opts, "setRelayer", _address)
}

// SetRelayer is a paid mutator transaction binding the contract method 0x6548e9bc.
//
// Solidity: function setRelayer(address _address) returns()
func (_Relayer *RelayerSession) SetRelayer(_address common.Address) (*types.Transaction, error) {
	return _Relayer.Contract.SetRelayer(&_Relayer.TransactOpts, _address)
}

// SetRelayer is a paid mutator transaction binding the contract method 0x6548e9bc.
//
// Solidity: function setRelayer(address _address) returns()
func (_Relayer *RelayerTransactorSession) SetRelayer(_address common.Address) (*types.Transaction, error) {
	return _Relayer.Contract.SetRelayer(&_Relayer.TransactOpts, _address)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Relayer *RelayerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Relayer.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Relayer *RelayerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Relayer.Contract.TransferOwnership(&_Relayer.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Relayer *RelayerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Relayer.Contract.TransferOwnership(&_Relayer.TransactOpts, newOwner)
}

// UnsetRelayer is a paid mutator transaction binding the contract method 0xa521d4df.
//
// Solidity: function unsetRelayer() returns()
func (_Relayer *RelayerTransactor) UnsetRelayer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Relayer.contract.Transact(opts, "unsetRelayer")
}

// UnsetRelayer is a paid mutator transaction binding the contract method 0xa521d4df.
//
// Solidity: function unsetRelayer() returns()
func (_Relayer *RelayerSession) UnsetRelayer() (*types.Transaction, error) {
	return _Relayer.Contract.UnsetRelayer(&_Relayer.TransactOpts)
}

// UnsetRelayer is a paid mutator transaction binding the contract method 0xa521d4df.
//
// Solidity: function unsetRelayer() returns()
func (_Relayer *RelayerTransactorSession) UnsetRelayer() (*types.Transaction, error) {
	return _Relayer.Contract.UnsetRelayer(&_Relayer.TransactOpts)
}

// RelayerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Relayer contract.
type RelayerOwnershipTransferredIterator struct {
	Event *RelayerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RelayerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayerOwnershipTransferred)
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
		it.Event = new(RelayerOwnershipTransferred)
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
func (it *RelayerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayerOwnershipTransferred represents a OwnershipTransferred event raised by the Relayer contract.
type RelayerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Relayer *RelayerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RelayerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Relayer.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RelayerOwnershipTransferredIterator{contract: _Relayer.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Relayer *RelayerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RelayerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Relayer.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayerOwnershipTransferred)
				if err := _Relayer.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Relayer *RelayerFilterer) ParseOwnershipTransferred(log types.Log) (*RelayerOwnershipTransferred, error) {
	event := new(RelayerOwnershipTransferred)
	if err := _Relayer.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
