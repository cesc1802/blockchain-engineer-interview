// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package blockchain

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
	_ = abi.ConvertType
)

// PostCovidStrokePreventionMetaData contains all meta data concerning the PostCovidStrokePrevention contract.
var PostCovidStrokePreventionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"riskScore\",\"type\":\"uint256\"}],\"name\":\"reward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PostCovidStrokePreventionABI is the input ABI used to generate the binding from.
// Deprecated: Use PostCovidStrokePreventionMetaData.ABI instead.
var PostCovidStrokePreventionABI = PostCovidStrokePreventionMetaData.ABI

// PostCovidStrokePrevention is an auto generated Go binding around an Ethereum contract.
type PostCovidStrokePrevention struct {
	PostCovidStrokePreventionCaller     // Read-only binding to the contract
	PostCovidStrokePreventionTransactor // Write-only binding to the contract
	PostCovidStrokePreventionFilterer   // Log filterer for contract events
}

// PostCovidStrokePreventionCaller is an auto generated read-only Go binding around an Ethereum contract.
type PostCovidStrokePreventionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PostCovidStrokePreventionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PostCovidStrokePreventionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PostCovidStrokePreventionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PostCovidStrokePreventionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PostCovidStrokePreventionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PostCovidStrokePreventionSession struct {
	Contract     *PostCovidStrokePrevention // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// PostCovidStrokePreventionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PostCovidStrokePreventionCallerSession struct {
	Contract *PostCovidStrokePreventionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// PostCovidStrokePreventionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PostCovidStrokePreventionTransactorSession struct {
	Contract     *PostCovidStrokePreventionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// PostCovidStrokePreventionRaw is an auto generated low-level Go binding around an Ethereum contract.
type PostCovidStrokePreventionRaw struct {
	Contract *PostCovidStrokePrevention // Generic contract binding to access the raw methods on
}

// PostCovidStrokePreventionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PostCovidStrokePreventionCallerRaw struct {
	Contract *PostCovidStrokePreventionCaller // Generic read-only contract binding to access the raw methods on
}

// PostCovidStrokePreventionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PostCovidStrokePreventionTransactorRaw struct {
	Contract *PostCovidStrokePreventionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPostCovidStrokePrevention creates a new instance of PostCovidStrokePrevention, bound to a specific deployed contract.
