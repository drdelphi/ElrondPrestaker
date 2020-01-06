// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package prestake

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// TokenABI is the input ABI used to generate the binding from.
const TokenABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rewardsAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakeDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"WithdrawExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawInitiated\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"baseReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"baseRewardHistory\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"baseRewardHistoryLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"baseRewardsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"currentReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"initialDeposit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentStakingLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentStatus\",\"outputs\":[{\"internalType\":\"enumStakingContract.Status\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentTotalStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"executeWithdrawal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStakeDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startDate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endDate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startCheckpointIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endCheckpointIndex\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"initiateWithdrawal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"launchTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardConfig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"multiplier\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardsAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"multiplier\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"anualRewardRates\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"lowerBounds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"upperBounds\",\"type\":\"uint256[]\"}],\"name\":\"setupRewards\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"daysInterval\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakingPeriod\",\"type\":\"uint256\"}],\"name\":\"setupStakingLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"setupState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"staking\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"rewards\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakingLimitConfig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maxAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"daysInterval\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxIntervals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakingPeriod\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"toggleRewards\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Token is an auto generated Go binding around an Ethereum contract.
type Token struct {
	TokenCaller     // Read-only binding to the contract
	TokenTransactor // Write-only binding to the contract
	TokenFilterer   // Log filterer for contract events
}

// TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenSession struct {
	Contract     *Token            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenCallerSession struct {
	Contract *TokenCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenTransactorSession struct {
	Contract     *TokenTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenRaw struct {
	Contract *Token // Generic contract binding to access the raw methods on
}

// TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenCallerRaw struct {
	Contract *TokenCaller // Generic read-only contract binding to access the raw methods on
}

// TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenTransactorRaw struct {
	Contract *TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewToken creates a new instance of Token, bound to a specific deployed contract.
func NewToken(address common.Address, backend bind.ContractBackend) (*Token, error) {
	contract, err := bindToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}, TokenFilterer: TokenFilterer{contract: contract}}, nil
}

// NewTokenCaller creates a new read-only instance of Token, bound to a specific deployed contract.
func NewTokenCaller(address common.Address, caller bind.ContractCaller) (*TokenCaller, error) {
	contract, err := bindToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenCaller{contract: contract}, nil
}

// NewTokenTransactor creates a new write-only instance of Token, bound to a specific deployed contract.
func NewTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenTransactor, error) {
	contract, err := bindToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenTransactor{contract: contract}, nil
}

// NewTokenFilterer creates a new log filterer instance of Token, bound to a specific deployed contract.
func NewTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenFilterer, error) {
	contract, err := bindToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenFilterer{contract: contract}, nil
}

// bindToken binds a generic wrapper to an already deployed contract.
func bindToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Token.Contract.TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.contract.Transact(opts, method, params...)
}

// BaseReward is a free data retrieval call binding the contract method 0xc3977fcd.
//
// Solidity: function baseReward(uint256 index) constant returns(uint256, uint256, uint256)
func (_Token *TokenCaller) BaseReward(opts *bind.CallOpts, index *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _Token.contract.Call(opts, out, "baseReward", index)
	return *ret0, *ret1, *ret2, err
}

// BaseReward is a free data retrieval call binding the contract method 0xc3977fcd.
//
// Solidity: function baseReward(uint256 index) constant returns(uint256, uint256, uint256)
func (_Token *TokenSession) BaseReward(index *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _Token.Contract.BaseReward(&_Token.CallOpts, index)
}

// BaseReward is a free data retrieval call binding the contract method 0xc3977fcd.
//
// Solidity: function baseReward(uint256 index) constant returns(uint256, uint256, uint256)
func (_Token *TokenCallerSession) BaseReward(index *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _Token.Contract.BaseReward(&_Token.CallOpts, index)
}

// BaseRewardHistory is a free data retrieval call binding the contract method 0x3769209b.
//
// Solidity: function baseRewardHistory(uint256 index) constant returns(uint256, uint256, uint256, uint256)
func (_Token *TokenCaller) BaseRewardHistory(opts *bind.CallOpts, index *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
		ret3 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
	}
	err := _Token.contract.Call(opts, out, "baseRewardHistory", index)
	return *ret0, *ret1, *ret2, *ret3, err
}

// BaseRewardHistory is a free data retrieval call binding the contract method 0x3769209b.
//
// Solidity: function baseRewardHistory(uint256 index) constant returns(uint256, uint256, uint256, uint256)
func (_Token *TokenSession) BaseRewardHistory(index *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Token.Contract.BaseRewardHistory(&_Token.CallOpts, index)
}

