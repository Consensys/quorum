// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package permission

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

// VoterManagerABI is the input ABI used to generate the binding from.
const VoterManagerABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"getPendingOpDetails\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_vAccount\",\"type\":\"address\"}],\"name\":\"addVoter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_vAccount\",\"type\":\"address\"}],\"name\":\"deleteVoter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_authOrg\",\"type\":\"string\"},{\"name\":\"_vAccount\",\"type\":\"address\"},{\"name\":\"_pendingOp\",\"type\":\"uint256\"}],\"name\":\"processVote\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_authOrg\",\"type\":\"string\"},{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_enodeId\",\"type\":\"string\"},{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_pendingOp\",\"type\":\"uint256\"}],\"name\":\"addVotingItem\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_permUpgradable\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_vAccount\",\"type\":\"address\"}],\"name\":\"VoterAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_vAccount\",\"type\":\"address\"}],\"name\":\"VoterDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"VotingItemAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"VoteProcessed\",\"type\":\"event\"}]"

var VoterManagerParsedABI, _ = abi.JSON(strings.NewReader(VoterManagerABI))

// VoterManagerBin is the compiled bytecode used for deploying new contracts.
var VoterManagerBin = "0x6080604052600060035534801561001557600080fd5b5060405160208061289c8339810180604052602081101561003557600080fd5b8101908080519060200190929190505050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050612806806100966000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c8063014e6acc1461005c5780635607395b146101f457806359cbd6fe1461028d578063b021386414610326578063e98ac22d146103e1575b600080fd5b6100d36004803603602081101561007257600080fd5b810190808035906020019064010000000081111561008f57600080fd5b8201836020820111156100a157600080fd5b803590602001918460018302840111640100000000831117156100c357600080fd5b909192939192939050505061052e565b6040518080602001806020018573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001848152602001838103835287818151815260200191508051906020019080838360005b8381101561014f578082015181840152602081019050610134565b50505050905090810190601f16801561017c5780820380516001836020036101000a031916815260200191505b50838103825286818151815260200191508051906020019080838360005b838110156101b557808201518184015260208101905061019a565b50505050905090810190601f1680156101e25780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390f35b61028b6004803603604081101561020a57600080fd5b810190808035906020019064010000000081111561022757600080fd5b82018360208201111561023957600080fd5b8035906020019184600183028401116401000000008311171561025b57600080fd5b9091929391929390803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506108c1565b005b610324600480360360408110156102a357600080fd5b81019080803590602001906401000000008111156102c057600080fd5b8201836020820111156102d257600080fd5b803590602001918460018302840111640100000000831117156102f457600080fd5b9091929391929390803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506112df565b005b6103c76004803603606081101561033c57600080fd5b810190808035906020019064010000000081111561035957600080fd5b82018360208201111561036b57600080fd5b8035906020019184600183028401116401000000008311171561038d57600080fd5b9091929391929390803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506116b5565b604051808215151515815260200191505060405180910390f35b61052c600480360360a08110156103f757600080fd5b810190808035906020019064010000000081111561041457600080fd5b82018360208201111561042657600080fd5b8035906020019184600183028401116401000000008311171561044857600080fd5b90919293919293908035906020019064010000000081111561046957600080fd5b82018360208201111561047b57600080fd5b8035906020019184600183028401116401000000008311171561049d57600080fd5b9091929391929390803590602001906401000000008111156104be57600080fd5b8201836020820111156104d057600080fd5b803590602001918460018302840111640100000000831117156104f257600080fd5b9091929391929390803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050611d79565b005b6060806000806000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630e32cf906040518163ffffffff1660e01b815260040160206040518083038186803b15801561059b57600080fd5b505afa1580156105af573d6000803e3d6000fd5b505050506040513d60208110156105c557600080fd5b810190808051906020019092919050505073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610678576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600e8152602001807f696e76616c69642063616c6c657200000000000000000000000000000000000081525060200191505060405180910390fd5b60006106c787878080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050506122a6565b90506001818154811015156106d857fe5b90600052602060002090600b02016004016000016001828154811015156106fb57fe5b90600052602060002090600b020160040160010160018381548110151561071e57fe5b90600052602060002090600b020160040160020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660018481548110151561076257fe5b90600052602060002090600b020160040160030154838054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561080c5780601f106107e15761010080835404028352916020019161080c565b820191906000526020600020905b8154815290600101906020018083116107ef57829003601f168201915b50505050509350828054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156108a85780601f1061087d576101008083540402835291602001916108a8565b820191906000526020600020905b81548152906001019060200180831161088b57829003601f168201915b5050505050925094509450945094505092959194509250565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630e32cf906040518163ffffffff1660e01b815260040160206040518083038186803b15801561092857600080fd5b505afa15801561093c573d6000803e3d6000fd5b505050506040513d602081101561095257600080fd5b810190808051906020019092919050505073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610a05576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600e8152602001807f696e76616c69642063616c6c657200000000000000000000000000000000000081525060200191505060405180910390fd5b600060026000858560405160200180806020018281038252848482818152602001925080828437600081840152601f19601f8201169050808301925050509350505050604051602081830303815290604052805190602001208152602001908152602001600020541415610e2a5760036000815480929190600101919050555060035460026000858560405160200180806020018281038252848482818152602001925080828437600081840152601f19601f820116905080830192505050935050505060405160208183030381529060405280519060200120815260200190815260200160002081905550600060018054809190600101610b0791906124df565b90508383600183815481101515610b1a57fe5b90600052602060002090600b02016000019190610b38929190612511565b5060018082815481101515610b4957fe5b90600052602060002090600b02016001018190555060018082815481101515610b6e57fe5b90600052602060002090600b0201600201819055506000600182815481101515610b9457fe5b90600052602060002090600b02016003018190555060405180602001604052806000815250600182815481101515610bc857fe5b90600052602060002090600b02016004016000019080519060200190610bef929190612591565b5060405180602001604052806000815250600182815481101515610c0f57fe5b90600052602060002090600b02016004016001019080519060200190610c36929190612591565b506000600182815481101515610c4857fe5b90600052602060002090600b020160040160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600182815481101515610cab57fe5b90600052602060002090600b020160040160030181905550600181815481101515610cd257fe5b90600052602060002090600b020160010154600182815481101515610cf357fe5b90600052602060002090600b020160090160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600181815481101515610d5457fe5b90600052602060002090600b020160080160405180604001604052808473ffffffffffffffffffffffffffffffffffffffff1681526020016001151581525090806001815401808255809150509060018203906000526020600020016000909192909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160000160146101000a81548160ff02191690831515021790555050505050611243565b6000610e7984848080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050506122a6565b90506000600182815481101515610e8c57fe5b90600052602060002090600b020160090160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205414156110a757600181815481101515610ef057fe5b90600052602060002090600b020160010160008154809291906001019190505550600181815481101515610f2057fe5b90600052602060002090600b020160010154600182815481101515610f4157fe5b90600052602060002090600b020160090160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600181815481101515610fa257fe5b90600052602060002090600b020160080160405180604001604052808473ffffffffffffffffffffffffffffffffffffffff1681526020016001151581525090806001815401808255809150509060018203906000526020600020016000909192909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160000160146101000a81548160ff02191690831515021790555050505060018181548110151561108157fe5b90600052602060002090600b020160020160008154809291906001019190505550611241565b60006110f785858080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505084612351565b90506001151560018381548110151561110c57fe5b90600052602060002090600b02016008018281548110151561112a57fe5b9060005260206000200160000160149054906101000a900460ff161515141515156111bd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600f8152602001807f616c7265616479206120766f746572000000000000000000000000000000000081525060200191505060405180910390fd5b600180838154811015156111cd57fe5b90600052602060002090600b0201600801828154811015156111eb57fe5b9060005260206000200160000160146101000a81548160ff02191690831515021790555060018281548110151561121e57fe5b90600052602060002090600b020160020160008154809291906001019190505550505b505b7f424f3ad05c61ea35cad66f22b70b1fad7250d8229921238078c401db36d3457483838360405180806020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281038252858582818152602001925080828437600081840152601f19601f82011690508083019250505094505050505060405180910390a1505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630e32cf906040518163ffffffff1660e01b815260040160206040518083038186803b15801561134657600080fd5b505afa15801561135a573d6000803e3d6000fd5b505050506040513d602081101561137057600080fd5b810190808051906020019092919050505073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515611423576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600e8152602001807f696e76616c69642063616c6c657200000000000000000000000000000000000081525060200191505060405180910390fd5b82828080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050816001151561147783836123c8565b15151415156114ee576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600f8152602001807f6d757374206265206120766f746572000000000000000000000000000000000081525060200191505060405180910390fd5b600061153d86868080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050506122a6565b9050600061158f87878080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505086612351565b90506001828154811015156115a057fe5b90600052602060002090600b0201600201600081548092919060019003919050555060006001838154811015156115d357fe5b90600052602060002090600b0201600801828154811015156115f157fe5b9060005260206000200160000160146101000a81548160ff0219169083151502179055507f654cd85d9b2abaf3affef0a047625d088e6e4d0448935c9b5016b5f5aa0ca3b687878760405180806020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281038252858582818152602001925080828437600081840152601f19601f82011690508083019250505094505050505060405180910390a150505050505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630e32cf906040518163ffffffff1660e01b815260040160206040518083038186803b15801561171e57600080fd5b505afa158015611732573d6000803e3d6000fd5b505050506040513d602081101561174857600080fd5b810190808051906020019092919050505073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156117fb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600e8152602001807f696e76616c69642063616c6c657200000000000000000000000000000000000081525060200191505060405180910390fd5b84848080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050836001151561184f83836123c8565b15151415156118c6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600f8152602001807f6d757374206265206120766f746572000000000000000000000000000000000081525060200191505060405180910390fd5b6001151561191888888080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050866124a7565b151514151561198f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f6e6f7468696e6720746f20617070726f7665000000000000000000000000000081525060200191505060405180910390fd5b60006119de88888080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050506122a6565b9050600115156001828154811015156119f357fe5b90600052602060002090600b0201600a01600083815260200190815260200160002060008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151514151515611ad8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f63616e6e6f7420646f75626c6520766f7465000000000000000000000000000081525060200191505060405180910390fd5b600181815481101515611ae757fe5b90600052602060002090600b02016003016000815480929190600101919050555060018082815481101515611b1857fe5b90600052602060002090600b0201600a01600083815260200190815260200160002060008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055507f87999b54e45aa02834a1265e356d7bcdceb72b8cbb4396ebaeba32a103b43508888860405180806020018281038252848482818152602001925080828437600081840152601f19601f820116905080830192505050935050505060405180910390a16002600182815481101515611c0257fe5b90600052602060002090600b020160020154811515611c1d57fe5b04600182815481101515611c2d57fe5b90600052602060002090600b0201600301541115611d695760405180602001604052806000815250600182815481101515611c6457fe5b90600052602060002090600b02016004016000019080519060200190611c8b929190612591565b5060405180602001604052806000815250600182815481101515611cab57fe5b90600052602060002090600b02016004016001019080519060200190611cd2929190612591565b506000600182815481101515611ce457fe5b90600052602060002090600b020160040160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600182815481101515611d4757fe5b90600052602060002090600b0201600401600301819055506001935050611d6f565b60009350505b5050949350505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630e32cf906040518163ffffffff1660e01b815260040160206040518083038186803b158015611de057600080fd5b505afa158015611df4573d6000803e3d6000fd5b505050506040513d6020811015611e0a57600080fd5b810190808051906020019092919050505073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515611ebd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600e8152602001807f696e76616c69642063616c6c657200000000000000000000000000000000000081525060200191505060405180910390fd5b611f0c88888080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505060006124a7565b1515611f63576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260348152602001806127a76034913960400191505060405180910390fd5b6000611fb289898080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050506122a6565b90508686600183815481101515611fc557fe5b90600052602060002090600b02016004016000019190611fe6929190612511565b508484600183815481101515611ff857fe5b90600052602060002090600b02016004016001019190612019929190612511565b508260018281548110151561202a57fe5b90600052602060002090600b020160040160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508160018281548110151561208c57fe5b90600052602060002090600b02016004016003018190555060008090505b6001828154811015156120b957fe5b90600052602060002090600b020160080180549050811015612211576001828154811015156120e457fe5b90600052602060002090600b02016008018181548110151561210257fe5b9060005260206000200160000160149054906101000a900460ff161561220457600060018381548110151561213357fe5b90600052602060002090600b0201600a016000848152602001908152602001600020600060018581548110151561216657fe5b90600052602060002090600b02016008018481548110151561218457fe5b9060005260206000200160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505b80806001019150506120aa565b50600060018281548110151561222357fe5b90600052602060002090600b0201600301819055507f5bfaebb5931145594f63236d2a59314c4dc6035b65d0ca4cee9c7298e2f06ca3898960405180806020018281038252848482818152602001925080828437600081840152601f19601f820116905080830192505050935050505060405180910390a1505050505050505050565b6000600160026000846040516020018080602001828103825283818151815260200191508051906020019080838360005b838110156122f25780820151818401526020810190506122d7565b50505050905090810190601f16801561231f5780820380516001836020036101000a031916815260200191505b509250505060405160208183030381529060405280519060200120815260200190815260200160002054039050919050565b60008061235d846122a6565b90506001808281548110151561236f57fe5b90600052602060002090600b020160090160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020540391505092915050565b6000806123d4846122a6565b905060006001828154811015156123e757fe5b90600052602060002090600b020160090160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205414156124465760009150506124a1565b60006124528585612351565b905060018281548110151561246357fe5b90600052602060002090600b02016008018181548110151561248157fe5b9060005260206000200160000160149054906101000a900460ff16925050505b92915050565b60008160016124b5856122a6565b8154811015156124c157fe5b90600052602060002090600b02016004016003015414905092915050565b81548183558181111561250c57600b0281600b02836000526020600020918201910161250b9190612611565b5b505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061255257803560ff1916838001178555612580565b82800160010185558215612580579182015b8281111561257f578235825591602001919060010190612564565b5b50905061258d91906126be565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106125d257805160ff1916838001178555612600565b82800160010185558215612600579182015b828111156125ff5782518255916020019190600101906125e4565b5b50905061260d91906126be565b5090565b6126bb91905b808211156126b7576000808201600061263091906126e3565b600182016000905560028201600090556003820160009055600482016000808201600061265d91906126e3565b60018201600061266d91906126e3565b6002820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055600382016000905550506008820160006126ae919061272b565b50600b01612617565b5090565b90565b6126e091905b808211156126dc5760008160009055506001016126c4565b5090565b90565b50805460018160011615610100020316600290046000825580601f106127095750612728565b601f01602090049060005260206000209081019061272791906126be565b5b50565b5080546000825590600052602060002090810190612749919061274c565b50565b6127a391905b8082111561279f57600080820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556000820160146101000a81549060ff021916905550600101612752565b5090565b9056fe6974656d732070656e64696e6720666f7220617070726f76616c2e206e6577206974656d2063616e6e6f74206265206164646564a165627a7a72305820713fbb6bfba50ab709d791050af07bfb4f0bcd7cd358685edfd1e8db1b1a93f30029"

