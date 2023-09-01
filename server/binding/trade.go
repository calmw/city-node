// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package trade

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

// TradeOrder is an auto generated low-level Go binding around an user-defined struct.
type TradeOrder struct {
	Id       *big.Int
	Seller   common.Address
	Buyer    common.Address
	Price    *big.Int
	Quantity *big.Int
	IsDael   bool
	SaleTime *big.Int
	BuyTime  *big.Int
}

// TradeMetaData contains all meta data concerning the Trade contract.
var TradeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"saleTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"buyTime\",\"type\":\"uint256\"}],\"name\":\"Buy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"saleTime\",\"type\":\"uint256\"}],\"name\":\"Sale\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"adminSetOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress_\",\"type\":\"address\"}],\"name\":\"adminSetTokenAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderId_\",\"type\":\"uint256\"}],\"name\":\"buy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderId_\",\"type\":\"uint256\"}],\"name\":\"getOrderInfoById\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isDael\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"saleTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"buyTime\",\"type\":\"uint256\"}],\"internalType\":\"structTrade.Order\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"orderInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isDael\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"saleTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"buyTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"sale\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// TradeABI is the input ABI used to generate the binding from.
// Deprecated: Use TradeMetaData.ABI instead.
var TradeABI = TradeMetaData.ABI

// Trade is an auto generated Go binding around an Ethereum contract.
type Trade struct {
	TradeCaller     // Read-only binding to the contract
	TradeTransactor // Write-only binding to the contract
	TradeFilterer   // Log filterer for contract events
}