// BaseRewardHistory is a free data retrieval call binding the contract method 0x3769209b.
//
// Solidity: function baseRewardHistory(uint256 index) constant returns(uint256, uint256, uint256, uint256)
func (_Token *TokenCallerSession) BaseRewardHistory(index *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Token.Contract.BaseRewardHistory(&_Token.CallOpts, index)
}

// BaseRewardHistoryLength is a free data retrieval call binding the contract method 0xbc5ca93c.
//
// Solidity: function baseRewardHistoryLength() constant returns(uint256)
func (_Token *TokenCaller) BaseRewardHistoryLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "baseRewardHistoryLength")
	return *ret0, err
}

// BaseRewardHistoryLength is a free data retrieval call binding the contract method 0xbc5ca93c.
//
// Solidity: function baseRewardHistoryLength() constant returns(uint256)
func (_Token *TokenSession) BaseRewardHistoryLength() (*big.Int, error) {
	return _Token.Contract.BaseRewardHistoryLength(&_Token.CallOpts)
}

// BaseRewardHistoryLength is a free data retrieval call binding the contract method 0xbc5ca93c.
//
// Solidity: function baseRewardHistoryLength() constant returns(uint256)
func (_Token *TokenCallerSession) BaseRewardHistoryLength() (*big.Int, error) {
	return _Token.Contract.BaseRewardHistoryLength(&_Token.CallOpts)
}

// BaseRewardsLength is a free data retrieval call binding the contract method 0xc3bc0c3e.
//
// Solidity: function baseRewardsLength() constant returns(uint256)
func (_Token *TokenCaller) BaseRewardsLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "baseRewardsLength")
	return *ret0, err
}

// BaseRewardsLength is a free data retrieval call binding the contract method 0xc3bc0c3e.
//
// Solidity: function baseRewardsLength() constant returns(uint256)
func (_Token *TokenSession) BaseRewardsLength() (*big.Int, error) {
	return _Token.Contract.BaseRewardsLength(&_Token.CallOpts)
}

// BaseRewardsLength is a free data retrieval call binding the contract method 0xc3bc0c3e.
//
// Solidity: function baseRewardsLength() constant returns(uint256)
func (_Token *TokenCallerSession) BaseRewardsLength() (*big.Int, error) {
	return _Token.Contract.BaseRewardsLength(&_Token.CallOpts)
}

// CurrentReward is a free data retrieval call binding the contract method 0x9d18e4b0.
//
// Solidity: function currentReward(address account) constant returns(uint256 initialDeposit, uint256 reward)
func (_Token *TokenCaller) CurrentReward(opts *bind.CallOpts, account common.Address) (struct {
	InitialDeposit *big.Int
	Reward         *big.Int
}, error) {
	ret := new(struct {
		InitialDeposit *big.Int
		Reward         *big.Int
	})
	out := ret
	err := _Token.contract.Call(opts, out, "currentReward", account)
	return *ret, err
}

// CurrentReward is a free data retrieval call binding the contract method 0x9d18e4b0.
//
// Solidity: function currentReward(address account) constant returns(uint256 initialDeposit, uint256 reward)
func (_Token *TokenSession) CurrentReward(account common.Address) (struct {
	InitialDeposit *big.Int
	Reward         *big.Int
}, error) {
	return _Token.Contract.CurrentReward(&_Token.CallOpts, account)
}

// CurrentReward is a free data retrieval call binding the contract method 0x9d18e4b0.
//
// Solidity: function currentReward(address account) constant returns(uint256 initialDeposit, uint256 reward)
func (_Token *TokenCallerSession) CurrentReward(account common.Address) (struct {
	InitialDeposit *big.Int
	Reward         *big.Int
}, error) {
	return _Token.Contract.CurrentReward(&_Token.CallOpts, account)
}

// CurrentStakingLimit is a free data retrieval call binding the contract method 0x3818548b.
//
// Solidity: function currentStakingLimit() constant returns(uint256)
func (_Token *TokenCaller) CurrentStakingLimit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "currentStakingLimit")
	return *ret0, err
}

// CurrentStakingLimit is a free data retrieval call binding the contract method 0x3818548b.
//
// Solidity: function currentStakingLimit() constant returns(uint256)
func (_Token *TokenSession) CurrentStakingLimit() (*big.Int, error) {
	return _Token.Contract.CurrentStakingLimit(&_Token.CallOpts)
}

// CurrentStakingLimit is a free data retrieval call binding the contract method 0x3818548b.
//
// Solidity: function currentStakingLimit() constant returns(uint256)
func (_Token *TokenCallerSession) CurrentStakingLimit() (*big.Int, error) {
	return _Token.Contract.CurrentStakingLimit(&_Token.CallOpts)
}