// DeployVoterManager deploys a new Ethereum contract, binding an instance of VoterManager to it.
func DeployVoterManager(auth *bind.TransactOpts, backend bind.ContractBackend, _permUpgradable common.Address) (common.Address, *types.Transaction, *VoterManager, error) {
	parsed, err := abi.JSON(strings.NewReader(VoterManagerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VoterManagerBin), backend, _permUpgradable)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VoterManager{VoterManagerCaller: VoterManagerCaller{contract: contract}, VoterManagerTransactor: VoterManagerTransactor{contract: contract}, VoterManagerFilterer: VoterManagerFilterer{contract: contract}}, nil
}

// VoterManager is an auto generated Go binding around an Ethereum contract.
type VoterManager struct {
	VoterManagerCaller     // Read-only binding to the contract
	VoterManagerTransactor // Write-only binding to the contract
	VoterManagerFilterer   // Log filterer for contract events
}

// VoterManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type VoterManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoterManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VoterManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoterManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VoterManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoterManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VoterManagerSession struct {
	Contract     *VoterManager     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VoterManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VoterManagerCallerSession struct {
	Contract *VoterManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// VoterManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VoterManagerTransactorSession struct {
	Contract     *VoterManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// VoterManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type VoterManagerRaw struct {
	Contract *VoterManager // Generic contract binding to access the raw methods on
}

// VoterManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VoterManagerCallerRaw struct {
	Contract *VoterManagerCaller // Generic read-only contract binding to access the raw methods on
}

// VoterManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VoterManagerTransactorRaw struct {
	Contract *VoterManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVoterManager creates a new instance of VoterManager, bound to a specific deployed contract.
func NewVoterManager(address common.Address, backend bind.ContractBackend) (*VoterManager, error) {
	contract, err := bindVoterManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VoterManager{VoterManagerCaller: VoterManagerCaller{contract: contract}, VoterManagerTransactor: VoterManagerTransactor{contract: contract}, VoterManagerFilterer: VoterManagerFilterer{contract: contract}}, nil
}

// NewVoterManagerCaller creates a new read-only instance of VoterManager, bound to a specific deployed contract.
func NewVoterManagerCaller(address common.Address, caller bind.ContractCaller) (*VoterManagerCaller, error) {
	contract, err := bindVoterManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VoterManagerCaller{contract: contract}, nil
}

// NewVoterManagerTransactor creates a new write-only instance of VoterManager, bound to a specific deployed contract.
func NewVoterManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*VoterManagerTransactor, error) {
	contract, err := bindVoterManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VoterManagerTransactor{contract: contract}, nil
}

// NewVoterManagerFilterer creates a new log filterer instance of VoterManager, bound to a specific deployed contract.
func NewVoterManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*VoterManagerFilterer, error) {
	contract, err := bindVoterManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VoterManagerFilterer{contract: contract}, nil
}

