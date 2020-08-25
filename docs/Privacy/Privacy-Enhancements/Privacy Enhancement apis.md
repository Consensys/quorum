# Privacy Enhancement  APIs
## APIs
### `getContractPrivacyMetadata` 
The api to query the Privacy Metadata for a contract account address
#### Parameter
* contract address

#### Returns
* `creationTxHash`: ACOTH
* `privacyFlag`: unsigned integer with `1` for PP, `3` for PSV contracts (default value is `0`).

#### Examples

```jshelllanguage tab="JSON RPC"
// Request
curl -X POST http://127.0.0.1:22000 --data '{"jsonrpc":"2.0","method":"eth_getContractPrivacyMetadata","params":["0x1932c48b2bf8102ba33b4a6b545c32236e342f34"],"id":15}' --header "Content-Type: application/json"

// Response
{"jsonrpc":"2.0","id":15,"result":{"creationTxHash": [246,124,116,139,190,217,33,16,203,102,81,13,65,58,249,145,68,180,67,79,163,37,119,27,99,35,247,240,12,53,25,45,47,134,16,118,246,128,97,237,45,50,79,97,78,221,47,1 89,11,165,238,36,8,187,66,64,42,135,108,75,41,85,152,183],"privacyFlag":3}}
```

```javascript tab="geth console"
> eth.getContractPrivacyMetadata("0x1932c48b2bf8102ba33b4a6b545c32236e342f34");

{
  creationTxHash: [246,124,116,139,190,217,33,16,203,102,81,13,65,58,249,145,68,180,67,79,163,37,119,27,99,35,247,240,12,53,25,45,47,134,16,118,246,128,97,237,45,50,79,97,78,221,47,1 89,11,165,238,36,8,187,66,64,42,135,108,75,41,85,152,183],
  privacyFlag: 3
}

```
