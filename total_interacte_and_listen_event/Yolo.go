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
	_ = abi.ConvertType
)

// YoloVRFBet is an auto generated low-level Go binding around an user-defined struct.
type YoloVRFBet struct {
	Uid    uint64
	Amount uint64
}

// YoloMetaData contains all meta data concerning the Yolo contract.
var YoloMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dedicatedMsgSender\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"RequestedRandomness\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"uid\",\"type\":\"uint64\"}],\"name\":\"Win\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_round\",\"type\":\"uint64\"}],\"name\":\"betInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"uid\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"amount\",\"type\":\"uint64\"}],\"internalType\":\"structYoloVRF.Bet[]\",\"name\":\"bets\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_round\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"uid\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"amount\",\"type\":\"uint64\"}],\"internalType\":\"structYoloVRF.Bet[]\",\"name\":\"_bets\",\"type\":\"tuple[]\"}],\"name\":\"betting\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"callers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"randomness\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"dataWithRound\",\"type\":\"bytes\"}],\"name\":\"fulfillRandomness\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"infos\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"totalAmount\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"winnerId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"randomness\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastRound\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"requestPending\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"requestedHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"enable\",\"type\":\"bool\"}],\"name\":\"setCaller\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_round\",\"type\":\"uint64\"}],\"name\":\"winner\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"uid\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// YoloABI is the input ABI used to generate the binding from.
// Deprecated: Use YoloMetaData.ABI instead.
var YoloABI = YoloMetaData.ABI

// Yolo is an auto generated Go binding around an Ethereum contract.
type Yolo struct {
	YoloCaller     // Read-only binding to the contract
	YoloTransactor // Write-only binding to the contract
	YoloFilterer   // Log filterer for contract events
}

// YoloCaller is an auto generated read-only Go binding around an Ethereum contract.
type YoloCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YoloTransactor is an auto generated write-only Go binding around an Ethereum contract.
type YoloTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YoloFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type YoloFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YoloSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type YoloSession struct {
	Contract     *Yolo             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// YoloCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type YoloCallerSession struct {
	Contract *YoloCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// YoloTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type YoloTransactorSession struct {
	Contract     *YoloTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// YoloRaw is an auto generated low-level Go binding around an Ethereum contract.
type YoloRaw struct {
	Contract *Yolo // Generic contract binding to access the raw methods on
}

// YoloCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type YoloCallerRaw struct {
	Contract *YoloCaller // Generic read-only contract binding to access the raw methods on
}

// YoloTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type YoloTransactorRaw struct {
	Contract *YoloTransactor // Generic write-only contract binding to access the raw methods on
}

// NewYolo creates a new instance of Yolo, bound to a specific deployed contract.
func NewYolo(address common.Address, backend bind.ContractBackend) (*Yolo, error) {
	contract, err := bindYolo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Yolo{YoloCaller: YoloCaller{contract: contract}, YoloTransactor: YoloTransactor{contract: contract}, YoloFilterer: YoloFilterer{contract: contract}}, nil
}

// NewYoloCaller creates a new read-only instance of Yolo, bound to a specific deployed contract.
func NewYoloCaller(address common.Address, caller bind.ContractCaller) (*YoloCaller, error) {
	contract, err := bindYolo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &YoloCaller{contract: contract}, nil
}

// NewYoloTransactor creates a new write-only instance of Yolo, bound to a specific deployed contract.
func NewYoloTransactor(address common.Address, transactor bind.ContractTransactor) (*YoloTransactor, error) {
	contract, err := bindYolo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &YoloTransactor{contract: contract}, nil
}

// NewYoloFilterer creates a new log filterer instance of Yolo, bound to a specific deployed contract.
func NewYoloFilterer(address common.Address, filterer bind.ContractFilterer) (*YoloFilterer, error) {
	contract, err := bindYolo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &YoloFilterer{contract: contract}, nil
}

// bindYolo binds a generic wrapper to an already deployed contract.
func bindYolo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := YoloMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Yolo *YoloRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Yolo.Contract.YoloCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Yolo *YoloRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yolo.Contract.YoloTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Yolo *YoloRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Yolo.Contract.YoloTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Yolo *YoloCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Yolo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Yolo *YoloTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yolo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Yolo *YoloTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Yolo.Contract.contract.Transact(opts, method, params...)
}