// bindVoterManager binds a generic wrapper to an already deployed contract.
func bindVoterManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VoterManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VoterManager *VoterManagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VoterManager.Contract.VoterManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VoterManager *VoterManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VoterManager.Contract.VoterManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VoterManager *VoterManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VoterManager.Contract.VoterManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VoterManager *VoterManagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VoterManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VoterManager *VoterManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VoterManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VoterManager *VoterManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VoterManager.Contract.contract.Transact(opts, method, params...)
}

// GetPendingOpDetails is a free data retrieval call binding the contract method 0x014e6acc.
//
// Solidity: function getPendingOpDetails(string _orgId) constant returns(string, string, address, uint256)
func (_VoterManager *VoterManagerCaller) GetPendingOpDetails(opts *bind.CallOpts, _orgId string) (string, string, common.Address, *big.Int, error) {
	var (
		ret0 = new(string)
		ret1 = new(string)
		ret2 = new(common.Address)
		ret3 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
	}
	err := _VoterManager.contract.Call(opts, out, "getPendingOpDetails", _orgId)
	return *ret0, *ret1, *ret2, *ret3, err
}

// GetPendingOpDetails is a free data retrieval call binding the contract method 0x014e6acc.
//
// Solidity: function getPendingOpDetails(string _orgId) constant returns(string, string, address, uint256)
func (_VoterManager *VoterManagerSession) GetPendingOpDetails(_orgId string) (string, string, common.Address, *big.Int, error) {
	return _VoterManager.Contract.GetPendingOpDetails(&_VoterManager.CallOpts, _orgId)
}

