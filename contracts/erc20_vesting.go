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

// ERC20VestingMetaData contains all meta data concerning the ERC20Vesting contract.
var ERC20VestingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_v1_address\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token_v2_address\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"old_addresses\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"new_addresses\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"new_controller\",\"type\":\"address\"}],\"name\":\"Controller_Set\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Issuer_Permitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"}],\"name\":\"Issuer_Revoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"vega_public_key\",\"type\":\"bytes32\"}],\"name\":\"Stake_Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"vega_public_key\",\"type\":\"bytes32\"}],\"name\":\"Stake_Removed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"vega_public_key\",\"type\":\"bytes32\"}],\"name\":\"Stake_Transferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"tranche_id\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Tranche_Balance_Added\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"tranche_id\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Tranche_Balance_Removed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"tranche_id\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cliff_start\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"Tranche_Created\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"accuracy_scale\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"address_migration\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"tranche_id\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"assisted_withdraw_from_tranche\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"controller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cliff_start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"create_tranche\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"default_tranche_id\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"tranche_id\",\"type\":\"uint8\"}],\"name\":\"get_tranche_balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"tranche_id\",\"type\":\"uint8\"}],\"name\":\"get_vested_for_tranche\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"tranche_id\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"issue_into_tranche\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"tranche_id\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"move_into_tranche\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"permit_issuer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"permitted_issuance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"vega_public_key\",\"type\":\"bytes32\"}],\"name\":\"remove_stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"}],\"name\":\"revoke_issuer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"new_controller\",\"type\":\"address\"}],\"name\":\"set_controller\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"vega_public_key\",\"type\":\"bytes32\"}],\"name\":\"stake_balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"vega_public_key\",\"type\":\"bytes32\"}],\"name\":\"stake_tokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"staking_token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"total_locked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"total_staked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tranche_count\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"tranches\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"cliff_start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"user_stats\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"total_in_all_tranches\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lien\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"user_total_all_tranches\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"v1_address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"v1_migrated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"v2_address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"tranche_id\",\"type\":\"uint8\"}],\"name\":\"withdraw_from_tranche\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ERC20VestingABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20VestingMetaData.ABI instead.
var ERC20VestingABI = ERC20VestingMetaData.ABI

// ERC20Vesting is an auto generated Go binding around an Ethereum contract.
type ERC20Vesting struct {
	ERC20VestingCaller     // Read-only binding to the contract
	ERC20VestingTransactor // Write-only binding to the contract
	ERC20VestingFilterer   // Log filterer for contract events
}

// ERC20VestingCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20VestingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20VestingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20VestingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20VestingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20VestingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20VestingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20VestingSession struct {
	Contract     *ERC20Vesting     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20VestingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20VestingCallerSession struct {
	Contract *ERC20VestingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ERC20VestingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20VestingTransactorSession struct {
	Contract     *ERC20VestingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ERC20VestingRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20VestingRaw struct {
	Contract *ERC20Vesting // Generic contract binding to access the raw methods on
}

// ERC20VestingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20VestingCallerRaw struct {
	Contract *ERC20VestingCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20VestingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20VestingTransactorRaw struct {
	Contract *ERC20VestingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Vesting creates a new instance of ERC20Vesting, bound to a specific deployed contract.
func NewERC20Vesting(address common.Address, backend bind.ContractBackend) (*ERC20Vesting, error) {
	contract, err := bindERC20Vesting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Vesting{ERC20VestingCaller: ERC20VestingCaller{contract: contract}, ERC20VestingTransactor: ERC20VestingTransactor{contract: contract}, ERC20VestingFilterer: ERC20VestingFilterer{contract: contract}}, nil
}

// NewERC20VestingCaller creates a new read-only instance of ERC20Vesting, bound to a specific deployed contract.
func NewERC20VestingCaller(address common.Address, caller bind.ContractCaller) (*ERC20VestingCaller, error) {
	contract, err := bindERC20Vesting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20VestingCaller{contract: contract}, nil
}

// NewERC20VestingTransactor creates a new write-only instance of ERC20Vesting, bound to a specific deployed contract.
func NewERC20VestingTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20VestingTransactor, error) {
	contract, err := bindERC20Vesting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20VestingTransactor{contract: contract}, nil
}

// NewERC20VestingFilterer creates a new log filterer instance of ERC20Vesting, bound to a specific deployed contract.
func NewERC20VestingFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20VestingFilterer, error) {
	contract, err := bindERC20Vesting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20VestingFilterer{contract: contract}, nil
}

// bindERC20Vesting binds a generic wrapper to an already deployed contract.
func bindERC20Vesting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20VestingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Vesting *ERC20VestingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Vesting.Contract.ERC20VestingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Vesting *ERC20VestingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Vesting.Contract.ERC20VestingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Vesting *ERC20VestingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Vesting.Contract.ERC20VestingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Vesting *ERC20VestingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Vesting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Vesting *ERC20VestingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Vesting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Vesting *ERC20VestingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Vesting.Contract.contract.Transact(opts, method, params...)
}