// BetInfo is a free data retrieval call binding the contract method 0x4db4c98d.
//
// Solidity: function betInfo(uint64 _round) view returns((uint64,uint64)[] bets)
func (_Yolo *YoloCaller) BetInfo(opts *bind.CallOpts, _round uint64) ([]YoloVRFBet, error) {
	var out []interface{}
	err := _Yolo.contract.Call(opts, &out, "betInfo", _round)

	if err != nil {
		return *new([]YoloVRFBet), err
	}

	out0 := *abi.ConvertType(out[0], new([]YoloVRFBet)).(*[]YoloVRFBet)

	return out0, err

}

// BetInfo is a free data retrieval call binding the contract method 0x4db4c98d.
//
// Solidity: function betInfo(uint64 _round) view returns((uint64,uint64)[] bets)
func (_Yolo *YoloSession) BetInfo(_round uint64) ([]YoloVRFBet, error) {
	return _Yolo.Contract.BetInfo(&_Yolo.CallOpts, _round)
}

// BetInfo is a free data retrieval call binding the contract method 0x4db4c98d.
//
// Solidity: function betInfo(uint64 _round) view returns((uint64,uint64)[] bets)
func (_Yolo *YoloCallerSession) BetInfo(_round uint64) ([]YoloVRFBet, error) {
	return _Yolo.Contract.BetInfo(&_Yolo.CallOpts, _round)
}

