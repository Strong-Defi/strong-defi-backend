// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sch_stake

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

// SCHStakePool is an auto generated low-level Go binding around an user-defined struct.
type SCHStakePool struct {
	TokenAddress           common.Address
	Wight                  *big.Int
	TotalStakeAmount       *big.Int
	LastBlockNumber        *big.Int
	AccRCCPerST            *big.Int
	TotalReward            *big.Int
	MinStakeAmount         *big.Int
	UnStakeLockBlockNumber *big.Int
}

// SCHStakeRequest is an auto generated low-level Go binding around an user-defined struct.
type SCHStakeRequest struct {
	Amount              *big.Int
	RewardSCH           *big.Int
	IsRewardBlockNumber *big.Int
}

// SCHStakeUser is an auto generated low-level Go binding around an user-defined struct.
type SCHStakeUser struct {
	UserAddress common.Address
	FinishedRCC *big.Int
	PendingRCC  *big.Int
	StakeAmount *big.Int
	Requests    []SCHStakeRequest
}

// SCHStakeMetaData contains all meta data concerning the SCHStake contract.
var SCHStakeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rewardSCH\",\"type\":\"uint256\"}],\"name\":\"GetRewardFail\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"}],\"name\":\"AddPool\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardSCH\",\"type\":\"uint256\"}],\"name\":\"GetReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RequestStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"rewardLock\",\"type\":\"bool\"}],\"name\":\"RewardLock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"schToken\",\"type\":\"address\"}],\"name\":\"SCHToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Stake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"stakeLock\",\"type\":\"bool\"}],\"name\":\"StakeLock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalStakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accRCCPerST\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minStakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unStakeLockBlockNumber\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structSCHStake.Pool\",\"name\":\"pool\",\"type\":\"tuple\"}],\"name\":\"UpdatePool\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SCH\",\"outputs\":[{\"internalType\":\"contractSCToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_reward_Lock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_reward_Unlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_stake_Lock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_stake_Unlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wight\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minStakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isUpdatePool\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_unStakeLockBlockNumber\",\"type\":\"uint256\"}],\"name\":\"addPool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getContractBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"getPool\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalStakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accRCCPerST\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minStakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unStakeLockBlockNumber\",\"type\":\"uint256\"}],\"internalType\":\"structSCHStake.Pool\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRewardLock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSCTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStakeLock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUserInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"finishedRCC\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingRCC\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardSCH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"isRewardBlockNumber\",\"type\":\"uint256\"}],\"internalType\":\"structSCHStake.Request[]\",\"name\":\"requests\",\"type\":\"tuple[]\"}],\"internalType\":\"structSCHStake.User\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractSCToken\",\"name\":\"_SCH\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_rewardPerBlock\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"pauseStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// SCHStakeABI is the input ABI used to generate the binding from.
// Deprecated: Use SCHStakeMetaData.ABI instead.
var SCHStakeABI = SCHStakeMetaData.ABI

// SCHStake is an auto generated Go binding around an Ethereum contract.
type SCHStake struct {
	SCHStakeCaller     // Read-only binding to the contract
	SCHStakeTransactor // Write-only binding to the contract
	SCHStakeFilterer   // Log filterer for contract events
}

