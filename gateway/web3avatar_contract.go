// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// Web3AvatarMetaData contains all meta data concerning the Web3Avatar contract.
var Web3AvatarMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_key\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_avatar\",\"type\":\"string\"}],\"name\":\"AvatarUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"getAvatar\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"getAvatar\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_avatar\",\"type\":\"string\"}],\"name\":\"setAvatar\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"b48f88c3": "getAvatar(address,string)",
		"7d840ed5": "getAvatar(string)",
		"d3189e06": "setAvatar(string,string)",
	},
	Bin: "0x608060405234801561001057600080fd5b506105de806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80637d840ed514610046578063b48f88c31461006f578063d3189e0614610082575b600080fd5b61005961005436600461039e565b610097565b60405161006691906103e0565b60405180910390f35b61005961007d366004610435565b61015a565b610095610090366004610496565b610230565b005b33600090815260208190526040908190209051606091906100bb9085908590610502565b908152602001604051809103902080546100d490610512565b80601f016020809104026020016040519081016040528092919081815260200182805461010090610512565b801561014d5780601f106101225761010080835404028352916020019161014d565b820191906000526020600020905b81548152906001019060200180831161013057829003601f168201915b5050505050905092915050565b6060600080856001600160a01b03166001600160a01b031681526020019081526020016000208383604051610190929190610502565b908152602001604051809103902080546101a990610512565b80601f01602080910402602001604051908101604052809291908181526020018280546101d590610512565b80156102225780601f106101f757610100808354040283529160200191610222565b820191906000526020600020905b81548152906001019060200180831161020557829003601f168201915b505050505090509392505050565b33600090815260208190526040908190209051839183916102549088908890610502565b90815260405190819003602001902061026e9290916102bc565b50336001600160a01b03167fa6c0024d6e53437883ab318166d16d7b67b59998f53bcd57333ec53900c3a219858585856040516102ae9493929190610576565b60405180910390a250505050565b8280546102c890610512565b90600052602060002090601f0160209004810192826102ea5760008555610330565b82601f106103035782800160ff19823516178555610330565b82800160010185558215610330579182015b82811115610330578235825591602001919060010190610315565b5061033c929150610340565b5090565b5b8082111561033c5760008155600101610341565b60008083601f84011261036757600080fd5b50813567ffffffffffffffff81111561037f57600080fd5b60208301915083602082850101111561039757600080fd5b9250929050565b600080602083850312156103b157600080fd5b823567ffffffffffffffff8111156103c857600080fd5b6103d485828601610355565b90969095509350505050565b600060208083528351808285015260005b8181101561040d578581018301518582016040015282016103f1565b8181111561041f576000604083870101525b50601f01601f1916929092016040019392505050565b60008060006040848603121561044a57600080fd5b83356001600160a01b038116811461046157600080fd5b9250602084013567ffffffffffffffff81111561047d57600080fd5b61048986828701610355565b9497909650939450505050565b600080600080604085870312156104ac57600080fd5b843567ffffffffffffffff808211156104c457600080fd5b6104d088838901610355565b909650945060208701359150808211156104e957600080fd5b506104f687828801610355565b95989497509550505050565b8183823760009101908152919050565b600181811c9082168061052657607f821691505b6020821081141561054757634e487b7160e01b600052602260045260246000fd5b50919050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b60408152600061058a60408301868861054d565b828103602084015261059d81858761054d565b97965050505050505056fea2646970667358221220d9c4fb3a211af329bb636a3cc8558948c1e6babe8d3d2178611f65b52a16d59f64736f6c634300080a0033",
}

// Web3AvatarABI is the input ABI used to generate the binding from.
// Deprecated: Use Web3AvatarMetaData.ABI instead.
var Web3AvatarABI = Web3AvatarMetaData.ABI

// Deprecated: Use Web3AvatarMetaData.Sigs instead.
// Web3AvatarFuncSigs maps the 4-byte function signature to its string representation.
var Web3AvatarFuncSigs = Web3AvatarMetaData.Sigs

// Web3AvatarBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Web3AvatarMetaData.Bin instead.
var Web3AvatarBin = Web3AvatarMetaData.Bin

// DeployWeb3Avatar deploys a new Ethereum contract, binding an instance of Web3Avatar to it.
func DeployWeb3Avatar(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Web3Avatar, error) {
	parsed, err := Web3AvatarMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Web3AvatarBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Web3Avatar{Web3AvatarCaller: Web3AvatarCaller{contract: contract}, Web3AvatarTransactor: Web3AvatarTransactor{contract: contract}, Web3AvatarFilterer: Web3AvatarFilterer{contract: contract}}, nil
}

// Web3Avatar is an auto generated Go binding around an Ethereum contract.
type Web3Avatar struct {
	Web3AvatarCaller     // Read-only binding to the contract
	Web3AvatarTransactor // Write-only binding to the contract
	Web3AvatarFilterer   // Log filterer for contract events
}

