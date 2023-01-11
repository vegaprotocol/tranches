// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// TokenV2MetaData contains all meta data concerning the TokenV2 contract.
var TokenV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"total_supply_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mint_lock_expiry_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"new_controller\",\"type\":\"address\"}],\"name\":\"Controller_Changed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"new_controller\",\"type\":\"address\"}],\"name\":\"change_controller\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"controller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint_and_issue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mint_lock_expiry\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TokenV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenV2MetaData.ABI instead.
var TokenV2ABI = TokenV2MetaData.ABI

// TokenV2 is an auto generated Go binding around an Ethereum contract.
type TokenV2 struct {
	TokenV2Caller     // Read-only binding to the contract
	TokenV2Transactor // Write-only binding to the contract
	TokenV2Filterer   // Log filterer for contract events
}

// TokenV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type TokenV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenV2Session struct {
	Contract     *TokenV2          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenV2CallerSession struct {
	Contract *TokenV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// TokenV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenV2TransactorSession struct {
	Contract     *TokenV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// TokenV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type TokenV2Raw struct {
	Contract *TokenV2 // Generic contract binding to access the raw methods on
}

// TokenV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenV2CallerRaw struct {
	Contract *TokenV2Caller // Generic read-only contract binding to access the raw methods on
}

// TokenV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenV2TransactorRaw struct {
	Contract *TokenV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenV2 creates a new instance of TokenV2, bound to a specific deployed contract.
func NewTokenV2(address common.Address, backend bind.ContractBackend) (*TokenV2, error) {
	contract, err := bindTokenV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenV2{TokenV2Caller: TokenV2Caller{contract: contract}, TokenV2Transactor: TokenV2Transactor{contract: contract}, TokenV2Filterer: TokenV2Filterer{contract: contract}}, nil
}

// NewTokenV2Caller creates a new read-only instance of TokenV2, bound to a specific deployed contract.
func NewTokenV2Caller(address common.Address, caller bind.ContractCaller) (*TokenV2Caller, error) {
	contract, err := bindTokenV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenV2Caller{contract: contract}, nil
}

// NewTokenV2Transactor creates a new write-only instance of TokenV2, bound to a specific deployed contract.
func NewTokenV2Transactor(address common.Address, transactor bind.ContractTransactor) (*TokenV2Transactor, error) {
	contract, err := bindTokenV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenV2Transactor{contract: contract}, nil
}

// NewTokenV2Filterer creates a new log filterer instance of TokenV2, bound to a specific deployed contract.
func NewTokenV2Filterer(address common.Address, filterer bind.ContractFilterer) (*TokenV2Filterer, error) {
	contract, err := bindTokenV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenV2Filterer{contract: contract}, nil
}

