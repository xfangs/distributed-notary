// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// LockedEthereumABI is the input ABI used to generate the binding from.
const LockedEthereumABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"secret\",\"type\":\"bytes32\"}],\"name\":\"lockout\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"cancleLockOut\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"secret_hash\",\"type\":\"bytes32\"},{\"name\":\"expiration\",\"type\":\"uint256\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"prepareLockoutHTLC\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"lockin_htlc\",\"outputs\":[{\"name\":\"SecretHash\",\"type\":\"bytes32\"},{\"name\":\"Expiration\",\"type\":\"uint256\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"queryLockin\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"cancelLockin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"secret\",\"type\":\"bytes32\"}],\"name\":\"lockin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"queryLockout\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"lockout_htlc\",\"outputs\":[{\"name\":\"SecretHash\",\"type\":\"bytes32\"},{\"name\":\"Expiration\",\"type\":\"uint256\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"secret_hash\",\"type\":\"bytes32\"},{\"name\":\"expiration\",\"type\":\"uint256\"}],\"name\":\"prepareLockin\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"PrepareLockin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"secret\",\"type\":\"bytes32\"}],\"name\":\"LockoutSecret\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"PrepareLockout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"secretHash\",\"type\":\"bytes32\"}],\"name\":\"Lockin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"secretHash\",\"type\":\"bytes32\"}],\"name\":\"CancelLockin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"secretHash\",\"type\":\"bytes32\"}],\"name\":\"CancelLockout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_prevOwner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"OwnerUpdate\",\"type\":\"event\"}]"

// LockedEthereumBin is the compiled bytecode used for deploying new contracts.
const LockedEthereumBin = `0x60018054600160a060020a031916905560c0604052601d60808190527f4c6f636b6564457468657265756d20666f722061746d6f73706865726500000060a090815261004e91600291906100b8565b506040805180820190915260048082527f76302e31000000000000000000000000000000000000000000000000000000006020909201918252610093916003916100b8565b503480156100a057600080fd5b5060008054600160a060020a03191633179055610153565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100f957805160ff1916838001178555610126565b82800160010185558215610126579182015b8281111561012657825182559160200191906001019061010b565b50610132929150610136565b5090565b61015091905b80821115610132576000815560010161013c565b90565b610be1806101626000396000f3006080604052600436106100e55763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663043d918081146100f757806306fdde031461011d57806310a276eb146101a75780631a10a238146101c85780631e0ef9a4146101f257806354fd4d501461023157806357e1ee591461024657806376188aa51461026757806379ba5097146102885780637fd408d21461029d5780638caa80f7146102c15780638da5cb5b146102e257806395d89b4114610313578063a6f9dae114610328578063b852876114610349578063e0ae1a811461036a575b3480156100f157600080fd5b50600080fd5b34801561010357600080fd5b5061011b600160a060020a0360043516602435610378565b005b34801561012957600080fd5b506101326104f3565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561016c578181015183820152602001610154565b50505050905090810190601f1680156101995780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156101b357600080fd5b5061011b600160a060020a036004351661057e565b3480156101d457600080fd5b5061011b600160a060020a0360043516602435604435606435610617565b3480156101fe57600080fd5b50610213600160a060020a03600435166106e4565b60408051938452602084019290925282820152519081900360600190f35b34801561023d57600080fd5b50610132610705565b34801561025257600080fd5b50610213600160a060020a0360043516610760565b34801561027357600080fd5b5061011b600160a060020a0360043516610788565b34801561029457600080fd5b5061011b610858565b3480156102a957600080fd5b5061011b600160a060020a03600435166024356108ef565b3480156102cd57600080fd5b50610213600160a060020a0360043516610a37565b3480156102ee57600080fd5b506102f7610a5f565b60408051600160a060020a039092168252519081900360200190f35b34801561031f57600080fd5b50610132610a6e565b34801561033457600080fd5b5061011b600160a060020a0360043516610aa5565b34801561035557600080fd5b50610213600160a060020a0360043516610b06565b61011b600435602435610b27565b600160a060020a03821660009081526005602052604081206002810154909181116103a257600080fd5b600182015443106103b257600080fd5b604080516020808201869052825180830382018152918301928390528151600293918291908401908083835b602083106103fd5780518252601f1990920191602091820191016103de565b51815160209384036101000a600019018019909216911617905260405191909301945091925050808303816000865af115801561043e573d6000803e3d6000fd5b5050506040513d602081101561045357600080fd5b505182541461046157600080fd5b60006002830181905580835560018301819055604051600160a060020a0386169183156108fc02918491818181858888f193505050501580156104a8573d6000803e3d6000fd5b5060408051600160a060020a03861681526020810185905281517fa0cfd4562aeab0234916ed60532417d84246c70a7f817dfc44e9c3d3423a84d3929181900390910190a150505050565b6002805460408051602060018416156101000260001901909316849004601f810184900484028201840190925281815292918301828280156105765780601f1061054b57610100808354040283529160200191610576565b820191906000526020600020905b81548152906001019060200180831161055957829003601f168201915b505050505081565b600160a060020a0381166000908152600560205260408120600281015490918082116105a957600080fd5b600183015443116105b957600080fd5b508154600060028401819055808455600184015560408051600160a060020a03861681526020810183905281517f625a628f697109c4cbee05890a8e5accf4b75c40503b6d8480b11715d148c2db929181900390910190a150505050565b60008054600160a060020a0316331461062f57600080fd5b600160a060020a038516151561064457600080fd5b50600160a060020a038416600090815260056020526040812090821161066957600080fd5b60028101541561067857600080fd5b61012c4301831161068857600080fd5b600281018290558381556001810183905560408051600160a060020a03871681526020810184905281517fbbae3304c67c8fbb052efa093374fc235534c3d862512a40007e7e35062a0475929181900390910190a15050505050565b60046020526000908152604090208054600182015460029092015490919083565b6003805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156105765780601f1061054b57610100808354040283529160200191610576565b600160a060020a03166000908152600460205260409020805460018201546002909201549092565b600160a060020a0381166000908152600460205260408120600281015490918082116107b357600080fd5b600183015443116107c357600080fd5b50815460006002840181905580845560018401819055604051600160a060020a0386169184156108fc02918591818181858888f1935050505015801561080d573d6000803e3d6000fd5b5060408051600160a060020a03861681526020810183905281517f026b98a8ac743c75f99f54b50949aa5e66574f9b73738858c62935046e4aa6c9929181900390910190a150505050565b600154600160a060020a0316331461086f57600080fd5b60005460015460408051600160a060020a03938416815292909116602083015280517f343765429aea5a34b3ff6a3785a98a5abb2597aca87bfbb58632c173d585373a9281900390910190a1600180546000805473ffffffffffffffffffffffffffffffffffffffff19908116600160a060020a03841617909155169055565b600160a060020a03821660009081526004602052604081206002810154909190811061091a57600080fd5b604080516020808201869052825180830382018152918301928390528151600293918291908401908083835b602083106109655780518252601f199092019160209182019101610946565b51815160209384036101000a600019018019909216911617905260405191909301945091925050808303816000865af11580156109a6573d6000803e3d6000fd5b5050506040513d60208110156109bb57600080fd5b50518254146109c957600080fd5b600182015443106109d957600080fd5b508054600060028301819055808355600183015560408051600160a060020a03861681526020810183905281517f0c89a242247566f6482a4febbbda97a1676fb18de194f38bf8f53d2d7a792c15929181900390910190a150505050565b600160a060020a03166000908152600560205260409020805460018201546002909201549092565b600054600160a060020a031681565b60408051808201909152600481527f4c45544800000000000000000000000000000000000000000000000000000000602082015281565b600054600160a060020a03163314610abc57600080fd5b600054600160a060020a0382811691161415610ad757600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60056020526000908152604090208054600182015460029092015490919083565b3360009081526004602052604081206002015415610b4457600080fd5b60003411610b5157600080fd5b5033600081815260046020908152604091829020858155600181018590553460028201819055835194855291840191909152815190927f1cc3ff93fb861f5fb2869fc15945f233d14ea7a4afa5721ad3c9804be90f3c6a92908290030190a15050505600a165627a7a7230582041cacc35ab048f90930c8bc0934a3b8e74b518df97cabd2534ef005c87cf92900029`

