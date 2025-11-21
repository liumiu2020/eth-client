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
	_ = abi.ConvertType
)

// BookMetaData contains all meta data concerning the Book contract.
var BookMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"bid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"judge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BookABI is the input ABI used to generate the binding from.
// Deprecated: Use BookMetaData.ABI instead.
var BookABI = BookMetaData.ABI

// Book is an auto generated Go binding around an Ethereum contract.
type Book struct {
	BookCaller     // Read-only binding to the contract
	BookTransactor // Write-only binding to the contract
	BookFilterer   // Log filterer for contract events
}

// BookCaller is an auto generated read-only Go binding around an Ethereum contract.
type BookCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BookTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BookTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BookFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BookFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BookSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BookSession struct {
	Contract     *Book             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BookCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BookCallerSession struct {
	Contract *BookCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BookTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BookTransactorSession struct {
	Contract     *BookTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BookRaw is an auto generated low-level Go binding around an Ethereum contract.
type BookRaw struct {
	Contract *Book // Generic contract binding to access the raw methods on
}

// BookCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BookCallerRaw struct {
	Contract *BookCaller // Generic read-only contract binding to access the raw methods on
}

// BookTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BookTransactorRaw struct {
	Contract *BookTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBook creates a new instance of Book, bound to a specific deployed contract.
func NewBook(address common.Address, backend bind.ContractBackend) (*Book, error) {
	contract, err := bindBook(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Book{BookCaller: BookCaller{contract: contract}, BookTransactor: BookTransactor{contract: contract}, BookFilterer: BookFilterer{contract: contract}}, nil
}

// NewBookCaller creates a new read-only instance of Book, bound to a specific deployed contract.
func NewBookCaller(address common.Address, caller bind.ContractCaller) (*BookCaller, error) {
	contract, err := bindBook(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BookCaller{contract: contract}, nil
}

// NewBookTransactor creates a new write-only instance of Book, bound to a specific deployed contract.
func NewBookTransactor(address common.Address, transactor bind.ContractTransactor) (*BookTransactor, error) {
	contract, err := bindBook(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BookTransactor{contract: contract}, nil
}

// NewBookFilterer creates a new log filterer instance of Book, bound to a specific deployed contract.
func NewBookFilterer(address common.Address, filterer bind.ContractFilterer) (*BookFilterer, error) {
	contract, err := bindBook(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BookFilterer{contract: contract}, nil
}

// bindBook binds a generic wrapper to an already deployed contract.
func bindBook(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BookMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Book *BookRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Book.Contract.BookCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Book *BookRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Book.Contract.BookTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Book *BookRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Book.Contract.BookTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Book *BookCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Book.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Book *BookTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Book.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Book *BookTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Book.Contract.contract.Transact(opts, method, params...)
}

// Bid is a free data retrieval call binding the contract method 0x454a2ab3.
//
// Solidity: function bid(uint256 amount) view returns(uint256 price)
func (_Book *BookCaller) Bid(opts *bind.CallOpts, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Book.contract.Call(opts, &out, "bid", amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Bid is a free data retrieval call binding the contract method 0x454a2ab3.
//
// Solidity: function bid(uint256 amount) view returns(uint256 price)
func (_Book *BookSession) Bid(amount *big.Int) (*big.Int, error) {
	return _Book.Contract.Bid(&_Book.CallOpts, amount)
}

// Bid is a free data retrieval call binding the contract method 0x454a2ab3.
//
// Solidity: function bid(uint256 amount) view returns(uint256 price)
func (_Book *BookCallerSession) Bid(amount *big.Int) (*big.Int, error) {
	return _Book.Contract.Bid(&_Book.CallOpts, amount)
}

// Judge is a paid mutator transaction binding the contract method 0xcddb8df2.
//
// Solidity: function judge(uint256 amount) returns()
func (_Book *BookTransactor) Judge(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Book.contract.Transact(opts, "judge", amount)
}

// Judge is a paid mutator transaction binding the contract method 0xcddb8df2.
//
// Solidity: function judge(uint256 amount) returns()
func (_Book *BookSession) Judge(amount *big.Int) (*types.Transaction, error) {
	return _Book.Contract.Judge(&_Book.TransactOpts, amount)
}

// Judge is a paid mutator transaction binding the contract method 0xcddb8df2.
//
// Solidity: function judge(uint256 amount) returns()
func (_Book *BookTransactorSession) Judge(amount *big.Int) (*types.Transaction, error) {
	return _Book.Contract.Judge(&_Book.TransactOpts, amount)
}
