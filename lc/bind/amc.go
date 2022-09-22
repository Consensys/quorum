// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bind

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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// AMCABI is the input ABI used to generate the binding from.
const AMCABI = "[{\"inputs\":[{\"internalType\":\"contractIPermissionsInterface\",\"name\":\"_permission\",\"type\":\"address\"},{\"internalType\":\"contractIMode\",\"name\":\"_mode\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_org\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"amendRequest\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isAuthorized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isManager\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isNetworkAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"managementOrg\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mode\",\"outputs\":[{\"internalType\":\"contractIMode\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"permission\",\"outputs\":[{\"internalType\":\"contractIPermissionsInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"router\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_amendRequest\",\"type\":\"address\"}],\"name\":\"setAmendRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_newOrg\",\"type\":\"string\"}],\"name\":\"setManagementOrg\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_mode\",\"type\":\"address\"}],\"name\":\"setMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_permission\",\"type\":\"address\"}],\"name\":\"setPermission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"}],\"name\":\"setRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_standardFactory\",\"type\":\"address\"}],\"name\":\"setStandardFactory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_upasFactory\",\"type\":\"address\"}],\"name\":\"setUPASFactory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"standardFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upasFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_org\",\"type\":\"bytes32\"}],\"name\":\"verifyIdentity\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

var AMCParsedABI, _ = abi.JSON(strings.NewReader(AMCABI))