// CurrentStatus is a free data retrieval call binding the contract method 0xef8a9235.
//
// Solidity: function currentStatus() constant returns(uint8)
func (_Token *TokenCaller) CurrentStatus(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "currentStatus")
	return *ret0, err
}

// CurrentStatus is a free data retrieval call binding the contract method 0xef8a9235.
//
// Solidity: function currentStatus() constant returns(uint8)
func (_Token *TokenSession) CurrentStatus() (uint8, error) {
	return _Token.Contract.CurrentStatus(&_Token.CallOpts)
}

// CurrentStatus is a free data retrieval call binding the contract method 0xef8a9235.
//
// Solidity: function currentStatus() constant returns(uint8)
func (_Token *TokenCallerSession) CurrentStatus() (uint8, error) {
	return _Token.Contract.CurrentStatus(&_Token.CallOpts)
}

// CurrentTotalStake is a free data retrieval call binding the contract method 0xce4843e9.
//
// Solidity: function currentTotalStake() constant returns(uint256)
func (_Token *TokenCaller) CurrentTotalStake(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "currentTotalStake")
	return *ret0, err
}

// CurrentTotalStake is a free data retrieval call binding the contract method 0xce4843e9.
//
// Solidity: function currentTotalStake() constant returns(uint256)
func (_Token *TokenSession) CurrentTotalStake() (*big.Int, error) {
	return _Token.Contract.CurrentTotalStake(&_Token.CallOpts)
}

// CurrentTotalStake is a free data retrieval call binding the contract method 0xce4843e9.
//
// Solidity: function currentTotalStake() constant returns(uint256)
func (_Token *TokenCallerSession) CurrentTotalStake() (*big.Int, error) {
	return _Token.Contract.CurrentTotalStake(&_Token.CallOpts)
}

// GetStakeDeposit is a free data retrieval call binding the contract method 0xaef91d5f.
//
// Solidity: function getStakeDeposit() constant returns(uint256 amount, uint256 startDate, uint256 endDate, uint256 startCheckpointIndex, uint256 endCheckpointIndex)
func (_Token *TokenCaller) GetStakeDeposit(opts *bind.CallOpts) (struct {
	Amount               *big.Int
	StartDate            *big.Int
	EndDate              *big.Int
	StartCheckpointIndex *big.Int
	EndCheckpointIndex   *big.Int
}, error) {
	ret := new(struct {
		Amount               *big.Int
		StartDate            *big.Int
		EndDate              *big.Int
		StartCheckpointIndex *big.Int
		EndCheckpointIndex   *big.Int
	})
	out := ret
	err := _Token.contract.Call(opts, out, "getStakeDeposit")
	return *ret, err
}

// GetStakeDeposit is a free data retrieval call binding the contract method 0xaef91d5f.
//
// Solidity: function getStakeDeposit() constant returns(uint256 amount, uint256 startDate, uint256 endDate, uint256 startCheckpointIndex, uint256 endCheckpointIndex)
func (_Token *TokenSession) GetStakeDeposit() (struct {
	Amount               *big.Int
	StartDate            *big.Int
	EndDate              *big.Int
	StartCheckpointIndex *big.Int
	EndCheckpointIndex   *big.Int
}, error) {
	return _Token.Contract.GetStakeDeposit(&_Token.CallOpts)
}

// GetStakeDeposit is a free data retrieval call binding the contract method 0xaef91d5f.
//
// Solidity: function getStakeDeposit() constant returns(uint256 amount, uint256 startDate, uint256 endDate, uint256 startCheckpointIndex, uint256 endCheckpointIndex)
func (_Token *TokenCallerSession) GetStakeDeposit() (struct {
	Amount               *big.Int
	StartDate            *big.Int
	EndDate              *big.Int
	StartCheckpointIndex *big.Int
	EndCheckpointIndex   *big.Int
}, error) {
	return _Token.Contract.GetStakeDeposit(&_Token.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Token *TokenCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Token *TokenSession) IsOwner() (bool, error) {
	return _Token.Contract.IsOwner(&_Token.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Token *TokenCallerSession) IsOwner() (bool, error) {
	return _Token.Contract.IsOwner(&_Token.CallOpts)
}

// LaunchTimestamp is a free data retrieval call binding the contract method 0x65cf7c9b.
//
// Solidity: function launchTimestamp() constant returns(uint256)
func (_Token *TokenCaller) LaunchTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "launchTimestamp")
	return *ret0, err
}