// TradeCaller is an auto generated read-only Go binding around an Ethereum contract.
type TradeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TradeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TradeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TradeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TradeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TradeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TradeSession struct {
	Contract     *Trade            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TradeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TradeCallerSession struct {
	Contract *TradeCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TradeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TradeTransactorSession struct {
	Contract     *TradeTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TradeRaw is an auto generated low-level Go binding around an Ethereum contract.
type TradeRaw struct {
	Contract *Trade // Generic contract binding to access the raw methods on
}

// TradeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TradeCallerRaw struct {
	Contract *TradeCaller // Generic read-only contract binding to access the raw methods on
}

// TradeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TradeTransactorRaw struct {
	Contract *TradeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTrade creates a new instance of Trade, bound to a specific deployed contract.
func NewTrade(address common.Address, backend bind.ContractBackend) (*Trade, error) {
	contract, err := bindTrade(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Trade{TradeCaller: TradeCaller{contract: contract}, TradeTransactor: TradeTransactor{contract: contract}, TradeFilterer: TradeFilterer{contract: contract}}, nil
}

// NewTradeCaller creates a new read-only instance of Trade, bound to a specific deployed contract.
func NewTradeCaller(address common.Address, caller bind.ContractCaller) (*TradeCaller, error) {
	contract, err := bindTrade(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TradeCaller{contract: contract}, nil
}

// NewTradeTransactor creates a new write-only instance of Trade, bound to a specific deployed contract.
func NewTradeTransactor(address common.Address, transactor bind.ContractTransactor) (*TradeTransactor, error) {
	contract, err := bindTrade(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TradeTransactor{contract: contract}, nil
}

// NewTradeFilterer creates a new log filterer instance of Trade, bound to a specific deployed contract.
func NewTradeFilterer(address common.Address, filterer bind.ContractFilterer) (*TradeFilterer, error) {
	contract, err := bindTrade(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TradeFilterer{contract: contract}, nil
}

// bindTrade binds a generic wrapper to an already deployed contract.
func bindTrade(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TradeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Trade *TradeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Trade.Contract.TradeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Trade *TradeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trade.Contract.TradeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Trade *TradeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Trade.Contract.TradeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Trade *TradeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Trade.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Trade *TradeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trade.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Trade *TradeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Trade.Contract.contract.Transact(opts, method, params...)
}

// GetOrderInfoById is a free data retrieval call binding the contract method 0x5bef7212.
//
// Solidity: function getOrderInfoById(uint256 orderId_) view returns((uint256,address,address,uint256,uint256,bool,uint256,uint256))
func (_Trade *TradeCaller) GetOrderInfoById(opts *bind.CallOpts, orderId_ *big.Int) (TradeOrder, error) {
	var out []interface{}
	err := _Trade.contract.Call(opts, &out, "getOrderInfoById", orderId_)

	if err != nil {
		return *new(TradeOrder), err
	}

	out0 := *abi.ConvertType(out[0], new(TradeOrder)).(*TradeOrder)

	return out0, err

}

// GetOrderInfoById is a free data retrieval call binding the contract method 0x5bef7212.
//
// Solidity: function getOrderInfoById(uint256 orderId_) view returns((uint256,address,address,uint256,uint256,bool,uint256,uint256))
func (_Trade *TradeSession) GetOrderInfoById(orderId_ *big.Int) (TradeOrder, error) {
	return _Trade.Contract.GetOrderInfoById(&_Trade.CallOpts, orderId_)
}

// GetOrderInfoById is a free data retrieval call binding the contract method 0x5bef7212.
//
// Solidity: function getOrderInfoById(uint256 orderId_) view returns((uint256,address,address,uint256,uint256,bool,uint256,uint256))
func (_Trade *TradeCallerSession) GetOrderInfoById(orderId_ *big.Int) (TradeOrder, error) {
	return _Trade.Contract.GetOrderInfoById(&_Trade.CallOpts, orderId_)
}

// OrderInfo is a free data retrieval call binding the contract method 0x6ab34a4a.
//
// Solidity: function orderInfo(uint256 ) view returns(uint256 id, address seller, address buyer, uint256 price, uint256 quantity, bool isDael, uint256 saleTime, uint256 buyTime)
func (_Trade *TradeCaller) OrderInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id       *big.Int
	Seller   common.Address
	Buyer    common.Address
	Price    *big.Int
	Quantity *big.Int
	IsDael   bool
	SaleTime *big.Int
	BuyTime  *big.Int
}, error) {
	var out []interface{}
	err := _Trade.contract.Call(opts, &out, "orderInfo", arg0)

	outstruct := new(struct {
		Id       *big.Int
		Seller   common.Address
		Buyer    common.Address
		Price    *big.Int
		Quantity *big.Int
		IsDael   bool
		SaleTime *big.Int
		BuyTime  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Seller = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Buyer = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Price = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Quantity = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.IsDael = *abi.ConvertType(out[5], new(bool)).(*bool)
	outstruct.SaleTime = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.BuyTime = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// OrderInfo is a free data retrieval call binding the contract method 0x6ab34a4a.
//
// Solidity: function orderInfo(uint256 ) view returns(uint256 id, address seller, address buyer, uint256 price, uint256 quantity, bool isDael, uint256 saleTime, uint256 buyTime)
func (_Trade *TradeSession) OrderInfo(arg0 *big.Int) (struct {
	Id       *big.Int
	Seller   common.Address
	Buyer    common.Address
	Price    *big.Int
	Quantity *big.Int
	IsDael   bool
	SaleTime *big.Int
	BuyTime  *big.Int
}, error) {
	return _Trade.Contract.OrderInfo(&_Trade.CallOpts, arg0)
}

// OrderInfo is a free data retrieval call binding the contract method 0x6ab34a4a.
//
// Solidity: function orderInfo(uint256 ) view returns(uint256 id, address seller, address buyer, uint256 price, uint256 quantity, bool isDael, uint256 saleTime, uint256 buyTime)
func (_Trade *TradeCallerSession) OrderInfo(arg0 *big.Int) (struct {
	Id       *big.Int
	Seller   common.Address
	Buyer    common.Address
	Price    *big.Int
	Quantity *big.Int
	IsDael   bool
	SaleTime *big.Int
	BuyTime  *big.Int
}, error) {
	return _Trade.Contract.OrderInfo(&_Trade.CallOpts, arg0)
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() view returns(address)
func (_Trade *TradeCaller) TokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Trade.contract.Call(opts, &out, "tokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() view returns(address)
func (_Trade *TradeSession) TokenAddress() (common.Address, error) {
	return _Trade.Contract.TokenAddress(&_Trade.CallOpts)
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() view returns(address)
func (_Trade *TradeCallerSession) TokenAddress() (common.Address, error) {
	return _Trade.Contract.TokenAddress(&_Trade.CallOpts)
}

// AdminSetOwner is a paid mutator transaction binding the contract method 0x824744dd.
//
// Solidity: function adminSetOwner(address owner_) returns()
func (_Trade *TradeTransactor) AdminSetOwner(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _Trade.contract.Transact(opts, "adminSetOwner", owner_)
}

// AdminSetOwner is a paid mutator transaction binding the contract method 0x824744dd.
//
// Solidity: function adminSetOwner(address owner_) returns()
func (_Trade *TradeSession) AdminSetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _Trade.Contract.AdminSetOwner(&_Trade.TransactOpts, owner_)
}

// AdminSetOwner is a paid mutator transaction binding the contract method 0x824744dd.
//
// Solidity: function adminSetOwner(address owner_) returns()
func (_Trade *TradeTransactorSession) AdminSetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _Trade.Contract.AdminSetOwner(&_Trade.TransactOpts, owner_)
}

// AdminSetTokenAddress is a paid mutator transaction binding the contract method 0x559a89e0.
//
// Solidity: function adminSetTokenAddress(address tokenAddress_) returns()
func (_Trade *TradeTransactor) AdminSetTokenAddress(opts *bind.TransactOpts, tokenAddress_ common.Address) (*types.Transaction, error) {
	return _Trade.contract.Transact(opts, "adminSetTokenAddress", tokenAddress_)
}

// AdminSetTokenAddress is a paid mutator transaction binding the contract method 0x559a89e0.
//
// Solidity: function adminSetTokenAddress(address tokenAddress_) returns()
func (_Trade *TradeSession) AdminSetTokenAddress(tokenAddress_ common.Address) (*types.Transaction, error) {
	return _Trade.Contract.AdminSetTokenAddress(&_Trade.TransactOpts, tokenAddress_)
}

// AdminSetTokenAddress is a paid mutator transaction binding the contract method 0x559a89e0.
//
// Solidity: function adminSetTokenAddress(address tokenAddress_) returns()
func (_Trade *TradeTransactorSession) AdminSetTokenAddress(tokenAddress_ common.Address) (*types.Transaction, error) {
	return _Trade.Contract.AdminSetTokenAddress(&_Trade.TransactOpts, tokenAddress_)
}

// Buy is a paid mutator transaction binding the contract method 0xd96a094a.
//
// Solidity: function buy(uint256 orderId_) returns()
func (_Trade *TradeTransactor) Buy(opts *bind.TransactOpts, orderId_ *big.Int) (*types.Transaction, error) {
	return _Trade.contract.Transact(opts, "buy", orderId_)
}

// Buy is a paid mutator transaction binding the contract method 0xd96a094a.
//
// Solidity: function buy(uint256 orderId_) returns()
func (_Trade *TradeSession) Buy(orderId_ *big.Int) (*types.Transaction, error) {
	return _Trade.Contract.Buy(&_Trade.TransactOpts, orderId_)
}

// Buy is a paid mutator transaction binding the contract method 0xd96a094a.
//
// Solidity: function buy(uint256 orderId_) returns()
func (_Trade *TradeTransactorSession) Buy(orderId_ *big.Int) (*types.Transaction, error) {
	return _Trade.Contract.Buy(&_Trade.TransactOpts, orderId_)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Trade *TradeTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trade.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Trade *TradeSession) Initialize() (*types.Transaction, error) {
	return _Trade.Contract.Initialize(&_Trade.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Trade *TradeTransactorSession) Initialize() (*types.Transaction, error) {
	return _Trade.Contract.Initialize(&_Trade.TransactOpts)
}

// Sale is a paid mutator transaction binding the contract method 0x1c35f962.
//
// Solidity: function sale(uint256 quantity, uint256 price) returns()
func (_Trade *TradeTransactor) Sale(opts *bind.TransactOpts, quantity *big.Int, price *big.Int) (*types.Transaction, error) {
	return _Trade.contract.Transact(opts, "sale", quantity, price)
}

// Sale is a paid mutator transaction binding the contract method 0x1c35f962.
//
// Solidity: function sale(uint256 quantity, uint256 price) returns()
func (_Trade *TradeSession) Sale(quantity *big.Int, price *big.Int) (*types.Transaction, error) {
	return _Trade.Contract.Sale(&_Trade.TransactOpts, quantity, price)
}

// Sale is a paid mutator transaction binding the contract method 0x1c35f962.
//
// Solidity: function sale(uint256 quantity, uint256 price) returns()
func (_Trade *TradeTransactorSession) Sale(quantity *big.Int, price *big.Int) (*types.Transaction, error) {
	return _Trade.Contract.Sale(&_Trade.TransactOpts, quantity, price)
}

// TradeBuyIterator is returned from FilterBuy and is used to iterate over the raw logs and unpacked data for Buy events raised by the Trade contract.
type TradeBuyIterator struct {
	Event *TradeBuy // Event containing the contract specifics and raw log

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
func (it *TradeBuyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradeBuy)
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
		it.Event = new(TradeBuy)
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
func (it *TradeBuyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradeBuyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradeBuy represents a Buy event raised by the Trade contract.
type TradeBuy struct {
	OrderId  *big.Int
	Seller   common.Address
	Buyer    common.Address
	Price    *big.Int
	Quantity *big.Int
	SaleTime *big.Int
	BuyTime  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBuy is a free log retrieval operation binding the contract event 0xb2b1af4e7bd168298631b9bd6d83e147029d0fa009f109d0160aadc82823b299.
//
// Solidity: event Buy(uint256 indexed orderId, address seller, address buyer, uint256 price, uint256 quantity, uint256 saleTime, uint256 buyTime)
func (_Trade *TradeFilterer) FilterBuy(opts *bind.FilterOpts, orderId []*big.Int) (*TradeBuyIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _Trade.contract.FilterLogs(opts, "Buy", orderIdRule)
	if err != nil {
		return nil, err
	}
	return &TradeBuyIterator{contract: _Trade.contract, event: "Buy", logs: logs, sub: sub}, nil
}

// WatchBuy is a free log subscription operation binding the contract event 0xb2b1af4e7bd168298631b9bd6d83e147029d0fa009f109d0160aadc82823b299.
//
// Solidity: event Buy(uint256 indexed orderId, address seller, address buyer, uint256 price, uint256 quantity, uint256 saleTime, uint256 buyTime)
func (_Trade *TradeFilterer) WatchBuy(opts *bind.WatchOpts, sink chan<- *TradeBuy, orderId []*big.Int) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _Trade.contract.WatchLogs(opts, "Buy", orderIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradeBuy)
				if err := _Trade.contract.UnpackLog(event, "Buy", log); err != nil {
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

// ParseBuy is a log parse operation binding the contract event 0xb2b1af4e7bd168298631b9bd6d83e147029d0fa009f109d0160aadc82823b299.
//
// Solidity: event Buy(uint256 indexed orderId, address seller, address buyer, uint256 price, uint256 quantity, uint256 saleTime, uint256 buyTime)
func (_Trade *TradeFilterer) ParseBuy(log types.Log) (*TradeBuy, error) {
	event := new(TradeBuy)
	if err := _Trade.contract.UnpackLog(event, "Buy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TradeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Trade contract.
type TradeInitializedIterator struct {
	Event *TradeInitialized // Event containing the contract specifics and raw log

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
func (it *TradeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradeInitialized)
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
		it.Event = new(TradeInitialized)
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
func (it *TradeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradeInitialized represents a Initialized event raised by the Trade contract.
type TradeInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Trade *TradeFilterer) FilterInitialized(opts *bind.FilterOpts) (*TradeInitializedIterator, error) {

	logs, sub, err := _Trade.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TradeInitializedIterator{contract: _Trade.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Trade *TradeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TradeInitialized) (event.Subscription, error) {

	logs, sub, err := _Trade.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradeInitialized)
				if err := _Trade.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Trade *TradeFilterer) ParseInitialized(log types.Log) (*TradeInitialized, error) {
	event := new(TradeInitialized)
	if err := _Trade.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TradeSaleIterator is returned from FilterSale and is used to iterate over the raw logs and unpacked data for Sale events raised by the Trade contract.
type TradeSaleIterator struct {
	Event *TradeSale // Event containing the contract specifics and raw log

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
func (it *TradeSaleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradeSale)
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
		it.Event = new(TradeSale)
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
func (it *TradeSaleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradeSaleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradeSale represents a Sale event raised by the Trade contract.
type TradeSale struct {
	OrderId  *big.Int
	Seller   common.Address
	Price    *big.Int
	Quantity *big.Int
	SaleTime *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSale is a free log retrieval operation binding the contract event 0xa9646de4555cf5a664323384c5a0a2668cd32f009d8db246c31b66a914c2d894.
//
// Solidity: event Sale(uint256 indexed orderId, address seller, uint256 price, uint256 quantity, uint256 saleTime)
func (_Trade *TradeFilterer) FilterSale(opts *bind.FilterOpts, orderId []*big.Int) (*TradeSaleIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _Trade.contract.FilterLogs(opts, "Sale", orderIdRule)
	if err != nil {
		return nil, err
	}
	return &TradeSaleIterator{contract: _Trade.contract, event: "Sale", logs: logs, sub: sub}, nil
}

// WatchSale is a free log subscription operation binding the contract event 0xa9646de4555cf5a664323384c5a0a2668cd32f009d8db246c31b66a914c2d894.
//
// Solidity: event Sale(uint256 indexed orderId, address seller, uint256 price, uint256 quantity, uint256 saleTime)
func (_Trade *TradeFilterer) WatchSale(opts *bind.WatchOpts, sink chan<- *TradeSale, orderId []*big.Int) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _Trade.contract.WatchLogs(opts, "Sale", orderIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradeSale)
				if err := _Trade.contract.UnpackLog(event, "Sale", log); err != nil {
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

// ParseSale is a log parse operation binding the contract event 0xa9646de4555cf5a664323384c5a0a2668cd32f009d8db246c31b66a914c2d894.
//
// Solidity: event Sale(uint256 indexed orderId, address seller, uint256 price, uint256 quantity, uint256 saleTime)
func (_Trade *TradeFilterer) ParseSale(log types.Log) (*TradeSale, error) {
	event := new(TradeSale)
	if err := _Trade.contract.UnpackLog(event, "Sale", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