// GetPendingOpDetails is a free data retrieval call binding the contract method 0x014e6acc.
//
// Solidity: function getPendingOpDetails(string _orgId) constant returns(string, string, address, uint256)
func (_VoterManager *VoterManagerCallerSession) GetPendingOpDetails(_orgId string) (string, string, common.Address, *big.Int, error) {
	return _VoterManager.Contract.GetPendingOpDetails(&_VoterManager.CallOpts, _orgId)
}

// AddVoter is a paid mutator transaction binding the contract method 0x5607395b.
//
// Solidity: function addVoter(string _orgId, address _vAccount) returns()
func (_VoterManager *VoterManagerTransactor) AddVoter(opts *bind.TransactOpts, _orgId string, _vAccount common.Address) (*types.Transaction, error) {
	return _VoterManager.contract.Transact(opts, "addVoter", _orgId, _vAccount)
}

// AddVoter is a paid mutator transaction binding the contract method 0x5607395b.
//
// Solidity: function addVoter(string _orgId, address _vAccount) returns()
func (_VoterManager *VoterManagerSession) AddVoter(_orgId string, _vAccount common.Address) (*types.Transaction, error) {
	return _VoterManager.Contract.AddVoter(&_VoterManager.TransactOpts, _orgId, _vAccount)
}

// AddVoter is a paid mutator transaction binding the contract method 0x5607395b.
//
// Solidity: function addVoter(string _orgId, address _vAccount) returns()
func (_VoterManager *VoterManagerTransactorSession) AddVoter(_orgId string, _vAccount common.Address) (*types.Transaction, error) {
	return _VoterManager.Contract.AddVoter(&_VoterManager.TransactOpts, _orgId, _vAccount)
}

// AddVotingItem is a paid mutator transaction binding the contract method 0xe98ac22d.
//
// Solidity: function addVotingItem(string _authOrg, string _orgId, string _enodeId, address _account, uint256 _pendingOp) returns()
func (_VoterManager *VoterManagerTransactor) AddVotingItem(opts *bind.TransactOpts, _authOrg string, _orgId string, _enodeId string, _account common.Address, _pendingOp *big.Int) (*types.Transaction, error) {
	return _VoterManager.contract.Transact(opts, "addVotingItem", _authOrg, _orgId, _enodeId, _account, _pendingOp)
}

// AddVotingItem is a paid mutator transaction binding the contract method 0xe98ac22d.
//
// Solidity: function addVotingItem(string _authOrg, string _orgId, string _enodeId, address _account, uint256 _pendingOp) returns()
func (_VoterManager *VoterManagerSession) AddVotingItem(_authOrg string, _orgId string, _enodeId string, _account common.Address, _pendingOp *big.Int) (*types.Transaction, error) {
	return _VoterManager.Contract.AddVotingItem(&_VoterManager.TransactOpts, _authOrg, _orgId, _enodeId, _account, _pendingOp)
}

// AddVotingItem is a paid mutator transaction binding the contract method 0xe98ac22d.
//
// Solidity: function addVotingItem(string _authOrg, string _orgId, string _enodeId, address _account, uint256 _pendingOp) returns()
func (_VoterManager *VoterManagerTransactorSession) AddVotingItem(_authOrg string, _orgId string, _enodeId string, _account common.Address, _pendingOp *big.Int) (*types.Transaction, error) {
	return _VoterManager.Contract.AddVotingItem(&_VoterManager.TransactOpts, _authOrg, _orgId, _enodeId, _account, _pendingOp)
}

