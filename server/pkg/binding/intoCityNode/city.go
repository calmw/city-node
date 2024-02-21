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

// CityMetaData contains all meta data concerning the City contract.
var CityMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AdminAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AdminRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"cityId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DecreaseCityDelegate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"cityId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"IncreaseCityDelegate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pioneer\",\"type\":\"address\"}],\"name\":\"PioneerBlack\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"countyId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"RechargeRecord\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"chengShiId_\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"pioneer_\",\"type\":\"address\"}],\"name\":\"adminChangePioneerCityId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"chengShiId_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"level_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"surety_\",\"type\":\"uint256\"}],\"name\":\"adminEditSurety\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adminPopPioneerChengShiId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"chengShiId_\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"pioneer_\",\"type\":\"address\"}],\"name\":\"adminRemovePioneer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"authAddress_\",\"type\":\"address\"}],\"name\":\"adminSetAuthAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"chengShiId_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"level_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"surety_\",\"type\":\"uint256\"}],\"name\":\"adminSetChengShiLevelAndSurety\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cityPioneerAddress_\",\"type\":\"address\"}],\"name\":\"adminSetCityPioneerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"setType\",\"type\":\"uint256\"}],\"name\":\"adminSetDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"foundsAddress_\",\"type\":\"address\"}],\"name\":\"adminSetFoundsAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"chengShiId_\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"pioneer_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pioneerBatch_\",\"type\":\"uint256\"}],\"name\":\"adminSetPioneer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pioneer_\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"status_\",\"type\":\"bool\"}],\"name\":\"adminSetPioneerStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"adminSetRechargeAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pioneer_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"weight_\",\"type\":\"uint256\"}],\"name\":\"adminSetRechargeWeightAdditional\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint56\",\"name\":\"secondsPerDay_\",\"type\":\"uint56\"}],\"name\":\"adminSetSecondsPerDay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userLocationAddress_\",\"type\":\"address\"}],\"name\":\"adminSetUserLocationAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"withdrawLimitAddress_\",\"type\":\"address\"}],\"name\":\"adminSetWithdrawLimitAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allAdmins\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"admins\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allCityDailyFoundsTotal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allCityDelegateTotal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allCityFoundsTotal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allCityIds\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allDailyFoundsTotal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allFoundsTotal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auth\",\"outputs\":[{\"internalType\":\"contractIAuth\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"authAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"amounts\",\"type\":\"address[]\"}],\"name\":\"batchAddAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPioneerAddress_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oldPioneerAddress_\",\"type\":\"address\"}],\"name\":\"changePioneerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"chengShiLevel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"chengShiLevelSurety\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"chengShiPioneer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"chengShiRechargeWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"cityDelegate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cityDelegateRecord\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"cityDelegateTotal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cityFoundsRecord\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"cityFoundsTotal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"cityLevel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cityMaxDelegate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cityNewlyDelegateRecord\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"cityOrChengShiWeightTotal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"cityPioneer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cityPioneerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"cityPioneerAssessment\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"cityRechargeTotal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"countyNewlyPioneerDelegateRecord\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"dailyTaskStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"countyId_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"today\",\"type\":\"uint256\"}],\"name\":\"descCountyOrChengShiDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"earnestMoney\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"foundsAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"chengShiId\",\"type\":\"bytes32\"}],\"name\":\"getChengShiRechargeWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"chengShiId_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"day\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isCityPioneer\",\"type\":\"bool\"}],\"name\":\"getDailyDelegateByChengShiID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"cityId_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"day\",\"type\":\"uint256\"}],\"name\":\"getDelegate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFifteenDayAverageFounds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"cityId_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"day\",\"type\":\"uint256\"}],\"name\":\"getFounds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPioneerCityNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pioneer_\",\"type\":\"address\"}],\"name\":\"getSuretyByPioneerAddress\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"hasSetPioneer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"countyId_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"today\",\"type\":\"uint256\"}],\"name\":\"incrCountyOrChengShiDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"chengShiId_\",\"type\":\"bytes32\"}],\"name\":\"initCityRechargeWeight\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pioneerChengShi\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"chengShiId\",\"type\":\"bytes32\"}],\"name\":\"pioneerChengShiIdExits\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pioneerChengShiIds\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pioneerCity\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pioneerCityIds\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pioneerAddress\",\"type\":\"address\"}],\"name\":\"pioneerDailyTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pioneerStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rechargeDailyTotalWeightRecord\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rechargeDailyWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rechargeDailyWeightRecord\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rechargeWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rechargeWeightAdditional\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"secondsPerDay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"chengShiId_\",\"type\":\"bytes32\"}],\"name\":\"setChengShiPioneerAssessment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"cityId_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"day_\",\"type\":\"uint256\"}],\"name\":\"setCityMaxDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"surety\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"userLocationAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawLimit\",\"outputs\":[{\"internalType\":\"contractIWithdrawLimit\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawLimitAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// CityABI is the input ABI used to generate the binding from.
// Deprecated: Use CityMetaData.ABI instead.
var CityABI = CityMetaData.ABI

// City is an auto generated Go binding around an Ethereum contract.
type City struct {
	CityCaller     // Read-only binding to the contract
	CityTransactor // Write-only binding to the contract
	CityFilterer   // Log filterer for contract events
}