func NewPostCovidStrokePrevention(address common.Address, backend bind.ContractBackend) (*PostCovidStrokePrevention, error) {
	contract, err := bindPostCovidStrokePrevention(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PostCovidStrokePrevention{PostCovidStrokePreventionCaller: PostCovidStrokePreventionCaller{contract: contract}, PostCovidStrokePreventionTransactor: PostCovidStrokePreventionTransactor{contract: contract}, PostCovidStrokePreventionFilterer: PostCovidStrokePreventionFilterer{contract: contract}}, nil
}

// NewPostCovidStrokePreventionCaller creates a new read-only instance of PostCovidStrokePrevention, bound to a specific deployed contract.
func NewPostCovidStrokePreventionCaller(address common.Address, caller bind.ContractCaller) (*PostCovidStrokePreventionCaller, error) {
	contract, err := bindPostCovidStrokePrevention(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PostCovidStrokePreventionCaller{contract: contract}, nil
}

// NewPostCovidStrokePreventionTransactor creates a new write-only instance of PostCovidStrokePrevention, bound to a specific deployed contract.
func NewPostCovidStrokePreventionTransactor(address common.Address, transactor bind.ContractTransactor) (*PostCovidStrokePreventionTransactor, error) {
	contract, err := bindPostCovidStrokePrevention(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PostCovidStrokePreventionTransactor{contract: contract}, nil
}

// NewPostCovidStrokePreventionFilterer creates a new log filterer instance of PostCovidStrokePrevention, bound to a specific deployed contract.
func NewPostCovidStrokePreventionFilterer(address common.Address, filterer bind.ContractFilterer) (*PostCovidStrokePreventionFilterer, error) {
	contract, err := bindPostCovidStrokePrevention(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PostCovidStrokePreventionFilterer{contract: contract}, nil
}

// bindPostCovidStrokePrevention binds a generic wrapper to an already deployed contract.
func bindPostCovidStrokePrevention(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PostCovidStrokePreventionMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PostCovidStrokePrevention *PostCovidStrokePreventionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PostCovidStrokePrevention.Contract.PostCovidStrokePreventionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PostCovidStrokePrevention *PostCovidStrokePreventionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PostCovidStrokePrevention.Contract.PostCovidStrokePreventionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PostCovidStrokePrevention *PostCovidStrokePreventionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PostCovidStrokePrevention.Contract.PostCovidStrokePreventionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PostCovidStrokePrevention *PostCovidStrokePreventionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PostCovidStrokePrevention.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PostCovidStrokePrevention *PostCovidStrokePreventionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PostCovidStrokePrevention.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PostCovidStrokePrevention *PostCovidStrokePreventionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PostCovidStrokePrevention.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_PostCovidStrokePrevention *PostCovidStrokePreventionCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PostCovidStrokePrevention.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_PostCovidStrokePrevention *PostCovidStrokePreventionSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _PostCovidStrokePrevention.Contract.Allowance(&_PostCovidStrokePrevention.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_PostCovidStrokePrevention *PostCovidStrokePreventionCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _PostCovidStrokePrevention.Contract.Allowance(&_PostCovidStrokePrevention.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_PostCovidStrokePrevention *PostCovidStrokePreventionCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PostCovidStrokePrevention.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_PostCovidStrokePrevention *PostCovidStrokePreventionSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _PostCovidStrokePrevention.Contract.BalanceOf(&_PostCovidStrokePrevention.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_PostCovidStrokePrevention *PostCovidStrokePreventionCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _PostCovidStrokePrevention.Contract.BalanceOf(&_PostCovidStrokePrevention.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_PostCovidStrokePrevention *PostCovidStrokePreventionCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _PostCovidStrokePrevention.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_PostCovidStrokePrevention *PostCovidStrokePreventionSession) Decimals() (uint8, error) {
	return _PostCovidStrokePrevention.Contract.Decimals(&_PostCovidStrokePrevention.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_PostCovidStrokePrevention *PostCovidStrokePreventionCallerSession) Decimals() (uint8, error) {
	return _PostCovidStrokePrevention.Contract.Decimals(&_PostCovidStrokePrevention.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PostCovidStrokePrevention *PostCovidStrokePreventionCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PostCovidStrokePrevention.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PostCovidStrokePrevention *PostCovidStrokePreventionSession) Name() (string, error) {
	return _PostCovidStrokePrevention.Contract.Name(&_PostCovidStrokePrevention.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PostCovidStrokePrevention *PostCovidStrokePreventionCallerSession) Name() (string, error) {
	return _PostCovidStrokePrevention.Contract.Name(&_PostCovidStrokePrevention.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PostCovidStrokePrevention *PostCovidStrokePreventionCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PostCovidStrokePrevention.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PostCovidStrokePrevention *PostCovidStrokePreventionSession) Owner() (common.Address, error) {
	return _PostCovidStrokePrevention.Contract.Owner(&_PostCovidStrokePrevention.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PostCovidStrokePrevention *PostCovidStrokePreventionCallerSession) Owner() (common.Address, error) {
	return _PostCovidStrokePrevention.Contract.Owner(&_PostCovidStrokePrevention.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PostCovidStrokePrevention *PostCovidStrokePreventionCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PostCovidStrokePrevention.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PostCovidStrokePrevention *PostCovidStrokePreventionSession) Symbol() (string, error) {
	return _PostCovidStrokePrevention.Contract.Symbol(&_PostCovidStrokePrevention.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PostCovidStrokePrevention *PostCovidStrokePreventionCallerSession) Symbol() (string, error) {
	return _PostCovidStrokePrevention.Contract.Symbol(&_PostCovidStrokePrevention.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PostCovidStrokePrevention *PostCovidStrokePreventionCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PostCovidStrokePrevention.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PostCovidStrokePrevention *PostCovidStrokePreventionSession) TotalSupply() (*big.Int, error) {
	return _PostCovidStrokePrevention.Contract.TotalSupply(&_PostCovidStrokePrevention.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PostCovidStrokePrevention *PostCovidStrokePreventionCallerSession) TotalSupply() (*big.Int, error) {
	return _PostCovidStrokePrevention.Contract.TotalSupply(&_PostCovidStrokePrevention.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_PostCovidStrokePrevention *PostCovidStrokePreventionTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PostCovidStrokePrevention.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_PostCovidStrokePrevention *PostCovidStrokePreventionSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PostCovidStrokePrevention.Contract.Approve(&_PostCovidStrokePrevention.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_PostCovidStrokePrevention *PostCovidStrokePreventionTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PostCovidStrokePrevention.Contract.Approve(&_PostCovidStrokePrevention.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_PostCovidStrokePrevention *PostCovidStrokePreventionTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _PostCovidStrokePrevention.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_PostCovidStrokePrevention *PostCovidStrokePreventionSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _PostCovidStrokePrevention.Contract.Burn(&_PostCovidStrokePrevention.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_PostCovidStrokePrevention *PostCovidStrokePreventionTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _PostCovidStrokePrevention.Contract.Burn(&_PostCovidStrokePrevention.TransactOpts, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_PostCovidStrokePrevention *PostCovidStrokePreventionTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PostCovidStrokePrevention.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_PostCovidStrokePrevention *PostCovidStrokePreventionSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PostCovidStrokePrevention.Contract.BurnFrom(&_PostCovidStrokePrevention.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_PostCovidStrokePrevention *PostCovidStrokePreventionTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PostCovidStrokePrevention.Contract.BurnFrom(&_PostCovidStrokePrevention.TransactOpts, account, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_PostCovidStrokePrevention *PostCovidStrokePreventionTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _PostCovidStrokePrevention.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_PostCovidStrokePrevention *PostCovidStrokePreventionSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _PostCovidStrokePrevention.Contract.DecreaseAllowance(&_PostCovidStrokePrevention.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_PostCovidStrokePrevention *PostCovidStrokePreventionTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _PostCovidStrokePrevention.Contract.DecreaseAllowance(&_PostCovidStrokePrevention.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_PostCovidStrokePrevention *PostCovidStrokePreventionTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _PostCovidStrokePrevention.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_PostCovidStrokePrevention *PostCovidStrokePreventionSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _PostCovidStrokePrevention.Contract.IncreaseAllowance(&_PostCovidStrokePrevention.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_PostCovidStrokePrevention *PostCovidStrokePreventionTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _PostCovidStrokePrevention.Contract.IncreaseAllowance(&_PostCovidStrokePrevention.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_PostCovidStrokePrevention *PostCovidStrokePreventionTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PostCovidStrokePrevention.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_PostCovidStrokePrevention *PostCovidStrokePreventionSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PostCovidStrokePrevention.Contract.Mint(&_PostCovidStrokePrevention.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_PostCovidStrokePrevention *PostCovidStrokePreventionTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PostCovidStrokePrevention.Contract.Mint(&_PostCovidStrokePrevention.TransactOpts, to, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PostCovidStrokePrevention *PostCovidStrokePreventionTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PostCovidStrokePrevention.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PostCovidStrokePrevention *PostCovidStrokePreventionSession) RenounceOwnership() (*types.Transaction, error) {
	return _PostCovidStrokePrevention.Contract.RenounceOwnership(&_PostCovidStrokePrevention.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PostCovidStrokePrevention *PostCovidStrokePreventionTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PostCovidStrokePrevention.Contract.RenounceOwnership(&_PostCovidStrokePrevention.TransactOpts)
}

// Reward is a paid mutator transaction binding the contract method 0x21670f22.
//
// Solidity: function reward(address to, uint256 riskScore) returns()
func (_PostCovidStrokePrevention *PostCovidStrokePreventionTransactor) Reward(opts *bind.TransactOpts, to common.Address, riskScore *big.Int) (*types.Transaction, error) {
	return _PostCovidStrokePrevention.contract.Transact(opts, "reward", to, riskScore)
}

// Reward is a paid mutator transaction binding the contract method 0x21670f22.
//
// Solidity: function reward(address to, uint256 riskScore) returns()
func (_PostCovidStrokePrevention *PostCovidStrokePreventionSession) Reward(to common.Address, riskScore *big.Int) (*types.Transaction, error) {
	return _PostCovidStrokePrevention.Contract.Reward(&_PostCovidStrokePrevention.TransactOpts, to, riskScore)
}

// Reward is a paid mutator transaction binding the contract method 0x21670f22.
//
// Solidity: function reward(address to, uint256 riskScore) returns()
func (_PostCovidStrokePrevention *PostCovidStrokePreventionTransactorSession) Reward(to common.Address, riskScore *big.Int) (*types.Transaction, error) {
	return _PostCovidStrokePrevention.Contract.Reward(&_PostCovidStrokePrevention.TransactOpts, to, riskScore)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_PostCovidStrokePrevention *PostCovidStrokePreventionTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PostCovidStrokePrevention.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_PostCovidStrokePrevention *PostCovidStrokePreventionSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PostCovidStrokePrevention.Contract.Transfer(&_PostCovidStrokePrevention.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_PostCovidStrokePrevention *PostCovidStrokePreventionTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PostCovidStrokePrevention.Contract.Transfer(&_PostCovidStrokePrevention.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_PostCovidStrokePrevention *PostCovidStrokePreventionTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PostCovidStrokePrevention.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_PostCovidStrokePrevention *PostCovidStrokePreventionSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PostCovidStrokePrevention.Contract.TransferFrom(&_PostCovidStrokePrevention.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_PostCovidStrokePrevention *PostCovidStrokePreventionTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PostCovidStrokePrevention.Contract.TransferFrom(&_PostCovidStrokePrevention.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PostCovidStrokePrevention *PostCovidStrokePreventionTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PostCovidStrokePrevention.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PostCovidStrokePrevention *PostCovidStrokePreventionSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PostCovidStrokePrevention.Contract.TransferOwnership(&_PostCovidStrokePrevention.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PostCovidStrokePrevention *PostCovidStrokePreventionTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PostCovidStrokePrevention.Contract.TransferOwnership(&_PostCovidStrokePrevention.TransactOpts, newOwner)
}

// PostCovidStrokePreventionApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the PostCovidStrokePrevention contract.
type PostCovidStrokePreventionApprovalIterator struct {
	Event *PostCovidStrokePreventionApproval // Event containing the contract specifics and raw log

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
func (it *PostCovidStrokePreventionApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PostCovidStrokePreventionApproval)
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
		it.Event = new(PostCovidStrokePreventionApproval)
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
func (it *PostCovidStrokePreventionApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PostCovidStrokePreventionApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PostCovidStrokePreventionApproval represents a Approval event raised by the PostCovidStrokePrevention contract.
type PostCovidStrokePreventionApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_PostCovidStrokePrevention *PostCovidStrokePreventionFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*PostCovidStrokePreventionApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _PostCovidStrokePrevention.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &PostCovidStrokePreventionApprovalIterator{contract: _PostCovidStrokePrevention.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_PostCovidStrokePrevention *PostCovidStrokePreventionFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *PostCovidStrokePreventionApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _PostCovidStrokePrevention.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PostCovidStrokePreventionApproval)
				if err := _PostCovidStrokePrevention.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_PostCovidStrokePrevention *PostCovidStrokePreventionFilterer) ParseApproval(log types.Log) (*PostCovidStrokePreventionApproval, error) {
	event := new(PostCovidStrokePreventionApproval)
	if err := _PostCovidStrokePrevention.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PostCovidStrokePreventionOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PostCovidStrokePrevention contract.
type PostCovidStrokePreventionOwnershipTransferredIterator struct {
	Event *PostCovidStrokePreventionOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PostCovidStrokePreventionOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PostCovidStrokePreventionOwnershipTransferred)
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
		it.Event = new(PostCovidStrokePreventionOwnershipTransferred)
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
func (it *PostCovidStrokePreventionOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PostCovidStrokePreventionOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PostCovidStrokePreventionOwnershipTransferred represents a OwnershipTransferred event raised by the PostCovidStrokePrevention contract.
type PostCovidStrokePreventionOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PostCovidStrokePrevention *PostCovidStrokePreventionFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PostCovidStrokePreventionOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PostCovidStrokePrevention.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PostCovidStrokePreventionOwnershipTransferredIterator{contract: _PostCovidStrokePrevention.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PostCovidStrokePrevention *PostCovidStrokePreventionFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PostCovidStrokePreventionOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PostCovidStrokePrevention.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PostCovidStrokePreventionOwnershipTransferred)
				if err := _PostCovidStrokePrevention.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_PostCovidStrokePrevention *PostCovidStrokePreventionFilterer) ParseOwnershipTransferred(log types.Log) (*PostCovidStrokePreventionOwnershipTransferred, error) {
	event := new(PostCovidStrokePreventionOwnershipTransferred)
	if err := _PostCovidStrokePrevention.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PostCovidStrokePreventionTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the PostCovidStrokePrevention contract.
type PostCovidStrokePreventionTransferIterator struct {
	Event *PostCovidStrokePreventionTransfer // Event containing the contract specifics and raw log

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
func (it *PostCovidStrokePreventionTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PostCovidStrokePreventionTransfer)
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
		it.Event = new(PostCovidStrokePreventionTransfer)
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
func (it *PostCovidStrokePreventionTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PostCovidStrokePreventionTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PostCovidStrokePreventionTransfer represents a Transfer event raised by the PostCovidStrokePrevention contract.
type PostCovidStrokePreventionTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_PostCovidStrokePrevention *PostCovidStrokePreventionFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*PostCovidStrokePreventionTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PostCovidStrokePrevention.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PostCovidStrokePreventionTransferIterator{contract: _PostCovidStrokePrevention.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_PostCovidStrokePrevention *PostCovidStrokePreventionFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PostCovidStrokePreventionTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PostCovidStrokePrevention.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PostCovidStrokePreventionTransfer)
				if err := _PostCovidStrokePrevention.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_PostCovidStrokePrevention *PostCovidStrokePreventionFilterer) ParseTransfer(log types.Log) (*PostCovidStrokePreventionTransfer, error) {
	event := new(PostCovidStrokePreventionTransfer)
	if err := _PostCovidStrokePrevention.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