// SCHStakeCaller is an auto generated read-only Go binding around an Ethereum contract.
type SCHStakeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SCHStakeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SCHStakeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SCHStakeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SCHStakeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SCHStakeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SCHStakeSession struct {
	Contract     *SCHStake         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SCHStakeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SCHStakeCallerSession struct {
	Contract *SCHStakeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SCHStakeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SCHStakeTransactorSession struct {
	Contract     *SCHStakeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SCHStakeRaw is an auto generated low-level Go binding around an Ethereum contract.
type SCHStakeRaw struct {
	Contract *SCHStake // Generic contract binding to access the raw methods on
}

// SCHStakeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SCHStakeCallerRaw struct {
	Contract *SCHStakeCaller // Generic read-only contract binding to access the raw methods on
}

// SCHStakeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SCHStakeTransactorRaw struct {
	Contract *SCHStakeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSCHStake creates a new instance of SCHStake, bound to a specific deployed contract.
func NewSCHStake(address common.Address, backend bind.ContractBackend) (*SCHStake, error) {
	contract, err := bindSCHStake(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SCHStake{SCHStakeCaller: SCHStakeCaller{contract: contract}, SCHStakeTransactor: SCHStakeTransactor{contract: contract}, SCHStakeFilterer: SCHStakeFilterer{contract: contract}}, nil
}

// NewSCHStakeCaller creates a new read-only instance of SCHStake, bound to a specific deployed contract.
func NewSCHStakeCaller(address common.Address, caller bind.ContractCaller) (*SCHStakeCaller, error) {
	contract, err := bindSCHStake(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SCHStakeCaller{contract: contract}, nil
}

// NewSCHStakeTransactor creates a new write-only instance of SCHStake, bound to a specific deployed contract.
func NewSCHStakeTransactor(address common.Address, transactor bind.ContractTransactor) (*SCHStakeTransactor, error) {
	contract, err := bindSCHStake(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SCHStakeTransactor{contract: contract}, nil
}

// NewSCHStakeFilterer creates a new log filterer instance of SCHStake, bound to a specific deployed contract.
func NewSCHStakeFilterer(address common.Address, filterer bind.ContractFilterer) (*SCHStakeFilterer, error) {
	contract, err := bindSCHStake(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SCHStakeFilterer{contract: contract}, nil
}

// bindSCHStake binds a generic wrapper to an already deployed contract.
func bindSCHStake(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SCHStakeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SCHStake *SCHStakeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SCHStake.Contract.SCHStakeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SCHStake *SCHStakeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SCHStake.Contract.SCHStakeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SCHStake *SCHStakeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SCHStake.Contract.SCHStakeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SCHStake *SCHStakeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SCHStake.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SCHStake *SCHStakeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SCHStake.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SCHStake *SCHStakeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SCHStake.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_SCHStake *SCHStakeCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SCHStake.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_SCHStake *SCHStakeSession) ADMINROLE() ([32]byte, error) {
	return _SCHStake.Contract.ADMINROLE(&_SCHStake.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_SCHStake *SCHStakeCallerSession) ADMINROLE() ([32]byte, error) {
	return _SCHStake.Contract.ADMINROLE(&_SCHStake.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_SCHStake *SCHStakeCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SCHStake.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_SCHStake *SCHStakeSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _SCHStake.Contract.DEFAULTADMINROLE(&_SCHStake.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_SCHStake *SCHStakeCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _SCHStake.Contract.DEFAULTADMINROLE(&_SCHStake.CallOpts)
}

// SCH is a free data retrieval call binding the contract method 0xc5494b82.
//
// Solidity: function SCH() view returns(address)
func (_SCHStake *SCHStakeCaller) SCH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SCHStake.contract.Call(opts, &out, "SCH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SCH is a free data retrieval call binding the contract method 0xc5494b82.
//
// Solidity: function SCH() view returns(address)
func (_SCHStake *SCHStakeSession) SCH() (common.Address, error) {
	return _SCHStake.Contract.SCH(&_SCHStake.CallOpts)
}

// SCH is a free data retrieval call binding the contract method 0xc5494b82.
//
// Solidity: function SCH() view returns(address)
func (_SCHStake *SCHStakeCallerSession) SCH() (common.Address, error) {
	return _SCHStake.Contract.SCH(&_SCHStake.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_SCHStake *SCHStakeCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SCHStake.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_SCHStake *SCHStakeSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _SCHStake.Contract.UPGRADEINTERFACEVERSION(&_SCHStake.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_SCHStake *SCHStakeCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _SCHStake.Contract.UPGRADEINTERFACEVERSION(&_SCHStake.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address _user) view returns(uint256)
func (_SCHStake *SCHStakeCaller) GetBalance(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SCHStake.contract.Call(opts, &out, "getBalance", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address _user) view returns(uint256)
func (_SCHStake *SCHStakeSession) GetBalance(_user common.Address) (*big.Int, error) {
	return _SCHStake.Contract.GetBalance(&_SCHStake.CallOpts, _user)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address _user) view returns(uint256)
func (_SCHStake *SCHStakeCallerSession) GetBalance(_user common.Address) (*big.Int, error) {
	return _SCHStake.Contract.GetBalance(&_SCHStake.CallOpts, _user)
}

// GetContractBalance is a free data retrieval call binding the contract method 0x38ae12b2.
//
// Solidity: function getContractBalance(uint256 _pid, address addr) view returns(uint256)
func (_SCHStake *SCHStakeCaller) GetContractBalance(opts *bind.CallOpts, _pid *big.Int, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SCHStake.contract.Call(opts, &out, "getContractBalance", _pid, addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetContractBalance is a free data retrieval call binding the contract method 0x38ae12b2.
//
// Solidity: function getContractBalance(uint256 _pid, address addr) view returns(uint256)
func (_SCHStake *SCHStakeSession) GetContractBalance(_pid *big.Int, addr common.Address) (*big.Int, error) {
	return _SCHStake.Contract.GetContractBalance(&_SCHStake.CallOpts, _pid, addr)
}

// GetContractBalance is a free data retrieval call binding the contract method 0x38ae12b2.
//
// Solidity: function getContractBalance(uint256 _pid, address addr) view returns(uint256)
func (_SCHStake *SCHStakeCallerSession) GetContractBalance(_pid *big.Int, addr common.Address) (*big.Int, error) {
	return _SCHStake.Contract.GetContractBalance(&_SCHStake.CallOpts, _pid, addr)
}

// GetPool is a free data retrieval call binding the contract method 0x068bcd8d.
//
// Solidity: function getPool(uint256 _pid) view returns((address,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_SCHStake *SCHStakeCaller) GetPool(opts *bind.CallOpts, _pid *big.Int) (SCHStakePool, error) {
	var out []interface{}
	err := _SCHStake.contract.Call(opts, &out, "getPool", _pid)

	if err != nil {
		return *new(SCHStakePool), err
	}

	out0 := *abi.ConvertType(out[0], new(SCHStakePool)).(*SCHStakePool)

	return out0, err

}

// GetPool is a free data retrieval call binding the contract method 0x068bcd8d.
//
// Solidity: function getPool(uint256 _pid) view returns((address,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_SCHStake *SCHStakeSession) GetPool(_pid *big.Int) (SCHStakePool, error) {
	return _SCHStake.Contract.GetPool(&_SCHStake.CallOpts, _pid)
}

// GetPool is a free data retrieval call binding the contract method 0x068bcd8d.
//
// Solidity: function getPool(uint256 _pid) view returns((address,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_SCHStake *SCHStakeCallerSession) GetPool(_pid *big.Int) (SCHStakePool, error) {
	return _SCHStake.Contract.GetPool(&_SCHStake.CallOpts, _pid)
}

// GetPoolLength is a free data retrieval call binding the contract method 0xb3944d52.
//
// Solidity: function getPoolLength() view returns(uint256)
func (_SCHStake *SCHStakeCaller) GetPoolLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SCHStake.contract.Call(opts, &out, "getPoolLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPoolLength is a free data retrieval call binding the contract method 0xb3944d52.
//
// Solidity: function getPoolLength() view returns(uint256)
func (_SCHStake *SCHStakeSession) GetPoolLength() (*big.Int, error) {
	return _SCHStake.Contract.GetPoolLength(&_SCHStake.CallOpts)
}

// GetPoolLength is a free data retrieval call binding the contract method 0xb3944d52.
//
// Solidity: function getPoolLength() view returns(uint256)
func (_SCHStake *SCHStakeCallerSession) GetPoolLength() (*big.Int, error) {
	return _SCHStake.Contract.GetPoolLength(&_SCHStake.CallOpts)
}

// GetRewardLock is a free data retrieval call binding the contract method 0x6c1aecc8.
//
// Solidity: function getRewardLock() view returns(bool)
func (_SCHStake *SCHStakeCaller) GetRewardLock(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SCHStake.contract.Call(opts, &out, "getRewardLock")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetRewardLock is a free data retrieval call binding the contract method 0x6c1aecc8.
//
// Solidity: function getRewardLock() view returns(bool)
func (_SCHStake *SCHStakeSession) GetRewardLock() (bool, error) {
	return _SCHStake.Contract.GetRewardLock(&_SCHStake.CallOpts)
}

// GetRewardLock is a free data retrieval call binding the contract method 0x6c1aecc8.
//
// Solidity: function getRewardLock() view returns(bool)
func (_SCHStake *SCHStakeCallerSession) GetRewardLock() (bool, error) {
	return _SCHStake.Contract.GetRewardLock(&_SCHStake.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_SCHStake *SCHStakeCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _SCHStake.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_SCHStake *SCHStakeSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _SCHStake.Contract.GetRoleAdmin(&_SCHStake.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_SCHStake *SCHStakeCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _SCHStake.Contract.GetRoleAdmin(&_SCHStake.CallOpts, role)
}

// GetSCTokenAddress is a free data retrieval call binding the contract method 0xd01102ca.
//
// Solidity: function getSCTokenAddress() view returns(address)
func (_SCHStake *SCHStakeCaller) GetSCTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SCHStake.contract.Call(opts, &out, "getSCTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSCTokenAddress is a free data retrieval call binding the contract method 0xd01102ca.
//
// Solidity: function getSCTokenAddress() view returns(address)
func (_SCHStake *SCHStakeSession) GetSCTokenAddress() (common.Address, error) {
	return _SCHStake.Contract.GetSCTokenAddress(&_SCHStake.CallOpts)
}

// GetSCTokenAddress is a free data retrieval call binding the contract method 0xd01102ca.
//
// Solidity: function getSCTokenAddress() view returns(address)
func (_SCHStake *SCHStakeCallerSession) GetSCTokenAddress() (common.Address, error) {
	return _SCHStake.Contract.GetSCTokenAddress(&_SCHStake.CallOpts)
}

// GetStakeLock is a free data retrieval call binding the contract method 0x55cdc558.
//
// Solidity: function getStakeLock() view returns(bool)
func (_SCHStake *SCHStakeCaller) GetStakeLock(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SCHStake.contract.Call(opts, &out, "getStakeLock")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetStakeLock is a free data retrieval call binding the contract method 0x55cdc558.
//
// Solidity: function getStakeLock() view returns(bool)
func (_SCHStake *SCHStakeSession) GetStakeLock() (bool, error) {
	return _SCHStake.Contract.GetStakeLock(&_SCHStake.CallOpts)
}

// GetStakeLock is a free data retrieval call binding the contract method 0x55cdc558.
//
// Solidity: function getStakeLock() view returns(bool)
func (_SCHStake *SCHStakeCallerSession) GetStakeLock() (bool, error) {
	return _SCHStake.Contract.GetStakeLock(&_SCHStake.CallOpts)
}

// GetUserInfo is a free data retrieval call binding the contract method 0x1069f3b5.
//
// Solidity: function getUserInfo(uint256 _pid, address _user) view returns((address,uint256,uint256,uint256,(uint256,uint256,uint256)[]))
func (_SCHStake *SCHStakeCaller) GetUserInfo(opts *bind.CallOpts, _pid *big.Int, _user common.Address) (SCHStakeUser, error) {
	var out []interface{}
	err := _SCHStake.contract.Call(opts, &out, "getUserInfo", _pid, _user)

	if err != nil {
		return *new(SCHStakeUser), err
	}

	out0 := *abi.ConvertType(out[0], new(SCHStakeUser)).(*SCHStakeUser)

	return out0, err

}

// GetUserInfo is a free data retrieval call binding the contract method 0x1069f3b5.
//
// Solidity: function getUserInfo(uint256 _pid, address _user) view returns((address,uint256,uint256,uint256,(uint256,uint256,uint256)[]))
func (_SCHStake *SCHStakeSession) GetUserInfo(_pid *big.Int, _user common.Address) (SCHStakeUser, error) {
	return _SCHStake.Contract.GetUserInfo(&_SCHStake.CallOpts, _pid, _user)
}

// GetUserInfo is a free data retrieval call binding the contract method 0x1069f3b5.
//
// Solidity: function getUserInfo(uint256 _pid, address _user) view returns((address,uint256,uint256,uint256,(uint256,uint256,uint256)[]))
func (_SCHStake *SCHStakeCallerSession) GetUserInfo(_pid *big.Int, _user common.Address) (SCHStakeUser, error) {
	return _SCHStake.Contract.GetUserInfo(&_SCHStake.CallOpts, _pid, _user)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_SCHStake *SCHStakeCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _SCHStake.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_SCHStake *SCHStakeSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _SCHStake.Contract.HasRole(&_SCHStake.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_SCHStake *SCHStakeCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _SCHStake.Contract.HasRole(&_SCHStake.CallOpts, role, account)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_SCHStake *SCHStakeCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SCHStake.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_SCHStake *SCHStakeSession) ProxiableUUID() ([32]byte, error) {
	return _SCHStake.Contract.ProxiableUUID(&_SCHStake.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_SCHStake *SCHStakeCallerSession) ProxiableUUID() ([32]byte, error) {
	return _SCHStake.Contract.ProxiableUUID(&_SCHStake.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SCHStake *SCHStakeCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _SCHStake.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SCHStake *SCHStakeSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SCHStake.Contract.SupportsInterface(&_SCHStake.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SCHStake *SCHStakeCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SCHStake.Contract.SupportsInterface(&_SCHStake.CallOpts, interfaceId)
}

// RewardLock is a paid mutator transaction binding the contract method 0xfa8dacf2.
//
// Solidity: function _reward_Lock() returns()
func (_SCHStake *SCHStakeTransactor) RewardLock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SCHStake.contract.Transact(opts, "_reward_Lock")
}

// RewardLock is a paid mutator transaction binding the contract method 0xfa8dacf2.
//
// Solidity: function _reward_Lock() returns()
func (_SCHStake *SCHStakeSession) RewardLock() (*types.Transaction, error) {
	return _SCHStake.Contract.RewardLock(&_SCHStake.TransactOpts)
}

// RewardLock is a paid mutator transaction binding the contract method 0xfa8dacf2.
//
// Solidity: function _reward_Lock() returns()
func (_SCHStake *SCHStakeTransactorSession) RewardLock() (*types.Transaction, error) {
	return _SCHStake.Contract.RewardLock(&_SCHStake.TransactOpts)
}

// RewardUnlock is a paid mutator transaction binding the contract method 0x4ac9a62c.
//
// Solidity: function _reward_Unlock() returns()
func (_SCHStake *SCHStakeTransactor) RewardUnlock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SCHStake.contract.Transact(opts, "_reward_Unlock")
}

// RewardUnlock is a paid mutator transaction binding the contract method 0x4ac9a62c.
//
// Solidity: function _reward_Unlock() returns()
func (_SCHStake *SCHStakeSession) RewardUnlock() (*types.Transaction, error) {
	return _SCHStake.Contract.RewardUnlock(&_SCHStake.TransactOpts)
}

// RewardUnlock is a paid mutator transaction binding the contract method 0x4ac9a62c.
//
// Solidity: function _reward_Unlock() returns()
func (_SCHStake *SCHStakeTransactorSession) RewardUnlock() (*types.Transaction, error) {
	return _SCHStake.Contract.RewardUnlock(&_SCHStake.TransactOpts)
}

// StakeLock is a paid mutator transaction binding the contract method 0x478f1d12.
//
// Solidity: function _stake_Lock() returns()
func (_SCHStake *SCHStakeTransactor) StakeLock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SCHStake.contract.Transact(opts, "_stake_Lock")
}

// StakeLock is a paid mutator transaction binding the contract method 0x478f1d12.
//
// Solidity: function _stake_Lock() returns()
func (_SCHStake *SCHStakeSession) StakeLock() (*types.Transaction, error) {
	return _SCHStake.Contract.StakeLock(&_SCHStake.TransactOpts)
}

// StakeLock is a paid mutator transaction binding the contract method 0x478f1d12.
//
// Solidity: function _stake_Lock() returns()
func (_SCHStake *SCHStakeTransactorSession) StakeLock() (*types.Transaction, error) {
	return _SCHStake.Contract.StakeLock(&_SCHStake.TransactOpts)
}

// StakeUnlock is a paid mutator transaction binding the contract method 0x287d1ec3.
//
// Solidity: function _stake_Unlock() returns()
func (_SCHStake *SCHStakeTransactor) StakeUnlock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SCHStake.contract.Transact(opts, "_stake_Unlock")
}

// StakeUnlock is a paid mutator transaction binding the contract method 0x287d1ec3.
//
// Solidity: function _stake_Unlock() returns()
func (_SCHStake *SCHStakeSession) StakeUnlock() (*types.Transaction, error) {
	return _SCHStake.Contract.StakeUnlock(&_SCHStake.TransactOpts)
}

// StakeUnlock is a paid mutator transaction binding the contract method 0x287d1ec3.
//
// Solidity: function _stake_Unlock() returns()
func (_SCHStake *SCHStakeTransactorSession) StakeUnlock() (*types.Transaction, error) {
	return _SCHStake.Contract.StakeUnlock(&_SCHStake.TransactOpts)
}

// AddPool is a paid mutator transaction binding the contract method 0x6418d260.
//
// Solidity: function addPool(uint256 wight, address _tokenAddress, uint256 _minStakeAmount, bool isUpdatePool, uint256 _unStakeLockBlockNumber) returns(uint256)
func (_SCHStake *SCHStakeTransactor) AddPool(opts *bind.TransactOpts, wight *big.Int, _tokenAddress common.Address, _minStakeAmount *big.Int, isUpdatePool bool, _unStakeLockBlockNumber *big.Int) (*types.Transaction, error) {
	return _SCHStake.contract.Transact(opts, "addPool", wight, _tokenAddress, _minStakeAmount, isUpdatePool, _unStakeLockBlockNumber)
}

// AddPool is a paid mutator transaction binding the contract method 0x6418d260.
//
// Solidity: function addPool(uint256 wight, address _tokenAddress, uint256 _minStakeAmount, bool isUpdatePool, uint256 _unStakeLockBlockNumber) returns(uint256)
func (_SCHStake *SCHStakeSession) AddPool(wight *big.Int, _tokenAddress common.Address, _minStakeAmount *big.Int, isUpdatePool bool, _unStakeLockBlockNumber *big.Int) (*types.Transaction, error) {
	return _SCHStake.Contract.AddPool(&_SCHStake.TransactOpts, wight, _tokenAddress, _minStakeAmount, isUpdatePool, _unStakeLockBlockNumber)
}

// AddPool is a paid mutator transaction binding the contract method 0x6418d260.
//
// Solidity: function addPool(uint256 wight, address _tokenAddress, uint256 _minStakeAmount, bool isUpdatePool, uint256 _unStakeLockBlockNumber) returns(uint256)
func (_SCHStake *SCHStakeTransactorSession) AddPool(wight *big.Int, _tokenAddress common.Address, _minStakeAmount *big.Int, isUpdatePool bool, _unStakeLockBlockNumber *big.Int) (*types.Transaction, error) {
	return _SCHStake.Contract.AddPool(&_SCHStake.TransactOpts, wight, _tokenAddress, _minStakeAmount, isUpdatePool, _unStakeLockBlockNumber)
}

// GetReward is a paid mutator transaction binding the contract method 0x008f33d7.
//
// Solidity: function getReward(uint256 _pid, address _user) returns()
func (_SCHStake *SCHStakeTransactor) GetReward(opts *bind.TransactOpts, _pid *big.Int, _user common.Address) (*types.Transaction, error) {
	return _SCHStake.contract.Transact(opts, "getReward", _pid, _user)
}

// GetReward is a paid mutator transaction binding the contract method 0x008f33d7.
//
// Solidity: function getReward(uint256 _pid, address _user) returns()
func (_SCHStake *SCHStakeSession) GetReward(_pid *big.Int, _user common.Address) (*types.Transaction, error) {
	return _SCHStake.Contract.GetReward(&_SCHStake.TransactOpts, _pid, _user)
}

// GetReward is a paid mutator transaction binding the contract method 0x008f33d7.
//
// Solidity: function getReward(uint256 _pid, address _user) returns()
func (_SCHStake *SCHStakeTransactorSession) GetReward(_pid *big.Int, _user common.Address) (*types.Transaction, error) {
	return _SCHStake.Contract.GetReward(&_SCHStake.TransactOpts, _pid, _user)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_SCHStake *SCHStakeTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SCHStake.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_SCHStake *SCHStakeSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SCHStake.Contract.GrantRole(&_SCHStake.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_SCHStake *SCHStakeTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SCHStake.Contract.GrantRole(&_SCHStake.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address _SCH, uint256 _rewardPerBlock) returns()
func (_SCHStake *SCHStakeTransactor) Initialize(opts *bind.TransactOpts, _SCH common.Address, _rewardPerBlock *big.Int) (*types.Transaction, error) {
	return _SCHStake.contract.Transact(opts, "initialize", _SCH, _rewardPerBlock)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address _SCH, uint256 _rewardPerBlock) returns()
func (_SCHStake *SCHStakeSession) Initialize(_SCH common.Address, _rewardPerBlock *big.Int) (*types.Transaction, error) {
	return _SCHStake.Contract.Initialize(&_SCHStake.TransactOpts, _SCH, _rewardPerBlock)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address _SCH, uint256 _rewardPerBlock) returns()
func (_SCHStake *SCHStakeTransactorSession) Initialize(_SCH common.Address, _rewardPerBlock *big.Int) (*types.Transaction, error) {
	return _SCHStake.Contract.Initialize(&_SCHStake.TransactOpts, _SCH, _rewardPerBlock)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_SCHStake *SCHStakeTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SCHStake.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_SCHStake *SCHStakeSession) Pause() (*types.Transaction, error) {
	return _SCHStake.Contract.Pause(&_SCHStake.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_SCHStake *SCHStakeTransactorSession) Pause() (*types.Transaction, error) {
	return _SCHStake.Contract.Pause(&_SCHStake.TransactOpts)
}

// PauseStake is a paid mutator transaction binding the contract method 0xa79473fe.
//
// Solidity: function pauseStake(uint256 _pid, address _user, uint256 _amount) returns()
func (_SCHStake *SCHStakeTransactor) PauseStake(opts *bind.TransactOpts, _pid *big.Int, _user common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SCHStake.contract.Transact(opts, "pauseStake", _pid, _user, _amount)
}

// PauseStake is a paid mutator transaction binding the contract method 0xa79473fe.
//
// Solidity: function pauseStake(uint256 _pid, address _user, uint256 _amount) returns()
func (_SCHStake *SCHStakeSession) PauseStake(_pid *big.Int, _user common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SCHStake.Contract.PauseStake(&_SCHStake.TransactOpts, _pid, _user, _amount)
}

// PauseStake is a paid mutator transaction binding the contract method 0xa79473fe.
//
// Solidity: function pauseStake(uint256 _pid, address _user, uint256 _amount) returns()
func (_SCHStake *SCHStakeTransactorSession) PauseStake(_pid *big.Int, _user common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SCHStake.Contract.PauseStake(&_SCHStake.TransactOpts, _pid, _user, _amount)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_SCHStake *SCHStakeTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _SCHStake.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_SCHStake *SCHStakeSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _SCHStake.Contract.RenounceRole(&_SCHStake.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_SCHStake *SCHStakeTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _SCHStake.Contract.RenounceRole(&_SCHStake.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_SCHStake *SCHStakeTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SCHStake.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_SCHStake *SCHStakeSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SCHStake.Contract.RevokeRole(&_SCHStake.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_SCHStake *SCHStakeTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SCHStake.Contract.RevokeRole(&_SCHStake.TransactOpts, role, account)
}

// Stake is a paid mutator transaction binding the contract method 0x6e9c931c.
//
// Solidity: function stake(uint256 _pid, address _user, uint256 _amount) returns()
func (_SCHStake *SCHStakeTransactor) Stake(opts *bind.TransactOpts, _pid *big.Int, _user common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SCHStake.contract.Transact(opts, "stake", _pid, _user, _amount)
}

// Stake is a paid mutator transaction binding the contract method 0x6e9c931c.
//
// Solidity: function stake(uint256 _pid, address _user, uint256 _amount) returns()
func (_SCHStake *SCHStakeSession) Stake(_pid *big.Int, _user common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SCHStake.Contract.Stake(&_SCHStake.TransactOpts, _pid, _user, _amount)
}

// Stake is a paid mutator transaction binding the contract method 0x6e9c931c.
//
// Solidity: function stake(uint256 _pid, address _user, uint256 _amount) returns()
func (_SCHStake *SCHStakeTransactorSession) Stake(_pid *big.Int, _user common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SCHStake.Contract.Stake(&_SCHStake.TransactOpts, _pid, _user, _amount)
}

// UnPause is a paid mutator transaction binding the contract method 0xf7b188a5.
//
// Solidity: function unPause() returns()
func (_SCHStake *SCHStakeTransactor) UnPause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SCHStake.contract.Transact(opts, "unPause")
}

// UnPause is a paid mutator transaction binding the contract method 0xf7b188a5.
//
// Solidity: function unPause() returns()
func (_SCHStake *SCHStakeSession) UnPause() (*types.Transaction, error) {
	return _SCHStake.Contract.UnPause(&_SCHStake.TransactOpts)
}

// UnPause is a paid mutator transaction binding the contract method 0xf7b188a5.
//
// Solidity: function unPause() returns()
func (_SCHStake *SCHStakeTransactorSession) UnPause() (*types.Transaction, error) {
	return _SCHStake.Contract.UnPause(&_SCHStake.TransactOpts)
}

// UpdateReward is a paid mutator transaction binding the contract method 0xf36c0a72.
//
// Solidity: function updateReward() returns()
func (_SCHStake *SCHStakeTransactor) UpdateReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SCHStake.contract.Transact(opts, "updateReward")
}

// UpdateReward is a paid mutator transaction binding the contract method 0xf36c0a72.
//
// Solidity: function updateReward() returns()
func (_SCHStake *SCHStakeSession) UpdateReward() (*types.Transaction, error) {
	return _SCHStake.Contract.UpdateReward(&_SCHStake.TransactOpts)
}

// UpdateReward is a paid mutator transaction binding the contract method 0xf36c0a72.
//
// Solidity: function updateReward() returns()
func (_SCHStake *SCHStakeTransactorSession) UpdateReward() (*types.Transaction, error) {
	return _SCHStake.Contract.UpdateReward(&_SCHStake.TransactOpts)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_SCHStake *SCHStakeTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _SCHStake.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_SCHStake *SCHStakeSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _SCHStake.Contract.UpgradeToAndCall(&_SCHStake.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_SCHStake *SCHStakeTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _SCHStake.Contract.UpgradeToAndCall(&_SCHStake.TransactOpts, newImplementation, data)
}

// SCHStakeAddPoolIterator is returned from FilterAddPool and is used to iterate over the raw logs and unpacked data for AddPool events raised by the SCHStake contract.
type SCHStakeAddPoolIterator struct {
	Event *SCHStakeAddPool // Event containing the contract specifics and raw log

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
func (it *SCHStakeAddPoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SCHStakeAddPool)
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
		it.Event = new(SCHStakeAddPool)
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
func (it *SCHStakeAddPoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SCHStakeAddPoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SCHStakeAddPool represents a AddPool event raised by the SCHStake contract.
type SCHStakeAddPool struct {
	PoolId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAddPool is a free log retrieval operation binding the contract event 0x6f92874181ba07c8e988c53b7d3c2fdcff7154a3500137bf6a350ebac65c0870.
//
// Solidity: event AddPool(uint256 poolId)
func (_SCHStake *SCHStakeFilterer) FilterAddPool(opts *bind.FilterOpts) (*SCHStakeAddPoolIterator, error) {

	logs, sub, err := _SCHStake.contract.FilterLogs(opts, "AddPool")
	if err != nil {
		return nil, err
	}
	return &SCHStakeAddPoolIterator{contract: _SCHStake.contract, event: "AddPool", logs: logs, sub: sub}, nil
}

// WatchAddPool is a free log subscription operation binding the contract event 0x6f92874181ba07c8e988c53b7d3c2fdcff7154a3500137bf6a350ebac65c0870.
//
// Solidity: event AddPool(uint256 poolId)
func (_SCHStake *SCHStakeFilterer) WatchAddPool(opts *bind.WatchOpts, sink chan<- *SCHStakeAddPool) (event.Subscription, error) {

	logs, sub, err := _SCHStake.contract.WatchLogs(opts, "AddPool")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SCHStakeAddPool)
				if err := _SCHStake.contract.UnpackLog(event, "AddPool", log); err != nil {
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

// ParseAddPool is a log parse operation binding the contract event 0x6f92874181ba07c8e988c53b7d3c2fdcff7154a3500137bf6a350ebac65c0870.
//
// Solidity: event AddPool(uint256 poolId)
func (_SCHStake *SCHStakeFilterer) ParseAddPool(log types.Log) (*SCHStakeAddPool, error) {
	event := new(SCHStakeAddPool)
	if err := _SCHStake.contract.UnpackLog(event, "AddPool", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SCHStakeGetRewardIterator is returned from FilterGetReward and is used to iterate over the raw logs and unpacked data for GetReward events raised by the SCHStake contract.
type SCHStakeGetRewardIterator struct {
	Event *SCHStakeGetReward // Event containing the contract specifics and raw log

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
func (it *SCHStakeGetRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SCHStakeGetReward)
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
		it.Event = new(SCHStakeGetReward)
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
func (it *SCHStakeGetRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SCHStakeGetRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SCHStakeGetReward represents a GetReward event raised by the SCHStake contract.
type SCHStakeGetReward struct {
	User      common.Address
	Pid       *big.Int
	RewardSCH *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterGetReward is a free log retrieval operation binding the contract event 0x971ef1dba2d1e17950e4d8a3028c07cf038d0f5bf1ea9394125e258dd867d488.
//
// Solidity: event GetReward(address indexed user, uint256 indexed pid, uint256 rewardSCH)
func (_SCHStake *SCHStakeFilterer) FilterGetReward(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*SCHStakeGetRewardIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _SCHStake.contract.FilterLogs(opts, "GetReward", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &SCHStakeGetRewardIterator{contract: _SCHStake.contract, event: "GetReward", logs: logs, sub: sub}, nil
}

// WatchGetReward is a free log subscription operation binding the contract event 0x971ef1dba2d1e17950e4d8a3028c07cf038d0f5bf1ea9394125e258dd867d488.
//
// Solidity: event GetReward(address indexed user, uint256 indexed pid, uint256 rewardSCH)
func (_SCHStake *SCHStakeFilterer) WatchGetReward(opts *bind.WatchOpts, sink chan<- *SCHStakeGetReward, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _SCHStake.contract.WatchLogs(opts, "GetReward", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SCHStakeGetReward)
				if err := _SCHStake.contract.UnpackLog(event, "GetReward", log); err != nil {
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

// ParseGetReward is a log parse operation binding the contract event 0x971ef1dba2d1e17950e4d8a3028c07cf038d0f5bf1ea9394125e258dd867d488.
//
// Solidity: event GetReward(address indexed user, uint256 indexed pid, uint256 rewardSCH)
func (_SCHStake *SCHStakeFilterer) ParseGetReward(log types.Log) (*SCHStakeGetReward, error) {
	event := new(SCHStakeGetReward)
	if err := _SCHStake.contract.UnpackLog(event, "GetReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SCHStakeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the SCHStake contract.
type SCHStakeInitializedIterator struct {
	Event *SCHStakeInitialized // Event containing the contract specifics and raw log

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
func (it *SCHStakeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SCHStakeInitialized)
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
		it.Event = new(SCHStakeInitialized)
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
func (it *SCHStakeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SCHStakeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SCHStakeInitialized represents a Initialized event raised by the SCHStake contract.
type SCHStakeInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_SCHStake *SCHStakeFilterer) FilterInitialized(opts *bind.FilterOpts) (*SCHStakeInitializedIterator, error) {

	logs, sub, err := _SCHStake.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SCHStakeInitializedIterator{contract: _SCHStake.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_SCHStake *SCHStakeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SCHStakeInitialized) (event.Subscription, error) {

	logs, sub, err := _SCHStake.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SCHStakeInitialized)
				if err := _SCHStake.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_SCHStake *SCHStakeFilterer) ParseInitialized(log types.Log) (*SCHStakeInitialized, error) {
	event := new(SCHStakeInitialized)
	if err := _SCHStake.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SCHStakeRequestStakeIterator is returned from FilterRequestStake and is used to iterate over the raw logs and unpacked data for RequestStake events raised by the SCHStake contract.
type SCHStakeRequestStakeIterator struct {
	Event *SCHStakeRequestStake // Event containing the contract specifics and raw log

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
func (it *SCHStakeRequestStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SCHStakeRequestStake)
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
		it.Event = new(SCHStakeRequestStake)
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
func (it *SCHStakeRequestStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SCHStakeRequestStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SCHStakeRequestStake represents a RequestStake event raised by the SCHStake contract.
type SCHStakeRequestStake struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRequestStake is a free log retrieval operation binding the contract event 0x77cf89c46ebd91051f6234ef74f00116e50723d51433996767eba74df5cde9f0.
//
// Solidity: event RequestStake(address indexed user, uint256 indexed pid, uint256 amount)
func (_SCHStake *SCHStakeFilterer) FilterRequestStake(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*SCHStakeRequestStakeIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _SCHStake.contract.FilterLogs(opts, "RequestStake", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &SCHStakeRequestStakeIterator{contract: _SCHStake.contract, event: "RequestStake", logs: logs, sub: sub}, nil
}

// WatchRequestStake is a free log subscription operation binding the contract event 0x77cf89c46ebd91051f6234ef74f00116e50723d51433996767eba74df5cde9f0.
//
// Solidity: event RequestStake(address indexed user, uint256 indexed pid, uint256 amount)
func (_SCHStake *SCHStakeFilterer) WatchRequestStake(opts *bind.WatchOpts, sink chan<- *SCHStakeRequestStake, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _SCHStake.contract.WatchLogs(opts, "RequestStake", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SCHStakeRequestStake)
				if err := _SCHStake.contract.UnpackLog(event, "RequestStake", log); err != nil {
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

// ParseRequestStake is a log parse operation binding the contract event 0x77cf89c46ebd91051f6234ef74f00116e50723d51433996767eba74df5cde9f0.
//
// Solidity: event RequestStake(address indexed user, uint256 indexed pid, uint256 amount)
func (_SCHStake *SCHStakeFilterer) ParseRequestStake(log types.Log) (*SCHStakeRequestStake, error) {
	event := new(SCHStakeRequestStake)
	if err := _SCHStake.contract.UnpackLog(event, "RequestStake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SCHStakeRewardLockIterator is returned from FilterRewardLock and is used to iterate over the raw logs and unpacked data for RewardLock events raised by the SCHStake contract.
type SCHStakeRewardLockIterator struct {
	Event *SCHStakeRewardLock // Event containing the contract specifics and raw log

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
func (it *SCHStakeRewardLockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SCHStakeRewardLock)
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
		it.Event = new(SCHStakeRewardLock)
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
func (it *SCHStakeRewardLockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SCHStakeRewardLockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SCHStakeRewardLock represents a RewardLock event raised by the SCHStake contract.
type SCHStakeRewardLock struct {
	RewardLock bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRewardLock is a free log retrieval operation binding the contract event 0xc43b85b876df0e1fcd072756642cda721ec62416510229af71d47e3a3513b730.
//
// Solidity: event RewardLock(bool rewardLock)
func (_SCHStake *SCHStakeFilterer) FilterRewardLock(opts *bind.FilterOpts) (*SCHStakeRewardLockIterator, error) {

	logs, sub, err := _SCHStake.contract.FilterLogs(opts, "RewardLock")
	if err != nil {
		return nil, err
	}
	return &SCHStakeRewardLockIterator{contract: _SCHStake.contract, event: "RewardLock", logs: logs, sub: sub}, nil
}

// WatchRewardLock is a free log subscription operation binding the contract event 0xc43b85b876df0e1fcd072756642cda721ec62416510229af71d47e3a3513b730.
//
// Solidity: event RewardLock(bool rewardLock)
func (_SCHStake *SCHStakeFilterer) WatchRewardLock(opts *bind.WatchOpts, sink chan<- *SCHStakeRewardLock) (event.Subscription, error) {

	logs, sub, err := _SCHStake.contract.WatchLogs(opts, "RewardLock")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SCHStakeRewardLock)
				if err := _SCHStake.contract.UnpackLog(event, "RewardLock", log); err != nil {
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

// ParseRewardLock is a log parse operation binding the contract event 0xc43b85b876df0e1fcd072756642cda721ec62416510229af71d47e3a3513b730.
//
// Solidity: event RewardLock(bool rewardLock)
func (_SCHStake *SCHStakeFilterer) ParseRewardLock(log types.Log) (*SCHStakeRewardLock, error) {
	event := new(SCHStakeRewardLock)
	if err := _SCHStake.contract.UnpackLog(event, "RewardLock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SCHStakeRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the SCHStake contract.
type SCHStakeRoleAdminChangedIterator struct {
	Event *SCHStakeRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *SCHStakeRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SCHStakeRoleAdminChanged)
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
		it.Event = new(SCHStakeRoleAdminChanged)
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
func (it *SCHStakeRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SCHStakeRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SCHStakeRoleAdminChanged represents a RoleAdminChanged event raised by the SCHStake contract.
type SCHStakeRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_SCHStake *SCHStakeFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*SCHStakeRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _SCHStake.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &SCHStakeRoleAdminChangedIterator{contract: _SCHStake.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_SCHStake *SCHStakeFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *SCHStakeRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _SCHStake.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SCHStakeRoleAdminChanged)
				if err := _SCHStake.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_SCHStake *SCHStakeFilterer) ParseRoleAdminChanged(log types.Log) (*SCHStakeRoleAdminChanged, error) {
	event := new(SCHStakeRoleAdminChanged)
	if err := _SCHStake.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SCHStakeRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the SCHStake contract.
type SCHStakeRoleGrantedIterator struct {
	Event *SCHStakeRoleGranted // Event containing the contract specifics and raw log

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
func (it *SCHStakeRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SCHStakeRoleGranted)
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
		it.Event = new(SCHStakeRoleGranted)
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
func (it *SCHStakeRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SCHStakeRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SCHStakeRoleGranted represents a RoleGranted event raised by the SCHStake contract.
type SCHStakeRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_SCHStake *SCHStakeFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*SCHStakeRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SCHStake.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SCHStakeRoleGrantedIterator{contract: _SCHStake.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_SCHStake *SCHStakeFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *SCHStakeRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SCHStake.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SCHStakeRoleGranted)
				if err := _SCHStake.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_SCHStake *SCHStakeFilterer) ParseRoleGranted(log types.Log) (*SCHStakeRoleGranted, error) {
	event := new(SCHStakeRoleGranted)
	if err := _SCHStake.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SCHStakeRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the SCHStake contract.
type SCHStakeRoleRevokedIterator struct {
	Event *SCHStakeRoleRevoked // Event containing the contract specifics and raw log

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
func (it *SCHStakeRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SCHStakeRoleRevoked)
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
		it.Event = new(SCHStakeRoleRevoked)
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
func (it *SCHStakeRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SCHStakeRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SCHStakeRoleRevoked represents a RoleRevoked event raised by the SCHStake contract.
type SCHStakeRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_SCHStake *SCHStakeFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*SCHStakeRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SCHStake.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SCHStakeRoleRevokedIterator{contract: _SCHStake.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_SCHStake *SCHStakeFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *SCHStakeRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SCHStake.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SCHStakeRoleRevoked)
				if err := _SCHStake.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_SCHStake *SCHStakeFilterer) ParseRoleRevoked(log types.Log) (*SCHStakeRoleRevoked, error) {
	event := new(SCHStakeRoleRevoked)
	if err := _SCHStake.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SCHStakeSCHTokenIterator is returned from FilterSCHToken and is used to iterate over the raw logs and unpacked data for SCHToken events raised by the SCHStake contract.
type SCHStakeSCHTokenIterator struct {
	Event *SCHStakeSCHToken // Event containing the contract specifics and raw log

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
func (it *SCHStakeSCHTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SCHStakeSCHToken)
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
		it.Event = new(SCHStakeSCHToken)
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
func (it *SCHStakeSCHTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SCHStakeSCHTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SCHStakeSCHToken represents a SCHToken event raised by the SCHStake contract.
type SCHStakeSCHToken struct {
	SchToken common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSCHToken is a free log retrieval operation binding the contract event 0xce93fbbaadd3acd0d0b97ae44df0cf7efd123699184e04db523373302246700b.
//
// Solidity: event SCHToken(address schToken)
func (_SCHStake *SCHStakeFilterer) FilterSCHToken(opts *bind.FilterOpts) (*SCHStakeSCHTokenIterator, error) {

	logs, sub, err := _SCHStake.contract.FilterLogs(opts, "SCHToken")
	if err != nil {
		return nil, err
	}
	return &SCHStakeSCHTokenIterator{contract: _SCHStake.contract, event: "SCHToken", logs: logs, sub: sub}, nil
}

// WatchSCHToken is a free log subscription operation binding the contract event 0xce93fbbaadd3acd0d0b97ae44df0cf7efd123699184e04db523373302246700b.
//
// Solidity: event SCHToken(address schToken)
func (_SCHStake *SCHStakeFilterer) WatchSCHToken(opts *bind.WatchOpts, sink chan<- *SCHStakeSCHToken) (event.Subscription, error) {

	logs, sub, err := _SCHStake.contract.WatchLogs(opts, "SCHToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SCHStakeSCHToken)
				if err := _SCHStake.contract.UnpackLog(event, "SCHToken", log); err != nil {
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

// ParseSCHToken is a log parse operation binding the contract event 0xce93fbbaadd3acd0d0b97ae44df0cf7efd123699184e04db523373302246700b.
//
// Solidity: event SCHToken(address schToken)
func (_SCHStake *SCHStakeFilterer) ParseSCHToken(log types.Log) (*SCHStakeSCHToken, error) {
	event := new(SCHStakeSCHToken)
	if err := _SCHStake.contract.UnpackLog(event, "SCHToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SCHStakeStakeIterator is returned from FilterStake and is used to iterate over the raw logs and unpacked data for Stake events raised by the SCHStake contract.
type SCHStakeStakeIterator struct {
	Event *SCHStakeStake // Event containing the contract specifics and raw log

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
func (it *SCHStakeStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SCHStakeStake)
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
		it.Event = new(SCHStakeStake)
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
func (it *SCHStakeStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SCHStakeStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SCHStakeStake represents a Stake event raised by the SCHStake contract.
type SCHStakeStake struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStake is a free log retrieval operation binding the contract event 0x5af417134f72a9d41143ace85b0a26dce6f550f894f2cbc1eeee8810603d91b6.
//
// Solidity: event Stake(address indexed user, uint256 indexed pid, uint256 amount)
func (_SCHStake *SCHStakeFilterer) FilterStake(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*SCHStakeStakeIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _SCHStake.contract.FilterLogs(opts, "Stake", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &SCHStakeStakeIterator{contract: _SCHStake.contract, event: "Stake", logs: logs, sub: sub}, nil
}

// WatchStake is a free log subscription operation binding the contract event 0x5af417134f72a9d41143ace85b0a26dce6f550f894f2cbc1eeee8810603d91b6.
//
// Solidity: event Stake(address indexed user, uint256 indexed pid, uint256 amount)
func (_SCHStake *SCHStakeFilterer) WatchStake(opts *bind.WatchOpts, sink chan<- *SCHStakeStake, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _SCHStake.contract.WatchLogs(opts, "Stake", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SCHStakeStake)
				if err := _SCHStake.contract.UnpackLog(event, "Stake", log); err != nil {
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

// ParseStake is a log parse operation binding the contract event 0x5af417134f72a9d41143ace85b0a26dce6f550f894f2cbc1eeee8810603d91b6.
//
// Solidity: event Stake(address indexed user, uint256 indexed pid, uint256 amount)
func (_SCHStake *SCHStakeFilterer) ParseStake(log types.Log) (*SCHStakeStake, error) {
	event := new(SCHStakeStake)
	if err := _SCHStake.contract.UnpackLog(event, "Stake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SCHStakeStakeLockIterator is returned from FilterStakeLock and is used to iterate over the raw logs and unpacked data for StakeLock events raised by the SCHStake contract.
type SCHStakeStakeLockIterator struct {
	Event *SCHStakeStakeLock // Event containing the contract specifics and raw log

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
func (it *SCHStakeStakeLockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SCHStakeStakeLock)
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
		it.Event = new(SCHStakeStakeLock)
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
func (it *SCHStakeStakeLockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SCHStakeStakeLockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SCHStakeStakeLock represents a StakeLock event raised by the SCHStake contract.
type SCHStakeStakeLock struct {
	StakeLock bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStakeLock is a free log retrieval operation binding the contract event 0x395f8887af53af3d58612c4aaac86bfedaae72c4525be58f44e39ddfb1218158.
//
// Solidity: event StakeLock(bool stakeLock)
func (_SCHStake *SCHStakeFilterer) FilterStakeLock(opts *bind.FilterOpts) (*SCHStakeStakeLockIterator, error) {

	logs, sub, err := _SCHStake.contract.FilterLogs(opts, "StakeLock")
	if err != nil {
		return nil, err
	}
	return &SCHStakeStakeLockIterator{contract: _SCHStake.contract, event: "StakeLock", logs: logs, sub: sub}, nil
}

// WatchStakeLock is a free log subscription operation binding the contract event 0x395f8887af53af3d58612c4aaac86bfedaae72c4525be58f44e39ddfb1218158.
//
// Solidity: event StakeLock(bool stakeLock)
func (_SCHStake *SCHStakeFilterer) WatchStakeLock(opts *bind.WatchOpts, sink chan<- *SCHStakeStakeLock) (event.Subscription, error) {

	logs, sub, err := _SCHStake.contract.WatchLogs(opts, "StakeLock")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SCHStakeStakeLock)
				if err := _SCHStake.contract.UnpackLog(event, "StakeLock", log); err != nil {
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

// ParseStakeLock is a log parse operation binding the contract event 0x395f8887af53af3d58612c4aaac86bfedaae72c4525be58f44e39ddfb1218158.
//
// Solidity: event StakeLock(bool stakeLock)
func (_SCHStake *SCHStakeFilterer) ParseStakeLock(log types.Log) (*SCHStakeStakeLock, error) {
	event := new(SCHStakeStakeLock)
	if err := _SCHStake.contract.UnpackLog(event, "StakeLock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SCHStakeUpdatePoolIterator is returned from FilterUpdatePool and is used to iterate over the raw logs and unpacked data for UpdatePool events raised by the SCHStake contract.
type SCHStakeUpdatePoolIterator struct {
	Event *SCHStakeUpdatePool // Event containing the contract specifics and raw log

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
func (it *SCHStakeUpdatePoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SCHStakeUpdatePool)
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
		it.Event = new(SCHStakeUpdatePool)
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
func (it *SCHStakeUpdatePoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SCHStakeUpdatePoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SCHStakeUpdatePool represents a UpdatePool event raised by the SCHStake contract.
type SCHStakeUpdatePool struct {
	Pool SCHStakePool
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterUpdatePool is a free log retrieval operation binding the contract event 0xebcfe84f1a9a2860880852b549450ba0499407518875bd9bef41fd671b3da7c4.
//
// Solidity: event UpdatePool((address,uint256,uint256,uint256,uint256,uint256,uint256,uint256) pool)
func (_SCHStake *SCHStakeFilterer) FilterUpdatePool(opts *bind.FilterOpts) (*SCHStakeUpdatePoolIterator, error) {

	logs, sub, err := _SCHStake.contract.FilterLogs(opts, "UpdatePool")
	if err != nil {
		return nil, err
	}
	return &SCHStakeUpdatePoolIterator{contract: _SCHStake.contract, event: "UpdatePool", logs: logs, sub: sub}, nil
}

// WatchUpdatePool is a free log subscription operation binding the contract event 0xebcfe84f1a9a2860880852b549450ba0499407518875bd9bef41fd671b3da7c4.
//
// Solidity: event UpdatePool((address,uint256,uint256,uint256,uint256,uint256,uint256,uint256) pool)
func (_SCHStake *SCHStakeFilterer) WatchUpdatePool(opts *bind.WatchOpts, sink chan<- *SCHStakeUpdatePool) (event.Subscription, error) {

	logs, sub, err := _SCHStake.contract.WatchLogs(opts, "UpdatePool")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SCHStakeUpdatePool)
				if err := _SCHStake.contract.UnpackLog(event, "UpdatePool", log); err != nil {
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

// ParseUpdatePool is a log parse operation binding the contract event 0xebcfe84f1a9a2860880852b549450ba0499407518875bd9bef41fd671b3da7c4.
//
// Solidity: event UpdatePool((address,uint256,uint256,uint256,uint256,uint256,uint256,uint256) pool)
func (_SCHStake *SCHStakeFilterer) ParseUpdatePool(log types.Log) (*SCHStakeUpdatePool, error) {
	event := new(SCHStakeUpdatePool)
	if err := _SCHStake.contract.UnpackLog(event, "UpdatePool", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SCHStakeUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the SCHStake contract.
type SCHStakeUpgradedIterator struct {
	Event *SCHStakeUpgraded // Event containing the contract specifics and raw log

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
func (it *SCHStakeUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SCHStakeUpgraded)
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
		it.Event = new(SCHStakeUpgraded)
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
func (it *SCHStakeUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SCHStakeUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SCHStakeUpgraded represents a Upgraded event raised by the SCHStake contract.
type SCHStakeUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_SCHStake *SCHStakeFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*SCHStakeUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _SCHStake.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &SCHStakeUpgradedIterator{contract: _SCHStake.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_SCHStake *SCHStakeFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *SCHStakeUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _SCHStake.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SCHStakeUpgraded)
				if err := _SCHStake.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_SCHStake *SCHStakeFilterer) ParseUpgraded(log types.Log) (*SCHStakeUpgraded, error) {
	event := new(SCHStakeUpgraded)
	if err := _SCHStake.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