// Web3AvatarCaller is an auto generated read-only Go binding around an Ethereum contract.
type Web3AvatarCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Web3AvatarTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Web3AvatarTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Web3AvatarFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Web3AvatarFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Web3AvatarSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Web3AvatarSession struct {
	Contract     *Web3Avatar       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Web3AvatarCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Web3AvatarCallerSession struct {
	Contract *Web3AvatarCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// Web3AvatarTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Web3AvatarTransactorSession struct {
	Contract     *Web3AvatarTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// Web3AvatarRaw is an auto generated low-level Go binding around an Ethereum contract.
type Web3AvatarRaw struct {
	Contract *Web3Avatar // Generic contract binding to access the raw methods on
}

// Web3AvatarCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Web3AvatarCallerRaw struct {
	Contract *Web3AvatarCaller // Generic read-only contract binding to access the raw methods on
}

// Web3AvatarTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Web3AvatarTransactorRaw struct {
	Contract *Web3AvatarTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWeb3Avatar creates a new instance of Web3Avatar, bound to a specific deployed contract.
func NewWeb3Avatar(address common.Address, backend bind.ContractBackend) (*Web3Avatar, error) {
	contract, err := bindWeb3Avatar(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Web3Avatar{Web3AvatarCaller: Web3AvatarCaller{contract: contract}, Web3AvatarTransactor: Web3AvatarTransactor{contract: contract}, Web3AvatarFilterer: Web3AvatarFilterer{contract: contract}}, nil
}

// NewWeb3AvatarCaller creates a new read-only instance of Web3Avatar, bound to a specific deployed contract.
func NewWeb3AvatarCaller(address common.Address, caller bind.ContractCaller) (*Web3AvatarCaller, error) {
	contract, err := bindWeb3Avatar(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Web3AvatarCaller{contract: contract}, nil
}

// NewWeb3AvatarTransactor creates a new write-only instance of Web3Avatar, bound to a specific deployed contract.
func NewWeb3AvatarTransactor(address common.Address, transactor bind.ContractTransactor) (*Web3AvatarTransactor, error) {
	contract, err := bindWeb3Avatar(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Web3AvatarTransactor{contract: contract}, nil
}

// NewWeb3AvatarFilterer creates a new log filterer instance of Web3Avatar, bound to a specific deployed contract.
func NewWeb3AvatarFilterer(address common.Address, filterer bind.ContractFilterer) (*Web3AvatarFilterer, error) {
	contract, err := bindWeb3Avatar(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Web3AvatarFilterer{contract: contract}, nil
}

// bindWeb3Avatar binds a generic wrapper to an already deployed contract.
func bindWeb3Avatar(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Web3AvatarABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Web3Avatar *Web3AvatarRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Web3Avatar.Contract.Web3AvatarCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Web3Avatar *Web3AvatarRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Web3Avatar.Contract.Web3AvatarTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Web3Avatar *Web3AvatarRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Web3Avatar.Contract.Web3AvatarTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Web3Avatar *Web3AvatarCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Web3Avatar.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Web3Avatar *Web3AvatarTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Web3Avatar.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Web3Avatar *Web3AvatarTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Web3Avatar.Contract.contract.Transact(opts, method, params...)
}

// GetAvatar is a free data retrieval call binding the contract method 0x7d840ed5.
//
// Solidity: function getAvatar(string _name) view returns(string)
func (_Web3Avatar *Web3AvatarCaller) GetAvatar(opts *bind.CallOpts, _name string) (string, error) {
	var out []interface{}
	err := _Web3Avatar.contract.Call(opts, &out, "getAvatar", _name)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetAvatar is a free data retrieval call binding the contract method 0x7d840ed5.
//
// Solidity: function getAvatar(string _name) view returns(string)
func (_Web3Avatar *Web3AvatarSession) GetAvatar(_name string) (string, error) {
	return _Web3Avatar.Contract.GetAvatar(&_Web3Avatar.CallOpts, _name)
}

// GetAvatar is a free data retrieval call binding the contract method 0x7d840ed5.
//
// Solidity: function getAvatar(string _name) view returns(string)
func (_Web3Avatar *Web3AvatarCallerSession) GetAvatar(_name string) (string, error) {
	return _Web3Avatar.Contract.GetAvatar(&_Web3Avatar.CallOpts, _name)
}

// GetAvatar0 is a free data retrieval call binding the contract method 0xb48f88c3.
//
// Solidity: function getAvatar(address _addr, string _name) view returns(string)
func (_Web3Avatar *Web3AvatarCaller) GetAvatar0(opts *bind.CallOpts, _addr common.Address, _name string) (string, error) {
	var out []interface{}
	err := _Web3Avatar.contract.Call(opts, &out, "getAvatar0", _addr, _name)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetAvatar0 is a free data retrieval call binding the contract method 0xb48f88c3.
//
// Solidity: function getAvatar(address _addr, string _name) view returns(string)
func (_Web3Avatar *Web3AvatarSession) GetAvatar0(_addr common.Address, _name string) (string, error) {
	return _Web3Avatar.Contract.GetAvatar0(&_Web3Avatar.CallOpts, _addr, _name)
}

// GetAvatar0 is a free data retrieval call binding the contract method 0xb48f88c3.
//
// Solidity: function getAvatar(address _addr, string _name) view returns(string)
func (_Web3Avatar *Web3AvatarCallerSession) GetAvatar0(_addr common.Address, _name string) (string, error) {
	return _Web3Avatar.Contract.GetAvatar0(&_Web3Avatar.CallOpts, _addr, _name)
}

// SetAvatar is a paid mutator transaction binding the contract method 0xd3189e06.
//
// Solidity: function setAvatar(string _name, string _avatar) returns()
func (_Web3Avatar *Web3AvatarTransactor) SetAvatar(opts *bind.TransactOpts, _name string, _avatar string) (*types.Transaction, error) {
	return _Web3Avatar.contract.Transact(opts, "setAvatar", _name, _avatar)
}

// SetAvatar is a paid mutator transaction binding the contract method 0xd3189e06.
//
// Solidity: function setAvatar(string _name, string _avatar) returns()
func (_Web3Avatar *Web3AvatarSession) SetAvatar(_name string, _avatar string) (*types.Transaction, error) {
	return _Web3Avatar.Contract.SetAvatar(&_Web3Avatar.TransactOpts, _name, _avatar)
}

// SetAvatar is a paid mutator transaction binding the contract method 0xd3189e06.
//
// Solidity: function setAvatar(string _name, string _avatar) returns()
func (_Web3Avatar *Web3AvatarTransactorSession) SetAvatar(_name string, _avatar string) (*types.Transaction, error) {
	return _Web3Avatar.Contract.SetAvatar(&_Web3Avatar.TransactOpts, _name, _avatar)
}

// Web3AvatarAvatarUpdatedIterator is returned from FilterAvatarUpdated and is used to iterate over the raw logs and unpacked data for AvatarUpdated events raised by the Web3Avatar contract.
type Web3AvatarAvatarUpdatedIterator struct {
	Event *Web3AvatarAvatarUpdated // Event containing the contract specifics and raw log

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
func (it *Web3AvatarAvatarUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Web3AvatarAvatarUpdated)
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
		it.Event = new(Web3AvatarAvatarUpdated)
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
func (it *Web3AvatarAvatarUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Web3AvatarAvatarUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Web3AvatarAvatarUpdated represents a AvatarUpdated event raised by the Web3Avatar contract.
type Web3AvatarAvatarUpdated struct {
	Key    common.Address
	Name   string
	Avatar string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAvatarUpdated is a free log retrieval operation binding the contract event 0xa6c0024d6e53437883ab318166d16d7b67b59998f53bcd57333ec53900c3a219.
//
// Solidity: event AvatarUpdated(address indexed _key, string _name, string _avatar)
func (_Web3Avatar *Web3AvatarFilterer) FilterAvatarUpdated(opts *bind.FilterOpts, _key []common.Address) (*Web3AvatarAvatarUpdatedIterator, error) {

	var _keyRule []interface{}
	for _, _keyItem := range _key {
		_keyRule = append(_keyRule, _keyItem)
	}

	logs, sub, err := _Web3Avatar.contract.FilterLogs(opts, "AvatarUpdated", _keyRule)
	if err != nil {
		return nil, err
	}
	return &Web3AvatarAvatarUpdatedIterator{contract: _Web3Avatar.contract, event: "AvatarUpdated", logs: logs, sub: sub}, nil
}

// WatchAvatarUpdated is a free log subscription operation binding the contract event 0xa6c0024d6e53437883ab318166d16d7b67b59998f53bcd57333ec53900c3a219.
//
// Solidity: event AvatarUpdated(address indexed _key, string _name, string _avatar)
func (_Web3Avatar *Web3AvatarFilterer) WatchAvatarUpdated(opts *bind.WatchOpts, sink chan<- *Web3AvatarAvatarUpdated, _key []common.Address) (event.Subscription, error) {

	var _keyRule []interface{}
	for _, _keyItem := range _key {
		_keyRule = append(_keyRule, _keyItem)
	}

	logs, sub, err := _Web3Avatar.contract.WatchLogs(opts, "AvatarUpdated", _keyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Web3AvatarAvatarUpdated)
				if err := _Web3Avatar.contract.UnpackLog(event, "AvatarUpdated", log); err != nil {
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

// ParseAvatarUpdated is a log parse operation binding the contract event 0xa6c0024d6e53437883ab318166d16d7b67b59998f53bcd57333ec53900c3a219.
//
// Solidity: event AvatarUpdated(address indexed _key, string _name, string _avatar)
func (_Web3Avatar *Web3AvatarFilterer) ParseAvatarUpdated(log types.Log) (*Web3AvatarAvatarUpdated, error) {
	event := new(Web3AvatarAvatarUpdated)
	if err := _Web3Avatar.contract.UnpackLog(event, "AvatarUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
