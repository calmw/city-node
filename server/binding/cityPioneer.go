// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package intoCityNode

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

// CityPioneerMetaData contains all meta data concerning the CityPioneer contract.
var CityPioneerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AdminAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AdminRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pioneerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"cityId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ctime\",\"type\":\"uint256\"}],\"name\":\"DailyIncreaseDelegateRecord\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pioneerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toxReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"foundsReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delegateReward\",\"type\":\"uint256\"}],\"name\":\"DailyRewardRecord\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"cityId\",\"type\":\"bytes32\"}],\"name\":\"LogBytes32\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LogUint256\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pioneerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SuretyRecord\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pioneerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardType\",\"type\":\"uint256\"}],\"name\":\"WithdrawalRewardRecord\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"TOXAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cityLevel_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"month_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"point_\",\"type\":\"uint256\"}],\"name\":\"adminSetAssessmentCriteria\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cityAddress_\",\"type\":\"address\"}],\"name\":\"adminSetCityAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"miningAddress_\",\"type\":\"address\"}],\"name\":\"adminSetMiningAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"TOXAddress_\",\"type\":\"address\"}],\"name\":\"adminSetTOXAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allAdmins\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"admins\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"assessmentCriteria\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"assessmentReturnCriteria\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"assessmentReturnRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"amounts\",\"type\":\"address[]\"}],\"name\":\"batchAddAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"benefitPackageReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"benefitPackageRewardReceived\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"benefitPackageRewardStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cityAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"dailyTaskStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"delegateReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"delegateRewardReceived\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"delegateRewardStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositSurety\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pioneerTest_\",\"type\":\"address\"}],\"name\":\"depositSuretyTest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"fundsReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"fundsRewardReceived\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"fundsRewardStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSurety\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"isPioneer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"miningAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pioneerInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pioneerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"ctime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cityLevel\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"assessmentMonthStatus\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"assessmentStatus\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"returnSuretyStatus\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"returnSuretyRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"returnSuretyTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pioneerAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"cityId_\",\"type\":\"bytes32\"}],\"name\":\"setPioneerDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"suretyReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"suretyRewardRecord\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"withdrawalBenefitPackageReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"withdrawalDelegateRewardReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"withdrawalFundsRewardReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"withdrawalSuretyReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// CityPioneerABI is the input ABI used to generate the binding from.
// Deprecated: Use CityPioneerMetaData.ABI instead.
var CityPioneerABI = CityPioneerMetaData.ABI

// CityPioneer is an auto generated Go binding around an Ethereum contract.
type CityPioneer struct {
	CityPioneerCaller     // Read-only binding to the contract
	CityPioneerTransactor // Write-only binding to the contract
	CityPioneerFilterer   // Log filterer for contract events
}

// CityPioneerCaller is an auto generated read-only Go binding around an Ethereum contract.
type CityPioneerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CityPioneerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CityPioneerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CityPioneerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CityPioneerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CityPioneerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CityPioneerSession struct {
	Contract     *CityPioneer      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CityPioneerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CityPioneerCallerSession struct {
	Contract *CityPioneerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// CityPioneerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CityPioneerTransactorSession struct {
	Contract     *CityPioneerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// CityPioneerRaw is an auto generated low-level Go binding around an Ethereum contract.
type CityPioneerRaw struct {
	Contract *CityPioneer // Generic contract binding to access the raw methods on
}

// CityPioneerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CityPioneerCallerRaw struct {
	Contract *CityPioneerCaller // Generic read-only contract binding to access the raw methods on
}

// CityPioneerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CityPioneerTransactorRaw struct {
	Contract *CityPioneerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCityPioneer creates a new instance of CityPioneer, bound to a specific deployed contract.
func NewCityPioneer(address common.Address, backend bind.ContractBackend) (*CityPioneer, error) {
	contract, err := bindCityPioneer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CityPioneer{CityPioneerCaller: CityPioneerCaller{contract: contract}, CityPioneerTransactor: CityPioneerTransactor{contract: contract}, CityPioneerFilterer: CityPioneerFilterer{contract: contract}}, nil
}

// NewCityPioneerCaller creates a new read-only instance of CityPioneer, bound to a specific deployed contract.
func NewCityPioneerCaller(address common.Address, caller bind.ContractCaller) (*CityPioneerCaller, error) {
	contract, err := bindCityPioneer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CityPioneerCaller{contract: contract}, nil
}

// NewCityPioneerTransactor creates a new write-only instance of CityPioneer, bound to a specific deployed contract.
func NewCityPioneerTransactor(address common.Address, transactor bind.ContractTransactor) (*CityPioneerTransactor, error) {
	contract, err := bindCityPioneer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CityPioneerTransactor{contract: contract}, nil
}

// NewCityPioneerFilterer creates a new log filterer instance of CityPioneer, bound to a specific deployed contract.
func NewCityPioneerFilterer(address common.Address, filterer bind.ContractFilterer) (*CityPioneerFilterer, error) {
	contract, err := bindCityPioneer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CityPioneerFilterer{contract: contract}, nil
}

// bindCityPioneer binds a generic wrapper to an already deployed contract.
func bindCityPioneer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CityPioneerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CityPioneer *CityPioneerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CityPioneer.Contract.CityPioneerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CityPioneer *CityPioneerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CityPioneer.Contract.CityPioneerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CityPioneer *CityPioneerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CityPioneer.Contract.CityPioneerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CityPioneer *CityPioneerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CityPioneer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CityPioneer *CityPioneerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CityPioneer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CityPioneer *CityPioneerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CityPioneer.Contract.contract.Transact(opts, method, params...)
}