// AMCBin is the compiled bytecode used for deploying new contracts.
var AMCBin = "0x60806040523480156200001157600080fd5b5060405162001ea038038062001ea0833981810160405281019062000037919062000236565b80600690805190602001906200004f929190620000da565b5081600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550826000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050620004c5565b828054620000e890620003a2565b90600052602060002090601f0160209004810192826200010c576000855562000158565b82601f106200012757805160ff191683800117855562000158565b8280016001018555821562000158579182015b82811115620001575782518255916020019190600101906200013a565b5b5090506200016791906200016b565b5090565b5b80821115620001865760008160009055506001016200016c565b5090565b6000620001a16200019b84620002da565b620002b1565b905082815260208101848484011115620001c057620001bf62000471565b5b620001cd8482856200036c565b509392505050565b600081519050620001e68162000491565b92915050565b600081519050620001fd81620004ab565b92915050565b600082601f8301126200021b576200021a6200046c565b5b81516200022d8482602086016200018a565b91505092915050565b6000806000606084860312156200025257620002516200047b565b5b60006200026286828701620001ec565b93505060206200027586828701620001d5565b925050604084015167ffffffffffffffff81111562000299576200029862000476565b5b620002a78682870162000203565b9150509250925092565b6000620002bd620002d0565b9050620002cb8282620003d8565b919050565b6000604051905090565b600067ffffffffffffffff821115620002f857620002f76200043d565b5b620003038262000480565b9050602081019050919050565b60006200031d826200034c565b9050919050565b6000620003318262000310565b9050919050565b6000620003458262000310565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60005b838110156200038c5780820151818401526020810190506200036f565b838111156200039c576000848401525b50505050565b60006002820490506001821680620003bb57607f821691505b60208210811415620003d257620003d16200040e565b5b50919050565b620003e38262000480565b810181811067ffffffffffffffff821117156200040557620004046200043d565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b6200049c8162000324565b8114620004a857600080fd5b50565b620004b68162000338565b8114620004c257600080fd5b50565b6119cb80620004d56000396000f3fe608060405234801561001057600080fd5b50600436106101155760003560e01c806386db5f8f116100a2578063d1aa0c2011610071578063d1aa0c20146102a4578063f3ae2415146102d4578063f3b0c8b714610304578063f887ea4014610322578063fe9fbb801461034057610115565b806386db5f8f146102325780639e694cea14610250578063b85a35d21461026c578063c0d786551461028857610115565b80635581f372116100e95780635581f3721461018e5780635f8501c8146101be57806366fba795146101dc57806371e771a2146101fa578063829297921461021657610115565b80625fa9391461011a578063295a521214610136578063317f86381461015457806339cd8e9614610172575b600080fd5b610134600480360381019061012f9190611209565b610370565b005b61013e61046e565b60405161014b91906114e4565b60405180910390f35b61015c610494565b604051610169919061144e565b60405180910390f35b61018c60048036038101906101879190611209565b6104ba565b005b6101a860048036038101906101a39190611236565b6105b8565b6040516101b591906114c9565b60405180910390f35b6101c6610679565b6040516101d3919061144e565b60405180910390f35b6101e461069f565b6040516101f1919061144e565b60405180910390f35b610214600480360381019061020f9190611209565b6106c5565b005b610230600480360381019061022b91906112a3565b6107c3565b005b61023a610821565b604051610247919061151a565b60405180910390f35b61026a60048036038101906102659190611209565b6108af565b005b61028660048036038101906102819190611209565b6109ad565b005b6102a2600480360381019061029d9190611209565b610aaa565b005b6102be60048036038101906102b99190611209565b610ba8565b6040516102cb91906114c9565b60405180910390f35b6102ee60048036038101906102e99190611209565b610c5b565b6040516102fb91906114c9565b60405180910390f35b61030c610d11565b60405161031991906114ff565b60405180910390f35b61032a610d35565b604051610337919061144e565b60405180910390f35b61035a60048036038101906103559190611209565b610d5b565b60405161036791906114c9565b60405180910390f35b61037933610d5b565b6103b8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103af9061155c565b60405180910390fd5b80600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610429576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104209061157c565b60405180910390fd5b81600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6104c333610d5b565b610502576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104f99061155c565b60405180910390fd5b80600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610573576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161056a9061157c565b60405180910390fd5b81600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636b568d76846106048560001c610e0f565b6040518363ffffffff1660e01b8152600401610621929190611469565b60206040518083038186803b15801561063957600080fd5b505afa15801561064d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106719190611276565b905092915050565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6106ce33610d5b565b61070d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107049061155c565b60405180910390fd5b80600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141561077e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107759061157c565b60405180910390fd5b81600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b6107cc33610d5b565b61080b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108029061155c565b60405180910390fd5b81816006919061081c9291906110d1565b505050565b6006805461082e90611774565b80601f016020809104026020016040519081016040528092919081815260200182805461085a90611774565b80156108a75780601f1061087c576101008083540402835291602001916108a7565b820191906000526020600020905b81548152906001019060200180831161088a57829003601f168201915b505050505081565b6108b833610d5b565b6108f7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108ee9061155c565b60405180910390fd5b80600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610968576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161095f9061157c565b60405180910390fd5b81600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b6109b633610d5b565b6109f5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109ec9061155c565b60405180910390fd5b80600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610a66576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a5d9061157c565b60405180910390fd5b816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b610ab333610d5b565b610af2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ae99061155c565b60405180910390fd5b80600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610b63576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b5a9061157c565b60405180910390fd5b81600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d1aa0c20836040518263ffffffff1660e01b8152600401610c04919061144e565b60206040518083038186803b158015610c1c57600080fd5b505afa158015610c30573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c549190611276565b9050919050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636b568d768360066040518363ffffffff1660e01b8152600401610cba929190611499565b60206040518083038186803b158015610cd257600080fd5b505afa158015610ce6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d0a9190611276565b9050919050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166389f4dd47836040518263ffffffff1660e01b8152600401610db8919061144e565b60206040518083038186803b158015610dd057600080fd5b505afa158015610de4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e089190611276565b9050919050565b60606000821415610e57576040518060400160405280600481526020017f30783030000000000000000000000000000000000000000000000000000000008152509050610e90565b600082905060005b60008214610e81578080610e72906117a6565b915050600882901c9150610e5f565b610e8b8482610e95565b925050505b919050565b606060006002836002610ea89190611623565b610eb291906115cd565b67ffffffffffffffff811115610ecb57610eca61187c565b5b6040519080825280601f01601f191660200182016040528015610efd5781602001600182028036833780820191505090505b5090507f300000000000000000000000000000000000000000000000000000000000000081600081518110610f3557610f3461184d565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053507f780000000000000000000000000000000000000000000000000000000000000081600181518110610f9957610f9861184d565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535060006001846002610fd99190611623565b610fe391906115cd565b90505b6001811115611083577f3031323334353637383961626364656600000000000000000000000000000000600f8616601081106110255761102461184d565b5b1a60f81b82828151811061103c5761103b61184d565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600485901c94508061107c9061174a565b9050610fe6565b50600084146110c7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110be9061153c565b60405180910390fd5b8091505092915050565b8280546110dd90611774565b90600052602060002090601f0160209004810192826110ff5760008555611146565b82601f1061111857803560ff1916838001178555611146565b82800160010185558215611146579182015b8281111561114557823582559160200191906001019061112a565b5b5090506111539190611157565b5090565b5b80821115611170576000816000905550600101611158565b5090565b60008135905061118381611950565b92915050565b60008151905061119881611967565b92915050565b6000813590506111ad8161197e565b92915050565b60008083601f8401126111c9576111c86118b0565b5b8235905067ffffffffffffffff8111156111e6576111e56118ab565b5b602083019150836001820283011115611202576112016118b5565b5b9250929050565b60006020828403121561121f5761121e6118bf565b5b600061122d84828501611174565b91505092915050565b6000806040838503121561124d5761124c6118bf565b5b600061125b85828601611174565b925050602061126c8582860161119e565b9150509250929050565b60006020828403121561128c5761128b6118bf565b5b600061129a84828501611189565b91505092915050565b600080602083850312156112ba576112b96118bf565b5b600083013567ffffffffffffffff8111156112d8576112d76118ba565b5b6112e4858286016111b3565b92509250509250929050565b6112f98161167d565b82525050565b6113088161168f565b82525050565b611317816116cf565b82525050565b611326816116f3565b82525050565b6000611337826115b1565b61134181856115bc565b9350611351818560208601611717565b61135a816118c4565b840191505092915050565b6000815461137281611774565b61137c81866115bc565b9450600182166000811461139757600181146113a9576113dc565b60ff19831686526020860193506113dc565b6113b28561159c565b60005b838110156113d4578154818901526001820191506020810190506113b5565b808801955050505b50505092915050565b60006113f26020836115bc565b91506113fd826118d5565b602082019050919050565b6000611415600c836115bc565b9150611420826118fe565b602082019050919050565b60006114386010836115bc565b915061144382611927565b602082019050919050565b600060208201905061146360008301846112f0565b92915050565b600060408201905061147e60008301856112f0565b8181036020830152611490818461132c565b90509392505050565b60006040820190506114ae60008301856112f0565b81810360208301526114c08184611365565b90509392505050565b60006020820190506114de60008301846112ff565b92915050565b60006020820190506114f9600083018461130e565b92915050565b6000602082019050611514600083018461131d565b92915050565b60006020820190508181036000830152611534818461132c565b905092915050565b60006020820190508181036000830152611555816113e5565b9050919050565b6000602082019050818103600083015261157581611408565b9050919050565b600060208201905081810360008301526115958161142b565b9050919050565b60008190508160005260206000209050919050565b600081519050919050565b600082825260208201905092915050565b60006115d8826116c5565b91506115e3836116c5565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115611618576116176117ef565b5b828201905092915050565b600061162e826116c5565b9150611639836116c5565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615611672576116716117ef565b5b828202905092915050565b6000611688826116a5565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b60006116da826116e1565b9050919050565b60006116ec826116a5565b9050919050565b60006116fe82611705565b9050919050565b6000611710826116a5565b9050919050565b60005b8381101561173557808201518184015260208101905061171a565b83811115611744576000848401525b50505050565b6000611755826116c5565b91506000821415611769576117686117ef565b5b600182039050919050565b6000600282049050600182168061178c57607f821691505b602082108114156117a05761179f61181e565b5b50919050565b60006117b1826116c5565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8214156117e4576117e36117ef565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f537472696e67733a20686578206c656e67746820696e73756666696369656e74600082015250565b7f556e617574686f72697a65640000000000000000000000000000000000000000600082015250565b7f536574207a65726f206164647265737300000000000000000000000000000000600082015250565b6119598161167d565b811461196457600080fd5b50565b6119708161168f565b811461197b57600080fd5b50565b6119878161169b565b811461199257600080fd5b5056fea264697066735822122005a018f47a18fc669e30613f9507b0bca50234a6595427552c3ed10474078b2164736f6c63430008060033"