// LaunchTimestamp is a free data retrieval call binding the contract method 0x65cf7c9b.
//
// Solidity: function launchTimestamp() constant returns(uint256)
func (_Token *TokenSession) LaunchTimestamp() (*big.Int, error) {
	return _Token.Contract.LaunchTimestamp(&_Token.CallOpts)
}

// LaunchTimestamp is a free data retrieval call binding the contract method 0x65cf7c9b.
//
// Solidity: function launchTimestamp() constant returns(uint256)
func (_Token *TokenCallerSession) LaunchTimestamp() (*big.Int, error) {
	return _Token.Contract.LaunchTimestamp(&_Token.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Token *TokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Token *TokenSession) Owner() (common.Address, error) {
	return _Token.Contract.Owner(&_Token.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Token *TokenCallerSession) Owner() (common.Address, error) {
	return _Token.Contract.Owner(&_Token.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Token *TokenCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Token *TokenSession) Paused() (bool, error) {
	return _Token.Contract.Paused(&_Token.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Token *TokenCallerSession) Paused() (bool, error) {
	return _Token.Contract.Paused(&_Token.CallOpts)
}

// RewardConfig is a free data retrieval call binding the contract method 0x4e94c285.
//
// Solidity: function rewardConfig() constant returns(uint256 multiplier)
func (_Token *TokenCaller) RewardConfig(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "rewardConfig")
	return *ret0, err
}

// RewardConfig is a free data retrieval call binding the contract method 0x4e94c285.
//
// Solidity: function rewardConfig() constant returns(uint256 multiplier)
func (_Token *TokenSession) RewardConfig() (*big.Int, error) {
	return _Token.Contract.RewardConfig(&_Token.CallOpts)
}

// RewardConfig is a free data retrieval call binding the contract method 0x4e94c285.
//
// Solidity: function rewardConfig() constant returns(uint256 multiplier)
func (_Token *TokenCallerSession) RewardConfig() (*big.Int, error) {
	return _Token.Contract.RewardConfig(&_Token.CallOpts)
}

// RewardsAddress is a free data retrieval call binding the contract method 0xc0973eed.
//
// Solidity: function rewardsAddress() constant returns(address)
func (_Token *TokenCaller) RewardsAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "rewardsAddress")
	return *ret0, err
}

// RewardsAddress is a free data retrieval call binding the contract method 0xc0973eed.
//
// Solidity: function rewardsAddress() constant returns(address)
func (_Token *TokenSession) RewardsAddress() (common.Address, error) {
	return _Token.Contract.RewardsAddress(&_Token.CallOpts)
}

// RewardsAddress is a free data retrieval call binding the contract method 0xc0973eed.
//
// Solidity: function rewardsAddress() constant returns(address)
func (_Token *TokenCallerSession) RewardsAddress() (common.Address, error) {
	return _Token.Contract.RewardsAddress(&_Token.CallOpts)
}

// SetupState is a free data retrieval call binding the contract method 0x0f31859a.
//
// Solidity: function setupState() constant returns(bool staking, bool rewards)
func (_Token *TokenCaller) SetupState(opts *bind.CallOpts) (struct {
	Staking bool
	Rewards bool
}, error) {
	ret := new(struct {
		Staking bool
		Rewards bool
	})
	out := ret
	err := _Token.contract.Call(opts, out, "setupState")
	return *ret, err
}

// SetupState is a free data retrieval call binding the contract method 0x0f31859a.
//
// Solidity: function setupState() constant returns(bool staking, bool rewards)
func (_Token *TokenSession) SetupState() (struct {
	Staking bool
	Rewards bool
}, error) {
	return _Token.Contract.SetupState(&_Token.CallOpts)
}

// SetupState is a free data retrieval call binding the contract method 0x0f31859a.
//
// Solidity: function setupState() constant returns(bool staking, bool rewards)
func (_Token *TokenCallerSession) SetupState() (struct {
	Staking bool
	Rewards bool
}, error) {
	return _Token.Contract.SetupState(&_Token.CallOpts)
}

// StakingLimitConfig is a free data retrieval call binding the contract method 0x73fe4a04.
//
// Solidity: function stakingLimitConfig() constant returns(uint256 maxAmount, uint256 initialAmount, uint256 daysInterval, uint256 maxIntervals, uint256 unstakingPeriod)
func (_Token *TokenCaller) StakingLimitConfig(opts *bind.CallOpts) (struct {
	MaxAmount       *big.Int
	InitialAmount   *big.Int
	DaysInterval    *big.Int
	MaxIntervals    *big.Int
	UnstakingPeriod *big.Int
}, error) {
	ret := new(struct {
		MaxAmount       *big.Int
		InitialAmount   *big.Int
		DaysInterval    *big.Int
		MaxIntervals    *big.Int
		UnstakingPeriod *big.Int
	})
	out := ret
	err := _Token.contract.Call(opts, out, "stakingLimitConfig")
	return *ret, err
}

// StakingLimitConfig is a free data retrieval call binding the contract method 0x73fe4a04.
//
// Solidity: function stakingLimitConfig() constant returns(uint256 maxAmount, uint256 initialAmount, uint256 daysInterval, uint256 maxIntervals, uint256 unstakingPeriod)
func (_Token *TokenSession) StakingLimitConfig() (struct {
	MaxAmount       *big.Int
	InitialAmount   *big.Int
	DaysInterval    *big.Int
	MaxIntervals    *big.Int
	UnstakingPeriod *big.Int
}, error) {
	return _Token.Contract.StakingLimitConfig(&_Token.CallOpts)
}

// StakingLimitConfig is a free data retrieval call binding the contract method 0x73fe4a04.
//
// Solidity: function stakingLimitConfig() constant returns(uint256 maxAmount, uint256 initialAmount, uint256 daysInterval, uint256 maxIntervals, uint256 unstakingPeriod)
func (_Token *TokenCallerSession) StakingLimitConfig() (struct {
	MaxAmount       *big.Int
	InitialAmount   *big.Int
	DaysInterval    *big.Int
	MaxIntervals    *big.Int
	UnstakingPeriod *big.Int
}, error) {
	return _Token.Contract.StakingLimitConfig(&_Token.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_Token *TokenCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_Token *TokenSession) Token() (common.Address, error) {
	return _Token.Contract.Token(&_Token.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_Token *TokenCallerSession) Token() (common.Address, error) {
	return _Token.Contract.Token(&_Token.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_Token *TokenTransactor) Deposit(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "deposit", amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_Token *TokenSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Deposit(&_Token.TransactOpts, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_Token *TokenTransactorSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Deposit(&_Token.TransactOpts, amount)
}

// ExecuteWithdrawal is a paid mutator transaction binding the contract method 0xfa45aa00.
//
// Solidity: function executeWithdrawal() returns()
func (_Token *TokenTransactor) ExecuteWithdrawal(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "executeWithdrawal")
}

// ExecuteWithdrawal is a paid mutator transaction binding the contract method 0xfa45aa00.
//
// Solidity: function executeWithdrawal() returns()
func (_Token *TokenSession) ExecuteWithdrawal() (*types.Transaction, error) {
	return _Token.Contract.ExecuteWithdrawal(&_Token.TransactOpts)
}

// ExecuteWithdrawal is a paid mutator transaction binding the contract method 0xfa45aa00.
//
// Solidity: function executeWithdrawal() returns()
func (_Token *TokenTransactorSession) ExecuteWithdrawal() (*types.Transaction, error) {
	return _Token.Contract.ExecuteWithdrawal(&_Token.TransactOpts)
}

// InitiateWithdrawal is a paid mutator transaction binding the contract method 0xb51d1d4f.
//
// Solidity: function initiateWithdrawal() returns()
func (_Token *TokenTransactor) InitiateWithdrawal(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "initiateWithdrawal")
}

// InitiateWithdrawal is a paid mutator transaction binding the contract method 0xb51d1d4f.
//
// Solidity: function initiateWithdrawal() returns()
func (_Token *TokenSession) InitiateWithdrawal() (*types.Transaction, error) {
	return _Token.Contract.InitiateWithdrawal(&_Token.TransactOpts)
}

// InitiateWithdrawal is a paid mutator transaction binding the contract method 0xb51d1d4f.
//
// Solidity: function initiateWithdrawal() returns()
func (_Token *TokenTransactorSession) InitiateWithdrawal() (*types.Transaction, error) {
	return _Token.Contract.InitiateWithdrawal(&_Token.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Token *TokenTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Token *TokenSession) Pause() (*types.Transaction, error) {
	return _Token.Contract.Pause(&_Token.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Token *TokenTransactorSession) Pause() (*types.Transaction, error) {
	return _Token.Contract.Pause(&_Token.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Token *TokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Token *TokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _Token.Contract.RenounceOwnership(&_Token.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Token *TokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Token.Contract.RenounceOwnership(&_Token.TransactOpts)
}

// SetupRewards is a paid mutator transaction binding the contract method 0x4716e2c4.
//
// Solidity: function setupRewards(uint256 multiplier, uint256[] anualRewardRates, uint256[] lowerBounds, uint256[] upperBounds) returns()
func (_Token *TokenTransactor) SetupRewards(opts *bind.TransactOpts, multiplier *big.Int, anualRewardRates []*big.Int, lowerBounds []*big.Int, upperBounds []*big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setupRewards", multiplier, anualRewardRates, lowerBounds, upperBounds)
}

// SetupRewards is a paid mutator transaction binding the contract method 0x4716e2c4.
//
// Solidity: function setupRewards(uint256 multiplier, uint256[] anualRewardRates, uint256[] lowerBounds, uint256[] upperBounds) returns()
func (_Token *TokenSession) SetupRewards(multiplier *big.Int, anualRewardRates []*big.Int, lowerBounds []*big.Int, upperBounds []*big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetupRewards(&_Token.TransactOpts, multiplier, anualRewardRates, lowerBounds, upperBounds)
}

// SetupRewards is a paid mutator transaction binding the contract method 0x4716e2c4.
//
// Solidity: function setupRewards(uint256 multiplier, uint256[] anualRewardRates, uint256[] lowerBounds, uint256[] upperBounds) returns()
func (_Token *TokenTransactorSession) SetupRewards(multiplier *big.Int, anualRewardRates []*big.Int, lowerBounds []*big.Int, upperBounds []*big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetupRewards(&_Token.TransactOpts, multiplier, anualRewardRates, lowerBounds, upperBounds)
}

// SetupStakingLimit is a paid mutator transaction binding the contract method 0x893bec43.
//
// Solidity: function setupStakingLimit(uint256 maxAmount, uint256 initialAmount, uint256 daysInterval, uint256 unstakingPeriod) returns()
func (_Token *TokenTransactor) SetupStakingLimit(opts *bind.TransactOpts, maxAmount *big.Int, initialAmount *big.Int, daysInterval *big.Int, unstakingPeriod *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setupStakingLimit", maxAmount, initialAmount, daysInterval, unstakingPeriod)
}

// SetupStakingLimit is a paid mutator transaction binding the contract method 0x893bec43.
//
// Solidity: function setupStakingLimit(uint256 maxAmount, uint256 initialAmount, uint256 daysInterval, uint256 unstakingPeriod) returns()
func (_Token *TokenSession) SetupStakingLimit(maxAmount *big.Int, initialAmount *big.Int, daysInterval *big.Int, unstakingPeriod *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetupStakingLimit(&_Token.TransactOpts, maxAmount, initialAmount, daysInterval, unstakingPeriod)
}

// SetupStakingLimit is a paid mutator transaction binding the contract method 0x893bec43.
//
// Solidity: function setupStakingLimit(uint256 maxAmount, uint256 initialAmount, uint256 daysInterval, uint256 unstakingPeriod) returns()
func (_Token *TokenTransactorSession) SetupStakingLimit(maxAmount *big.Int, initialAmount *big.Int, daysInterval *big.Int, unstakingPeriod *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetupStakingLimit(&_Token.TransactOpts, maxAmount, initialAmount, daysInterval, unstakingPeriod)
}

// ToggleRewards is a paid mutator transaction binding the contract method 0x8c19ec8b.
//
// Solidity: function toggleRewards(bool enabled) returns()
func (_Token *TokenTransactor) ToggleRewards(opts *bind.TransactOpts, enabled bool) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "toggleRewards", enabled)
}

// ToggleRewards is a paid mutator transaction binding the contract method 0x8c19ec8b.
//
// Solidity: function toggleRewards(bool enabled) returns()
func (_Token *TokenSession) ToggleRewards(enabled bool) (*types.Transaction, error) {
	return _Token.Contract.ToggleRewards(&_Token.TransactOpts, enabled)
}

// ToggleRewards is a paid mutator transaction binding the contract method 0x8c19ec8b.
//
// Solidity: function toggleRewards(bool enabled) returns()
func (_Token *TokenTransactorSession) ToggleRewards(enabled bool) (*types.Transaction, error) {
	return _Token.Contract.ToggleRewards(&_Token.TransactOpts, enabled)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Token *TokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Token *TokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Token.Contract.TransferOwnership(&_Token.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Token *TokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Token.Contract.TransferOwnership(&_Token.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Token *TokenTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Token *TokenSession) Unpause() (*types.Transaction, error) {
	return _Token.Contract.Unpause(&_Token.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Token *TokenTransactorSession) Unpause() (*types.Transaction, error) {
	return _Token.Contract.Unpause(&_Token.TransactOpts)
}

// TokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Token contract.
type TokenOwnershipTransferredIterator struct {
	Event *TokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenOwnershipTransferred)
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
		it.Event = new(TokenOwnershipTransferred)
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
func (it *TokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenOwnershipTransferred represents a OwnershipTransferred event raised by the Token contract.
type TokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Token *TokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TokenOwnershipTransferredIterator{contract: _Token.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Token *TokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenOwnershipTransferred)
				if err := _Token.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Token *TokenFilterer) ParseOwnershipTransferred(log types.Log) (*TokenOwnershipTransferred, error) {
	event := new(TokenOwnershipTransferred)
	if err := _Token.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Token contract.
type TokenPausedIterator struct {
	Event *TokenPaused // Event containing the contract specifics and raw log

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
func (it *TokenPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenPaused)
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
		it.Event = new(TokenPaused)
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
func (it *TokenPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenPaused represents a Paused event raised by the Token contract.
type TokenPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Token *TokenFilterer) FilterPaused(opts *bind.FilterOpts) (*TokenPausedIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &TokenPausedIterator{contract: _Token.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Token *TokenFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *TokenPaused) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenPaused)
				if err := _Token.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Token *TokenFilterer) ParsePaused(log types.Log) (*TokenPaused, error) {
	event := new(TokenPaused)
	if err := _Token.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenStakeDepositedIterator is returned from FilterStakeDeposited and is used to iterate over the raw logs and unpacked data for StakeDeposited events raised by the Token contract.
type TokenStakeDepositedIterator struct {
	Event *TokenStakeDeposited // Event containing the contract specifics and raw log

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
func (it *TokenStakeDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenStakeDeposited)
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
		it.Event = new(TokenStakeDeposited)
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
func (it *TokenStakeDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenStakeDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenStakeDeposited represents a StakeDeposited event raised by the Token contract.
type TokenStakeDeposited struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStakeDeposited is a free log retrieval operation binding the contract event 0x0a7bb2e28cc4698aac06db79cf9163bfcc20719286cf59fa7d492ceda1b8edc2.
//
// Solidity: event StakeDeposited(address indexed account, uint256 amount)
func (_Token *TokenFilterer) FilterStakeDeposited(opts *bind.FilterOpts, account []common.Address) (*TokenStakeDepositedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "StakeDeposited", accountRule)
	if err != nil {
		return nil, err
	}
	return &TokenStakeDepositedIterator{contract: _Token.contract, event: "StakeDeposited", logs: logs, sub: sub}, nil
}

// WatchStakeDeposited is a free log subscription operation binding the contract event 0x0a7bb2e28cc4698aac06db79cf9163bfcc20719286cf59fa7d492ceda1b8edc2.
//
// Solidity: event StakeDeposited(address indexed account, uint256 amount)
func (_Token *TokenFilterer) WatchStakeDeposited(opts *bind.WatchOpts, sink chan<- *TokenStakeDeposited, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "StakeDeposited", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenStakeDeposited)
				if err := _Token.contract.UnpackLog(event, "StakeDeposited", log); err != nil {
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

// ParseStakeDeposited is a log parse operation binding the contract event 0x0a7bb2e28cc4698aac06db79cf9163bfcc20719286cf59fa7d492ceda1b8edc2.
//
// Solidity: event StakeDeposited(address indexed account, uint256 amount)
func (_Token *TokenFilterer) ParseStakeDeposited(log types.Log) (*TokenStakeDeposited, error) {
	event := new(TokenStakeDeposited)
	if err := _Token.contract.UnpackLog(event, "StakeDeposited", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Token contract.
type TokenUnpausedIterator struct {
	Event *TokenUnpaused // Event containing the contract specifics and raw log

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
func (it *TokenUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenUnpaused)
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
		it.Event = new(TokenUnpaused)
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
func (it *TokenUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenUnpaused represents a Unpaused event raised by the Token contract.
type TokenUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Token *TokenFilterer) FilterUnpaused(opts *bind.FilterOpts) (*TokenUnpausedIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &TokenUnpausedIterator{contract: _Token.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Token *TokenFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *TokenUnpaused) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenUnpaused)
				if err := _Token.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Token *TokenFilterer) ParseUnpaused(log types.Log) (*TokenUnpaused, error) {
	event := new(TokenUnpaused)
	if err := _Token.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenWithdrawExecutedIterator is returned from FilterWithdrawExecuted and is used to iterate over the raw logs and unpacked data for WithdrawExecuted events raised by the Token contract.
type TokenWithdrawExecutedIterator struct {
	Event *TokenWithdrawExecuted // Event containing the contract specifics and raw log

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
func (it *TokenWithdrawExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenWithdrawExecuted)
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
		it.Event = new(TokenWithdrawExecuted)
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
func (it *TokenWithdrawExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenWithdrawExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenWithdrawExecuted represents a WithdrawExecuted event raised by the Token contract.
type TokenWithdrawExecuted struct {
	Account common.Address
	Amount  *big.Int
	Reward  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawExecuted is a free log retrieval operation binding the contract event 0xf6c0fc94edc8fd2a37514f1fb85af88fca386b00a54127b912273e6b83322e08.
//
// Solidity: event WithdrawExecuted(address indexed account, uint256 amount, uint256 reward)
func (_Token *TokenFilterer) FilterWithdrawExecuted(opts *bind.FilterOpts, account []common.Address) (*TokenWithdrawExecutedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "WithdrawExecuted", accountRule)
	if err != nil {
		return nil, err
	}
	return &TokenWithdrawExecutedIterator{contract: _Token.contract, event: "WithdrawExecuted", logs: logs, sub: sub}, nil
}

// WatchWithdrawExecuted is a free log subscription operation binding the contract event 0xf6c0fc94edc8fd2a37514f1fb85af88fca386b00a54127b912273e6b83322e08.
//
// Solidity: event WithdrawExecuted(address indexed account, uint256 amount, uint256 reward)
func (_Token *TokenFilterer) WatchWithdrawExecuted(opts *bind.WatchOpts, sink chan<- *TokenWithdrawExecuted, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "WithdrawExecuted", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenWithdrawExecuted)
				if err := _Token.contract.UnpackLog(event, "WithdrawExecuted", log); err != nil {
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

// ParseWithdrawExecuted is a log parse operation binding the contract event 0xf6c0fc94edc8fd2a37514f1fb85af88fca386b00a54127b912273e6b83322e08.
//
// Solidity: event WithdrawExecuted(address indexed account, uint256 amount, uint256 reward)
func (_Token *TokenFilterer) ParseWithdrawExecuted(log types.Log) (*TokenWithdrawExecuted, error) {
	event := new(TokenWithdrawExecuted)
	if err := _Token.contract.UnpackLog(event, "WithdrawExecuted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenWithdrawInitiatedIterator is returned from FilterWithdrawInitiated and is used to iterate over the raw logs and unpacked data for WithdrawInitiated events raised by the Token contract.
type TokenWithdrawInitiatedIterator struct {
	Event *TokenWithdrawInitiated // Event containing the contract specifics and raw log

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
func (it *TokenWithdrawInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenWithdrawInitiated)
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
		it.Event = new(TokenWithdrawInitiated)
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
func (it *TokenWithdrawInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenWithdrawInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenWithdrawInitiated represents a WithdrawInitiated event raised by the Token contract.
type TokenWithdrawInitiated struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawInitiated is a free log retrieval operation binding the contract event 0x15e1f18cd8cc91e09c1a1a056f1d8e61884aec50ad8502f88defaa64820e8fc2.
//
// Solidity: event WithdrawInitiated(address indexed account, uint256 amount)
func (_Token *TokenFilterer) FilterWithdrawInitiated(opts *bind.FilterOpts, account []common.Address) (*TokenWithdrawInitiatedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "WithdrawInitiated", accountRule)
	if err != nil {
		return nil, err
	}
	return &TokenWithdrawInitiatedIterator{contract: _Token.contract, event: "WithdrawInitiated", logs: logs, sub: sub}, nil
}

// WatchWithdrawInitiated is a free log subscription operation binding the contract event 0x15e1f18cd8cc91e09c1a1a056f1d8e61884aec50ad8502f88defaa64820e8fc2.
//
// Solidity: event WithdrawInitiated(address indexed account, uint256 amount)
func (_Token *TokenFilterer) WatchWithdrawInitiated(opts *bind.WatchOpts, sink chan<- *TokenWithdrawInitiated, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "WithdrawInitiated", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenWithdrawInitiated)
				if err := _Token.contract.UnpackLog(event, "WithdrawInitiated", log); err != nil {
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

// ParseWithdrawInitiated is a log parse operation binding the contract event 0x15e1f18cd8cc91e09c1a1a056f1d8e61884aec50ad8502f88defaa64820e8fc2.
//
// Solidity: event WithdrawInitiated(address indexed account, uint256 amount)
func (_Token *TokenFilterer) ParseWithdrawInitiated(log types.Log) (*TokenWithdrawInitiated, error) {
	event := new(TokenWithdrawInitiated)
	if err := _Token.contract.UnpackLog(event, "WithdrawInitiated", log); err != nil {
		return nil, err
	}
	return event, nil
}
