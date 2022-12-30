// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dmeet

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

// DmeetMetaData contains all meta data concerning the Dmeet contract.
var DmeetMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"viewer\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"participant\",\"type\":\"uint256\"}],\"name\":\"Membership\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"creators\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\",\"payable\":true},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"buyMembership\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\",\"payable\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_viewerPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_participantPrice\",\"type\":\"uint256\"}],\"name\":\"createMembership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// DmeetABI is the input ABI used to generate the binding from.
// Deprecated: Use DmeetMetaData.ABI instead.
var DmeetABI = DmeetMetaData.ABI

// Dmeet is an auto generated Go binding around an Ethereum contract.
type Dmeet struct {
	DmeetCaller     // Read-only binding to the contract
	DmeetTransactor // Write-only binding to the contract
	DmeetFilterer   // Log filterer for contract events
}

// DmeetCaller is an auto generated read-only Go binding around an Ethereum contract.
type DmeetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DmeetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DmeetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DmeetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DmeetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DmeetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DmeetSession struct {
	Contract     *Dmeet            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DmeetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DmeetCallerSession struct {
	Contract *DmeetCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DmeetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DmeetTransactorSession struct {
	Contract     *DmeetTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DmeetRaw is an auto generated low-level Go binding around an Ethereum contract.
type DmeetRaw struct {
	Contract *Dmeet // Generic contract binding to access the raw methods on
}

// DmeetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DmeetCallerRaw struct {
	Contract *DmeetCaller // Generic read-only contract binding to access the raw methods on
}

// DmeetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DmeetTransactorRaw struct {
	Contract *DmeetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDmeet creates a new instance of Dmeet, bound to a specific deployed contract.
func NewDmeet(address common.Address, backend bind.ContractBackend) (*Dmeet, error) {
	contract, err := bindDmeet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dmeet{DmeetCaller: DmeetCaller{contract: contract}, DmeetTransactor: DmeetTransactor{contract: contract}, DmeetFilterer: DmeetFilterer{contract: contract}}, nil
}

// NewDmeetCaller creates a new read-only instance of Dmeet, bound to a specific deployed contract.
func NewDmeetCaller(address common.Address, caller bind.ContractCaller) (*DmeetCaller, error) {
	contract, err := bindDmeet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DmeetCaller{contract: contract}, nil
}

// NewDmeetTransactor creates a new write-only instance of Dmeet, bound to a specific deployed contract.
func NewDmeetTransactor(address common.Address, transactor bind.ContractTransactor) (*DmeetTransactor, error) {
	contract, err := bindDmeet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DmeetTransactor{contract: contract}, nil
}

// NewDmeetFilterer creates a new log filterer instance of Dmeet, bound to a specific deployed contract.
func NewDmeetFilterer(address common.Address, filterer bind.ContractFilterer) (*DmeetFilterer, error) {
	contract, err := bindDmeet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DmeetFilterer{contract: contract}, nil
}

// bindDmeet binds a generic wrapper to an already deployed contract.
func bindDmeet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DmeetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dmeet *DmeetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dmeet.Contract.DmeetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dmeet *DmeetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dmeet.Contract.DmeetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dmeet *DmeetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dmeet.Contract.DmeetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dmeet *DmeetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dmeet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dmeet *DmeetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dmeet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dmeet *DmeetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dmeet.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address _account, uint256 _id) view returns(uint256)
func (_Dmeet *DmeetCaller) BalanceOf(opts *bind.CallOpts, _account common.Address, _id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Dmeet.contract.Call(opts, &out, "balanceOf", _account, _id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address _account, uint256 _id) view returns(uint256)
func (_Dmeet *DmeetSession) BalanceOf(_account common.Address, _id *big.Int) (*big.Int, error) {
	return _Dmeet.Contract.BalanceOf(&_Dmeet.CallOpts, _account, _id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address _account, uint256 _id) view returns(uint256)
func (_Dmeet *DmeetCallerSession) BalanceOf(_account common.Address, _id *big.Int) (*big.Int, error) {
	return _Dmeet.Contract.BalanceOf(&_Dmeet.CallOpts, _account, _id)
}

// Creators is a free data retrieval call binding the contract method 0xcd53d08e.
//
// Solidity: function creators(uint256 ) view returns(address)
func (_Dmeet *DmeetCaller) Creators(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Dmeet.contract.Call(opts, &out, "creators", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Creators is a free data retrieval call binding the contract method 0xcd53d08e.
//
// Solidity: function creators(uint256 ) view returns(address)
func (_Dmeet *DmeetSession) Creators(arg0 *big.Int) (common.Address, error) {
	return _Dmeet.Contract.Creators(&_Dmeet.CallOpts, arg0)
}

// Creators is a free data retrieval call binding the contract method 0xcd53d08e.
//
// Solidity: function creators(uint256 ) view returns(address)
func (_Dmeet *DmeetCallerSession) Creators(arg0 *big.Int) (common.Address, error) {
	return _Dmeet.Contract.Creators(&_Dmeet.CallOpts, arg0)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Dmeet *DmeetCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dmeet.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Dmeet *DmeetSession) Implementation() (common.Address, error) {
	return _Dmeet.Contract.Implementation(&_Dmeet.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Dmeet *DmeetCallerSession) Implementation() (common.Address, error) {
	return _Dmeet.Contract.Implementation(&_Dmeet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dmeet *DmeetCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dmeet.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dmeet *DmeetSession) Owner() (common.Address, error) {
	return _Dmeet.Contract.Owner(&_Dmeet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dmeet *DmeetCallerSession) Owner() (common.Address, error) {
	return _Dmeet.Contract.Owner(&_Dmeet.CallOpts)
}

// Owner0 is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dmeet *DmeetCaller) Owner0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dmeet.contract.Call(opts, &out, "owner0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner0 is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dmeet *DmeetSession) Owner0() (common.Address, error) {
	return _Dmeet.Contract.Owner0(&_Dmeet.CallOpts)
}

// Owner0 is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dmeet *DmeetCallerSession) Owner0() (common.Address, error) {
	return _Dmeet.Contract.Owner0(&_Dmeet.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Dmeet *DmeetCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Dmeet.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Dmeet *DmeetSession) ProxiableUUID() ([32]byte, error) {
	return _Dmeet.Contract.ProxiableUUID(&_Dmeet.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Dmeet *DmeetCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Dmeet.Contract.ProxiableUUID(&_Dmeet.CallOpts)
}

// ProxiableUUID0 is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Dmeet *DmeetCaller) ProxiableUUID0(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Dmeet.contract.Call(opts, &out, "proxiableUUID0")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID0 is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Dmeet *DmeetSession) ProxiableUUID0() ([32]byte, error) {
	return _Dmeet.Contract.ProxiableUUID0(&_Dmeet.CallOpts)
}

// ProxiableUUID0 is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Dmeet *DmeetCallerSession) ProxiableUUID0() ([32]byte, error) {
	return _Dmeet.Contract.ProxiableUUID0(&_Dmeet.CallOpts)
}

// ProxiableUUID1 is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Dmeet *DmeetCaller) ProxiableUUID1(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Dmeet.contract.Call(opts, &out, "proxiableUUID1")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID1 is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Dmeet *DmeetSession) ProxiableUUID1() ([32]byte, error) {
	return _Dmeet.Contract.ProxiableUUID1(&_Dmeet.CallOpts)
}

// ProxiableUUID1 is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Dmeet *DmeetCallerSession) ProxiableUUID1() ([32]byte, error) {
	return _Dmeet.Contract.ProxiableUUID1(&_Dmeet.CallOpts)
}

// TokenPrice is a free data retrieval call binding the contract method 0xd4ddce8a.
//
// Solidity: function tokenPrice(uint256 ) view returns(uint256)
func (_Dmeet *DmeetCaller) TokenPrice(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Dmeet.contract.Call(opts, &out, "tokenPrice", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenPrice is a free data retrieval call binding the contract method 0xd4ddce8a.
//
// Solidity: function tokenPrice(uint256 ) view returns(uint256)
func (_Dmeet *DmeetSession) TokenPrice(arg0 *big.Int) (*big.Int, error) {
	return _Dmeet.Contract.TokenPrice(&_Dmeet.CallOpts, arg0)
}

// TokenPrice is a free data retrieval call binding the contract method 0xd4ddce8a.
//
// Solidity: function tokenPrice(uint256 ) view returns(uint256)
func (_Dmeet *DmeetCallerSession) TokenPrice(arg0 *big.Int) (*big.Int, error) {
	return _Dmeet.Contract.TokenPrice(&_Dmeet.CallOpts, arg0)
}

// BuyMembership is a paid mutator transaction binding the contract method 0xa4323128.
//
// Solidity: function buyMembership(uint256 _id) payable returns()
func (_Dmeet *DmeetTransactor) BuyMembership(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _Dmeet.contract.Transact(opts, "buyMembership", _id)
}

// BuyMembership is a paid mutator transaction binding the contract method 0xa4323128.
//
// Solidity: function buyMembership(uint256 _id) payable returns()
func (_Dmeet *DmeetSession) BuyMembership(_id *big.Int) (*types.Transaction, error) {
	return _Dmeet.Contract.BuyMembership(&_Dmeet.TransactOpts, _id)
}

// BuyMembership is a paid mutator transaction binding the contract method 0xa4323128.
//
// Solidity: function buyMembership(uint256 _id) payable returns()
func (_Dmeet *DmeetTransactorSession) BuyMembership(_id *big.Int) (*types.Transaction, error) {
	return _Dmeet.Contract.BuyMembership(&_Dmeet.TransactOpts, _id)
}

// CreateMembership is a paid mutator transaction binding the contract method 0x471eb94c.
//
// Solidity: function createMembership(uint256 _viewerPrice, uint256 _participantPrice) returns()
func (_Dmeet *DmeetTransactor) CreateMembership(opts *bind.TransactOpts, _viewerPrice *big.Int, _participantPrice *big.Int) (*types.Transaction, error) {
	return _Dmeet.contract.Transact(opts, "createMembership", _viewerPrice, _participantPrice)
}

// CreateMembership is a paid mutator transaction binding the contract method 0x471eb94c.
//
// Solidity: function createMembership(uint256 _viewerPrice, uint256 _participantPrice) returns()
func (_Dmeet *DmeetSession) CreateMembership(_viewerPrice *big.Int, _participantPrice *big.Int) (*types.Transaction, error) {
	return _Dmeet.Contract.CreateMembership(&_Dmeet.TransactOpts, _viewerPrice, _participantPrice)
}

// CreateMembership is a paid mutator transaction binding the contract method 0x471eb94c.
//
// Solidity: function createMembership(uint256 _viewerPrice, uint256 _participantPrice) returns()
func (_Dmeet *DmeetTransactorSession) CreateMembership(_viewerPrice *big.Int, _participantPrice *big.Int) (*types.Transaction, error) {
	return _Dmeet.Contract.CreateMembership(&_Dmeet.TransactOpts, _viewerPrice, _participantPrice)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Dmeet *DmeetTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dmeet.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Dmeet *DmeetSession) Initialize() (*types.Transaction, error) {
	return _Dmeet.Contract.Initialize(&_Dmeet.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Dmeet *DmeetTransactorSession) Initialize() (*types.Transaction, error) {
	return _Dmeet.Contract.Initialize(&_Dmeet.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dmeet *DmeetTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dmeet.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dmeet *DmeetSession) RenounceOwnership() (*types.Transaction, error) {
	return _Dmeet.Contract.RenounceOwnership(&_Dmeet.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dmeet *DmeetTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Dmeet.Contract.RenounceOwnership(&_Dmeet.TransactOpts)
}

// RenounceOwnership0 is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dmeet *DmeetTransactor) RenounceOwnership0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dmeet.contract.Transact(opts, "renounceOwnership0")
}

// RenounceOwnership0 is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dmeet *DmeetSession) RenounceOwnership0() (*types.Transaction, error) {
	return _Dmeet.Contract.RenounceOwnership0(&_Dmeet.TransactOpts)
}

// RenounceOwnership0 is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dmeet *DmeetTransactorSession) RenounceOwnership0() (*types.Transaction, error) {
	return _Dmeet.Contract.RenounceOwnership0(&_Dmeet.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dmeet *DmeetTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Dmeet.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dmeet *DmeetSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dmeet.Contract.TransferOwnership(&_Dmeet.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dmeet *DmeetTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dmeet.Contract.TransferOwnership(&_Dmeet.TransactOpts, newOwner)
}

// TransferOwnership0 is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dmeet *DmeetTransactor) TransferOwnership0(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Dmeet.contract.Transact(opts, "transferOwnership0", newOwner)
}

// TransferOwnership0 is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dmeet *DmeetSession) TransferOwnership0(newOwner common.Address) (*types.Transaction, error) {
	return _Dmeet.Contract.TransferOwnership0(&_Dmeet.TransactOpts, newOwner)
}

// TransferOwnership0 is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dmeet *DmeetTransactorSession) TransferOwnership0(newOwner common.Address) (*types.Transaction, error) {
	return _Dmeet.Contract.TransferOwnership0(&_Dmeet.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Dmeet *DmeetTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Dmeet.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Dmeet *DmeetSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Dmeet.Contract.UpgradeTo(&_Dmeet.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Dmeet *DmeetTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Dmeet.Contract.UpgradeTo(&_Dmeet.TransactOpts, newImplementation)
}

// UpgradeTo0 is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Dmeet *DmeetTransactor) UpgradeTo0(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Dmeet.contract.Transact(opts, "upgradeTo0", newImplementation)
}

// UpgradeTo0 is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Dmeet *DmeetSession) UpgradeTo0(newImplementation common.Address) (*types.Transaction, error) {
	return _Dmeet.Contract.UpgradeTo0(&_Dmeet.TransactOpts, newImplementation)
}

// UpgradeTo0 is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Dmeet *DmeetTransactorSession) UpgradeTo0(newImplementation common.Address) (*types.Transaction, error) {
	return _Dmeet.Contract.UpgradeTo0(&_Dmeet.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Dmeet *DmeetTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Dmeet.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Dmeet *DmeetSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Dmeet.Contract.UpgradeToAndCall(&_Dmeet.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Dmeet *DmeetTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Dmeet.Contract.UpgradeToAndCall(&_Dmeet.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall0 is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Dmeet *DmeetTransactor) UpgradeToAndCall0(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Dmeet.contract.Transact(opts, "upgradeToAndCall0", newImplementation, data)
}

// UpgradeToAndCall0 is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Dmeet *DmeetSession) UpgradeToAndCall0(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Dmeet.Contract.UpgradeToAndCall0(&_Dmeet.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall0 is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Dmeet *DmeetTransactorSession) UpgradeToAndCall0(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Dmeet.Contract.UpgradeToAndCall0(&_Dmeet.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Dmeet *DmeetTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dmeet.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Dmeet *DmeetSession) Withdraw() (*types.Transaction, error) {
	return _Dmeet.Contract.Withdraw(&_Dmeet.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Dmeet *DmeetTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Dmeet.Contract.Withdraw(&_Dmeet.TransactOpts)
}

// DmeetAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Dmeet contract.
type DmeetAdminChangedIterator struct {
	Event *DmeetAdminChanged // Event containing the contract specifics and raw log

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
func (it *DmeetAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DmeetAdminChanged)
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
		it.Event = new(DmeetAdminChanged)
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
func (it *DmeetAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DmeetAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DmeetAdminChanged represents a AdminChanged event raised by the Dmeet contract.
type DmeetAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Dmeet *DmeetFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*DmeetAdminChangedIterator, error) {

	logs, sub, err := _Dmeet.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &DmeetAdminChangedIterator{contract: _Dmeet.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Dmeet *DmeetFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *DmeetAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Dmeet.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DmeetAdminChanged)
				if err := _Dmeet.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Dmeet *DmeetFilterer) ParseAdminChanged(log types.Log) (*DmeetAdminChanged, error) {
	event := new(DmeetAdminChanged)
	if err := _Dmeet.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DmeetAdminChanged0Iterator is returned from FilterAdminChanged0 and is used to iterate over the raw logs and unpacked data for AdminChanged0 events raised by the Dmeet contract.
type DmeetAdminChanged0Iterator struct {
	Event *DmeetAdminChanged0 // Event containing the contract specifics and raw log

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
func (it *DmeetAdminChanged0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DmeetAdminChanged0)
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
		it.Event = new(DmeetAdminChanged0)
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
func (it *DmeetAdminChanged0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DmeetAdminChanged0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DmeetAdminChanged0 represents a AdminChanged0 event raised by the Dmeet contract.
type DmeetAdminChanged0 struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged0 is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Dmeet *DmeetFilterer) FilterAdminChanged0(opts *bind.FilterOpts) (*DmeetAdminChanged0Iterator, error) {

	logs, sub, err := _Dmeet.contract.FilterLogs(opts, "AdminChanged0")
	if err != nil {
		return nil, err
	}
	return &DmeetAdminChanged0Iterator{contract: _Dmeet.contract, event: "AdminChanged0", logs: logs, sub: sub}, nil
}

// WatchAdminChanged0 is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Dmeet *DmeetFilterer) WatchAdminChanged0(opts *bind.WatchOpts, sink chan<- *DmeetAdminChanged0) (event.Subscription, error) {

	logs, sub, err := _Dmeet.contract.WatchLogs(opts, "AdminChanged0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DmeetAdminChanged0)
				if err := _Dmeet.contract.UnpackLog(event, "AdminChanged0", log); err != nil {
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

// ParseAdminChanged0 is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Dmeet *DmeetFilterer) ParseAdminChanged0(log types.Log) (*DmeetAdminChanged0, error) {
	event := new(DmeetAdminChanged0)
	if err := _Dmeet.contract.UnpackLog(event, "AdminChanged0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DmeetAdminChanged1Iterator is returned from FilterAdminChanged1 and is used to iterate over the raw logs and unpacked data for AdminChanged1 events raised by the Dmeet contract.
type DmeetAdminChanged1Iterator struct {
	Event *DmeetAdminChanged1 // Event containing the contract specifics and raw log

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
func (it *DmeetAdminChanged1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DmeetAdminChanged1)
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
		it.Event = new(DmeetAdminChanged1)
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
func (it *DmeetAdminChanged1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DmeetAdminChanged1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DmeetAdminChanged1 represents a AdminChanged1 event raised by the Dmeet contract.
type DmeetAdminChanged1 struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged1 is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Dmeet *DmeetFilterer) FilterAdminChanged1(opts *bind.FilterOpts) (*DmeetAdminChanged1Iterator, error) {

	logs, sub, err := _Dmeet.contract.FilterLogs(opts, "AdminChanged1")
	if err != nil {
		return nil, err
	}
	return &DmeetAdminChanged1Iterator{contract: _Dmeet.contract, event: "AdminChanged1", logs: logs, sub: sub}, nil
}

// WatchAdminChanged1 is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Dmeet *DmeetFilterer) WatchAdminChanged1(opts *bind.WatchOpts, sink chan<- *DmeetAdminChanged1) (event.Subscription, error) {

	logs, sub, err := _Dmeet.contract.WatchLogs(opts, "AdminChanged1")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DmeetAdminChanged1)
				if err := _Dmeet.contract.UnpackLog(event, "AdminChanged1", log); err != nil {
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

// ParseAdminChanged1 is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Dmeet *DmeetFilterer) ParseAdminChanged1(log types.Log) (*DmeetAdminChanged1, error) {
	event := new(DmeetAdminChanged1)
	if err := _Dmeet.contract.UnpackLog(event, "AdminChanged1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DmeetBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Dmeet contract.
type DmeetBeaconUpgradedIterator struct {
	Event *DmeetBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *DmeetBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DmeetBeaconUpgraded)
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
		it.Event = new(DmeetBeaconUpgraded)
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
func (it *DmeetBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DmeetBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DmeetBeaconUpgraded represents a BeaconUpgraded event raised by the Dmeet contract.
type DmeetBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Dmeet *DmeetFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*DmeetBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Dmeet.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &DmeetBeaconUpgradedIterator{contract: _Dmeet.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Dmeet *DmeetFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *DmeetBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Dmeet.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DmeetBeaconUpgraded)
				if err := _Dmeet.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Dmeet *DmeetFilterer) ParseBeaconUpgraded(log types.Log) (*DmeetBeaconUpgraded, error) {
	event := new(DmeetBeaconUpgraded)
	if err := _Dmeet.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DmeetBeaconUpgraded0Iterator is returned from FilterBeaconUpgraded0 and is used to iterate over the raw logs and unpacked data for BeaconUpgraded0 events raised by the Dmeet contract.
type DmeetBeaconUpgraded0Iterator struct {
	Event *DmeetBeaconUpgraded0 // Event containing the contract specifics and raw log

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
func (it *DmeetBeaconUpgraded0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DmeetBeaconUpgraded0)
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
		it.Event = new(DmeetBeaconUpgraded0)
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
func (it *DmeetBeaconUpgraded0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DmeetBeaconUpgraded0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DmeetBeaconUpgraded0 represents a BeaconUpgraded0 event raised by the Dmeet contract.
type DmeetBeaconUpgraded0 struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded0 is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Dmeet *DmeetFilterer) FilterBeaconUpgraded0(opts *bind.FilterOpts, beacon []common.Address) (*DmeetBeaconUpgraded0Iterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Dmeet.contract.FilterLogs(opts, "BeaconUpgraded0", beaconRule)
	if err != nil {
		return nil, err
	}
	return &DmeetBeaconUpgraded0Iterator{contract: _Dmeet.contract, event: "BeaconUpgraded0", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded0 is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Dmeet *DmeetFilterer) WatchBeaconUpgraded0(opts *bind.WatchOpts, sink chan<- *DmeetBeaconUpgraded0, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Dmeet.contract.WatchLogs(opts, "BeaconUpgraded0", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DmeetBeaconUpgraded0)
				if err := _Dmeet.contract.UnpackLog(event, "BeaconUpgraded0", log); err != nil {
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

// ParseBeaconUpgraded0 is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Dmeet *DmeetFilterer) ParseBeaconUpgraded0(log types.Log) (*DmeetBeaconUpgraded0, error) {
	event := new(DmeetBeaconUpgraded0)
	if err := _Dmeet.contract.UnpackLog(event, "BeaconUpgraded0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DmeetBeaconUpgraded1Iterator is returned from FilterBeaconUpgraded1 and is used to iterate over the raw logs and unpacked data for BeaconUpgraded1 events raised by the Dmeet contract.
type DmeetBeaconUpgraded1Iterator struct {
	Event *DmeetBeaconUpgraded1 // Event containing the contract specifics and raw log

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
func (it *DmeetBeaconUpgraded1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DmeetBeaconUpgraded1)
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
		it.Event = new(DmeetBeaconUpgraded1)
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
func (it *DmeetBeaconUpgraded1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DmeetBeaconUpgraded1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DmeetBeaconUpgraded1 represents a BeaconUpgraded1 event raised by the Dmeet contract.
type DmeetBeaconUpgraded1 struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded1 is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Dmeet *DmeetFilterer) FilterBeaconUpgraded1(opts *bind.FilterOpts, beacon []common.Address) (*DmeetBeaconUpgraded1Iterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Dmeet.contract.FilterLogs(opts, "BeaconUpgraded1", beaconRule)
	if err != nil {
		return nil, err
	}
	return &DmeetBeaconUpgraded1Iterator{contract: _Dmeet.contract, event: "BeaconUpgraded1", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded1 is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Dmeet *DmeetFilterer) WatchBeaconUpgraded1(opts *bind.WatchOpts, sink chan<- *DmeetBeaconUpgraded1, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Dmeet.contract.WatchLogs(opts, "BeaconUpgraded1", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DmeetBeaconUpgraded1)
				if err := _Dmeet.contract.UnpackLog(event, "BeaconUpgraded1", log); err != nil {
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

// ParseBeaconUpgraded1 is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Dmeet *DmeetFilterer) ParseBeaconUpgraded1(log types.Log) (*DmeetBeaconUpgraded1, error) {
	event := new(DmeetBeaconUpgraded1)
	if err := _Dmeet.contract.UnpackLog(event, "BeaconUpgraded1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DmeetMembershipIterator is returned from FilterMembership and is used to iterate over the raw logs and unpacked data for Membership events raised by the Dmeet contract.
type DmeetMembershipIterator struct {
	Event *DmeetMembership // Event containing the contract specifics and raw log

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
func (it *DmeetMembershipIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DmeetMembership)
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
		it.Event = new(DmeetMembership)
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
func (it *DmeetMembershipIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DmeetMembershipIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DmeetMembership represents a Membership event raised by the Dmeet contract.
type DmeetMembership struct {
	Creator     common.Address
	Viewer      *big.Int
	Participant *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMembership is a free log retrieval operation binding the contract event 0x9752103f34b3dc3370aec0ba081b56d3149e6d7f6b750c28d340961ddd006570.
//
// Solidity: event Membership(address indexed creator, uint256 viewer, uint256 participant)
func (_Dmeet *DmeetFilterer) FilterMembership(opts *bind.FilterOpts, creator []common.Address) (*DmeetMembershipIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _Dmeet.contract.FilterLogs(opts, "Membership", creatorRule)
	if err != nil {
		return nil, err
	}
	return &DmeetMembershipIterator{contract: _Dmeet.contract, event: "Membership", logs: logs, sub: sub}, nil
}

// WatchMembership is a free log subscription operation binding the contract event 0x9752103f34b3dc3370aec0ba081b56d3149e6d7f6b750c28d340961ddd006570.
//
// Solidity: event Membership(address indexed creator, uint256 viewer, uint256 participant)
func (_Dmeet *DmeetFilterer) WatchMembership(opts *bind.WatchOpts, sink chan<- *DmeetMembership, creator []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _Dmeet.contract.WatchLogs(opts, "Membership", creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DmeetMembership)
				if err := _Dmeet.contract.UnpackLog(event, "Membership", log); err != nil {
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

// ParseMembership is a log parse operation binding the contract event 0x9752103f34b3dc3370aec0ba081b56d3149e6d7f6b750c28d340961ddd006570.
//
// Solidity: event Membership(address indexed creator, uint256 viewer, uint256 participant)
func (_Dmeet *DmeetFilterer) ParseMembership(log types.Log) (*DmeetMembership, error) {
	event := new(DmeetMembership)
	if err := _Dmeet.contract.UnpackLog(event, "Membership", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DmeetOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Dmeet contract.
type DmeetOwnershipTransferredIterator struct {
	Event *DmeetOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DmeetOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DmeetOwnershipTransferred)
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
		it.Event = new(DmeetOwnershipTransferred)
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
func (it *DmeetOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DmeetOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DmeetOwnershipTransferred represents a OwnershipTransferred event raised by the Dmeet contract.
type DmeetOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Dmeet *DmeetFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DmeetOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dmeet.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DmeetOwnershipTransferredIterator{contract: _Dmeet.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Dmeet *DmeetFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DmeetOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dmeet.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DmeetOwnershipTransferred)
				if err := _Dmeet.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Dmeet *DmeetFilterer) ParseOwnershipTransferred(log types.Log) (*DmeetOwnershipTransferred, error) {
	event := new(DmeetOwnershipTransferred)
	if err := _Dmeet.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DmeetOwnershipTransferred0Iterator is returned from FilterOwnershipTransferred0 and is used to iterate over the raw logs and unpacked data for OwnershipTransferred0 events raised by the Dmeet contract.
type DmeetOwnershipTransferred0Iterator struct {
	Event *DmeetOwnershipTransferred0 // Event containing the contract specifics and raw log

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
func (it *DmeetOwnershipTransferred0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DmeetOwnershipTransferred0)
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
		it.Event = new(DmeetOwnershipTransferred0)
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
func (it *DmeetOwnershipTransferred0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DmeetOwnershipTransferred0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DmeetOwnershipTransferred0 represents a OwnershipTransferred0 event raised by the Dmeet contract.
type DmeetOwnershipTransferred0 struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred0 is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Dmeet *DmeetFilterer) FilterOwnershipTransferred0(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DmeetOwnershipTransferred0Iterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dmeet.contract.FilterLogs(opts, "OwnershipTransferred0", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DmeetOwnershipTransferred0Iterator{contract: _Dmeet.contract, event: "OwnershipTransferred0", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred0 is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Dmeet *DmeetFilterer) WatchOwnershipTransferred0(opts *bind.WatchOpts, sink chan<- *DmeetOwnershipTransferred0, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dmeet.contract.WatchLogs(opts, "OwnershipTransferred0", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DmeetOwnershipTransferred0)
				if err := _Dmeet.contract.UnpackLog(event, "OwnershipTransferred0", log); err != nil {
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

// ParseOwnershipTransferred0 is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Dmeet *DmeetFilterer) ParseOwnershipTransferred0(log types.Log) (*DmeetOwnershipTransferred0, error) {
	event := new(DmeetOwnershipTransferred0)
	if err := _Dmeet.contract.UnpackLog(event, "OwnershipTransferred0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DmeetUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Dmeet contract.
type DmeetUpgradedIterator struct {
	Event *DmeetUpgraded // Event containing the contract specifics and raw log

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
func (it *DmeetUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DmeetUpgraded)
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
		it.Event = new(DmeetUpgraded)
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
func (it *DmeetUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DmeetUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DmeetUpgraded represents a Upgraded event raised by the Dmeet contract.
type DmeetUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Dmeet *DmeetFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*DmeetUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Dmeet.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &DmeetUpgradedIterator{contract: _Dmeet.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Dmeet *DmeetFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *DmeetUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Dmeet.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DmeetUpgraded)
				if err := _Dmeet.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Dmeet *DmeetFilterer) ParseUpgraded(log types.Log) (*DmeetUpgraded, error) {
	event := new(DmeetUpgraded)
	if err := _Dmeet.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DmeetUpgraded0Iterator is returned from FilterUpgraded0 and is used to iterate over the raw logs and unpacked data for Upgraded0 events raised by the Dmeet contract.
type DmeetUpgraded0Iterator struct {
	Event *DmeetUpgraded0 // Event containing the contract specifics and raw log

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
func (it *DmeetUpgraded0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DmeetUpgraded0)
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
		it.Event = new(DmeetUpgraded0)
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
func (it *DmeetUpgraded0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DmeetUpgraded0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DmeetUpgraded0 represents a Upgraded0 event raised by the Dmeet contract.
type DmeetUpgraded0 struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded0 is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Dmeet *DmeetFilterer) FilterUpgraded0(opts *bind.FilterOpts, implementation []common.Address) (*DmeetUpgraded0Iterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Dmeet.contract.FilterLogs(opts, "Upgraded0", implementationRule)
	if err != nil {
		return nil, err
	}
	return &DmeetUpgraded0Iterator{contract: _Dmeet.contract, event: "Upgraded0", logs: logs, sub: sub}, nil
}

// WatchUpgraded0 is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Dmeet *DmeetFilterer) WatchUpgraded0(opts *bind.WatchOpts, sink chan<- *DmeetUpgraded0, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Dmeet.contract.WatchLogs(opts, "Upgraded0", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DmeetUpgraded0)
				if err := _Dmeet.contract.UnpackLog(event, "Upgraded0", log); err != nil {
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

// ParseUpgraded0 is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Dmeet *DmeetFilterer) ParseUpgraded0(log types.Log) (*DmeetUpgraded0, error) {
	event := new(DmeetUpgraded0)
	if err := _Dmeet.contract.UnpackLog(event, "Upgraded0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DmeetUpgraded1Iterator is returned from FilterUpgraded1 and is used to iterate over the raw logs and unpacked data for Upgraded1 events raised by the Dmeet contract.
type DmeetUpgraded1Iterator struct {
	Event *DmeetUpgraded1 // Event containing the contract specifics and raw log

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
func (it *DmeetUpgraded1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DmeetUpgraded1)
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
		it.Event = new(DmeetUpgraded1)
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
func (it *DmeetUpgraded1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DmeetUpgraded1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DmeetUpgraded1 represents a Upgraded1 event raised by the Dmeet contract.
type DmeetUpgraded1 struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded1 is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Dmeet *DmeetFilterer) FilterUpgraded1(opts *bind.FilterOpts, implementation []common.Address) (*DmeetUpgraded1Iterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Dmeet.contract.FilterLogs(opts, "Upgraded1", implementationRule)
	if err != nil {
		return nil, err
	}
	return &DmeetUpgraded1Iterator{contract: _Dmeet.contract, event: "Upgraded1", logs: logs, sub: sub}, nil
}

// WatchUpgraded1 is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Dmeet *DmeetFilterer) WatchUpgraded1(opts *bind.WatchOpts, sink chan<- *DmeetUpgraded1, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Dmeet.contract.WatchLogs(opts, "Upgraded1", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DmeetUpgraded1)
				if err := _Dmeet.contract.UnpackLog(event, "Upgraded1", log); err != nil {
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

// ParseUpgraded1 is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Dmeet *DmeetFilterer) ParseUpgraded1(log types.Log) (*DmeetUpgraded1, error) {
	event := new(DmeetUpgraded1)
	if err := _Dmeet.contract.UnpackLog(event, "Upgraded1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