// TOXAddress is a free data retrieval call binding the contract method 0xccff8899.
//
// Solidity: function TOXAddress() view returns(address)
func (_CityPioneer *CityPioneerCaller) TOXAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CityPioneer.contract.Call(opts, &out, "TOXAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TOXAddress is a free data retrieval call binding the contract method 0xccff8899.
//
// Solidity: function TOXAddress() view returns(address)
func (_CityPioneer *CityPioneerSession) TOXAddress() (common.Address, error) {
	return _CityPioneer.Contract.TOXAddress(&_CityPioneer.CallOpts)
}

// TOXAddress is a free data retrieval call binding the contract method 0xccff8899.
//
// Solidity: function TOXAddress() view returns(address)
func (_CityPioneer *CityPioneerCallerSession) TOXAddress() (common.Address, error) {
	return _CityPioneer.Contract.TOXAddress(&_CityPioneer.CallOpts)
}

// AllAdmins is a free data retrieval call binding the contract method 0x40f32be6.
//
// Solidity: function allAdmins() view returns(address[] admins)
func (_CityPioneer *CityPioneerCaller) AllAdmins(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _CityPioneer.contract.Call(opts, &out, "allAdmins")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// AllAdmins is a free data retrieval call binding the contract method 0x40f32be6.
//
// Solidity: function allAdmins() view returns(address[] admins)
func (_CityPioneer *CityPioneerSession) AllAdmins() ([]common.Address, error) {
	return _CityPioneer.Contract.AllAdmins(&_CityPioneer.CallOpts)
}

// AllAdmins is a free data retrieval call binding the contract method 0x40f32be6.
//
// Solidity: function allAdmins() view returns(address[] admins)
func (_CityPioneer *CityPioneerCallerSession) AllAdmins() ([]common.Address, error) {
	return _CityPioneer.Contract.AllAdmins(&_CityPioneer.CallOpts)
}

// AssessmentCriteria is a free data retrieval call binding the contract method 0xc1e3d40b.
//
// Solidity: function assessmentCriteria(uint256 , uint256 ) view returns(uint256)
func (_CityPioneer *CityPioneerCaller) AssessmentCriteria(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CityPioneer.contract.Call(opts, &out, "assessmentCriteria", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AssessmentCriteria is a free data retrieval call binding the contract method 0xc1e3d40b.
//
// Solidity: function assessmentCriteria(uint256 , uint256 ) view returns(uint256)
func (_CityPioneer *CityPioneerSession) AssessmentCriteria(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _CityPioneer.Contract.AssessmentCriteria(&_CityPioneer.CallOpts, arg0, arg1)
}

// AssessmentCriteria is a free data retrieval call binding the contract method 0xc1e3d40b.
//
// Solidity: function assessmentCriteria(uint256 , uint256 ) view returns(uint256)
func (_CityPioneer *CityPioneerCallerSession) AssessmentCriteria(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _CityPioneer.Contract.AssessmentCriteria(&_CityPioneer.CallOpts, arg0, arg1)
}

// AssessmentReturnCriteria is a free data retrieval call binding the contract method 0x55fd6ce1.
//
// Solidity: function assessmentReturnCriteria(uint256 , uint256 ) view returns(uint256)
func (_CityPioneer *CityPioneerCaller) AssessmentReturnCriteria(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CityPioneer.contract.Call(opts, &out, "assessmentReturnCriteria", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AssessmentReturnCriteria is a free data retrieval call binding the contract method 0x55fd6ce1.
//
// Solidity: function assessmentReturnCriteria(uint256 , uint256 ) view returns(uint256)
func (_CityPioneer *CityPioneerSession) AssessmentReturnCriteria(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _CityPioneer.Contract.AssessmentReturnCriteria(&_CityPioneer.CallOpts, arg0, arg1)
}

// AssessmentReturnCriteria is a free data retrieval call binding the contract method 0x55fd6ce1.
//
// Solidity: function assessmentReturnCriteria(uint256 , uint256 ) view returns(uint256)
func (_CityPioneer *CityPioneerCallerSession) AssessmentReturnCriteria(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _CityPioneer.Contract.AssessmentReturnCriteria(&_CityPioneer.CallOpts, arg0, arg1)
}

// AssessmentReturnRate is a free data retrieval call binding the contract method 0x71753070.
//
// Solidity: function assessmentReturnRate(uint256 , uint256 ) view returns(uint256)
func (_CityPioneer *CityPioneerCaller) AssessmentReturnRate(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CityPioneer.contract.Call(opts, &out, "assessmentReturnRate", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AssessmentReturnRate is a free data retrieval call binding the contract method 0x71753070.
//
// Solidity: function assessmentReturnRate(uint256 , uint256 ) view returns(uint256)
func (_CityPioneer *CityPioneerSession) AssessmentReturnRate(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _CityPioneer.Contract.AssessmentReturnRate(&_CityPioneer.CallOpts, arg0, arg1)
}

// AssessmentReturnRate is a free data retrieval call binding the contract method 0x71753070.
//
// Solidity: function assessmentReturnRate(uint256 , uint256 ) view returns(uint256)
func (_CityPioneer *CityPioneerCallerSession) AssessmentReturnRate(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _CityPioneer.Contract.AssessmentReturnRate(&_CityPioneer.CallOpts, arg0, arg1)
}

// BenefitPackageReward is a free data retrieval call binding the contract method 0x7c1412f5.
//
// Solidity: function benefitPackageReward(address ) view returns(uint256)
func (_CityPioneer *CityPioneerCaller) BenefitPackageReward(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CityPioneer.contract.Call(opts, &out, "benefitPackageReward", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BenefitPackageReward is a free data retrieval call binding the contract method 0x7c1412f5.
//
// Solidity: function benefitPackageReward(address ) view returns(uint256)
func (_CityPioneer *CityPioneerSession) BenefitPackageReward(arg0 common.Address) (*big.Int, error) {
	return _CityPioneer.Contract.BenefitPackageReward(&_CityPioneer.CallOpts, arg0)
}

// BenefitPackageReward is a free data retrieval call binding the contract method 0x7c1412f5.
//
// Solidity: function benefitPackageReward(address ) view returns(uint256)
func (_CityPioneer *CityPioneerCallerSession) BenefitPackageReward(arg0 common.Address) (*big.Int, error) {
	return _CityPioneer.Contract.BenefitPackageReward(&_CityPioneer.CallOpts, arg0)
}

// BenefitPackageRewardReceived is a free data retrieval call binding the contract method 0xcc9ce5ae.
//
// Solidity: function benefitPackageRewardReceived(address ) view returns(uint256)
func (_CityPioneer *CityPioneerCaller) BenefitPackageRewardReceived(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CityPioneer.contract.Call(opts, &out, "benefitPackageRewardReceived", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BenefitPackageRewardReceived is a free data retrieval call binding the contract method 0xcc9ce5ae.
//
// Solidity: function benefitPackageRewardReceived(address ) view returns(uint256)
func (_CityPioneer *CityPioneerSession) BenefitPackageRewardReceived(arg0 common.Address) (*big.Int, error) {
	return _CityPioneer.Contract.BenefitPackageRewardReceived(&_CityPioneer.CallOpts, arg0)
}

// BenefitPackageRewardReceived is a free data retrieval call binding the contract method 0xcc9ce5ae.
//
// Solidity: function benefitPackageRewardReceived(address ) view returns(uint256)
func (_CityPioneer *CityPioneerCallerSession) BenefitPackageRewardReceived(arg0 common.Address) (*big.Int, error) {
	return _CityPioneer.Contract.BenefitPackageRewardReceived(&_CityPioneer.CallOpts, arg0)
}

// BenefitPackageRewardStatus is a free data retrieval call binding the contract method 0x5c1a0cfb.
//
// Solidity: function benefitPackageRewardStatus(address ) view returns(bool)
func (_CityPioneer *CityPioneerCaller) BenefitPackageRewardStatus(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _CityPioneer.contract.Call(opts, &out, "benefitPackageRewardStatus", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BenefitPackageRewardStatus is a free data retrieval call binding the contract method 0x5c1a0cfb.
//
// Solidity: function benefitPackageRewardStatus(address ) view returns(bool)
func (_CityPioneer *CityPioneerSession) BenefitPackageRewardStatus(arg0 common.Address) (bool, error) {
	return _CityPioneer.Contract.BenefitPackageRewardStatus(&_CityPioneer.CallOpts, arg0)
}

// BenefitPackageRewardStatus is a free data retrieval call binding the contract method 0x5c1a0cfb.
//
// Solidity: function benefitPackageRewardStatus(address ) view returns(bool)
func (_CityPioneer *CityPioneerCallerSession) BenefitPackageRewardStatus(arg0 common.Address) (bool, error) {
	return _CityPioneer.Contract.BenefitPackageRewardStatus(&_CityPioneer.CallOpts, arg0)
}

// CityAddress is a free data retrieval call binding the contract method 0x3e9de606.
//
// Solidity: function cityAddress() view returns(address)
func (_CityPioneer *CityPioneerCaller) CityAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CityPioneer.contract.Call(opts, &out, "cityAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CityAddress is a free data retrieval call binding the contract method 0x3e9de606.
//
// Solidity: function cityAddress() view returns(address)
func (_CityPioneer *CityPioneerSession) CityAddress() (common.Address, error) {
	return _CityPioneer.Contract.CityAddress(&_CityPioneer.CallOpts)
}

// CityAddress is a free data retrieval call binding the contract method 0x3e9de606.
//
// Solidity: function cityAddress() view returns(address)
func (_CityPioneer *CityPioneerCallerSession) CityAddress() (common.Address, error) {
	return _CityPioneer.Contract.CityAddress(&_CityPioneer.CallOpts)
}

// DailyTaskStatus is a free data retrieval call binding the contract method 0x0c069666.
//
// Solidity: function dailyTaskStatus(uint256 ) view returns(bool)
func (_CityPioneer *CityPioneerCaller) DailyTaskStatus(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _CityPioneer.contract.Call(opts, &out, "dailyTaskStatus", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DailyTaskStatus is a free data retrieval call binding the contract method 0x0c069666.
//
// Solidity: function dailyTaskStatus(uint256 ) view returns(bool)
func (_CityPioneer *CityPioneerSession) DailyTaskStatus(arg0 *big.Int) (bool, error) {
	return _CityPioneer.Contract.DailyTaskStatus(&_CityPioneer.CallOpts, arg0)
}

// DailyTaskStatus is a free data retrieval call binding the contract method 0x0c069666.
//
// Solidity: function dailyTaskStatus(uint256 ) view returns(bool)
func (_CityPioneer *CityPioneerCallerSession) DailyTaskStatus(arg0 *big.Int) (bool, error) {
	return _CityPioneer.Contract.DailyTaskStatus(&_CityPioneer.CallOpts, arg0)
}

// DelegateReward is a free data retrieval call binding the contract method 0xa9db12e0.
//
// Solidity: function delegateReward(address ) view returns(uint256)
func (_CityPioneer *CityPioneerCaller) DelegateReward(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CityPioneer.contract.Call(opts, &out, "delegateReward", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelegateReward is a free data retrieval call binding the contract method 0xa9db12e0.
//
// Solidity: function delegateReward(address ) view returns(uint256)
func (_CityPioneer *CityPioneerSession) DelegateReward(arg0 common.Address) (*big.Int, error) {
	return _CityPioneer.Contract.DelegateReward(&_CityPioneer.CallOpts, arg0)
}

// DelegateReward is a free data retrieval call binding the contract method 0xa9db12e0.
//
// Solidity: function delegateReward(address ) view returns(uint256)
func (_CityPioneer *CityPioneerCallerSession) DelegateReward(arg0 common.Address) (*big.Int, error) {
	return _CityPioneer.Contract.DelegateReward(&_CityPioneer.CallOpts, arg0)
}

// DelegateRewardReceived is a free data retrieval call binding the contract method 0xda6f62ef.
//
// Solidity: function delegateRewardReceived(address ) view returns(uint256)
func (_CityPioneer *CityPioneerCaller) DelegateRewardReceived(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CityPioneer.contract.Call(opts, &out, "delegateRewardReceived", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelegateRewardReceived is a free data retrieval call binding the contract method 0xda6f62ef.
//
// Solidity: function delegateRewardReceived(address ) view returns(uint256)
func (_CityPioneer *CityPioneerSession) DelegateRewardReceived(arg0 common.Address) (*big.Int, error) {
	return _CityPioneer.Contract.DelegateRewardReceived(&_CityPioneer.CallOpts, arg0)
}

// DelegateRewardReceived is a free data retrieval call binding the contract method 0xda6f62ef.
//
// Solidity: function delegateRewardReceived(address ) view returns(uint256)
func (_CityPioneer *CityPioneerCallerSession) DelegateRewardReceived(arg0 common.Address) (*big.Int, error) {
	return _CityPioneer.Contract.DelegateRewardReceived(&_CityPioneer.CallOpts, arg0)
}

// DelegateRewardStatus is a free data retrieval call binding the contract method 0xb019e616.
//
// Solidity: function delegateRewardStatus(address ) view returns(bool)
func (_CityPioneer *CityPioneerCaller) DelegateRewardStatus(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _CityPioneer.contract.Call(opts, &out, "delegateRewardStatus", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DelegateRewardStatus is a free data retrieval call binding the contract method 0xb019e616.
//
// Solidity: function delegateRewardStatus(address ) view returns(bool)
func (_CityPioneer *CityPioneerSession) DelegateRewardStatus(arg0 common.Address) (bool, error) {
	return _CityPioneer.Contract.DelegateRewardStatus(&_CityPioneer.CallOpts, arg0)
}

// DelegateRewardStatus is a free data retrieval call binding the contract method 0xb019e616.
//
// Solidity: function delegateRewardStatus(address ) view returns(bool)
func (_CityPioneer *CityPioneerCallerSession) DelegateRewardStatus(arg0 common.Address) (bool, error) {
	return _CityPioneer.Contract.DelegateRewardStatus(&_CityPioneer.CallOpts, arg0)
}

// FundsReward is a free data retrieval call binding the contract method 0xc2f3afa9.
//
// Solidity: function fundsReward(address ) view returns(uint256)
func (_CityPioneer *CityPioneerCaller) FundsReward(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CityPioneer.contract.Call(opts, &out, "fundsReward", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FundsReward is a free data retrieval call binding the contract method 0xc2f3afa9.
//
// Solidity: function fundsReward(address ) view returns(uint256)
func (_CityPioneer *CityPioneerSession) FundsReward(arg0 common.Address) (*big.Int, error) {
	return _CityPioneer.Contract.FundsReward(&_CityPioneer.CallOpts, arg0)
}

// FundsReward is a free data retrieval call binding the contract method 0xc2f3afa9.
//
// Solidity: function fundsReward(address ) view returns(uint256)
func (_CityPioneer *CityPioneerCallerSession) FundsReward(arg0 common.Address) (*big.Int, error) {
	return _CityPioneer.Contract.FundsReward(&_CityPioneer.CallOpts, arg0)
}

// FundsRewardReceived is a free data retrieval call binding the contract method 0xf9e33bee.
//
// Solidity: function fundsRewardReceived(address ) view returns(uint256)
func (_CityPioneer *CityPioneerCaller) FundsRewardReceived(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CityPioneer.contract.Call(opts, &out, "fundsRewardReceived", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FundsRewardReceived is a free data retrieval call binding the contract method 0xf9e33bee.
//
// Solidity: function fundsRewardReceived(address ) view returns(uint256)
func (_CityPioneer *CityPioneerSession) FundsRewardReceived(arg0 common.Address) (*big.Int, error) {
	return _CityPioneer.Contract.FundsRewardReceived(&_CityPioneer.CallOpts, arg0)
}

// FundsRewardReceived is a free data retrieval call binding the contract method 0xf9e33bee.
//
// Solidity: function fundsRewardReceived(address ) view returns(uint256)
func (_CityPioneer *CityPioneerCallerSession) FundsRewardReceived(arg0 common.Address) (*big.Int, error) {
	return _CityPioneer.Contract.FundsRewardReceived(&_CityPioneer.CallOpts, arg0)
}

// FundsRewardStatus is a free data retrieval call binding the contract method 0xdfcdb468.
//
// Solidity: function fundsRewardStatus(address ) view returns(bool)
func (_CityPioneer *CityPioneerCaller) FundsRewardStatus(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _CityPioneer.contract.Call(opts, &out, "fundsRewardStatus", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// FundsRewardStatus is a free data retrieval call binding the contract method 0xdfcdb468.
//
// Solidity: function fundsRewardStatus(address ) view returns(bool)
func (_CityPioneer *CityPioneerSession) FundsRewardStatus(arg0 common.Address) (bool, error) {
	return _CityPioneer.Contract.FundsRewardStatus(&_CityPioneer.CallOpts, arg0)
}

// FundsRewardStatus is a free data retrieval call binding the contract method 0xdfcdb468.
//
// Solidity: function fundsRewardStatus(address ) view returns(bool)
func (_CityPioneer *CityPioneerCallerSession) FundsRewardStatus(arg0 common.Address) (bool, error) {
	return _CityPioneer.Contract.FundsRewardStatus(&_CityPioneer.CallOpts, arg0)
}

// GetDay is a free data retrieval call binding the contract method 0x14ba5c09.
//
// Solidity: function getDay() view returns(uint256)
func (_CityPioneer *CityPioneerCaller) GetDay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CityPioneer.contract.Call(opts, &out, "getDay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDay is a free data retrieval call binding the contract method 0x14ba5c09.
//
// Solidity: function getDay() view returns(uint256)
func (_CityPioneer *CityPioneerSession) GetDay() (*big.Int, error) {
	return _CityPioneer.Contract.GetDay(&_CityPioneer.CallOpts)
}

// GetDay is a free data retrieval call binding the contract method 0x14ba5c09.
//
// Solidity: function getDay() view returns(uint256)
func (_CityPioneer *CityPioneerCallerSession) GetDay() (*big.Int, error) {
	return _CityPioneer.Contract.GetDay(&_CityPioneer.CallOpts)
}

// GetSurety is a free data retrieval call binding the contract method 0xe9914078.
//
// Solidity: function getSurety() view returns(uint256)
func (_CityPioneer *CityPioneerCaller) GetSurety(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CityPioneer.contract.Call(opts, &out, "getSurety")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSurety is a free data retrieval call binding the contract method 0xe9914078.
//
// Solidity: function getSurety() view returns(uint256)
func (_CityPioneer *CityPioneerSession) GetSurety() (*big.Int, error) {
	return _CityPioneer.Contract.GetSurety(&_CityPioneer.CallOpts)
}

// GetSurety is a free data retrieval call binding the contract method 0xe9914078.
//
// Solidity: function getSurety() view returns(uint256)
func (_CityPioneer *CityPioneerCallerSession) GetSurety() (*big.Int, error) {
	return _CityPioneer.Contract.GetSurety(&_CityPioneer.CallOpts)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address account) view returns(bool)
func (_CityPioneer *CityPioneerCaller) IsAdmin(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _CityPioneer.contract.Call(opts, &out, "isAdmin", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address account) view returns(bool)
func (_CityPioneer *CityPioneerSession) IsAdmin(account common.Address) (bool, error) {
	return _CityPioneer.Contract.IsAdmin(&_CityPioneer.CallOpts, account)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address account) view returns(bool)
func (_CityPioneer *CityPioneerCallerSession) IsAdmin(account common.Address) (bool, error) {
	return _CityPioneer.Contract.IsAdmin(&_CityPioneer.CallOpts, account)
}

// IsPioneer is a free data retrieval call binding the contract method 0x197d63a1.
//
// Solidity: function isPioneer(address user) view returns(bool)
func (_CityPioneer *CityPioneerCaller) IsPioneer(opts *bind.CallOpts, user common.Address) (bool, error) {
	var out []interface{}
	err := _CityPioneer.contract.Call(opts, &out, "isPioneer", user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPioneer is a free data retrieval call binding the contract method 0x197d63a1.
//
// Solidity: function isPioneer(address user) view returns(bool)
func (_CityPioneer *CityPioneerSession) IsPioneer(user common.Address) (bool, error) {
	return _CityPioneer.Contract.IsPioneer(&_CityPioneer.CallOpts, user)
}

// IsPioneer is a free data retrieval call binding the contract method 0x197d63a1.
//
// Solidity: function isPioneer(address user) view returns(bool)
func (_CityPioneer *CityPioneerCallerSession) IsPioneer(user common.Address) (bool, error) {
	return _CityPioneer.Contract.IsPioneer(&_CityPioneer.CallOpts, user)
}

// MiningAddress is a free data retrieval call binding the contract method 0x88ece43f.
//
// Solidity: function miningAddress() view returns(address)
func (_CityPioneer *CityPioneerCaller) MiningAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CityPioneer.contract.Call(opts, &out, "miningAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MiningAddress is a free data retrieval call binding the contract method 0x88ece43f.
//
// Solidity: function miningAddress() view returns(address)
func (_CityPioneer *CityPioneerSession) MiningAddress() (common.Address, error) {
	return _CityPioneer.Contract.MiningAddress(&_CityPioneer.CallOpts)
}

// MiningAddress is a free data retrieval call binding the contract method 0x88ece43f.
//
// Solidity: function miningAddress() view returns(address)
func (_CityPioneer *CityPioneerCallerSession) MiningAddress() (common.Address, error) {
	return _CityPioneer.Contract.MiningAddress(&_CityPioneer.CallOpts)
}

// PioneerInfo is a free data retrieval call binding the contract method 0x0980a959.
//
// Solidity: function pioneerInfo(address ) view returns(address pioneerAddress, uint256 ctime, uint256 cityLevel, bool assessmentMonthStatus, bool assessmentStatus, bool returnSuretyStatus, uint256 returnSuretyRate, uint256 returnSuretyTime)
func (_CityPioneer *CityPioneerCaller) PioneerInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	PioneerAddress        common.Address
	Ctime                 *big.Int
	CityLevel             *big.Int
	AssessmentMonthStatus bool
	AssessmentStatus      bool
	ReturnSuretyStatus    bool
	ReturnSuretyRate      *big.Int
	ReturnSuretyTime      *big.Int
}, error) {
	var out []interface{}
	err := _CityPioneer.contract.Call(opts, &out, "pioneerInfo", arg0)

	outstruct := new(struct {
		PioneerAddress        common.Address
		Ctime                 *big.Int
		CityLevel             *big.Int
		AssessmentMonthStatus bool
		AssessmentStatus      bool
		ReturnSuretyStatus    bool
		ReturnSuretyRate      *big.Int
		ReturnSuretyTime      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PioneerAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Ctime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.CityLevel = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.AssessmentMonthStatus = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.AssessmentStatus = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.ReturnSuretyStatus = *abi.ConvertType(out[5], new(bool)).(*bool)
	outstruct.ReturnSuretyRate = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.ReturnSuretyTime = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PioneerInfo is a free data retrieval call binding the contract method 0x0980a959.
//
// Solidity: function pioneerInfo(address ) view returns(address pioneerAddress, uint256 ctime, uint256 cityLevel, bool assessmentMonthStatus, bool assessmentStatus, bool returnSuretyStatus, uint256 returnSuretyRate, uint256 returnSuretyTime)
func (_CityPioneer *CityPioneerSession) PioneerInfo(arg0 common.Address) (struct {
	PioneerAddress        common.Address
	Ctime                 *big.Int
	CityLevel             *big.Int
	AssessmentMonthStatus bool
	AssessmentStatus      bool
	ReturnSuretyStatus    bool
	ReturnSuretyRate      *big.Int
	ReturnSuretyTime      *big.Int
}, error) {
	return _CityPioneer.Contract.PioneerInfo(&_CityPioneer.CallOpts, arg0)
}

// PioneerInfo is a free data retrieval call binding the contract method 0x0980a959.
//
// Solidity: function pioneerInfo(address ) view returns(address pioneerAddress, uint256 ctime, uint256 cityLevel, bool assessmentMonthStatus, bool assessmentStatus, bool returnSuretyStatus, uint256 returnSuretyRate, uint256 returnSuretyTime)
func (_CityPioneer *CityPioneerCallerSession) PioneerInfo(arg0 common.Address) (struct {
	PioneerAddress        common.Address
	Ctime                 *big.Int
	CityLevel             *big.Int
	AssessmentMonthStatus bool
	AssessmentStatus      bool
	ReturnSuretyStatus    bool
	ReturnSuretyRate      *big.Int
	ReturnSuretyTime      *big.Int
}, error) {
	return _CityPioneer.Contract.PioneerInfo(&_CityPioneer.CallOpts, arg0)
}

// SuretyReward is a free data retrieval call binding the contract method 0xbad9aaa6.
//
// Solidity: function suretyReward(address ) view returns(uint256)
func (_CityPioneer *CityPioneerCaller) SuretyReward(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CityPioneer.contract.Call(opts, &out, "suretyReward", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SuretyReward is a free data retrieval call binding the contract method 0xbad9aaa6.
//
// Solidity: function suretyReward(address ) view returns(uint256)
func (_CityPioneer *CityPioneerSession) SuretyReward(arg0 common.Address) (*big.Int, error) {
	return _CityPioneer.Contract.SuretyReward(&_CityPioneer.CallOpts, arg0)
}

// SuretyReward is a free data retrieval call binding the contract method 0xbad9aaa6.
//
// Solidity: function suretyReward(address ) view returns(uint256)
func (_CityPioneer *CityPioneerCallerSession) SuretyReward(arg0 common.Address) (*big.Int, error) {
	return _CityPioneer.Contract.SuretyReward(&_CityPioneer.CallOpts, arg0)
}

// SuretyRewardRecord is a free data retrieval call binding the contract method 0xa597c8f4.
//
// Solidity: function suretyRewardRecord(address ) view returns(uint256)
func (_CityPioneer *CityPioneerCaller) SuretyRewardRecord(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CityPioneer.contract.Call(opts, &out, "suretyRewardRecord", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SuretyRewardRecord is a free data retrieval call binding the contract method 0xa597c8f4.
//
// Solidity: function suretyRewardRecord(address ) view returns(uint256)
func (_CityPioneer *CityPioneerSession) SuretyRewardRecord(arg0 common.Address) (*big.Int, error) {
	return _CityPioneer.Contract.SuretyRewardRecord(&_CityPioneer.CallOpts, arg0)
}

// SuretyRewardRecord is a free data retrieval call binding the contract method 0xa597c8f4.
//
// Solidity: function suretyRewardRecord(address ) view returns(uint256)
func (_CityPioneer *CityPioneerCallerSession) SuretyRewardRecord(arg0 common.Address) (*big.Int, error) {
	return _CityPioneer.Contract.SuretyRewardRecord(&_CityPioneer.CallOpts, arg0)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address account) returns()
func (_CityPioneer *CityPioneerTransactor) AddAdmin(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _CityPioneer.contract.Transact(opts, "addAdmin", account)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address account) returns()
func (_CityPioneer *CityPioneerSession) AddAdmin(account common.Address) (*types.Transaction, error) {
	return _CityPioneer.Contract.AddAdmin(&_CityPioneer.TransactOpts, account)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address account) returns()
func (_CityPioneer *CityPioneerTransactorSession) AddAdmin(account common.Address) (*types.Transaction, error) {
	return _CityPioneer.Contract.AddAdmin(&_CityPioneer.TransactOpts, account)
}

// AdminSetAssessmentCriteria is a paid mutator transaction binding the contract method 0x98432900.
//
// Solidity: function adminSetAssessmentCriteria(uint256 cityLevel_, uint256 month_, uint256 point_) returns()
func (_CityPioneer *CityPioneerTransactor) AdminSetAssessmentCriteria(opts *bind.TransactOpts, cityLevel_ *big.Int, month_ *big.Int, point_ *big.Int) (*types.Transaction, error) {
	return _CityPioneer.contract.Transact(opts, "adminSetAssessmentCriteria", cityLevel_, month_, point_)
}

// AdminSetAssessmentCriteria is a paid mutator transaction binding the contract method 0x98432900.
//
// Solidity: function adminSetAssessmentCriteria(uint256 cityLevel_, uint256 month_, uint256 point_) returns()
func (_CityPioneer *CityPioneerSession) AdminSetAssessmentCriteria(cityLevel_ *big.Int, month_ *big.Int, point_ *big.Int) (*types.Transaction, error) {
	return _CityPioneer.Contract.AdminSetAssessmentCriteria(&_CityPioneer.TransactOpts, cityLevel_, month_, point_)
}

// AdminSetAssessmentCriteria is a paid mutator transaction binding the contract method 0x98432900.
//
// Solidity: function adminSetAssessmentCriteria(uint256 cityLevel_, uint256 month_, uint256 point_) returns()
func (_CityPioneer *CityPioneerTransactorSession) AdminSetAssessmentCriteria(cityLevel_ *big.Int, month_ *big.Int, point_ *big.Int) (*types.Transaction, error) {
	return _CityPioneer.Contract.AdminSetAssessmentCriteria(&_CityPioneer.TransactOpts, cityLevel_, month_, point_)
}

// AdminSetCityAddress is a paid mutator transaction binding the contract method 0xa94c843d.
//
// Solidity: function adminSetCityAddress(address cityAddress_) returns()
func (_CityPioneer *CityPioneerTransactor) AdminSetCityAddress(opts *bind.TransactOpts, cityAddress_ common.Address) (*types.Transaction, error) {
	return _CityPioneer.contract.Transact(opts, "adminSetCityAddress", cityAddress_)
}

// AdminSetCityAddress is a paid mutator transaction binding the contract method 0xa94c843d.
//
// Solidity: function adminSetCityAddress(address cityAddress_) returns()
func (_CityPioneer *CityPioneerSession) AdminSetCityAddress(cityAddress_ common.Address) (*types.Transaction, error) {
	return _CityPioneer.Contract.AdminSetCityAddress(&_CityPioneer.TransactOpts, cityAddress_)
}

// AdminSetCityAddress is a paid mutator transaction binding the contract method 0xa94c843d.
//
// Solidity: function adminSetCityAddress(address cityAddress_) returns()
func (_CityPioneer *CityPioneerTransactorSession) AdminSetCityAddress(cityAddress_ common.Address) (*types.Transaction, error) {
	return _CityPioneer.Contract.AdminSetCityAddress(&_CityPioneer.TransactOpts, cityAddress_)
}

// AdminSetMiningAddress is a paid mutator transaction binding the contract method 0x3ed16cd1.
//
// Solidity: function adminSetMiningAddress(address miningAddress_) returns()
func (_CityPioneer *CityPioneerTransactor) AdminSetMiningAddress(opts *bind.TransactOpts, miningAddress_ common.Address) (*types.Transaction, error) {
	return _CityPioneer.contract.Transact(opts, "adminSetMiningAddress", miningAddress_)
}

// AdminSetMiningAddress is a paid mutator transaction binding the contract method 0x3ed16cd1.
//
// Solidity: function adminSetMiningAddress(address miningAddress_) returns()
func (_CityPioneer *CityPioneerSession) AdminSetMiningAddress(miningAddress_ common.Address) (*types.Transaction, error) {
	return _CityPioneer.Contract.AdminSetMiningAddress(&_CityPioneer.TransactOpts, miningAddress_)
}

// AdminSetMiningAddress is a paid mutator transaction binding the contract method 0x3ed16cd1.
//
// Solidity: function adminSetMiningAddress(address miningAddress_) returns()
func (_CityPioneer *CityPioneerTransactorSession) AdminSetMiningAddress(miningAddress_ common.Address) (*types.Transaction, error) {
	return _CityPioneer.Contract.AdminSetMiningAddress(&_CityPioneer.TransactOpts, miningAddress_)
}

// AdminSetTOXAddress is a paid mutator transaction binding the contract method 0x036c1009.
//
// Solidity: function adminSetTOXAddress(address TOXAddress_) returns()
func (_CityPioneer *CityPioneerTransactor) AdminSetTOXAddress(opts *bind.TransactOpts, TOXAddress_ common.Address) (*types.Transaction, error) {
	return _CityPioneer.contract.Transact(opts, "adminSetTOXAddress", TOXAddress_)
}

// AdminSetTOXAddress is a paid mutator transaction binding the contract method 0x036c1009.
//
// Solidity: function adminSetTOXAddress(address TOXAddress_) returns()
func (_CityPioneer *CityPioneerSession) AdminSetTOXAddress(TOXAddress_ common.Address) (*types.Transaction, error) {
	return _CityPioneer.Contract.AdminSetTOXAddress(&_CityPioneer.TransactOpts, TOXAddress_)
}

// AdminSetTOXAddress is a paid mutator transaction binding the contract method 0x036c1009.
//
// Solidity: function adminSetTOXAddress(address TOXAddress_) returns()
func (_CityPioneer *CityPioneerTransactorSession) AdminSetTOXAddress(TOXAddress_ common.Address) (*types.Transaction, error) {
	return _CityPioneer.Contract.AdminSetTOXAddress(&_CityPioneer.TransactOpts, TOXAddress_)
}

// BatchAddAdmin is a paid mutator transaction binding the contract method 0x2c9ab42b.
//
// Solidity: function batchAddAdmin(address[] amounts) returns()
func (_CityPioneer *CityPioneerTransactor) BatchAddAdmin(opts *bind.TransactOpts, amounts []common.Address) (*types.Transaction, error) {
	return _CityPioneer.contract.Transact(opts, "batchAddAdmin", amounts)
}

// BatchAddAdmin is a paid mutator transaction binding the contract method 0x2c9ab42b.
//
// Solidity: function batchAddAdmin(address[] amounts) returns()
func (_CityPioneer *CityPioneerSession) BatchAddAdmin(amounts []common.Address) (*types.Transaction, error) {
	return _CityPioneer.Contract.BatchAddAdmin(&_CityPioneer.TransactOpts, amounts)
}

// BatchAddAdmin is a paid mutator transaction binding the contract method 0x2c9ab42b.
//
// Solidity: function batchAddAdmin(address[] amounts) returns()
func (_CityPioneer *CityPioneerTransactorSession) BatchAddAdmin(amounts []common.Address) (*types.Transaction, error) {
	return _CityPioneer.Contract.BatchAddAdmin(&_CityPioneer.TransactOpts, amounts)
}

// DepositSurety is a paid mutator transaction binding the contract method 0x07edb9c4.
//
// Solidity: function depositSurety() returns(uint256, uint256, bytes32)
func (_CityPioneer *CityPioneerTransactor) DepositSurety(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CityPioneer.contract.Transact(opts, "depositSurety")
}

// DepositSurety is a paid mutator transaction binding the contract method 0x07edb9c4.
//
// Solidity: function depositSurety() returns(uint256, uint256, bytes32)
func (_CityPioneer *CityPioneerSession) DepositSurety() (*types.Transaction, error) {
	return _CityPioneer.Contract.DepositSurety(&_CityPioneer.TransactOpts)
}

// DepositSurety is a paid mutator transaction binding the contract method 0x07edb9c4.
//
// Solidity: function depositSurety() returns(uint256, uint256, bytes32)
func (_CityPioneer *CityPioneerTransactorSession) DepositSurety() (*types.Transaction, error) {
	return _CityPioneer.Contract.DepositSurety(&_CityPioneer.TransactOpts)
}

// DepositSuretyTest is a paid mutator transaction binding the contract method 0x3cfba34d.
//
// Solidity: function depositSuretyTest(address pioneerTest_) returns()
func (_CityPioneer *CityPioneerTransactor) DepositSuretyTest(opts *bind.TransactOpts, pioneerTest_ common.Address) (*types.Transaction, error) {
	return _CityPioneer.contract.Transact(opts, "depositSuretyTest", pioneerTest_)
}

// DepositSuretyTest is a paid mutator transaction binding the contract method 0x3cfba34d.
//
// Solidity: function depositSuretyTest(address pioneerTest_) returns()
func (_CityPioneer *CityPioneerSession) DepositSuretyTest(pioneerTest_ common.Address) (*types.Transaction, error) {
	return _CityPioneer.Contract.DepositSuretyTest(&_CityPioneer.TransactOpts, pioneerTest_)
}

// DepositSuretyTest is a paid mutator transaction binding the contract method 0x3cfba34d.
//
// Solidity: function depositSuretyTest(address pioneerTest_) returns()
func (_CityPioneer *CityPioneerTransactorSession) DepositSuretyTest(pioneerTest_ common.Address) (*types.Transaction, error) {
	return _CityPioneer.Contract.DepositSuretyTest(&_CityPioneer.TransactOpts, pioneerTest_)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_CityPioneer *CityPioneerTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CityPioneer.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_CityPioneer *CityPioneerSession) Initialize() (*types.Transaction, error) {
	return _CityPioneer.Contract.Initialize(&_CityPioneer.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_CityPioneer *CityPioneerTransactorSession) Initialize() (*types.Transaction, error) {
	return _CityPioneer.Contract.Initialize(&_CityPioneer.TransactOpts)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address account) returns()
func (_CityPioneer *CityPioneerTransactor) RemoveAdmin(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _CityPioneer.contract.Transact(opts, "removeAdmin", account)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address account) returns()
func (_CityPioneer *CityPioneerSession) RemoveAdmin(account common.Address) (*types.Transaction, error) {
	return _CityPioneer.Contract.RemoveAdmin(&_CityPioneer.TransactOpts, account)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address account) returns()
func (_CityPioneer *CityPioneerTransactorSession) RemoveAdmin(account common.Address) (*types.Transaction, error) {
	return _CityPioneer.Contract.RemoveAdmin(&_CityPioneer.TransactOpts, account)
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x8bad0c0a.
//
// Solidity: function renounceAdmin() returns()
func (_CityPioneer *CityPioneerTransactor) RenounceAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CityPioneer.contract.Transact(opts, "renounceAdmin")
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x8bad0c0a.
//
// Solidity: function renounceAdmin() returns()
func (_CityPioneer *CityPioneerSession) RenounceAdmin() (*types.Transaction, error) {
	return _CityPioneer.Contract.RenounceAdmin(&_CityPioneer.TransactOpts)
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x8bad0c0a.
//
// Solidity: function renounceAdmin() returns()
func (_CityPioneer *CityPioneerTransactorSession) RenounceAdmin() (*types.Transaction, error) {
	return _CityPioneer.Contract.RenounceAdmin(&_CityPioneer.TransactOpts)
}

// SetPioneerDelegate is a paid mutator transaction binding the contract method 0x845c0205.
//
// Solidity: function setPioneerDelegate(address pioneerAddress_, uint256 amount_, bytes32 cityId_) returns()
func (_CityPioneer *CityPioneerTransactor) SetPioneerDelegate(opts *bind.TransactOpts, pioneerAddress_ common.Address, amount_ *big.Int, cityId_ [32]byte) (*types.Transaction, error) {
	return _CityPioneer.contract.Transact(opts, "setPioneerDelegate", pioneerAddress_, amount_, cityId_)
}

// SetPioneerDelegate is a paid mutator transaction binding the contract method 0x845c0205.
//
// Solidity: function setPioneerDelegate(address pioneerAddress_, uint256 amount_, bytes32 cityId_) returns()
func (_CityPioneer *CityPioneerSession) SetPioneerDelegate(pioneerAddress_ common.Address, amount_ *big.Int, cityId_ [32]byte) (*types.Transaction, error) {
	return _CityPioneer.Contract.SetPioneerDelegate(&_CityPioneer.TransactOpts, pioneerAddress_, amount_, cityId_)
}

// SetPioneerDelegate is a paid mutator transaction binding the contract method 0x845c0205.
//
// Solidity: function setPioneerDelegate(address pioneerAddress_, uint256 amount_, bytes32 cityId_) returns()
func (_CityPioneer *CityPioneerTransactorSession) SetPioneerDelegate(pioneerAddress_ common.Address, amount_ *big.Int, cityId_ [32]byte) (*types.Transaction, error) {
	return _CityPioneer.Contract.SetPioneerDelegate(&_CityPioneer.TransactOpts, pioneerAddress_, amount_, cityId_)
}

// WithdrawalBenefitPackageReward is a paid mutator transaction binding the contract method 0x8f22826c.
//
// Solidity: function withdrawalBenefitPackageReward(uint256 amount_) returns()
func (_CityPioneer *CityPioneerTransactor) WithdrawalBenefitPackageReward(opts *bind.TransactOpts, amount_ *big.Int) (*types.Transaction, error) {
	return _CityPioneer.contract.Transact(opts, "withdrawalBenefitPackageReward", amount_)
}

// WithdrawalBenefitPackageReward is a paid mutator transaction binding the contract method 0x8f22826c.
//
// Solidity: function withdrawalBenefitPackageReward(uint256 amount_) returns()
func (_CityPioneer *CityPioneerSession) WithdrawalBenefitPackageReward(amount_ *big.Int) (*types.Transaction, error) {
	return _CityPioneer.Contract.WithdrawalBenefitPackageReward(&_CityPioneer.TransactOpts, amount_)
}

// WithdrawalBenefitPackageReward is a paid mutator transaction binding the contract method 0x8f22826c.
//
// Solidity: function withdrawalBenefitPackageReward(uint256 amount_) returns()
func (_CityPioneer *CityPioneerTransactorSession) WithdrawalBenefitPackageReward(amount_ *big.Int) (*types.Transaction, error) {
	return _CityPioneer.Contract.WithdrawalBenefitPackageReward(&_CityPioneer.TransactOpts, amount_)
}

// WithdrawalDelegateRewardReward is a paid mutator transaction binding the contract method 0x972c5ce4.
//
// Solidity: function withdrawalDelegateRewardReward(uint256 amount_) returns()
func (_CityPioneer *CityPioneerTransactor) WithdrawalDelegateRewardReward(opts *bind.TransactOpts, amount_ *big.Int) (*types.Transaction, error) {
	return _CityPioneer.contract.Transact(opts, "withdrawalDelegateRewardReward", amount_)
}

// WithdrawalDelegateRewardReward is a paid mutator transaction binding the contract method 0x972c5ce4.
//
// Solidity: function withdrawalDelegateRewardReward(uint256 amount_) returns()
func (_CityPioneer *CityPioneerSession) WithdrawalDelegateRewardReward(amount_ *big.Int) (*types.Transaction, error) {
	return _CityPioneer.Contract.WithdrawalDelegateRewardReward(&_CityPioneer.TransactOpts, amount_)
}

// WithdrawalDelegateRewardReward is a paid mutator transaction binding the contract method 0x972c5ce4.
//
// Solidity: function withdrawalDelegateRewardReward(uint256 amount_) returns()
func (_CityPioneer *CityPioneerTransactorSession) WithdrawalDelegateRewardReward(amount_ *big.Int) (*types.Transaction, error) {
	return _CityPioneer.Contract.WithdrawalDelegateRewardReward(&_CityPioneer.TransactOpts, amount_)
}

// WithdrawalFundsRewardReward is a paid mutator transaction binding the contract method 0x298018ea.
//
// Solidity: function withdrawalFundsRewardReward(uint256 amount_) returns()
func (_CityPioneer *CityPioneerTransactor) WithdrawalFundsRewardReward(opts *bind.TransactOpts, amount_ *big.Int) (*types.Transaction, error) {
	return _CityPioneer.contract.Transact(opts, "withdrawalFundsRewardReward", amount_)
}

// WithdrawalFundsRewardReward is a paid mutator transaction binding the contract method 0x298018ea.
//
// Solidity: function withdrawalFundsRewardReward(uint256 amount_) returns()
func (_CityPioneer *CityPioneerSession) WithdrawalFundsRewardReward(amount_ *big.Int) (*types.Transaction, error) {
	return _CityPioneer.Contract.WithdrawalFundsRewardReward(&_CityPioneer.TransactOpts, amount_)
}

// WithdrawalFundsRewardReward is a paid mutator transaction binding the contract method 0x298018ea.
//
// Solidity: function withdrawalFundsRewardReward(uint256 amount_) returns()
func (_CityPioneer *CityPioneerTransactorSession) WithdrawalFundsRewardReward(amount_ *big.Int) (*types.Transaction, error) {
	return _CityPioneer.Contract.WithdrawalFundsRewardReward(&_CityPioneer.TransactOpts, amount_)
}

// WithdrawalSuretyReward is a paid mutator transaction binding the contract method 0xbad5a393.
//
// Solidity: function withdrawalSuretyReward(uint256 amount_) returns()
func (_CityPioneer *CityPioneerTransactor) WithdrawalSuretyReward(opts *bind.TransactOpts, amount_ *big.Int) (*types.Transaction, error) {
	return _CityPioneer.contract.Transact(opts, "withdrawalSuretyReward", amount_)
}

// WithdrawalSuretyReward is a paid mutator transaction binding the contract method 0xbad5a393.
//
// Solidity: function withdrawalSuretyReward(uint256 amount_) returns()
func (_CityPioneer *CityPioneerSession) WithdrawalSuretyReward(amount_ *big.Int) (*types.Transaction, error) {
	return _CityPioneer.Contract.WithdrawalSuretyReward(&_CityPioneer.TransactOpts, amount_)
}

// WithdrawalSuretyReward is a paid mutator transaction binding the contract method 0xbad5a393.
//
// Solidity: function withdrawalSuretyReward(uint256 amount_) returns()
func (_CityPioneer *CityPioneerTransactorSession) WithdrawalSuretyReward(amount_ *big.Int) (*types.Transaction, error) {
	return _CityPioneer.Contract.WithdrawalSuretyReward(&_CityPioneer.TransactOpts, amount_)
}

// CityPioneerAdminAddedIterator is returned from FilterAdminAdded and is used to iterate over the raw logs and unpacked data for AdminAdded events raised by the CityPioneer contract.
type CityPioneerAdminAddedIterator struct {
	Event *CityPioneerAdminAdded // Event containing the contract specifics and raw log

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
func (it *CityPioneerAdminAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CityPioneerAdminAdded)
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
		it.Event = new(CityPioneerAdminAdded)
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
func (it *CityPioneerAdminAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CityPioneerAdminAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CityPioneerAdminAdded represents a AdminAdded event raised by the CityPioneer contract.
type CityPioneerAdminAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAdminAdded is a free log retrieval operation binding the contract event 0x44d6d25963f097ad14f29f06854a01f575648a1ef82f30e562ccd3889717e339.
//
// Solidity: event AdminAdded(address indexed account)
func (_CityPioneer *CityPioneerFilterer) FilterAdminAdded(opts *bind.FilterOpts, account []common.Address) (*CityPioneerAdminAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _CityPioneer.contract.FilterLogs(opts, "AdminAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &CityPioneerAdminAddedIterator{contract: _CityPioneer.contract, event: "AdminAdded", logs: logs, sub: sub}, nil
}

// WatchAdminAdded is a free log subscription operation binding the contract event 0x44d6d25963f097ad14f29f06854a01f575648a1ef82f30e562ccd3889717e339.
//
// Solidity: event AdminAdded(address indexed account)
func (_CityPioneer *CityPioneerFilterer) WatchAdminAdded(opts *bind.WatchOpts, sink chan<- *CityPioneerAdminAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _CityPioneer.contract.WatchLogs(opts, "AdminAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CityPioneerAdminAdded)
				if err := _CityPioneer.contract.UnpackLog(event, "AdminAdded", log); err != nil {
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

// ParseAdminAdded is a log parse operation binding the contract event 0x44d6d25963f097ad14f29f06854a01f575648a1ef82f30e562ccd3889717e339.
//
// Solidity: event AdminAdded(address indexed account)
func (_CityPioneer *CityPioneerFilterer) ParseAdminAdded(log types.Log) (*CityPioneerAdminAdded, error) {
	event := new(CityPioneerAdminAdded)
	if err := _CityPioneer.contract.UnpackLog(event, "AdminAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CityPioneerAdminRemovedIterator is returned from FilterAdminRemoved and is used to iterate over the raw logs and unpacked data for AdminRemoved events raised by the CityPioneer contract.
type CityPioneerAdminRemovedIterator struct {
	Event *CityPioneerAdminRemoved // Event containing the contract specifics and raw log

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
func (it *CityPioneerAdminRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CityPioneerAdminRemoved)
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
		it.Event = new(CityPioneerAdminRemoved)
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
func (it *CityPioneerAdminRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CityPioneerAdminRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CityPioneerAdminRemoved represents a AdminRemoved event raised by the CityPioneer contract.
type CityPioneerAdminRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAdminRemoved is a free log retrieval operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: event AdminRemoved(address indexed account)
func (_CityPioneer *CityPioneerFilterer) FilterAdminRemoved(opts *bind.FilterOpts, account []common.Address) (*CityPioneerAdminRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _CityPioneer.contract.FilterLogs(opts, "AdminRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &CityPioneerAdminRemovedIterator{contract: _CityPioneer.contract, event: "AdminRemoved", logs: logs, sub: sub}, nil
}

// WatchAdminRemoved is a free log subscription operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: event AdminRemoved(address indexed account)
func (_CityPioneer *CityPioneerFilterer) WatchAdminRemoved(opts *bind.WatchOpts, sink chan<- *CityPioneerAdminRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _CityPioneer.contract.WatchLogs(opts, "AdminRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CityPioneerAdminRemoved)
				if err := _CityPioneer.contract.UnpackLog(event, "AdminRemoved", log); err != nil {
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

// ParseAdminRemoved is a log parse operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: event AdminRemoved(address indexed account)
func (_CityPioneer *CityPioneerFilterer) ParseAdminRemoved(log types.Log) (*CityPioneerAdminRemoved, error) {
	event := new(CityPioneerAdminRemoved)
	if err := _CityPioneer.contract.UnpackLog(event, "AdminRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CityPioneerDailyIncreaseDelegateRecordIterator is returned from FilterDailyIncreaseDelegateRecord and is used to iterate over the raw logs and unpacked data for DailyIncreaseDelegateRecord events raised by the CityPioneer contract.
type CityPioneerDailyIncreaseDelegateRecordIterator struct {
	Event *CityPioneerDailyIncreaseDelegateRecord // Event containing the contract specifics and raw log

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
func (it *CityPioneerDailyIncreaseDelegateRecordIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CityPioneerDailyIncreaseDelegateRecord)
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
		it.Event = new(CityPioneerDailyIncreaseDelegateRecord)
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
func (it *CityPioneerDailyIncreaseDelegateRecordIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CityPioneerDailyIncreaseDelegateRecordIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CityPioneerDailyIncreaseDelegateRecord represents a DailyIncreaseDelegateRecord event raised by the CityPioneer contract.
type CityPioneerDailyIncreaseDelegateRecord struct {
	PioneerAddress common.Address
	CityId         [32]byte
	Amount         *big.Int
	Ctime          *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDailyIncreaseDelegateRecord is a free log retrieval operation binding the contract event 0x1ebf328440508d9e865037e77390e7f290e59186a438c71b4189b8b006e3699a.
//
// Solidity: event DailyIncreaseDelegateRecord(address pioneerAddress, bytes32 cityId, uint256 amount, uint256 ctime)
func (_CityPioneer *CityPioneerFilterer) FilterDailyIncreaseDelegateRecord(opts *bind.FilterOpts) (*CityPioneerDailyIncreaseDelegateRecordIterator, error) {

	logs, sub, err := _CityPioneer.contract.FilterLogs(opts, "DailyIncreaseDelegateRecord")
	if err != nil {
		return nil, err
	}
	return &CityPioneerDailyIncreaseDelegateRecordIterator{contract: _CityPioneer.contract, event: "DailyIncreaseDelegateRecord", logs: logs, sub: sub}, nil
}

// WatchDailyIncreaseDelegateRecord is a free log subscription operation binding the contract event 0x1ebf328440508d9e865037e77390e7f290e59186a438c71b4189b8b006e3699a.
//
// Solidity: event DailyIncreaseDelegateRecord(address pioneerAddress, bytes32 cityId, uint256 amount, uint256 ctime)
func (_CityPioneer *CityPioneerFilterer) WatchDailyIncreaseDelegateRecord(opts *bind.WatchOpts, sink chan<- *CityPioneerDailyIncreaseDelegateRecord) (event.Subscription, error) {

	logs, sub, err := _CityPioneer.contract.WatchLogs(opts, "DailyIncreaseDelegateRecord")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CityPioneerDailyIncreaseDelegateRecord)
				if err := _CityPioneer.contract.UnpackLog(event, "DailyIncreaseDelegateRecord", log); err != nil {
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

// ParseDailyIncreaseDelegateRecord is a log parse operation binding the contract event 0x1ebf328440508d9e865037e77390e7f290e59186a438c71b4189b8b006e3699a.
//
// Solidity: event DailyIncreaseDelegateRecord(address pioneerAddress, bytes32 cityId, uint256 amount, uint256 ctime)
func (_CityPioneer *CityPioneerFilterer) ParseDailyIncreaseDelegateRecord(log types.Log) (*CityPioneerDailyIncreaseDelegateRecord, error) {
	event := new(CityPioneerDailyIncreaseDelegateRecord)
	if err := _CityPioneer.contract.UnpackLog(event, "DailyIncreaseDelegateRecord", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CityPioneerDailyRewardRecordIterator is returned from FilterDailyRewardRecord and is used to iterate over the raw logs and unpacked data for DailyRewardRecord events raised by the CityPioneer contract.
type CityPioneerDailyRewardRecordIterator struct {
	Event *CityPioneerDailyRewardRecord // Event containing the contract specifics and raw log

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
func (it *CityPioneerDailyRewardRecordIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CityPioneerDailyRewardRecord)
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
		it.Event = new(CityPioneerDailyRewardRecord)
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
func (it *CityPioneerDailyRewardRecordIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CityPioneerDailyRewardRecordIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CityPioneerDailyRewardRecord represents a DailyRewardRecord event raised by the CityPioneer contract.
type CityPioneerDailyRewardRecord struct {
	PioneerAddress common.Address
	ToxReward      *big.Int
	FoundsReward   *big.Int
	DelegateReward *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDailyRewardRecord is a free log retrieval operation binding the contract event 0xf68c222bf3134ae094ffb2a8502e2254a18dc93c697adda0858dc06cf6a89217.
//
// Solidity: event DailyRewardRecord(address pioneerAddress, uint256 toxReward, uint256 foundsReward, uint256 delegateReward)
func (_CityPioneer *CityPioneerFilterer) FilterDailyRewardRecord(opts *bind.FilterOpts) (*CityPioneerDailyRewardRecordIterator, error) {

	logs, sub, err := _CityPioneer.contract.FilterLogs(opts, "DailyRewardRecord")
	if err != nil {
		return nil, err
	}
	return &CityPioneerDailyRewardRecordIterator{contract: _CityPioneer.contract, event: "DailyRewardRecord", logs: logs, sub: sub}, nil
}

// WatchDailyRewardRecord is a free log subscription operation binding the contract event 0xf68c222bf3134ae094ffb2a8502e2254a18dc93c697adda0858dc06cf6a89217.
//
// Solidity: event DailyRewardRecord(address pioneerAddress, uint256 toxReward, uint256 foundsReward, uint256 delegateReward)
func (_CityPioneer *CityPioneerFilterer) WatchDailyRewardRecord(opts *bind.WatchOpts, sink chan<- *CityPioneerDailyRewardRecord) (event.Subscription, error) {

	logs, sub, err := _CityPioneer.contract.WatchLogs(opts, "DailyRewardRecord")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CityPioneerDailyRewardRecord)
				if err := _CityPioneer.contract.UnpackLog(event, "DailyRewardRecord", log); err != nil {
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

// ParseDailyRewardRecord is a log parse operation binding the contract event 0xf68c222bf3134ae094ffb2a8502e2254a18dc93c697adda0858dc06cf6a89217.
//
// Solidity: event DailyRewardRecord(address pioneerAddress, uint256 toxReward, uint256 foundsReward, uint256 delegateReward)
func (_CityPioneer *CityPioneerFilterer) ParseDailyRewardRecord(log types.Log) (*CityPioneerDailyRewardRecord, error) {
	event := new(CityPioneerDailyRewardRecord)
	if err := _CityPioneer.contract.UnpackLog(event, "DailyRewardRecord", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CityPioneerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the CityPioneer contract.
type CityPioneerInitializedIterator struct {
	Event *CityPioneerInitialized // Event containing the contract specifics and raw log

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
func (it *CityPioneerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CityPioneerInitialized)
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
		it.Event = new(CityPioneerInitialized)
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
func (it *CityPioneerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CityPioneerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CityPioneerInitialized represents a Initialized event raised by the CityPioneer contract.
type CityPioneerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_CityPioneer *CityPioneerFilterer) FilterInitialized(opts *bind.FilterOpts) (*CityPioneerInitializedIterator, error) {

	logs, sub, err := _CityPioneer.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &CityPioneerInitializedIterator{contract: _CityPioneer.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_CityPioneer *CityPioneerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *CityPioneerInitialized) (event.Subscription, error) {

	logs, sub, err := _CityPioneer.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CityPioneerInitialized)
				if err := _CityPioneer.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_CityPioneer *CityPioneerFilterer) ParseInitialized(log types.Log) (*CityPioneerInitialized, error) {
	event := new(CityPioneerInitialized)
	if err := _CityPioneer.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CityPioneerLogBytes32Iterator is returned from FilterLogBytes32 and is used to iterate over the raw logs and unpacked data for LogBytes32 events raised by the CityPioneer contract.
type CityPioneerLogBytes32Iterator struct {
	Event *CityPioneerLogBytes32 // Event containing the contract specifics and raw log

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
func (it *CityPioneerLogBytes32Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CityPioneerLogBytes32)
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
		it.Event = new(CityPioneerLogBytes32)
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
func (it *CityPioneerLogBytes32Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CityPioneerLogBytes32Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CityPioneerLogBytes32 represents a LogBytes32 event raised by the CityPioneer contract.
type CityPioneerLogBytes32 struct {
	CityId [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLogBytes32 is a free log retrieval operation binding the contract event 0x009fd52f05c0ded31d6fb0ee580b923f85e99cf1a5a1da342f25e73c45829c83.
//
// Solidity: event LogBytes32(bytes32 cityId)
func (_CityPioneer *CityPioneerFilterer) FilterLogBytes32(opts *bind.FilterOpts) (*CityPioneerLogBytes32Iterator, error) {

	logs, sub, err := _CityPioneer.contract.FilterLogs(opts, "LogBytes32")
	if err != nil {
		return nil, err
	}
	return &CityPioneerLogBytes32Iterator{contract: _CityPioneer.contract, event: "LogBytes32", logs: logs, sub: sub}, nil
}

// WatchLogBytes32 is a free log subscription operation binding the contract event 0x009fd52f05c0ded31d6fb0ee580b923f85e99cf1a5a1da342f25e73c45829c83.
//
// Solidity: event LogBytes32(bytes32 cityId)
func (_CityPioneer *CityPioneerFilterer) WatchLogBytes32(opts *bind.WatchOpts, sink chan<- *CityPioneerLogBytes32) (event.Subscription, error) {

	logs, sub, err := _CityPioneer.contract.WatchLogs(opts, "LogBytes32")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CityPioneerLogBytes32)
				if err := _CityPioneer.contract.UnpackLog(event, "LogBytes32", log); err != nil {
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

// ParseLogBytes32 is a log parse operation binding the contract event 0x009fd52f05c0ded31d6fb0ee580b923f85e99cf1a5a1da342f25e73c45829c83.
//
// Solidity: event LogBytes32(bytes32 cityId)
func (_CityPioneer *CityPioneerFilterer) ParseLogBytes32(log types.Log) (*CityPioneerLogBytes32, error) {
	event := new(CityPioneerLogBytes32)
	if err := _CityPioneer.contract.UnpackLog(event, "LogBytes32", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CityPioneerLogUint256Iterator is returned from FilterLogUint256 and is used to iterate over the raw logs and unpacked data for LogUint256 events raised by the CityPioneer contract.
type CityPioneerLogUint256Iterator struct {
	Event *CityPioneerLogUint256 // Event containing the contract specifics and raw log

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
func (it *CityPioneerLogUint256Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CityPioneerLogUint256)
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
		it.Event = new(CityPioneerLogUint256)
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
func (it *CityPioneerLogUint256Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CityPioneerLogUint256Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CityPioneerLogUint256 represents a LogUint256 event raised by the CityPioneer contract.
type CityPioneerLogUint256 struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLogUint256 is a free log retrieval operation binding the contract event 0x535266f26566acd2ef175615d9f1140b36f149b810b33fb93143236a69912c32.
//
// Solidity: event LogUint256(uint256 amount)
func (_CityPioneer *CityPioneerFilterer) FilterLogUint256(opts *bind.FilterOpts) (*CityPioneerLogUint256Iterator, error) {

	logs, sub, err := _CityPioneer.contract.FilterLogs(opts, "LogUint256")
	if err != nil {
		return nil, err
	}
	return &CityPioneerLogUint256Iterator{contract: _CityPioneer.contract, event: "LogUint256", logs: logs, sub: sub}, nil
}

// WatchLogUint256 is a free log subscription operation binding the contract event 0x535266f26566acd2ef175615d9f1140b36f149b810b33fb93143236a69912c32.
//
// Solidity: event LogUint256(uint256 amount)
func (_CityPioneer *CityPioneerFilterer) WatchLogUint256(opts *bind.WatchOpts, sink chan<- *CityPioneerLogUint256) (event.Subscription, error) {

	logs, sub, err := _CityPioneer.contract.WatchLogs(opts, "LogUint256")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CityPioneerLogUint256)
				if err := _CityPioneer.contract.UnpackLog(event, "LogUint256", log); err != nil {
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

// ParseLogUint256 is a log parse operation binding the contract event 0x535266f26566acd2ef175615d9f1140b36f149b810b33fb93143236a69912c32.
//
// Solidity: event LogUint256(uint256 amount)
func (_CityPioneer *CityPioneerFilterer) ParseLogUint256(log types.Log) (*CityPioneerLogUint256, error) {
	event := new(CityPioneerLogUint256)
	if err := _CityPioneer.contract.UnpackLog(event, "LogUint256", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CityPioneerSuretyRecordIterator is returned from FilterSuretyRecord and is used to iterate over the raw logs and unpacked data for SuretyRecord events raised by the CityPioneer contract.
type CityPioneerSuretyRecordIterator struct {
	Event *CityPioneerSuretyRecord // Event containing the contract specifics and raw log

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
func (it *CityPioneerSuretyRecordIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CityPioneerSuretyRecord)
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
		it.Event = new(CityPioneerSuretyRecord)
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
func (it *CityPioneerSuretyRecordIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CityPioneerSuretyRecordIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CityPioneerSuretyRecord represents a SuretyRecord event raised by the CityPioneer contract.
type CityPioneerSuretyRecord struct {
	PioneerAddress common.Address
	Amount         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSuretyRecord is a free log retrieval operation binding the contract event 0x1b1af76c73b9b87dfae9023ac3db8c6d94ad03d433aa9feec6f86509c6c501c7.
//
// Solidity: event SuretyRecord(address pioneerAddress, uint256 amount)
func (_CityPioneer *CityPioneerFilterer) FilterSuretyRecord(opts *bind.FilterOpts) (*CityPioneerSuretyRecordIterator, error) {

	logs, sub, err := _CityPioneer.contract.FilterLogs(opts, "SuretyRecord")
	if err != nil {
		return nil, err
	}
	return &CityPioneerSuretyRecordIterator{contract: _CityPioneer.contract, event: "SuretyRecord", logs: logs, sub: sub}, nil
}

// WatchSuretyRecord is a free log subscription operation binding the contract event 0x1b1af76c73b9b87dfae9023ac3db8c6d94ad03d433aa9feec6f86509c6c501c7.
//
// Solidity: event SuretyRecord(address pioneerAddress, uint256 amount)
func (_CityPioneer *CityPioneerFilterer) WatchSuretyRecord(opts *bind.WatchOpts, sink chan<- *CityPioneerSuretyRecord) (event.Subscription, error) {

	logs, sub, err := _CityPioneer.contract.WatchLogs(opts, "SuretyRecord")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CityPioneerSuretyRecord)
				if err := _CityPioneer.contract.UnpackLog(event, "SuretyRecord", log); err != nil {
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

// ParseSuretyRecord is a log parse operation binding the contract event 0x1b1af76c73b9b87dfae9023ac3db8c6d94ad03d433aa9feec6f86509c6c501c7.
//
// Solidity: event SuretyRecord(address pioneerAddress, uint256 amount)
func (_CityPioneer *CityPioneerFilterer) ParseSuretyRecord(log types.Log) (*CityPioneerSuretyRecord, error) {
	event := new(CityPioneerSuretyRecord)
	if err := _CityPioneer.contract.UnpackLog(event, "SuretyRecord", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CityPioneerWithdrawalRewardRecordIterator is returned from FilterWithdrawalRewardRecord and is used to iterate over the raw logs and unpacked data for WithdrawalRewardRecord events raised by the CityPioneer contract.
type CityPioneerWithdrawalRewardRecordIterator struct {
	Event *CityPioneerWithdrawalRewardRecord // Event containing the contract specifics and raw log

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
func (it *CityPioneerWithdrawalRewardRecordIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CityPioneerWithdrawalRewardRecord)
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
		it.Event = new(CityPioneerWithdrawalRewardRecord)
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
func (it *CityPioneerWithdrawalRewardRecordIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CityPioneerWithdrawalRewardRecordIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CityPioneerWithdrawalRewardRecord represents a WithdrawalRewardRecord event raised by the CityPioneer contract.
type CityPioneerWithdrawalRewardRecord struct {
	PioneerAddress common.Address
	Reward         *big.Int
	RewardType     *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalRewardRecord is a free log retrieval operation binding the contract event 0xd380455fa21ce288b613282a4d9830c5f32aad091426cab0aca190df48ec750a.
//
// Solidity: event WithdrawalRewardRecord(address pioneerAddress, uint256 reward, uint256 rewardType)
func (_CityPioneer *CityPioneerFilterer) FilterWithdrawalRewardRecord(opts *bind.FilterOpts) (*CityPioneerWithdrawalRewardRecordIterator, error) {

	logs, sub, err := _CityPioneer.contract.FilterLogs(opts, "WithdrawalRewardRecord")
	if err != nil {
		return nil, err
	}
	return &CityPioneerWithdrawalRewardRecordIterator{contract: _CityPioneer.contract, event: "WithdrawalRewardRecord", logs: logs, sub: sub}, nil
}

// WatchWithdrawalRewardRecord is a free log subscription operation binding the contract event 0xd380455fa21ce288b613282a4d9830c5f32aad091426cab0aca190df48ec750a.
//
// Solidity: event WithdrawalRewardRecord(address pioneerAddress, uint256 reward, uint256 rewardType)
func (_CityPioneer *CityPioneerFilterer) WatchWithdrawalRewardRecord(opts *bind.WatchOpts, sink chan<- *CityPioneerWithdrawalRewardRecord) (event.Subscription, error) {

	logs, sub, err := _CityPioneer.contract.WatchLogs(opts, "WithdrawalRewardRecord")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CityPioneerWithdrawalRewardRecord)
				if err := _CityPioneer.contract.UnpackLog(event, "WithdrawalRewardRecord", log); err != nil {
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

// ParseWithdrawalRewardRecord is a log parse operation binding the contract event 0xd380455fa21ce288b613282a4d9830c5f32aad091426cab0aca190df48ec750a.
//
// Solidity: event WithdrawalRewardRecord(address pioneerAddress, uint256 reward, uint256 rewardType)
func (_CityPioneer *CityPioneerFilterer) ParseWithdrawalRewardRecord(log types.Log) (*CityPioneerWithdrawalRewardRecord, error) {
	event := new(CityPioneerWithdrawalRewardRecord)
	if err := _CityPioneer.contract.UnpackLog(event, "WithdrawalRewardRecord", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
