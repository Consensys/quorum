//
// Create a public contract
//

a = eth.accounts[0]
web3.eth.defaultAccount = a;

var simpleSource = 'contract simplestorage { uint public storedData; function simplestorage(uint initVal) { storedData = initVal; } function set(uint x) { storedData = x; } function get() constant returns (uint retVal) { return storedData; } }'
var simpleCompiled = web3.eth.compile.solidity(simpleSource);
var simpleRoot = Object.keys(simpleCompiled)[0];
var simpleContract = web3.eth.contract(simpleCompiled[simpleRoot].info.abiDefinition);
var simple = simpleContract.new(42, {from:web3.eth.accounts[0], data: simpleCompiled[simpleRoot].code, gas: 300000}, function(e, contract) {
	if (e) {
		console.log("err creating contract", e);
	} else {
		if (!contract.address) {
			console.log("Contract transaction send: TransactionHash: " + contract.transactionHash + " waiting to be mined...");
		} else {
			console.log("Contract mined! Address: " + contract.address);
			console.log(contract);
		}
	}
});