// DeleteVoter is a paid mutator transaction binding the contract method 0x59cbd6fe.
//
// Solidity: function deleteVoter(string _orgId, address _vAccount) returns()
func (_VoterManager *VoterManagerTransactor) DeleteVoter(opts *bind.TransactOpts, _orgId string, _vAccount common.Address) (*types.Transaction, error) {
	return _VoterManager.contract.Transact(opts, "deleteVoter", _orgId, _vAccount)
}

// DeleteVoter is a paid mutator transaction binding the contract method 0x59cbd6fe.
//
// Solidity: function deleteVoter(string _orgId, address _vAccount) returns()
func (_VoterManager *VoterManagerSession) DeleteVoter(_orgId string, _vAccount common.Address) (*types.Transaction, error) {
	return _VoterManager.Contract.DeleteVoter(&_VoterManager.TransactOpts, _orgId, _vAccount)
}

// DeleteVoter is a paid mutator transaction binding the contract method 0x59cbd6fe.
//
// Solidity: function deleteVoter(string _orgId, address _vAccount) returns()
func (_VoterManager *VoterManagerTransactorSession) DeleteVoter(_orgId string, _vAccount common.Address) (*types.Transaction, error) {
	return _VoterManager.Contract.DeleteVoter(&_VoterManager.TransactOpts, _orgId, _vAccount)
}

// ProcessVote is a paid mutator transaction binding the contract method 0xb0213864.
//
// Solidity: function processVote(string _authOrg, address _vAccount, uint256 _pendingOp) returns(bool)
func (_VoterManager *VoterManagerTransactor) ProcessVote(opts *bind.TransactOpts, _authOrg string, _vAccount common.Address, _pendingOp *big.Int) (*types.Transaction, error) {
	return _VoterManager.contract.Transact(opts, "processVote", _authOrg, _vAccount, _pendingOp)
}

// ProcessVote is a paid mutator transaction binding the contract method 0xb0213864.
//
// Solidity: function processVote(string _authOrg, address _vAccount, uint256 _pendingOp) returns(bool)
func (_VoterManager *VoterManagerSession) ProcessVote(_authOrg string, _vAccount common.Address, _pendingOp *big.Int) (*types.Transaction, error) {
	return _VoterManager.Contract.ProcessVote(&_VoterManager.TransactOpts, _authOrg, _vAccount, _pendingOp)
}

// ProcessVote is a paid mutator transaction binding the contract method 0xb0213864.
//
// Solidity: function processVote(string _authOrg, address _vAccount, uint256 _pendingOp) returns(bool)
func (_VoterManager *VoterManagerTransactorSession) ProcessVote(_authOrg string, _vAccount common.Address, _pendingOp *big.Int) (*types.Transaction, error) {
	return _VoterManager.Contract.ProcessVote(&_VoterManager.TransactOpts, _authOrg, _vAccount, _pendingOp)
}