// DeployAMC deploys a new Ethereum contract, binding an instance of AMC to it.
func DeployAMC(auth *bind.TransactOpts, backend bind.ContractBackend, _permission common.Address, _mode common.Address, _org string) (common.Address, *types.Transaction, *AMC, error) {
	parsed, err := abi.JSON(strings.NewReader(AMCABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AMCBin), backend, _permission, _mode, _org)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AMC{AMCCaller: AMCCaller{contract: contract}, AMCTransactor: AMCTransactor{contract: contract}, AMCFilterer: AMCFilterer{contract: contract}}, nil
}

// AMC is an auto generated Go binding around an Ethereum contract.
type AMC struct {
	AMCCaller     // Read-only binding to the contract
	AMCTransactor // Write-only binding to the contract
	AMCFilterer   // Log filterer for contract events
}

// AMCCaller is an auto generated read-only Go binding around an Ethereum contract.
type AMCCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AMCTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AMCTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AMCFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AMCFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AMCSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AMCSession struct {
	Contract     *AMC              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AMCCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AMCCallerSession struct {
	Contract *AMCCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AMCTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AMCTransactorSession struct {
	Contract     *AMCTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AMCRaw is an auto generated low-level Go binding around an Ethereum contract.
type AMCRaw struct {
	Contract *AMC // Generic contract binding to access the raw methods on
}

// AMCCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AMCCallerRaw struct {
	Contract *AMCCaller // Generic read-only contract binding to access the raw methods on
}

// AMCTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AMCTransactorRaw struct {
	Contract *AMCTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAMC creates a new instance of AMC, bound to a specific deployed contract.
func NewAMC(address common.Address, backend bind.ContractBackend) (*AMC, error) {
	contract, err := bindAMC(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AMC{AMCCaller: AMCCaller{contract: contract}, AMCTransactor: AMCTransactor{contract: contract}, AMCFilterer: AMCFilterer{contract: contract}}, nil
}

// NewAMCCaller creates a new read-only instance of AMC, bound to a specific deployed contract.
func NewAMCCaller(address common.Address, caller bind.ContractCaller) (*AMCCaller, error) {
	contract, err := bindAMC(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AMCCaller{contract: contract}, nil
}

// NewAMCTransactor creates a new write-only instance of AMC, bound to a specific deployed contract.
func NewAMCTransactor(address common.Address, transactor bind.ContractTransactor) (*AMCTransactor, error) {
	contract, err := bindAMC(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AMCTransactor{contract: contract}, nil
}

// NewAMCFilterer creates a new log filterer instance of AMC, bound to a specific deployed contract.
func NewAMCFilterer(address common.Address, filterer bind.ContractFilterer) (*AMCFilterer, error) {
	contract, err := bindAMC(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AMCFilterer{contract: contract}, nil
}

// bindAMC binds a generic wrapper to an already deployed contract.
func bindAMC(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AMCABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AMC *AMCRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AMC.Contract.AMCCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AMC *AMCRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AMC.Contract.AMCTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AMC *AMCRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AMC.Contract.AMCTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AMC *AMCCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AMC.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AMC *AMCTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AMC.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AMC *AMCTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AMC.Contract.contract.Transact(opts, method, params...)
}

// AmendRequest is a free data retrieval call binding the contract method 0x66fba795.
//
// Solidity: function amendRequest() view returns(address)
func (_AMC *AMCCaller) AmendRequest(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AMC.contract.Call(opts, &out, "amendRequest")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AmendRequest is a free data retrieval call binding the contract method 0x66fba795.
//
// Solidity: function amendRequest() view returns(address)
func (_AMC *AMCSession) AmendRequest() (common.Address, error) {
	return _AMC.Contract.AmendRequest(&_AMC.CallOpts)
}

// AmendRequest is a free data retrieval call binding the contract method 0x66fba795.
//
// Solidity: function amendRequest() view returns(address)
func (_AMC *AMCCallerSession) AmendRequest() (common.Address, error) {
	return _AMC.Contract.AmendRequest(&_AMC.CallOpts)
}

// IsAuthorized is a free data retrieval call binding the contract method 0xfe9fbb80.
//
// Solidity: function isAuthorized(address _account) view returns(bool)
func (_AMC *AMCCaller) IsAuthorized(opts *bind.CallOpts, _account common.Address) (bool, error) {
	var out []interface{}
	err := _AMC.contract.Call(opts, &out, "isAuthorized", _account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAuthorized is a free data retrieval call binding the contract method 0xfe9fbb80.
//
// Solidity: function isAuthorized(address _account) view returns(bool)
func (_AMC *AMCSession) IsAuthorized(_account common.Address) (bool, error) {
	return _AMC.Contract.IsAuthorized(&_AMC.CallOpts, _account)
}

// IsAuthorized is a free data retrieval call binding the contract method 0xfe9fbb80.
//
// Solidity: function isAuthorized(address _account) view returns(bool)
func (_AMC *AMCCallerSession) IsAuthorized(_account common.Address) (bool, error) {
	return _AMC.Contract.IsAuthorized(&_AMC.CallOpts, _account)
}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address _account) view returns(bool)
func (_AMC *AMCCaller) IsManager(opts *bind.CallOpts, _account common.Address) (bool, error) {
	var out []interface{}
	err := _AMC.contract.Call(opts, &out, "isManager", _account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address _account) view returns(bool)
func (_AMC *AMCSession) IsManager(_account common.Address) (bool, error) {
	return _AMC.Contract.IsManager(&_AMC.CallOpts, _account)
}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address _account) view returns(bool)
func (_AMC *AMCCallerSession) IsManager(_account common.Address) (bool, error) {
	return _AMC.Contract.IsManager(&_AMC.CallOpts, _account)
}

// IsNetworkAdmin is a free data retrieval call binding the contract method 0xd1aa0c20.
//
// Solidity: function isNetworkAdmin(address _account) view returns(bool)
func (_AMC *AMCCaller) IsNetworkAdmin(opts *bind.CallOpts, _account common.Address) (bool, error) {
	var out []interface{}
	err := _AMC.contract.Call(opts, &out, "isNetworkAdmin", _account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsNetworkAdmin is a free data retrieval call binding the contract method 0xd1aa0c20.
//
// Solidity: function isNetworkAdmin(address _account) view returns(bool)
func (_AMC *AMCSession) IsNetworkAdmin(_account common.Address) (bool, error) {
	return _AMC.Contract.IsNetworkAdmin(&_AMC.CallOpts, _account)
}

// IsNetworkAdmin is a free data retrieval call binding the contract method 0xd1aa0c20.
//
// Solidity: function isNetworkAdmin(address _account) view returns(bool)
func (_AMC *AMCCallerSession) IsNetworkAdmin(_account common.Address) (bool, error) {
	return _AMC.Contract.IsNetworkAdmin(&_AMC.CallOpts, _account)
}

// ManagementOrg is a free data retrieval call binding the contract method 0x86db5f8f.
//
// Solidity: function managementOrg() view returns(string)
func (_AMC *AMCCaller) ManagementOrg(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AMC.contract.Call(opts, &out, "managementOrg")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ManagementOrg is a free data retrieval call binding the contract method 0x86db5f8f.
//
// Solidity: function managementOrg() view returns(string)
func (_AMC *AMCSession) ManagementOrg() (string, error) {
	return _AMC.Contract.ManagementOrg(&_AMC.CallOpts)
}

// ManagementOrg is a free data retrieval call binding the contract method 0x86db5f8f.
//
// Solidity: function managementOrg() view returns(string)
func (_AMC *AMCCallerSession) ManagementOrg() (string, error) {
	return _AMC.Contract.ManagementOrg(&_AMC.CallOpts)
}

// Mode is a free data retrieval call binding the contract method 0x295a5212.
//
// Solidity: function mode() view returns(address)
func (_AMC *AMCCaller) Mode(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AMC.contract.Call(opts, &out, "mode")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Mode is a free data retrieval call binding the contract method 0x295a5212.
//
// Solidity: function mode() view returns(address)
func (_AMC *AMCSession) Mode() (common.Address, error) {
	return _AMC.Contract.Mode(&_AMC.CallOpts)
}

// Mode is a free data retrieval call binding the contract method 0x295a5212.
//
// Solidity: function mode() view returns(address)
func (_AMC *AMCCallerSession) Mode() (common.Address, error) {
	return _AMC.Contract.Mode(&_AMC.CallOpts)
}

// Permission is a free data retrieval call binding the contract method 0xf3b0c8b7.
//
// Solidity: function permission() view returns(address)
func (_AMC *AMCCaller) Permission(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AMC.contract.Call(opts, &out, "permission")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Permission is a free data retrieval call binding the contract method 0xf3b0c8b7.
//
// Solidity: function permission() view returns(address)
func (_AMC *AMCSession) Permission() (common.Address, error) {
	return _AMC.Contract.Permission(&_AMC.CallOpts)
}

// Permission is a free data retrieval call binding the contract method 0xf3b0c8b7.
//
// Solidity: function permission() view returns(address)
func (_AMC *AMCCallerSession) Permission() (common.Address, error) {
	return _AMC.Contract.Permission(&_AMC.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_AMC *AMCCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AMC.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_AMC *AMCSession) Router() (common.Address, error) {
	return _AMC.Contract.Router(&_AMC.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_AMC *AMCCallerSession) Router() (common.Address, error) {
	return _AMC.Contract.Router(&_AMC.CallOpts)
}

// StandardFactory is a free data retrieval call binding the contract method 0x317f8638.
//
// Solidity: function standardFactory() view returns(address)
func (_AMC *AMCCaller) StandardFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AMC.contract.Call(opts, &out, "standardFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StandardFactory is a free data retrieval call binding the contract method 0x317f8638.
//
// Solidity: function standardFactory() view returns(address)
func (_AMC *AMCSession) StandardFactory() (common.Address, error) {
	return _AMC.Contract.StandardFactory(&_AMC.CallOpts)
}

// StandardFactory is a free data retrieval call binding the contract method 0x317f8638.
//
// Solidity: function standardFactory() view returns(address)
func (_AMC *AMCCallerSession) StandardFactory() (common.Address, error) {
	return _AMC.Contract.StandardFactory(&_AMC.CallOpts)
}

// UpasFactory is a free data retrieval call binding the contract method 0x5f8501c8.
//
// Solidity: function upasFactory() view returns(address)
func (_AMC *AMCCaller) UpasFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AMC.contract.Call(opts, &out, "upasFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UpasFactory is a free data retrieval call binding the contract method 0x5f8501c8.
//
// Solidity: function upasFactory() view returns(address)
func (_AMC *AMCSession) UpasFactory() (common.Address, error) {
	return _AMC.Contract.UpasFactory(&_AMC.CallOpts)
}

// UpasFactory is a free data retrieval call binding the contract method 0x5f8501c8.
//
// Solidity: function upasFactory() view returns(address)
func (_AMC *AMCCallerSession) UpasFactory() (common.Address, error) {
	return _AMC.Contract.UpasFactory(&_AMC.CallOpts)
}

// VerifyIdentity is a free data retrieval call binding the contract method 0x5581f372.
//
// Solidity: function verifyIdentity(address _account, bytes32 _org) view returns(bool)
func (_AMC *AMCCaller) VerifyIdentity(opts *bind.CallOpts, _account common.Address, _org [32]byte) (bool, error) {
	var out []interface{}
	err := _AMC.contract.Call(opts, &out, "verifyIdentity", _account, _org)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyIdentity is a free data retrieval call binding the contract method 0x5581f372.
//
// Solidity: function verifyIdentity(address _account, bytes32 _org) view returns(bool)
func (_AMC *AMCSession) VerifyIdentity(_account common.Address, _org [32]byte) (bool, error) {
	return _AMC.Contract.VerifyIdentity(&_AMC.CallOpts, _account, _org)
}

// VerifyIdentity is a free data retrieval call binding the contract method 0x5581f372.
//
// Solidity: function verifyIdentity(address _account, bytes32 _org) view returns(bool)
func (_AMC *AMCCallerSession) VerifyIdentity(_account common.Address, _org [32]byte) (bool, error) {
	return _AMC.Contract.VerifyIdentity(&_AMC.CallOpts, _account, _org)
}

// SetAmendRequest is a paid mutator transaction binding the contract method 0x71e771a2.
//
// Solidity: function setAmendRequest(address _amendRequest) returns()
func (_AMC *AMCTransactor) SetAmendRequest(opts *bind.TransactOpts, _amendRequest common.Address) (*types.Transaction, error) {
	return _AMC.contract.Transact(opts, "setAmendRequest", _amendRequest)
}

// SetAmendRequest is a paid mutator transaction binding the contract method 0x71e771a2.
//
// Solidity: function setAmendRequest(address _amendRequest) returns()
func (_AMC *AMCSession) SetAmendRequest(_amendRequest common.Address) (*types.Transaction, error) {
	return _AMC.Contract.SetAmendRequest(&_AMC.TransactOpts, _amendRequest)
}

// SetAmendRequest is a paid mutator transaction binding the contract method 0x71e771a2.
//
// Solidity: function setAmendRequest(address _amendRequest) returns()
func (_AMC *AMCTransactorSession) SetAmendRequest(_amendRequest common.Address) (*types.Transaction, error) {
	return _AMC.Contract.SetAmendRequest(&_AMC.TransactOpts, _amendRequest)
}

// SetManagementOrg is a paid mutator transaction binding the contract method 0x82929792.
//
// Solidity: function setManagementOrg(string _newOrg) returns()
func (_AMC *AMCTransactor) SetManagementOrg(opts *bind.TransactOpts, _newOrg string) (*types.Transaction, error) {
	return _AMC.contract.Transact(opts, "setManagementOrg", _newOrg)
}

// SetManagementOrg is a paid mutator transaction binding the contract method 0x82929792.
//
// Solidity: function setManagementOrg(string _newOrg) returns()
func (_AMC *AMCSession) SetManagementOrg(_newOrg string) (*types.Transaction, error) {
	return _AMC.Contract.SetManagementOrg(&_AMC.TransactOpts, _newOrg)
}

// SetManagementOrg is a paid mutator transaction binding the contract method 0x82929792.
//
// Solidity: function setManagementOrg(string _newOrg) returns()
func (_AMC *AMCTransactorSession) SetManagementOrg(_newOrg string) (*types.Transaction, error) {
	return _AMC.Contract.SetManagementOrg(&_AMC.TransactOpts, _newOrg)
}

// SetMode is a paid mutator transaction binding the contract method 0x9e694cea.
//
// Solidity: function setMode(address _mode) returns()
func (_AMC *AMCTransactor) SetMode(opts *bind.TransactOpts, _mode common.Address) (*types.Transaction, error) {
	return _AMC.contract.Transact(opts, "setMode", _mode)
}

// SetMode is a paid mutator transaction binding the contract method 0x9e694cea.
//
// Solidity: function setMode(address _mode) returns()
func (_AMC *AMCSession) SetMode(_mode common.Address) (*types.Transaction, error) {
	return _AMC.Contract.SetMode(&_AMC.TransactOpts, _mode)
}

// SetMode is a paid mutator transaction binding the contract method 0x9e694cea.
//
// Solidity: function setMode(address _mode) returns()
func (_AMC *AMCTransactorSession) SetMode(_mode common.Address) (*types.Transaction, error) {
	return _AMC.Contract.SetMode(&_AMC.TransactOpts, _mode)
}

// SetPermission is a paid mutator transaction binding the contract method 0xb85a35d2.
//
// Solidity: function setPermission(address _permission) returns()
func (_AMC *AMCTransactor) SetPermission(opts *bind.TransactOpts, _permission common.Address) (*types.Transaction, error) {
	return _AMC.contract.Transact(opts, "setPermission", _permission)
}

// SetPermission is a paid mutator transaction binding the contract method 0xb85a35d2.
//
// Solidity: function setPermission(address _permission) returns()
func (_AMC *AMCSession) SetPermission(_permission common.Address) (*types.Transaction, error) {
	return _AMC.Contract.SetPermission(&_AMC.TransactOpts, _permission)
}

// SetPermission is a paid mutator transaction binding the contract method 0xb85a35d2.
//
// Solidity: function setPermission(address _permission) returns()
func (_AMC *AMCTransactorSession) SetPermission(_permission common.Address) (*types.Transaction, error) {
	return _AMC.Contract.SetPermission(&_AMC.TransactOpts, _permission)
}

// SetRouter is a paid mutator transaction binding the contract method 0xc0d78655.
//
// Solidity: function setRouter(address _router) returns()
func (_AMC *AMCTransactor) SetRouter(opts *bind.TransactOpts, _router common.Address) (*types.Transaction, error) {
	return _AMC.contract.Transact(opts, "setRouter", _router)
}

// SetRouter is a paid mutator transaction binding the contract method 0xc0d78655.
//
// Solidity: function setRouter(address _router) returns()
func (_AMC *AMCSession) SetRouter(_router common.Address) (*types.Transaction, error) {
	return _AMC.Contract.SetRouter(&_AMC.TransactOpts, _router)
}

// SetRouter is a paid mutator transaction binding the contract method 0xc0d78655.
//
// Solidity: function setRouter(address _router) returns()
func (_AMC *AMCTransactorSession) SetRouter(_router common.Address) (*types.Transaction, error) {
	return _AMC.Contract.SetRouter(&_AMC.TransactOpts, _router)
}

// SetStandardFactory is a paid mutator transaction binding the contract method 0x005fa939.
//
// Solidity: function setStandardFactory(address _standardFactory) returns()
func (_AMC *AMCTransactor) SetStandardFactory(opts *bind.TransactOpts, _standardFactory common.Address) (*types.Transaction, error) {
	return _AMC.contract.Transact(opts, "setStandardFactory", _standardFactory)
}

// SetStandardFactory is a paid mutator transaction binding the contract method 0x005fa939.
//
// Solidity: function setStandardFactory(address _standardFactory) returns()
func (_AMC *AMCSession) SetStandardFactory(_standardFactory common.Address) (*types.Transaction, error) {
	return _AMC.Contract.SetStandardFactory(&_AMC.TransactOpts, _standardFactory)
}

// SetStandardFactory is a paid mutator transaction binding the contract method 0x005fa939.
//
// Solidity: function setStandardFactory(address _standardFactory) returns()
func (_AMC *AMCTransactorSession) SetStandardFactory(_standardFactory common.Address) (*types.Transaction, error) {
	return _AMC.Contract.SetStandardFactory(&_AMC.TransactOpts, _standardFactory)
}

// SetUPASFactory is a paid mutator transaction binding the contract method 0x39cd8e96.
//
// Solidity: function setUPASFactory(address _upasFactory) returns()
func (_AMC *AMCTransactor) SetUPASFactory(opts *bind.TransactOpts, _upasFactory common.Address) (*types.Transaction, error) {
	return _AMC.contract.Transact(opts, "setUPASFactory", _upasFactory)
}

// SetUPASFactory is a paid mutator transaction binding the contract method 0x39cd8e96.
//
// Solidity: function setUPASFactory(address _upasFactory) returns()
func (_AMC *AMCSession) SetUPASFactory(_upasFactory common.Address) (*types.Transaction, error) {
	return _AMC.Contract.SetUPASFactory(&_AMC.TransactOpts, _upasFactory)
}

// SetUPASFactory is a paid mutator transaction binding the contract method 0x39cd8e96.
//
// Solidity: function setUPASFactory(address _upasFactory) returns()
func (_AMC *AMCTransactorSession) SetUPASFactory(_upasFactory common.Address) (*types.Transaction, error) {
	return _AMC.Contract.SetUPASFactory(&_AMC.TransactOpts, _upasFactory)
}