// bindTokenV2 binds a generic wrapper to an already deployed contract.
func bindTokenV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenV2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenV2 *TokenV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenV2.Contract.TokenV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenV2 *TokenV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenV2.Contract.TokenV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenV2 *TokenV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenV2.Contract.TokenV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenV2 *TokenV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenV2 *TokenV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenV2 *TokenV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenV2.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_TokenV2 *TokenV2Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TokenV2.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_TokenV2 *TokenV2Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _TokenV2.Contract.Allowance(&_TokenV2.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_TokenV2 *TokenV2CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _TokenV2.Contract.Allowance(&_TokenV2.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_TokenV2 *TokenV2Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TokenV2.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_TokenV2 *TokenV2Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _TokenV2.Contract.BalanceOf(&_TokenV2.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_TokenV2 *TokenV2CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _TokenV2.Contract.BalanceOf(&_TokenV2.CallOpts, account)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_TokenV2 *TokenV2Caller) Controller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenV2.contract.Call(opts, &out, "controller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_TokenV2 *TokenV2Session) Controller() (common.Address, error) {
	return _TokenV2.Contract.Controller(&_TokenV2.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_TokenV2 *TokenV2CallerSession) Controller() (common.Address, error) {
	return _TokenV2.Contract.Controller(&_TokenV2.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TokenV2 *TokenV2Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _TokenV2.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TokenV2 *TokenV2Session) Decimals() (uint8, error) {
	return _TokenV2.Contract.Decimals(&_TokenV2.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TokenV2 *TokenV2CallerSession) Decimals() (uint8, error) {
	return _TokenV2.Contract.Decimals(&_TokenV2.CallOpts)
}

// MintLockExpiry is a free data retrieval call binding the contract method 0x66304afd.
//
// Solidity: function mint_lock_expiry() view returns(uint256)
func (_TokenV2 *TokenV2Caller) MintLockExpiry(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenV2.contract.Call(opts, &out, "mint_lock_expiry")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintLockExpiry is a free data retrieval call binding the contract method 0x66304afd.
//
// Solidity: function mint_lock_expiry() view returns(uint256)
func (_TokenV2 *TokenV2Session) MintLockExpiry() (*big.Int, error) {
	return _TokenV2.Contract.MintLockExpiry(&_TokenV2.CallOpts)
}

// MintLockExpiry is a free data retrieval call binding the contract method 0x66304afd.
//
// Solidity: function mint_lock_expiry() view returns(uint256)
func (_TokenV2 *TokenV2CallerSession) MintLockExpiry() (*big.Int, error) {
	return _TokenV2.Contract.MintLockExpiry(&_TokenV2.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TokenV2 *TokenV2Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TokenV2.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TokenV2 *TokenV2Session) Name() (string, error) {
	return _TokenV2.Contract.Name(&_TokenV2.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TokenV2 *TokenV2CallerSession) Name() (string, error) {
	return _TokenV2.Contract.Name(&_TokenV2.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TokenV2 *TokenV2Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TokenV2.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TokenV2 *TokenV2Session) Symbol() (string, error) {
	return _TokenV2.Contract.Symbol(&_TokenV2.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TokenV2 *TokenV2CallerSession) Symbol() (string, error) {
	return _TokenV2.Contract.Symbol(&_TokenV2.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TokenV2 *TokenV2Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenV2.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TokenV2 *TokenV2Session) TotalSupply() (*big.Int, error) {
	return _TokenV2.Contract.TotalSupply(&_TokenV2.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TokenV2 *TokenV2CallerSession) TotalSupply() (*big.Int, error) {
	return _TokenV2.Contract.TotalSupply(&_TokenV2.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_TokenV2 *TokenV2Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenV2.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_TokenV2 *TokenV2Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenV2.Contract.Approve(&_TokenV2.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_TokenV2 *TokenV2TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenV2.Contract.Approve(&_TokenV2.TransactOpts, spender, amount)
}

// ChangeController is a paid mutator transaction binding the contract method 0x711d03ce.
//
// Solidity: function change_controller(address new_controller) returns()
func (_TokenV2 *TokenV2Transactor) ChangeController(opts *bind.TransactOpts, new_controller common.Address) (*types.Transaction, error) {
	return _TokenV2.contract.Transact(opts, "change_controller", new_controller)
}

// ChangeController is a paid mutator transaction binding the contract method 0x711d03ce.
//
// Solidity: function change_controller(address new_controller) returns()
func (_TokenV2 *TokenV2Session) ChangeController(new_controller common.Address) (*types.Transaction, error) {
	return _TokenV2.Contract.ChangeController(&_TokenV2.TransactOpts, new_controller)
}

// ChangeController is a paid mutator transaction binding the contract method 0x711d03ce.
//
// Solidity: function change_controller(address new_controller) returns()
func (_TokenV2 *TokenV2TransactorSession) ChangeController(new_controller common.Address) (*types.Transaction, error) {
	return _TokenV2.Contract.ChangeController(&_TokenV2.TransactOpts, new_controller)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_TokenV2 *TokenV2Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _TokenV2.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_TokenV2 *TokenV2Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _TokenV2.Contract.DecreaseAllowance(&_TokenV2.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_TokenV2 *TokenV2TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _TokenV2.Contract.DecreaseAllowance(&_TokenV2.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_TokenV2 *TokenV2Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _TokenV2.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_TokenV2 *TokenV2Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _TokenV2.Contract.IncreaseAllowance(&_TokenV2.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_TokenV2 *TokenV2TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _TokenV2.Contract.IncreaseAllowance(&_TokenV2.TransactOpts, spender, addedValue)
}

// MintAndIssue is a paid mutator transaction binding the contract method 0x28231bf2.
//
// Solidity: function mint_and_issue(address target, uint256 amount) returns()
func (_TokenV2 *TokenV2Transactor) MintAndIssue(opts *bind.TransactOpts, target common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenV2.contract.Transact(opts, "mint_and_issue", target, amount)
}

// MintAndIssue is a paid mutator transaction binding the contract method 0x28231bf2.
//
// Solidity: function mint_and_issue(address target, uint256 amount) returns()
func (_TokenV2 *TokenV2Session) MintAndIssue(target common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenV2.Contract.MintAndIssue(&_TokenV2.TransactOpts, target, amount)
}

// MintAndIssue is a paid mutator transaction binding the contract method 0x28231bf2.
//
// Solidity: function mint_and_issue(address target, uint256 amount) returns()
func (_TokenV2 *TokenV2TransactorSession) MintAndIssue(target common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenV2.Contract.MintAndIssue(&_TokenV2.TransactOpts, target, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_TokenV2 *TokenV2Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenV2.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_TokenV2 *TokenV2Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenV2.Contract.Transfer(&_TokenV2.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_TokenV2 *TokenV2TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenV2.Contract.Transfer(&_TokenV2.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_TokenV2 *TokenV2Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenV2.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_TokenV2 *TokenV2Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenV2.Contract.TransferFrom(&_TokenV2.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_TokenV2 *TokenV2TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenV2.Contract.TransferFrom(&_TokenV2.TransactOpts, sender, recipient, amount)
}

// TokenV2ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the TokenV2 contract.
type TokenV2ApprovalIterator struct {
	Event *TokenV2Approval // Event containing the contract specifics and raw log

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
func (it *TokenV2ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenV2Approval)
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
		it.Event = new(TokenV2Approval)
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
func (it *TokenV2ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenV2ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenV2Approval represents a Approval event raised by the TokenV2 contract.
type TokenV2Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TokenV2 *TokenV2Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*TokenV2ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _TokenV2.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &TokenV2ApprovalIterator{contract: _TokenV2.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TokenV2 *TokenV2Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TokenV2Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _TokenV2.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenV2Approval)
				if err := _TokenV2.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_TokenV2 *TokenV2Filterer) ParseApproval(log types.Log) (*TokenV2Approval, error) {
	event := new(TokenV2Approval)
	if err := _TokenV2.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenV2ControllerChangedIterator is returned from FilterControllerChanged and is used to iterate over the raw logs and unpacked data for ControllerChanged events raised by the TokenV2 contract.
type TokenV2ControllerChangedIterator struct {
	Event *TokenV2ControllerChanged // Event containing the contract specifics and raw log

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
func (it *TokenV2ControllerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenV2ControllerChanged)
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
		it.Event = new(TokenV2ControllerChanged)
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
func (it *TokenV2ControllerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenV2ControllerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenV2ControllerChanged represents a ControllerChanged event raised by the TokenV2 contract.
type TokenV2ControllerChanged struct {
	NewController common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterControllerChanged is a free log retrieval operation binding the contract event 0xdacf3394e265c5a9b5a0b8b7f9c64619dd468f6abe1fc900bd429e6a8700bed9.
//
// Solidity: event Controller_Changed(address indexed new_controller)
func (_TokenV2 *TokenV2Filterer) FilterControllerChanged(opts *bind.FilterOpts, new_controller []common.Address) (*TokenV2ControllerChangedIterator, error) {

	var new_controllerRule []interface{}
	for _, new_controllerItem := range new_controller {
		new_controllerRule = append(new_controllerRule, new_controllerItem)
	}

	logs, sub, err := _TokenV2.contract.FilterLogs(opts, "Controller_Changed", new_controllerRule)
	if err != nil {
		return nil, err
	}
	return &TokenV2ControllerChangedIterator{contract: _TokenV2.contract, event: "Controller_Changed", logs: logs, sub: sub}, nil
}

// WatchControllerChanged is a free log subscription operation binding the contract event 0xdacf3394e265c5a9b5a0b8b7f9c64619dd468f6abe1fc900bd429e6a8700bed9.
//
// Solidity: event Controller_Changed(address indexed new_controller)
func (_TokenV2 *TokenV2Filterer) WatchControllerChanged(opts *bind.WatchOpts, sink chan<- *TokenV2ControllerChanged, new_controller []common.Address) (event.Subscription, error) {

	var new_controllerRule []interface{}
	for _, new_controllerItem := range new_controller {
		new_controllerRule = append(new_controllerRule, new_controllerItem)
	}

	logs, sub, err := _TokenV2.contract.WatchLogs(opts, "Controller_Changed", new_controllerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenV2ControllerChanged)
				if err := _TokenV2.contract.UnpackLog(event, "Controller_Changed", log); err != nil {
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

// ParseControllerChanged is a log parse operation binding the contract event 0xdacf3394e265c5a9b5a0b8b7f9c64619dd468f6abe1fc900bd429e6a8700bed9.
//
// Solidity: event Controller_Changed(address indexed new_controller)
func (_TokenV2 *TokenV2Filterer) ParseControllerChanged(log types.Log) (*TokenV2ControllerChanged, error) {
	event := new(TokenV2ControllerChanged)
	if err := _TokenV2.contract.UnpackLog(event, "Controller_Changed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenV2TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the TokenV2 contract.
type TokenV2TransferIterator struct {
	Event *TokenV2Transfer // Event containing the contract specifics and raw log

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
func (it *TokenV2TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenV2Transfer)
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
		it.Event = new(TokenV2Transfer)
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
func (it *TokenV2TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenV2TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenV2Transfer represents a Transfer event raised by the TokenV2 contract.
type TokenV2Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TokenV2 *TokenV2Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TokenV2TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TokenV2.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TokenV2TransferIterator{contract: _TokenV2.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TokenV2 *TokenV2Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TokenV2Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TokenV2.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenV2Transfer)
				if err := _TokenV2.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_TokenV2 *TokenV2Filterer) ParseTransfer(log types.Log) (*TokenV2Transfer, error) {
	event := new(TokenV2Transfer)
	if err := _TokenV2.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