// Callers is a free data retrieval call binding the contract method 0x7bbf4a3f.
//
// Solidity: function callers(address ) view returns(bool)
func (_Yolo *YoloCaller) Callers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Yolo.contract.Call(opts, &out, "callers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Callers is a free data retrieval call binding the contract method 0x7bbf4a3f.
//
// Solidity: function callers(address ) view returns(bool)
func (_Yolo *YoloSession) Callers(arg0 common.Address) (bool, error) {
	return _Yolo.Contract.Callers(&_Yolo.CallOpts, arg0)
}

// Callers is a free data retrieval call binding the contract method 0x7bbf4a3f.
//
// Solidity: function callers(address ) view returns(bool)
func (_Yolo *YoloCallerSession) Callers(arg0 common.Address) (bool, error) {
	return _Yolo.Contract.Callers(&_Yolo.CallOpts, arg0)
}

// Infos is a free data retrieval call binding the contract method 0x1d7771de.
//
// Solidity: function infos(uint64 ) view returns(uint64 totalAmount, uint64 winnerId, uint256 randomness)
func (_Yolo *YoloCaller) Infos(opts *bind.CallOpts, arg0 uint64) (struct {
	TotalAmount uint64
	WinnerId    uint64
	Randomness  *big.Int
}, error) {
	var out []interface{}
	err := _Yolo.contract.Call(opts, &out, "infos", arg0)

	outstruct := new(struct {
		TotalAmount uint64
		WinnerId    uint64
		Randomness  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalAmount = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.WinnerId = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.Randomness = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Infos is a free data retrieval call binding the contract method 0x1d7771de.
//
// Solidity: function infos(uint64 ) view returns(uint64 totalAmount, uint64 winnerId, uint256 randomness)
func (_Yolo *YoloSession) Infos(arg0 uint64) (struct {
	TotalAmount uint64
	WinnerId    uint64
	Randomness  *big.Int
}, error) {
	return _Yolo.Contract.Infos(&_Yolo.CallOpts, arg0)
}

// Infos is a free data retrieval call binding the contract method 0x1d7771de.
//
// Solidity: function infos(uint64 ) view returns(uint64 totalAmount, uint64 winnerId, uint256 randomness)
func (_Yolo *YoloCallerSession) Infos(arg0 uint64) (struct {
	TotalAmount uint64
	WinnerId    uint64
	Randomness  *big.Int
}, error) {
	return _Yolo.Contract.Infos(&_Yolo.CallOpts, arg0)
}

// LastRound is a free data retrieval call binding the contract method 0x82bc07e6.
//
// Solidity: function lastRound() view returns(uint64)
func (_Yolo *YoloCaller) LastRound(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Yolo.contract.Call(opts, &out, "lastRound")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastRound is a free data retrieval call binding the contract method 0x82bc07e6.
//
// Solidity: function lastRound() view returns(uint64)
func (_Yolo *YoloSession) LastRound() (uint64, error) {
	return _Yolo.Contract.LastRound(&_Yolo.CallOpts)
}

// LastRound is a free data retrieval call binding the contract method 0x82bc07e6.
//
// Solidity: function lastRound() view returns(uint64)
func (_Yolo *YoloCallerSession) LastRound() (uint64, error) {
	return _Yolo.Contract.LastRound(&_Yolo.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Yolo *YoloCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Yolo.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Yolo *YoloSession) Owner() (common.Address, error) {
	return _Yolo.Contract.Owner(&_Yolo.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Yolo *YoloCallerSession) Owner() (common.Address, error) {
	return _Yolo.Contract.Owner(&_Yolo.CallOpts)
}

// RequestPending is a free data retrieval call binding the contract method 0x75ce7fff.
//
// Solidity: function requestPending(uint256 ) view returns(bool)
func (_Yolo *YoloCaller) RequestPending(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _Yolo.contract.Call(opts, &out, "requestPending", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RequestPending is a free data retrieval call binding the contract method 0x75ce7fff.
//
// Solidity: function requestPending(uint256 ) view returns(bool)
func (_Yolo *YoloSession) RequestPending(arg0 *big.Int) (bool, error) {
	return _Yolo.Contract.RequestPending(&_Yolo.CallOpts, arg0)
}

// RequestPending is a free data retrieval call binding the contract method 0x75ce7fff.
//
// Solidity: function requestPending(uint256 ) view returns(bool)
func (_Yolo *YoloCallerSession) RequestPending(arg0 *big.Int) (bool, error) {
	return _Yolo.Contract.RequestPending(&_Yolo.CallOpts, arg0)
}

// RequestedHash is a free data retrieval call binding the contract method 0xc4f8f27b.
//
// Solidity: function requestedHash(uint256 ) view returns(bytes32)
func (_Yolo *YoloCaller) RequestedHash(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Yolo.contract.Call(opts, &out, "requestedHash", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RequestedHash is a free data retrieval call binding the contract method 0xc4f8f27b.
//
// Solidity: function requestedHash(uint256 ) view returns(bytes32)
func (_Yolo *YoloSession) RequestedHash(arg0 *big.Int) ([32]byte, error) {
	return _Yolo.Contract.RequestedHash(&_Yolo.CallOpts, arg0)
}

// RequestedHash is a free data retrieval call binding the contract method 0xc4f8f27b.
//
// Solidity: function requestedHash(uint256 ) view returns(bytes32)
func (_Yolo *YoloCallerSession) RequestedHash(arg0 *big.Int) ([32]byte, error) {
	return _Yolo.Contract.RequestedHash(&_Yolo.CallOpts, arg0)
}

// Winner is a free data retrieval call binding the contract method 0x6a131136.
//
// Solidity: function winner(uint64 _round) view returns(uint64 uid)
func (_Yolo *YoloCaller) Winner(opts *bind.CallOpts, _round uint64) (uint64, error) {
	var out []interface{}
	err := _Yolo.contract.Call(opts, &out, "winner", _round)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Winner is a free data retrieval call binding the contract method 0x6a131136.
//
// Solidity: function winner(uint64 _round) view returns(uint64 uid)
func (_Yolo *YoloSession) Winner(_round uint64) (uint64, error) {
	return _Yolo.Contract.Winner(&_Yolo.CallOpts, _round)
}

// Winner is a free data retrieval call binding the contract method 0x6a131136.
//
// Solidity: function winner(uint64 _round) view returns(uint64 uid)
func (_Yolo *YoloCallerSession) Winner(_round uint64) (uint64, error) {
	return _Yolo.Contract.Winner(&_Yolo.CallOpts, _round)
}

// Betting is a paid mutator transaction binding the contract method 0xde85d7b8.
//
// Solidity: function betting(uint64 _round, (uint64,uint64)[] _bets) returns()
func (_Yolo *YoloTransactor) Betting(opts *bind.TransactOpts, _round uint64, _bets []YoloVRFBet) (*types.Transaction, error) {
	return _Yolo.contract.Transact(opts, "betting", _round, _bets)
}

// Betting is a paid mutator transaction binding the contract method 0xde85d7b8.
//
// Solidity: function betting(uint64 _round, (uint64,uint64)[] _bets) returns()
func (_Yolo *YoloSession) Betting(_round uint64, _bets []YoloVRFBet) (*types.Transaction, error) {
	return _Yolo.Contract.Betting(&_Yolo.TransactOpts, _round, _bets)
}

// Betting is a paid mutator transaction binding the contract method 0xde85d7b8.
//
// Solidity: function betting(uint64 _round, (uint64,uint64)[] _bets) returns()
func (_Yolo *YoloTransactorSession) Betting(_round uint64, _bets []YoloVRFBet) (*types.Transaction, error) {
	return _Yolo.Contract.Betting(&_Yolo.TransactOpts, _round, _bets)
}

// FulfillRandomness is a paid mutator transaction binding the contract method 0xb3f6b99a.
//
// Solidity: function fulfillRandomness(uint256 randomness, bytes dataWithRound) returns()
func (_Yolo *YoloTransactor) FulfillRandomness(opts *bind.TransactOpts, randomness *big.Int, dataWithRound []byte) (*types.Transaction, error) {
	return _Yolo.contract.Transact(opts, "fulfillRandomness", randomness, dataWithRound)
}

// FulfillRandomness is a paid mutator transaction binding the contract method 0xb3f6b99a.
//
// Solidity: function fulfillRandomness(uint256 randomness, bytes dataWithRound) returns()
func (_Yolo *YoloSession) FulfillRandomness(randomness *big.Int, dataWithRound []byte) (*types.Transaction, error) {
	return _Yolo.Contract.FulfillRandomness(&_Yolo.TransactOpts, randomness, dataWithRound)
}

// FulfillRandomness is a paid mutator transaction binding the contract method 0xb3f6b99a.
//
// Solidity: function fulfillRandomness(uint256 randomness, bytes dataWithRound) returns()
func (_Yolo *YoloTransactorSession) FulfillRandomness(randomness *big.Int, dataWithRound []byte) (*types.Transaction, error) {
	return _Yolo.Contract.FulfillRandomness(&_Yolo.TransactOpts, randomness, dataWithRound)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Yolo *YoloTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yolo.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Yolo *YoloSession) RenounceOwnership() (*types.Transaction, error) {
	return _Yolo.Contract.RenounceOwnership(&_Yolo.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Yolo *YoloTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Yolo.Contract.RenounceOwnership(&_Yolo.TransactOpts)
}

// SetCaller is a paid mutator transaction binding the contract method 0x9cae6eae.
//
// Solidity: function setCaller(address addr, bool enable) returns()
func (_Yolo *YoloTransactor) SetCaller(opts *bind.TransactOpts, addr common.Address, enable bool) (*types.Transaction, error) {
	return _Yolo.contract.Transact(opts, "setCaller", addr, enable)
}

// SetCaller is a paid mutator transaction binding the contract method 0x9cae6eae.
//
// Solidity: function setCaller(address addr, bool enable) returns()
func (_Yolo *YoloSession) SetCaller(addr common.Address, enable bool) (*types.Transaction, error) {
	return _Yolo.Contract.SetCaller(&_Yolo.TransactOpts, addr, enable)
}

// SetCaller is a paid mutator transaction binding the contract method 0x9cae6eae.
//
// Solidity: function setCaller(address addr, bool enable) returns()
func (_Yolo *YoloTransactorSession) SetCaller(addr common.Address, enable bool) (*types.Transaction, error) {
	return _Yolo.Contract.SetCaller(&_Yolo.TransactOpts, addr, enable)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Yolo *YoloTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Yolo.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Yolo *YoloSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Yolo.Contract.TransferOwnership(&_Yolo.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Yolo *YoloTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Yolo.Contract.TransferOwnership(&_Yolo.TransactOpts, newOwner)
}

// YoloOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Yolo contract.
type YoloOwnershipTransferredIterator struct {
	Event *YoloOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *YoloOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YoloOwnershipTransferred)
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
		it.Event = new(YoloOwnershipTransferred)
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
func (it *YoloOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YoloOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YoloOwnershipTransferred represents a OwnershipTransferred event raised by the Yolo contract.
type YoloOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Yolo *YoloFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*YoloOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Yolo.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &YoloOwnershipTransferredIterator{contract: _Yolo.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Yolo *YoloFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *YoloOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Yolo.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YoloOwnershipTransferred)
				if err := _Yolo.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Yolo *YoloFilterer) ParseOwnershipTransferred(log types.Log) (*YoloOwnershipTransferred, error) {
	event := new(YoloOwnershipTransferred)
	if err := _Yolo.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YoloRequestedRandomnessIterator is returned from FilterRequestedRandomness and is used to iterate over the raw logs and unpacked data for RequestedRandomness events raised by the Yolo contract.
type YoloRequestedRandomnessIterator struct {
	Event *YoloRequestedRandomness // Event containing the contract specifics and raw log

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
func (it *YoloRequestedRandomnessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YoloRequestedRandomness)
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
		it.Event = new(YoloRequestedRandomness)
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
func (it *YoloRequestedRandomnessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YoloRequestedRandomnessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YoloRequestedRandomness represents a RequestedRandomness event raised by the Yolo contract.
type YoloRequestedRandomness struct {
	Round *big.Int
	Data  []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRequestedRandomness is a free log retrieval operation binding the contract event 0xd91fc3685b930310b008ec37d2334870cab88a023ed8cc628a2e2ccd4e55d202.
//
// Solidity: event RequestedRandomness(uint256 round, bytes data)
func (_Yolo *YoloFilterer) FilterRequestedRandomness(opts *bind.FilterOpts) (*YoloRequestedRandomnessIterator, error) {

	logs, sub, err := _Yolo.contract.FilterLogs(opts, "RequestedRandomness")
	if err != nil {
		return nil, err
	}
	return &YoloRequestedRandomnessIterator{contract: _Yolo.contract, event: "RequestedRandomness", logs: logs, sub: sub}, nil
}

// WatchRequestedRandomness is a free log subscription operation binding the contract event 0xd91fc3685b930310b008ec37d2334870cab88a023ed8cc628a2e2ccd4e55d202.
//
// Solidity: event RequestedRandomness(uint256 round, bytes data)
func (_Yolo *YoloFilterer) WatchRequestedRandomness(opts *bind.WatchOpts, sink chan<- *YoloRequestedRandomness) (event.Subscription, error) {

	logs, sub, err := _Yolo.contract.WatchLogs(opts, "RequestedRandomness")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YoloRequestedRandomness)
				if err := _Yolo.contract.UnpackLog(event, "RequestedRandomness", log); err != nil {
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

// ParseRequestedRandomness is a log parse operation binding the contract event 0xd91fc3685b930310b008ec37d2334870cab88a023ed8cc628a2e2ccd4e55d202.
//
// Solidity: event RequestedRandomness(uint256 round, bytes data)
func (_Yolo *YoloFilterer) ParseRequestedRandomness(log types.Log) (*YoloRequestedRandomness, error) {
	event := new(YoloRequestedRandomness)
	if err := _Yolo.contract.UnpackLog(event, "RequestedRandomness", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YoloWinIterator is returned from FilterWin and is used to iterate over the raw logs and unpacked data for Win events raised by the Yolo contract.
type YoloWinIterator struct {
	Event *YoloWin // Event containing the contract specifics and raw log

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
func (it *YoloWinIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YoloWin)
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
		it.Event = new(YoloWin)
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
func (it *YoloWinIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YoloWinIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YoloWin represents a Win event raised by the Yolo contract.
type YoloWin struct {
	Round uint64
	Uid   uint64
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterWin is a free log retrieval operation binding the contract event 0xd80a925b1b3de17571108034e9364dbf148bcda93c9f30a22f127c73e19e4d51.
//
// Solidity: event Win(uint64 indexed round, uint64 indexed uid)
func (_Yolo *YoloFilterer) FilterWin(opts *bind.FilterOpts, round []uint64, uid []uint64) (*YoloWinIterator, error) {

	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}
	var uidRule []interface{}
	for _, uidItem := range uid {
		uidRule = append(uidRule, uidItem)
	}

	logs, sub, err := _Yolo.contract.FilterLogs(opts, "Win", roundRule, uidRule)
	if err != nil {
		return nil, err
	}
	return &YoloWinIterator{contract: _Yolo.contract, event: "Win", logs: logs, sub: sub}, nil
}

// WatchWin is a free log subscription operation binding the contract event 0xd80a925b1b3de17571108034e9364dbf148bcda93c9f30a22f127c73e19e4d51.
//
// Solidity: event Win(uint64 indexed round, uint64 indexed uid)
func (_Yolo *YoloFilterer) WatchWin(opts *bind.WatchOpts, sink chan<- *YoloWin, round []uint64, uid []uint64) (event.Subscription, error) {

	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}
	var uidRule []interface{}
	for _, uidItem := range uid {
		uidRule = append(uidRule, uidItem)
	}

	logs, sub, err := _Yolo.contract.WatchLogs(opts, "Win", roundRule, uidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YoloWin)
				if err := _Yolo.contract.UnpackLog(event, "Win", log); err != nil {
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

// ParseWin is a log parse operation binding the contract event 0xd80a925b1b3de17571108034e9364dbf148bcda93c9f30a22f127c73e19e4d51.
//
// Solidity: event Win(uint64 indexed round, uint64 indexed uid)
func (_Yolo *YoloFilterer) ParseWin(log types.Log) (*YoloWin, error) {
	event := new(YoloWin)
	if err := _Yolo.contract.UnpackLog(event, "Win", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
