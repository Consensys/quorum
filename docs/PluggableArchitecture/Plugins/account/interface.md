<!-- This is auto generated file from running `go generate` in plugin/proto folder. Please do not edit -->



<a name="account.proto"></a>

## account.proto


### Services


<a name="proto.AccountService"></a>

#### `AccountService`


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Status | [StatusRequest](#proto.StatusRequest) | [StatusResponse](#proto.StatusResponse) | Status provides a string message describing the current state of the plugin. The content of the returned state is left to the developer's discretion. Examples include information on the business-level state (e.g. currently unlocked accounts) and process-level state (e.g. is plugin healthy) |
| Open | [OpenRequest](#proto.OpenRequest) | [OpenResponse](#proto.OpenResponse) | The specific behaviour of Open will depend on the backend being implemented. It may be that Open is not required in which case it should be implemented as a no-op. It should not unlock or decrypt account keys. Open may be used to initialize resources or create a connection to the storage backend. It should be expected that any allocated resources or connections can be released by calling Close. |
| Close | [CloseRequest](#proto.CloseRequest) | [CloseResponse](#proto.CloseResponse) | The specific behaviour of Close will depend on the backend being implemented. It may be that Close is not required, in which case it should be implemented as a no-op. Close releases any resources or connections allocated by Open. |
| Accounts | [AccountsRequest](#proto.AccountsRequest) | [AccountsResponse](#proto.AccountsResponse) | Accounts returns the currently available accounts managed by the plugin. |
| Contains | [ContainsRequest](#proto.ContainsRequest) | [ContainsResponse](#proto.ContainsResponse) | Contains returns whether the provided account is managed by the plugin. |
| Sign | [SignRequest](#proto.SignRequest) | [SignResponse](#proto.SignResponse) | Sign signs the provided data with the specified account |
| UnlockAndSign | [UnlockAndSignRequest](#proto.UnlockAndSignRequest) | [SignResponse](#proto.SignResponse) | UnlockAndSign unlocks the specified account with the provided passphrase and uses it to sign the provided data. The account will be locked once the signing is complete. It may be that the storage backend being implemented does not rely on passphrase-encryption, in which case the passphrase parameter should be ignored when unlocking. |
| TimedUnlock | [TimedUnlockRequest](#proto.TimedUnlockRequest) | [TimedUnlockResponse](#proto.TimedUnlockResponse) | TimedUnlock unlocks the specified account with the provided passphrase for the duration provided. The duration is provided in nanoseconds. It may be that the storage backend being implemented does not rely on passphrase-encryption, in which case the passphrase parameter should be ignored when unlocking. |
| Lock | [LockRequest](#proto.LockRequest) | [LockResponse](#proto.LockResponse) | Lock immediately locks the specified account, overriding any existing timed unlocks. |
| NewAccount | [NewAccountRequest](#proto.NewAccountRequest) | [NewAccountResponse](#proto.NewAccountResponse) | NewAccount creates a new account and stores it in the backend. The newAccountConfig is provided as a generic json-encoded byte array to allow for the structure of the config to be left to the developer's discretion. |
| ImportRawKey | [ImportRawKeyRequest](#proto.ImportRawKeyRequest) | [ImportRawKeyResponse](#proto.ImportRawKeyResponse) | ImportRawKey creates a new account from the provided hex-encoded private key and stores it in the backend. Validation of the hex string private key is not required as this handled by Quorum. The newAccountConfig is provided as a generic json-encoded byte array to allow for the structure of the config to be left to the developer's discretion. |

 <!-- end services -->

### Messsages


<a name="proto.Account"></a>

#### `Account`
Note: The Account type is used only in Response types.
All Request types which require an account to be specified use only the address to identify the account.
At the Quorum-level there is no knowledge of account URLs so the url field has been excluded from Requests to simplify the protocol.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [bytes](#bytes) |  | 20-byte ethereum address of the account |
| url | [string](#string) |  | URL for the stored account; format and structure will depend on the storage backend and is left to the developer's discretion |






<a name="proto.AccountsRequest"></a>

#### `AccountsRequest`







<a name="proto.AccountsResponse"></a>

#### `AccountsResponse`



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| accounts | [Account](#proto.Account) | repeated | list of accounts managed by the plugin |






<a name="proto.CloseRequest"></a>

#### `CloseRequest`







<a name="proto.CloseResponse"></a>

#### `CloseResponse`







<a name="proto.ContainsRequest"></a>

#### `ContainsRequest`



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [bytes](#bytes) |  | 20-byte ethereum address of the account to search for |






<a name="proto.ContainsResponse"></a>

#### `ContainsResponse`



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| isContained | [bool](#bool) |  | whether the account was found |






<a name="proto.ImportRawKeyRequest"></a>

#### `ImportRawKeyRequest`



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| rawKey | [string](#string) |  | hex-encoded private key to import and store in the storage backend |
| newAccountConfig | [bytes](#bytes) |  | json-encoded byte array providing the config necessary to create a new account; the config will be dependent on the storage backend and so its definition is left to the developer's discretion |






<a name="proto.ImportRawKeyResponse"></a>

#### `ImportRawKeyResponse`



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| account | [Account](#proto.Account) |  | the imported account |






<a name="proto.LockRequest"></a>

#### `LockRequest`



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [bytes](#bytes) |  | 20-byte ethereum address of the account to lock |






<a name="proto.LockResponse"></a>

#### `LockResponse`







<a name="proto.NewAccountRequest"></a>

#### `NewAccountRequest`



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| newAccountConfig | [bytes](#bytes) |  | json-encoded byte array providing the config necessary to create a new account; the config will be dependent on the storage backend and so its definition is left to the developer's discretion |






<a name="proto.NewAccountResponse"></a>

#### `NewAccountResponse`



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| account | [Account](#proto.Account) |  | the created account |






<a name="proto.OpenRequest"></a>

#### `OpenRequest`



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| passphrase | [string](#string) |  | passphrase required to open; may not be required by all storage backends, so its use is left to the developer's discretion |






<a name="proto.OpenResponse"></a>

#### `OpenResponse`







<a name="proto.SignRequest"></a>

#### `SignRequest`



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [bytes](#bytes) |  | 20-byte ethereum address of the account to use |
| toSign | [bytes](#bytes) |  | data to sign |






<a name="proto.SignResponse"></a>

#### `SignResponse`



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sig | [bytes](#bytes) |  | secp256k1 ECDSA signature in the 65-byte [R || S || V] format where V is 0 or 1 |






<a name="proto.StatusRequest"></a>

#### `StatusRequest`







<a name="proto.StatusResponse"></a>

#### `StatusResponse`



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [string](#string) |  | message describing the status of the plugin; message content is left to the developer's discretion |






<a name="proto.TimedUnlockRequest"></a>

#### `TimedUnlockRequest`



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [bytes](#bytes) |  | 20-byte ethereum address of the account to unlock |
| password | [string](#string) |  | passphrase required to unlock the account; may not be required by all storage backends, so its use is left to the developer's discretion |
| duration | [int64](#int64) |  | number of nanoseconds the account should be unlocked for |






<a name="proto.TimedUnlockResponse"></a>

#### `TimedUnlockResponse`







<a name="proto.UnlockAndSignRequest"></a>

#### `UnlockAndSignRequest`



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [bytes](#bytes) |  | 20-byte ethereum address of the account to use |
| toSign | [bytes](#bytes) |  | data to sign |
| passphrase | [string](#string) |  | passphrase required to unlock the account; may not be required by all storage backends, so its use is left to the developer's discretion |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

