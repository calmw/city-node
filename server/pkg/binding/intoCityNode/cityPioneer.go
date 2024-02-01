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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AdminAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AdminRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pioneerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"cityId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ctime\",\"type\":\"uint256\"}],\"name\":\"DailyIncreaseDelegateRecord\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pioneerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toxReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"foundsReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delegateReward\",\"type\":\"uint256\"}],\"name\":\"DailyRewardRecord\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pioneerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toxReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"foundsReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delegateReward\",\"type\":\"uint256\"}],\"name\":\"DailyRewardRecordV2\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pioneerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"month\",\"type\":\"uint256\"}],\"name\":\"SuretyRecord\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pioneerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardType\",\"type\":\"uint256\"}],\"name\":\"WithdrawalRewardRecord\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"TOXAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPioneerAddress_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oldPioneerAddress_\",\"type\":\"address\"}],\"name\":\"adminChangePioneerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allAdmins\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"admins\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"alreadyRewardRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"alreadyRewardRateTotal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"appraise\",\"outputs\":[{\"internalType\":\"contractIAppraise\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"assessmentCriteria\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"assessmentReturnCriteria\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"assessmentReturnRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"amounts\",\"type\":\"address[]\"}],\"name\":\"batchAddAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"benefitPackageReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"benefitPackageRewardReceived\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"benefitPackageRewardStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"checkPioneerDailyStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"checkPioneerStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"checkRewardStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pioneer_\",\"type\":\"address\"}],\"name\":\"checkSurety\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cityAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cityPioneerData\",\"outputs\":[{\"internalType\":\"contractCityPioneerData\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"dailyTaskStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"delegateReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"delegateRewardReceived\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"delegateRewardStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositSurety\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"failedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"failedDelegate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"fundsReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"fundsRewardReceived\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"fundsRewardStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPioneerNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSurety\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pioneer\",\"type\":\"address\"}],\"name\":\"initPioneer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"isPioneer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isPioneerReturnSurety\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"miningAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paySurety\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pioneerInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pioneerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"ctime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cityLevel\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"assessmentMonthStatus\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"assessmentStatus\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"returnSuretyStatus\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"returnSuretyRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"returnSuretyTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"suretyTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pioneerPaySurety\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pioneerAddress_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"chengShiId_\",\"type\":\"bytes32\"}],\"name\":\"pioneerTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pioneers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"presidencyTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pioneer_\",\"type\":\"address\"}],\"name\":\"removePioneer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"secondsPerDay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pioneer_\",\"type\":\"address\"}],\"name\":\"setIsPioneerReturnSurety\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"successAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"successTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"suretyMonthWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"suretyReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"suretyRewardRecord\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"testValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"userLocationAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawalBenefitPackageReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawalDelegateReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawalFundsReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawalSuretyReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// AlreadyRewardRate is a free data retrieval call binding the contract method 0x858db1f3.
//
// Solidity: function alreadyRewardRate(address , uint256 ) view returns(uint256)
func (_CityPioneer *CityPioneerCaller) AlreadyRewardRate(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CityPioneer.contract.Call(opts, &out, "alreadyRewardRate", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AlreadyRewardRate is a free data retrieval call binding the contract method 0x858db1f3.
//
// Solidity: function alreadyRewardRate(address , uint256 ) view returns(uint256)
func (_CityPioneer *CityPioneerSession) AlreadyRewardRate(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _CityPioneer.Contract.AlreadyRewardRate(&_CityPioneer.CallOpts, arg0, arg1)
}

// AlreadyRewardRate is a free data retrieval call binding the contract method 0x858db1f3.
//
// Solidity: function alreadyRewardRate(address , uint256 ) view returns(uint256)
func (_CityPioneer *CityPioneerCallerSession) AlreadyRewardRate(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _CityPioneer.Contract.AlreadyRewardRate(&_CityPioneer.CallOpts, arg0, arg1)
}

// AlreadyRewardRateTotal is a free data retrieval call binding the contract method 0xd87cc67b.
//
// Solidity: function alreadyRewardRateTotal(address ) view returns(uint256)
func (_CityPioneer *CityPioneerCaller) AlreadyRewardRateTotal(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CityPioneer.contract.Call(opts, &out, "alreadyRewardRateTotal", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AlreadyRewardRateTotal is a free data retrieval call binding the contract method 0xd87cc67b.
//
// Solidity: function alreadyRewardRateTotal(address ) view returns(uint256)
func (_CityPioneer *CityPioneerSession) AlreadyRewardRateTotal(arg0 common.Address) (*big.Int, error) {
	return _CityPioneer.Contract.AlreadyRewardRateTotal(&_CityPioneer.CallOpts, arg0)
}

// AlreadyRewardRateTotal is a free data retrieval call binding the contract method 0xd87cc67b.
//
// Solidity: function alreadyRewardRateTotal(address ) view returns(uint256)
func (_CityPioneer *CityPioneerCallerSession) AlreadyRewardRateTotal(arg0 common.Address) (*big.Int, error) {
	return _CityPioneer.Contract.AlreadyRewardRateTotal(&_CityPioneer.CallOpts, arg0)
}

// Appraise is a free data retrieval call binding the contract method 0x144f067d.
//
// Solidity: function appraise() view returns(address)
func (_CityPioneer *CityPioneerCaller) Appraise(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CityPioneer.contract.Call(opts, &out, "appraise")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Appraise is a free data retrieval call binding the contract method 0x144f067d.
//
// Solidity: function appraise() view returns(address)
func (_CityPioneer *CityPioneerSession) Appraise() (common.Address, error) {
	return _CityPioneer.Contract.Appraise(&_CityPioneer.CallOpts)
}

// Appraise is a free data retrieval call binding the contract method 0x144f067d.
//
// Solidity: function appraise() view returns(address)
func (_CityPioneer *CityPioneerCallerSession) Appraise() (common.Address, error) {
	return _CityPioneer.Contract.Appraise(&_CityPioneer.CallOpts)
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

// CheckPioneerDailyStatus is a free data retrieval call binding the contract method 0x4fcc0fd4.
//
// Solidity: function checkPioneerDailyStatus(uint256 , address ) view returns(bool)
func (_CityPioneer *CityPioneerCaller) CheckPioneerDailyStatus(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _CityPioneer.contract.Call(opts, &out, "checkPioneerDailyStatus", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckPioneerDailyStatus is a free data retrieval call binding the contract method 0x4fcc0fd4.
//
// Solidity: function checkPioneerDailyStatus(uint256 , address ) view returns(bool)
func (_CityPioneer *CityPioneerSession) CheckPioneerDailyStatus(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _CityPioneer.Contract.CheckPioneerDailyStatus(&_CityPioneer.CallOpts, arg0, arg1)
}

// CheckPioneerDailyStatus is a free data retrieval call binding the contract method 0x4fcc0fd4.
//
// Solidity: function checkPioneerDailyStatus(uint256 , address ) view returns(bool)
func (_CityPioneer *CityPioneerCallerSession) CheckPioneerDailyStatus(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _CityPioneer.Contract.CheckPioneerDailyStatus(&_CityPioneer.CallOpts, arg0, arg1)
}

// CheckPioneerStatus is a free data retrieval call binding the contract method 0xcb4dfd1d.
//
// Solidity: function checkPioneerStatus(address , uint256 ) view returns(bool)
func (_CityPioneer *CityPioneerCaller) CheckPioneerStatus(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (bool, error) {
	var out []interface{}
	err := _CityPioneer.contract.Call(opts, &out, "checkPioneerStatus", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckPioneerStatus is a free data retrieval call binding the contract method 0xcb4dfd1d.
//
// Solidity: function checkPioneerStatus(address , uint256 ) view returns(bool)
func (_CityPioneer *CityPioneerSession) CheckPioneerStatus(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _CityPioneer.Contract.CheckPioneerStatus(&_CityPioneer.CallOpts, arg0, arg1)
}

// CheckPioneerStatus is a free data retrieval call binding the contract method 0xcb4dfd1d.
//
// Solidity: function checkPioneerStatus(address , uint256 ) view returns(bool)
func (_CityPioneer *CityPioneerCallerSession) CheckPioneerStatus(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _CityPioneer.Contract.CheckPioneerStatus(&_CityPioneer.CallOpts, arg0, arg1)
}

// CheckRewardStatus is a free data retrieval call binding the contract method 0x9d3a649b.
//
// Solidity: function checkRewardStatus(address , uint256 ) view returns(bool)
func (_CityPioneer *CityPioneerCaller) CheckRewardStatus(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (bool, error) {
	var out []interface{}
	err := _CityPioneer.contract.Call(opts, &out, "checkRewardStatus", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckRewardStatus is a free data retrieval call binding the contract method 0x9d3a649b.
//
// Solidity: function checkRewardStatus(address , uint256 ) view returns(bool)
func (_CityPioneer *CityPioneerSession) CheckRewardStatus(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _CityPioneer.Contract.CheckRewardStatus(&_CityPioneer.CallOpts, arg0, arg1)
}

// CheckRewardStatus is a free data retrieval call binding the contract method 0x9d3a649b.
//
// Solidity: function checkRewardStatus(address , uint256 ) view returns(bool)
func (_CityPioneer *CityPioneerCallerSession) CheckRewardStatus(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _CityPioneer.Contract.CheckRewardStatus(&_CityPioneer.CallOpts, arg0, arg1)
}

// CheckSurety is a free data retrieval call binding the contract method 0x7a933435.
//
// Solidity: function checkSurety(address pioneer_) view returns(uint256)
func (_CityPioneer *CityPioneerCaller) CheckSurety(opts *bind.CallOpts, pioneer_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CityPioneer.contract.Call(opts, &out, "checkSurety", pioneer_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CheckSurety is a free data retrieval call binding the contract method 0x7a933435.
//
// Solidity: function checkSurety(address pioneer_) view returns(uint256)
func (_CityPioneer *CityPioneerSession) CheckSurety(pioneer_ common.Address) (*big.Int, error) {
	return _CityPioneer.Contract.CheckSurety(&_CityPioneer.CallOpts, pioneer_)
}

// CheckSurety is a free data retrieval call binding the contract method 0x7a933435.
//
// Solidity: function checkSurety(address pioneer_) view returns(uint256)
func (_CityPioneer *CityPioneerCallerSession) CheckSurety(pioneer_ common.Address) (*big.Int, error) {
	return _CityPioneer.Contract.CheckSurety(&_CityPioneer.CallOpts, pioneer_)
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

// CityPioneerData is a free data retrieval call binding the contract method 0x483d0611.
//
// Solidity: function cityPioneerData() view returns(address)
func (_CityPioneer *CityPioneerCaller) CityPioneerData(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CityPioneer.contract.Call(opts, &out, "cityPioneerData")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CityPioneerData is a free data retrieval call binding the contract method 0x483d0611.
//
// Solidity: function cityPioneerData() view returns(address)
func (_CityPioneer *CityPioneerSession) CityPioneerData() (common.Address, error) {
	return _CityPioneer.Contract.CityPioneerData(&_CityPioneer.CallOpts)
}

// CityPioneerData is a free data retrieval call binding the contract method 0x483d0611.
//
// Solidity: function cityPioneerData() view returns(address)
func (_CityPioneer *CityPioneerCallerSession) CityPioneerData() (common.Address, error) {
	return _CityPioneer.Contract.CityPioneerData(&_CityPioneer.CallOpts)
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

// FailedAt is a free data retrieval call binding the contract method 0x0af98575.
//
// Solidity: function failedAt(address ) view returns(uint256)
func (_CityPioneer *CityPioneerCaller) FailedAt(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CityPioneer.contract.Call(opts, &out, "failedAt", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FailedAt is a free data retrieval call binding the contract method 0x0af98575.
//
// Solidity: function failedAt(address ) view returns(uint256)
func (_CityPioneer *CityPioneerSession) FailedAt(arg0 common.Address) (*big.Int, error) {
	return _CityPioneer.Contract.FailedAt(&_CityPioneer.CallOpts, arg0)
}

// FailedAt is a free data retrieval call binding the contract method 0x0af98575.
//
// Solidity: function failedAt(address ) view returns(uint256)
func (_CityPioneer *CityPioneerCallerSession) FailedAt(arg0 common.Address) (*big.Int, error) {
	return _CityPioneer.Contract.FailedAt(&_CityPioneer.CallOpts, arg0)
}

// FailedDelegate is a free data retrieval call binding the contract method 0x69b93a36.
//
// Solidity: function failedDelegate(bytes32 ) view returns(uint256)
func (_CityPioneer *CityPioneerCaller) FailedDelegate(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _CityPioneer.contract.Call(opts, &out, "failedDelegate", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FailedDelegate is a free data retrieval call binding the contract method 0x69b93a36.
//
// Solidity: function failedDelegate(bytes32 ) view returns(uint256)
func (_CityPioneer *CityPioneerSession) FailedDelegate(arg0 [32]byte) (*big.Int, error) {
	return _CityPioneer.Contract.FailedDelegate(&_CityPioneer.CallOpts, arg0)
}

// FailedDelegate is a free data retrieval call binding the contract method 0x69b93a36.
//
// Solidity: function failedDelegate(bytes32 ) view returns(uint256)
func (_CityPioneer *CityPioneerCallerSession) FailedDelegate(arg0 [32]byte) (*big.Int, error) {
	return _CityPioneer.Contract.FailedDelegate(&_CityPioneer.CallOpts, arg0)
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

// GetPioneerNumber is a free data retrieval call binding the contract method 0xc2325b84.
//
// Solidity: function getPioneerNumber() view returns(uint256)
func (_CityPioneer *CityPioneerCaller) GetPioneerNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CityPioneer.contract.Call(opts, &out, "getPioneerNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPioneerNumber is a free data retrieval call binding the contract method 0xc2325b84.
//
// Solidity: function getPioneerNumber() view returns(uint256)
func (_CityPioneer *CityPioneerSession) GetPioneerNumber() (*big.Int, error) {
	return _CityPioneer.Contract.GetPioneerNumber(&_CityPioneer.CallOpts)
}

// GetPioneerNumber is a free data retrieval call binding the contract method 0xc2325b84.
//
// Solidity: function getPioneerNumber() view returns(uint256)
func (_CityPioneer *CityPioneerCallerSession) GetPioneerNumber() (*big.Int, error) {
	return _CityPioneer.Contract.GetPioneerNumber(&_CityPioneer.CallOpts)
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

// IsPioneerReturnSurety is a free data retrieval call binding the contract method 0x23f3d945.
//
// Solidity: function isPioneerReturnSurety(address ) view returns(bool)
func (_CityPioneer *CityPioneerCaller) IsPioneerReturnSurety(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _CityPioneer.contract.Call(opts, &out, "isPioneerReturnSurety", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPioneerReturnSurety is a free data retrieval call binding the contract method 0x23f3d945.
//
// Solidity: function isPioneerReturnSurety(address ) view returns(bool)
func (_CityPioneer *CityPioneerSession) IsPioneerReturnSurety(arg0 common.Address) (bool, error) {
	return _CityPioneer.Contract.IsPioneerReturnSurety(&_CityPioneer.CallOpts, arg0)
}

// IsPioneerReturnSurety is a free data retrieval call binding the contract method 0x23f3d945.
//
// Solidity: function isPioneerReturnSurety(address ) view returns(bool)
func (_CityPioneer *CityPioneerCallerSession) IsPioneerReturnSurety(arg0 common.Address) (bool, error) {
	return _CityPioneer.Contract.IsPioneerReturnSurety(&_CityPioneer.CallOpts, arg0)
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
// Solidity: function pioneerInfo(address ) view returns(address pioneerAddress, uint256 ctime, uint256 cityLevel, bool assessmentMonthStatus, bool assessmentStatus, bool returnSuretyStatus, uint256 returnSuretyRate, uint256 returnSuretyTime, uint256 suretyTime)
func (_CityPioneer *CityPioneerCaller) PioneerInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	PioneerAddress        common.Address
	Ctime                 *big.Int
	CityLevel             *big.Int
	AssessmentMonthStatus bool
	AssessmentStatus      bool
	ReturnSuretyStatus    bool
	ReturnSuretyRate      *big.Int
	ReturnSuretyTime      *big.Int
	SuretyTime            *big.Int
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
		SuretyTime            *big.Int
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
	outstruct.SuretyTime = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PioneerInfo is a free data retrieval call binding the contract method 0x0980a959.
//
// Solidity: function pioneerInfo(address ) view returns(address pioneerAddress, uint256 ctime, uint256 cityLevel, bool assessmentMonthStatus, bool assessmentStatus, bool returnSuretyStatus, uint256 returnSuretyRate, uint256 returnSuretyTime, uint256 suretyTime)
func (_CityPioneer *CityPioneerSession) PioneerInfo(arg0 common.Address) (struct {
	PioneerAddress        common.Address
	Ctime                 *big.Int
	CityLevel             *big.Int
	AssessmentMonthStatus bool
	AssessmentStatus      bool
	ReturnSuretyStatus    bool
	ReturnSuretyRate      *big.Int
	ReturnSuretyTime      *big.Int
	SuretyTime            *big.Int
}, error) {
	return _CityPioneer.Contract.PioneerInfo(&_CityPioneer.CallOpts, arg0)
}

// PioneerInfo is a free data retrieval call binding the contract method 0x0980a959.
//
// Solidity: function pioneerInfo(address ) view returns(address pioneerAddress, uint256 ctime, uint256 cityLevel, bool assessmentMonthStatus, bool assessmentStatus, bool returnSuretyStatus, uint256 returnSuretyRate, uint256 returnSuretyTime, uint256 suretyTime)
func (_CityPioneer *CityPioneerCallerSession) PioneerInfo(arg0 common.Address) (struct {
	PioneerAddress        common.Address
	Ctime                 *big.Int
	CityLevel             *big.Int
	AssessmentMonthStatus bool
	AssessmentStatus      bool
	ReturnSuretyStatus    bool
	ReturnSuretyRate      *big.Int
	ReturnSuretyTime      *big.Int
	SuretyTime            *big.Int
}, error) {
	return _CityPioneer.Contract.PioneerInfo(&_CityPioneer.CallOpts, arg0)
}

// PioneerPaySurety is a free data retrieval call binding the contract method 0x1fcedfd0.
//
// Solidity: function pioneerPaySurety(address ) view returns(uint256)
func (_CityPioneer *CityPioneerCaller) PioneerPaySurety(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CityPioneer.contract.Call(opts, &out, "pioneerPaySurety", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PioneerPaySurety is a free data retrieval call binding the contract method 0x1fcedfd0.
//
// Solidity: function pioneerPaySurety(address ) view returns(uint256)
func (_CityPioneer *CityPioneerSession) PioneerPaySurety(arg0 common.Address) (*big.Int, error) {
	return _CityPioneer.Contract.PioneerPaySurety(&_CityPioneer.CallOpts, arg0)
}

// PioneerPaySurety is a free data retrieval call binding the contract method 0x1fcedfd0.
//
// Solidity: function pioneerPaySurety(address ) view returns(uint256)
func (_CityPioneer *CityPioneerCallerSession) PioneerPaySurety(arg0 common.Address) (*big.Int, error) {
	return _CityPioneer.Contract.PioneerPaySurety(&_CityPioneer.CallOpts, arg0)
}

// Pioneers is a free data retrieval call binding the contract method 0xf98775fe.
//
// Solidity: function pioneers(uint256 ) view returns(address)
func (_CityPioneer *CityPioneerCaller) Pioneers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CityPioneer.contract.Call(opts, &out, "pioneers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pioneers is a free data retrieval call binding the contract method 0xf98775fe.
//
// Solidity: function pioneers(uint256 ) view returns(address)
func (_CityPioneer *CityPioneerSession) Pioneers(arg0 *big.Int) (common.Address, error) {
	return _CityPioneer.Contract.Pioneers(&_CityPioneer.CallOpts, arg0)
}

// Pioneers is a free data retrieval call binding the contract method 0xf98775fe.
//
// Solidity: function pioneers(uint256 ) view returns(address)
func (_CityPioneer *CityPioneerCallerSession) Pioneers(arg0 *big.Int) (common.Address, error) {
	return _CityPioneer.Contract.Pioneers(&_CityPioneer.CallOpts, arg0)
}

// PresidencyTime is a free data retrieval call binding the contract method 0xf47510e1.
//
// Solidity: function presidencyTime() view returns(uint256)
func (_CityPioneer *CityPioneerCaller) PresidencyTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CityPioneer.contract.Call(opts, &out, "presidencyTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PresidencyTime is a free data retrieval call binding the contract method 0xf47510e1.
//
// Solidity: function presidencyTime() view returns(uint256)
func (_CityPioneer *CityPioneerSession) PresidencyTime() (*big.Int, error) {
	return _CityPioneer.Contract.PresidencyTime(&_CityPioneer.CallOpts)
}

// PresidencyTime is a free data retrieval call binding the contract method 0xf47510e1.
//
// Solidity: function presidencyTime() view returns(uint256)
func (_CityPioneer *CityPioneerCallerSession) PresidencyTime() (*big.Int, error) {
	return _CityPioneer.Contract.PresidencyTime(&_CityPioneer.CallOpts)
}

// SecondsPerDay is a free data retrieval call binding the contract method 0x63809953.
//
// Solidity: function secondsPerDay() view returns(uint256)
func (_CityPioneer *CityPioneerCaller) SecondsPerDay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CityPioneer.contract.Call(opts, &out, "secondsPerDay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SecondsPerDay is a free data retrieval call binding the contract method 0x63809953.
//
// Solidity: function secondsPerDay() view returns(uint256)
func (_CityPioneer *CityPioneerSession) SecondsPerDay() (*big.Int, error) {
	return _CityPioneer.Contract.SecondsPerDay(&_CityPioneer.CallOpts)
}

// SecondsPerDay is a free data retrieval call binding the contract method 0x63809953.
//
// Solidity: function secondsPerDay() view returns(uint256)
func (_CityPioneer *CityPioneerCallerSession) SecondsPerDay() (*big.Int, error) {
	return _CityPioneer.Contract.SecondsPerDay(&_CityPioneer.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_CityPioneer *CityPioneerCaller) StartTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CityPioneer.contract.Call(opts, &out, "startTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_CityPioneer *CityPioneerSession) StartTime() (*big.Int, error) {
	return _CityPioneer.Contract.StartTime(&_CityPioneer.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_CityPioneer *CityPioneerCallerSession) StartTime() (*big.Int, error) {
	return _CityPioneer.Contract.StartTime(&_CityPioneer.CallOpts)
}

// SuccessAt is a free data retrieval call binding the contract method 0xbe7e0aaa.
//
// Solidity: function successAt(address , uint256 ) view returns(uint256)
func (_CityPioneer *CityPioneerCaller) SuccessAt(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CityPioneer.contract.Call(opts, &out, "successAt", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SuccessAt is a free data retrieval call binding the contract method 0xbe7e0aaa.
//
// Solidity: function successAt(address , uint256 ) view returns(uint256)
func (_CityPioneer *CityPioneerSession) SuccessAt(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _CityPioneer.Contract.SuccessAt(&_CityPioneer.CallOpts, arg0, arg1)
}

// SuccessAt is a free data retrieval call binding the contract method 0xbe7e0aaa.
//
// Solidity: function successAt(address , uint256 ) view returns(uint256)
func (_CityPioneer *CityPioneerCallerSession) SuccessAt(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _CityPioneer.Contract.SuccessAt(&_CityPioneer.CallOpts, arg0, arg1)
}

// SuccessTime is a free data retrieval call binding the contract method 0x622b9a9f.
//
// Solidity: function successTime(address ) view returns(uint256)
func (_CityPioneer *CityPioneerCaller) SuccessTime(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CityPioneer.contract.Call(opts, &out, "successTime", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SuccessTime is a free data retrieval call binding the contract method 0x622b9a9f.
//
// Solidity: function successTime(address ) view returns(uint256)
func (_CityPioneer *CityPioneerSession) SuccessTime(arg0 common.Address) (*big.Int, error) {
	return _CityPioneer.Contract.SuccessTime(&_CityPioneer.CallOpts, arg0)
}

// SuccessTime is a free data retrieval call binding the contract method 0x622b9a9f.
//
// Solidity: function successTime(address ) view returns(uint256)
func (_CityPioneer *CityPioneerCallerSession) SuccessTime(arg0 common.Address) (*big.Int, error) {
	return _CityPioneer.Contract.SuccessTime(&_CityPioneer.CallOpts, arg0)
}

// SuretyMonthWeight is a free data retrieval call binding the contract method 0xf9c270bf.
//
// Solidity: function suretyMonthWeight(address , uint256 ) view returns(uint256)
func (_CityPioneer *CityPioneerCaller) SuretyMonthWeight(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CityPioneer.contract.Call(opts, &out, "suretyMonthWeight", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SuretyMonthWeight is a free data retrieval call binding the contract method 0xf9c270bf.
//
// Solidity: function suretyMonthWeight(address , uint256 ) view returns(uint256)
func (_CityPioneer *CityPioneerSession) SuretyMonthWeight(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _CityPioneer.Contract.SuretyMonthWeight(&_CityPioneer.CallOpts, arg0, arg1)
}

// SuretyMonthWeight is a free data retrieval call binding the contract method 0xf9c270bf.
//
// Solidity: function suretyMonthWeight(address , uint256 ) view returns(uint256)
func (_CityPioneer *CityPioneerCallerSession) SuretyMonthWeight(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _CityPioneer.Contract.SuretyMonthWeight(&_CityPioneer.CallOpts, arg0, arg1)
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

// TestValue is a free data retrieval call binding the contract method 0x121739e3.
//
// Solidity: function testValue(address , uint256 ) view returns(uint256)
func (_CityPioneer *CityPioneerCaller) TestValue(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CityPioneer.contract.Call(opts, &out, "testValue", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TestValue is a free data retrieval call binding the contract method 0x121739e3.
//
// Solidity: function testValue(address , uint256 ) view returns(uint256)
func (_CityPioneer *CityPioneerSession) TestValue(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _CityPioneer.Contract.TestValue(&_CityPioneer.CallOpts, arg0, arg1)
}

// TestValue is a free data retrieval call binding the contract method 0x121739e3.
//
// Solidity: function testValue(address , uint256 ) view returns(uint256)
func (_CityPioneer *CityPioneerCallerSession) TestValue(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _CityPioneer.Contract.TestValue(&_CityPioneer.CallOpts, arg0, arg1)
}

// UserLocationAddress is a free data retrieval call binding the contract method 0x12516fc6.
//
// Solidity: function userLocationAddress() view returns(address)
func (_CityPioneer *CityPioneerCaller) UserLocationAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CityPioneer.contract.Call(opts, &out, "userLocationAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UserLocationAddress is a free data retrieval call binding the contract method 0x12516fc6.
//
// Solidity: function userLocationAddress() view returns(address)
func (_CityPioneer *CityPioneerSession) UserLocationAddress() (common.Address, error) {
	return _CityPioneer.Contract.UserLocationAddress(&_CityPioneer.CallOpts)
}

// UserLocationAddress is a free data retrieval call binding the contract method 0x12516fc6.
//
// Solidity: function userLocationAddress() view returns(address)
func (_CityPioneer *CityPioneerCallerSession) UserLocationAddress() (common.Address, error) {
	return _CityPioneer.Contract.UserLocationAddress(&_CityPioneer.CallOpts)
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

// AdminChangePioneerAddress is a paid mutator transaction binding the contract method 0x9016179e.
//
// Solidity: function adminChangePioneerAddress(address newPioneerAddress_, address oldPioneerAddress_) returns()
func (_CityPioneer *CityPioneerTransactor) AdminChangePioneerAddress(opts *bind.TransactOpts, newPioneerAddress_ common.Address, oldPioneerAddress_ common.Address) (*types.Transaction, error) {
	return _CityPioneer.contract.Transact(opts, "adminChangePioneerAddress", newPioneerAddress_, oldPioneerAddress_)
}

// AdminChangePioneerAddress is a paid mutator transaction binding the contract method 0x9016179e.
//
// Solidity: function adminChangePioneerAddress(address newPioneerAddress_, address oldPioneerAddress_) returns()
func (_CityPioneer *CityPioneerSession) AdminChangePioneerAddress(newPioneerAddress_ common.Address, oldPioneerAddress_ common.Address) (*types.Transaction, error) {
	return _CityPioneer.Contract.AdminChangePioneerAddress(&_CityPioneer.TransactOpts, newPioneerAddress_, oldPioneerAddress_)
}

// AdminChangePioneerAddress is a paid mutator transaction binding the contract method 0x9016179e.
//
// Solidity: function adminChangePioneerAddress(address newPioneerAddress_, address oldPioneerAddress_) returns()
func (_CityPioneer *CityPioneerTransactorSession) AdminChangePioneerAddress(newPioneerAddress_ common.Address, oldPioneerAddress_ common.Address) (*types.Transaction, error) {
	return _CityPioneer.Contract.AdminChangePioneerAddress(&_CityPioneer.TransactOpts, newPioneerAddress_, oldPioneerAddress_)
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
// Solidity: function depositSurety() returns()
func (_CityPioneer *CityPioneerTransactor) DepositSurety(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CityPioneer.contract.Transact(opts, "depositSurety")
}

// DepositSurety is a paid mutator transaction binding the contract method 0x07edb9c4.
//
// Solidity: function depositSurety() returns()
func (_CityPioneer *CityPioneerSession) DepositSurety() (*types.Transaction, error) {
	return _CityPioneer.Contract.DepositSurety(&_CityPioneer.TransactOpts)
}

// DepositSurety is a paid mutator transaction binding the contract method 0x07edb9c4.
//
// Solidity: function depositSurety() returns()
func (_CityPioneer *CityPioneerTransactorSession) DepositSurety() (*types.Transaction, error) {
	return _CityPioneer.Contract.DepositSurety(&_CityPioneer.TransactOpts)
}

// InitPioneer is a paid mutator transaction binding the contract method 0x31622c13.
//
// Solidity: function initPioneer(address pioneer) returns()
func (_CityPioneer *CityPioneerTransactor) InitPioneer(opts *bind.TransactOpts, pioneer common.Address) (*types.Transaction, error) {
	return _CityPioneer.contract.Transact(opts, "initPioneer", pioneer)
}

// InitPioneer is a paid mutator transaction binding the contract method 0x31622c13.
//
// Solidity: function initPioneer(address pioneer) returns()
func (_CityPioneer *CityPioneerSession) InitPioneer(pioneer common.Address) (*types.Transaction, error) {
	return _CityPioneer.Contract.InitPioneer(&_CityPioneer.TransactOpts, pioneer)
}

// InitPioneer is a paid mutator transaction binding the contract method 0x31622c13.
//
// Solidity: function initPioneer(address pioneer) returns()
func (_CityPioneer *CityPioneerTransactorSession) InitPioneer(pioneer common.Address) (*types.Transaction, error) {
	return _CityPioneer.Contract.InitPioneer(&_CityPioneer.TransactOpts, pioneer)
}

// PaySurety is a paid mutator transaction binding the contract method 0xf59f7b77.
//
// Solidity: function paySurety() returns()
func (_CityPioneer *CityPioneerTransactor) PaySurety(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CityPioneer.contract.Transact(opts, "paySurety")
}

// PaySurety is a paid mutator transaction binding the contract method 0xf59f7b77.
//
// Solidity: function paySurety() returns()
func (_CityPioneer *CityPioneerSession) PaySurety() (*types.Transaction, error) {
	return _CityPioneer.Contract.PaySurety(&_CityPioneer.TransactOpts)
}

// PaySurety is a paid mutator transaction binding the contract method 0xf59f7b77.
//
// Solidity: function paySurety() returns()
func (_CityPioneer *CityPioneerTransactorSession) PaySurety() (*types.Transaction, error) {
	return _CityPioneer.Contract.PaySurety(&_CityPioneer.TransactOpts)
}

// PioneerTask is a paid mutator transaction binding the contract method 0x3bfdbd39.
//
// Solidity: function pioneerTask(address pioneerAddress_, bytes32 chengShiId_) returns()
func (_CityPioneer *CityPioneerTransactor) PioneerTask(opts *bind.TransactOpts, pioneerAddress_ common.Address, chengShiId_ [32]byte) (*types.Transaction, error) {
	return _CityPioneer.contract.Transact(opts, "pioneerTask", pioneerAddress_, chengShiId_)
}

// PioneerTask is a paid mutator transaction binding the contract method 0x3bfdbd39.
//
// Solidity: function pioneerTask(address pioneerAddress_, bytes32 chengShiId_) returns()
func (_CityPioneer *CityPioneerSession) PioneerTask(pioneerAddress_ common.Address, chengShiId_ [32]byte) (*types.Transaction, error) {
	return _CityPioneer.Contract.PioneerTask(&_CityPioneer.TransactOpts, pioneerAddress_, chengShiId_)
}

// PioneerTask is a paid mutator transaction binding the contract method 0x3bfdbd39.
//
// Solidity: function pioneerTask(address pioneerAddress_, bytes32 chengShiId_) returns()
func (_CityPioneer *CityPioneerTransactorSession) PioneerTask(pioneerAddress_ common.Address, chengShiId_ [32]byte) (*types.Transaction, error) {
	return _CityPioneer.Contract.PioneerTask(&_CityPioneer.TransactOpts, pioneerAddress_, chengShiId_)
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

// RemovePioneer is a paid mutator transaction binding the contract method 0x06de5967.
//
// Solidity: function removePioneer(address pioneer_) returns()
func (_CityPioneer *CityPioneerTransactor) RemovePioneer(opts *bind.TransactOpts, pioneer_ common.Address) (*types.Transaction, error) {
	return _CityPioneer.contract.Transact(opts, "removePioneer", pioneer_)
}

// RemovePioneer is a paid mutator transaction binding the contract method 0x06de5967.
//
// Solidity: function removePioneer(address pioneer_) returns()
func (_CityPioneer *CityPioneerSession) RemovePioneer(pioneer_ common.Address) (*types.Transaction, error) {
	return _CityPioneer.Contract.RemovePioneer(&_CityPioneer.TransactOpts, pioneer_)
}

// RemovePioneer is a paid mutator transaction binding the contract method 0x06de5967.
//
// Solidity: function removePioneer(address pioneer_) returns()
func (_CityPioneer *CityPioneerTransactorSession) RemovePioneer(pioneer_ common.Address) (*types.Transaction, error) {
	return _CityPioneer.Contract.RemovePioneer(&_CityPioneer.TransactOpts, pioneer_)
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

// SetIsPioneerReturnSurety is a paid mutator transaction binding the contract method 0x6328866f.
//
// Solidity: function setIsPioneerReturnSurety(address pioneer_) returns()
func (_CityPioneer *CityPioneerTransactor) SetIsPioneerReturnSurety(opts *bind.TransactOpts, pioneer_ common.Address) (*types.Transaction, error) {
	return _CityPioneer.contract.Transact(opts, "setIsPioneerReturnSurety", pioneer_)
}

// SetIsPioneerReturnSurety is a paid mutator transaction binding the contract method 0x6328866f.
//
// Solidity: function setIsPioneerReturnSurety(address pioneer_) returns()
func (_CityPioneer *CityPioneerSession) SetIsPioneerReturnSurety(pioneer_ common.Address) (*types.Transaction, error) {
	return _CityPioneer.Contract.SetIsPioneerReturnSurety(&_CityPioneer.TransactOpts, pioneer_)
}

// SetIsPioneerReturnSurety is a paid mutator transaction binding the contract method 0x6328866f.
//
// Solidity: function setIsPioneerReturnSurety(address pioneer_) returns()
func (_CityPioneer *CityPioneerTransactorSession) SetIsPioneerReturnSurety(pioneer_ common.Address) (*types.Transaction, error) {
	return _CityPioneer.Contract.SetIsPioneerReturnSurety(&_CityPioneer.TransactOpts, pioneer_)
}

// WithdrawalBenefitPackageReward is a paid mutator transaction binding the contract method 0x335822b4.
//
// Solidity: function withdrawalBenefitPackageReward() returns()
func (_CityPioneer *CityPioneerTransactor) WithdrawalBenefitPackageReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CityPioneer.contract.Transact(opts, "withdrawalBenefitPackageReward")
}

// WithdrawalBenefitPackageReward is a paid mutator transaction binding the contract method 0x335822b4.
//
// Solidity: function withdrawalBenefitPackageReward() returns()
func (_CityPioneer *CityPioneerSession) WithdrawalBenefitPackageReward() (*types.Transaction, error) {
	return _CityPioneer.Contract.WithdrawalBenefitPackageReward(&_CityPioneer.TransactOpts)
}

// WithdrawalBenefitPackageReward is a paid mutator transaction binding the contract method 0x335822b4.
//
// Solidity: function withdrawalBenefitPackageReward() returns()
func (_CityPioneer *CityPioneerTransactorSession) WithdrawalBenefitPackageReward() (*types.Transaction, error) {
	return _CityPioneer.Contract.WithdrawalBenefitPackageReward(&_CityPioneer.TransactOpts)
}

// WithdrawalDelegateReward is a paid mutator transaction binding the contract method 0xd459bc5b.
//
// Solidity: function withdrawalDelegateReward() returns()
func (_CityPioneer *CityPioneerTransactor) WithdrawalDelegateReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CityPioneer.contract.Transact(opts, "withdrawalDelegateReward")
}

// WithdrawalDelegateReward is a paid mutator transaction binding the contract method 0xd459bc5b.
//
// Solidity: function withdrawalDelegateReward() returns()
func (_CityPioneer *CityPioneerSession) WithdrawalDelegateReward() (*types.Transaction, error) {
	return _CityPioneer.Contract.WithdrawalDelegateReward(&_CityPioneer.TransactOpts)
}

// WithdrawalDelegateReward is a paid mutator transaction binding the contract method 0xd459bc5b.
//
// Solidity: function withdrawalDelegateReward() returns()
func (_CityPioneer *CityPioneerTransactorSession) WithdrawalDelegateReward() (*types.Transaction, error) {
	return _CityPioneer.Contract.WithdrawalDelegateReward(&_CityPioneer.TransactOpts)
}

// WithdrawalFundsReward is a paid mutator transaction binding the contract method 0xf5dc805c.
//
// Solidity: function withdrawalFundsReward() returns()
func (_CityPioneer *CityPioneerTransactor) WithdrawalFundsReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CityPioneer.contract.Transact(opts, "withdrawalFundsReward")
}

// WithdrawalFundsReward is a paid mutator transaction binding the contract method 0xf5dc805c.
//
// Solidity: function withdrawalFundsReward() returns()
func (_CityPioneer *CityPioneerSession) WithdrawalFundsReward() (*types.Transaction, error) {
	return _CityPioneer.Contract.WithdrawalFundsReward(&_CityPioneer.TransactOpts)
}

// WithdrawalFundsReward is a paid mutator transaction binding the contract method 0xf5dc805c.
//
// Solidity: function withdrawalFundsReward() returns()
func (_CityPioneer *CityPioneerTransactorSession) WithdrawalFundsReward() (*types.Transaction, error) {
	return _CityPioneer.Contract.WithdrawalFundsReward(&_CityPioneer.TransactOpts)
}

// WithdrawalSuretyReward is a paid mutator transaction binding the contract method 0x92d962ec.
//
// Solidity: function withdrawalSuretyReward() returns()
func (_CityPioneer *CityPioneerTransactor) WithdrawalSuretyReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CityPioneer.contract.Transact(opts, "withdrawalSuretyReward")
}

// WithdrawalSuretyReward is a paid mutator transaction binding the contract method 0x92d962ec.
//
// Solidity: function withdrawalSuretyReward() returns()
func (_CityPioneer *CityPioneerSession) WithdrawalSuretyReward() (*types.Transaction, error) {
	return _CityPioneer.Contract.WithdrawalSuretyReward(&_CityPioneer.TransactOpts)
}

// WithdrawalSuretyReward is a paid mutator transaction binding the contract method 0x92d962ec.
//
// Solidity: function withdrawalSuretyReward() returns()
func (_CityPioneer *CityPioneerTransactorSession) WithdrawalSuretyReward() (*types.Transaction, error) {
	return _CityPioneer.Contract.WithdrawalSuretyReward(&_CityPioneer.TransactOpts)
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

// CityPioneerDailyRewardRecordV2Iterator is returned from FilterDailyRewardRecordV2 and is used to iterate over the raw logs and unpacked data for DailyRewardRecordV2 events raised by the CityPioneer contract.
type CityPioneerDailyRewardRecordV2Iterator struct {
	Event *CityPioneerDailyRewardRecordV2 // Event containing the contract specifics and raw log

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
func (it *CityPioneerDailyRewardRecordV2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CityPioneerDailyRewardRecordV2)
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
		it.Event = new(CityPioneerDailyRewardRecordV2)
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
func (it *CityPioneerDailyRewardRecordV2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CityPioneerDailyRewardRecordV2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CityPioneerDailyRewardRecordV2 represents a DailyRewardRecordV2 event raised by the CityPioneer contract.
type CityPioneerDailyRewardRecordV2 struct {
	PioneerAddress common.Address
	ToxReward      *big.Int
	FoundsReward   *big.Int
	DelegateReward *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDailyRewardRecordV2 is a free log retrieval operation binding the contract event 0x593c84866e560a2b3b81272108b1158935911ccd81692fac808cd1acc9b6cdea.
//
// Solidity: event DailyRewardRecordV2(address pioneerAddress, uint256 toxReward, uint256 foundsReward, uint256 delegateReward)
func (_CityPioneer *CityPioneerFilterer) FilterDailyRewardRecordV2(opts *bind.FilterOpts) (*CityPioneerDailyRewardRecordV2Iterator, error) {

	logs, sub, err := _CityPioneer.contract.FilterLogs(opts, "DailyRewardRecordV2")
	if err != nil {
		return nil, err
	}
	return &CityPioneerDailyRewardRecordV2Iterator{contract: _CityPioneer.contract, event: "DailyRewardRecordV2", logs: logs, sub: sub}, nil
}

// WatchDailyRewardRecordV2 is a free log subscription operation binding the contract event 0x593c84866e560a2b3b81272108b1158935911ccd81692fac808cd1acc9b6cdea.
//
// Solidity: event DailyRewardRecordV2(address pioneerAddress, uint256 toxReward, uint256 foundsReward, uint256 delegateReward)
func (_CityPioneer *CityPioneerFilterer) WatchDailyRewardRecordV2(opts *bind.WatchOpts, sink chan<- *CityPioneerDailyRewardRecordV2) (event.Subscription, error) {

	logs, sub, err := _CityPioneer.contract.WatchLogs(opts, "DailyRewardRecordV2")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CityPioneerDailyRewardRecordV2)
				if err := _CityPioneer.contract.UnpackLog(event, "DailyRewardRecordV2", log); err != nil {
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

// ParseDailyRewardRecordV2 is a log parse operation binding the contract event 0x593c84866e560a2b3b81272108b1158935911ccd81692fac808cd1acc9b6cdea.
//
// Solidity: event DailyRewardRecordV2(address pioneerAddress, uint256 toxReward, uint256 foundsReward, uint256 delegateReward)
func (_CityPioneer *CityPioneerFilterer) ParseDailyRewardRecordV2(log types.Log) (*CityPioneerDailyRewardRecordV2, error) {
	event := new(CityPioneerDailyRewardRecordV2)
	if err := _CityPioneer.contract.UnpackLog(event, "DailyRewardRecordV2", log); err != nil {
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
	Month          *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSuretyRecord is a free log retrieval operation binding the contract event 0x4d1b08d9b6dee0ccfd52569e86af5fca351aa92061ee3333f18bcbd3ec6f2713.
//
// Solidity: event SuretyRecord(address pioneerAddress, uint256 amount, uint256 month)
func (_CityPioneer *CityPioneerFilterer) FilterSuretyRecord(opts *bind.FilterOpts) (*CityPioneerSuretyRecordIterator, error) {

	logs, sub, err := _CityPioneer.contract.FilterLogs(opts, "SuretyRecord")
	if err != nil {
		return nil, err
	}
	return &CityPioneerSuretyRecordIterator{contract: _CityPioneer.contract, event: "SuretyRecord", logs: logs, sub: sub}, nil
}

// WatchSuretyRecord is a free log subscription operation binding the contract event 0x4d1b08d9b6dee0ccfd52569e86af5fca351aa92061ee3333f18bcbd3ec6f2713.
//
// Solidity: event SuretyRecord(address pioneerAddress, uint256 amount, uint256 month)
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

// ParseSuretyRecord is a log parse operation binding the contract event 0x4d1b08d9b6dee0ccfd52569e86af5fca351aa92061ee3333f18bcbd3ec6f2713.
//
// Solidity: event SuretyRecord(address pioneerAddress, uint256 amount, uint256 month)
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
