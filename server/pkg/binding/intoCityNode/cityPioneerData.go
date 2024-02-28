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

// CityPioneerDataMetaData contains all meta data concerning the CityPioneerData contract.
var CityPioneerDataMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AdminAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AdminRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"chengShiId_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"suretyUSDT_\",\"type\":\"uint256\"}],\"name\":\"adminSetChengShiSuretyUSDT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cityAddress_\",\"type\":\"address\"}],\"name\":\"adminSetCityAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cityPioneerAddress_\",\"type\":\"address\"}],\"name\":\"adminSetCityPioneerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"globalNodeAddress_\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isGlobalNode_\",\"type\":\"bool\"}],\"name\":\"adminSetIsGlobalNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"reward_\",\"type\":\"uint256\"}],\"name\":\"adminSetSubReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"toxAddress_\",\"type\":\"address\"}],\"name\":\"adminSetTOX\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usdtAddress_\",\"type\":\"address\"}],\"name\":\"adminSetUSDT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"reward_\",\"type\":\"uint256\"}],\"name\":\"adminSubReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allAdmins\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"admins\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"amounts\",\"type\":\"address[]\"}],\"name\":\"batchAddAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"city\",\"outputs\":[{\"internalType\":\"contractCity\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cityContract\",\"outputs\":[{\"internalType\":\"contractICity\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cityPioneer\",\"outputs\":[{\"internalType\":\"contractCityPioneer\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cityPioneerContract\",\"outputs\":[{\"internalType\":\"contractICityPioneer\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositSuretyTOX\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositSuretyUSDT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getSurety\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isGlobalNode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"subReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"suretyDepositTOX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"suretyDepositUSDT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"suretyUSDT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tox\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdt\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// CityPioneerDataABI is the input ABI used to generate the binding from.
// Deprecated: Use CityPioneerDataMetaData.ABI instead.
var CityPioneerDataABI = CityPioneerDataMetaData.ABI

// CityPioneerData is an auto generated Go binding around an Ethereum contract.
type CityPioneerData struct {
	CityPioneerDataCaller     // Read-only binding to the contract
	CityPioneerDataTransactor // Write-only binding to the contract
	CityPioneerDataFilterer   // Log filterer for contract events
}