// AccuracyScale is a free data retrieval call binding the contract method 0x903223bf.
//
// Solidity: function accuracy_scale() view returns(uint256)
func (_ERC20Vesting *ERC20VestingCaller) AccuracyScale(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Vesting.contract.Call(opts, &out, "accuracy_scale")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccuracyScale is a free data retrieval call binding the contract method 0x903223bf.
//
// Solidity: function accuracy_scale() view returns(uint256)
func (_ERC20Vesting *ERC20VestingSession) AccuracyScale() (*big.Int, error) {
	return _ERC20Vesting.Contract.AccuracyScale(&_ERC20Vesting.CallOpts)
}

// AccuracyScale is a free data retrieval call binding the contract method 0x903223bf.
//
// Solidity: function accuracy_scale() view returns(uint256)
func (_ERC20Vesting *ERC20VestingCallerSession) AccuracyScale() (*big.Int, error) {
	return _ERC20Vesting.Contract.AccuracyScale(&_ERC20Vesting.CallOpts)
}

// AddressMigration is a free data retrieval call binding the contract method 0x59e7454b.
//
// Solidity: function address_migration(address ) view returns(address)
func (_ERC20Vesting *ERC20VestingCaller) AddressMigration(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _ERC20Vesting.contract.Call(opts, &out, "address_migration", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressMigration is a free data retrieval call binding the contract method 0x59e7454b.
//
// Solidity: function address_migration(address ) view returns(address)
func (_ERC20Vesting *ERC20VestingSession) AddressMigration(arg0 common.Address) (common.Address, error) {
	return _ERC20Vesting.Contract.AddressMigration(&_ERC20Vesting.CallOpts, arg0)
}

// AddressMigration is a free data retrieval call binding the contract method 0x59e7454b.
//
// Solidity: function address_migration(address ) view returns(address)
func (_ERC20Vesting *ERC20VestingCallerSession) AddressMigration(arg0 common.Address) (common.Address, error) {
	return _ERC20Vesting.Contract.AddressMigration(&_ERC20Vesting.CallOpts, arg0)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_ERC20Vesting *ERC20VestingCaller) Controller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20Vesting.contract.Call(opts, &out, "controller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_ERC20Vesting *ERC20VestingSession) Controller() (common.Address, error) {
	return _ERC20Vesting.Contract.Controller(&_ERC20Vesting.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_ERC20Vesting *ERC20VestingCallerSession) Controller() (common.Address, error) {
	return _ERC20Vesting.Contract.Controller(&_ERC20Vesting.CallOpts)
}

// DefaultTrancheId is a free data retrieval call binding the contract method 0xed1cac5c.
//
// Solidity: function default_tranche_id() view returns(uint8)
func (_ERC20Vesting *ERC20VestingCaller) DefaultTrancheId(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ERC20Vesting.contract.Call(opts, &out, "default_tranche_id")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// DefaultTrancheId is a free data retrieval call binding the contract method 0xed1cac5c.
//
// Solidity: function default_tranche_id() view returns(uint8)
func (_ERC20Vesting *ERC20VestingSession) DefaultTrancheId() (uint8, error) {
	return _ERC20Vesting.Contract.DefaultTrancheId(&_ERC20Vesting.CallOpts)
}

// DefaultTrancheId is a free data retrieval call binding the contract method 0xed1cac5c.
//
// Solidity: function default_tranche_id() view returns(uint8)
func (_ERC20Vesting *ERC20VestingCallerSession) DefaultTrancheId() (uint8, error) {
	return _ERC20Vesting.Contract.DefaultTrancheId(&_ERC20Vesting.CallOpts)
}

// GetTrancheBalance is a free data retrieval call binding the contract method 0x02e8e239.
//
// Solidity: function get_tranche_balance(address user, uint8 tranche_id) view returns(uint256)
func (_ERC20Vesting *ERC20VestingCaller) GetTrancheBalance(opts *bind.CallOpts, user common.Address, tranche_id uint8) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Vesting.contract.Call(opts, &out, "get_tranche_balance", user, tranche_id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTrancheBalance is a free data retrieval call binding the contract method 0x02e8e239.
//
// Solidity: function get_tranche_balance(address user, uint8 tranche_id) view returns(uint256)
func (_ERC20Vesting *ERC20VestingSession) GetTrancheBalance(user common.Address, tranche_id uint8) (*big.Int, error) {
	return _ERC20Vesting.Contract.GetTrancheBalance(&_ERC20Vesting.CallOpts, user, tranche_id)
}

// GetTrancheBalance is a free data retrieval call binding the contract method 0x02e8e239.
//
// Solidity: function get_tranche_balance(address user, uint8 tranche_id) view returns(uint256)
func (_ERC20Vesting *ERC20VestingCallerSession) GetTrancheBalance(user common.Address, tranche_id uint8) (*big.Int, error) {
	return _ERC20Vesting.Contract.GetTrancheBalance(&_ERC20Vesting.CallOpts, user, tranche_id)
}

// GetVestedForTranche is a free data retrieval call binding the contract method 0x29f7de26.
//
// Solidity: function get_vested_for_tranche(address user, uint8 tranche_id) view returns(uint256)
func (_ERC20Vesting *ERC20VestingCaller) GetVestedForTranche(opts *bind.CallOpts, user common.Address, tranche_id uint8) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Vesting.contract.Call(opts, &out, "get_vested_for_tranche", user, tranche_id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVestedForTranche is a free data retrieval call binding the contract method 0x29f7de26.
//
// Solidity: function get_vested_for_tranche(address user, uint8 tranche_id) view returns(uint256)
func (_ERC20Vesting *ERC20VestingSession) GetVestedForTranche(user common.Address, tranche_id uint8) (*big.Int, error) {
	return _ERC20Vesting.Contract.GetVestedForTranche(&_ERC20Vesting.CallOpts, user, tranche_id)
}

// GetVestedForTranche is a free data retrieval call binding the contract method 0x29f7de26.
//
// Solidity: function get_vested_for_tranche(address user, uint8 tranche_id) view returns(uint256)
func (_ERC20Vesting *ERC20VestingCallerSession) GetVestedForTranche(user common.Address, tranche_id uint8) (*big.Int, error) {
	return _ERC20Vesting.Contract.GetVestedForTranche(&_ERC20Vesting.CallOpts, user, tranche_id)
}

// PermittedIssuance is a free data retrieval call binding the contract method 0xa8be47e4.
//
// Solidity: function permitted_issuance(address ) view returns(uint256)
func (_ERC20Vesting *ERC20VestingCaller) PermittedIssuance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Vesting.contract.Call(opts, &out, "permitted_issuance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PermittedIssuance is a free data retrieval call binding the contract method 0xa8be47e4.
//
// Solidity: function permitted_issuance(address ) view returns(uint256)
func (_ERC20Vesting *ERC20VestingSession) PermittedIssuance(arg0 common.Address) (*big.Int, error) {
	return _ERC20Vesting.Contract.PermittedIssuance(&_ERC20Vesting.CallOpts, arg0)
}

// PermittedIssuance is a free data retrieval call binding the contract method 0xa8be47e4.
//
// Solidity: function permitted_issuance(address ) view returns(uint256)
func (_ERC20Vesting *ERC20VestingCallerSession) PermittedIssuance(arg0 common.Address) (*big.Int, error) {
	return _ERC20Vesting.Contract.PermittedIssuance(&_ERC20Vesting.CallOpts, arg0)
}

// StakeBalance is a free data retrieval call binding the contract method 0x274abf34.
//
// Solidity: function stake_balance(address target, bytes32 vega_public_key) view returns(uint256)
func (_ERC20Vesting *ERC20VestingCaller) StakeBalance(opts *bind.CallOpts, target common.Address, vega_public_key [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Vesting.contract.Call(opts, &out, "stake_balance", target, vega_public_key)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeBalance is a free data retrieval call binding the contract method 0x274abf34.
//
// Solidity: function stake_balance(address target, bytes32 vega_public_key) view returns(uint256)
func (_ERC20Vesting *ERC20VestingSession) StakeBalance(target common.Address, vega_public_key [32]byte) (*big.Int, error) {
	return _ERC20Vesting.Contract.StakeBalance(&_ERC20Vesting.CallOpts, target, vega_public_key)
}

// StakeBalance is a free data retrieval call binding the contract method 0x274abf34.
//
// Solidity: function stake_balance(address target, bytes32 vega_public_key) view returns(uint256)
func (_ERC20Vesting *ERC20VestingCallerSession) StakeBalance(target common.Address, vega_public_key [32]byte) (*big.Int, error) {
	return _ERC20Vesting.Contract.StakeBalance(&_ERC20Vesting.CallOpts, target, vega_public_key)
}

// StakingToken is a free data retrieval call binding the contract method 0x2dc7d74c.
//
// Solidity: function staking_token() view returns(address)
func (_ERC20Vesting *ERC20VestingCaller) StakingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20Vesting.contract.Call(opts, &out, "staking_token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingToken is a free data retrieval call binding the contract method 0x2dc7d74c.
//
// Solidity: function staking_token() view returns(address)
func (_ERC20Vesting *ERC20VestingSession) StakingToken() (common.Address, error) {
	return _ERC20Vesting.Contract.StakingToken(&_ERC20Vesting.CallOpts)
}

// StakingToken is a free data retrieval call binding the contract method 0x2dc7d74c.
//
// Solidity: function staking_token() view returns(address)
func (_ERC20Vesting *ERC20VestingCallerSession) StakingToken() (common.Address, error) {
	return _ERC20Vesting.Contract.StakingToken(&_ERC20Vesting.CallOpts)
}

// TotalLocked is a free data retrieval call binding the contract method 0x3c48a620.
//
// Solidity: function total_locked() view returns(uint256)
func (_ERC20Vesting *ERC20VestingCaller) TotalLocked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Vesting.contract.Call(opts, &out, "total_locked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalLocked is a free data retrieval call binding the contract method 0x3c48a620.
//
// Solidity: function total_locked() view returns(uint256)
func (_ERC20Vesting *ERC20VestingSession) TotalLocked() (*big.Int, error) {
	return _ERC20Vesting.Contract.TotalLocked(&_ERC20Vesting.CallOpts)
}

// TotalLocked is a free data retrieval call binding the contract method 0x3c48a620.
//
// Solidity: function total_locked() view returns(uint256)
func (_ERC20Vesting *ERC20VestingCallerSession) TotalLocked() (*big.Int, error) {
	return _ERC20Vesting.Contract.TotalLocked(&_ERC20Vesting.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0xaf7568dd.
//
// Solidity: function total_staked() view returns(uint256)
func (_ERC20Vesting *ERC20VestingCaller) TotalStaked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Vesting.contract.Call(opts, &out, "total_staked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStaked is a free data retrieval call binding the contract method 0xaf7568dd.
//
// Solidity: function total_staked() view returns(uint256)
func (_ERC20Vesting *ERC20VestingSession) TotalStaked() (*big.Int, error) {
	return _ERC20Vesting.Contract.TotalStaked(&_ERC20Vesting.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0xaf7568dd.
//
// Solidity: function total_staked() view returns(uint256)
func (_ERC20Vesting *ERC20VestingCallerSession) TotalStaked() (*big.Int, error) {
	return _ERC20Vesting.Contract.TotalStaked(&_ERC20Vesting.CallOpts)
}

// TrancheCount is a free data retrieval call binding the contract method 0xeb1cbd27.
//
// Solidity: function tranche_count() view returns(uint8)
func (_ERC20Vesting *ERC20VestingCaller) TrancheCount(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ERC20Vesting.contract.Call(opts, &out, "tranche_count")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TrancheCount is a free data retrieval call binding the contract method 0xeb1cbd27.
//
// Solidity: function tranche_count() view returns(uint8)
func (_ERC20Vesting *ERC20VestingSession) TrancheCount() (uint8, error) {
	return _ERC20Vesting.Contract.TrancheCount(&_ERC20Vesting.CallOpts)
}

// TrancheCount is a free data retrieval call binding the contract method 0xeb1cbd27.
//
// Solidity: function tranche_count() view returns(uint8)
func (_ERC20Vesting *ERC20VestingCallerSession) TrancheCount() (uint8, error) {
	return _ERC20Vesting.Contract.TrancheCount(&_ERC20Vesting.CallOpts)
}

// Tranches is a free data retrieval call binding the contract method 0x572dd5a9.
//
// Solidity: function tranches(uint8 ) view returns(uint256 cliff_start, uint256 duration)
func (_ERC20Vesting *ERC20VestingCaller) Tranches(opts *bind.CallOpts, arg0 uint8) (struct {
	CliffStart *big.Int
	Duration   *big.Int
}, error) {
	var out []interface{}
	err := _ERC20Vesting.contract.Call(opts, &out, "tranches", arg0)

	outstruct := new(struct {
		CliffStart *big.Int
		Duration   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CliffStart = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Duration = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Tranches is a free data retrieval call binding the contract method 0x572dd5a9.
//
// Solidity: function tranches(uint8 ) view returns(uint256 cliff_start, uint256 duration)
func (_ERC20Vesting *ERC20VestingSession) Tranches(arg0 uint8) (struct {
	CliffStart *big.Int
	Duration   *big.Int
}, error) {
	return _ERC20Vesting.Contract.Tranches(&_ERC20Vesting.CallOpts, arg0)
}

// Tranches is a free data retrieval call binding the contract method 0x572dd5a9.
//
// Solidity: function tranches(uint8 ) view returns(uint256 cliff_start, uint256 duration)
func (_ERC20Vesting *ERC20VestingCallerSession) Tranches(arg0 uint8) (struct {
	CliffStart *big.Int
	Duration   *big.Int
}, error) {
	return _ERC20Vesting.Contract.Tranches(&_ERC20Vesting.CallOpts, arg0)
}

// UserStats is a free data retrieval call binding the contract method 0xe96842ec.
//
// Solidity: function user_stats(address ) view returns(uint256 total_in_all_tranches, uint256 lien)
func (_ERC20Vesting *ERC20VestingCaller) UserStats(opts *bind.CallOpts, arg0 common.Address) (struct {
	TotalInAllTranches *big.Int
	Lien               *big.Int
}, error) {
	var out []interface{}
	err := _ERC20Vesting.contract.Call(opts, &out, "user_stats", arg0)

	outstruct := new(struct {
		TotalInAllTranches *big.Int
		Lien               *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalInAllTranches = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Lien = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserStats is a free data retrieval call binding the contract method 0xe96842ec.
//
// Solidity: function user_stats(address ) view returns(uint256 total_in_all_tranches, uint256 lien)
func (_ERC20Vesting *ERC20VestingSession) UserStats(arg0 common.Address) (struct {
	TotalInAllTranches *big.Int
	Lien               *big.Int
}, error) {
	return _ERC20Vesting.Contract.UserStats(&_ERC20Vesting.CallOpts, arg0)
}

// UserStats is a free data retrieval call binding the contract method 0xe96842ec.
//
// Solidity: function user_stats(address ) view returns(uint256 total_in_all_tranches, uint256 lien)
func (_ERC20Vesting *ERC20VestingCallerSession) UserStats(arg0 common.Address) (struct {
	TotalInAllTranches *big.Int
	Lien               *big.Int
}, error) {
	return _ERC20Vesting.Contract.UserStats(&_ERC20Vesting.CallOpts, arg0)
}

// UserTotalAllTranches is a free data retrieval call binding the contract method 0x15e1a382.
//
// Solidity: function user_total_all_tranches(address user) view returns(uint256)
func (_ERC20Vesting *ERC20VestingCaller) UserTotalAllTranches(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Vesting.contract.Call(opts, &out, "user_total_all_tranches", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserTotalAllTranches is a free data retrieval call binding the contract method 0x15e1a382.
//
// Solidity: function user_total_all_tranches(address user) view returns(uint256)
func (_ERC20Vesting *ERC20VestingSession) UserTotalAllTranches(user common.Address) (*big.Int, error) {
	return _ERC20Vesting.Contract.UserTotalAllTranches(&_ERC20Vesting.CallOpts, user)
}

// UserTotalAllTranches is a free data retrieval call binding the contract method 0x15e1a382.
//
// Solidity: function user_total_all_tranches(address user) view returns(uint256)
func (_ERC20Vesting *ERC20VestingCallerSession) UserTotalAllTranches(user common.Address) (*big.Int, error) {
	return _ERC20Vesting.Contract.UserTotalAllTranches(&_ERC20Vesting.CallOpts, user)
}

// V1Address is a free data retrieval call binding the contract method 0xa7cf676e.
//
// Solidity: function v1_address() view returns(address)
func (_ERC20Vesting *ERC20VestingCaller) V1Address(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20Vesting.contract.Call(opts, &out, "v1_address")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// V1Address is a free data retrieval call binding the contract method 0xa7cf676e.
//
// Solidity: function v1_address() view returns(address)
func (_ERC20Vesting *ERC20VestingSession) V1Address() (common.Address, error) {
	return _ERC20Vesting.Contract.V1Address(&_ERC20Vesting.CallOpts)
}

// V1Address is a free data retrieval call binding the contract method 0xa7cf676e.
//
// Solidity: function v1_address() view returns(address)
func (_ERC20Vesting *ERC20VestingCallerSession) V1Address() (common.Address, error) {
	return _ERC20Vesting.Contract.V1Address(&_ERC20Vesting.CallOpts)
}

// V1Migrated is a free data retrieval call binding the contract method 0x6bb7d856.
//
// Solidity: function v1_migrated(address ) view returns(bool)
func (_ERC20Vesting *ERC20VestingCaller) V1Migrated(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ERC20Vesting.contract.Call(opts, &out, "v1_migrated", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// V1Migrated is a free data retrieval call binding the contract method 0x6bb7d856.
//
// Solidity: function v1_migrated(address ) view returns(bool)
func (_ERC20Vesting *ERC20VestingSession) V1Migrated(arg0 common.Address) (bool, error) {
	return _ERC20Vesting.Contract.V1Migrated(&_ERC20Vesting.CallOpts, arg0)
}

// V1Migrated is a free data retrieval call binding the contract method 0x6bb7d856.
//
// Solidity: function v1_migrated(address ) view returns(bool)
func (_ERC20Vesting *ERC20VestingCallerSession) V1Migrated(arg0 common.Address) (bool, error) {
	return _ERC20Vesting.Contract.V1Migrated(&_ERC20Vesting.CallOpts, arg0)
}

// V2Address is a free data retrieval call binding the contract method 0x7d89651a.
//
// Solidity: function v2_address() view returns(address)
func (_ERC20Vesting *ERC20VestingCaller) V2Address(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20Vesting.contract.Call(opts, &out, "v2_address")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// V2Address is a free data retrieval call binding the contract method 0x7d89651a.
//
// Solidity: function v2_address() view returns(address)
func (_ERC20Vesting *ERC20VestingSession) V2Address() (common.Address, error) {
	return _ERC20Vesting.Contract.V2Address(&_ERC20Vesting.CallOpts)
}

// V2Address is a free data retrieval call binding the contract method 0x7d89651a.
//
// Solidity: function v2_address() view returns(address)
func (_ERC20Vesting *ERC20VestingCallerSession) V2Address() (common.Address, error) {
	return _ERC20Vesting.Contract.V2Address(&_ERC20Vesting.CallOpts)
}

// AssistedWithdrawFromTranche is a paid mutator transaction binding the contract method 0x3fd60603.
//
// Solidity: function assisted_withdraw_from_tranche(uint8 tranche_id, address target) returns()
func (_ERC20Vesting *ERC20VestingTransactor) AssistedWithdrawFromTranche(opts *bind.TransactOpts, tranche_id uint8, target common.Address) (*types.Transaction, error) {
	return _ERC20Vesting.contract.Transact(opts, "assisted_withdraw_from_tranche", tranche_id, target)
}

// AssistedWithdrawFromTranche is a paid mutator transaction binding the contract method 0x3fd60603.
//
// Solidity: function assisted_withdraw_from_tranche(uint8 tranche_id, address target) returns()
func (_ERC20Vesting *ERC20VestingSession) AssistedWithdrawFromTranche(tranche_id uint8, target common.Address) (*types.Transaction, error) {
	return _ERC20Vesting.Contract.AssistedWithdrawFromTranche(&_ERC20Vesting.TransactOpts, tranche_id, target)
}

// AssistedWithdrawFromTranche is a paid mutator transaction binding the contract method 0x3fd60603.
//
// Solidity: function assisted_withdraw_from_tranche(uint8 tranche_id, address target) returns()
func (_ERC20Vesting *ERC20VestingTransactorSession) AssistedWithdrawFromTranche(tranche_id uint8, target common.Address) (*types.Transaction, error) {
	return _ERC20Vesting.Contract.AssistedWithdrawFromTranche(&_ERC20Vesting.TransactOpts, tranche_id, target)
}

// CreateTranche is a paid mutator transaction binding the contract method 0x66891093.
//
// Solidity: function create_tranche(uint256 cliff_start, uint256 duration) returns()
func (_ERC20Vesting *ERC20VestingTransactor) CreateTranche(opts *bind.TransactOpts, cliff_start *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _ERC20Vesting.contract.Transact(opts, "create_tranche", cliff_start, duration)
}

// CreateTranche is a paid mutator transaction binding the contract method 0x66891093.
//
// Solidity: function create_tranche(uint256 cliff_start, uint256 duration) returns()
func (_ERC20Vesting *ERC20VestingSession) CreateTranche(cliff_start *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _ERC20Vesting.Contract.CreateTranche(&_ERC20Vesting.TransactOpts, cliff_start, duration)
}

// CreateTranche is a paid mutator transaction binding the contract method 0x66891093.
//
// Solidity: function create_tranche(uint256 cliff_start, uint256 duration) returns()
func (_ERC20Vesting *ERC20VestingTransactorSession) CreateTranche(cliff_start *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _ERC20Vesting.Contract.CreateTranche(&_ERC20Vesting.TransactOpts, cliff_start, duration)
}

// IssueIntoTranche is a paid mutator transaction binding the contract method 0xe2de6e6d.
//
// Solidity: function issue_into_tranche(address user, uint8 tranche_id, uint256 amount) returns()
func (_ERC20Vesting *ERC20VestingTransactor) IssueIntoTranche(opts *bind.TransactOpts, user common.Address, tranche_id uint8, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Vesting.contract.Transact(opts, "issue_into_tranche", user, tranche_id, amount)
}

// IssueIntoTranche is a paid mutator transaction binding the contract method 0xe2de6e6d.
//
// Solidity: function issue_into_tranche(address user, uint8 tranche_id, uint256 amount) returns()
func (_ERC20Vesting *ERC20VestingSession) IssueIntoTranche(user common.Address, tranche_id uint8, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Vesting.Contract.IssueIntoTranche(&_ERC20Vesting.TransactOpts, user, tranche_id, amount)
}

// IssueIntoTranche is a paid mutator transaction binding the contract method 0xe2de6e6d.
//
// Solidity: function issue_into_tranche(address user, uint8 tranche_id, uint256 amount) returns()
func (_ERC20Vesting *ERC20VestingTransactorSession) IssueIntoTranche(user common.Address, tranche_id uint8, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Vesting.Contract.IssueIntoTranche(&_ERC20Vesting.TransactOpts, user, tranche_id, amount)
}

// MoveIntoTranche is a paid mutator transaction binding the contract method 0x912d66ff.
//
// Solidity: function move_into_tranche(address user, uint8 tranche_id, uint256 amount) returns()
func (_ERC20Vesting *ERC20VestingTransactor) MoveIntoTranche(opts *bind.TransactOpts, user common.Address, tranche_id uint8, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Vesting.contract.Transact(opts, "move_into_tranche", user, tranche_id, amount)
}

// MoveIntoTranche is a paid mutator transaction binding the contract method 0x912d66ff.
//
// Solidity: function move_into_tranche(address user, uint8 tranche_id, uint256 amount) returns()
func (_ERC20Vesting *ERC20VestingSession) MoveIntoTranche(user common.Address, tranche_id uint8, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Vesting.Contract.MoveIntoTranche(&_ERC20Vesting.TransactOpts, user, tranche_id, amount)
}

// MoveIntoTranche is a paid mutator transaction binding the contract method 0x912d66ff.
//
// Solidity: function move_into_tranche(address user, uint8 tranche_id, uint256 amount) returns()
func (_ERC20Vesting *ERC20VestingTransactorSession) MoveIntoTranche(user common.Address, tranche_id uint8, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Vesting.Contract.MoveIntoTranche(&_ERC20Vesting.TransactOpts, user, tranche_id, amount)
}

// PermitIssuer is a paid mutator transaction binding the contract method 0x277817c7.
//
// Solidity: function permit_issuer(address issuer, uint256 amount) returns()
func (_ERC20Vesting *ERC20VestingTransactor) PermitIssuer(opts *bind.TransactOpts, issuer common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Vesting.contract.Transact(opts, "permit_issuer", issuer, amount)
}

// PermitIssuer is a paid mutator transaction binding the contract method 0x277817c7.
//
// Solidity: function permit_issuer(address issuer, uint256 amount) returns()
func (_ERC20Vesting *ERC20VestingSession) PermitIssuer(issuer common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Vesting.Contract.PermitIssuer(&_ERC20Vesting.TransactOpts, issuer, amount)
}

// PermitIssuer is a paid mutator transaction binding the contract method 0x277817c7.
//
// Solidity: function permit_issuer(address issuer, uint256 amount) returns()
func (_ERC20Vesting *ERC20VestingTransactorSession) PermitIssuer(issuer common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Vesting.Contract.PermitIssuer(&_ERC20Vesting.TransactOpts, issuer, amount)
}

// RemoveStake is a paid mutator transaction binding the contract method 0xdd01ba0b.
//
// Solidity: function remove_stake(uint256 amount, bytes32 vega_public_key) returns()
func (_ERC20Vesting *ERC20VestingTransactor) RemoveStake(opts *bind.TransactOpts, amount *big.Int, vega_public_key [32]byte) (*types.Transaction, error) {
	return _ERC20Vesting.contract.Transact(opts, "remove_stake", amount, vega_public_key)
}

// RemoveStake is a paid mutator transaction binding the contract method 0xdd01ba0b.
//
// Solidity: function remove_stake(uint256 amount, bytes32 vega_public_key) returns()
func (_ERC20Vesting *ERC20VestingSession) RemoveStake(amount *big.Int, vega_public_key [32]byte) (*types.Transaction, error) {
	return _ERC20Vesting.Contract.RemoveStake(&_ERC20Vesting.TransactOpts, amount, vega_public_key)
}

// RemoveStake is a paid mutator transaction binding the contract method 0xdd01ba0b.
//
// Solidity: function remove_stake(uint256 amount, bytes32 vega_public_key) returns()
func (_ERC20Vesting *ERC20VestingTransactorSession) RemoveStake(amount *big.Int, vega_public_key [32]byte) (*types.Transaction, error) {
	return _ERC20Vesting.Contract.RemoveStake(&_ERC20Vesting.TransactOpts, amount, vega_public_key)
}

// RevokeIssuer is a paid mutator transaction binding the contract method 0x94a4ff3f.
//
// Solidity: function revoke_issuer(address issuer) returns()
func (_ERC20Vesting *ERC20VestingTransactor) RevokeIssuer(opts *bind.TransactOpts, issuer common.Address) (*types.Transaction, error) {
	return _ERC20Vesting.contract.Transact(opts, "revoke_issuer", issuer)
}

// RevokeIssuer is a paid mutator transaction binding the contract method 0x94a4ff3f.
//
// Solidity: function revoke_issuer(address issuer) returns()
func (_ERC20Vesting *ERC20VestingSession) RevokeIssuer(issuer common.Address) (*types.Transaction, error) {
	return _ERC20Vesting.Contract.RevokeIssuer(&_ERC20Vesting.TransactOpts, issuer)
}

// RevokeIssuer is a paid mutator transaction binding the contract method 0x94a4ff3f.
//
// Solidity: function revoke_issuer(address issuer) returns()
func (_ERC20Vesting *ERC20VestingTransactorSession) RevokeIssuer(issuer common.Address) (*types.Transaction, error) {
	return _ERC20Vesting.Contract.RevokeIssuer(&_ERC20Vesting.TransactOpts, issuer)
}

// SetController is a paid mutator transaction binding the contract method 0x91b10ffa.
//
// Solidity: function set_controller(address new_controller) returns()
func (_ERC20Vesting *ERC20VestingTransactor) SetController(opts *bind.TransactOpts, new_controller common.Address) (*types.Transaction, error) {
	return _ERC20Vesting.contract.Transact(opts, "set_controller", new_controller)
}

// SetController is a paid mutator transaction binding the contract method 0x91b10ffa.
//
// Solidity: function set_controller(address new_controller) returns()
func (_ERC20Vesting *ERC20VestingSession) SetController(new_controller common.Address) (*types.Transaction, error) {
	return _ERC20Vesting.Contract.SetController(&_ERC20Vesting.TransactOpts, new_controller)
}

// SetController is a paid mutator transaction binding the contract method 0x91b10ffa.
//
// Solidity: function set_controller(address new_controller) returns()
func (_ERC20Vesting *ERC20VestingTransactorSession) SetController(new_controller common.Address) (*types.Transaction, error) {
	return _ERC20Vesting.Contract.SetController(&_ERC20Vesting.TransactOpts, new_controller)
}

// StakeTokens is a paid mutator transaction binding the contract method 0xeea70b42.
//
// Solidity: function stake_tokens(uint256 amount, bytes32 vega_public_key) returns()
func (_ERC20Vesting *ERC20VestingTransactor) StakeTokens(opts *bind.TransactOpts, amount *big.Int, vega_public_key [32]byte) (*types.Transaction, error) {
	return _ERC20Vesting.contract.Transact(opts, "stake_tokens", amount, vega_public_key)
}

// StakeTokens is a paid mutator transaction binding the contract method 0xeea70b42.
//
// Solidity: function stake_tokens(uint256 amount, bytes32 vega_public_key) returns()
func (_ERC20Vesting *ERC20VestingSession) StakeTokens(amount *big.Int, vega_public_key [32]byte) (*types.Transaction, error) {
	return _ERC20Vesting.Contract.StakeTokens(&_ERC20Vesting.TransactOpts, amount, vega_public_key)
}

// StakeTokens is a paid mutator transaction binding the contract method 0xeea70b42.
//
// Solidity: function stake_tokens(uint256 amount, bytes32 vega_public_key) returns()
func (_ERC20Vesting *ERC20VestingTransactorSession) StakeTokens(amount *big.Int, vega_public_key [32]byte) (*types.Transaction, error) {
	return _ERC20Vesting.Contract.StakeTokens(&_ERC20Vesting.TransactOpts, amount, vega_public_key)
}

// WithdrawFromTranche is a paid mutator transaction binding the contract method 0xbc6d385d.
//
// Solidity: function withdraw_from_tranche(uint8 tranche_id) returns()
func (_ERC20Vesting *ERC20VestingTransactor) WithdrawFromTranche(opts *bind.TransactOpts, tranche_id uint8) (*types.Transaction, error) {
	return _ERC20Vesting.contract.Transact(opts, "withdraw_from_tranche", tranche_id)
}

// WithdrawFromTranche is a paid mutator transaction binding the contract method 0xbc6d385d.
//
// Solidity: function withdraw_from_tranche(uint8 tranche_id) returns()
func (_ERC20Vesting *ERC20VestingSession) WithdrawFromTranche(tranche_id uint8) (*types.Transaction, error) {
	return _ERC20Vesting.Contract.WithdrawFromTranche(&_ERC20Vesting.TransactOpts, tranche_id)
}

// WithdrawFromTranche is a paid mutator transaction binding the contract method 0xbc6d385d.
//
// Solidity: function withdraw_from_tranche(uint8 tranche_id) returns()
func (_ERC20Vesting *ERC20VestingTransactorSession) WithdrawFromTranche(tranche_id uint8) (*types.Transaction, error) {
	return _ERC20Vesting.Contract.WithdrawFromTranche(&_ERC20Vesting.TransactOpts, tranche_id)
}

// ERC20VestingControllerSetIterator is returned from FilterControllerSet and is used to iterate over the raw logs and unpacked data for ControllerSet events raised by the ERC20Vesting contract.
type ERC20VestingControllerSetIterator struct {
	Event *ERC20VestingControllerSet // Event containing the contract specifics and raw log

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
func (it *ERC20VestingControllerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20VestingControllerSet)
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
		it.Event = new(ERC20VestingControllerSet)
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
func (it *ERC20VestingControllerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20VestingControllerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20VestingControllerSet represents a ControllerSet event raised by the ERC20Vesting contract.
type ERC20VestingControllerSet struct {
	NewController common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterControllerSet is a free log retrieval operation binding the contract event 0x587f8c3da67289e1b7fa2679aafa43e8ef08d5d8d9adfcb67cc825264c431c13.
//
// Solidity: event Controller_Set(address indexed new_controller)
func (_ERC20Vesting *ERC20VestingFilterer) FilterControllerSet(opts *bind.FilterOpts, new_controller []common.Address) (*ERC20VestingControllerSetIterator, error) {

	var new_controllerRule []interface{}
	for _, new_controllerItem := range new_controller {
		new_controllerRule = append(new_controllerRule, new_controllerItem)
	}

	logs, sub, err := _ERC20Vesting.contract.FilterLogs(opts, "Controller_Set", new_controllerRule)
	if err != nil {
		return nil, err
	}
	return &ERC20VestingControllerSetIterator{contract: _ERC20Vesting.contract, event: "Controller_Set", logs: logs, sub: sub}, nil
}

// WatchControllerSet is a free log subscription operation binding the contract event 0x587f8c3da67289e1b7fa2679aafa43e8ef08d5d8d9adfcb67cc825264c431c13.
//
// Solidity: event Controller_Set(address indexed new_controller)
func (_ERC20Vesting *ERC20VestingFilterer) WatchControllerSet(opts *bind.WatchOpts, sink chan<- *ERC20VestingControllerSet, new_controller []common.Address) (event.Subscription, error) {

	var new_controllerRule []interface{}
	for _, new_controllerItem := range new_controller {
		new_controllerRule = append(new_controllerRule, new_controllerItem)
	}

	logs, sub, err := _ERC20Vesting.contract.WatchLogs(opts, "Controller_Set", new_controllerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20VestingControllerSet)
				if err := _ERC20Vesting.contract.UnpackLog(event, "Controller_Set", log); err != nil {
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

// ParseControllerSet is a log parse operation binding the contract event 0x587f8c3da67289e1b7fa2679aafa43e8ef08d5d8d9adfcb67cc825264c431c13.
//
// Solidity: event Controller_Set(address indexed new_controller)
func (_ERC20Vesting *ERC20VestingFilterer) ParseControllerSet(log types.Log) (*ERC20VestingControllerSet, error) {
	event := new(ERC20VestingControllerSet)
	if err := _ERC20Vesting.contract.UnpackLog(event, "Controller_Set", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20VestingIssuerPermittedIterator is returned from FilterIssuerPermitted and is used to iterate over the raw logs and unpacked data for IssuerPermitted events raised by the ERC20Vesting contract.
type ERC20VestingIssuerPermittedIterator struct {
	Event *ERC20VestingIssuerPermitted // Event containing the contract specifics and raw log

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
func (it *ERC20VestingIssuerPermittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20VestingIssuerPermitted)
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
		it.Event = new(ERC20VestingIssuerPermitted)
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
func (it *ERC20VestingIssuerPermittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20VestingIssuerPermittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20VestingIssuerPermitted represents a IssuerPermitted event raised by the ERC20Vesting contract.
type ERC20VestingIssuerPermitted struct {
	Issuer common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterIssuerPermitted is a free log retrieval operation binding the contract event 0xb1dde3dfc5e8df8a1ad5d278b36a91192c9106fc9ae0d1730c82455692640cd8.
//
// Solidity: event Issuer_Permitted(address indexed issuer, uint256 amount)
func (_ERC20Vesting *ERC20VestingFilterer) FilterIssuerPermitted(opts *bind.FilterOpts, issuer []common.Address) (*ERC20VestingIssuerPermittedIterator, error) {

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _ERC20Vesting.contract.FilterLogs(opts, "Issuer_Permitted", issuerRule)
	if err != nil {
		return nil, err
	}
	return &ERC20VestingIssuerPermittedIterator{contract: _ERC20Vesting.contract, event: "Issuer_Permitted", logs: logs, sub: sub}, nil
}

// WatchIssuerPermitted is a free log subscription operation binding the contract event 0xb1dde3dfc5e8df8a1ad5d278b36a91192c9106fc9ae0d1730c82455692640cd8.
//
// Solidity: event Issuer_Permitted(address indexed issuer, uint256 amount)
func (_ERC20Vesting *ERC20VestingFilterer) WatchIssuerPermitted(opts *bind.WatchOpts, sink chan<- *ERC20VestingIssuerPermitted, issuer []common.Address) (event.Subscription, error) {

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _ERC20Vesting.contract.WatchLogs(opts, "Issuer_Permitted", issuerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20VestingIssuerPermitted)
				if err := _ERC20Vesting.contract.UnpackLog(event, "Issuer_Permitted", log); err != nil {
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

// ParseIssuerPermitted is a log parse operation binding the contract event 0xb1dde3dfc5e8df8a1ad5d278b36a91192c9106fc9ae0d1730c82455692640cd8.
//
// Solidity: event Issuer_Permitted(address indexed issuer, uint256 amount)
func (_ERC20Vesting *ERC20VestingFilterer) ParseIssuerPermitted(log types.Log) (*ERC20VestingIssuerPermitted, error) {
	event := new(ERC20VestingIssuerPermitted)
	if err := _ERC20Vesting.contract.UnpackLog(event, "Issuer_Permitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20VestingIssuerRevokedIterator is returned from FilterIssuerRevoked and is used to iterate over the raw logs and unpacked data for IssuerRevoked events raised by the ERC20Vesting contract.
type ERC20VestingIssuerRevokedIterator struct {
	Event *ERC20VestingIssuerRevoked // Event containing the contract specifics and raw log

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
func (it *ERC20VestingIssuerRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20VestingIssuerRevoked)
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
		it.Event = new(ERC20VestingIssuerRevoked)
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
func (it *ERC20VestingIssuerRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20VestingIssuerRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20VestingIssuerRevoked represents a IssuerRevoked event raised by the ERC20Vesting contract.
type ERC20VestingIssuerRevoked struct {
	Issuer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterIssuerRevoked is a free log retrieval operation binding the contract event 0xaa482319d1110d84c35ff519b80c0449540dfbdbe6ee1daf21c2483378a1bed0.
//
// Solidity: event Issuer_Revoked(address indexed issuer)
func (_ERC20Vesting *ERC20VestingFilterer) FilterIssuerRevoked(opts *bind.FilterOpts, issuer []common.Address) (*ERC20VestingIssuerRevokedIterator, error) {

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _ERC20Vesting.contract.FilterLogs(opts, "Issuer_Revoked", issuerRule)
	if err != nil {
		return nil, err
	}
	return &ERC20VestingIssuerRevokedIterator{contract: _ERC20Vesting.contract, event: "Issuer_Revoked", logs: logs, sub: sub}, nil
}

// WatchIssuerRevoked is a free log subscription operation binding the contract event 0xaa482319d1110d84c35ff519b80c0449540dfbdbe6ee1daf21c2483378a1bed0.
//
// Solidity: event Issuer_Revoked(address indexed issuer)
func (_ERC20Vesting *ERC20VestingFilterer) WatchIssuerRevoked(opts *bind.WatchOpts, sink chan<- *ERC20VestingIssuerRevoked, issuer []common.Address) (event.Subscription, error) {

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _ERC20Vesting.contract.WatchLogs(opts, "Issuer_Revoked", issuerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20VestingIssuerRevoked)
				if err := _ERC20Vesting.contract.UnpackLog(event, "Issuer_Revoked", log); err != nil {
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

// ParseIssuerRevoked is a log parse operation binding the contract event 0xaa482319d1110d84c35ff519b80c0449540dfbdbe6ee1daf21c2483378a1bed0.
//
// Solidity: event Issuer_Revoked(address indexed issuer)
func (_ERC20Vesting *ERC20VestingFilterer) ParseIssuerRevoked(log types.Log) (*ERC20VestingIssuerRevoked, error) {
	event := new(ERC20VestingIssuerRevoked)
	if err := _ERC20Vesting.contract.UnpackLog(event, "Issuer_Revoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20VestingStakeDepositedIterator is returned from FilterStakeDeposited and is used to iterate over the raw logs and unpacked data for StakeDeposited events raised by the ERC20Vesting contract.
type ERC20VestingStakeDepositedIterator struct {
	Event *ERC20VestingStakeDeposited // Event containing the contract specifics and raw log

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
func (it *ERC20VestingStakeDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20VestingStakeDeposited)
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
		it.Event = new(ERC20VestingStakeDeposited)
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
func (it *ERC20VestingStakeDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20VestingStakeDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20VestingStakeDeposited represents a StakeDeposited event raised by the ERC20Vesting contract.
type ERC20VestingStakeDeposited struct {
	User          common.Address
	Amount        *big.Int
	VegaPublicKey [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterStakeDeposited is a free log retrieval operation binding the contract event 0x9e3e33edf5dcded4adabc51b1266225d00fa41516bfcad69513fa4eca69519da.
//
// Solidity: event Stake_Deposited(address indexed user, uint256 amount, bytes32 indexed vega_public_key)
func (_ERC20Vesting *ERC20VestingFilterer) FilterStakeDeposited(opts *bind.FilterOpts, user []common.Address, vega_public_key [][32]byte) (*ERC20VestingStakeDepositedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	var vega_public_keyRule []interface{}
	for _, vega_public_keyItem := range vega_public_key {
		vega_public_keyRule = append(vega_public_keyRule, vega_public_keyItem)
	}

	logs, sub, err := _ERC20Vesting.contract.FilterLogs(opts, "Stake_Deposited", userRule, vega_public_keyRule)
	if err != nil {
		return nil, err
	}
	return &ERC20VestingStakeDepositedIterator{contract: _ERC20Vesting.contract, event: "Stake_Deposited", logs: logs, sub: sub}, nil
}

// WatchStakeDeposited is a free log subscription operation binding the contract event 0x9e3e33edf5dcded4adabc51b1266225d00fa41516bfcad69513fa4eca69519da.
//
// Solidity: event Stake_Deposited(address indexed user, uint256 amount, bytes32 indexed vega_public_key)
func (_ERC20Vesting *ERC20VestingFilterer) WatchStakeDeposited(opts *bind.WatchOpts, sink chan<- *ERC20VestingStakeDeposited, user []common.Address, vega_public_key [][32]byte) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	var vega_public_keyRule []interface{}
	for _, vega_public_keyItem := range vega_public_key {
		vega_public_keyRule = append(vega_public_keyRule, vega_public_keyItem)
	}

	logs, sub, err := _ERC20Vesting.contract.WatchLogs(opts, "Stake_Deposited", userRule, vega_public_keyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20VestingStakeDeposited)
				if err := _ERC20Vesting.contract.UnpackLog(event, "Stake_Deposited", log); err != nil {
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

// ParseStakeDeposited is a log parse operation binding the contract event 0x9e3e33edf5dcded4adabc51b1266225d00fa41516bfcad69513fa4eca69519da.
//
// Solidity: event Stake_Deposited(address indexed user, uint256 amount, bytes32 indexed vega_public_key)
func (_ERC20Vesting *ERC20VestingFilterer) ParseStakeDeposited(log types.Log) (*ERC20VestingStakeDeposited, error) {
	event := new(ERC20VestingStakeDeposited)
	if err := _ERC20Vesting.contract.UnpackLog(event, "Stake_Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20VestingStakeRemovedIterator is returned from FilterStakeRemoved and is used to iterate over the raw logs and unpacked data for StakeRemoved events raised by the ERC20Vesting contract.
type ERC20VestingStakeRemovedIterator struct {
	Event *ERC20VestingStakeRemoved // Event containing the contract specifics and raw log

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
func (it *ERC20VestingStakeRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20VestingStakeRemoved)
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
		it.Event = new(ERC20VestingStakeRemoved)
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
func (it *ERC20VestingStakeRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20VestingStakeRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20VestingStakeRemoved represents a StakeRemoved event raised by the ERC20Vesting contract.
type ERC20VestingStakeRemoved struct {
	User          common.Address
	Amount        *big.Int
	VegaPublicKey [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterStakeRemoved is a free log retrieval operation binding the contract event 0xa131d16963736e4c641f27a7f82f2e350b5971e555ae06ae906892bbba0a0939.
//
// Solidity: event Stake_Removed(address indexed user, uint256 amount, bytes32 indexed vega_public_key)
func (_ERC20Vesting *ERC20VestingFilterer) FilterStakeRemoved(opts *bind.FilterOpts, user []common.Address, vega_public_key [][32]byte) (*ERC20VestingStakeRemovedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	var vega_public_keyRule []interface{}
	for _, vega_public_keyItem := range vega_public_key {
		vega_public_keyRule = append(vega_public_keyRule, vega_public_keyItem)
	}

	logs, sub, err := _ERC20Vesting.contract.FilterLogs(opts, "Stake_Removed", userRule, vega_public_keyRule)
	if err != nil {
		return nil, err
	}
	return &ERC20VestingStakeRemovedIterator{contract: _ERC20Vesting.contract, event: "Stake_Removed", logs: logs, sub: sub}, nil
}

// WatchStakeRemoved is a free log subscription operation binding the contract event 0xa131d16963736e4c641f27a7f82f2e350b5971e555ae06ae906892bbba0a0939.
//
// Solidity: event Stake_Removed(address indexed user, uint256 amount, bytes32 indexed vega_public_key)
func (_ERC20Vesting *ERC20VestingFilterer) WatchStakeRemoved(opts *bind.WatchOpts, sink chan<- *ERC20VestingStakeRemoved, user []common.Address, vega_public_key [][32]byte) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	var vega_public_keyRule []interface{}
	for _, vega_public_keyItem := range vega_public_key {
		vega_public_keyRule = append(vega_public_keyRule, vega_public_keyItem)
	}

	logs, sub, err := _ERC20Vesting.contract.WatchLogs(opts, "Stake_Removed", userRule, vega_public_keyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20VestingStakeRemoved)
				if err := _ERC20Vesting.contract.UnpackLog(event, "Stake_Removed", log); err != nil {
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

// ParseStakeRemoved is a log parse operation binding the contract event 0xa131d16963736e4c641f27a7f82f2e350b5971e555ae06ae906892bbba0a0939.
//
// Solidity: event Stake_Removed(address indexed user, uint256 amount, bytes32 indexed vega_public_key)
func (_ERC20Vesting *ERC20VestingFilterer) ParseStakeRemoved(log types.Log) (*ERC20VestingStakeRemoved, error) {
	event := new(ERC20VestingStakeRemoved)
	if err := _ERC20Vesting.contract.UnpackLog(event, "Stake_Removed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20VestingStakeTransferredIterator is returned from FilterStakeTransferred and is used to iterate over the raw logs and unpacked data for StakeTransferred events raised by the ERC20Vesting contract.
type ERC20VestingStakeTransferredIterator struct {
	Event *ERC20VestingStakeTransferred // Event containing the contract specifics and raw log

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
func (it *ERC20VestingStakeTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20VestingStakeTransferred)
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
		it.Event = new(ERC20VestingStakeTransferred)
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
func (it *ERC20VestingStakeTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20VestingStakeTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20VestingStakeTransferred represents a StakeTransferred event raised by the ERC20Vesting contract.
type ERC20VestingStakeTransferred struct {
	From          common.Address
	Amount        *big.Int
	To            common.Address
	VegaPublicKey [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterStakeTransferred is a free log retrieval operation binding the contract event 0x296aca09e6f616abedcd9cd45ac378207310452b7a713289374fd1b35e2c2fbe.
//
// Solidity: event Stake_Transferred(address indexed from, uint256 amount, address indexed to, bytes32 indexed vega_public_key)
func (_ERC20Vesting *ERC20VestingFilterer) FilterStakeTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address, vega_public_key [][32]byte) (*ERC20VestingStakeTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var vega_public_keyRule []interface{}
	for _, vega_public_keyItem := range vega_public_key {
		vega_public_keyRule = append(vega_public_keyRule, vega_public_keyItem)
	}

	logs, sub, err := _ERC20Vesting.contract.FilterLogs(opts, "Stake_Transferred", fromRule, toRule, vega_public_keyRule)
	if err != nil {
		return nil, err
	}
	return &ERC20VestingStakeTransferredIterator{contract: _ERC20Vesting.contract, event: "Stake_Transferred", logs: logs, sub: sub}, nil
}

// WatchStakeTransferred is a free log subscription operation binding the contract event 0x296aca09e6f616abedcd9cd45ac378207310452b7a713289374fd1b35e2c2fbe.
//
// Solidity: event Stake_Transferred(address indexed from, uint256 amount, address indexed to, bytes32 indexed vega_public_key)
func (_ERC20Vesting *ERC20VestingFilterer) WatchStakeTransferred(opts *bind.WatchOpts, sink chan<- *ERC20VestingStakeTransferred, from []common.Address, to []common.Address, vega_public_key [][32]byte) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var vega_public_keyRule []interface{}
	for _, vega_public_keyItem := range vega_public_key {
		vega_public_keyRule = append(vega_public_keyRule, vega_public_keyItem)
	}

	logs, sub, err := _ERC20Vesting.contract.WatchLogs(opts, "Stake_Transferred", fromRule, toRule, vega_public_keyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20VestingStakeTransferred)
				if err := _ERC20Vesting.contract.UnpackLog(event, "Stake_Transferred", log); err != nil {
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

// ParseStakeTransferred is a log parse operation binding the contract event 0x296aca09e6f616abedcd9cd45ac378207310452b7a713289374fd1b35e2c2fbe.
//
// Solidity: event Stake_Transferred(address indexed from, uint256 amount, address indexed to, bytes32 indexed vega_public_key)
func (_ERC20Vesting *ERC20VestingFilterer) ParseStakeTransferred(log types.Log) (*ERC20VestingStakeTransferred, error) {
	event := new(ERC20VestingStakeTransferred)
	if err := _ERC20Vesting.contract.UnpackLog(event, "Stake_Transferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20VestingTrancheBalanceAddedIterator is returned from FilterTrancheBalanceAdded and is used to iterate over the raw logs and unpacked data for TrancheBalanceAdded events raised by the ERC20Vesting contract.
type ERC20VestingTrancheBalanceAddedIterator struct {
	Event *ERC20VestingTrancheBalanceAdded // Event containing the contract specifics and raw log

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
func (it *ERC20VestingTrancheBalanceAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20VestingTrancheBalanceAdded)
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
		it.Event = new(ERC20VestingTrancheBalanceAdded)
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
func (it *ERC20VestingTrancheBalanceAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20VestingTrancheBalanceAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20VestingTrancheBalanceAdded represents a TrancheBalanceAdded event raised by the ERC20Vesting contract.
type ERC20VestingTrancheBalanceAdded struct {
	User      common.Address
	TrancheId uint8
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTrancheBalanceAdded is a free log retrieval operation binding the contract event 0xd574b9b920ef951710b35bf90459120b667d6c810777f945e3e82996f75824c4.
//
// Solidity: event Tranche_Balance_Added(address indexed user, uint8 indexed tranche_id, uint256 amount)
func (_ERC20Vesting *ERC20VestingFilterer) FilterTrancheBalanceAdded(opts *bind.FilterOpts, user []common.Address, tranche_id []uint8) (*ERC20VestingTrancheBalanceAddedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tranche_idRule []interface{}
	for _, tranche_idItem := range tranche_id {
		tranche_idRule = append(tranche_idRule, tranche_idItem)
	}

	logs, sub, err := _ERC20Vesting.contract.FilterLogs(opts, "Tranche_Balance_Added", userRule, tranche_idRule)
	if err != nil {
		return nil, err
	}
	return &ERC20VestingTrancheBalanceAddedIterator{contract: _ERC20Vesting.contract, event: "Tranche_Balance_Added", logs: logs, sub: sub}, nil
}

// WatchTrancheBalanceAdded is a free log subscription operation binding the contract event 0xd574b9b920ef951710b35bf90459120b667d6c810777f945e3e82996f75824c4.
//
// Solidity: event Tranche_Balance_Added(address indexed user, uint8 indexed tranche_id, uint256 amount)
func (_ERC20Vesting *ERC20VestingFilterer) WatchTrancheBalanceAdded(opts *bind.WatchOpts, sink chan<- *ERC20VestingTrancheBalanceAdded, user []common.Address, tranche_id []uint8) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tranche_idRule []interface{}
	for _, tranche_idItem := range tranche_id {
		tranche_idRule = append(tranche_idRule, tranche_idItem)
	}

	logs, sub, err := _ERC20Vesting.contract.WatchLogs(opts, "Tranche_Balance_Added", userRule, tranche_idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20VestingTrancheBalanceAdded)
				if err := _ERC20Vesting.contract.UnpackLog(event, "Tranche_Balance_Added", log); err != nil {
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

// ParseTrancheBalanceAdded is a log parse operation binding the contract event 0xd574b9b920ef951710b35bf90459120b667d6c810777f945e3e82996f75824c4.
//
// Solidity: event Tranche_Balance_Added(address indexed user, uint8 indexed tranche_id, uint256 amount)
func (_ERC20Vesting *ERC20VestingFilterer) ParseTrancheBalanceAdded(log types.Log) (*ERC20VestingTrancheBalanceAdded, error) {
	event := new(ERC20VestingTrancheBalanceAdded)
	if err := _ERC20Vesting.contract.UnpackLog(event, "Tranche_Balance_Added", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20VestingTrancheBalanceRemovedIterator is returned from FilterTrancheBalanceRemoved and is used to iterate over the raw logs and unpacked data for TrancheBalanceRemoved events raised by the ERC20Vesting contract.
type ERC20VestingTrancheBalanceRemovedIterator struct {
	Event *ERC20VestingTrancheBalanceRemoved // Event containing the contract specifics and raw log

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
func (it *ERC20VestingTrancheBalanceRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20VestingTrancheBalanceRemoved)
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
		it.Event = new(ERC20VestingTrancheBalanceRemoved)
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
func (it *ERC20VestingTrancheBalanceRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20VestingTrancheBalanceRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20VestingTrancheBalanceRemoved represents a TrancheBalanceRemoved event raised by the ERC20Vesting contract.
type ERC20VestingTrancheBalanceRemoved struct {
	User      common.Address
	TrancheId uint8
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTrancheBalanceRemoved is a free log retrieval operation binding the contract event 0x69ff77f8ae4662e6d47e9ec4c054fa45e4ed9cd4d875497ea162e3cb95ee2fd5.
//
// Solidity: event Tranche_Balance_Removed(address indexed user, uint8 indexed tranche_id, uint256 amount)
func (_ERC20Vesting *ERC20VestingFilterer) FilterTrancheBalanceRemoved(opts *bind.FilterOpts, user []common.Address, tranche_id []uint8) (*ERC20VestingTrancheBalanceRemovedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tranche_idRule []interface{}
	for _, tranche_idItem := range tranche_id {
		tranche_idRule = append(tranche_idRule, tranche_idItem)
	}

	logs, sub, err := _ERC20Vesting.contract.FilterLogs(opts, "Tranche_Balance_Removed", userRule, tranche_idRule)
	if err != nil {
		return nil, err
	}
	return &ERC20VestingTrancheBalanceRemovedIterator{contract: _ERC20Vesting.contract, event: "Tranche_Balance_Removed", logs: logs, sub: sub}, nil
}

// WatchTrancheBalanceRemoved is a free log subscription operation binding the contract event 0x69ff77f8ae4662e6d47e9ec4c054fa45e4ed9cd4d875497ea162e3cb95ee2fd5.
//
// Solidity: event Tranche_Balance_Removed(address indexed user, uint8 indexed tranche_id, uint256 amount)
func (_ERC20Vesting *ERC20VestingFilterer) WatchTrancheBalanceRemoved(opts *bind.WatchOpts, sink chan<- *ERC20VestingTrancheBalanceRemoved, user []common.Address, tranche_id []uint8) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tranche_idRule []interface{}
	for _, tranche_idItem := range tranche_id {
		tranche_idRule = append(tranche_idRule, tranche_idItem)
	}

	logs, sub, err := _ERC20Vesting.contract.WatchLogs(opts, "Tranche_Balance_Removed", userRule, tranche_idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20VestingTrancheBalanceRemoved)
				if err := _ERC20Vesting.contract.UnpackLog(event, "Tranche_Balance_Removed", log); err != nil {
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

// ParseTrancheBalanceRemoved is a log parse operation binding the contract event 0x69ff77f8ae4662e6d47e9ec4c054fa45e4ed9cd4d875497ea162e3cb95ee2fd5.
//
// Solidity: event Tranche_Balance_Removed(address indexed user, uint8 indexed tranche_id, uint256 amount)
func (_ERC20Vesting *ERC20VestingFilterer) ParseTrancheBalanceRemoved(log types.Log) (*ERC20VestingTrancheBalanceRemoved, error) {
	event := new(ERC20VestingTrancheBalanceRemoved)
	if err := _ERC20Vesting.contract.UnpackLog(event, "Tranche_Balance_Removed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20VestingTrancheCreatedIterator is returned from FilterTrancheCreated and is used to iterate over the raw logs and unpacked data for TrancheCreated events raised by the ERC20Vesting contract.
type ERC20VestingTrancheCreatedIterator struct {
	Event *ERC20VestingTrancheCreated // Event containing the contract specifics and raw log

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
func (it *ERC20VestingTrancheCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20VestingTrancheCreated)
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
		it.Event = new(ERC20VestingTrancheCreated)
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
func (it *ERC20VestingTrancheCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20VestingTrancheCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20VestingTrancheCreated represents a TrancheCreated event raised by the ERC20Vesting contract.
type ERC20VestingTrancheCreated struct {
	TrancheId  uint8
	CliffStart *big.Int
	Duration   *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTrancheCreated is a free log retrieval operation binding the contract event 0xc1aade57980687511ea5382b793f5b68cf97b354aef89addfbf9c59d23f1fc72.
//
// Solidity: event Tranche_Created(uint8 indexed tranche_id, uint256 cliff_start, uint256 duration)
func (_ERC20Vesting *ERC20VestingFilterer) FilterTrancheCreated(opts *bind.FilterOpts, tranche_id []uint8) (*ERC20VestingTrancheCreatedIterator, error) {

	var tranche_idRule []interface{}
	for _, tranche_idItem := range tranche_id {
		tranche_idRule = append(tranche_idRule, tranche_idItem)
	}

	logs, sub, err := _ERC20Vesting.contract.FilterLogs(opts, "Tranche_Created", tranche_idRule)
	if err != nil {
		return nil, err
	}
	return &ERC20VestingTrancheCreatedIterator{contract: _ERC20Vesting.contract, event: "Tranche_Created", logs: logs, sub: sub}, nil
}

// WatchTrancheCreated is a free log subscription operation binding the contract event 0xc1aade57980687511ea5382b793f5b68cf97b354aef89addfbf9c59d23f1fc72.
//
// Solidity: event Tranche_Created(uint8 indexed tranche_id, uint256 cliff_start, uint256 duration)
func (_ERC20Vesting *ERC20VestingFilterer) WatchTrancheCreated(opts *bind.WatchOpts, sink chan<- *ERC20VestingTrancheCreated, tranche_id []uint8) (event.Subscription, error) {

	var tranche_idRule []interface{}
	for _, tranche_idItem := range tranche_id {
		tranche_idRule = append(tranche_idRule, tranche_idItem)
	}

	logs, sub, err := _ERC20Vesting.contract.WatchLogs(opts, "Tranche_Created", tranche_idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20VestingTrancheCreated)
				if err := _ERC20Vesting.contract.UnpackLog(event, "Tranche_Created", log); err != nil {
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

// ParseTrancheCreated is a log parse operation binding the contract event 0xc1aade57980687511ea5382b793f5b68cf97b354aef89addfbf9c59d23f1fc72.
//
// Solidity: event Tranche_Created(uint8 indexed tranche_id, uint256 cliff_start, uint256 duration)
func (_ERC20Vesting *ERC20VestingFilterer) ParseTrancheCreated(log types.Log) (*ERC20VestingTrancheCreated, error) {
	event := new(ERC20VestingTrancheCreated)
	if err := _ERC20Vesting.contract.UnpackLog(event, "Tranche_Created", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