// DeployLockedEthereum deploys a new Ethereum contract, binding an instance of LockedEthereum to it.
func DeployLockedEthereum(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *LockedEthereum, error) {
	parsed, err := abi.JSON(strings.NewReader(LockedEthereumABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(LockedEthereumBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LockedEthereum{LockedEthereumCaller: LockedEthereumCaller{contract: contract}, LockedEthereumTransactor: LockedEthereumTransactor{contract: contract}, LockedEthereumFilterer: LockedEthereumFilterer{contract: contract}}, nil
}

// LockedEthereum is an auto generated Go binding around an Ethereum contract.
type LockedEthereum struct {
	LockedEthereumCaller     // Read-only binding to the contract
	LockedEthereumTransactor // Write-only binding to the contract
	LockedEthereumFilterer   // Log filterer for contract events
}

// LockedEthereumCaller is an auto generated read-only Go binding around an Ethereum contract.
type LockedEthereumCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LockedEthereumTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LockedEthereumTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LockedEthereumFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LockedEthereumFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LockedEthereumSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LockedEthereumSession struct {
	Contract     *LockedEthereum   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LockedEthereumCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LockedEthereumCallerSession struct {
	Contract *LockedEthereumCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// LockedEthereumTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LockedEthereumTransactorSession struct {
	Contract     *LockedEthereumTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// LockedEthereumRaw is an auto generated low-level Go binding around an Ethereum contract.
type LockedEthereumRaw struct {
	Contract *LockedEthereum // Generic contract binding to access the raw methods on
}

// LockedEthereumCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LockedEthereumCallerRaw struct {
	Contract *LockedEthereumCaller // Generic read-only contract binding to access the raw methods on
}

// LockedEthereumTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LockedEthereumTransactorRaw struct {
	Contract *LockedEthereumTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLockedEthereum creates a new instance of LockedEthereum, bound to a specific deployed contract.
func NewLockedEthereum(address common.Address, backend bind.ContractBackend) (*LockedEthereum, error) {
	contract, err := bindLockedEthereum(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LockedEthereum{LockedEthereumCaller: LockedEthereumCaller{contract: contract}, LockedEthereumTransactor: LockedEthereumTransactor{contract: contract}, LockedEthereumFilterer: LockedEthereumFilterer{contract: contract}}, nil
}

// NewLockedEthereumCaller creates a new read-only instance of LockedEthereum, bound to a specific deployed contract.
func NewLockedEthereumCaller(address common.Address, caller bind.ContractCaller) (*LockedEthereumCaller, error) {
	contract, err := bindLockedEthereum(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LockedEthereumCaller{contract: contract}, nil
}

// NewLockedEthereumTransactor creates a new write-only instance of LockedEthereum, bound to a specific deployed contract.
func NewLockedEthereumTransactor(address common.Address, transactor bind.ContractTransactor) (*LockedEthereumTransactor, error) {
	contract, err := bindLockedEthereum(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LockedEthereumTransactor{contract: contract}, nil
}

// NewLockedEthereumFilterer creates a new log filterer instance of LockedEthereum, bound to a specific deployed contract.
func NewLockedEthereumFilterer(address common.Address, filterer bind.ContractFilterer) (*LockedEthereumFilterer, error) {
	contract, err := bindLockedEthereum(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LockedEthereumFilterer{contract: contract}, nil
}

// bindLockedEthereum binds a generic wrapper to an already deployed contract.
func bindLockedEthereum(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LockedEthereumABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LockedEthereum *LockedEthereumRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LockedEthereum.Contract.LockedEthereumCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LockedEthereum *LockedEthereumRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LockedEthereum.Contract.LockedEthereumTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LockedEthereum *LockedEthereumRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LockedEthereum.Contract.LockedEthereumTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LockedEthereum *LockedEthereumCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LockedEthereum.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LockedEthereum *LockedEthereumTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LockedEthereum.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LockedEthereum *LockedEthereumTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LockedEthereum.Contract.contract.Transact(opts, method, params...)
}

// LockinHtlc is a free data retrieval call binding the contract method 0x1e0ef9a4.
//
// Solidity: function lockin_htlc( address) constant returns(SecretHash bytes32, Expiration uint256, value uint256)
func (_LockedEthereum *LockedEthereumCaller) LockinHtlc(opts *bind.CallOpts, arg0 common.Address) (struct {
	SecretHash [32]byte
	Expiration *big.Int
	Value      *big.Int
}, error) {
	ret := new(struct {
		SecretHash [32]byte
		Expiration *big.Int
		Value      *big.Int
	})
	out := ret
	err := _LockedEthereum.contract.Call(opts, out, "lockin_htlc", arg0)
	return *ret, err
}

// LockinHtlc is a free data retrieval call binding the contract method 0x1e0ef9a4.
//
// Solidity: function lockin_htlc( address) constant returns(SecretHash bytes32, Expiration uint256, value uint256)
func (_LockedEthereum *LockedEthereumSession) LockinHtlc(arg0 common.Address) (struct {
	SecretHash [32]byte
	Expiration *big.Int
	Value      *big.Int
}, error) {
	return _LockedEthereum.Contract.LockinHtlc(&_LockedEthereum.CallOpts, arg0)
}

// LockinHtlc is a free data retrieval call binding the contract method 0x1e0ef9a4.
//
// Solidity: function lockin_htlc( address) constant returns(SecretHash bytes32, Expiration uint256, value uint256)
func (_LockedEthereum *LockedEthereumCallerSession) LockinHtlc(arg0 common.Address) (struct {
	SecretHash [32]byte
	Expiration *big.Int
	Value      *big.Int
}, error) {
	return _LockedEthereum.Contract.LockinHtlc(&_LockedEthereum.CallOpts, arg0)
}

// LockoutHtlc is a free data retrieval call binding the contract method 0xb8528761.
//
// Solidity: function lockout_htlc( address) constant returns(SecretHash bytes32, Expiration uint256, value uint256)
func (_LockedEthereum *LockedEthereumCaller) LockoutHtlc(opts *bind.CallOpts, arg0 common.Address) (struct {
	SecretHash [32]byte
	Expiration *big.Int
	Value      *big.Int
}, error) {
	ret := new(struct {
		SecretHash [32]byte
		Expiration *big.Int
		Value      *big.Int
	})
	out := ret
	err := _LockedEthereum.contract.Call(opts, out, "lockout_htlc", arg0)
	return *ret, err
}

// LockoutHtlc is a free data retrieval call binding the contract method 0xb8528761.
//
// Solidity: function lockout_htlc( address) constant returns(SecretHash bytes32, Expiration uint256, value uint256)
func (_LockedEthereum *LockedEthereumSession) LockoutHtlc(arg0 common.Address) (struct {
	SecretHash [32]byte
	Expiration *big.Int
	Value      *big.Int
}, error) {
	return _LockedEthereum.Contract.LockoutHtlc(&_LockedEthereum.CallOpts, arg0)
}

// LockoutHtlc is a free data retrieval call binding the contract method 0xb8528761.
//
// Solidity: function lockout_htlc( address) constant returns(SecretHash bytes32, Expiration uint256, value uint256)
func (_LockedEthereum *LockedEthereumCallerSession) LockoutHtlc(arg0 common.Address) (struct {
	SecretHash [32]byte
	Expiration *big.Int
	Value      *big.Int
}, error) {
	return _LockedEthereum.Contract.LockoutHtlc(&_LockedEthereum.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_LockedEthereum *LockedEthereumCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _LockedEthereum.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_LockedEthereum *LockedEthereumSession) Name() (string, error) {
	return _LockedEthereum.Contract.Name(&_LockedEthereum.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_LockedEthereum *LockedEthereumCallerSession) Name() (string, error) {
	return _LockedEthereum.Contract.Name(&_LockedEthereum.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_LockedEthereum *LockedEthereumCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LockedEthereum.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_LockedEthereum *LockedEthereumSession) Owner() (common.Address, error) {
	return _LockedEthereum.Contract.Owner(&_LockedEthereum.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_LockedEthereum *LockedEthereumCallerSession) Owner() (common.Address, error) {
	return _LockedEthereum.Contract.Owner(&_LockedEthereum.CallOpts)
}

// QueryLockin is a free data retrieval call binding the contract method 0x57e1ee59.
//
// Solidity: function queryLockin(account address) constant returns(bytes32, uint256, uint256)
func (_LockedEthereum *LockedEthereumCaller) QueryLockin(opts *bind.CallOpts, account common.Address) ([32]byte, *big.Int, *big.Int, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _LockedEthereum.contract.Call(opts, out, "queryLockin", account)
	return *ret0, *ret1, *ret2, err
}

// QueryLockin is a free data retrieval call binding the contract method 0x57e1ee59.
//
// Solidity: function queryLockin(account address) constant returns(bytes32, uint256, uint256)
func (_LockedEthereum *LockedEthereumSession) QueryLockin(account common.Address) ([32]byte, *big.Int, *big.Int, error) {
	return _LockedEthereum.Contract.QueryLockin(&_LockedEthereum.CallOpts, account)
}

// QueryLockin is a free data retrieval call binding the contract method 0x57e1ee59.
//
// Solidity: function queryLockin(account address) constant returns(bytes32, uint256, uint256)
func (_LockedEthereum *LockedEthereumCallerSession) QueryLockin(account common.Address) ([32]byte, *big.Int, *big.Int, error) {
	return _LockedEthereum.Contract.QueryLockin(&_LockedEthereum.CallOpts, account)
}

// QueryLockout is a free data retrieval call binding the contract method 0x8caa80f7.
//
// Solidity: function queryLockout(account address) constant returns(bytes32, uint256, uint256)
func (_LockedEthereum *LockedEthereumCaller) QueryLockout(opts *bind.CallOpts, account common.Address) ([32]byte, *big.Int, *big.Int, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _LockedEthereum.contract.Call(opts, out, "queryLockout", account)
	return *ret0, *ret1, *ret2, err
}

// QueryLockout is a free data retrieval call binding the contract method 0x8caa80f7.
//
// Solidity: function queryLockout(account address) constant returns(bytes32, uint256, uint256)
func (_LockedEthereum *LockedEthereumSession) QueryLockout(account common.Address) ([32]byte, *big.Int, *big.Int, error) {
	return _LockedEthereum.Contract.QueryLockout(&_LockedEthereum.CallOpts, account)
}

// QueryLockout is a free data retrieval call binding the contract method 0x8caa80f7.
//
// Solidity: function queryLockout(account address) constant returns(bytes32, uint256, uint256)
func (_LockedEthereum *LockedEthereumCallerSession) QueryLockout(account common.Address) ([32]byte, *big.Int, *big.Int, error) {
	return _LockedEthereum.Contract.QueryLockout(&_LockedEthereum.CallOpts, account)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_LockedEthereum *LockedEthereumCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _LockedEthereum.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_LockedEthereum *LockedEthereumSession) Symbol() (string, error) {
	return _LockedEthereum.Contract.Symbol(&_LockedEthereum.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_LockedEthereum *LockedEthereumCallerSession) Symbol() (string, error) {
	return _LockedEthereum.Contract.Symbol(&_LockedEthereum.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_LockedEthereum *LockedEthereumCaller) Version(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _LockedEthereum.contract.Call(opts, out, "version")
	return *ret0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_LockedEthereum *LockedEthereumSession) Version() (string, error) {
	return _LockedEthereum.Contract.Version(&_LockedEthereum.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_LockedEthereum *LockedEthereumCallerSession) Version() (string, error) {
	return _LockedEthereum.Contract.Version(&_LockedEthereum.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_LockedEthereum *LockedEthereumTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LockedEthereum.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_LockedEthereum *LockedEthereumSession) AcceptOwnership() (*types.Transaction, error) {
	return _LockedEthereum.Contract.AcceptOwnership(&_LockedEthereum.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_LockedEthereum *LockedEthereumTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _LockedEthereum.Contract.AcceptOwnership(&_LockedEthereum.TransactOpts)
}

// CancelLockin is a paid mutator transaction binding the contract method 0x76188aa5.
//
// Solidity: function cancelLockin(account address) returns()
func (_LockedEthereum *LockedEthereumTransactor) CancelLockin(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _LockedEthereum.contract.Transact(opts, "cancelLockin", account)
}

// CancelLockin is a paid mutator transaction binding the contract method 0x76188aa5.
//
// Solidity: function cancelLockin(account address) returns()
func (_LockedEthereum *LockedEthereumSession) CancelLockin(account common.Address) (*types.Transaction, error) {
	return _LockedEthereum.Contract.CancelLockin(&_LockedEthereum.TransactOpts, account)
}

// CancelLockin is a paid mutator transaction binding the contract method 0x76188aa5.
//
// Solidity: function cancelLockin(account address) returns()
func (_LockedEthereum *LockedEthereumTransactorSession) CancelLockin(account common.Address) (*types.Transaction, error) {
	return _LockedEthereum.Contract.CancelLockin(&_LockedEthereum.TransactOpts, account)
}

// CancleLockOut is a paid mutator transaction binding the contract method 0x10a276eb.
//
// Solidity: function cancleLockOut(account address) returns()
func (_LockedEthereum *LockedEthereumTransactor) CancleLockOut(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _LockedEthereum.contract.Transact(opts, "cancleLockOut", account)
}

// CancleLockOut is a paid mutator transaction binding the contract method 0x10a276eb.
//
// Solidity: function cancleLockOut(account address) returns()
func (_LockedEthereum *LockedEthereumSession) CancleLockOut(account common.Address) (*types.Transaction, error) {
	return _LockedEthereum.Contract.CancleLockOut(&_LockedEthereum.TransactOpts, account)
}

// CancleLockOut is a paid mutator transaction binding the contract method 0x10a276eb.
//
// Solidity: function cancleLockOut(account address) returns()
func (_LockedEthereum *LockedEthereumTransactorSession) CancleLockOut(account common.Address) (*types.Transaction, error) {
	return _LockedEthereum.Contract.CancleLockOut(&_LockedEthereum.TransactOpts, account)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_newOwner address) returns()
func (_LockedEthereum *LockedEthereumTransactor) ChangeOwner(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _LockedEthereum.contract.Transact(opts, "changeOwner", _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_newOwner address) returns()
func (_LockedEthereum *LockedEthereumSession) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _LockedEthereum.Contract.ChangeOwner(&_LockedEthereum.TransactOpts, _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_newOwner address) returns()
func (_LockedEthereum *LockedEthereumTransactorSession) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _LockedEthereum.Contract.ChangeOwner(&_LockedEthereum.TransactOpts, _newOwner)
}

// Lockin is a paid mutator transaction binding the contract method 0x7fd408d2.
//
// Solidity: function lockin(account address, secret bytes32) returns()
func (_LockedEthereum *LockedEthereumTransactor) Lockin(opts *bind.TransactOpts, account common.Address, secret [32]byte) (*types.Transaction, error) {
	return _LockedEthereum.contract.Transact(opts, "lockin", account, secret)
}

// Lockin is a paid mutator transaction binding the contract method 0x7fd408d2.
//
// Solidity: function lockin(account address, secret bytes32) returns()
func (_LockedEthereum *LockedEthereumSession) Lockin(account common.Address, secret [32]byte) (*types.Transaction, error) {
	return _LockedEthereum.Contract.Lockin(&_LockedEthereum.TransactOpts, account, secret)
}

// Lockin is a paid mutator transaction binding the contract method 0x7fd408d2.
//
// Solidity: function lockin(account address, secret bytes32) returns()
func (_LockedEthereum *LockedEthereumTransactorSession) Lockin(account common.Address, secret [32]byte) (*types.Transaction, error) {
	return _LockedEthereum.Contract.Lockin(&_LockedEthereum.TransactOpts, account, secret)
}

// Lockout is a paid mutator transaction binding the contract method 0x043d9180.
//
// Solidity: function lockout(account address, secret bytes32) returns()
func (_LockedEthereum *LockedEthereumTransactor) Lockout(opts *bind.TransactOpts, account common.Address, secret [32]byte) (*types.Transaction, error) {
	return _LockedEthereum.contract.Transact(opts, "lockout", account, secret)
}

// Lockout is a paid mutator transaction binding the contract method 0x043d9180.
//
// Solidity: function lockout(account address, secret bytes32) returns()
func (_LockedEthereum *LockedEthereumSession) Lockout(account common.Address, secret [32]byte) (*types.Transaction, error) {
	return _LockedEthereum.Contract.Lockout(&_LockedEthereum.TransactOpts, account, secret)
}

// Lockout is a paid mutator transaction binding the contract method 0x043d9180.
//
// Solidity: function lockout(account address, secret bytes32) returns()
func (_LockedEthereum *LockedEthereumTransactorSession) Lockout(account common.Address, secret [32]byte) (*types.Transaction, error) {
	return _LockedEthereum.Contract.Lockout(&_LockedEthereum.TransactOpts, account, secret)
}

// PrepareLockin is a paid mutator transaction binding the contract method 0xe0ae1a81.
//
// Solidity: function prepareLockin(secret_hash bytes32, expiration uint256) returns()
func (_LockedEthereum *LockedEthereumTransactor) PrepareLockin(opts *bind.TransactOpts, secret_hash [32]byte, expiration *big.Int) (*types.Transaction, error) {
	return _LockedEthereum.contract.Transact(opts, "prepareLockin", secret_hash, expiration)
}

// PrepareLockin is a paid mutator transaction binding the contract method 0xe0ae1a81.
//
// Solidity: function prepareLockin(secret_hash bytes32, expiration uint256) returns()
func (_LockedEthereum *LockedEthereumSession) PrepareLockin(secret_hash [32]byte, expiration *big.Int) (*types.Transaction, error) {
	return _LockedEthereum.Contract.PrepareLockin(&_LockedEthereum.TransactOpts, secret_hash, expiration)
}

// PrepareLockin is a paid mutator transaction binding the contract method 0xe0ae1a81.
//
// Solidity: function prepareLockin(secret_hash bytes32, expiration uint256) returns()
func (_LockedEthereum *LockedEthereumTransactorSession) PrepareLockin(secret_hash [32]byte, expiration *big.Int) (*types.Transaction, error) {
	return _LockedEthereum.Contract.PrepareLockin(&_LockedEthereum.TransactOpts, secret_hash, expiration)
}

// PrepareLockoutHTLC is a paid mutator transaction binding the contract method 0x1a10a238.
//
// Solidity: function prepareLockoutHTLC(account address, secret_hash bytes32, expiration uint256, value uint256) returns()
func (_LockedEthereum *LockedEthereumTransactor) PrepareLockoutHTLC(opts *bind.TransactOpts, account common.Address, secret_hash [32]byte, expiration *big.Int, value *big.Int) (*types.Transaction, error) {
	return _LockedEthereum.contract.Transact(opts, "prepareLockoutHTLC", account, secret_hash, expiration, value)
}

// PrepareLockoutHTLC is a paid mutator transaction binding the contract method 0x1a10a238.
//
// Solidity: function prepareLockoutHTLC(account address, secret_hash bytes32, expiration uint256, value uint256) returns()
func (_LockedEthereum *LockedEthereumSession) PrepareLockoutHTLC(account common.Address, secret_hash [32]byte, expiration *big.Int, value *big.Int) (*types.Transaction, error) {
	return _LockedEthereum.Contract.PrepareLockoutHTLC(&_LockedEthereum.TransactOpts, account, secret_hash, expiration, value)
}

// PrepareLockoutHTLC is a paid mutator transaction binding the contract method 0x1a10a238.
//
// Solidity: function prepareLockoutHTLC(account address, secret_hash bytes32, expiration uint256, value uint256) returns()
func (_LockedEthereum *LockedEthereumTransactorSession) PrepareLockoutHTLC(account common.Address, secret_hash [32]byte, expiration *big.Int, value *big.Int) (*types.Transaction, error) {
	return _LockedEthereum.Contract.PrepareLockoutHTLC(&_LockedEthereum.TransactOpts, account, secret_hash, expiration, value)
}

// LockedEthereumCancelLockinIterator is returned from FilterCancelLockin and is used to iterate over the raw logs and unpacked data for CancelLockin events raised by the LockedEthereum contract.
type LockedEthereumCancelLockinIterator struct {
	Event *LockedEthereumCancelLockin // Event containing the contract specifics and raw log

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
func (it *LockedEthereumCancelLockinIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockedEthereumCancelLockin)
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
		it.Event = new(LockedEthereumCancelLockin)
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
func (it *LockedEthereumCancelLockinIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockedEthereumCancelLockinIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockedEthereumCancelLockin represents a CancelLockin event raised by the LockedEthereum contract.
type LockedEthereumCancelLockin struct {
	Account    common.Address
	SecretHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCancelLockin is a free log retrieval operation binding the contract event 0x026b98a8ac743c75f99f54b50949aa5e66574f9b73738858c62935046e4aa6c9.
//
// Solidity: e CancelLockin(account address, secretHash bytes32)
func (_LockedEthereum *LockedEthereumFilterer) FilterCancelLockin(opts *bind.FilterOpts) (*LockedEthereumCancelLockinIterator, error) {

	logs, sub, err := _LockedEthereum.contract.FilterLogs(opts, "CancelLockin")
	if err != nil {
		return nil, err
	}
	return &LockedEthereumCancelLockinIterator{contract: _LockedEthereum.contract, event: "CancelLockin", logs: logs, sub: sub}, nil
}

// WatchCancelLockin is a free log subscription operation binding the contract event 0x026b98a8ac743c75f99f54b50949aa5e66574f9b73738858c62935046e4aa6c9.
//
// Solidity: e CancelLockin(account address, secretHash bytes32)
func (_LockedEthereum *LockedEthereumFilterer) WatchCancelLockin(opts *bind.WatchOpts, sink chan<- *LockedEthereumCancelLockin) (event.Subscription, error) {

	logs, sub, err := _LockedEthereum.contract.WatchLogs(opts, "CancelLockin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockedEthereumCancelLockin)
				if err := _LockedEthereum.contract.UnpackLog(event, "CancelLockin", log); err != nil {
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

// LockedEthereumCancelLockoutIterator is returned from FilterCancelLockout and is used to iterate over the raw logs and unpacked data for CancelLockout events raised by the LockedEthereum contract.
type LockedEthereumCancelLockoutIterator struct {
	Event *LockedEthereumCancelLockout // Event containing the contract specifics and raw log

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
func (it *LockedEthereumCancelLockoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockedEthereumCancelLockout)
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
		it.Event = new(LockedEthereumCancelLockout)
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
func (it *LockedEthereumCancelLockoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockedEthereumCancelLockoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockedEthereumCancelLockout represents a CancelLockout event raised by the LockedEthereum contract.
type LockedEthereumCancelLockout struct {
	Account    common.Address
	SecretHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCancelLockout is a free log retrieval operation binding the contract event 0x625a628f697109c4cbee05890a8e5accf4b75c40503b6d8480b11715d148c2db.
//
// Solidity: e CancelLockout(account address, secretHash bytes32)
func (_LockedEthereum *LockedEthereumFilterer) FilterCancelLockout(opts *bind.FilterOpts) (*LockedEthereumCancelLockoutIterator, error) {

	logs, sub, err := _LockedEthereum.contract.FilterLogs(opts, "CancelLockout")
	if err != nil {
		return nil, err
	}
	return &LockedEthereumCancelLockoutIterator{contract: _LockedEthereum.contract, event: "CancelLockout", logs: logs, sub: sub}, nil
}

// WatchCancelLockout is a free log subscription operation binding the contract event 0x625a628f697109c4cbee05890a8e5accf4b75c40503b6d8480b11715d148c2db.
//
// Solidity: e CancelLockout(account address, secretHash bytes32)
func (_LockedEthereum *LockedEthereumFilterer) WatchCancelLockout(opts *bind.WatchOpts, sink chan<- *LockedEthereumCancelLockout) (event.Subscription, error) {

	logs, sub, err := _LockedEthereum.contract.WatchLogs(opts, "CancelLockout")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockedEthereumCancelLockout)
				if err := _LockedEthereum.contract.UnpackLog(event, "CancelLockout", log); err != nil {
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

// LockedEthereumLockinIterator is returned from FilterLockin and is used to iterate over the raw logs and unpacked data for Lockin events raised by the LockedEthereum contract.
type LockedEthereumLockinIterator struct {
	Event *LockedEthereumLockin // Event containing the contract specifics and raw log

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
func (it *LockedEthereumLockinIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockedEthereumLockin)
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
		it.Event = new(LockedEthereumLockin)
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
func (it *LockedEthereumLockinIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockedEthereumLockinIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockedEthereumLockin represents a Lockin event raised by the LockedEthereum contract.
type LockedEthereumLockin struct {
	Account    common.Address
	SecretHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLockin is a free log retrieval operation binding the contract event 0x0c89a242247566f6482a4febbbda97a1676fb18de194f38bf8f53d2d7a792c15.
//
// Solidity: e Lockin(account address, secretHash bytes32)
func (_LockedEthereum *LockedEthereumFilterer) FilterLockin(opts *bind.FilterOpts) (*LockedEthereumLockinIterator, error) {

	logs, sub, err := _LockedEthereum.contract.FilterLogs(opts, "Lockin")
	if err != nil {
		return nil, err
	}
	return &LockedEthereumLockinIterator{contract: _LockedEthereum.contract, event: "Lockin", logs: logs, sub: sub}, nil
}

// WatchLockin is a free log subscription operation binding the contract event 0x0c89a242247566f6482a4febbbda97a1676fb18de194f38bf8f53d2d7a792c15.
//
// Solidity: e Lockin(account address, secretHash bytes32)
func (_LockedEthereum *LockedEthereumFilterer) WatchLockin(opts *bind.WatchOpts, sink chan<- *LockedEthereumLockin) (event.Subscription, error) {

	logs, sub, err := _LockedEthereum.contract.WatchLogs(opts, "Lockin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockedEthereumLockin)
				if err := _LockedEthereum.contract.UnpackLog(event, "Lockin", log); err != nil {
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

// LockedEthereumLockoutSecretIterator is returned from FilterLockoutSecret and is used to iterate over the raw logs and unpacked data for LockoutSecret events raised by the LockedEthereum contract.
type LockedEthereumLockoutSecretIterator struct {
	Event *LockedEthereumLockoutSecret // Event containing the contract specifics and raw log

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
func (it *LockedEthereumLockoutSecretIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockedEthereumLockoutSecret)
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
		it.Event = new(LockedEthereumLockoutSecret)
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
func (it *LockedEthereumLockoutSecretIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockedEthereumLockoutSecretIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockedEthereumLockoutSecret represents a LockoutSecret event raised by the LockedEthereum contract.
type LockedEthereumLockoutSecret struct {
	Account common.Address
	Secret  [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLockoutSecret is a free log retrieval operation binding the contract event 0xa0cfd4562aeab0234916ed60532417d84246c70a7f817dfc44e9c3d3423a84d3.
//
// Solidity: e LockoutSecret(account address, secret bytes32)
func (_LockedEthereum *LockedEthereumFilterer) FilterLockoutSecret(opts *bind.FilterOpts) (*LockedEthereumLockoutSecretIterator, error) {

	logs, sub, err := _LockedEthereum.contract.FilterLogs(opts, "LockoutSecret")
	if err != nil {
		return nil, err
	}
	return &LockedEthereumLockoutSecretIterator{contract: _LockedEthereum.contract, event: "LockoutSecret", logs: logs, sub: sub}, nil
}

// WatchLockoutSecret is a free log subscription operation binding the contract event 0xa0cfd4562aeab0234916ed60532417d84246c70a7f817dfc44e9c3d3423a84d3.
//
// Solidity: e LockoutSecret(account address, secret bytes32)
func (_LockedEthereum *LockedEthereumFilterer) WatchLockoutSecret(opts *bind.WatchOpts, sink chan<- *LockedEthereumLockoutSecret) (event.Subscription, error) {

	logs, sub, err := _LockedEthereum.contract.WatchLogs(opts, "LockoutSecret")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockedEthereumLockoutSecret)
				if err := _LockedEthereum.contract.UnpackLog(event, "LockoutSecret", log); err != nil {
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

// LockedEthereumOwnerUpdateIterator is returned from FilterOwnerUpdate and is used to iterate over the raw logs and unpacked data for OwnerUpdate events raised by the LockedEthereum contract.
type LockedEthereumOwnerUpdateIterator struct {
	Event *LockedEthereumOwnerUpdate // Event containing the contract specifics and raw log

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
func (it *LockedEthereumOwnerUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockedEthereumOwnerUpdate)
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
		it.Event = new(LockedEthereumOwnerUpdate)
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
func (it *LockedEthereumOwnerUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockedEthereumOwnerUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockedEthereumOwnerUpdate represents a OwnerUpdate event raised by the LockedEthereum contract.
type LockedEthereumOwnerUpdate struct {
	PrevOwner common.Address
	NewOwner  common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOwnerUpdate is a free log retrieval operation binding the contract event 0x343765429aea5a34b3ff6a3785a98a5abb2597aca87bfbb58632c173d585373a.
//
// Solidity: e OwnerUpdate(_prevOwner address, _newOwner address)
func (_LockedEthereum *LockedEthereumFilterer) FilterOwnerUpdate(opts *bind.FilterOpts) (*LockedEthereumOwnerUpdateIterator, error) {

	logs, sub, err := _LockedEthereum.contract.FilterLogs(opts, "OwnerUpdate")
	if err != nil {
		return nil, err
	}
	return &LockedEthereumOwnerUpdateIterator{contract: _LockedEthereum.contract, event: "OwnerUpdate", logs: logs, sub: sub}, nil
}

// WatchOwnerUpdate is a free log subscription operation binding the contract event 0x343765429aea5a34b3ff6a3785a98a5abb2597aca87bfbb58632c173d585373a.
//
// Solidity: e OwnerUpdate(_prevOwner address, _newOwner address)
func (_LockedEthereum *LockedEthereumFilterer) WatchOwnerUpdate(opts *bind.WatchOpts, sink chan<- *LockedEthereumOwnerUpdate) (event.Subscription, error) {

	logs, sub, err := _LockedEthereum.contract.WatchLogs(opts, "OwnerUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockedEthereumOwnerUpdate)
				if err := _LockedEthereum.contract.UnpackLog(event, "OwnerUpdate", log); err != nil {
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

// LockedEthereumPrepareLockinIterator is returned from FilterPrepareLockin and is used to iterate over the raw logs and unpacked data for PrepareLockin events raised by the LockedEthereum contract.
type LockedEthereumPrepareLockinIterator struct {
	Event *LockedEthereumPrepareLockin // Event containing the contract specifics and raw log

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
func (it *LockedEthereumPrepareLockinIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockedEthereumPrepareLockin)
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
		it.Event = new(LockedEthereumPrepareLockin)
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
func (it *LockedEthereumPrepareLockinIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockedEthereumPrepareLockinIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockedEthereumPrepareLockin represents a PrepareLockin event raised by the LockedEthereum contract.
type LockedEthereumPrepareLockin struct {
	Account common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPrepareLockin is a free log retrieval operation binding the contract event 0x1cc3ff93fb861f5fb2869fc15945f233d14ea7a4afa5721ad3c9804be90f3c6a.
//
// Solidity: e PrepareLockin(account address, value uint256)
func (_LockedEthereum *LockedEthereumFilterer) FilterPrepareLockin(opts *bind.FilterOpts) (*LockedEthereumPrepareLockinIterator, error) {

	logs, sub, err := _LockedEthereum.contract.FilterLogs(opts, "PrepareLockin")
	if err != nil {
		return nil, err
	}
	return &LockedEthereumPrepareLockinIterator{contract: _LockedEthereum.contract, event: "PrepareLockin", logs: logs, sub: sub}, nil
}

// WatchPrepareLockin is a free log subscription operation binding the contract event 0x1cc3ff93fb861f5fb2869fc15945f233d14ea7a4afa5721ad3c9804be90f3c6a.
//
// Solidity: e PrepareLockin(account address, value uint256)
func (_LockedEthereum *LockedEthereumFilterer) WatchPrepareLockin(opts *bind.WatchOpts, sink chan<- *LockedEthereumPrepareLockin) (event.Subscription, error) {

	logs, sub, err := _LockedEthereum.contract.WatchLogs(opts, "PrepareLockin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockedEthereumPrepareLockin)
				if err := _LockedEthereum.contract.UnpackLog(event, "PrepareLockin", log); err != nil {
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

// LockedEthereumPrepareLockoutIterator is returned from FilterPrepareLockout and is used to iterate over the raw logs and unpacked data for PrepareLockout events raised by the LockedEthereum contract.
type LockedEthereumPrepareLockoutIterator struct {
	Event *LockedEthereumPrepareLockout // Event containing the contract specifics and raw log

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
func (it *LockedEthereumPrepareLockoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockedEthereumPrepareLockout)
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
		it.Event = new(LockedEthereumPrepareLockout)
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
func (it *LockedEthereumPrepareLockoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockedEthereumPrepareLockoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockedEthereumPrepareLockout represents a PrepareLockout event raised by the LockedEthereum contract.
type LockedEthereumPrepareLockout struct {
	Account common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPrepareLockout is a free log retrieval operation binding the contract event 0xbbae3304c67c8fbb052efa093374fc235534c3d862512a40007e7e35062a0475.
//
// Solidity: e PrepareLockout(account address, _value uint256)
func (_LockedEthereum *LockedEthereumFilterer) FilterPrepareLockout(opts *bind.FilterOpts) (*LockedEthereumPrepareLockoutIterator, error) {

	logs, sub, err := _LockedEthereum.contract.FilterLogs(opts, "PrepareLockout")
	if err != nil {
		return nil, err
	}
	return &LockedEthereumPrepareLockoutIterator{contract: _LockedEthereum.contract, event: "PrepareLockout", logs: logs, sub: sub}, nil
}

// WatchPrepareLockout is a free log subscription operation binding the contract event 0xbbae3304c67c8fbb052efa093374fc235534c3d862512a40007e7e35062a0475.
//
// Solidity: e PrepareLockout(account address, _value uint256)
func (_LockedEthereum *LockedEthereumFilterer) WatchPrepareLockout(opts *bind.WatchOpts, sink chan<- *LockedEthereumPrepareLockout) (event.Subscription, error) {

	logs, sub, err := _LockedEthereum.contract.WatchLogs(opts, "PrepareLockout")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockedEthereumPrepareLockout)
				if err := _LockedEthereum.contract.UnpackLog(event, "PrepareLockout", log); err != nil {
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

// OwnedABI is the input ABI used to generate the binding from.
const OwnedABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_prevOwner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"OwnerUpdate\",\"type\":\"event\"}]"

// OwnedBin is the compiled bytecode used for deploying new contracts.
const OwnedBin = `0x608060405260018054600160a060020a031916905534801561002057600080fd5b5060008054600160a060020a031916331790556101f7806100426000396000f3006080604052600436106100565763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166379ba5097811461005b5780638da5cb5b14610072578063a6f9dae1146100a3575b600080fd5b34801561006757600080fd5b506100706100c4565b005b34801561007e57600080fd5b5061008761015b565b60408051600160a060020a039092168252519081900360200190f35b3480156100af57600080fd5b50610070600160a060020a036004351661016a565b600154600160a060020a031633146100db57600080fd5b60005460015460408051600160a060020a03938416815292909116602083015280517f343765429aea5a34b3ff6a3785a98a5abb2597aca87bfbb58632c173d585373a9281900390910190a1600180546000805473ffffffffffffffffffffffffffffffffffffffff19908116600160a060020a03841617909155169055565b600054600160a060020a031681565b600054600160a060020a0316331461018157600080fd5b600054600160a060020a038281169116141561019c57600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a7230582089d2fd0d94a351b7ddca464a05d9fc8fbe7034defb995b4552c05ea8f6b9f2650029`

// DeployOwned deploys a new Ethereum contract, binding an instance of Owned to it.
func DeployOwned(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Owned, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnedABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OwnedBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Owned{OwnedCaller: OwnedCaller{contract: contract}, OwnedTransactor: OwnedTransactor{contract: contract}, OwnedFilterer: OwnedFilterer{contract: contract}}, nil
}

// Owned is an auto generated Go binding around an Ethereum contract.
type Owned struct {
	OwnedCaller     // Read-only binding to the contract
	OwnedTransactor // Write-only binding to the contract
	OwnedFilterer   // Log filterer for contract events
}

// OwnedCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnedSession struct {
	Contract     *Owned            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnedCallerSession struct {
	Contract *OwnedCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OwnedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnedTransactorSession struct {
	Contract     *OwnedTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnedRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnedRaw struct {
	Contract *Owned // Generic contract binding to access the raw methods on
}

// OwnedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnedCallerRaw struct {
	Contract *OwnedCaller // Generic read-only contract binding to access the raw methods on
}

// OwnedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnedTransactorRaw struct {
	Contract *OwnedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwned creates a new instance of Owned, bound to a specific deployed contract.
func NewOwned(address common.Address, backend bind.ContractBackend) (*Owned, error) {
	contract, err := bindOwned(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Owned{OwnedCaller: OwnedCaller{contract: contract}, OwnedTransactor: OwnedTransactor{contract: contract}, OwnedFilterer: OwnedFilterer{contract: contract}}, nil
}

// NewOwnedCaller creates a new read-only instance of Owned, bound to a specific deployed contract.
func NewOwnedCaller(address common.Address, caller bind.ContractCaller) (*OwnedCaller, error) {
	contract, err := bindOwned(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnedCaller{contract: contract}, nil
}

// NewOwnedTransactor creates a new write-only instance of Owned, bound to a specific deployed contract.
func NewOwnedTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnedTransactor, error) {
	contract, err := bindOwned(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnedTransactor{contract: contract}, nil
}

// NewOwnedFilterer creates a new log filterer instance of Owned, bound to a specific deployed contract.
func NewOwnedFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnedFilterer, error) {
	contract, err := bindOwned(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnedFilterer{contract: contract}, nil
}

// bindOwned binds a generic wrapper to an already deployed contract.
func bindOwned(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnedABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Owned *OwnedRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Owned.Contract.OwnedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Owned *OwnedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Owned.Contract.OwnedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Owned *OwnedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Owned.Contract.OwnedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Owned *OwnedCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Owned.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Owned *OwnedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Owned.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Owned *OwnedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Owned.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Owned *OwnedCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Owned.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Owned *OwnedSession) Owner() (common.Address, error) {
	return _Owned.Contract.Owner(&_Owned.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Owned *OwnedCallerSession) Owner() (common.Address, error) {
	return _Owned.Contract.Owner(&_Owned.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Owned *OwnedTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Owned.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Owned *OwnedSession) AcceptOwnership() (*types.Transaction, error) {
	return _Owned.Contract.AcceptOwnership(&_Owned.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Owned *OwnedTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Owned.Contract.AcceptOwnership(&_Owned.TransactOpts)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_newOwner address) returns()
func (_Owned *OwnedTransactor) ChangeOwner(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Owned.contract.Transact(opts, "changeOwner", _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_newOwner address) returns()
func (_Owned *OwnedSession) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _Owned.Contract.ChangeOwner(&_Owned.TransactOpts, _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_newOwner address) returns()
func (_Owned *OwnedTransactorSession) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _Owned.Contract.ChangeOwner(&_Owned.TransactOpts, _newOwner)
}

// OwnedOwnerUpdateIterator is returned from FilterOwnerUpdate and is used to iterate over the raw logs and unpacked data for OwnerUpdate events raised by the Owned contract.
type OwnedOwnerUpdateIterator struct {
	Event *OwnedOwnerUpdate // Event containing the contract specifics and raw log

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
func (it *OwnedOwnerUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnedOwnerUpdate)
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
		it.Event = new(OwnedOwnerUpdate)
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
func (it *OwnedOwnerUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnedOwnerUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnedOwnerUpdate represents a OwnerUpdate event raised by the Owned contract.
type OwnedOwnerUpdate struct {
	PrevOwner common.Address
	NewOwner  common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOwnerUpdate is a free log retrieval operation binding the contract event 0x343765429aea5a34b3ff6a3785a98a5abb2597aca87bfbb58632c173d585373a.
//
// Solidity: e OwnerUpdate(_prevOwner address, _newOwner address)
func (_Owned *OwnedFilterer) FilterOwnerUpdate(opts *bind.FilterOpts) (*OwnedOwnerUpdateIterator, error) {

	logs, sub, err := _Owned.contract.FilterLogs(opts, "OwnerUpdate")
	if err != nil {
		return nil, err
	}
	return &OwnedOwnerUpdateIterator{contract: _Owned.contract, event: "OwnerUpdate", logs: logs, sub: sub}, nil
}

// WatchOwnerUpdate is a free log subscription operation binding the contract event 0x343765429aea5a34b3ff6a3785a98a5abb2597aca87bfbb58632c173d585373a.
//
// Solidity: e OwnerUpdate(_prevOwner address, _newOwner address)
func (_Owned *OwnedFilterer) WatchOwnerUpdate(opts *bind.WatchOpts, sink chan<- *OwnedOwnerUpdate) (event.Subscription, error) {

	logs, sub, err := _Owned.contract.WatchLogs(opts, "OwnerUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnedOwnerUpdate)
				if err := _Owned.contract.UnpackLog(event, "OwnerUpdate", log); err != nil {
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