// CityCaller is an auto generated read-only Go binding around an Ethereum contract.
type CityCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CityTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CityTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CityFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CityFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CitySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CitySession struct {
	Contract     *City             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CityCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CityCallerSession struct {
	Contract *CityCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CityTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CityTransactorSession struct {
	Contract     *CityTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CityRaw is an auto generated low-level Go binding around an Ethereum contract.
type CityRaw struct {
	Contract *City // Generic contract binding to access the raw methods on
}

// CityCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CityCallerRaw struct {
	Contract *CityCaller // Generic read-only contract binding to access the raw methods on
}

// CityTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CityTransactorRaw struct {
	Contract *CityTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCity creates a new instance of City, bound to a specific deployed contract.
func NewCity(address common.Address, backend bind.ContractBackend) (*City, error) {
	contract, err := bindCity(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &City{CityCaller: CityCaller{contract: contract}, CityTransactor: CityTransactor{contract: contract}, CityFilterer: CityFilterer{contract: contract}}, nil
}

// NewCityCaller creates a new read-only instance of City, bound to a specific deployed contract.
func NewCityCaller(address common.Address, caller bind.ContractCaller) (*CityCaller, error) {
	contract, err := bindCity(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CityCaller{contract: contract}, nil
}

// NewCityTransactor creates a new write-only instance of City, bound to a specific deployed contract.
func NewCityTransactor(address common.Address, transactor bind.ContractTransactor) (*CityTransactor, error) {
	contract, err := bindCity(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CityTransactor{contract: contract}, nil
}

// NewCityFilterer creates a new log filterer instance of City, bound to a specific deployed contract.
func NewCityFilterer(address common.Address, filterer bind.ContractFilterer) (*CityFilterer, error) {
	contract, err := bindCity(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CityFilterer{contract: contract}, nil
}

// bindCity binds a generic wrapper to an already deployed contract.
func bindCity(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CityMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_City *CityRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _City.Contract.CityCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_City *CityRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _City.Contract.CityTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_City *CityRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _City.Contract.CityTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_City *CityCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _City.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_City *CityTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _City.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_City *CityTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _City.Contract.contract.Transact(opts, method, params...)
}

// AllAdmins is a free data retrieval call binding the contract method 0x40f32be6.
//
// Solidity: function allAdmins() view returns(address[] admins)
func (_City *CityCaller) AllAdmins(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _City.contract.Call(opts, &out, "allAdmins")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// AllAdmins is a free data retrieval call binding the contract method 0x40f32be6.
//
// Solidity: function allAdmins() view returns(address[] admins)
func (_City *CitySession) AllAdmins() ([]common.Address, error) {
	return _City.Contract.AllAdmins(&_City.CallOpts)
}

// AllAdmins is a free data retrieval call binding the contract method 0x40f32be6.
//
// Solidity: function allAdmins() view returns(address[] admins)
func (_City *CityCallerSession) AllAdmins() ([]common.Address, error) {
	return _City.Contract.AllAdmins(&_City.CallOpts)
}

// AllCityDailyFoundsTotal is a free data retrieval call binding the contract method 0x52c3abc6.
//
// Solidity: function allCityDailyFoundsTotal(uint256 ) view returns(uint256)
func (_City *CityCaller) AllCityDailyFoundsTotal(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _City.contract.Call(opts, &out, "allCityDailyFoundsTotal", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllCityDailyFoundsTotal is a free data retrieval call binding the contract method 0x52c3abc6.
//
// Solidity: function allCityDailyFoundsTotal(uint256 ) view returns(uint256)
func (_City *CitySession) AllCityDailyFoundsTotal(arg0 *big.Int) (*big.Int, error) {
	return _City.Contract.AllCityDailyFoundsTotal(&_City.CallOpts, arg0)
}

// AllCityDailyFoundsTotal is a free data retrieval call binding the contract method 0x52c3abc6.
//
// Solidity: function allCityDailyFoundsTotal(uint256 ) view returns(uint256)
func (_City *CityCallerSession) AllCityDailyFoundsTotal(arg0 *big.Int) (*big.Int, error) {
	return _City.Contract.AllCityDailyFoundsTotal(&_City.CallOpts, arg0)
}

// AllCityDelegateTotal is a free data retrieval call binding the contract method 0xb0bbc130.
//
// Solidity: function allCityDelegateTotal() view returns(uint256)
func (_City *CityCaller) AllCityDelegateTotal(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _City.contract.Call(opts, &out, "allCityDelegateTotal")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllCityDelegateTotal is a free data retrieval call binding the contract method 0xb0bbc130.
//
// Solidity: function allCityDelegateTotal() view returns(uint256)
func (_City *CitySession) AllCityDelegateTotal() (*big.Int, error) {
	return _City.Contract.AllCityDelegateTotal(&_City.CallOpts)
}

// AllCityDelegateTotal is a free data retrieval call binding the contract method 0xb0bbc130.
//
// Solidity: function allCityDelegateTotal() view returns(uint256)
func (_City *CityCallerSession) AllCityDelegateTotal() (*big.Int, error) {
	return _City.Contract.AllCityDelegateTotal(&_City.CallOpts)
}

// AllCityFoundsTotal is a free data retrieval call binding the contract method 0x6088a497.
//
// Solidity: function allCityFoundsTotal() view returns(uint256)
func (_City *CityCaller) AllCityFoundsTotal(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _City.contract.Call(opts, &out, "allCityFoundsTotal")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllCityFoundsTotal is a free data retrieval call binding the contract method 0x6088a497.
//
// Solidity: function allCityFoundsTotal() view returns(uint256)
func (_City *CitySession) AllCityFoundsTotal() (*big.Int, error) {
	return _City.Contract.AllCityFoundsTotal(&_City.CallOpts)
}

// AllCityFoundsTotal is a free data retrieval call binding the contract method 0x6088a497.
//
// Solidity: function allCityFoundsTotal() view returns(uint256)
func (_City *CityCallerSession) AllCityFoundsTotal() (*big.Int, error) {
	return _City.Contract.AllCityFoundsTotal(&_City.CallOpts)
}

// AllCityIds is a free data retrieval call binding the contract method 0x394b2e3d.
//
// Solidity: function allCityIds(uint256 ) view returns(bytes32)
func (_City *CityCaller) AllCityIds(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _City.contract.Call(opts, &out, "allCityIds", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AllCityIds is a free data retrieval call binding the contract method 0x394b2e3d.
//
// Solidity: function allCityIds(uint256 ) view returns(bytes32)
func (_City *CitySession) AllCityIds(arg0 *big.Int) ([32]byte, error) {
	return _City.Contract.AllCityIds(&_City.CallOpts, arg0)
}

// AllCityIds is a free data retrieval call binding the contract method 0x394b2e3d.
//
// Solidity: function allCityIds(uint256 ) view returns(bytes32)
func (_City *CityCallerSession) AllCityIds(arg0 *big.Int) ([32]byte, error) {
	return _City.Contract.AllCityIds(&_City.CallOpts, arg0)
}

// AllDailyFoundsTotal is a free data retrieval call binding the contract method 0x55cb6a61.
//
// Solidity: function allDailyFoundsTotal(uint256 ) view returns(uint256)
func (_City *CityCaller) AllDailyFoundsTotal(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _City.contract.Call(opts, &out, "allDailyFoundsTotal", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllDailyFoundsTotal is a free data retrieval call binding the contract method 0x55cb6a61.
//
// Solidity: function allDailyFoundsTotal(uint256 ) view returns(uint256)
func (_City *CitySession) AllDailyFoundsTotal(arg0 *big.Int) (*big.Int, error) {
	return _City.Contract.AllDailyFoundsTotal(&_City.CallOpts, arg0)
}

// AllDailyFoundsTotal is a free data retrieval call binding the contract method 0x55cb6a61.
//
// Solidity: function allDailyFoundsTotal(uint256 ) view returns(uint256)
func (_City *CityCallerSession) AllDailyFoundsTotal(arg0 *big.Int) (*big.Int, error) {
	return _City.Contract.AllDailyFoundsTotal(&_City.CallOpts, arg0)
}

// AllFoundsTotal is a free data retrieval call binding the contract method 0x32ba48f8.
//
// Solidity: function allFoundsTotal() view returns(uint256)
func (_City *CityCaller) AllFoundsTotal(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _City.contract.Call(opts, &out, "allFoundsTotal")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllFoundsTotal is a free data retrieval call binding the contract method 0x32ba48f8.
//
// Solidity: function allFoundsTotal() view returns(uint256)
func (_City *CitySession) AllFoundsTotal() (*big.Int, error) {
	return _City.Contract.AllFoundsTotal(&_City.CallOpts)
}

// AllFoundsTotal is a free data retrieval call binding the contract method 0x32ba48f8.
//
// Solidity: function allFoundsTotal() view returns(uint256)
func (_City *CityCallerSession) AllFoundsTotal() (*big.Int, error) {
	return _City.Contract.AllFoundsTotal(&_City.CallOpts)
}

// Auth is a free data retrieval call binding the contract method 0xde9375f2.
//
// Solidity: function auth() view returns(address)
func (_City *CityCaller) Auth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _City.contract.Call(opts, &out, "auth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Auth is a free data retrieval call binding the contract method 0xde9375f2.
//
// Solidity: function auth() view returns(address)
func (_City *CitySession) Auth() (common.Address, error) {
	return _City.Contract.Auth(&_City.CallOpts)
}

// Auth is a free data retrieval call binding the contract method 0xde9375f2.
//
// Solidity: function auth() view returns(address)
func (_City *CityCallerSession) Auth() (common.Address, error) {
	return _City.Contract.Auth(&_City.CallOpts)
}

// AuthAddress is a free data retrieval call binding the contract method 0xa64719d7.
//
// Solidity: function authAddress() view returns(address)
func (_City *CityCaller) AuthAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _City.contract.Call(opts, &out, "authAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AuthAddress is a free data retrieval call binding the contract method 0xa64719d7.
//
// Solidity: function authAddress() view returns(address)
func (_City *CitySession) AuthAddress() (common.Address, error) {
	return _City.Contract.AuthAddress(&_City.CallOpts)
}

// AuthAddress is a free data retrieval call binding the contract method 0xa64719d7.
//
// Solidity: function authAddress() view returns(address)
func (_City *CityCallerSession) AuthAddress() (common.Address, error) {
	return _City.Contract.AuthAddress(&_City.CallOpts)
}

// ChengShiLevel is a free data retrieval call binding the contract method 0xd8806373.
//
// Solidity: function chengShiLevel(bytes32 ) view returns(uint256)
func (_City *CityCaller) ChengShiLevel(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _City.contract.Call(opts, &out, "chengShiLevel", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChengShiLevel is a free data retrieval call binding the contract method 0xd8806373.
//
// Solidity: function chengShiLevel(bytes32 ) view returns(uint256)
func (_City *CitySession) ChengShiLevel(arg0 [32]byte) (*big.Int, error) {
	return _City.Contract.ChengShiLevel(&_City.CallOpts, arg0)
}

// ChengShiLevel is a free data retrieval call binding the contract method 0xd8806373.
//
// Solidity: function chengShiLevel(bytes32 ) view returns(uint256)
func (_City *CityCallerSession) ChengShiLevel(arg0 [32]byte) (*big.Int, error) {
	return _City.Contract.ChengShiLevel(&_City.CallOpts, arg0)
}

// ChengShiLevelSurety is a free data retrieval call binding the contract method 0xf50d0835.
//
// Solidity: function chengShiLevelSurety(uint256 ) view returns(uint256)
func (_City *CityCaller) ChengShiLevelSurety(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _City.contract.Call(opts, &out, "chengShiLevelSurety", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChengShiLevelSurety is a free data retrieval call binding the contract method 0xf50d0835.
//
// Solidity: function chengShiLevelSurety(uint256 ) view returns(uint256)
func (_City *CitySession) ChengShiLevelSurety(arg0 *big.Int) (*big.Int, error) {
	return _City.Contract.ChengShiLevelSurety(&_City.CallOpts, arg0)
}

// ChengShiLevelSurety is a free data retrieval call binding the contract method 0xf50d0835.
//
// Solidity: function chengShiLevelSurety(uint256 ) view returns(uint256)
func (_City *CityCallerSession) ChengShiLevelSurety(arg0 *big.Int) (*big.Int, error) {
	return _City.Contract.ChengShiLevelSurety(&_City.CallOpts, arg0)
}

// ChengShiPioneer is a free data retrieval call binding the contract method 0x157832e5.
//
// Solidity: function chengShiPioneer(bytes32 ) view returns(address)
func (_City *CityCaller) ChengShiPioneer(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _City.contract.Call(opts, &out, "chengShiPioneer", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ChengShiPioneer is a free data retrieval call binding the contract method 0x157832e5.
//
// Solidity: function chengShiPioneer(bytes32 ) view returns(address)
func (_City *CitySession) ChengShiPioneer(arg0 [32]byte) (common.Address, error) {
	return _City.Contract.ChengShiPioneer(&_City.CallOpts, arg0)
}

// ChengShiPioneer is a free data retrieval call binding the contract method 0x157832e5.
//
// Solidity: function chengShiPioneer(bytes32 ) view returns(address)
func (_City *CityCallerSession) ChengShiPioneer(arg0 [32]byte) (common.Address, error) {
	return _City.Contract.ChengShiPioneer(&_City.CallOpts, arg0)
}

// ChengShiRechargeWeight is a free data retrieval call binding the contract method 0xad9eeaee.
//
// Solidity: function chengShiRechargeWeight(bytes32 ) view returns(uint256)
func (_City *CityCaller) ChengShiRechargeWeight(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _City.contract.Call(opts, &out, "chengShiRechargeWeight", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChengShiRechargeWeight is a free data retrieval call binding the contract method 0xad9eeaee.
//
// Solidity: function chengShiRechargeWeight(bytes32 ) view returns(uint256)
func (_City *CitySession) ChengShiRechargeWeight(arg0 [32]byte) (*big.Int, error) {
	return _City.Contract.ChengShiRechargeWeight(&_City.CallOpts, arg0)
}

// ChengShiRechargeWeight is a free data retrieval call binding the contract method 0xad9eeaee.
//
// Solidity: function chengShiRechargeWeight(bytes32 ) view returns(uint256)
func (_City *CityCallerSession) ChengShiRechargeWeight(arg0 [32]byte) (*big.Int, error) {
	return _City.Contract.ChengShiRechargeWeight(&_City.CallOpts, arg0)
}

// CityDelegate is a free data retrieval call binding the contract method 0xb1014abb.
//
// Solidity: function cityDelegate(bytes32 ) view returns(uint256)
func (_City *CityCaller) CityDelegate(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _City.contract.Call(opts, &out, "cityDelegate", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CityDelegate is a free data retrieval call binding the contract method 0xb1014abb.
//
// Solidity: function cityDelegate(bytes32 ) view returns(uint256)
func (_City *CitySession) CityDelegate(arg0 [32]byte) (*big.Int, error) {
	return _City.Contract.CityDelegate(&_City.CallOpts, arg0)
}

// CityDelegate is a free data retrieval call binding the contract method 0xb1014abb.
//
// Solidity: function cityDelegate(bytes32 ) view returns(uint256)
func (_City *CityCallerSession) CityDelegate(arg0 [32]byte) (*big.Int, error) {
	return _City.Contract.CityDelegate(&_City.CallOpts, arg0)
}

// CityDelegateRecord is a free data retrieval call binding the contract method 0x19d944bb.
//
// Solidity: function cityDelegateRecord(bytes32 , uint256 ) view returns(uint256)
func (_City *CityCaller) CityDelegateRecord(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _City.contract.Call(opts, &out, "cityDelegateRecord", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CityDelegateRecord is a free data retrieval call binding the contract method 0x19d944bb.
//
// Solidity: function cityDelegateRecord(bytes32 , uint256 ) view returns(uint256)
func (_City *CitySession) CityDelegateRecord(arg0 [32]byte, arg1 *big.Int) (*big.Int, error) {
	return _City.Contract.CityDelegateRecord(&_City.CallOpts, arg0, arg1)
}

// CityDelegateRecord is a free data retrieval call binding the contract method 0x19d944bb.
//
// Solidity: function cityDelegateRecord(bytes32 , uint256 ) view returns(uint256)
func (_City *CityCallerSession) CityDelegateRecord(arg0 [32]byte, arg1 *big.Int) (*big.Int, error) {
	return _City.Contract.CityDelegateRecord(&_City.CallOpts, arg0, arg1)
}

// CityDelegateTotal is a free data retrieval call binding the contract method 0x557e3f1f.
//
// Solidity: function cityDelegateTotal(bytes32 ) view returns(uint256)
func (_City *CityCaller) CityDelegateTotal(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _City.contract.Call(opts, &out, "cityDelegateTotal", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CityDelegateTotal is a free data retrieval call binding the contract method 0x557e3f1f.
//
// Solidity: function cityDelegateTotal(bytes32 ) view returns(uint256)
func (_City *CitySession) CityDelegateTotal(arg0 [32]byte) (*big.Int, error) {
	return _City.Contract.CityDelegateTotal(&_City.CallOpts, arg0)
}

// CityDelegateTotal is a free data retrieval call binding the contract method 0x557e3f1f.
//
// Solidity: function cityDelegateTotal(bytes32 ) view returns(uint256)
func (_City *CityCallerSession) CityDelegateTotal(arg0 [32]byte) (*big.Int, error) {
	return _City.Contract.CityDelegateTotal(&_City.CallOpts, arg0)
}

// CityFoundsRecord is a free data retrieval call binding the contract method 0x6023cf25.
//
// Solidity: function cityFoundsRecord(bytes32 , uint256 ) view returns(uint256)
func (_City *CityCaller) CityFoundsRecord(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _City.contract.Call(opts, &out, "cityFoundsRecord", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CityFoundsRecord is a free data retrieval call binding the contract method 0x6023cf25.
//
// Solidity: function cityFoundsRecord(bytes32 , uint256 ) view returns(uint256)
func (_City *CitySession) CityFoundsRecord(arg0 [32]byte, arg1 *big.Int) (*big.Int, error) {
	return _City.Contract.CityFoundsRecord(&_City.CallOpts, arg0, arg1)
}

// CityFoundsRecord is a free data retrieval call binding the contract method 0x6023cf25.
//
// Solidity: function cityFoundsRecord(bytes32 , uint256 ) view returns(uint256)
func (_City *CityCallerSession) CityFoundsRecord(arg0 [32]byte, arg1 *big.Int) (*big.Int, error) {
	return _City.Contract.CityFoundsRecord(&_City.CallOpts, arg0, arg1)
}

// CityFoundsTotal is a free data retrieval call binding the contract method 0x6d0ae024.
//
// Solidity: function cityFoundsTotal(bytes32 ) view returns(uint256)
func (_City *CityCaller) CityFoundsTotal(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _City.contract.Call(opts, &out, "cityFoundsTotal", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CityFoundsTotal is a free data retrieval call binding the contract method 0x6d0ae024.
//
// Solidity: function cityFoundsTotal(bytes32 ) view returns(uint256)
func (_City *CitySession) CityFoundsTotal(arg0 [32]byte) (*big.Int, error) {
	return _City.Contract.CityFoundsTotal(&_City.CallOpts, arg0)
}

// CityFoundsTotal is a free data retrieval call binding the contract method 0x6d0ae024.
//
// Solidity: function cityFoundsTotal(bytes32 ) view returns(uint256)
func (_City *CityCallerSession) CityFoundsTotal(arg0 [32]byte) (*big.Int, error) {
	return _City.Contract.CityFoundsTotal(&_City.CallOpts, arg0)
}

// CityLevel is a free data retrieval call binding the contract method 0x0c56059d.
//
// Solidity: function cityLevel(bytes32 ) view returns(uint256)
func (_City *CityCaller) CityLevel(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _City.contract.Call(opts, &out, "cityLevel", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CityLevel is a free data retrieval call binding the contract method 0x0c56059d.
//
// Solidity: function cityLevel(bytes32 ) view returns(uint256)
func (_City *CitySession) CityLevel(arg0 [32]byte) (*big.Int, error) {
	return _City.Contract.CityLevel(&_City.CallOpts, arg0)
}

// CityLevel is a free data retrieval call binding the contract method 0x0c56059d.
//
// Solidity: function cityLevel(bytes32 ) view returns(uint256)
func (_City *CityCallerSession) CityLevel(arg0 [32]byte) (*big.Int, error) {
	return _City.Contract.CityLevel(&_City.CallOpts, arg0)
}

// CityMaxDelegate is a free data retrieval call binding the contract method 0x6deea120.
//
// Solidity: function cityMaxDelegate(bytes32 , uint256 ) view returns(uint256)
func (_City *CityCaller) CityMaxDelegate(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _City.contract.Call(opts, &out, "cityMaxDelegate", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CityMaxDelegate is a free data retrieval call binding the contract method 0x6deea120.
//
// Solidity: function cityMaxDelegate(bytes32 , uint256 ) view returns(uint256)
func (_City *CitySession) CityMaxDelegate(arg0 [32]byte, arg1 *big.Int) (*big.Int, error) {
	return _City.Contract.CityMaxDelegate(&_City.CallOpts, arg0, arg1)
}

// CityMaxDelegate is a free data retrieval call binding the contract method 0x6deea120.
//
// Solidity: function cityMaxDelegate(bytes32 , uint256 ) view returns(uint256)
func (_City *CityCallerSession) CityMaxDelegate(arg0 [32]byte, arg1 *big.Int) (*big.Int, error) {
	return _City.Contract.CityMaxDelegate(&_City.CallOpts, arg0, arg1)
}

// CityNewlyDelegateRecord is a free data retrieval call binding the contract method 0x60442166.
//
// Solidity: function cityNewlyDelegateRecord(bytes32 , uint256 ) view returns(uint256)
func (_City *CityCaller) CityNewlyDelegateRecord(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _City.contract.Call(opts, &out, "cityNewlyDelegateRecord", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CityNewlyDelegateRecord is a free data retrieval call binding the contract method 0x60442166.
//
// Solidity: function cityNewlyDelegateRecord(bytes32 , uint256 ) view returns(uint256)
func (_City *CitySession) CityNewlyDelegateRecord(arg0 [32]byte, arg1 *big.Int) (*big.Int, error) {
	return _City.Contract.CityNewlyDelegateRecord(&_City.CallOpts, arg0, arg1)
}

// CityNewlyDelegateRecord is a free data retrieval call binding the contract method 0x60442166.
//
// Solidity: function cityNewlyDelegateRecord(bytes32 , uint256 ) view returns(uint256)
func (_City *CityCallerSession) CityNewlyDelegateRecord(arg0 [32]byte, arg1 *big.Int) (*big.Int, error) {
	return _City.Contract.CityNewlyDelegateRecord(&_City.CallOpts, arg0, arg1)
}

// CityOrChengShiWeightTotal is a free data retrieval call binding the contract method 0x2139df5b.
//
// Solidity: function cityOrChengShiWeightTotal(bytes32 ) view returns(uint256)
func (_City *CityCaller) CityOrChengShiWeightTotal(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _City.contract.Call(opts, &out, "cityOrChengShiWeightTotal", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CityOrChengShiWeightTotal is a free data retrieval call binding the contract method 0x2139df5b.
//
// Solidity: function cityOrChengShiWeightTotal(bytes32 ) view returns(uint256)
func (_City *CitySession) CityOrChengShiWeightTotal(arg0 [32]byte) (*big.Int, error) {
	return _City.Contract.CityOrChengShiWeightTotal(&_City.CallOpts, arg0)
}

// CityOrChengShiWeightTotal is a free data retrieval call binding the contract method 0x2139df5b.
//
// Solidity: function cityOrChengShiWeightTotal(bytes32 ) view returns(uint256)
func (_City *CityCallerSession) CityOrChengShiWeightTotal(arg0 [32]byte) (*big.Int, error) {
	return _City.Contract.CityOrChengShiWeightTotal(&_City.CallOpts, arg0)
}

// CityPioneer is a free data retrieval call binding the contract method 0xa02018a3.
//
// Solidity: function cityPioneer(bytes32 ) view returns(address)
func (_City *CityCaller) CityPioneer(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _City.contract.Call(opts, &out, "cityPioneer", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CityPioneer is a free data retrieval call binding the contract method 0xa02018a3.
//
// Solidity: function cityPioneer(bytes32 ) view returns(address)
func (_City *CitySession) CityPioneer(arg0 [32]byte) (common.Address, error) {
	return _City.Contract.CityPioneer(&_City.CallOpts, arg0)
}

// CityPioneer is a free data retrieval call binding the contract method 0xa02018a3.
//
// Solidity: function cityPioneer(bytes32 ) view returns(address)
func (_City *CityCallerSession) CityPioneer(arg0 [32]byte) (common.Address, error) {
	return _City.Contract.CityPioneer(&_City.CallOpts, arg0)
}

// CityPioneerAddress is a free data retrieval call binding the contract method 0x2902e1f1.
//
// Solidity: function cityPioneerAddress() view returns(address)
func (_City *CityCaller) CityPioneerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _City.contract.Call(opts, &out, "cityPioneerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CityPioneerAddress is a free data retrieval call binding the contract method 0x2902e1f1.
//
// Solidity: function cityPioneerAddress() view returns(address)
func (_City *CitySession) CityPioneerAddress() (common.Address, error) {
	return _City.Contract.CityPioneerAddress(&_City.CallOpts)
}

// CityPioneerAddress is a free data retrieval call binding the contract method 0x2902e1f1.
//
// Solidity: function cityPioneerAddress() view returns(address)
func (_City *CityCallerSession) CityPioneerAddress() (common.Address, error) {
	return _City.Contract.CityPioneerAddress(&_City.CallOpts)
}

// CityPioneerAssessment is a free data retrieval call binding the contract method 0x4e75f821.
//
// Solidity: function cityPioneerAssessment(bytes32 ) view returns(bool)
func (_City *CityCaller) CityPioneerAssessment(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _City.contract.Call(opts, &out, "cityPioneerAssessment", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CityPioneerAssessment is a free data retrieval call binding the contract method 0x4e75f821.
//
// Solidity: function cityPioneerAssessment(bytes32 ) view returns(bool)
func (_City *CitySession) CityPioneerAssessment(arg0 [32]byte) (bool, error) {
	return _City.Contract.CityPioneerAssessment(&_City.CallOpts, arg0)
}

// CityPioneerAssessment is a free data retrieval call binding the contract method 0x4e75f821.
//
// Solidity: function cityPioneerAssessment(bytes32 ) view returns(bool)
func (_City *CityCallerSession) CityPioneerAssessment(arg0 [32]byte) (bool, error) {
	return _City.Contract.CityPioneerAssessment(&_City.CallOpts, arg0)
}

// CityRechargeTotal is a free data retrieval call binding the contract method 0x834e42ce.
//
// Solidity: function cityRechargeTotal(bytes32 ) view returns(uint256)
func (_City *CityCaller) CityRechargeTotal(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _City.contract.Call(opts, &out, "cityRechargeTotal", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CityRechargeTotal is a free data retrieval call binding the contract method 0x834e42ce.
//
// Solidity: function cityRechargeTotal(bytes32 ) view returns(uint256)
func (_City *CitySession) CityRechargeTotal(arg0 [32]byte) (*big.Int, error) {
	return _City.Contract.CityRechargeTotal(&_City.CallOpts, arg0)
}

// CityRechargeTotal is a free data retrieval call binding the contract method 0x834e42ce.
//
// Solidity: function cityRechargeTotal(bytes32 ) view returns(uint256)
func (_City *CityCallerSession) CityRechargeTotal(arg0 [32]byte) (*big.Int, error) {
	return _City.Contract.CityRechargeTotal(&_City.CallOpts, arg0)
}

// CountyNewlyPioneerDelegateRecord is a free data retrieval call binding the contract method 0x6c6865e4.
//
// Solidity: function countyNewlyPioneerDelegateRecord(bytes32 , uint256 ) view returns(uint256)
func (_City *CityCaller) CountyNewlyPioneerDelegateRecord(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _City.contract.Call(opts, &out, "countyNewlyPioneerDelegateRecord", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CountyNewlyPioneerDelegateRecord is a free data retrieval call binding the contract method 0x6c6865e4.
//
// Solidity: function countyNewlyPioneerDelegateRecord(bytes32 , uint256 ) view returns(uint256)
func (_City *CitySession) CountyNewlyPioneerDelegateRecord(arg0 [32]byte, arg1 *big.Int) (*big.Int, error) {
	return _City.Contract.CountyNewlyPioneerDelegateRecord(&_City.CallOpts, arg0, arg1)
}

// CountyNewlyPioneerDelegateRecord is a free data retrieval call binding the contract method 0x6c6865e4.
//
// Solidity: function countyNewlyPioneerDelegateRecord(bytes32 , uint256 ) view returns(uint256)
func (_City *CityCallerSession) CountyNewlyPioneerDelegateRecord(arg0 [32]byte, arg1 *big.Int) (*big.Int, error) {
	return _City.Contract.CountyNewlyPioneerDelegateRecord(&_City.CallOpts, arg0, arg1)
}

// DailyTaskStatus is a free data retrieval call binding the contract method 0x0c069666.
//
// Solidity: function dailyTaskStatus(uint256 ) view returns(bool)
func (_City *CityCaller) DailyTaskStatus(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _City.contract.Call(opts, &out, "dailyTaskStatus", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DailyTaskStatus is a free data retrieval call binding the contract method 0x0c069666.
//
// Solidity: function dailyTaskStatus(uint256 ) view returns(bool)
func (_City *CitySession) DailyTaskStatus(arg0 *big.Int) (bool, error) {
	return _City.Contract.DailyTaskStatus(&_City.CallOpts, arg0)
}

// DailyTaskStatus is a free data retrieval call binding the contract method 0x0c069666.
//
// Solidity: function dailyTaskStatus(uint256 ) view returns(bool)
func (_City *CityCallerSession) DailyTaskStatus(arg0 *big.Int) (bool, error) {
	return _City.Contract.DailyTaskStatus(&_City.CallOpts, arg0)
}

// EarnestMoney is a free data retrieval call binding the contract method 0x533264ef.
//
// Solidity: function earnestMoney(bytes32 ) view returns(uint256)
func (_City *CityCaller) EarnestMoney(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _City.contract.Call(opts, &out, "earnestMoney", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EarnestMoney is a free data retrieval call binding the contract method 0x533264ef.
//
// Solidity: function earnestMoney(bytes32 ) view returns(uint256)
func (_City *CitySession) EarnestMoney(arg0 [32]byte) (*big.Int, error) {
	return _City.Contract.EarnestMoney(&_City.CallOpts, arg0)
}

// EarnestMoney is a free data retrieval call binding the contract method 0x533264ef.
//
// Solidity: function earnestMoney(bytes32 ) view returns(uint256)
func (_City *CityCallerSession) EarnestMoney(arg0 [32]byte) (*big.Int, error) {
	return _City.Contract.EarnestMoney(&_City.CallOpts, arg0)
}

// FoundsAddress is a free data retrieval call binding the contract method 0x816467f1.
//
// Solidity: function foundsAddress() view returns(address)
func (_City *CityCaller) FoundsAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _City.contract.Call(opts, &out, "foundsAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FoundsAddress is a free data retrieval call binding the contract method 0x816467f1.
//
// Solidity: function foundsAddress() view returns(address)
func (_City *CitySession) FoundsAddress() (common.Address, error) {
	return _City.Contract.FoundsAddress(&_City.CallOpts)
}

// FoundsAddress is a free data retrieval call binding the contract method 0x816467f1.
//
// Solidity: function foundsAddress() view returns(address)
func (_City *CityCallerSession) FoundsAddress() (common.Address, error) {
	return _City.Contract.FoundsAddress(&_City.CallOpts)
}

// GetChengShiRechargeWeight is a free data retrieval call binding the contract method 0xcfb9a307.
//
// Solidity: function getChengShiRechargeWeight(bytes32 chengShiId) view returns(uint256)
func (_City *CityCaller) GetChengShiRechargeWeight(opts *bind.CallOpts, chengShiId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _City.contract.Call(opts, &out, "getChengShiRechargeWeight", chengShiId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetChengShiRechargeWeight is a free data retrieval call binding the contract method 0xcfb9a307.
//
// Solidity: function getChengShiRechargeWeight(bytes32 chengShiId) view returns(uint256)
func (_City *CitySession) GetChengShiRechargeWeight(chengShiId [32]byte) (*big.Int, error) {
	return _City.Contract.GetChengShiRechargeWeight(&_City.CallOpts, chengShiId)
}

// GetChengShiRechargeWeight is a free data retrieval call binding the contract method 0xcfb9a307.
//
// Solidity: function getChengShiRechargeWeight(bytes32 chengShiId) view returns(uint256)
func (_City *CityCallerSession) GetChengShiRechargeWeight(chengShiId [32]byte) (*big.Int, error) {
	return _City.Contract.GetChengShiRechargeWeight(&_City.CallOpts, chengShiId)
}

// GetDailyDelegateByChengShiID is a free data retrieval call binding the contract method 0x2180505b.
//
// Solidity: function getDailyDelegateByChengShiID(bytes32 chengShiId_, uint256 day, bool isCityPioneer) view returns(uint256)
func (_City *CityCaller) GetDailyDelegateByChengShiID(opts *bind.CallOpts, chengShiId_ [32]byte, day *big.Int, isCityPioneer bool) (*big.Int, error) {
	var out []interface{}
	err := _City.contract.Call(opts, &out, "getDailyDelegateByChengShiID", chengShiId_, day, isCityPioneer)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDailyDelegateByChengShiID is a free data retrieval call binding the contract method 0x2180505b.
//
// Solidity: function getDailyDelegateByChengShiID(bytes32 chengShiId_, uint256 day, bool isCityPioneer) view returns(uint256)
func (_City *CitySession) GetDailyDelegateByChengShiID(chengShiId_ [32]byte, day *big.Int, isCityPioneer bool) (*big.Int, error) {
	return _City.Contract.GetDailyDelegateByChengShiID(&_City.CallOpts, chengShiId_, day, isCityPioneer)
}

// GetDailyDelegateByChengShiID is a free data retrieval call binding the contract method 0x2180505b.
//
// Solidity: function getDailyDelegateByChengShiID(bytes32 chengShiId_, uint256 day, bool isCityPioneer) view returns(uint256)
func (_City *CityCallerSession) GetDailyDelegateByChengShiID(chengShiId_ [32]byte, day *big.Int, isCityPioneer bool) (*big.Int, error) {
	return _City.Contract.GetDailyDelegateByChengShiID(&_City.CallOpts, chengShiId_, day, isCityPioneer)
}

// GetDay is a free data retrieval call binding the contract method 0x14ba5c09.
//
// Solidity: function getDay() view returns(uint256)
func (_City *CityCaller) GetDay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _City.contract.Call(opts, &out, "getDay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDay is a free data retrieval call binding the contract method 0x14ba5c09.
//
// Solidity: function getDay() view returns(uint256)
func (_City *CitySession) GetDay() (*big.Int, error) {
	return _City.Contract.GetDay(&_City.CallOpts)
}

// GetDay is a free data retrieval call binding the contract method 0x14ba5c09.
//
// Solidity: function getDay() view returns(uint256)
func (_City *CityCallerSession) GetDay() (*big.Int, error) {
	return _City.Contract.GetDay(&_City.CallOpts)
}

// GetDelegate is a free data retrieval call binding the contract method 0xecaa7e3c.
//
// Solidity: function getDelegate(bytes32 cityId_, uint256 day) view returns(uint256)
func (_City *CityCaller) GetDelegate(opts *bind.CallOpts, cityId_ [32]byte, day *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _City.contract.Call(opts, &out, "getDelegate", cityId_, day)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDelegate is a free data retrieval call binding the contract method 0xecaa7e3c.
//
// Solidity: function getDelegate(bytes32 cityId_, uint256 day) view returns(uint256)
func (_City *CitySession) GetDelegate(cityId_ [32]byte, day *big.Int) (*big.Int, error) {
	return _City.Contract.GetDelegate(&_City.CallOpts, cityId_, day)
}

// GetDelegate is a free data retrieval call binding the contract method 0xecaa7e3c.
//
// Solidity: function getDelegate(bytes32 cityId_, uint256 day) view returns(uint256)
func (_City *CityCallerSession) GetDelegate(cityId_ [32]byte, day *big.Int) (*big.Int, error) {
	return _City.Contract.GetDelegate(&_City.CallOpts, cityId_, day)
}

// GetFifteenDayAverageFounds is a free data retrieval call binding the contract method 0x316a3a29.
//
// Solidity: function getFifteenDayAverageFounds() view returns(uint256)
func (_City *CityCaller) GetFifteenDayAverageFounds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _City.contract.Call(opts, &out, "getFifteenDayAverageFounds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFifteenDayAverageFounds is a free data retrieval call binding the contract method 0x316a3a29.
//
// Solidity: function getFifteenDayAverageFounds() view returns(uint256)
func (_City *CitySession) GetFifteenDayAverageFounds() (*big.Int, error) {
	return _City.Contract.GetFifteenDayAverageFounds(&_City.CallOpts)
}

// GetFifteenDayAverageFounds is a free data retrieval call binding the contract method 0x316a3a29.
//
// Solidity: function getFifteenDayAverageFounds() view returns(uint256)
func (_City *CityCallerSession) GetFifteenDayAverageFounds() (*big.Int, error) {
	return _City.Contract.GetFifteenDayAverageFounds(&_City.CallOpts)
}

// GetFounds is a free data retrieval call binding the contract method 0x67032599.
//
// Solidity: function getFounds(bytes32 cityId_, uint256 day) view returns(uint256)
func (_City *CityCaller) GetFounds(opts *bind.CallOpts, cityId_ [32]byte, day *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _City.contract.Call(opts, &out, "getFounds", cityId_, day)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFounds is a free data retrieval call binding the contract method 0x67032599.
//
// Solidity: function getFounds(bytes32 cityId_, uint256 day) view returns(uint256)
func (_City *CitySession) GetFounds(cityId_ [32]byte, day *big.Int) (*big.Int, error) {
	return _City.Contract.GetFounds(&_City.CallOpts, cityId_, day)
}

// GetFounds is a free data retrieval call binding the contract method 0x67032599.
//
// Solidity: function getFounds(bytes32 cityId_, uint256 day) view returns(uint256)
func (_City *CityCallerSession) GetFounds(cityId_ [32]byte, day *big.Int) (*big.Int, error) {
	return _City.Contract.GetFounds(&_City.CallOpts, cityId_, day)
}

// GetPioneerCityNumber is a free data retrieval call binding the contract method 0xe0db7746.
//
// Solidity: function getPioneerCityNumber() view returns(uint256)
func (_City *CityCaller) GetPioneerCityNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _City.contract.Call(opts, &out, "getPioneerCityNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPioneerCityNumber is a free data retrieval call binding the contract method 0xe0db7746.
//
// Solidity: function getPioneerCityNumber() view returns(uint256)
func (_City *CitySession) GetPioneerCityNumber() (*big.Int, error) {
	return _City.Contract.GetPioneerCityNumber(&_City.CallOpts)
}

// GetPioneerCityNumber is a free data retrieval call binding the contract method 0xe0db7746.
//
// Solidity: function getPioneerCityNumber() view returns(uint256)
func (_City *CityCallerSession) GetPioneerCityNumber() (*big.Int, error) {
	return _City.Contract.GetPioneerCityNumber(&_City.CallOpts)
}

// GetSuretyByPioneerAddress is a free data retrieval call binding the contract method 0xd9bc53b4.
//
// Solidity: function getSuretyByPioneerAddress(address pioneer_) view returns()
func (_City *CityCaller) GetSuretyByPioneerAddress(opts *bind.CallOpts, pioneer_ common.Address) error {
	var out []interface{}
	err := _City.contract.Call(opts, &out, "getSuretyByPioneerAddress", pioneer_)

	if err != nil {
		return err
	}

	return err

}

// GetSuretyByPioneerAddress is a free data retrieval call binding the contract method 0xd9bc53b4.
//
// Solidity: function getSuretyByPioneerAddress(address pioneer_) view returns()
func (_City *CitySession) GetSuretyByPioneerAddress(pioneer_ common.Address) error {
	return _City.Contract.GetSuretyByPioneerAddress(&_City.CallOpts, pioneer_)
}

// GetSuretyByPioneerAddress is a free data retrieval call binding the contract method 0xd9bc53b4.
//
// Solidity: function getSuretyByPioneerAddress(address pioneer_) view returns()
func (_City *CityCallerSession) GetSuretyByPioneerAddress(pioneer_ common.Address) error {
	return _City.Contract.GetSuretyByPioneerAddress(&_City.CallOpts, pioneer_)
}

// HasSetPioneer is a free data retrieval call binding the contract method 0x703d924c.
//
// Solidity: function hasSetPioneer(address ) view returns(bool)
func (_City *CityCaller) HasSetPioneer(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _City.contract.Call(opts, &out, "hasSetPioneer", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasSetPioneer is a free data retrieval call binding the contract method 0x703d924c.
//
// Solidity: function hasSetPioneer(address ) view returns(bool)
func (_City *CitySession) HasSetPioneer(arg0 common.Address) (bool, error) {
	return _City.Contract.HasSetPioneer(&_City.CallOpts, arg0)
}

// HasSetPioneer is a free data retrieval call binding the contract method 0x703d924c.
//
// Solidity: function hasSetPioneer(address ) view returns(bool)
func (_City *CityCallerSession) HasSetPioneer(arg0 common.Address) (bool, error) {
	return _City.Contract.HasSetPioneer(&_City.CallOpts, arg0)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address account) view returns(bool)
func (_City *CityCaller) IsAdmin(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _City.contract.Call(opts, &out, "isAdmin", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address account) view returns(bool)
func (_City *CitySession) IsAdmin(account common.Address) (bool, error) {
	return _City.Contract.IsAdmin(&_City.CallOpts, account)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address account) view returns(bool)
func (_City *CityCallerSession) IsAdmin(account common.Address) (bool, error) {
	return _City.Contract.IsAdmin(&_City.CallOpts, account)
}

// PioneerChengShi is a free data retrieval call binding the contract method 0x84ad5018.
//
// Solidity: function pioneerChengShi(address ) view returns(bytes32)
func (_City *CityCaller) PioneerChengShi(opts *bind.CallOpts, arg0 common.Address) ([32]byte, error) {
	var out []interface{}
	err := _City.contract.Call(opts, &out, "pioneerChengShi", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PioneerChengShi is a free data retrieval call binding the contract method 0x84ad5018.
//
// Solidity: function pioneerChengShi(address ) view returns(bytes32)
func (_City *CitySession) PioneerChengShi(arg0 common.Address) ([32]byte, error) {
	return _City.Contract.PioneerChengShi(&_City.CallOpts, arg0)
}

// PioneerChengShi is a free data retrieval call binding the contract method 0x84ad5018.
//
// Solidity: function pioneerChengShi(address ) view returns(bytes32)
func (_City *CityCallerSession) PioneerChengShi(arg0 common.Address) ([32]byte, error) {
	return _City.Contract.PioneerChengShi(&_City.CallOpts, arg0)
}

// PioneerChengShiIdExits is a free data retrieval call binding the contract method 0x1e853976.
//
// Solidity: function pioneerChengShiIdExits(bytes32 chengShiId) view returns(bool)
func (_City *CityCaller) PioneerChengShiIdExits(opts *bind.CallOpts, chengShiId [32]byte) (bool, error) {
	var out []interface{}
	err := _City.contract.Call(opts, &out, "pioneerChengShiIdExits", chengShiId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PioneerChengShiIdExits is a free data retrieval call binding the contract method 0x1e853976.
//
// Solidity: function pioneerChengShiIdExits(bytes32 chengShiId) view returns(bool)
func (_City *CitySession) PioneerChengShiIdExits(chengShiId [32]byte) (bool, error) {
	return _City.Contract.PioneerChengShiIdExits(&_City.CallOpts, chengShiId)
}

// PioneerChengShiIdExits is a free data retrieval call binding the contract method 0x1e853976.
//
// Solidity: function pioneerChengShiIdExits(bytes32 chengShiId) view returns(bool)
func (_City *CityCallerSession) PioneerChengShiIdExits(chengShiId [32]byte) (bool, error) {
	return _City.Contract.PioneerChengShiIdExits(&_City.CallOpts, chengShiId)
}

// PioneerChengShiIds is a free data retrieval call binding the contract method 0xf7aa0de1.
//
// Solidity: function pioneerChengShiIds(uint256 ) view returns(bytes32)
func (_City *CityCaller) PioneerChengShiIds(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _City.contract.Call(opts, &out, "pioneerChengShiIds", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PioneerChengShiIds is a free data retrieval call binding the contract method 0xf7aa0de1.
//
// Solidity: function pioneerChengShiIds(uint256 ) view returns(bytes32)
func (_City *CitySession) PioneerChengShiIds(arg0 *big.Int) ([32]byte, error) {
	return _City.Contract.PioneerChengShiIds(&_City.CallOpts, arg0)
}

// PioneerChengShiIds is a free data retrieval call binding the contract method 0xf7aa0de1.
//
// Solidity: function pioneerChengShiIds(uint256 ) view returns(bytes32)
func (_City *CityCallerSession) PioneerChengShiIds(arg0 *big.Int) ([32]byte, error) {
	return _City.Contract.PioneerChengShiIds(&_City.CallOpts, arg0)
}

// PioneerCity is a free data retrieval call binding the contract method 0x2d7fa49a.
//
// Solidity: function pioneerCity(address ) view returns(bytes32)
func (_City *CityCaller) PioneerCity(opts *bind.CallOpts, arg0 common.Address) ([32]byte, error) {
	var out []interface{}
	err := _City.contract.Call(opts, &out, "pioneerCity", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PioneerCity is a free data retrieval call binding the contract method 0x2d7fa49a.
//
// Solidity: function pioneerCity(address ) view returns(bytes32)
func (_City *CitySession) PioneerCity(arg0 common.Address) ([32]byte, error) {
	return _City.Contract.PioneerCity(&_City.CallOpts, arg0)
}

// PioneerCity is a free data retrieval call binding the contract method 0x2d7fa49a.
//
// Solidity: function pioneerCity(address ) view returns(bytes32)
func (_City *CityCallerSession) PioneerCity(arg0 common.Address) ([32]byte, error) {
	return _City.Contract.PioneerCity(&_City.CallOpts, arg0)
}

// PioneerCityIds is a free data retrieval call binding the contract method 0x592d4c0f.
//
// Solidity: function pioneerCityIds(uint256 ) view returns(bytes32)
func (_City *CityCaller) PioneerCityIds(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _City.contract.Call(opts, &out, "pioneerCityIds", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PioneerCityIds is a free data retrieval call binding the contract method 0x592d4c0f.
//
// Solidity: function pioneerCityIds(uint256 ) view returns(bytes32)
func (_City *CitySession) PioneerCityIds(arg0 *big.Int) ([32]byte, error) {
	return _City.Contract.PioneerCityIds(&_City.CallOpts, arg0)
}

// PioneerCityIds is a free data retrieval call binding the contract method 0x592d4c0f.
//
// Solidity: function pioneerCityIds(uint256 ) view returns(bytes32)
func (_City *CityCallerSession) PioneerCityIds(arg0 *big.Int) ([32]byte, error) {
	return _City.Contract.PioneerCityIds(&_City.CallOpts, arg0)
}

// PioneerStatus is a free data retrieval call binding the contract method 0xb5b1a11e.
//
// Solidity: function pioneerStatus(address ) view returns(bool)
func (_City *CityCaller) PioneerStatus(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _City.contract.Call(opts, &out, "pioneerStatus", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PioneerStatus is a free data retrieval call binding the contract method 0xb5b1a11e.
//
// Solidity: function pioneerStatus(address ) view returns(bool)
func (_City *CitySession) PioneerStatus(arg0 common.Address) (bool, error) {
	return _City.Contract.PioneerStatus(&_City.CallOpts, arg0)
}

// PioneerStatus is a free data retrieval call binding the contract method 0xb5b1a11e.
//
// Solidity: function pioneerStatus(address ) view returns(bool)
func (_City *CityCallerSession) PioneerStatus(arg0 common.Address) (bool, error) {
	return _City.Contract.PioneerStatus(&_City.CallOpts, arg0)
}

// RechargeDailyTotalWeightRecord is a free data retrieval call binding the contract method 0x9c84e444.
//
// Solidity: function rechargeDailyTotalWeightRecord(bytes32 , uint256 ) view returns(uint256)
func (_City *CityCaller) RechargeDailyTotalWeightRecord(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _City.contract.Call(opts, &out, "rechargeDailyTotalWeightRecord", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RechargeDailyTotalWeightRecord is a free data retrieval call binding the contract method 0x9c84e444.
//
// Solidity: function rechargeDailyTotalWeightRecord(bytes32 , uint256 ) view returns(uint256)
func (_City *CitySession) RechargeDailyTotalWeightRecord(arg0 [32]byte, arg1 *big.Int) (*big.Int, error) {
	return _City.Contract.RechargeDailyTotalWeightRecord(&_City.CallOpts, arg0, arg1)
}

// RechargeDailyTotalWeightRecord is a free data retrieval call binding the contract method 0x9c84e444.
//
// Solidity: function rechargeDailyTotalWeightRecord(bytes32 , uint256 ) view returns(uint256)
func (_City *CityCallerSession) RechargeDailyTotalWeightRecord(arg0 [32]byte, arg1 *big.Int) (*big.Int, error) {
	return _City.Contract.RechargeDailyTotalWeightRecord(&_City.CallOpts, arg0, arg1)
}

// RechargeDailyWeight is a free data retrieval call binding the contract method 0xe53efdaa.
//
// Solidity: function rechargeDailyWeight(uint256 ) view returns(uint256)
func (_City *CityCaller) RechargeDailyWeight(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _City.contract.Call(opts, &out, "rechargeDailyWeight", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RechargeDailyWeight is a free data retrieval call binding the contract method 0xe53efdaa.
//
// Solidity: function rechargeDailyWeight(uint256 ) view returns(uint256)
func (_City *CitySession) RechargeDailyWeight(arg0 *big.Int) (*big.Int, error) {
	return _City.Contract.RechargeDailyWeight(&_City.CallOpts, arg0)
}

// RechargeDailyWeight is a free data retrieval call binding the contract method 0xe53efdaa.
//
// Solidity: function rechargeDailyWeight(uint256 ) view returns(uint256)
func (_City *CityCallerSession) RechargeDailyWeight(arg0 *big.Int) (*big.Int, error) {
	return _City.Contract.RechargeDailyWeight(&_City.CallOpts, arg0)
}

// RechargeDailyWeightRecord is a free data retrieval call binding the contract method 0x1be94253.
//
// Solidity: function rechargeDailyWeightRecord(bytes32 , uint256 ) view returns(uint256)
func (_City *CityCaller) RechargeDailyWeightRecord(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _City.contract.Call(opts, &out, "rechargeDailyWeightRecord", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RechargeDailyWeightRecord is a free data retrieval call binding the contract method 0x1be94253.
//
// Solidity: function rechargeDailyWeightRecord(bytes32 , uint256 ) view returns(uint256)
func (_City *CitySession) RechargeDailyWeightRecord(arg0 [32]byte, arg1 *big.Int) (*big.Int, error) {
	return _City.Contract.RechargeDailyWeightRecord(&_City.CallOpts, arg0, arg1)
}

// RechargeDailyWeightRecord is a free data retrieval call binding the contract method 0x1be94253.
//
// Solidity: function rechargeDailyWeightRecord(bytes32 , uint256 ) view returns(uint256)
func (_City *CityCallerSession) RechargeDailyWeightRecord(arg0 [32]byte, arg1 *big.Int) (*big.Int, error) {
	return _City.Contract.RechargeDailyWeightRecord(&_City.CallOpts, arg0, arg1)
}

// RechargeWeight is a free data retrieval call binding the contract method 0xb61a2470.
//
// Solidity: function rechargeWeight() view returns(uint256)
func (_City *CityCaller) RechargeWeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _City.contract.Call(opts, &out, "rechargeWeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RechargeWeight is a free data retrieval call binding the contract method 0xb61a2470.
//
// Solidity: function rechargeWeight() view returns(uint256)
func (_City *CitySession) RechargeWeight() (*big.Int, error) {
	return _City.Contract.RechargeWeight(&_City.CallOpts)
}

// RechargeWeight is a free data retrieval call binding the contract method 0xb61a2470.
//
// Solidity: function rechargeWeight() view returns(uint256)
func (_City *CityCallerSession) RechargeWeight() (*big.Int, error) {
	return _City.Contract.RechargeWeight(&_City.CallOpts)
}

// RechargeWeightAdditional is a free data retrieval call binding the contract method 0xbe9475f2.
//
// Solidity: function rechargeWeightAdditional(address ) view returns(uint256)
func (_City *CityCaller) RechargeWeightAdditional(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _City.contract.Call(opts, &out, "rechargeWeightAdditional", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RechargeWeightAdditional is a free data retrieval call binding the contract method 0xbe9475f2.
//
// Solidity: function rechargeWeightAdditional(address ) view returns(uint256)
func (_City *CitySession) RechargeWeightAdditional(arg0 common.Address) (*big.Int, error) {
	return _City.Contract.RechargeWeightAdditional(&_City.CallOpts, arg0)
}

// RechargeWeightAdditional is a free data retrieval call binding the contract method 0xbe9475f2.
//
// Solidity: function rechargeWeightAdditional(address ) view returns(uint256)
func (_City *CityCallerSession) RechargeWeightAdditional(arg0 common.Address) (*big.Int, error) {
	return _City.Contract.RechargeWeightAdditional(&_City.CallOpts, arg0)
}

// SecondsPerDay is a free data retrieval call binding the contract method 0x63809953.
//
// Solidity: function secondsPerDay() view returns(uint256)
func (_City *CityCaller) SecondsPerDay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _City.contract.Call(opts, &out, "secondsPerDay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SecondsPerDay is a free data retrieval call binding the contract method 0x63809953.
//
// Solidity: function secondsPerDay() view returns(uint256)
func (_City *CitySession) SecondsPerDay() (*big.Int, error) {
	return _City.Contract.SecondsPerDay(&_City.CallOpts)
}

// SecondsPerDay is a free data retrieval call binding the contract method 0x63809953.
//
// Solidity: function secondsPerDay() view returns(uint256)
func (_City *CityCallerSession) SecondsPerDay() (*big.Int, error) {
	return _City.Contract.SecondsPerDay(&_City.CallOpts)
}

// Surety is a free data retrieval call binding the contract method 0x9d9ad2ca.
//
// Solidity: function surety(bytes32 ) view returns(uint256)
func (_City *CityCaller) Surety(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _City.contract.Call(opts, &out, "surety", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Surety is a free data retrieval call binding the contract method 0x9d9ad2ca.
//
// Solidity: function surety(bytes32 ) view returns(uint256)
func (_City *CitySession) Surety(arg0 [32]byte) (*big.Int, error) {
	return _City.Contract.Surety(&_City.CallOpts, arg0)
}

// Surety is a free data retrieval call binding the contract method 0x9d9ad2ca.
//
// Solidity: function surety(bytes32 ) view returns(uint256)
func (_City *CityCallerSession) Surety(arg0 [32]byte) (*big.Int, error) {
	return _City.Contract.Surety(&_City.CallOpts, arg0)
}

// UserLocationAddress is a free data retrieval call binding the contract method 0x12516fc6.
//
// Solidity: function userLocationAddress() view returns(address)
func (_City *CityCaller) UserLocationAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _City.contract.Call(opts, &out, "userLocationAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UserLocationAddress is a free data retrieval call binding the contract method 0x12516fc6.
//
// Solidity: function userLocationAddress() view returns(address)
func (_City *CitySession) UserLocationAddress() (common.Address, error) {
	return _City.Contract.UserLocationAddress(&_City.CallOpts)
}

// UserLocationAddress is a free data retrieval call binding the contract method 0x12516fc6.
//
// Solidity: function userLocationAddress() view returns(address)
func (_City *CityCallerSession) UserLocationAddress() (common.Address, error) {
	return _City.Contract.UserLocationAddress(&_City.CallOpts)
}

// WithdrawLimit is a free data retrieval call binding the contract method 0xf848d541.
//
// Solidity: function withdrawLimit() view returns(address)
func (_City *CityCaller) WithdrawLimit(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _City.contract.Call(opts, &out, "withdrawLimit")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WithdrawLimit is a free data retrieval call binding the contract method 0xf848d541.
//
// Solidity: function withdrawLimit() view returns(address)
func (_City *CitySession) WithdrawLimit() (common.Address, error) {
	return _City.Contract.WithdrawLimit(&_City.CallOpts)
}

// WithdrawLimit is a free data retrieval call binding the contract method 0xf848d541.
//
// Solidity: function withdrawLimit() view returns(address)
func (_City *CityCallerSession) WithdrawLimit() (common.Address, error) {
	return _City.Contract.WithdrawLimit(&_City.CallOpts)
}

// WithdrawLimitAddress is a free data retrieval call binding the contract method 0xa0c9ae2f.
//
// Solidity: function withdrawLimitAddress() view returns(address)
func (_City *CityCaller) WithdrawLimitAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _City.contract.Call(opts, &out, "withdrawLimitAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WithdrawLimitAddress is a free data retrieval call binding the contract method 0xa0c9ae2f.
//
// Solidity: function withdrawLimitAddress() view returns(address)
func (_City *CitySession) WithdrawLimitAddress() (common.Address, error) {
	return _City.Contract.WithdrawLimitAddress(&_City.CallOpts)
}

// WithdrawLimitAddress is a free data retrieval call binding the contract method 0xa0c9ae2f.
//
// Solidity: function withdrawLimitAddress() view returns(address)
func (_City *CityCallerSession) WithdrawLimitAddress() (common.Address, error) {
	return _City.Contract.WithdrawLimitAddress(&_City.CallOpts)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address account) returns()
func (_City *CityTransactor) AddAdmin(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _City.contract.Transact(opts, "addAdmin", account)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address account) returns()
func (_City *CitySession) AddAdmin(account common.Address) (*types.Transaction, error) {
	return _City.Contract.AddAdmin(&_City.TransactOpts, account)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address account) returns()
func (_City *CityTransactorSession) AddAdmin(account common.Address) (*types.Transaction, error) {
	return _City.Contract.AddAdmin(&_City.TransactOpts, account)
}

// AdminChangePioneerCityId is a paid mutator transaction binding the contract method 0x9961315e.
//
// Solidity: function adminChangePioneerCityId(bytes32 chengShiId_, address pioneer_) returns()
func (_City *CityTransactor) AdminChangePioneerCityId(opts *bind.TransactOpts, chengShiId_ [32]byte, pioneer_ common.Address) (*types.Transaction, error) {
	return _City.contract.Transact(opts, "adminChangePioneerCityId", chengShiId_, pioneer_)
}

// AdminChangePioneerCityId is a paid mutator transaction binding the contract method 0x9961315e.
//
// Solidity: function adminChangePioneerCityId(bytes32 chengShiId_, address pioneer_) returns()
func (_City *CitySession) AdminChangePioneerCityId(chengShiId_ [32]byte, pioneer_ common.Address) (*types.Transaction, error) {
	return _City.Contract.AdminChangePioneerCityId(&_City.TransactOpts, chengShiId_, pioneer_)
}

// AdminChangePioneerCityId is a paid mutator transaction binding the contract method 0x9961315e.
//
// Solidity: function adminChangePioneerCityId(bytes32 chengShiId_, address pioneer_) returns()
func (_City *CityTransactorSession) AdminChangePioneerCityId(chengShiId_ [32]byte, pioneer_ common.Address) (*types.Transaction, error) {
	return _City.Contract.AdminChangePioneerCityId(&_City.TransactOpts, chengShiId_, pioneer_)
}

// AdminEditSurety is a paid mutator transaction binding the contract method 0xdf000fec.
//
// Solidity: function adminEditSurety(bytes32 chengShiId_, uint256 level_, uint256 surety_) returns()
func (_City *CityTransactor) AdminEditSurety(opts *bind.TransactOpts, chengShiId_ [32]byte, level_ *big.Int, surety_ *big.Int) (*types.Transaction, error) {
	return _City.contract.Transact(opts, "adminEditSurety", chengShiId_, level_, surety_)
}

// AdminEditSurety is a paid mutator transaction binding the contract method 0xdf000fec.
//
// Solidity: function adminEditSurety(bytes32 chengShiId_, uint256 level_, uint256 surety_) returns()
func (_City *CitySession) AdminEditSurety(chengShiId_ [32]byte, level_ *big.Int, surety_ *big.Int) (*types.Transaction, error) {
	return _City.Contract.AdminEditSurety(&_City.TransactOpts, chengShiId_, level_, surety_)
}

// AdminEditSurety is a paid mutator transaction binding the contract method 0xdf000fec.
//
// Solidity: function adminEditSurety(bytes32 chengShiId_, uint256 level_, uint256 surety_) returns()
func (_City *CityTransactorSession) AdminEditSurety(chengShiId_ [32]byte, level_ *big.Int, surety_ *big.Int) (*types.Transaction, error) {
	return _City.Contract.AdminEditSurety(&_City.TransactOpts, chengShiId_, level_, surety_)
}

// AdminPopPioneerChengShiId is a paid mutator transaction binding the contract method 0xc9c7ba6f.
//
// Solidity: function adminPopPioneerChengShiId() returns()
func (_City *CityTransactor) AdminPopPioneerChengShiId(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _City.contract.Transact(opts, "adminPopPioneerChengShiId")
}

// AdminPopPioneerChengShiId is a paid mutator transaction binding the contract method 0xc9c7ba6f.
//
// Solidity: function adminPopPioneerChengShiId() returns()
func (_City *CitySession) AdminPopPioneerChengShiId() (*types.Transaction, error) {
	return _City.Contract.AdminPopPioneerChengShiId(&_City.TransactOpts)
}

// AdminPopPioneerChengShiId is a paid mutator transaction binding the contract method 0xc9c7ba6f.
//
// Solidity: function adminPopPioneerChengShiId() returns()
func (_City *CityTransactorSession) AdminPopPioneerChengShiId() (*types.Transaction, error) {
	return _City.Contract.AdminPopPioneerChengShiId(&_City.TransactOpts)
}

// AdminRemovePioneer is a paid mutator transaction binding the contract method 0xb11ad118.
//
// Solidity: function adminRemovePioneer(bytes32 chengShiId_, address pioneer_) returns()
func (_City *CityTransactor) AdminRemovePioneer(opts *bind.TransactOpts, chengShiId_ [32]byte, pioneer_ common.Address) (*types.Transaction, error) {
	return _City.contract.Transact(opts, "adminRemovePioneer", chengShiId_, pioneer_)
}

// AdminRemovePioneer is a paid mutator transaction binding the contract method 0xb11ad118.
//
// Solidity: function adminRemovePioneer(bytes32 chengShiId_, address pioneer_) returns()
func (_City *CitySession) AdminRemovePioneer(chengShiId_ [32]byte, pioneer_ common.Address) (*types.Transaction, error) {
	return _City.Contract.AdminRemovePioneer(&_City.TransactOpts, chengShiId_, pioneer_)
}

// AdminRemovePioneer is a paid mutator transaction binding the contract method 0xb11ad118.
//
// Solidity: function adminRemovePioneer(bytes32 chengShiId_, address pioneer_) returns()
func (_City *CityTransactorSession) AdminRemovePioneer(chengShiId_ [32]byte, pioneer_ common.Address) (*types.Transaction, error) {
	return _City.Contract.AdminRemovePioneer(&_City.TransactOpts, chengShiId_, pioneer_)
}

// AdminSetAuthAddress is a paid mutator transaction binding the contract method 0x236a8b2d.
//
// Solidity: function adminSetAuthAddress(address authAddress_) returns()
func (_City *CityTransactor) AdminSetAuthAddress(opts *bind.TransactOpts, authAddress_ common.Address) (*types.Transaction, error) {
	return _City.contract.Transact(opts, "adminSetAuthAddress", authAddress_)
}

// AdminSetAuthAddress is a paid mutator transaction binding the contract method 0x236a8b2d.
//
// Solidity: function adminSetAuthAddress(address authAddress_) returns()
func (_City *CitySession) AdminSetAuthAddress(authAddress_ common.Address) (*types.Transaction, error) {
	return _City.Contract.AdminSetAuthAddress(&_City.TransactOpts, authAddress_)
}

// AdminSetAuthAddress is a paid mutator transaction binding the contract method 0x236a8b2d.
//
// Solidity: function adminSetAuthAddress(address authAddress_) returns()
func (_City *CityTransactorSession) AdminSetAuthAddress(authAddress_ common.Address) (*types.Transaction, error) {
	return _City.Contract.AdminSetAuthAddress(&_City.TransactOpts, authAddress_)
}

// AdminSetChengShiLevelAndSurety is a paid mutator transaction binding the contract method 0xb9c74af7.
//
// Solidity: function adminSetChengShiLevelAndSurety(bytes32 chengShiId_, uint256 level_, uint256 surety_) returns()
func (_City *CityTransactor) AdminSetChengShiLevelAndSurety(opts *bind.TransactOpts, chengShiId_ [32]byte, level_ *big.Int, surety_ *big.Int) (*types.Transaction, error) {
	return _City.contract.Transact(opts, "adminSetChengShiLevelAndSurety", chengShiId_, level_, surety_)
}

// AdminSetChengShiLevelAndSurety is a paid mutator transaction binding the contract method 0xb9c74af7.
//
// Solidity: function adminSetChengShiLevelAndSurety(bytes32 chengShiId_, uint256 level_, uint256 surety_) returns()
func (_City *CitySession) AdminSetChengShiLevelAndSurety(chengShiId_ [32]byte, level_ *big.Int, surety_ *big.Int) (*types.Transaction, error) {
	return _City.Contract.AdminSetChengShiLevelAndSurety(&_City.TransactOpts, chengShiId_, level_, surety_)
}

// AdminSetChengShiLevelAndSurety is a paid mutator transaction binding the contract method 0xb9c74af7.
//
// Solidity: function adminSetChengShiLevelAndSurety(bytes32 chengShiId_, uint256 level_, uint256 surety_) returns()
func (_City *CityTransactorSession) AdminSetChengShiLevelAndSurety(chengShiId_ [32]byte, level_ *big.Int, surety_ *big.Int) (*types.Transaction, error) {
	return _City.Contract.AdminSetChengShiLevelAndSurety(&_City.TransactOpts, chengShiId_, level_, surety_)
}

// AdminSetCityPioneerAddress is a paid mutator transaction binding the contract method 0xa1532e85.
//
// Solidity: function adminSetCityPioneerAddress(address cityPioneerAddress_) returns()
func (_City *CityTransactor) AdminSetCityPioneerAddress(opts *bind.TransactOpts, cityPioneerAddress_ common.Address) (*types.Transaction, error) {
	return _City.contract.Transact(opts, "adminSetCityPioneerAddress", cityPioneerAddress_)
}

// AdminSetCityPioneerAddress is a paid mutator transaction binding the contract method 0xa1532e85.
//
// Solidity: function adminSetCityPioneerAddress(address cityPioneerAddress_) returns()
func (_City *CitySession) AdminSetCityPioneerAddress(cityPioneerAddress_ common.Address) (*types.Transaction, error) {
	return _City.Contract.AdminSetCityPioneerAddress(&_City.TransactOpts, cityPioneerAddress_)
}

// AdminSetCityPioneerAddress is a paid mutator transaction binding the contract method 0xa1532e85.
//
// Solidity: function adminSetCityPioneerAddress(address cityPioneerAddress_) returns()
func (_City *CityTransactorSession) AdminSetCityPioneerAddress(cityPioneerAddress_ common.Address) (*types.Transaction, error) {
	return _City.Contract.AdminSetCityPioneerAddress(&_City.TransactOpts, cityPioneerAddress_)
}

// AdminSetDelegate is a paid mutator transaction binding the contract method 0xe5317afa.
//
// Solidity: function adminSetDelegate(address userAddress_, uint256 amount_, uint256 setType) returns()
func (_City *CityTransactor) AdminSetDelegate(opts *bind.TransactOpts, userAddress_ common.Address, amount_ *big.Int, setType *big.Int) (*types.Transaction, error) {
	return _City.contract.Transact(opts, "adminSetDelegate", userAddress_, amount_, setType)
}

// AdminSetDelegate is a paid mutator transaction binding the contract method 0xe5317afa.
//
// Solidity: function adminSetDelegate(address userAddress_, uint256 amount_, uint256 setType) returns()
func (_City *CitySession) AdminSetDelegate(userAddress_ common.Address, amount_ *big.Int, setType *big.Int) (*types.Transaction, error) {
	return _City.Contract.AdminSetDelegate(&_City.TransactOpts, userAddress_, amount_, setType)
}

// AdminSetDelegate is a paid mutator transaction binding the contract method 0xe5317afa.
//
// Solidity: function adminSetDelegate(address userAddress_, uint256 amount_, uint256 setType) returns()
func (_City *CityTransactorSession) AdminSetDelegate(userAddress_ common.Address, amount_ *big.Int, setType *big.Int) (*types.Transaction, error) {
	return _City.Contract.AdminSetDelegate(&_City.TransactOpts, userAddress_, amount_, setType)
}

// AdminSetFoundsAddress is a paid mutator transaction binding the contract method 0xf823aa0e.
//
// Solidity: function adminSetFoundsAddress(address foundsAddress_) returns()
func (_City *CityTransactor) AdminSetFoundsAddress(opts *bind.TransactOpts, foundsAddress_ common.Address) (*types.Transaction, error) {
	return _City.contract.Transact(opts, "adminSetFoundsAddress", foundsAddress_)
}

// AdminSetFoundsAddress is a paid mutator transaction binding the contract method 0xf823aa0e.
//
// Solidity: function adminSetFoundsAddress(address foundsAddress_) returns()
func (_City *CitySession) AdminSetFoundsAddress(foundsAddress_ common.Address) (*types.Transaction, error) {
	return _City.Contract.AdminSetFoundsAddress(&_City.TransactOpts, foundsAddress_)
}

// AdminSetFoundsAddress is a paid mutator transaction binding the contract method 0xf823aa0e.
//
// Solidity: function adminSetFoundsAddress(address foundsAddress_) returns()
func (_City *CityTransactorSession) AdminSetFoundsAddress(foundsAddress_ common.Address) (*types.Transaction, error) {
	return _City.Contract.AdminSetFoundsAddress(&_City.TransactOpts, foundsAddress_)
}

// AdminSetPioneer is a paid mutator transaction binding the contract method 0x1305cad9.
//
// Solidity: function adminSetPioneer(bytes32 chengShiId_, address pioneer_, uint256 pioneerBatch_) returns()
func (_City *CityTransactor) AdminSetPioneer(opts *bind.TransactOpts, chengShiId_ [32]byte, pioneer_ common.Address, pioneerBatch_ *big.Int) (*types.Transaction, error) {
	return _City.contract.Transact(opts, "adminSetPioneer", chengShiId_, pioneer_, pioneerBatch_)
}

// AdminSetPioneer is a paid mutator transaction binding the contract method 0x1305cad9.
//
// Solidity: function adminSetPioneer(bytes32 chengShiId_, address pioneer_, uint256 pioneerBatch_) returns()
func (_City *CitySession) AdminSetPioneer(chengShiId_ [32]byte, pioneer_ common.Address, pioneerBatch_ *big.Int) (*types.Transaction, error) {
	return _City.Contract.AdminSetPioneer(&_City.TransactOpts, chengShiId_, pioneer_, pioneerBatch_)
}

// AdminSetPioneer is a paid mutator transaction binding the contract method 0x1305cad9.
//
// Solidity: function adminSetPioneer(bytes32 chengShiId_, address pioneer_, uint256 pioneerBatch_) returns()
func (_City *CityTransactorSession) AdminSetPioneer(chengShiId_ [32]byte, pioneer_ common.Address, pioneerBatch_ *big.Int) (*types.Transaction, error) {
	return _City.Contract.AdminSetPioneer(&_City.TransactOpts, chengShiId_, pioneer_, pioneerBatch_)
}

// AdminSetPioneerStatus is a paid mutator transaction binding the contract method 0xbaf1bc6a.
//
// Solidity: function adminSetPioneerStatus(address pioneer_, bool status_) returns()
func (_City *CityTransactor) AdminSetPioneerStatus(opts *bind.TransactOpts, pioneer_ common.Address, status_ bool) (*types.Transaction, error) {
	return _City.contract.Transact(opts, "adminSetPioneerStatus", pioneer_, status_)
}

// AdminSetPioneerStatus is a paid mutator transaction binding the contract method 0xbaf1bc6a.
//
// Solidity: function adminSetPioneerStatus(address pioneer_, bool status_) returns()
func (_City *CitySession) AdminSetPioneerStatus(pioneer_ common.Address, status_ bool) (*types.Transaction, error) {
	return _City.Contract.AdminSetPioneerStatus(&_City.TransactOpts, pioneer_, status_)
}

// AdminSetPioneerStatus is a paid mutator transaction binding the contract method 0xbaf1bc6a.
//
// Solidity: function adminSetPioneerStatus(address pioneer_, bool status_) returns()
func (_City *CityTransactorSession) AdminSetPioneerStatus(pioneer_ common.Address, status_ bool) (*types.Transaction, error) {
	return _City.Contract.AdminSetPioneerStatus(&_City.TransactOpts, pioneer_, status_)
}

// AdminSetRechargeAmount is a paid mutator transaction binding the contract method 0x5d048e61.
//
// Solidity: function adminSetRechargeAmount(address user, uint256 amount) returns()
func (_City *CityTransactor) AdminSetRechargeAmount(opts *bind.TransactOpts, user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _City.contract.Transact(opts, "adminSetRechargeAmount", user, amount)
}

// AdminSetRechargeAmount is a paid mutator transaction binding the contract method 0x5d048e61.
//
// Solidity: function adminSetRechargeAmount(address user, uint256 amount) returns()
func (_City *CitySession) AdminSetRechargeAmount(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _City.Contract.AdminSetRechargeAmount(&_City.TransactOpts, user, amount)
}

// AdminSetRechargeAmount is a paid mutator transaction binding the contract method 0x5d048e61.
//
// Solidity: function adminSetRechargeAmount(address user, uint256 amount) returns()
func (_City *CityTransactorSession) AdminSetRechargeAmount(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _City.Contract.AdminSetRechargeAmount(&_City.TransactOpts, user, amount)
}

// AdminSetRechargeWeightAdditional is a paid mutator transaction binding the contract method 0xa66dc64f.
//
// Solidity: function adminSetRechargeWeightAdditional(address pioneer_, uint256 weight_) returns()
func (_City *CityTransactor) AdminSetRechargeWeightAdditional(opts *bind.TransactOpts, pioneer_ common.Address, weight_ *big.Int) (*types.Transaction, error) {
	return _City.contract.Transact(opts, "adminSetRechargeWeightAdditional", pioneer_, weight_)
}

// AdminSetRechargeWeightAdditional is a paid mutator transaction binding the contract method 0xa66dc64f.
//
// Solidity: function adminSetRechargeWeightAdditional(address pioneer_, uint256 weight_) returns()
func (_City *CitySession) AdminSetRechargeWeightAdditional(pioneer_ common.Address, weight_ *big.Int) (*types.Transaction, error) {
	return _City.Contract.AdminSetRechargeWeightAdditional(&_City.TransactOpts, pioneer_, weight_)
}

// AdminSetRechargeWeightAdditional is a paid mutator transaction binding the contract method 0xa66dc64f.
//
// Solidity: function adminSetRechargeWeightAdditional(address pioneer_, uint256 weight_) returns()
func (_City *CityTransactorSession) AdminSetRechargeWeightAdditional(pioneer_ common.Address, weight_ *big.Int) (*types.Transaction, error) {
	return _City.Contract.AdminSetRechargeWeightAdditional(&_City.TransactOpts, pioneer_, weight_)
}

// AdminSetSecondsPerDay is a paid mutator transaction binding the contract method 0xace00fce.
//
// Solidity: function adminSetSecondsPerDay(uint56 secondsPerDay_) returns()
func (_City *CityTransactor) AdminSetSecondsPerDay(opts *bind.TransactOpts, secondsPerDay_ *big.Int) (*types.Transaction, error) {
	return _City.contract.Transact(opts, "adminSetSecondsPerDay", secondsPerDay_)
}

// AdminSetSecondsPerDay is a paid mutator transaction binding the contract method 0xace00fce.
//
// Solidity: function adminSetSecondsPerDay(uint56 secondsPerDay_) returns()
func (_City *CitySession) AdminSetSecondsPerDay(secondsPerDay_ *big.Int) (*types.Transaction, error) {
	return _City.Contract.AdminSetSecondsPerDay(&_City.TransactOpts, secondsPerDay_)
}

// AdminSetSecondsPerDay is a paid mutator transaction binding the contract method 0xace00fce.
//
// Solidity: function adminSetSecondsPerDay(uint56 secondsPerDay_) returns()
func (_City *CityTransactorSession) AdminSetSecondsPerDay(secondsPerDay_ *big.Int) (*types.Transaction, error) {
	return _City.Contract.AdminSetSecondsPerDay(&_City.TransactOpts, secondsPerDay_)
}

// AdminSetUserLocationAddress is a paid mutator transaction binding the contract method 0x267044e3.
//
// Solidity: function adminSetUserLocationAddress(address userLocationAddress_) returns()
func (_City *CityTransactor) AdminSetUserLocationAddress(opts *bind.TransactOpts, userLocationAddress_ common.Address) (*types.Transaction, error) {
	return _City.contract.Transact(opts, "adminSetUserLocationAddress", userLocationAddress_)
}

// AdminSetUserLocationAddress is a paid mutator transaction binding the contract method 0x267044e3.
//
// Solidity: function adminSetUserLocationAddress(address userLocationAddress_) returns()
func (_City *CitySession) AdminSetUserLocationAddress(userLocationAddress_ common.Address) (*types.Transaction, error) {
	return _City.Contract.AdminSetUserLocationAddress(&_City.TransactOpts, userLocationAddress_)
}

// AdminSetUserLocationAddress is a paid mutator transaction binding the contract method 0x267044e3.
//
// Solidity: function adminSetUserLocationAddress(address userLocationAddress_) returns()
func (_City *CityTransactorSession) AdminSetUserLocationAddress(userLocationAddress_ common.Address) (*types.Transaction, error) {
	return _City.Contract.AdminSetUserLocationAddress(&_City.TransactOpts, userLocationAddress_)
}

// AdminSetWithdrawLimitAddress is a paid mutator transaction binding the contract method 0x1746fa80.
//
// Solidity: function adminSetWithdrawLimitAddress(address withdrawLimitAddress_) returns()
func (_City *CityTransactor) AdminSetWithdrawLimitAddress(opts *bind.TransactOpts, withdrawLimitAddress_ common.Address) (*types.Transaction, error) {
	return _City.contract.Transact(opts, "adminSetWithdrawLimitAddress", withdrawLimitAddress_)
}

// AdminSetWithdrawLimitAddress is a paid mutator transaction binding the contract method 0x1746fa80.
//
// Solidity: function adminSetWithdrawLimitAddress(address withdrawLimitAddress_) returns()
func (_City *CitySession) AdminSetWithdrawLimitAddress(withdrawLimitAddress_ common.Address) (*types.Transaction, error) {
	return _City.Contract.AdminSetWithdrawLimitAddress(&_City.TransactOpts, withdrawLimitAddress_)
}

// AdminSetWithdrawLimitAddress is a paid mutator transaction binding the contract method 0x1746fa80.
//
// Solidity: function adminSetWithdrawLimitAddress(address withdrawLimitAddress_) returns()
func (_City *CityTransactorSession) AdminSetWithdrawLimitAddress(withdrawLimitAddress_ common.Address) (*types.Transaction, error) {
	return _City.Contract.AdminSetWithdrawLimitAddress(&_City.TransactOpts, withdrawLimitAddress_)
}

// BatchAddAdmin is a paid mutator transaction binding the contract method 0x2c9ab42b.
//
// Solidity: function batchAddAdmin(address[] amounts) returns()
func (_City *CityTransactor) BatchAddAdmin(opts *bind.TransactOpts, amounts []common.Address) (*types.Transaction, error) {
	return _City.contract.Transact(opts, "batchAddAdmin", amounts)
}

// BatchAddAdmin is a paid mutator transaction binding the contract method 0x2c9ab42b.
//
// Solidity: function batchAddAdmin(address[] amounts) returns()
func (_City *CitySession) BatchAddAdmin(amounts []common.Address) (*types.Transaction, error) {
	return _City.Contract.BatchAddAdmin(&_City.TransactOpts, amounts)
}

// BatchAddAdmin is a paid mutator transaction binding the contract method 0x2c9ab42b.
//
// Solidity: function batchAddAdmin(address[] amounts) returns()
func (_City *CityTransactorSession) BatchAddAdmin(amounts []common.Address) (*types.Transaction, error) {
	return _City.Contract.BatchAddAdmin(&_City.TransactOpts, amounts)
}

// ChangePioneerAddress is a paid mutator transaction binding the contract method 0xe820387e.
//
// Solidity: function changePioneerAddress(address newPioneerAddress_, address oldPioneerAddress_) returns()
func (_City *CityTransactor) ChangePioneerAddress(opts *bind.TransactOpts, newPioneerAddress_ common.Address, oldPioneerAddress_ common.Address) (*types.Transaction, error) {
	return _City.contract.Transact(opts, "changePioneerAddress", newPioneerAddress_, oldPioneerAddress_)
}

// ChangePioneerAddress is a paid mutator transaction binding the contract method 0xe820387e.
//
// Solidity: function changePioneerAddress(address newPioneerAddress_, address oldPioneerAddress_) returns()
func (_City *CitySession) ChangePioneerAddress(newPioneerAddress_ common.Address, oldPioneerAddress_ common.Address) (*types.Transaction, error) {
	return _City.Contract.ChangePioneerAddress(&_City.TransactOpts, newPioneerAddress_, oldPioneerAddress_)
}

// ChangePioneerAddress is a paid mutator transaction binding the contract method 0xe820387e.
//
// Solidity: function changePioneerAddress(address newPioneerAddress_, address oldPioneerAddress_) returns()
func (_City *CityTransactorSession) ChangePioneerAddress(newPioneerAddress_ common.Address, oldPioneerAddress_ common.Address) (*types.Transaction, error) {
	return _City.Contract.ChangePioneerAddress(&_City.TransactOpts, newPioneerAddress_, oldPioneerAddress_)
}

// DescCountyOrChengShiDelegate is a paid mutator transaction binding the contract method 0x4b8d74eb.
//
// Solidity: function descCountyOrChengShiDelegate(address user_, bytes32 countyId_, uint256 amount_, uint256 today) returns()
func (_City *CityTransactor) DescCountyOrChengShiDelegate(opts *bind.TransactOpts, user_ common.Address, countyId_ [32]byte, amount_ *big.Int, today *big.Int) (*types.Transaction, error) {
	return _City.contract.Transact(opts, "descCountyOrChengShiDelegate", user_, countyId_, amount_, today)
}

// DescCountyOrChengShiDelegate is a paid mutator transaction binding the contract method 0x4b8d74eb.
//
// Solidity: function descCountyOrChengShiDelegate(address user_, bytes32 countyId_, uint256 amount_, uint256 today) returns()
func (_City *CitySession) DescCountyOrChengShiDelegate(user_ common.Address, countyId_ [32]byte, amount_ *big.Int, today *big.Int) (*types.Transaction, error) {
	return _City.Contract.DescCountyOrChengShiDelegate(&_City.TransactOpts, user_, countyId_, amount_, today)
}

// DescCountyOrChengShiDelegate is a paid mutator transaction binding the contract method 0x4b8d74eb.
//
// Solidity: function descCountyOrChengShiDelegate(address user_, bytes32 countyId_, uint256 amount_, uint256 today) returns()
func (_City *CityTransactorSession) DescCountyOrChengShiDelegate(user_ common.Address, countyId_ [32]byte, amount_ *big.Int, today *big.Int) (*types.Transaction, error) {
	return _City.Contract.DescCountyOrChengShiDelegate(&_City.TransactOpts, user_, countyId_, amount_, today)
}

// IncrCountyOrChengShiDelegate is a paid mutator transaction binding the contract method 0xdf6cbf2b.
//
// Solidity: function incrCountyOrChengShiDelegate(address user_, bytes32 countyId_, uint256 amount_, uint256 today) returns()
func (_City *CityTransactor) IncrCountyOrChengShiDelegate(opts *bind.TransactOpts, user_ common.Address, countyId_ [32]byte, amount_ *big.Int, today *big.Int) (*types.Transaction, error) {
	return _City.contract.Transact(opts, "incrCountyOrChengShiDelegate", user_, countyId_, amount_, today)
}

// IncrCountyOrChengShiDelegate is a paid mutator transaction binding the contract method 0xdf6cbf2b.
//
// Solidity: function incrCountyOrChengShiDelegate(address user_, bytes32 countyId_, uint256 amount_, uint256 today) returns()
func (_City *CitySession) IncrCountyOrChengShiDelegate(user_ common.Address, countyId_ [32]byte, amount_ *big.Int, today *big.Int) (*types.Transaction, error) {
	return _City.Contract.IncrCountyOrChengShiDelegate(&_City.TransactOpts, user_, countyId_, amount_, today)
}

// IncrCountyOrChengShiDelegate is a paid mutator transaction binding the contract method 0xdf6cbf2b.
//
// Solidity: function incrCountyOrChengShiDelegate(address user_, bytes32 countyId_, uint256 amount_, uint256 today) returns()
func (_City *CityTransactorSession) IncrCountyOrChengShiDelegate(user_ common.Address, countyId_ [32]byte, amount_ *big.Int, today *big.Int) (*types.Transaction, error) {
	return _City.Contract.IncrCountyOrChengShiDelegate(&_City.TransactOpts, user_, countyId_, amount_, today)
}

// InitCityRechargeWeight is a paid mutator transaction binding the contract method 0x3e211318.
//
// Solidity: function initCityRechargeWeight(bytes32 chengShiId_) returns()
func (_City *CityTransactor) InitCityRechargeWeight(opts *bind.TransactOpts, chengShiId_ [32]byte) (*types.Transaction, error) {
	return _City.contract.Transact(opts, "initCityRechargeWeight", chengShiId_)
}

// InitCityRechargeWeight is a paid mutator transaction binding the contract method 0x3e211318.
//
// Solidity: function initCityRechargeWeight(bytes32 chengShiId_) returns()
func (_City *CitySession) InitCityRechargeWeight(chengShiId_ [32]byte) (*types.Transaction, error) {
	return _City.Contract.InitCityRechargeWeight(&_City.TransactOpts, chengShiId_)
}

// InitCityRechargeWeight is a paid mutator transaction binding the contract method 0x3e211318.
//
// Solidity: function initCityRechargeWeight(bytes32 chengShiId_) returns()
func (_City *CityTransactorSession) InitCityRechargeWeight(chengShiId_ [32]byte) (*types.Transaction, error) {
	return _City.Contract.InitCityRechargeWeight(&_City.TransactOpts, chengShiId_)
}

// PioneerDailyTask is a paid mutator transaction binding the contract method 0x1f5574b8.
//
// Solidity: function pioneerDailyTask(address pioneerAddress) returns()
func (_City *CityTransactor) PioneerDailyTask(opts *bind.TransactOpts, pioneerAddress common.Address) (*types.Transaction, error) {
	return _City.contract.Transact(opts, "pioneerDailyTask", pioneerAddress)
}

// PioneerDailyTask is a paid mutator transaction binding the contract method 0x1f5574b8.
//
// Solidity: function pioneerDailyTask(address pioneerAddress) returns()
func (_City *CitySession) PioneerDailyTask(pioneerAddress common.Address) (*types.Transaction, error) {
	return _City.Contract.PioneerDailyTask(&_City.TransactOpts, pioneerAddress)
}

// PioneerDailyTask is a paid mutator transaction binding the contract method 0x1f5574b8.
//
// Solidity: function pioneerDailyTask(address pioneerAddress) returns()
func (_City *CityTransactorSession) PioneerDailyTask(pioneerAddress common.Address) (*types.Transaction, error) {
	return _City.Contract.PioneerDailyTask(&_City.TransactOpts, pioneerAddress)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address account) returns()
func (_City *CityTransactor) RemoveAdmin(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _City.contract.Transact(opts, "removeAdmin", account)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address account) returns()
func (_City *CitySession) RemoveAdmin(account common.Address) (*types.Transaction, error) {
	return _City.Contract.RemoveAdmin(&_City.TransactOpts, account)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address account) returns()
func (_City *CityTransactorSession) RemoveAdmin(account common.Address) (*types.Transaction, error) {
	return _City.Contract.RemoveAdmin(&_City.TransactOpts, account)
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x8bad0c0a.
//
// Solidity: function renounceAdmin() returns()
func (_City *CityTransactor) RenounceAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _City.contract.Transact(opts, "renounceAdmin")
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x8bad0c0a.
//
// Solidity: function renounceAdmin() returns()
func (_City *CitySession) RenounceAdmin() (*types.Transaction, error) {
	return _City.Contract.RenounceAdmin(&_City.TransactOpts)
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x8bad0c0a.
//
// Solidity: function renounceAdmin() returns()
func (_City *CityTransactorSession) RenounceAdmin() (*types.Transaction, error) {
	return _City.Contract.RenounceAdmin(&_City.TransactOpts)
}

// SetChengShiPioneerAssessment is a paid mutator transaction binding the contract method 0x507a1c18.
//
// Solidity: function setChengShiPioneerAssessment(bytes32 chengShiId_) returns()
func (_City *CityTransactor) SetChengShiPioneerAssessment(opts *bind.TransactOpts, chengShiId_ [32]byte) (*types.Transaction, error) {
	return _City.contract.Transact(opts, "setChengShiPioneerAssessment", chengShiId_)
}

// SetChengShiPioneerAssessment is a paid mutator transaction binding the contract method 0x507a1c18.
//
// Solidity: function setChengShiPioneerAssessment(bytes32 chengShiId_) returns()
func (_City *CitySession) SetChengShiPioneerAssessment(chengShiId_ [32]byte) (*types.Transaction, error) {
	return _City.Contract.SetChengShiPioneerAssessment(&_City.TransactOpts, chengShiId_)
}

// SetChengShiPioneerAssessment is a paid mutator transaction binding the contract method 0x507a1c18.
//
// Solidity: function setChengShiPioneerAssessment(bytes32 chengShiId_) returns()
func (_City *CityTransactorSession) SetChengShiPioneerAssessment(chengShiId_ [32]byte) (*types.Transaction, error) {
	return _City.Contract.SetChengShiPioneerAssessment(&_City.TransactOpts, chengShiId_)
}

// SetCityMaxDelegate is a paid mutator transaction binding the contract method 0x8fd0b948.
//
// Solidity: function setCityMaxDelegate(bytes32 cityId_, uint256 amount_, uint256 day_) returns()
func (_City *CityTransactor) SetCityMaxDelegate(opts *bind.TransactOpts, cityId_ [32]byte, amount_ *big.Int, day_ *big.Int) (*types.Transaction, error) {
	return _City.contract.Transact(opts, "setCityMaxDelegate", cityId_, amount_, day_)
}

// SetCityMaxDelegate is a paid mutator transaction binding the contract method 0x8fd0b948.
//
// Solidity: function setCityMaxDelegate(bytes32 cityId_, uint256 amount_, uint256 day_) returns()
func (_City *CitySession) SetCityMaxDelegate(cityId_ [32]byte, amount_ *big.Int, day_ *big.Int) (*types.Transaction, error) {
	return _City.Contract.SetCityMaxDelegate(&_City.TransactOpts, cityId_, amount_, day_)
}

// SetCityMaxDelegate is a paid mutator transaction binding the contract method 0x8fd0b948.
//
// Solidity: function setCityMaxDelegate(bytes32 cityId_, uint256 amount_, uint256 day_) returns()
func (_City *CityTransactorSession) SetCityMaxDelegate(cityId_ [32]byte, amount_ *big.Int, day_ *big.Int) (*types.Transaction, error) {
	return _City.Contract.SetCityMaxDelegate(&_City.TransactOpts, cityId_, amount_, day_)
}

// CityAdminAddedIterator is returned from FilterAdminAdded and is used to iterate over the raw logs and unpacked data for AdminAdded events raised by the City contract.
type CityAdminAddedIterator struct {
	Event *CityAdminAdded // Event containing the contract specifics and raw log

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
func (it *CityAdminAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CityAdminAdded)
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
		it.Event = new(CityAdminAdded)
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
func (it *CityAdminAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CityAdminAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CityAdminAdded represents a AdminAdded event raised by the City contract.
type CityAdminAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAdminAdded is a free log retrieval operation binding the contract event 0x44d6d25963f097ad14f29f06854a01f575648a1ef82f30e562ccd3889717e339.
//
// Solidity: event AdminAdded(address indexed account)
func (_City *CityFilterer) FilterAdminAdded(opts *bind.FilterOpts, account []common.Address) (*CityAdminAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _City.contract.FilterLogs(opts, "AdminAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &CityAdminAddedIterator{contract: _City.contract, event: "AdminAdded", logs: logs, sub: sub}, nil
}

// WatchAdminAdded is a free log subscription operation binding the contract event 0x44d6d25963f097ad14f29f06854a01f575648a1ef82f30e562ccd3889717e339.
//
// Solidity: event AdminAdded(address indexed account)
func (_City *CityFilterer) WatchAdminAdded(opts *bind.WatchOpts, sink chan<- *CityAdminAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _City.contract.WatchLogs(opts, "AdminAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CityAdminAdded)
				if err := _City.contract.UnpackLog(event, "AdminAdded", log); err != nil {
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
func (_City *CityFilterer) ParseAdminAdded(log types.Log) (*CityAdminAdded, error) {
	event := new(CityAdminAdded)
	if err := _City.contract.UnpackLog(event, "AdminAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CityAdminRemovedIterator is returned from FilterAdminRemoved and is used to iterate over the raw logs and unpacked data for AdminRemoved events raised by the City contract.
type CityAdminRemovedIterator struct {
	Event *CityAdminRemoved // Event containing the contract specifics and raw log

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
func (it *CityAdminRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CityAdminRemoved)
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
		it.Event = new(CityAdminRemoved)
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
func (it *CityAdminRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CityAdminRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CityAdminRemoved represents a AdminRemoved event raised by the City contract.
type CityAdminRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAdminRemoved is a free log retrieval operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: event AdminRemoved(address indexed account)
func (_City *CityFilterer) FilterAdminRemoved(opts *bind.FilterOpts, account []common.Address) (*CityAdminRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _City.contract.FilterLogs(opts, "AdminRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &CityAdminRemovedIterator{contract: _City.contract, event: "AdminRemoved", logs: logs, sub: sub}, nil
}

// WatchAdminRemoved is a free log subscription operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: event AdminRemoved(address indexed account)
func (_City *CityFilterer) WatchAdminRemoved(opts *bind.WatchOpts, sink chan<- *CityAdminRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _City.contract.WatchLogs(opts, "AdminRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CityAdminRemoved)
				if err := _City.contract.UnpackLog(event, "AdminRemoved", log); err != nil {
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
func (_City *CityFilterer) ParseAdminRemoved(log types.Log) (*CityAdminRemoved, error) {
	event := new(CityAdminRemoved)
	if err := _City.contract.UnpackLog(event, "AdminRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CityDecreaseCityDelegateIterator is returned from FilterDecreaseCityDelegate and is used to iterate over the raw logs and unpacked data for DecreaseCityDelegate events raised by the City contract.
type CityDecreaseCityDelegateIterator struct {
	Event *CityDecreaseCityDelegate // Event containing the contract specifics and raw log

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
func (it *CityDecreaseCityDelegateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CityDecreaseCityDelegate)
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
		it.Event = new(CityDecreaseCityDelegate)
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
func (it *CityDecreaseCityDelegateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CityDecreaseCityDelegateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CityDecreaseCityDelegate represents a DecreaseCityDelegate event raised by the City contract.
type CityDecreaseCityDelegate struct {
	User   common.Address
	CityId [32]byte
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDecreaseCityDelegate is a free log retrieval operation binding the contract event 0x726ee283c47c8203b92dce3a1964d78c143bb26d63d4ce5f1ec3cdf61e1e1878.
//
// Solidity: event DecreaseCityDelegate(address user, bytes32 cityId, uint256 amount)
func (_City *CityFilterer) FilterDecreaseCityDelegate(opts *bind.FilterOpts) (*CityDecreaseCityDelegateIterator, error) {

	logs, sub, err := _City.contract.FilterLogs(opts, "DecreaseCityDelegate")
	if err != nil {
		return nil, err
	}
	return &CityDecreaseCityDelegateIterator{contract: _City.contract, event: "DecreaseCityDelegate", logs: logs, sub: sub}, nil
}

// WatchDecreaseCityDelegate is a free log subscription operation binding the contract event 0x726ee283c47c8203b92dce3a1964d78c143bb26d63d4ce5f1ec3cdf61e1e1878.
//
// Solidity: event DecreaseCityDelegate(address user, bytes32 cityId, uint256 amount)
func (_City *CityFilterer) WatchDecreaseCityDelegate(opts *bind.WatchOpts, sink chan<- *CityDecreaseCityDelegate) (event.Subscription, error) {

	logs, sub, err := _City.contract.WatchLogs(opts, "DecreaseCityDelegate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CityDecreaseCityDelegate)
				if err := _City.contract.UnpackLog(event, "DecreaseCityDelegate", log); err != nil {
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

// ParseDecreaseCityDelegate is a log parse operation binding the contract event 0x726ee283c47c8203b92dce3a1964d78c143bb26d63d4ce5f1ec3cdf61e1e1878.
//
// Solidity: event DecreaseCityDelegate(address user, bytes32 cityId, uint256 amount)
func (_City *CityFilterer) ParseDecreaseCityDelegate(log types.Log) (*CityDecreaseCityDelegate, error) {
	event := new(CityDecreaseCityDelegate)
	if err := _City.contract.UnpackLog(event, "DecreaseCityDelegate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CityIncreaseCityDelegateIterator is returned from FilterIncreaseCityDelegate and is used to iterate over the raw logs and unpacked data for IncreaseCityDelegate events raised by the City contract.
type CityIncreaseCityDelegateIterator struct {
	Event *CityIncreaseCityDelegate // Event containing the contract specifics and raw log

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
func (it *CityIncreaseCityDelegateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CityIncreaseCityDelegate)
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
		it.Event = new(CityIncreaseCityDelegate)
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
func (it *CityIncreaseCityDelegateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CityIncreaseCityDelegateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CityIncreaseCityDelegate represents a IncreaseCityDelegate event raised by the City contract.
type CityIncreaseCityDelegate struct {
	User   common.Address
	CityId [32]byte
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterIncreaseCityDelegate is a free log retrieval operation binding the contract event 0xa7900f2ee8d29e3abedc3205badaccdade1f4df880caa506193cec45e2c71c7e.
//
// Solidity: event IncreaseCityDelegate(address user, bytes32 cityId, uint256 amount)
func (_City *CityFilterer) FilterIncreaseCityDelegate(opts *bind.FilterOpts) (*CityIncreaseCityDelegateIterator, error) {

	logs, sub, err := _City.contract.FilterLogs(opts, "IncreaseCityDelegate")
	if err != nil {
		return nil, err
	}
	return &CityIncreaseCityDelegateIterator{contract: _City.contract, event: "IncreaseCityDelegate", logs: logs, sub: sub}, nil
}

// WatchIncreaseCityDelegate is a free log subscription operation binding the contract event 0xa7900f2ee8d29e3abedc3205badaccdade1f4df880caa506193cec45e2c71c7e.
//
// Solidity: event IncreaseCityDelegate(address user, bytes32 cityId, uint256 amount)
func (_City *CityFilterer) WatchIncreaseCityDelegate(opts *bind.WatchOpts, sink chan<- *CityIncreaseCityDelegate) (event.Subscription, error) {

	logs, sub, err := _City.contract.WatchLogs(opts, "IncreaseCityDelegate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CityIncreaseCityDelegate)
				if err := _City.contract.UnpackLog(event, "IncreaseCityDelegate", log); err != nil {
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

// ParseIncreaseCityDelegate is a log parse operation binding the contract event 0xa7900f2ee8d29e3abedc3205badaccdade1f4df880caa506193cec45e2c71c7e.
//
// Solidity: event IncreaseCityDelegate(address user, bytes32 cityId, uint256 amount)
func (_City *CityFilterer) ParseIncreaseCityDelegate(log types.Log) (*CityIncreaseCityDelegate, error) {
	event := new(CityIncreaseCityDelegate)
	if err := _City.contract.UnpackLog(event, "IncreaseCityDelegate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CityInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the City contract.
type CityInitializedIterator struct {
	Event *CityInitialized // Event containing the contract specifics and raw log

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
func (it *CityInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CityInitialized)
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
		it.Event = new(CityInitialized)
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
func (it *CityInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CityInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CityInitialized represents a Initialized event raised by the City contract.
type CityInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_City *CityFilterer) FilterInitialized(opts *bind.FilterOpts) (*CityInitializedIterator, error) {

	logs, sub, err := _City.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &CityInitializedIterator{contract: _City.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_City *CityFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *CityInitialized) (event.Subscription, error) {

	logs, sub, err := _City.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CityInitialized)
				if err := _City.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_City *CityFilterer) ParseInitialized(log types.Log) (*CityInitialized, error) {
	event := new(CityInitialized)
	if err := _City.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CityPioneerBlackIterator is returned from FilterPioneerBlack and is used to iterate over the raw logs and unpacked data for PioneerBlack events raised by the City contract.
type CityPioneerBlackIterator struct {
	Event *CityPioneerBlack // Event containing the contract specifics and raw log

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
func (it *CityPioneerBlackIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CityPioneerBlack)
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
		it.Event = new(CityPioneerBlack)
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
func (it *CityPioneerBlackIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CityPioneerBlackIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CityPioneerBlack represents a PioneerBlack event raised by the City contract.
type CityPioneerBlack struct {
	Pioneer common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPioneerBlack is a free log retrieval operation binding the contract event 0x7b69a9dc9c659903e8fa03173afce6d7212cbc8f69cc769730a061d7094f13aa.
//
// Solidity: event PioneerBlack(address pioneer)
func (_City *CityFilterer) FilterPioneerBlack(opts *bind.FilterOpts) (*CityPioneerBlackIterator, error) {

	logs, sub, err := _City.contract.FilterLogs(opts, "PioneerBlack")
	if err != nil {
		return nil, err
	}
	return &CityPioneerBlackIterator{contract: _City.contract, event: "PioneerBlack", logs: logs, sub: sub}, nil
}

// WatchPioneerBlack is a free log subscription operation binding the contract event 0x7b69a9dc9c659903e8fa03173afce6d7212cbc8f69cc769730a061d7094f13aa.
//
// Solidity: event PioneerBlack(address pioneer)
func (_City *CityFilterer) WatchPioneerBlack(opts *bind.WatchOpts, sink chan<- *CityPioneerBlack) (event.Subscription, error) {

	logs, sub, err := _City.contract.WatchLogs(opts, "PioneerBlack")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CityPioneerBlack)
				if err := _City.contract.UnpackLog(event, "PioneerBlack", log); err != nil {
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

// ParsePioneerBlack is a log parse operation binding the contract event 0x7b69a9dc9c659903e8fa03173afce6d7212cbc8f69cc769730a061d7094f13aa.
//
// Solidity: event PioneerBlack(address pioneer)
func (_City *CityFilterer) ParsePioneerBlack(log types.Log) (*CityPioneerBlack, error) {
	event := new(CityPioneerBlack)
	if err := _City.contract.UnpackLog(event, "PioneerBlack", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CityRechargeRecordIterator is returned from FilterRechargeRecord and is used to iterate over the raw logs and unpacked data for RechargeRecord events raised by the City contract.
type CityRechargeRecordIterator struct {
	Event *CityRechargeRecord // Event containing the contract specifics and raw log

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
func (it *CityRechargeRecordIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CityRechargeRecord)
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
		it.Event = new(CityRechargeRecord)
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
func (it *CityRechargeRecordIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CityRechargeRecordIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CityRechargeRecord represents a RechargeRecord event raised by the City contract.
type CityRechargeRecord struct {
	User     common.Address
	CountyId [32]byte
	Amount   *big.Int
	Time     *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRechargeRecord is a free log retrieval operation binding the contract event 0xb3e6f770a2a15dbbd8fd48dab33d9918a43e5a3c50669f989e2a1d1d62615dff.
//
// Solidity: event RechargeRecord(address user, bytes32 countyId, uint256 amount, uint256 time)
func (_City *CityFilterer) FilterRechargeRecord(opts *bind.FilterOpts) (*CityRechargeRecordIterator, error) {

	logs, sub, err := _City.contract.FilterLogs(opts, "RechargeRecord")
	if err != nil {
		return nil, err
	}
	return &CityRechargeRecordIterator{contract: _City.contract, event: "RechargeRecord", logs: logs, sub: sub}, nil
}

// WatchRechargeRecord is a free log subscription operation binding the contract event 0xb3e6f770a2a15dbbd8fd48dab33d9918a43e5a3c50669f989e2a1d1d62615dff.
//
// Solidity: event RechargeRecord(address user, bytes32 countyId, uint256 amount, uint256 time)
func (_City *CityFilterer) WatchRechargeRecord(opts *bind.WatchOpts, sink chan<- *CityRechargeRecord) (event.Subscription, error) {

	logs, sub, err := _City.contract.WatchLogs(opts, "RechargeRecord")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CityRechargeRecord)
				if err := _City.contract.UnpackLog(event, "RechargeRecord", log); err != nil {
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

// ParseRechargeRecord is a log parse operation binding the contract event 0xb3e6f770a2a15dbbd8fd48dab33d9918a43e5a3c50669f989e2a1d1d62615dff.
//
// Solidity: event RechargeRecord(address user, bytes32 countyId, uint256 amount, uint256 time)
func (_City *CityFilterer) ParseRechargeRecord(log types.Log) (*CityRechargeRecord, error) {
	event := new(CityRechargeRecord)
	if err := _City.contract.UnpackLog(event, "RechargeRecord", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