// VoterManagerVoteProcessedIterator is returned from FilterVoteProcessed and is used to iterate over the raw logs and unpacked data for VoteProcessed events raised by the VoterManager contract.
type VoterManagerVoteProcessedIterator struct {
	Event *VoterManagerVoteProcessed // Event containing the contract specifics and raw log

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
func (it *VoterManagerVoteProcessedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoterManagerVoteProcessed)
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
		it.Event = new(VoterManagerVoteProcessed)
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
func (it *VoterManagerVoteProcessedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoterManagerVoteProcessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoterManagerVoteProcessed represents a VoteProcessed event raised by the VoterManager contract.
type VoterManagerVoteProcessed struct {
	OrgId string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterVoteProcessed is a free log retrieval operation binding the contract event 0x87999b54e45aa02834a1265e356d7bcdceb72b8cbb4396ebaeba32a103b43508.
//
// Solidity: event VoteProcessed(string _orgId)
func (_VoterManager *VoterManagerFilterer) FilterVoteProcessed(opts *bind.FilterOpts) (*VoterManagerVoteProcessedIterator, error) {

	logs, sub, err := _VoterManager.contract.FilterLogs(opts, "VoteProcessed")
	if err != nil {
		return nil, err
	}
	return &VoterManagerVoteProcessedIterator{contract: _VoterManager.contract, event: "VoteProcessed", logs: logs, sub: sub}, nil
}

var VoteProcessedTopicHash = "0x87999b54e45aa02834a1265e356d7bcdceb72b8cbb4396ebaeba32a103b43508"

// WatchVoteProcessed is a free log subscription operation binding the contract event 0x87999b54e45aa02834a1265e356d7bcdceb72b8cbb4396ebaeba32a103b43508.
//
// Solidity: event VoteProcessed(string _orgId)
func (_VoterManager *VoterManagerFilterer) WatchVoteProcessed(opts *bind.WatchOpts, sink chan<- *VoterManagerVoteProcessed) (event.Subscription, error) {

	logs, sub, err := _VoterManager.contract.WatchLogs(opts, "VoteProcessed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoterManagerVoteProcessed)
				if err := _VoterManager.contract.UnpackLog(event, "VoteProcessed", log); err != nil {
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

// ParseVoteProcessed is a log parse operation binding the contract event 0x87999b54e45aa02834a1265e356d7bcdceb72b8cbb4396ebaeba32a103b43508.
//
// Solidity: event VoteProcessed(string _orgId)
func (_VoterManager *VoterManagerFilterer) ParseVoteProcessed(log types.Log) (*VoterManagerVoteProcessed, error) {
	event := new(VoterManagerVoteProcessed)
	if err := _VoterManager.contract.UnpackLog(event, "VoteProcessed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// VoterManagerVoterAddedIterator is returned from FilterVoterAdded and is used to iterate over the raw logs and unpacked data for VoterAdded events raised by the VoterManager contract.
type VoterManagerVoterAddedIterator struct {
	Event *VoterManagerVoterAdded // Event containing the contract specifics and raw log

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
func (it *VoterManagerVoterAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoterManagerVoterAdded)
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
		it.Event = new(VoterManagerVoterAdded)
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
func (it *VoterManagerVoterAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoterManagerVoterAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoterManagerVoterAdded represents a VoterAdded event raised by the VoterManager contract.
type VoterManagerVoterAdded struct {
	OrgId    string
	VAccount common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterVoterAdded is a free log retrieval operation binding the contract event 0x424f3ad05c61ea35cad66f22b70b1fad7250d8229921238078c401db36d34574.
//
// Solidity: event VoterAdded(string _orgId, address _vAccount)
func (_VoterManager *VoterManagerFilterer) FilterVoterAdded(opts *bind.FilterOpts) (*VoterManagerVoterAddedIterator, error) {

	logs, sub, err := _VoterManager.contract.FilterLogs(opts, "VoterAdded")
	if err != nil {
		return nil, err
	}
	return &VoterManagerVoterAddedIterator{contract: _VoterManager.contract, event: "VoterAdded", logs: logs, sub: sub}, nil
}

var VoterAddedTopicHash = "0x424f3ad05c61ea35cad66f22b70b1fad7250d8229921238078c401db36d34574"

// WatchVoterAdded is a free log subscription operation binding the contract event 0x424f3ad05c61ea35cad66f22b70b1fad7250d8229921238078c401db36d34574.
//
// Solidity: event VoterAdded(string _orgId, address _vAccount)
func (_VoterManager *VoterManagerFilterer) WatchVoterAdded(opts *bind.WatchOpts, sink chan<- *VoterManagerVoterAdded) (event.Subscription, error) {

	logs, sub, err := _VoterManager.contract.WatchLogs(opts, "VoterAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoterManagerVoterAdded)
				if err := _VoterManager.contract.UnpackLog(event, "VoterAdded", log); err != nil {
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

// ParseVoterAdded is a log parse operation binding the contract event 0x424f3ad05c61ea35cad66f22b70b1fad7250d8229921238078c401db36d34574.
//
// Solidity: event VoterAdded(string _orgId, address _vAccount)
func (_VoterManager *VoterManagerFilterer) ParseVoterAdded(log types.Log) (*VoterManagerVoterAdded, error) {
	event := new(VoterManagerVoterAdded)
	if err := _VoterManager.contract.UnpackLog(event, "VoterAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// VoterManagerVoterDeletedIterator is returned from FilterVoterDeleted and is used to iterate over the raw logs and unpacked data for VoterDeleted events raised by the VoterManager contract.
type VoterManagerVoterDeletedIterator struct {
	Event *VoterManagerVoterDeleted // Event containing the contract specifics and raw log

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
func (it *VoterManagerVoterDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoterManagerVoterDeleted)
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
		it.Event = new(VoterManagerVoterDeleted)
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
func (it *VoterManagerVoterDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoterManagerVoterDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoterManagerVoterDeleted represents a VoterDeleted event raised by the VoterManager contract.
type VoterManagerVoterDeleted struct {
	OrgId    string
	VAccount common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterVoterDeleted is a free log retrieval operation binding the contract event 0x654cd85d9b2abaf3affef0a047625d088e6e4d0448935c9b5016b5f5aa0ca3b6.
//
// Solidity: event VoterDeleted(string _orgId, address _vAccount)
func (_VoterManager *VoterManagerFilterer) FilterVoterDeleted(opts *bind.FilterOpts) (*VoterManagerVoterDeletedIterator, error) {

	logs, sub, err := _VoterManager.contract.FilterLogs(opts, "VoterDeleted")
	if err != nil {
		return nil, err
	}
	return &VoterManagerVoterDeletedIterator{contract: _VoterManager.contract, event: "VoterDeleted", logs: logs, sub: sub}, nil
}

var VoterDeletedTopicHash = "0x654cd85d9b2abaf3affef0a047625d088e6e4d0448935c9b5016b5f5aa0ca3b6"

// WatchVoterDeleted is a free log subscription operation binding the contract event 0x654cd85d9b2abaf3affef0a047625d088e6e4d0448935c9b5016b5f5aa0ca3b6.
//
// Solidity: event VoterDeleted(string _orgId, address _vAccount)
func (_VoterManager *VoterManagerFilterer) WatchVoterDeleted(opts *bind.WatchOpts, sink chan<- *VoterManagerVoterDeleted) (event.Subscription, error) {

	logs, sub, err := _VoterManager.contract.WatchLogs(opts, "VoterDeleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoterManagerVoterDeleted)
				if err := _VoterManager.contract.UnpackLog(event, "VoterDeleted", log); err != nil {
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

// ParseVoterDeleted is a log parse operation binding the contract event 0x654cd85d9b2abaf3affef0a047625d088e6e4d0448935c9b5016b5f5aa0ca3b6.
//
// Solidity: event VoterDeleted(string _orgId, address _vAccount)
func (_VoterManager *VoterManagerFilterer) ParseVoterDeleted(log types.Log) (*VoterManagerVoterDeleted, error) {
	event := new(VoterManagerVoterDeleted)
	if err := _VoterManager.contract.UnpackLog(event, "VoterDeleted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// VoterManagerVotingItemAddedIterator is returned from FilterVotingItemAdded and is used to iterate over the raw logs and unpacked data for VotingItemAdded events raised by the VoterManager contract.
type VoterManagerVotingItemAddedIterator struct {
	Event *VoterManagerVotingItemAdded // Event containing the contract specifics and raw log

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
func (it *VoterManagerVotingItemAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoterManagerVotingItemAdded)
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
		it.Event = new(VoterManagerVotingItemAdded)
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
func (it *VoterManagerVotingItemAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoterManagerVotingItemAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoterManagerVotingItemAdded represents a VotingItemAdded event raised by the VoterManager contract.
type VoterManagerVotingItemAdded struct {
	OrgId string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterVotingItemAdded is a free log retrieval operation binding the contract event 0x5bfaebb5931145594f63236d2a59314c4dc6035b65d0ca4cee9c7298e2f06ca3.
//
// Solidity: event VotingItemAdded(string _orgId)
func (_VoterManager *VoterManagerFilterer) FilterVotingItemAdded(opts *bind.FilterOpts) (*VoterManagerVotingItemAddedIterator, error) {

	logs, sub, err := _VoterManager.contract.FilterLogs(opts, "VotingItemAdded")
	if err != nil {
		return nil, err
	}
	return &VoterManagerVotingItemAddedIterator{contract: _VoterManager.contract, event: "VotingItemAdded", logs: logs, sub: sub}, nil
}

var VotingItemAddedTopicHash = "0x5bfaebb5931145594f63236d2a59314c4dc6035b65d0ca4cee9c7298e2f06ca3"

// WatchVotingItemAdded is a free log subscription operation binding the contract event 0x5bfaebb5931145594f63236d2a59314c4dc6035b65d0ca4cee9c7298e2f06ca3.
//
// Solidity: event VotingItemAdded(string _orgId)
func (_VoterManager *VoterManagerFilterer) WatchVotingItemAdded(opts *bind.WatchOpts, sink chan<- *VoterManagerVotingItemAdded) (event.Subscription, error) {

	logs, sub, err := _VoterManager.contract.WatchLogs(opts, "VotingItemAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoterManagerVotingItemAdded)
				if err := _VoterManager.contract.UnpackLog(event, "VotingItemAdded", log); err != nil {
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

// ParseVotingItemAdded is a log parse operation binding the contract event 0x5bfaebb5931145594f63236d2a59314c4dc6035b65d0ca4cee9c7298e2f06ca3.
//
// Solidity: event VotingItemAdded(string _orgId)
func (_VoterManager *VoterManagerFilterer) ParseVotingItemAdded(log types.Log) (*VoterManagerVotingItemAdded, error) {
	event := new(VoterManagerVotingItemAdded)
	if err := _VoterManager.contract.UnpackLog(event, "VotingItemAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}