// CityPioneerDataCaller is an auto generated read-only Go binding around an Ethereum contract.
type CityPioneerDataCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CityPioneerDataTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CityPioneerDataTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CityPioneerDataFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CityPioneerDataFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CityPioneerDataSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CityPioneerDataSession struct {
	Contract     *CityPioneerData  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CityPioneerDataCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CityPioneerDataCallerSession struct {
	Contract *CityPioneerDataCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// CityPioneerDataTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CityPioneerDataTransactorSession struct {
	Contract     *CityPioneerDataTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// CityPioneerDataRaw is an auto generated low-level Go binding around an Ethereum contract.
type CityPioneerDataRaw struct {
	Contract *CityPioneerData // Generic contract binding to access the raw methods on
}

// CityPioneerDataCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CityPioneerDataCallerRaw struct {
	Contract *CityPioneerDataCaller // Generic read-only contract binding to access the raw methods on
}

// CityPioneerDataTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CityPioneerDataTransactorRaw struct {
	Contract *CityPioneerDataTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCityPioneerData creates a new instance of CityPioneerData, bound to a specific deployed contract.
func NewCityPioneerData(address common.Address, backend bind.ContractBackend) (*CityPioneerData, error) {
	contract, err := bindCityPioneerData(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CityPioneerData{CityPioneerDataCaller: CityPioneerDataCaller{contract: contract}, CityPioneerDataTransactor: CityPioneerDataTransactor{contract: contract}, CityPioneerDataFilterer: CityPioneerDataFilterer{contract: contract}}, nil
}

// NewCityPioneerDataCaller creates a new read-only instance of CityPioneerData, bound to a specific deployed contract.
func NewCityPioneerDataCaller(address common.Address, caller bind.ContractCaller) (*CityPioneerDataCaller, error) {
	contract, err := bindCityPioneerData(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CityPioneerDataCaller{contract: contract}, nil
}

// NewCityPioneerDataTransactor creates a new write-only instance of CityPioneerData, bound to a specific deployed contract.
func NewCityPioneerDataTransactor(address common.Address, transactor bind.ContractTransactor) (*CityPioneerDataTransactor, error) {
	contract, err := bindCityPioneerData(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CityPioneerDataTransactor{contract: contract}, nil
}

// NewCityPioneerDataFilterer creates a new log filterer instance of CityPioneerData, bound to a specific deployed contract.
func NewCityPioneerDataFilterer(address common.Address, filterer bind.ContractFilterer) (*CityPioneerDataFilterer, error) {
	contract, err := bindCityPioneerData(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CityPioneerDataFilterer{contract: contract}, nil
}

// bindCityPioneerData binds a generic wrapper to an already deployed contract.
func bindCityPioneerData(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CityPioneerDataMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CityPioneerData *CityPioneerDataRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CityPioneerData.Contract.CityPioneerDataCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CityPioneerData *CityPioneerDataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CityPioneerData.Contract.CityPioneerDataTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CityPioneerData *CityPioneerDataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CityPioneerData.Contract.CityPioneerDataTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CityPioneerData *CityPioneerDataCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CityPioneerData.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CityPioneerData *CityPioneerDataTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CityPioneerData.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CityPioneerData *CityPioneerDataTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CityPioneerData.Contract.contract.Transact(opts, method, params...)
}

// AllAdmins is a free data retrieval call binding the contract method 0x40f32be6.
//
// Solidity: function allAdmins() view returns(address[] admins)
func (_CityPioneerData *CityPioneerDataCaller) AllAdmins(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _CityPioneerData.contract.Call(opts, &out, "allAdmins")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// AllAdmins is a free data retrieval call binding the contract method 0x40f32be6.
//
// Solidity: function allAdmins() view returns(address[] admins)
func (_CityPioneerData *CityPioneerDataSession) AllAdmins() ([]common.Address, error) {
	return _CityPioneerData.Contract.AllAdmins(&_CityPioneerData.CallOpts)
}

// AllAdmins is a free data retrieval call binding the contract method 0x40f32be6.
//
// Solidity: function allAdmins() view returns(address[] admins)
func (_CityPioneerData *CityPioneerDataCallerSession) AllAdmins() ([]common.Address, error) {
	return _CityPioneerData.Contract.AllAdmins(&_CityPioneerData.CallOpts)
}

// City is a free data retrieval call binding the contract method 0xfd0160fe.
//
// Solidity: function city() view returns(address)
func (_CityPioneerData *CityPioneerDataCaller) City(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CityPioneerData.contract.Call(opts, &out, "city")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// City is a free data retrieval call binding the contract method 0xfd0160fe.
//
// Solidity: function city() view returns(address)
func (_CityPioneerData *CityPioneerDataSession) City() (common.Address, error) {
	return _CityPioneerData.Contract.City(&_CityPioneerData.CallOpts)
}

// City is a free data retrieval call binding the contract method 0xfd0160fe.
//
// Solidity: function city() view returns(address)
func (_CityPioneerData *CityPioneerDataCallerSession) City() (common.Address, error) {
	return _CityPioneerData.Contract.City(&_CityPioneerData.CallOpts)
}

// CityContract is a free data retrieval call binding the contract method 0x6c99f57b.
//
// Solidity: function cityContract() view returns(address)
func (_CityPioneerData *CityPioneerDataCaller) CityContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CityPioneerData.contract.Call(opts, &out, "cityContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CityContract is a free data retrieval call binding the contract method 0x6c99f57b.
//
// Solidity: function cityContract() view returns(address)
func (_CityPioneerData *CityPioneerDataSession) CityContract() (common.Address, error) {
	return _CityPioneerData.Contract.CityContract(&_CityPioneerData.CallOpts)
}

// CityContract is a free data retrieval call binding the contract method 0x6c99f57b.
//
// Solidity: function cityContract() view returns(address)
func (_CityPioneerData *CityPioneerDataCallerSession) CityContract() (common.Address, error) {
	return _CityPioneerData.Contract.CityContract(&_CityPioneerData.CallOpts)
}

// CityPioneer is a free data retrieval call binding the contract method 0xa5bcc3b0.
//
// Solidity: function cityPioneer() view returns(address)
func (_CityPioneerData *CityPioneerDataCaller) CityPioneer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CityPioneerData.contract.Call(opts, &out, "cityPioneer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CityPioneer is a free data retrieval call binding the contract method 0xa5bcc3b0.
//
// Solidity: function cityPioneer() view returns(address)
func (_CityPioneerData *CityPioneerDataSession) CityPioneer() (common.Address, error) {
	return _CityPioneerData.Contract.CityPioneer(&_CityPioneerData.CallOpts)
}

// CityPioneer is a free data retrieval call binding the contract method 0xa5bcc3b0.
//
// Solidity: function cityPioneer() view returns(address)
func (_CityPioneerData *CityPioneerDataCallerSession) CityPioneer() (common.Address, error) {
	return _CityPioneerData.Contract.CityPioneer(&_CityPioneerData.CallOpts)
}

// CityPioneerContract is a free data retrieval call binding the contract method 0x4aa89d25.
//
// Solidity: function cityPioneerContract() view returns(address)
func (_CityPioneerData *CityPioneerDataCaller) CityPioneerContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CityPioneerData.contract.Call(opts, &out, "cityPioneerContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CityPioneerContract is a free data retrieval call binding the contract method 0x4aa89d25.
//
// Solidity: function cityPioneerContract() view returns(address)
func (_CityPioneerData *CityPioneerDataSession) CityPioneerContract() (common.Address, error) {
	return _CityPioneerData.Contract.CityPioneerContract(&_CityPioneerData.CallOpts)
}

// CityPioneerContract is a free data retrieval call binding the contract method 0x4aa89d25.
//
// Solidity: function cityPioneerContract() view returns(address)
func (_CityPioneerData *CityPioneerDataCallerSession) CityPioneerContract() (common.Address, error) {
	return _CityPioneerData.Contract.CityPioneerContract(&_CityPioneerData.CallOpts)
}

// GetSurety is a free data retrieval call binding the contract method 0x6fc5c56d.
//
// Solidity: function getSurety(address user) view returns(uint256, uint256)
func (_CityPioneerData *CityPioneerDataCaller) GetSurety(opts *bind.CallOpts, user common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _CityPioneerData.contract.Call(opts, &out, "getSurety", user)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetSurety is a free data retrieval call binding the contract method 0x6fc5c56d.
//
// Solidity: function getSurety(address user) view returns(uint256, uint256)
func (_CityPioneerData *CityPioneerDataSession) GetSurety(user common.Address) (*big.Int, *big.Int, error) {
	return _CityPioneerData.Contract.GetSurety(&_CityPioneerData.CallOpts, user)
}

// GetSurety is a free data retrieval call binding the contract method 0x6fc5c56d.
//
// Solidity: function getSurety(address user) view returns(uint256, uint256)
func (_CityPioneerData *CityPioneerDataCallerSession) GetSurety(user common.Address) (*big.Int, *big.Int, error) {
	return _CityPioneerData.Contract.GetSurety(&_CityPioneerData.CallOpts, user)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address account) view returns(bool)
func (_CityPioneerData *CityPioneerDataCaller) IsAdmin(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _CityPioneerData.contract.Call(opts, &out, "isAdmin", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address account) view returns(bool)
func (_CityPioneerData *CityPioneerDataSession) IsAdmin(account common.Address) (bool, error) {
	return _CityPioneerData.Contract.IsAdmin(&_CityPioneerData.CallOpts, account)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address account) view returns(bool)
func (_CityPioneerData *CityPioneerDataCallerSession) IsAdmin(account common.Address) (bool, error) {
	return _CityPioneerData.Contract.IsAdmin(&_CityPioneerData.CallOpts, account)
}

// IsGlobalNode is a free data retrieval call binding the contract method 0x724fbdb6.
//
// Solidity: function isGlobalNode(address ) view returns(bool)
func (_CityPioneerData *CityPioneerDataCaller) IsGlobalNode(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _CityPioneerData.contract.Call(opts, &out, "isGlobalNode", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsGlobalNode is a free data retrieval call binding the contract method 0x724fbdb6.
//
// Solidity: function isGlobalNode(address ) view returns(bool)
func (_CityPioneerData *CityPioneerDataSession) IsGlobalNode(arg0 common.Address) (bool, error) {
	return _CityPioneerData.Contract.IsGlobalNode(&_CityPioneerData.CallOpts, arg0)
}

// IsGlobalNode is a free data retrieval call binding the contract method 0x724fbdb6.
//
// Solidity: function isGlobalNode(address ) view returns(bool)
func (_CityPioneerData *CityPioneerDataCallerSession) IsGlobalNode(arg0 common.Address) (bool, error) {
	return _CityPioneerData.Contract.IsGlobalNode(&_CityPioneerData.CallOpts, arg0)
}

// SubReward is a free data retrieval call binding the contract method 0xd092918f.
//
// Solidity: function subReward(address ) view returns(uint256)
func (_CityPioneerData *CityPioneerDataCaller) SubReward(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CityPioneerData.contract.Call(opts, &out, "subReward", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SubReward is a free data retrieval call binding the contract method 0xd092918f.
//
// Solidity: function subReward(address ) view returns(uint256)
func (_CityPioneerData *CityPioneerDataSession) SubReward(arg0 common.Address) (*big.Int, error) {
	return _CityPioneerData.Contract.SubReward(&_CityPioneerData.CallOpts, arg0)
}

// SubReward is a free data retrieval call binding the contract method 0xd092918f.
//
// Solidity: function subReward(address ) view returns(uint256)
func (_CityPioneerData *CityPioneerDataCallerSession) SubReward(arg0 common.Address) (*big.Int, error) {
	return _CityPioneerData.Contract.SubReward(&_CityPioneerData.CallOpts, arg0)
}

// SuretyDepositTOX is a free data retrieval call binding the contract method 0x61e5789a.
//
// Solidity: function suretyDepositTOX(address ) view returns(uint256)
func (_CityPioneerData *CityPioneerDataCaller) SuretyDepositTOX(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CityPioneerData.contract.Call(opts, &out, "suretyDepositTOX", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SuretyDepositTOX is a free data retrieval call binding the contract method 0x61e5789a.
//
// Solidity: function suretyDepositTOX(address ) view returns(uint256)
func (_CityPioneerData *CityPioneerDataSession) SuretyDepositTOX(arg0 common.Address) (*big.Int, error) {
	return _CityPioneerData.Contract.SuretyDepositTOX(&_CityPioneerData.CallOpts, arg0)
}

// SuretyDepositTOX is a free data retrieval call binding the contract method 0x61e5789a.
//
// Solidity: function suretyDepositTOX(address ) view returns(uint256)
func (_CityPioneerData *CityPioneerDataCallerSession) SuretyDepositTOX(arg0 common.Address) (*big.Int, error) {
	return _CityPioneerData.Contract.SuretyDepositTOX(&_CityPioneerData.CallOpts, arg0)
}

// SuretyDepositUSDT is a free data retrieval call binding the contract method 0xfeae9d8a.
//
// Solidity: function suretyDepositUSDT(address ) view returns(uint256)
func (_CityPioneerData *CityPioneerDataCaller) SuretyDepositUSDT(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CityPioneerData.contract.Call(opts, &out, "suretyDepositUSDT", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SuretyDepositUSDT is a free data retrieval call binding the contract method 0xfeae9d8a.
//
// Solidity: function suretyDepositUSDT(address ) view returns(uint256)
func (_CityPioneerData *CityPioneerDataSession) SuretyDepositUSDT(arg0 common.Address) (*big.Int, error) {
	return _CityPioneerData.Contract.SuretyDepositUSDT(&_CityPioneerData.CallOpts, arg0)
}

// SuretyDepositUSDT is a free data retrieval call binding the contract method 0xfeae9d8a.
//
// Solidity: function suretyDepositUSDT(address ) view returns(uint256)
func (_CityPioneerData *CityPioneerDataCallerSession) SuretyDepositUSDT(arg0 common.Address) (*big.Int, error) {
	return _CityPioneerData.Contract.SuretyDepositUSDT(&_CityPioneerData.CallOpts, arg0)
}

// SuretyUSDT is a free data retrieval call binding the contract method 0xdf56ff5d.
//
// Solidity: function suretyUSDT(bytes32 ) view returns(uint256)
func (_CityPioneerData *CityPioneerDataCaller) SuretyUSDT(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _CityPioneerData.contract.Call(opts, &out, "suretyUSDT", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SuretyUSDT is a free data retrieval call binding the contract method 0xdf56ff5d.
//
// Solidity: function suretyUSDT(bytes32 ) view returns(uint256)
func (_CityPioneerData *CityPioneerDataSession) SuretyUSDT(arg0 [32]byte) (*big.Int, error) {
	return _CityPioneerData.Contract.SuretyUSDT(&_CityPioneerData.CallOpts, arg0)
}

// SuretyUSDT is a free data retrieval call binding the contract method 0xdf56ff5d.
//
// Solidity: function suretyUSDT(bytes32 ) view returns(uint256)
func (_CityPioneerData *CityPioneerDataCallerSession) SuretyUSDT(arg0 [32]byte) (*big.Int, error) {
	return _CityPioneerData.Contract.SuretyUSDT(&_CityPioneerData.CallOpts, arg0)
}

// Tox is a free data retrieval call binding the contract method 0x46c811fd.
//
// Solidity: function tox() view returns(address)
func (_CityPioneerData *CityPioneerDataCaller) Tox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CityPioneerData.contract.Call(opts, &out, "tox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Tox is a free data retrieval call binding the contract method 0x46c811fd.
//
// Solidity: function tox() view returns(address)
func (_CityPioneerData *CityPioneerDataSession) Tox() (common.Address, error) {
	return _CityPioneerData.Contract.Tox(&_CityPioneerData.CallOpts)
}

// Tox is a free data retrieval call binding the contract method 0x46c811fd.
//
// Solidity: function tox() view returns(address)
func (_CityPioneerData *CityPioneerDataCallerSession) Tox() (common.Address, error) {
	return _CityPioneerData.Contract.Tox(&_CityPioneerData.CallOpts)
}

// Usdt is a free data retrieval call binding the contract method 0x2f48ab7d.
//
// Solidity: function usdt() view returns(address)
func (_CityPioneerData *CityPioneerDataCaller) Usdt(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CityPioneerData.contract.Call(opts, &out, "usdt")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Usdt is a free data retrieval call binding the contract method 0x2f48ab7d.
//
// Solidity: function usdt() view returns(address)
func (_CityPioneerData *CityPioneerDataSession) Usdt() (common.Address, error) {
	return _CityPioneerData.Contract.Usdt(&_CityPioneerData.CallOpts)
}

// Usdt is a free data retrieval call binding the contract method 0x2f48ab7d.
//
// Solidity: function usdt() view returns(address)
func (_CityPioneerData *CityPioneerDataCallerSession) Usdt() (common.Address, error) {
	return _CityPioneerData.Contract.Usdt(&_CityPioneerData.CallOpts)
}

// UserPeriod is a free data retrieval call binding the contract method 0x159ed917.
//
// Solidity: function userPeriod(address ) view returns(uint256)
func (_CityPioneerData *CityPioneerDataCaller) UserPeriod(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CityPioneerData.contract.Call(opts, &out, "userPeriod", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserPeriod is a free data retrieval call binding the contract method 0x159ed917.
//
// Solidity: function userPeriod(address ) view returns(uint256)
func (_CityPioneerData *CityPioneerDataSession) UserPeriod(arg0 common.Address) (*big.Int, error) {
	return _CityPioneerData.Contract.UserPeriod(&_CityPioneerData.CallOpts, arg0)
}

// UserPeriod is a free data retrieval call binding the contract method 0x159ed917.
//
// Solidity: function userPeriod(address ) view returns(uint256)
func (_CityPioneerData *CityPioneerDataCallerSession) UserPeriod(arg0 common.Address) (*big.Int, error) {
	return _CityPioneerData.Contract.UserPeriod(&_CityPioneerData.CallOpts, arg0)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address account) returns()
func (_CityPioneerData *CityPioneerDataTransactor) AddAdmin(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _CityPioneerData.contract.Transact(opts, "addAdmin", account)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address account) returns()
func (_CityPioneerData *CityPioneerDataSession) AddAdmin(account common.Address) (*types.Transaction, error) {
	return _CityPioneerData.Contract.AddAdmin(&_CityPioneerData.TransactOpts, account)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address account) returns()
func (_CityPioneerData *CityPioneerDataTransactorSession) AddAdmin(account common.Address) (*types.Transaction, error) {
	return _CityPioneerData.Contract.AddAdmin(&_CityPioneerData.TransactOpts, account)
}

// AdminSetChengShiSuretyUSDT is a paid mutator transaction binding the contract method 0xe95e21b6.
//
// Solidity: function adminSetChengShiSuretyUSDT(bytes32 chengShiId_, uint256 suretyUSDT_) returns()
func (_CityPioneerData *CityPioneerDataTransactor) AdminSetChengShiSuretyUSDT(opts *bind.TransactOpts, chengShiId_ [32]byte, suretyUSDT_ *big.Int) (*types.Transaction, error) {
	return _CityPioneerData.contract.Transact(opts, "adminSetChengShiSuretyUSDT", chengShiId_, suretyUSDT_)
}

// AdminSetChengShiSuretyUSDT is a paid mutator transaction binding the contract method 0xe95e21b6.
//
// Solidity: function adminSetChengShiSuretyUSDT(bytes32 chengShiId_, uint256 suretyUSDT_) returns()
func (_CityPioneerData *CityPioneerDataSession) AdminSetChengShiSuretyUSDT(chengShiId_ [32]byte, suretyUSDT_ *big.Int) (*types.Transaction, error) {
	return _CityPioneerData.Contract.AdminSetChengShiSuretyUSDT(&_CityPioneerData.TransactOpts, chengShiId_, suretyUSDT_)
}

// AdminSetChengShiSuretyUSDT is a paid mutator transaction binding the contract method 0xe95e21b6.
//
// Solidity: function adminSetChengShiSuretyUSDT(bytes32 chengShiId_, uint256 suretyUSDT_) returns()
func (_CityPioneerData *CityPioneerDataTransactorSession) AdminSetChengShiSuretyUSDT(chengShiId_ [32]byte, suretyUSDT_ *big.Int) (*types.Transaction, error) {
	return _CityPioneerData.Contract.AdminSetChengShiSuretyUSDT(&_CityPioneerData.TransactOpts, chengShiId_, suretyUSDT_)
}

// AdminSetCityAddress is a paid mutator transaction binding the contract method 0xa94c843d.
//
// Solidity: function adminSetCityAddress(address cityAddress_) returns()
func (_CityPioneerData *CityPioneerDataTransactor) AdminSetCityAddress(opts *bind.TransactOpts, cityAddress_ common.Address) (*types.Transaction, error) {
	return _CityPioneerData.contract.Transact(opts, "adminSetCityAddress", cityAddress_)
}

// AdminSetCityAddress is a paid mutator transaction binding the contract method 0xa94c843d.
//
// Solidity: function adminSetCityAddress(address cityAddress_) returns()
func (_CityPioneerData *CityPioneerDataSession) AdminSetCityAddress(cityAddress_ common.Address) (*types.Transaction, error) {
	return _CityPioneerData.Contract.AdminSetCityAddress(&_CityPioneerData.TransactOpts, cityAddress_)
}

// AdminSetCityAddress is a paid mutator transaction binding the contract method 0xa94c843d.
//
// Solidity: function adminSetCityAddress(address cityAddress_) returns()
func (_CityPioneerData *CityPioneerDataTransactorSession) AdminSetCityAddress(cityAddress_ common.Address) (*types.Transaction, error) {
	return _CityPioneerData.Contract.AdminSetCityAddress(&_CityPioneerData.TransactOpts, cityAddress_)
}

// AdminSetCityPioneerAddress is a paid mutator transaction binding the contract method 0xa1532e85.
//
// Solidity: function adminSetCityPioneerAddress(address cityPioneerAddress_) returns()
func (_CityPioneerData *CityPioneerDataTransactor) AdminSetCityPioneerAddress(opts *bind.TransactOpts, cityPioneerAddress_ common.Address) (*types.Transaction, error) {
	return _CityPioneerData.contract.Transact(opts, "adminSetCityPioneerAddress", cityPioneerAddress_)
}

// AdminSetCityPioneerAddress is a paid mutator transaction binding the contract method 0xa1532e85.
//
// Solidity: function adminSetCityPioneerAddress(address cityPioneerAddress_) returns()
func (_CityPioneerData *CityPioneerDataSession) AdminSetCityPioneerAddress(cityPioneerAddress_ common.Address) (*types.Transaction, error) {
	return _CityPioneerData.Contract.AdminSetCityPioneerAddress(&_CityPioneerData.TransactOpts, cityPioneerAddress_)
}

// AdminSetCityPioneerAddress is a paid mutator transaction binding the contract method 0xa1532e85.
//
// Solidity: function adminSetCityPioneerAddress(address cityPioneerAddress_) returns()
func (_CityPioneerData *CityPioneerDataTransactorSession) AdminSetCityPioneerAddress(cityPioneerAddress_ common.Address) (*types.Transaction, error) {
	return _CityPioneerData.Contract.AdminSetCityPioneerAddress(&_CityPioneerData.TransactOpts, cityPioneerAddress_)
}

// AdminSetIsGlobalNode is a paid mutator transaction binding the contract method 0x65f0f793.
//
// Solidity: function adminSetIsGlobalNode(address globalNodeAddress_, bool isGlobalNode_) returns()
func (_CityPioneerData *CityPioneerDataTransactor) AdminSetIsGlobalNode(opts *bind.TransactOpts, globalNodeAddress_ common.Address, isGlobalNode_ bool) (*types.Transaction, error) {
	return _CityPioneerData.contract.Transact(opts, "adminSetIsGlobalNode", globalNodeAddress_, isGlobalNode_)
}

// AdminSetIsGlobalNode is a paid mutator transaction binding the contract method 0x65f0f793.
//
// Solidity: function adminSetIsGlobalNode(address globalNodeAddress_, bool isGlobalNode_) returns()
func (_CityPioneerData *CityPioneerDataSession) AdminSetIsGlobalNode(globalNodeAddress_ common.Address, isGlobalNode_ bool) (*types.Transaction, error) {
	return _CityPioneerData.Contract.AdminSetIsGlobalNode(&_CityPioneerData.TransactOpts, globalNodeAddress_, isGlobalNode_)
}

// AdminSetIsGlobalNode is a paid mutator transaction binding the contract method 0x65f0f793.
//
// Solidity: function adminSetIsGlobalNode(address globalNodeAddress_, bool isGlobalNode_) returns()
func (_CityPioneerData *CityPioneerDataTransactorSession) AdminSetIsGlobalNode(globalNodeAddress_ common.Address, isGlobalNode_ bool) (*types.Transaction, error) {
	return _CityPioneerData.Contract.AdminSetIsGlobalNode(&_CityPioneerData.TransactOpts, globalNodeAddress_, isGlobalNode_)
}

// AdminSetSubReward is a paid mutator transaction binding the contract method 0x57e7b3b0.
//
// Solidity: function adminSetSubReward(address user_, uint256 reward_) returns()
func (_CityPioneerData *CityPioneerDataTransactor) AdminSetSubReward(opts *bind.TransactOpts, user_ common.Address, reward_ *big.Int) (*types.Transaction, error) {
	return _CityPioneerData.contract.Transact(opts, "adminSetSubReward", user_, reward_)
}

// AdminSetSubReward is a paid mutator transaction binding the contract method 0x57e7b3b0.
//
// Solidity: function adminSetSubReward(address user_, uint256 reward_) returns()
func (_CityPioneerData *CityPioneerDataSession) AdminSetSubReward(user_ common.Address, reward_ *big.Int) (*types.Transaction, error) {
	return _CityPioneerData.Contract.AdminSetSubReward(&_CityPioneerData.TransactOpts, user_, reward_)
}

// AdminSetSubReward is a paid mutator transaction binding the contract method 0x57e7b3b0.
//
// Solidity: function adminSetSubReward(address user_, uint256 reward_) returns()
func (_CityPioneerData *CityPioneerDataTransactorSession) AdminSetSubReward(user_ common.Address, reward_ *big.Int) (*types.Transaction, error) {
	return _CityPioneerData.Contract.AdminSetSubReward(&_CityPioneerData.TransactOpts, user_, reward_)
}

// AdminSetTOX is a paid mutator transaction binding the contract method 0x0697161b.
//
// Solidity: function adminSetTOX(address toxAddress_) returns()
func (_CityPioneerData *CityPioneerDataTransactor) AdminSetTOX(opts *bind.TransactOpts, toxAddress_ common.Address) (*types.Transaction, error) {
	return _CityPioneerData.contract.Transact(opts, "adminSetTOX", toxAddress_)
}

// AdminSetTOX is a paid mutator transaction binding the contract method 0x0697161b.
//
// Solidity: function adminSetTOX(address toxAddress_) returns()
func (_CityPioneerData *CityPioneerDataSession) AdminSetTOX(toxAddress_ common.Address) (*types.Transaction, error) {
	return _CityPioneerData.Contract.AdminSetTOX(&_CityPioneerData.TransactOpts, toxAddress_)
}

// AdminSetTOX is a paid mutator transaction binding the contract method 0x0697161b.
//
// Solidity: function adminSetTOX(address toxAddress_) returns()
func (_CityPioneerData *CityPioneerDataTransactorSession) AdminSetTOX(toxAddress_ common.Address) (*types.Transaction, error) {
	return _CityPioneerData.Contract.AdminSetTOX(&_CityPioneerData.TransactOpts, toxAddress_)
}

// AdminSetUSDT is a paid mutator transaction binding the contract method 0x95bdbb74.
//
// Solidity: function adminSetUSDT(address usdtAddress_) returns()
func (_CityPioneerData *CityPioneerDataTransactor) AdminSetUSDT(opts *bind.TransactOpts, usdtAddress_ common.Address) (*types.Transaction, error) {
	return _CityPioneerData.contract.Transact(opts, "adminSetUSDT", usdtAddress_)
}

// AdminSetUSDT is a paid mutator transaction binding the contract method 0x95bdbb74.
//
// Solidity: function adminSetUSDT(address usdtAddress_) returns()
func (_CityPioneerData *CityPioneerDataSession) AdminSetUSDT(usdtAddress_ common.Address) (*types.Transaction, error) {
	return _CityPioneerData.Contract.AdminSetUSDT(&_CityPioneerData.TransactOpts, usdtAddress_)
}

// AdminSetUSDT is a paid mutator transaction binding the contract method 0x95bdbb74.
//
// Solidity: function adminSetUSDT(address usdtAddress_) returns()
func (_CityPioneerData *CityPioneerDataTransactorSession) AdminSetUSDT(usdtAddress_ common.Address) (*types.Transaction, error) {
	return _CityPioneerData.Contract.AdminSetUSDT(&_CityPioneerData.TransactOpts, usdtAddress_)
}

// AdminSubReward is a paid mutator transaction binding the contract method 0xc5b0363e.
//
// Solidity: function adminSubReward(address user_, uint256 reward_) returns()
func (_CityPioneerData *CityPioneerDataTransactor) AdminSubReward(opts *bind.TransactOpts, user_ common.Address, reward_ *big.Int) (*types.Transaction, error) {
	return _CityPioneerData.contract.Transact(opts, "adminSubReward", user_, reward_)
}

// AdminSubReward is a paid mutator transaction binding the contract method 0xc5b0363e.
//
// Solidity: function adminSubReward(address user_, uint256 reward_) returns()
func (_CityPioneerData *CityPioneerDataSession) AdminSubReward(user_ common.Address, reward_ *big.Int) (*types.Transaction, error) {
	return _CityPioneerData.Contract.AdminSubReward(&_CityPioneerData.TransactOpts, user_, reward_)
}

// AdminSubReward is a paid mutator transaction binding the contract method 0xc5b0363e.
//
// Solidity: function adminSubReward(address user_, uint256 reward_) returns()
func (_CityPioneerData *CityPioneerDataTransactorSession) AdminSubReward(user_ common.Address, reward_ *big.Int) (*types.Transaction, error) {
	return _CityPioneerData.Contract.AdminSubReward(&_CityPioneerData.TransactOpts, user_, reward_)
}

// BatchAddAdmin is a paid mutator transaction binding the contract method 0x2c9ab42b.
//
// Solidity: function batchAddAdmin(address[] amounts) returns()
func (_CityPioneerData *CityPioneerDataTransactor) BatchAddAdmin(opts *bind.TransactOpts, amounts []common.Address) (*types.Transaction, error) {
	return _CityPioneerData.contract.Transact(opts, "batchAddAdmin", amounts)
}

// BatchAddAdmin is a paid mutator transaction binding the contract method 0x2c9ab42b.
//
// Solidity: function batchAddAdmin(address[] amounts) returns()
func (_CityPioneerData *CityPioneerDataSession) BatchAddAdmin(amounts []common.Address) (*types.Transaction, error) {
	return _CityPioneerData.Contract.BatchAddAdmin(&_CityPioneerData.TransactOpts, amounts)
}

// BatchAddAdmin is a paid mutator transaction binding the contract method 0x2c9ab42b.
//
// Solidity: function batchAddAdmin(address[] amounts) returns()
func (_CityPioneerData *CityPioneerDataTransactorSession) BatchAddAdmin(amounts []common.Address) (*types.Transaction, error) {
	return _CityPioneerData.Contract.BatchAddAdmin(&_CityPioneerData.TransactOpts, amounts)
}

// DepositSuretyTOX is a paid mutator transaction binding the contract method 0x61b517e5.
//
// Solidity: function depositSuretyTOX() returns()
func (_CityPioneerData *CityPioneerDataTransactor) DepositSuretyTOX(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CityPioneerData.contract.Transact(opts, "depositSuretyTOX")
}

// DepositSuretyTOX is a paid mutator transaction binding the contract method 0x61b517e5.
//
// Solidity: function depositSuretyTOX() returns()
func (_CityPioneerData *CityPioneerDataSession) DepositSuretyTOX() (*types.Transaction, error) {
	return _CityPioneerData.Contract.DepositSuretyTOX(&_CityPioneerData.TransactOpts)
}

// DepositSuretyTOX is a paid mutator transaction binding the contract method 0x61b517e5.
//
// Solidity: function depositSuretyTOX() returns()
func (_CityPioneerData *CityPioneerDataTransactorSession) DepositSuretyTOX() (*types.Transaction, error) {
	return _CityPioneerData.Contract.DepositSuretyTOX(&_CityPioneerData.TransactOpts)
}

// DepositSuretyUSDT is a paid mutator transaction binding the contract method 0x4dcecc50.
//
// Solidity: function depositSuretyUSDT() returns()
func (_CityPioneerData *CityPioneerDataTransactor) DepositSuretyUSDT(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CityPioneerData.contract.Transact(opts, "depositSuretyUSDT")
}

// DepositSuretyUSDT is a paid mutator transaction binding the contract method 0x4dcecc50.
//
// Solidity: function depositSuretyUSDT() returns()
func (_CityPioneerData *CityPioneerDataSession) DepositSuretyUSDT() (*types.Transaction, error) {
	return _CityPioneerData.Contract.DepositSuretyUSDT(&_CityPioneerData.TransactOpts)
}

// DepositSuretyUSDT is a paid mutator transaction binding the contract method 0x4dcecc50.
//
// Solidity: function depositSuretyUSDT() returns()
func (_CityPioneerData *CityPioneerDataTransactorSession) DepositSuretyUSDT() (*types.Transaction, error) {
	return _CityPioneerData.Contract.DepositSuretyUSDT(&_CityPioneerData.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_CityPioneerData *CityPioneerDataTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CityPioneerData.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_CityPioneerData *CityPioneerDataSession) Initialize() (*types.Transaction, error) {
	return _CityPioneerData.Contract.Initialize(&_CityPioneerData.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_CityPioneerData *CityPioneerDataTransactorSession) Initialize() (*types.Transaction, error) {
	return _CityPioneerData.Contract.Initialize(&_CityPioneerData.TransactOpts)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address account) returns()
func (_CityPioneerData *CityPioneerDataTransactor) RemoveAdmin(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _CityPioneerData.contract.Transact(opts, "removeAdmin", account)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address account) returns()
func (_CityPioneerData *CityPioneerDataSession) RemoveAdmin(account common.Address) (*types.Transaction, error) {
	return _CityPioneerData.Contract.RemoveAdmin(&_CityPioneerData.TransactOpts, account)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address account) returns()
func (_CityPioneerData *CityPioneerDataTransactorSession) RemoveAdmin(account common.Address) (*types.Transaction, error) {
	return _CityPioneerData.Contract.RemoveAdmin(&_CityPioneerData.TransactOpts, account)
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x8bad0c0a.
//
// Solidity: function renounceAdmin() returns()
func (_CityPioneerData *CityPioneerDataTransactor) RenounceAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CityPioneerData.contract.Transact(opts, "renounceAdmin")
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x8bad0c0a.
//
// Solidity: function renounceAdmin() returns()
func (_CityPioneerData *CityPioneerDataSession) RenounceAdmin() (*types.Transaction, error) {
	return _CityPioneerData.Contract.RenounceAdmin(&_CityPioneerData.TransactOpts)
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x8bad0c0a.
//
// Solidity: function renounceAdmin() returns()
func (_CityPioneerData *CityPioneerDataTransactorSession) RenounceAdmin() (*types.Transaction, error) {
	return _CityPioneerData.Contract.RenounceAdmin(&_CityPioneerData.TransactOpts)
}

// CityPioneerDataAdminAddedIterator is returned from FilterAdminAdded and is used to iterate over the raw logs and unpacked data for AdminAdded events raised by the CityPioneerData contract.
type CityPioneerDataAdminAddedIterator struct {
	Event *CityPioneerDataAdminAdded // Event containing the contract specifics and raw log

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
func (it *CityPioneerDataAdminAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CityPioneerDataAdminAdded)
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
		it.Event = new(CityPioneerDataAdminAdded)
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
func (it *CityPioneerDataAdminAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CityPioneerDataAdminAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CityPioneerDataAdminAdded represents a AdminAdded event raised by the CityPioneerData contract.
type CityPioneerDataAdminAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAdminAdded is a free log retrieval operation binding the contract event 0x44d6d25963f097ad14f29f06854a01f575648a1ef82f30e562ccd3889717e339.
//
// Solidity: event AdminAdded(address indexed account)
func (_CityPioneerData *CityPioneerDataFilterer) FilterAdminAdded(opts *bind.FilterOpts, account []common.Address) (*CityPioneerDataAdminAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _CityPioneerData.contract.FilterLogs(opts, "AdminAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &CityPioneerDataAdminAddedIterator{contract: _CityPioneerData.contract, event: "AdminAdded", logs: logs, sub: sub}, nil
}

// WatchAdminAdded is a free log subscription operation binding the contract event 0x44d6d25963f097ad14f29f06854a01f575648a1ef82f30e562ccd3889717e339.
//
// Solidity: event AdminAdded(address indexed account)
func (_CityPioneerData *CityPioneerDataFilterer) WatchAdminAdded(opts *bind.WatchOpts, sink chan<- *CityPioneerDataAdminAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _CityPioneerData.contract.WatchLogs(opts, "AdminAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CityPioneerDataAdminAdded)
				if err := _CityPioneerData.contract.UnpackLog(event, "AdminAdded", log); err != nil {
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
func (_CityPioneerData *CityPioneerDataFilterer) ParseAdminAdded(log types.Log) (*CityPioneerDataAdminAdded, error) {
	event := new(CityPioneerDataAdminAdded)
	if err := _CityPioneerData.contract.UnpackLog(event, "AdminAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CityPioneerDataAdminRemovedIterator is returned from FilterAdminRemoved and is used to iterate over the raw logs and unpacked data for AdminRemoved events raised by the CityPioneerData contract.
type CityPioneerDataAdminRemovedIterator struct {
	Event *CityPioneerDataAdminRemoved // Event containing the contract specifics and raw log

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
func (it *CityPioneerDataAdminRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CityPioneerDataAdminRemoved)
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
		it.Event = new(CityPioneerDataAdminRemoved)
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
func (it *CityPioneerDataAdminRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CityPioneerDataAdminRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CityPioneerDataAdminRemoved represents a AdminRemoved event raised by the CityPioneerData contract.
type CityPioneerDataAdminRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAdminRemoved is a free log retrieval operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: event AdminRemoved(address indexed account)
func (_CityPioneerData *CityPioneerDataFilterer) FilterAdminRemoved(opts *bind.FilterOpts, account []common.Address) (*CityPioneerDataAdminRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _CityPioneerData.contract.FilterLogs(opts, "AdminRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &CityPioneerDataAdminRemovedIterator{contract: _CityPioneerData.contract, event: "AdminRemoved", logs: logs, sub: sub}, nil
}

// WatchAdminRemoved is a free log subscription operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: event AdminRemoved(address indexed account)
func (_CityPioneerData *CityPioneerDataFilterer) WatchAdminRemoved(opts *bind.WatchOpts, sink chan<- *CityPioneerDataAdminRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _CityPioneerData.contract.WatchLogs(opts, "AdminRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CityPioneerDataAdminRemoved)
				if err := _CityPioneerData.contract.UnpackLog(event, "AdminRemoved", log); err != nil {
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
func (_CityPioneerData *CityPioneerDataFilterer) ParseAdminRemoved(log types.Log) (*CityPioneerDataAdminRemoved, error) {
	event := new(CityPioneerDataAdminRemoved)
	if err := _CityPioneerData.contract.UnpackLog(event, "AdminRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CityPioneerDataInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the CityPioneerData contract.
type CityPioneerDataInitializedIterator struct {
	Event *CityPioneerDataInitialized // Event containing the contract specifics and raw log

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
func (it *CityPioneerDataInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CityPioneerDataInitialized)
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
		it.Event = new(CityPioneerDataInitialized)
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
func (it *CityPioneerDataInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CityPioneerDataInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CityPioneerDataInitialized represents a Initialized event raised by the CityPioneerData contract.
type CityPioneerDataInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_CityPioneerData *CityPioneerDataFilterer) FilterInitialized(opts *bind.FilterOpts) (*CityPioneerDataInitializedIterator, error) {

	logs, sub, err := _CityPioneerData.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &CityPioneerDataInitializedIterator{contract: _CityPioneerData.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_CityPioneerData *CityPioneerDataFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *CityPioneerDataInitialized) (event.Subscription, error) {

	logs, sub, err := _CityPioneerData.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CityPioneerDataInitialized)
				if err := _CityPioneerData.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_CityPioneerData *CityPioneerDataFilterer) ParseInitialized(log types.Log) (*CityPioneerDataInitialized, error) {
	event := new(CityPioneerDataInitialized)
	if err := _CityPioneerData.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
